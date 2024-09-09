package gqlfiltergen

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/ast"
)

type AstAndTemplateData struct {
	Ast      *ast.Definition
	TypeData *TypeData
}

func generateMainFilterDefinition(ot map[string]ast.FieldList) map[string]*AstAndTemplateData {
	out := make(map[string]*AstAndTemplateData)
	seen := make(map[string]bool)
	for objectName := range ot {
		generateMainFilterDefinitionLoop(ot, out, seen, false, objectName)
	}

	return out
}

func generateMainFilterDefinitionLoop(ot map[string]ast.FieldList, processed map[string]*AstAndTemplateData, seen map[string]bool, nested bool, objectName string) string {
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

	fields, ok := ot[objectName]
	if !ok {
		field := processField(objectName, objectName)
		return field.Type.NamedType
	}

	for _, f := range fields {
		fd := &ast.FieldDefinition{
			Description: fmt.Sprintf("filter for %s field.", f.Name),
			Name:        f.Name,
		}

		var isSlice bool
		var isSliceElemPointer bool
		var isNested bool

		field := processField(f.Name, f.Type.NamedType)
		switch {
		case field != nil:
			fd = field
		case f.Type.NamedType == "": // it is a list
			fd.Type = ast.NamedType(generateMainFilterDefinitionLoop(ot, processed, seen, true, f.Type.Elem.NamedType), nil)
			isSliceElemPointer = !f.Type.Elem.NonNull
			isSlice = true
		default: // It is a named custom named type
			isNested = true
			fd.Type = ast.NamedType(generateMainFilterDefinitionLoop(ot, processed, seen, true, f.Type.NamedType), nil)
		}

		objDef.Fields = append(objDef.Fields, fd)

		tf := &FieldMapping{
			FilterField:        f.Name,
			TypeName:           fd.Type.Name(),
			IsSlice:            isSlice,
			IsNested:           isNested,
			IsPointer:          !f.Type.NonNull,
			IsSliceElemPointer: isSliceElemPointer,
		}

		typeData.Fields = append(typeData.Fields, tf)
	}

	processed[filterName] = &AstAndTemplateData{
		Ast:      objDef,
		TypeData: typeData,
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
		fd.Type = ast.NamedType(filterNumberName, nil)
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
		Kind:        ast.Object,
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
	filterNumberName  = "FilterNumber"
	filterTimeName    = "FilterTime"
	filterBooleanName = "FilterBoolean"
)

var filterString = &ast.Definition{
	Kind:        ast.Object,
	Description: "Filter type for string fields. It contains a variety of filter types for string types. All added filters here are processed as AND operators.",
	Name:        filterStringName,
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

var filterNumber = &ast.Definition{
	Kind:        ast.Object,
	Description: "Filter type for number fields. All added filters here are processed as AND operators.",
	Name:        filterNumberName,
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

var filterTime = &ast.Definition{
	Kind:        ast.Object,
	Description: "Filter type for time fields. All added filters here are processed as AND operators.",
	Name:        filterTimeName,
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

var filterBoolean = &ast.Definition{
	Kind:        ast.Object,
	Description: "Filter type for boolean fields. All added filters here are processed as AND operators.",
	Name:        filterBooleanName,
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
