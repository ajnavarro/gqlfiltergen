package gqlfiltergen

// FieldMapping defines field mapping struct to pass data to the template
type FieldMapping struct {
	Pointer     bool   // Set if the field will be a pointer or not
	FilterField string // Name of the filter field in FilterTypeOne
	IsSlice     bool   // Whether the field is a slice or not
	IsNested    bool   // Whether the field is a nested struct or not
}

func (fm *FieldMapping) NeedMemAddr() string {
	if !fm.Pointer && !fm.IsNested {
		return "&"
	}

	return ""
}

// TypeData holds the data to be passed into the template for every type
type TypeData struct {
	TypeName   string
	FilterName string
	Fields     []*FieldMapping
}

type TemplateData struct {
	TypeDatas []*TypeData
}
