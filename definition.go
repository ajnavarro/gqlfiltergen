package gqlfiltergen

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/ast"
)

type ProcessingField struct {
	Field         *ast.FieldDefinition
	IsMinmaxeable bool
}

type ProcessingObject struct {
	Fields     []*ProcessingField
	Definition *ast.Definition
}

type AstAndTemplateData struct {
	Ast      *ast.Definition
	TypeData *TypeData
}

func generateMainFilterDefinition(ot map[string]*ProcessingObject) map[string]*AstAndTemplateData {
	out := make(map[string]*AstAndTemplateData)
	seen := make(map[string]bool)
	for objectName := range ot {
		generateMainFilterDefinitionLoop(ot, out, seen, false, objectName)
	}

	return out
}

func generateMainFilterDefinitionLoop(ot map[string]*ProcessingObject, processed map[string]*AstAndTemplateData, seen map[string]bool, nested bool, objectName string) string {
	field := processField(objectName, objectName)
	if field != nil {
		return field.Type.NamedType
	}

	filterName := fmt.Sprintf("Filter%s", objectName)
	if nested {
		filterName = fmt.Sprintf("NestedFilter%s", objectName)
	}

	if ok := seen[filterName]; ok {
		fmt.Println("filter already on the loop:", filterName)
		return filterName
	}

	seen[filterName] = true

	objDef := filterDefinition(filterName, objectName)

	typeData := &TypeData{
		TypeName:   objectName,
		FilterName: filterName,
	}

	pfs, ok := ot[objectName]
	if !ok {
		panic(fmt.Errorf("error creating field for type %q. Check that you added some fields as @filterable", objectName))
	}

	switch pfs.Definition.Kind {
	case ast.Union:
		isFiltered := false
		for _, t := range pfs.Definition.Types {
			if _, ok := ot[t]; !ok {
				fmt.Println("type in an union with no filters:", t)
				continue
			}

			isFiltered = true

			fd := &ast.FieldDefinition{
				Description: fmt.Sprintf("filter for %s union type.", t),
				Name:        t,
				Type:        ast.NamedType(generateMainFilterDefinitionLoop(ot, processed, seen, true, t), nil),
			}

			objDef.Fields = append(objDef.Fields, fd)

			tf := &FieldMapping{
				Field:       t,
				FilterField: t,
				TypeName:    fd.Type.Name(),
			}

			typeData.Fields = append(typeData.Fields, tf)
		}

		typeData.IsUnion = true

		if isFiltered {
			processed[filterName] = &AstAndTemplateData{
				Ast:      objDef,
				TypeData: typeData,
			}
		}

		return filterName
	case ast.Object:
		for _, pf := range pfs.Fields {
			f := pf.Field

			if f.Type.Name() != "Int" && pf.IsMinmaxeable {
				panic(fmt.Errorf("only Int types can be minmaxeables. Field: %s Type: %s", f.Name, f.Type.Name()))
			}

			fd := &ast.FieldDefinition{
				Description: fmt.Sprintf("filter for %s field.", f.Name),
				Name:        f.Name,
			}

			var isSlice bool
			var isSliceElemPointer bool
			var isNested bool
			var isUnion bool

			if possibleUnion, ok := ot[f.Type.NamedType]; ok {
				isUnion = possibleUnion.Definition.Kind == ast.Union
			}

			field := processField(f.Name, f.Type.NamedType)
			switch {
			case field != nil:
				fd = field
			case f.Type.NamedType == "": // it is a list
				fd.Type = ast.NamedType(generateMainFilterDefinitionLoop(ot, processed, seen, true, f.Type.Elem.NamedType), nil)
				isSliceElemPointer = !f.Type.Elem.NonNull
				isSlice = true
			case isUnion:
				fd.Type = ast.NamedType(generateMainFilterDefinitionLoop(ot, processed, seen, true, f.Type.NamedType), nil)
			default: // It is a named custom named type
				isNested = true
				fd.Type = ast.NamedType(generateMainFilterDefinitionLoop(ot, processed, seen, true, f.Type.NamedType), nil)
			}

			objDef.Fields = append(objDef.Fields, fd)

			tf := &FieldMapping{
				Field:              f.Name,
				TypeName:           fd.Type.Name(),
				IsSlice:            isSlice,
				IsNested:           isNested,
				IsPointer:          !f.Type.NonNull && !isUnion,
				IsSliceElemPointer: isSliceElemPointer,

				IsMinmaxeable: pf.IsMinmaxeable,
			}

			typeData.Fields = append(typeData.Fields, tf)
		}

		processed[filterName] = &AstAndTemplateData{
			Ast:      objDef,
			TypeData: typeData,
		}
	default:
		panic(fmt.Errorf("unsupported type: %q", pfs.Definition.Kind))
	}

	return filterName
}

func processField(name, typeName string) *ast.FieldDefinition {
	fd := &ast.FieldDefinition{
		Description: fmt.Sprintf("filter for %s field.", name),
		Name:        name,
	}
	switch typeName {
	case "String":
		fd.Type = ast.NamedType(filterStringName, nil)
	case "Int":
		fd.Type = ast.NamedType(filterIntName, nil)
	case "Time":
		fd.Type = ast.NamedType(filterTimeName, nil)
	case "Boolean":
		fd.Type = ast.NamedType(filterBooleanName, nil)
	default:
		return nil
	}

	return fd
}

func filterDefinition(filterName, objectName string) *ast.Definition {
	return &ast.Definition{
		Kind:        ast.InputObject,
		Name:        filterName,
		Description: fmt.Sprintf("filter for %s objects", objectName),
		Fields: ast.FieldList{
			&ast.FieldDefinition{
				Type:        ast.ListType(ast.NamedType(filterName, nil), nil),
				Name:        "_and",
				Description: fmt.Sprintf("logical operator for %s that will combine two or more conditions, returning true if all of them are true.", objectName),
			},
			&ast.FieldDefinition{
				Type:        ast.ListType(ast.NamedType(filterName, nil), nil),
				Name:        "_or",
				Description: fmt.Sprintf("logical operator for %s that will combine two or more conditions, returning true if at least one of them is true.", objectName),
			},
			&ast.FieldDefinition{
				Type:        ast.NamedType(filterName, nil),
				Name:        "_not",
				Description: fmt.Sprintf("logical operator for %s that will reverse conditions.", objectName),
			},
		},
	}
}

const (
	filterStringName  = "FilterString"
	filterIntName     = "FilterInt"
	filterTimeName    = "FilterTime"
	filterBooleanName = "FilterBoolean"
)

func filterString(name string) *ast.Definition {
	return &ast.Definition{
		Kind:        ast.InputObject,
		Description: "Filter type for string fields. It contains a variety of filter types for string types. All added filters here are processed as AND operators.",
		Name:        name,
		Fields: ast.FieldList{
			&ast.FieldDefinition{
				Description: "Filter a string field checking if it exists or not.",
				Name:        "exists",
				Type:        ast.NamedType("Boolean", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a string field checking if it is equals to the specified value.",
				Name:        "eq",
				Type:        ast.NamedType("String", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a string field checking if it is NOT equals to the specified value.",
				Name:        "neq",
				Type:        ast.NamedType("String", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a string field checking if it is like the specified value. You can use standard Go RegEx expressions here.",
				Name:        "like",
				Type:        ast.NamedType("String", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a string field checking if it is NOT like the specified value. You can use standard Go RegEx expressions here.",
				Name:        "nlike",
				Type:        ast.NamedType("String", nil),
			},
		},
	}

}

func filterNumber(name string) *ast.Definition {
	return &ast.Definition{
		Kind:        ast.InputObject,
		Description: "Filter type for number fields. All added filters here are processed as AND operators.",
		Name:        name,
		Fields: ast.FieldList{
			&ast.FieldDefinition{
				Description: "Filter a number field checking if it exists or not.",
				Name:        "exists",
				Type:        ast.NamedType("Boolean", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a number field checking if it is equals to the specified value.",
				Name:        "eq",
				Type:        ast.NamedType("Int", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a number field checking if it is NOT equals to the specified value.",
				Name:        "neq",
				Type:        ast.NamedType("Int", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a number field checking if it is greater than the specified value.",
				Name:        "gt",
				Type:        ast.NamedType("Int", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a number field checking if it is less than the specified value.",
				Name:        "lt",
				Type:        ast.NamedType("Int", nil),
			},
		},
	}

}

func filterTime(name string) *ast.Definition {
	return &ast.Definition{
		Kind:        ast.InputObject,
		Description: "Filter type for time fields. All added filters here are processed as AND operators.",
		Name:        name,
		Fields: ast.FieldList{
			&ast.FieldDefinition{
				Description: "Filter a time field checking if it exists or not.",
				Name:        "exists",
				Type:        ast.NamedType("Boolean", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a time field checking if it is equals to the specified value.",
				Name:        "eq",
				Type:        ast.NamedType("Time", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a time field checking if it is NOT equals to the specified value.",
				Name:        "neq",
				Type:        ast.NamedType("Time", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a time field checking if it is before than the specified value.",
				Name:        "before",
				Type:        ast.NamedType("Time", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a time field checking if it is after the specified value.",
				Name:        "after",
				Type:        ast.NamedType("Time", nil),
			},
		},
	}

}

func filterBoolean(name string) *ast.Definition {
	return &ast.Definition{
		Kind:        ast.InputObject,
		Description: "Filter type for boolean fields. All added filters here are processed as AND operators.",
		Name:        name,
		Fields: ast.FieldList{
			&ast.FieldDefinition{
				Description: "Filter a boolean field checking if it exists or not.",
				Name:        "exists",
				Type:        ast.NamedType("Boolean", nil),
			},
			&ast.FieldDefinition{
				Description: "Filter a boolean field checking if it is equals to the specified value.",
				Name:        "eq",
				Type:        ast.NamedType("Boolean", nil),
			},
		},
	}
}
