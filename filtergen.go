package gqlfiltergen

import (
	"bytes"
	_ "embed"
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"github.com/vektah/gqlparser/v2/parser"
)

const (
	filterableDirectiveName = "filterable"
	extrasArgumentName      = "extras"
	extraMinmaxName         = "MINMAX"
)

//go:embed filter.functions.go.tpl
var functionsTpl string

var _ plugin.EarlySourceInjector = &Plugin{}
var _ plugin.LateSourceInjector = &Plugin{}
var _ plugin.CodeGenerator = &Plugin{}
var _ plugin.ConfigMutator = &Plugin{}

type Plugin struct {
	templateData *TemplateData
	opts         *Options
}

type Options struct {
	InjectCodeAfter string
}

func NewPlugin(opts *Options) *Plugin {
	return &Plugin{
		opts: opts,
	}
}

func (f *Plugin) Name() string {
	return "filtergen"
}

func (f *Plugin) InjectSourceEarly() *ast.Source {
	return &ast.Source{
		Name: "filtergen.directives.graphql",
		Input: `
enum FilterableExtra {
  """
  Get minimum and maximum value used on all the filters for this field.
  Useful when you need to do a range query for performance reasons.
  """
  MINMAX
}

directive @filterable(
  """
  Add extra functionality to this field apart from the filtering capabilities.
  """
  extras: [FilterableExtra!]
) on FIELD_DEFINITION
`,
	}
}

func (f *Plugin) InjectSourceLate(schema *ast.Schema) *ast.Source {
	processingTypes := make(map[string]*ProcessingObject)
	for n, t := range schema.Types {
		if _, ok := processingTypes[n]; ok {
			continue
		}
		if t.Kind == ast.Union {
			// add unions just in case they have types with filters
			processingTypes[n] = &ProcessingObject{
				Definition: t,
			}

			continue
		}

		for _, f := range t.Fields {
			filterable, minmaxeable := getDirectives(f.Directives)
			if !filterable {
				continue
			}
			fl := processingTypes[n]
			if fl == nil {
				fl = &ProcessingObject{}
			}

			fl.Fields = append(fl.Fields, &ProcessingField{
				Field:         f,
				IsMinmaxeable: minmaxeable,
			})

			fl.Definition = t

			processingTypes[n] = fl
		}
	}

	initTypes := map[string]*ast.Definition{
		filterStringName:        filterString(filterStringName),
		nestedFilterStringName:  filterString(nestedFilterStringName),
		filterNumberName:        filterNumber(filterNumberName),
		nestedFilterNumberName:  filterNumber(nestedFilterNumberName),
		filterTimeName:          filterTime(filterTimeName),
		nestedFilterTimeName:    filterTime(nestedFilterTimeName),
		filterBooleanName:       filterBoolean(filterBooleanName),
		nestedFilterBooleanName: filterBoolean(nestedFilterBooleanName),
	}

	outSchema := &ast.Schema{
		Types: initTypes,
	}

	defMap := generateMainFilterDefinition(processingTypes)

	td := &TemplateData{}
	for k, v := range defMap {
		outSchema.Types[k] = v.Ast

		td.TypeDatas = append(td.TypeDatas, v.TypeData)
	}

	f.templateData = td

	var buf bytes.Buffer
	formatter := formatter.NewFormatter(&buf, formatter.WithComments())

	formatter.FormatSchema(outSchema)

	return &ast.Source{
		Name:    "filtergen.graphql",
		Input:   buf.String(),
		BuiltIn: false,
	}
}

func getDirectives(cd ast.DirectiveList) (filterable bool, minmaxeable bool) {
	if cd == nil {
		return
	}

	for _, d := range cd {
		if d.Name == filterableDirectiveName {
			filterable = true
		}

		if a := d.Arguments.ForName(extrasArgumentName); a != nil {
			if a.Value.Kind != ast.ListValue {
				continue
			}
			if ch := a.Value.Children; len(ch) != 0 {
				if vals := ch.ForName(""); vals != nil {
					if strings.Contains(vals.Raw, extraMinmaxName) {
						minmaxeable = true
					}
				}
			}
		}
	}

	return
}

func (f *Plugin) MutateConfig(cfg *config.Config) error {
	if err := f.injectUserQueries(cfg); err != nil {
		return err
	}

	return cfg.LoadSchema()
}

func (f *Plugin) GenerateCode(data *codegen.Data) error {
	// rebuild template data with Go field names
	for _, t := range f.templateData.TypeDatas {
		dataObj := data.Objects.ByName(t.TypeName)
		if dataObj == nil {
			continue
		}

		for _, f := range t.Fields {
			for _, of := range dataObj.Fields {
				if f.Field == of.FieldDefinition.Name {
					f.FilterField = of.GoFieldName
					f.IsMethod = of.IsMethod()
				}
			}
		}

		sort.Slice(t.Fields, func(i, j int) bool {
			return t.Fields[i].FilterField > t.Fields[j].FilterField
		})
	}

	sort.Slice(f.templateData.TypeDatas, func(i, j int) bool {
		return f.templateData.TypeDatas[i].FilterName > f.templateData.TypeDatas[j].FilterName
	})

	filename := path.Join(path.Dir(data.Config.Model.Filename), "filter_methods_gen.go")

	return templates.Render(templates.Options{
		PackageName:     data.Config.Model.Package,
		Filename:        filename,
		Data:            f.templateData,
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Template:        functionsTpl,
	})
}

const (
	queryName        = "Query"
	subscriptionName = "Subscription"
)

func (f *Plugin) injectUserQueries(cfg *config.Config) error {
	orig := cfg.Schema

	source := &ast.Source{
		Input: f.opts.InjectCodeAfter,
	}

	schema, err := parser.ParseSchema(source)
	if err != nil {
		return fmt.Errorf("error injecting user-defined queries: %w", err)
	}

	if qt := schema.Definitions.ForName(queryName); qt != nil {
		if oqt, ok := orig.Types[queryName]; ok {
			for _, f := range qt.Fields {
				oqt.Fields = append(oqt.Fields, f)
			}
		} else {
			orig.Types[queryName] = qt
		}
	}

	if st := schema.Definitions.ForName(subscriptionName); st != nil {
		if ost, ok := orig.Types[subscriptionName]; ok {
			for _, f := range st.Fields {
				ost.Fields = append(ost.Fields, f)
			}
		} else {
			orig.Types[subscriptionName] = st
		}
	}

	var buf bytes.Buffer
	form := formatter.NewFormatter(&buf, formatter.WithComments())
	form.FormatSchema(orig)

	// TODO: might be a better way to do it
	// we need to overwrite sources to correctly generate resolvers for the queries and subscriptions that we injected
	cfg.Sources = []*ast.Source{
		{
			Name:  "all",
			Input: buf.String(),
		},
	}

	return nil
}
