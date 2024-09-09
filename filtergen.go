package gqlfiltergen

import (
	"bytes"
	_ "embed"
	"path"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
)

const filterableDirectiveName = "filterable"

//go:embed filter.functions.go.tpl
var functionsTpl string

var _ plugin.EarlySourceInjector = &Plugin{}
var _ plugin.LateSourceInjector = &Plugin{}
var _ plugin.CodeGenerator = &Plugin{}

type Plugin struct {
	templateData *TemplateData
}

func (f *Plugin) Name() string {
	return "filtergen"
}

func (f *Plugin) InjectSourceEarly() *ast.Source {
	return &ast.Source{
		Name:  "filtergen.directives.graphql",
		Input: "directive @filterable on FIELD_DEFINITION",
	}
}

func (f *Plugin) InjectSourceLate(schema *ast.Schema) *ast.Source {
	processingTypes := make(map[string]ast.FieldList)
	for n, t := range schema.Types {
		if _, ok := processingTypes[n]; ok {
			continue
		}
		for _, f := range t.Fields {
			if !sholudProcess(f.Directives) {
				continue
			}
			fl := processingTypes[n]

			if fl == nil {
				fl = ast.FieldList{}
			}

			fl = append(fl, f)

			processingTypes[n] = fl
		}
	}

	initTypes := map[string]*ast.Definition{
		filterString.Name:  filterString,
		filterNumber.Name:  filterNumber,
		filterTime.Name:    filterTime,
		filterBoolean.Name: filterBoolean,
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

func sholudProcess(cd ast.DirectiveList) bool {
	if cd == nil {
		return false
	}
	for _, d := range cd {
		if d.Name == filterableDirectiveName {
			return true
		}
	}

	return false
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
	}

	filename := path.Join(path.Dir(data.Config.Model.Filename), "filter_methods.go")

	return templates.Render(templates.Options{
		PackageName:     data.Config.Model.Package,
		Filename:        filename,
		Data:            f.templateData,
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Template:        functionsTpl,
	})
}
