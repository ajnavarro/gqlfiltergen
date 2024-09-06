
{{ reserveImport "regexp" }}
{{ reserveImport "time" }}

///////////////////////////////// CUSTOM  TYPES /////////////////////////////////

{{- range .TypeDatas }}

	func (f *{{.FilterName}}) Eval(obj *{{.TypeName}}) bool {
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

	// Evaluate individual field filters
	{{- range .Fields }}
		{{if .IsSlice}}
		// Handle {{.FilterField}} slice
		if f.{{.FilterField}} != nil {
			for _, elem := range obj.{{.FilterField}} {
				if !f.{{.FilterField}}.Eval(elem) {
					return false
				}
			}
		}		
		{{else}}
		// Handle {{.FilterField}} field
		if f.{{.FilterField}} != nil && !f.{{.FilterField}}.Eval({{.NeedMemAddr}}obj.{{.FilterField}}) {
			return false
		}
		{{end}}
	{{- end }}

	return true
}
{{- end }}

///////////////////////////////// GENERIC TYPES /////////////////////////////////

func (f *FilterBoolean) Eval(val *bool) bool {
	if f == nil {
		return true
	}

	return rootEval(val, f.Exists, f.Eq, nil)
}

func (f *FilterNumber) Eval(val *int) bool {
	if f == nil {
		return true
	}

	if !rootEval(val, f.Exists, f.Eq, f.Neq) {
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

	if !rootEval(val, f.Exists, f.Eq, f.Neq) {
		return false
	}

	if val != nil && f.Like != nil {
		matched, err := regexp.MatchString(*f.Like, *val)
		if err != nil || !matched {
			return false
		}
	}

	if val != nil && f.Nlike != nil {
		matched, err := regexp.MatchString(*f.Nlike, *val)
		if err != nil || matched {
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

	if !rootEval(val, f.Exists, f.Eq, f.Neq) {
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
func rootEval[T comparable](val *T, exists *bool, eq *T, neq *T) bool {
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

	// Check the Neq filter
	if neq != nil && *neq == *val {
		return false
	}

	return true
}