
{{ reserveImport "regexp" }}
{{ reserveImport "time" }}
{{ reserveImport "math" }}

///////////////////////////////// CUSTOM  TYPES /////////////////////////////////

{{range $td := .TypeDatas }}

	func (f *{{$td.FilterName}}) Eval(obj *{{$td.TypeName}}) bool {
	// Evaluate logical operators first
	if len(f.And) > 0 {
		for _, subFilter := range f.And {
			if !subFilter.Eval(obj) {
				return false
			}
		}
	}

	if len(f.Or) > 0 {
		orResult := false
		for _, subFilter := range f.Or {
			if subFilter.Eval(obj) {
				orResult = true
				break
			}
		}
		if !orResult {
			return false
		}
	}

	if f.Not != nil {
		if f.Not.Eval(obj) {
			return false
		}
	}

{{if $td.IsUnion}}
	// Handle union objects depending of the type

	// Check if any filters are specified
	filtersSpecified := {{- range $td.Fields }} f.{{.FilterField}} != nil || {{end}} false

    // If no filters are specified for any types, accept all objects
    if !filtersSpecified {
        return true
    }

 	// Evaluate specified type filters
    matchedType := false

	tobj := *obj
	{{- range $td.Fields }}
	if uObj, ok := tobj.({{.FilterField}}); ok {
		matchedType = true
		if f.{{.FilterField}} != nil && f.{{.FilterField}}.Eval(&uObj) {
			return true
		}
	}
	if uObj, ok := tobj.(*{{.FilterField}}); ok {
		matchedType = true
		if f.{{.FilterField}} != nil && f.{{.FilterField}}.Eval(uObj) {
			return true
		}
	}
	{{end}}
	// If the object is of a type specified in filters but didn't match, return false.
    if matchedType {
        return false
    }
{{else}}
	// Evaluate individual field filters
	{{- range $td.Fields }}
		{{if .IsSlice}}
		// Handle {{.FilterField}} slice
		if f.{{.FilterField}} != nil {
			elemMatch{{.FilterField}} := false
			for _, elem := range obj.{{.CallWrapping .FilterField}} {
				if f.{{.FilterField}}.Eval({{.EvalCallWrapping "elem"}}) {
					elemMatch{{.FilterField}} = true
				}
			}

			if !elemMatch{{.FilterField}} {
				return false
			}

		}		
		{{else}}
		// Handle {{.FilterField}} field
		toEval{{.FilterField}} := {{.EvalVarWrapping (printf "obj.%s" .FilterField)}}
		if f.{{.FilterField}} != nil && !f.{{.FilterField}}.Eval({{.EvalCallWrapping (printf "toEval%s" .FilterField)}}) {
			return false
		}
		{{end}}
	{{- end }}
{{end}}

	return true
}

{{- range $td.Fields }}
	{{if .IsMinmaxeable}}
		// MinMax function for {{.FilterField}}
		func (f *{{$td.FilterName}}) MinMax{{.FilterField}}() (min *int, max *int) {
			// Recursively handle And conditions
			if len(f.And) > 0 {
				for _, subFilter := range f.And {
					subMin, subMax := subFilter.MinMax{{.FilterField}}()
					if subMin != nil && (min == nil || *subMin < *min) {
						min = subMin
					}
					if subMax != nil && (max == nil || *subMax > *max) {
						max = subMax
					}
				}
			}

			// Recursively handle Or conditions
			if len(f.Or) > 0 {
				for _, subFilter := range f.Or {
					subMin, subMax := subFilter.MinMax{{.FilterField}}()
					if subMin != nil && (min == nil || *subMin < *min) {
						min = subMin
					}
					if subMax != nil && (max == nil || *subMax > *max) {
						max = subMax
					}
				}
			}

			if f.{{.FilterField}} != nil {
				if f.{{.FilterField}}.Gt != nil {
					if min == nil || *f.{{.FilterField}}.Gt < *min {
						min = f.{{.FilterField}}.Gt
					}
					if max == nil || *f.{{.FilterField}}.Gt > *max {
						max = f.{{.FilterField}}.Gt
					}
				}

				if f.{{.FilterField}}.Lt != nil {
					if min == nil || *f.{{.FilterField}}.Lt < *min {
						min = f.{{.FilterField}}.Lt
					}
					if max == nil || *f.{{.FilterField}}.Lt > *max {
						max = f.{{.FilterField}}.Lt
					}
				}

				if f.{{.FilterField}}.Eq != nil {
					if min == nil || *f.{{.FilterField}}.Eq < *min {
						min = f.{{.FilterField}}.Eq
					}
					if max == nil || *f.{{.FilterField}}.Eq > *max {
						max = f.{{.FilterField}}.Eq
					}
				}
			}

			return min, max
		}
	{{end}}
{{- end}}

{{- end }}

func toIntPtr(val interface{}) *int {
	if val == nil {
		return nil
	}

	switch v := val.(type) {
	case int:
		return &v
	case int64:
		i := int(v)
		return &i
	case int32:
		i := int(v)
		return &i
	case int16:
		i := int(v)
		return &i
	case int8:
		i := int(v)
		return &i
	case *int:
		return v
	case *int64:
		i := int(*v)
		return &i
	case *int32:
		i := int(*v)
		return &i
	case *int16:
		i := int(*v)
		return &i
	case *int8:
		i := int(*v)
		return &i
	default:
		return nil
	}
}

///////////////////////////////// GENERIC TYPES /////////////////////////////////

func (f *FilterBoolean) Eval(val *bool) bool {
	if f == nil {
		return true
	}

	return rootEval(val, f.Exists, f.Eq)
}

func (f *FilterInt) Eval(val *int) bool {
	if f == nil {
		return true
	}

	if !rootEval(val, f.Exists, f.Eq) {
		return false
	}

	if val != nil && f.Gt != nil && *val <= *f.Gt {
		return false
	}

	if val != nil && f.Lt != nil && *val >= *f.Lt {
		return false
	}

	return true
}

func (f *FilterString) Eval(val *string) bool {
	if f == nil {
		return true
	}

	if !rootEval(val, f.Exists, f.Eq) {
		return false
	}

	if val != nil && f.Like != nil {
		matched, err := regexp.MatchString(*f.Like, *val)
		if err != nil || !matched {
			return false
		}
	}

	return true
}

// Eval evaluates the FilterTime conditions against a given time.Time value
func (f *FilterTime) Eval(val *time.Time) bool {
	if f == nil {
		return true
	}

	if !rootEval(val, f.Exists, f.Eq) {
		return false
	}

	// Check if the value is before the specified time
	if f.Before != nil && !val.Before(*f.Before) {
		return false
	}

	// Check if the value is after the specified time
	if f.After != nil && !val.After(*f.After) {
		return false
	}

	return true
}

// rootEval is a generic function that checks if the provided value matches the filter conditions.
func rootEval[T comparable](val *T, exists *bool, eq *T) bool {
	// Check the Exists filter
	if exists != nil {
		if *exists && val == nil {
			return false
		}
		if !*exists && val != nil {
			return false
		}
	}

	// If val is nil and we reach this point, skip the following checks
	if val == nil {
		return true
	}

	// Check the Eq filter
	if eq != nil && *eq != *val {
		return false
	}

	return true
}