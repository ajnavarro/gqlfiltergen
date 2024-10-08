package gqlfiltergen

import "fmt"

const (
	toIntPtrFunction = "toIntPtr"
)

// FieldMapping defines field mapping struct to pass data to the template
type FieldMapping struct {
	FilterField string // Name of the field from the struct
	Field       string // Name of the field from the schema
	TypeName    string // Type name

	IsPointer        bool // Set if the field will be a pointer or not
	IsSlice          bool // Whether the field is a slice or not
	IsSliceBasicType bool
	IsNested         bool // Whether the field is a nested struct or not
	IsMethod         bool // Check if we need to call a method instead of a var
	IsUnion          bool

	IsMinmaxeable bool
}

func (fm *FieldMapping) EvalVarWrapping(code string) string {
	if fm.IsMethod {
		code += "()"
	}

	if fm.TypeName == filterIntName {
		return fmt.Sprintf("%s(%s)", toIntPtrFunction, code)
	}

	return code
}

func (fm *FieldMapping) EvalCallWrapping(code string) string {
	if fm.IsUnion { // unions are defined by interfaces, so we need a pointer always
		return fmt.Sprintf("%s%s", "&", code)
	}

	if !fm.IsPointer && !fm.IsNested && !fm.IsSlice && fm.TypeName != filterIntName {
		return fmt.Sprintf("%s%s", "&", code)
	}

	if fm.IsSlice && fm.IsSliceBasicType { // Case for non pointers on slices for primitive types
		return fmt.Sprintf("%s%s", "&", code)
	}

	return code
}

func (fm *FieldMapping) CallWrapping(field string) string {
	funct := ""
	if fm.IsMethod {
		funct = "()"
	}

	return fmt.Sprintf("%s%s", field, funct)
}

// TypeData holds the data to be passed into the template for every type
type TypeData struct {
	TypeName   string
	FilterName string
	IsUnion    bool // Whether the type is coming from an union or not
	Fields     []*FieldMapping
}

type TemplateData struct {
	TypeDatas []*TypeData
}
