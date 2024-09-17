// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package out

import (
	"regexp"
	"time"
)

///////////////////////////////// CUSTOM  TYPES /////////////////////////////////

func (f *NestedFilterUnionTypeTwo) Eval(obj *UnionTypeTwo) bool {
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

	// Handle TypeTimeUnionTwo field
	toEvalTypeTimeUnionTwo := obj.TypeTimeUnionTwo
	if f.TypeTimeUnionTwo != nil && !f.TypeTimeUnionTwo.Eval(toEvalTypeTimeUnionTwo) {
		return false
	}

	// Handle TypeStringUnionTwo field
	toEvalTypeStringUnionTwo := obj.TypeStringUnionTwo
	if f.TypeStringUnionTwo != nil && !f.TypeStringUnionTwo.Eval(toEvalTypeStringUnionTwo) {
		return false
	}

	// Handle TypeStringSliceUnionTwo slice
	if f.TypeStringSliceUnionTwo != nil {
		for _, elem := range obj.TypeStringSliceUnionTwo {
			if !f.TypeStringSliceUnionTwo.Eval(&elem) {
				return false
			}
		}
	}

	// Handle TypeIntUnionTwo field
	toEvalTypeIntUnionTwo := toIntPtr(obj.TypeIntUnionTwo)
	if f.TypeIntUnionTwo != nil && !f.TypeIntUnionTwo.Eval(toEvalTypeIntUnionTwo) {
		return false
	}

	return true
}

func (f *NestedFilterUnionTypeOne) Eval(obj *UnionTypeOne) bool {
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

	// Handle TypeTimeUnionOne field
	toEvalTypeTimeUnionOne := obj.TypeTimeUnionOne
	if f.TypeTimeUnionOne != nil && !f.TypeTimeUnionOne.Eval(toEvalTypeTimeUnionOne) {
		return false
	}

	// Handle TypeStringUnionOne field
	toEvalTypeStringUnionOne := obj.TypeStringUnionOne
	if f.TypeStringUnionOne != nil && !f.TypeStringUnionOne.Eval(toEvalTypeStringUnionOne) {
		return false
	}

	// Handle TypeNested field
	toEvalTypeNested := obj.TypeNested
	if f.TypeNested != nil && !f.TypeNested.Eval(toEvalTypeNested) {
		return false
	}

	// Handle TypeIntUnionOne field
	toEvalTypeIntUnionOne := toIntPtr(obj.TypeIntUnionOne)
	if f.TypeIntUnionOne != nil && !f.TypeIntUnionOne.Eval(toEvalTypeIntUnionOne) {
		return false
	}

	return true
}

func (f *NestedFilterUnionOne) Eval(obj *UnionOne) bool {
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

	// Handle union objects depending of the type
	tobj := *obj
	switch objv := tobj.(type) {
	case UnionTypeOne:

		// Handle UnionTypeOne field
		toEvalUnionTypeOne := objv
		if f.UnionTypeOne != nil && !f.UnionTypeOne.Eval(&toEvalUnionTypeOne) {
			return false
		}

	case UnionTypeTwo:

		// Handle UnionTypeTwo field
		toEvalUnionTypeTwo := objv
		if f.UnionTypeTwo != nil && !f.UnionTypeTwo.Eval(&toEvalUnionTypeTwo) {
			return false
		}

	}

	return true
}

func (f *NestedFilterTypeTwo) Eval(obj *TypeTwo) bool {
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

	// Handle TypeTwoWithTypeThreeNotMandatory field
	toEvalTypeTwoWithTypeThreeNotMandatory := obj.TypeTwoWithTypeThreeNotMandatory
	if f.TypeTwoWithTypeThreeNotMandatory != nil && !f.TypeTwoWithTypeThreeNotMandatory.Eval(toEvalTypeTwoWithTypeThreeNotMandatory) {
		return false
	}

	// Handle TypeTwoWithTypeThree field
	toEvalTypeTwoWithTypeThree := obj.TypeTwoWithTypeThree
	if f.TypeTwoWithTypeThree != nil && !f.TypeTwoWithTypeThree.Eval(toEvalTypeTwoWithTypeThree) {
		return false
	}

	// Handle TypeTwoTimeFieldFiltered field
	toEvalTypeTwoTimeFieldFiltered := obj.TypeTwoTimeFieldFiltered
	if f.TypeTwoTimeFieldFiltered != nil && !f.TypeTwoTimeFieldFiltered.Eval(&toEvalTypeTwoTimeFieldFiltered) {
		return false
	}

	// Handle TypeTwoStringFieldFiltered field
	toEvalTypeTwoStringFieldFiltered := obj.TypeTwoStringFieldFiltered
	if f.TypeTwoStringFieldFiltered != nil && !f.TypeTwoStringFieldFiltered.Eval(&toEvalTypeTwoStringFieldFiltered) {
		return false
	}

	// Handle TypeTwoSliceWithTypeTwos slice
	if f.TypeTwoSliceWithTypeTwos != nil {
		for _, elem := range obj.TypeTwoSliceWithTypeTwos {
			if !f.TypeTwoSliceWithTypeTwos.Eval(elem) {
				return false
			}
		}
	}

	// Handle TypeTwoNumberFieldFiltered field
	toEvalTypeTwoNumberFieldFiltered := toIntPtr(obj.TypeTwoNumberFieldFiltered)
	if f.TypeTwoNumberFieldFiltered != nil && !f.TypeTwoNumberFieldFiltered.Eval(toEvalTypeTwoNumberFieldFiltered) {
		return false
	}

	// Handle TypeTwoBooleanFieldFiltered field
	toEvalTypeTwoBooleanFieldFiltered := obj.TypeTwoBooleanFieldFiltered
	if f.TypeTwoBooleanFieldFiltered != nil && !f.TypeTwoBooleanFieldFiltered.Eval(&toEvalTypeTwoBooleanFieldFiltered) {
		return false
	}

	return true
}

func (f *NestedFilterTypeThree) Eval(obj *TypeThree) bool {
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

	// Handle TypeUnionSlice slice
	if f.TypeUnionSlice != nil {
		for _, elem := range obj.TypeUnionSlice {
			if !f.TypeUnionSlice.Eval(&elem) {
				return false
			}
		}
	}

	// Handle TypeUnionNotMandatory field
	toEvalTypeUnionNotMandatory := obj.TypeUnionNotMandatory
	if f.TypeUnionNotMandatory != nil && !f.TypeUnionNotMandatory.Eval(&toEvalTypeUnionNotMandatory) {
		return false
	}

	// Handle TypeUnion field
	toEvalTypeUnion := obj.TypeUnion
	if f.TypeUnion != nil && !f.TypeUnion.Eval(&toEvalTypeUnion) {
		return false
	}

	// Handle TypeThreeTimeFieldFiltered field
	toEvalTypeThreeTimeFieldFiltered := obj.TypeThreeTimeFieldFiltered
	if f.TypeThreeTimeFieldFiltered != nil && !f.TypeThreeTimeFieldFiltered.Eval(&toEvalTypeThreeTimeFieldFiltered) {
		return false
	}

	// Handle TypeThreeStringFieldFiltered field
	toEvalTypeThreeStringFieldFiltered := obj.TypeThreeStringFieldFiltered
	if f.TypeThreeStringFieldFiltered != nil && !f.TypeThreeStringFieldFiltered.Eval(&toEvalTypeThreeStringFieldFiltered) {
		return false
	}

	// Handle TypeThreeNumberFieldFiltered field
	toEvalTypeThreeNumberFieldFiltered := toIntPtr(obj.TypeThreeNumberFieldFiltered)
	if f.TypeThreeNumberFieldFiltered != nil && !f.TypeThreeNumberFieldFiltered.Eval(toEvalTypeThreeNumberFieldFiltered) {
		return false
	}

	// Handle TypeThreeBooleanFieldFiltered field
	toEvalTypeThreeBooleanFieldFiltered := obj.TypeThreeBooleanFieldFiltered
	if f.TypeThreeBooleanFieldFiltered != nil && !f.TypeThreeBooleanFieldFiltered.Eval(&toEvalTypeThreeBooleanFieldFiltered) {
		return false
	}

	return true
}

func (f *NestedFilterTypeOne) Eval(obj *TypeOne) bool {
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

	// Handle TypeOneTimeFieldFilteredNotMandatory field
	toEvalTypeOneTimeFieldFilteredNotMandatory := obj.TypeOneTimeFieldFilteredNotMandatory
	if f.TypeOneTimeFieldFilteredNotMandatory != nil && !f.TypeOneTimeFieldFilteredNotMandatory.Eval(toEvalTypeOneTimeFieldFilteredNotMandatory) {
		return false
	}

	// Handle TypeOneTimeFieldFiltered field
	toEvalTypeOneTimeFieldFiltered := obj.TypeOneTimeFieldFiltered
	if f.TypeOneTimeFieldFiltered != nil && !f.TypeOneTimeFieldFiltered.Eval(&toEvalTypeOneTimeFieldFiltered) {
		return false
	}

	// Handle TypeOneStringSliceFiltered slice
	if f.TypeOneStringSliceFiltered != nil {
		for _, elem := range obj.TypeOneStringSliceFiltered {
			if !f.TypeOneStringSliceFiltered.Eval(&elem) {
				return false
			}
		}
	}

	// Handle TypeOneStringFieldFilteredNotMandatory field
	toEvalTypeOneStringFieldFilteredNotMandatory := obj.TypeOneStringFieldFilteredNotMandatory
	if f.TypeOneStringFieldFilteredNotMandatory != nil && !f.TypeOneStringFieldFilteredNotMandatory.Eval(toEvalTypeOneStringFieldFilteredNotMandatory) {
		return false
	}

	// Handle TypeOneStringFieldFiltered field
	toEvalTypeOneStringFieldFiltered := obj.TypeOneStringFieldFiltered
	if f.TypeOneStringFieldFiltered != nil && !f.TypeOneStringFieldFiltered.Eval(&toEvalTypeOneStringFieldFiltered) {
		return false
	}

	// Handle TypeOneSliceWithTypeTwos slice
	if f.TypeOneSliceWithTypeTwos != nil {
		for _, elem := range obj.TypeOneSliceWithTypeTwos {
			if !f.TypeOneSliceWithTypeTwos.Eval(elem) {
				return false
			}
		}
	}

	// Handle TypeOneNumberFieldFilteredNotMandatory field
	toEvalTypeOneNumberFieldFilteredNotMandatory := toIntPtr(obj.TypeOneNumberFieldFilteredNotMandatory)
	if f.TypeOneNumberFieldFilteredNotMandatory != nil && !f.TypeOneNumberFieldFilteredNotMandatory.Eval(toEvalTypeOneNumberFieldFilteredNotMandatory) {
		return false
	}

	// Handle TypeOneNumberFieldFiltered field
	toEvalTypeOneNumberFieldFiltered := toIntPtr(obj.TypeOneNumberFieldFiltered)
	if f.TypeOneNumberFieldFiltered != nil && !f.TypeOneNumberFieldFiltered.Eval(toEvalTypeOneNumberFieldFiltered) {
		return false
	}

	// Handle TypeOneBooleanFieldFilteredNotMandatory field
	toEvalTypeOneBooleanFieldFilteredNotMandatory := obj.TypeOneBooleanFieldFilteredNotMandatory
	if f.TypeOneBooleanFieldFilteredNotMandatory != nil && !f.TypeOneBooleanFieldFilteredNotMandatory.Eval(toEvalTypeOneBooleanFieldFilteredNotMandatory) {
		return false
	}

	// Handle TypeOneBooleanFieldFiltered field
	toEvalTypeOneBooleanFieldFiltered := obj.TypeOneBooleanFieldFiltered
	if f.TypeOneBooleanFieldFiltered != nil && !f.TypeOneBooleanFieldFiltered.Eval(&toEvalTypeOneBooleanFieldFiltered) {
		return false
	}

	return true
}

// MinMax function for TypeOneNumberFieldFiltered
func (f *NestedFilterTypeOne) MinMaxTypeOneNumberFieldFiltered() (min *int, max *int) {
	// Recursively handle And conditions
	if len(f.And) > 0 {
		for _, subFilter := range f.And {
			subMin, subMax := subFilter.MinMaxTypeOneNumberFieldFiltered()
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
			subMin, subMax := subFilter.MinMaxTypeOneNumberFieldFiltered()
			if subMin != nil && (min == nil || *subMin < *min) {
				min = subMin
			}
			if subMax != nil && (max == nil || *subMax > *max) {
				max = subMax
			}
		}
	}

	if f.TypeOneNumberFieldFiltered != nil {
		if f.TypeOneNumberFieldFiltered.Gt != nil {
			if min == nil || *f.TypeOneNumberFieldFiltered.Gt < *min {
				min = f.TypeOneNumberFieldFiltered.Gt
			}
			if max == nil || *f.TypeOneNumberFieldFiltered.Gt > *max {
				max = f.TypeOneNumberFieldFiltered.Gt
			}
		}

		if f.TypeOneNumberFieldFiltered.Lt != nil {
			if min == nil || *f.TypeOneNumberFieldFiltered.Lt < *min {
				min = f.TypeOneNumberFieldFiltered.Lt
			}
			if max == nil || *f.TypeOneNumberFieldFiltered.Lt > *max {
				max = f.TypeOneNumberFieldFiltered.Lt
			}
		}

		if f.TypeOneNumberFieldFiltered.Eq != nil {
			if min == nil || *f.TypeOneNumberFieldFiltered.Eq < *min {
				min = f.TypeOneNumberFieldFiltered.Eq
			}
			if max == nil || *f.TypeOneNumberFieldFiltered.Eq > *max {
				max = f.TypeOneNumberFieldFiltered.Eq
			}
		}
	}

	return min, max
}

func (f *NestedFilterNestedTypeTwo) Eval(obj *NestedTypeTwo) bool {
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

	// Handle ValString field
	toEvalValString := obj.ValString
	if f.ValString != nil && !f.ValString.Eval(&toEvalValString) {
		return false
	}

	return true
}

func (f *NestedFilterNestedType) Eval(obj *NestedType) bool {
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

	// Handle NestedOnNested slice
	if f.NestedOnNested != nil {
		for _, elem := range obj.NestedOnNested {
			if !f.NestedOnNested.Eval(elem) {
				return false
			}
		}
	}

	return true
}

func (f *FilterUnionTypeTwo) Eval(obj *UnionTypeTwo) bool {
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

	// Handle TypeTimeUnionTwo field
	toEvalTypeTimeUnionTwo := obj.TypeTimeUnionTwo
	if f.TypeTimeUnionTwo != nil && !f.TypeTimeUnionTwo.Eval(toEvalTypeTimeUnionTwo) {
		return false
	}

	// Handle TypeStringUnionTwo field
	toEvalTypeStringUnionTwo := obj.TypeStringUnionTwo
	if f.TypeStringUnionTwo != nil && !f.TypeStringUnionTwo.Eval(toEvalTypeStringUnionTwo) {
		return false
	}

	// Handle TypeStringSliceUnionTwo slice
	if f.TypeStringSliceUnionTwo != nil {
		for _, elem := range obj.TypeStringSliceUnionTwo {
			if !f.TypeStringSliceUnionTwo.Eval(&elem) {
				return false
			}
		}
	}

	// Handle TypeIntUnionTwo field
	toEvalTypeIntUnionTwo := toIntPtr(obj.TypeIntUnionTwo)
	if f.TypeIntUnionTwo != nil && !f.TypeIntUnionTwo.Eval(toEvalTypeIntUnionTwo) {
		return false
	}

	return true
}

func (f *FilterUnionTypeOne) Eval(obj *UnionTypeOne) bool {
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

	// Handle TypeTimeUnionOne field
	toEvalTypeTimeUnionOne := obj.TypeTimeUnionOne
	if f.TypeTimeUnionOne != nil && !f.TypeTimeUnionOne.Eval(toEvalTypeTimeUnionOne) {
		return false
	}

	// Handle TypeStringUnionOne field
	toEvalTypeStringUnionOne := obj.TypeStringUnionOne
	if f.TypeStringUnionOne != nil && !f.TypeStringUnionOne.Eval(toEvalTypeStringUnionOne) {
		return false
	}

	// Handle TypeNested field
	toEvalTypeNested := obj.TypeNested
	if f.TypeNested != nil && !f.TypeNested.Eval(toEvalTypeNested) {
		return false
	}

	// Handle TypeIntUnionOne field
	toEvalTypeIntUnionOne := toIntPtr(obj.TypeIntUnionOne)
	if f.TypeIntUnionOne != nil && !f.TypeIntUnionOne.Eval(toEvalTypeIntUnionOne) {
		return false
	}

	return true
}

func (f *FilterUnionOne) Eval(obj *UnionOne) bool {
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

	// Handle union objects depending of the type
	tobj := *obj
	switch objv := tobj.(type) {
	case UnionTypeOne:

		// Handle UnionTypeOne field
		toEvalUnionTypeOne := objv
		if f.UnionTypeOne != nil && !f.UnionTypeOne.Eval(&toEvalUnionTypeOne) {
			return false
		}

	case UnionTypeTwo:

		// Handle UnionTypeTwo field
		toEvalUnionTypeTwo := objv
		if f.UnionTypeTwo != nil && !f.UnionTypeTwo.Eval(&toEvalUnionTypeTwo) {
			return false
		}

	}

	return true
}

func (f *FilterTypeTwo) Eval(obj *TypeTwo) bool {
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

	// Handle TypeTwoWithTypeThreeNotMandatory field
	toEvalTypeTwoWithTypeThreeNotMandatory := obj.TypeTwoWithTypeThreeNotMandatory
	if f.TypeTwoWithTypeThreeNotMandatory != nil && !f.TypeTwoWithTypeThreeNotMandatory.Eval(toEvalTypeTwoWithTypeThreeNotMandatory) {
		return false
	}

	// Handle TypeTwoWithTypeThree field
	toEvalTypeTwoWithTypeThree := obj.TypeTwoWithTypeThree
	if f.TypeTwoWithTypeThree != nil && !f.TypeTwoWithTypeThree.Eval(toEvalTypeTwoWithTypeThree) {
		return false
	}

	// Handle TypeTwoTimeFieldFiltered field
	toEvalTypeTwoTimeFieldFiltered := obj.TypeTwoTimeFieldFiltered
	if f.TypeTwoTimeFieldFiltered != nil && !f.TypeTwoTimeFieldFiltered.Eval(&toEvalTypeTwoTimeFieldFiltered) {
		return false
	}

	// Handle TypeTwoStringFieldFiltered field
	toEvalTypeTwoStringFieldFiltered := obj.TypeTwoStringFieldFiltered
	if f.TypeTwoStringFieldFiltered != nil && !f.TypeTwoStringFieldFiltered.Eval(&toEvalTypeTwoStringFieldFiltered) {
		return false
	}

	// Handle TypeTwoSliceWithTypeTwos slice
	if f.TypeTwoSliceWithTypeTwos != nil {
		for _, elem := range obj.TypeTwoSliceWithTypeTwos {
			if !f.TypeTwoSliceWithTypeTwos.Eval(elem) {
				return false
			}
		}
	}

	// Handle TypeTwoNumberFieldFiltered field
	toEvalTypeTwoNumberFieldFiltered := toIntPtr(obj.TypeTwoNumberFieldFiltered)
	if f.TypeTwoNumberFieldFiltered != nil && !f.TypeTwoNumberFieldFiltered.Eval(toEvalTypeTwoNumberFieldFiltered) {
		return false
	}

	// Handle TypeTwoBooleanFieldFiltered field
	toEvalTypeTwoBooleanFieldFiltered := obj.TypeTwoBooleanFieldFiltered
	if f.TypeTwoBooleanFieldFiltered != nil && !f.TypeTwoBooleanFieldFiltered.Eval(&toEvalTypeTwoBooleanFieldFiltered) {
		return false
	}

	return true
}

func (f *FilterTypeThree) Eval(obj *TypeThree) bool {
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

	// Handle TypeUnionSlice slice
	if f.TypeUnionSlice != nil {
		for _, elem := range obj.TypeUnionSlice {
			if !f.TypeUnionSlice.Eval(&elem) {
				return false
			}
		}
	}

	// Handle TypeUnionNotMandatory field
	toEvalTypeUnionNotMandatory := obj.TypeUnionNotMandatory
	if f.TypeUnionNotMandatory != nil && !f.TypeUnionNotMandatory.Eval(&toEvalTypeUnionNotMandatory) {
		return false
	}

	// Handle TypeUnion field
	toEvalTypeUnion := obj.TypeUnion
	if f.TypeUnion != nil && !f.TypeUnion.Eval(&toEvalTypeUnion) {
		return false
	}

	// Handle TypeThreeTimeFieldFiltered field
	toEvalTypeThreeTimeFieldFiltered := obj.TypeThreeTimeFieldFiltered
	if f.TypeThreeTimeFieldFiltered != nil && !f.TypeThreeTimeFieldFiltered.Eval(&toEvalTypeThreeTimeFieldFiltered) {
		return false
	}

	// Handle TypeThreeStringFieldFiltered field
	toEvalTypeThreeStringFieldFiltered := obj.TypeThreeStringFieldFiltered
	if f.TypeThreeStringFieldFiltered != nil && !f.TypeThreeStringFieldFiltered.Eval(&toEvalTypeThreeStringFieldFiltered) {
		return false
	}

	// Handle TypeThreeNumberFieldFiltered field
	toEvalTypeThreeNumberFieldFiltered := toIntPtr(obj.TypeThreeNumberFieldFiltered)
	if f.TypeThreeNumberFieldFiltered != nil && !f.TypeThreeNumberFieldFiltered.Eval(toEvalTypeThreeNumberFieldFiltered) {
		return false
	}

	// Handle TypeThreeBooleanFieldFiltered field
	toEvalTypeThreeBooleanFieldFiltered := obj.TypeThreeBooleanFieldFiltered
	if f.TypeThreeBooleanFieldFiltered != nil && !f.TypeThreeBooleanFieldFiltered.Eval(&toEvalTypeThreeBooleanFieldFiltered) {
		return false
	}

	return true
}

func (f *FilterTypeOne) Eval(obj *TypeOne) bool {
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

	// Handle TypeOneTimeFieldFilteredNotMandatory field
	toEvalTypeOneTimeFieldFilteredNotMandatory := obj.TypeOneTimeFieldFilteredNotMandatory
	if f.TypeOneTimeFieldFilteredNotMandatory != nil && !f.TypeOneTimeFieldFilteredNotMandatory.Eval(toEvalTypeOneTimeFieldFilteredNotMandatory) {
		return false
	}

	// Handle TypeOneTimeFieldFiltered field
	toEvalTypeOneTimeFieldFiltered := obj.TypeOneTimeFieldFiltered
	if f.TypeOneTimeFieldFiltered != nil && !f.TypeOneTimeFieldFiltered.Eval(&toEvalTypeOneTimeFieldFiltered) {
		return false
	}

	// Handle TypeOneStringSliceFiltered slice
	if f.TypeOneStringSliceFiltered != nil {
		for _, elem := range obj.TypeOneStringSliceFiltered {
			if !f.TypeOneStringSliceFiltered.Eval(&elem) {
				return false
			}
		}
	}

	// Handle TypeOneStringFieldFilteredNotMandatory field
	toEvalTypeOneStringFieldFilteredNotMandatory := obj.TypeOneStringFieldFilteredNotMandatory
	if f.TypeOneStringFieldFilteredNotMandatory != nil && !f.TypeOneStringFieldFilteredNotMandatory.Eval(toEvalTypeOneStringFieldFilteredNotMandatory) {
		return false
	}

	// Handle TypeOneStringFieldFiltered field
	toEvalTypeOneStringFieldFiltered := obj.TypeOneStringFieldFiltered
	if f.TypeOneStringFieldFiltered != nil && !f.TypeOneStringFieldFiltered.Eval(&toEvalTypeOneStringFieldFiltered) {
		return false
	}

	// Handle TypeOneSliceWithTypeTwos slice
	if f.TypeOneSliceWithTypeTwos != nil {
		for _, elem := range obj.TypeOneSliceWithTypeTwos {
			if !f.TypeOneSliceWithTypeTwos.Eval(elem) {
				return false
			}
		}
	}

	// Handle TypeOneNumberFieldFilteredNotMandatory field
	toEvalTypeOneNumberFieldFilteredNotMandatory := toIntPtr(obj.TypeOneNumberFieldFilteredNotMandatory)
	if f.TypeOneNumberFieldFilteredNotMandatory != nil && !f.TypeOneNumberFieldFilteredNotMandatory.Eval(toEvalTypeOneNumberFieldFilteredNotMandatory) {
		return false
	}

	// Handle TypeOneNumberFieldFiltered field
	toEvalTypeOneNumberFieldFiltered := toIntPtr(obj.TypeOneNumberFieldFiltered)
	if f.TypeOneNumberFieldFiltered != nil && !f.TypeOneNumberFieldFiltered.Eval(toEvalTypeOneNumberFieldFiltered) {
		return false
	}

	// Handle TypeOneBooleanFieldFilteredNotMandatory field
	toEvalTypeOneBooleanFieldFilteredNotMandatory := obj.TypeOneBooleanFieldFilteredNotMandatory
	if f.TypeOneBooleanFieldFilteredNotMandatory != nil && !f.TypeOneBooleanFieldFilteredNotMandatory.Eval(toEvalTypeOneBooleanFieldFilteredNotMandatory) {
		return false
	}

	// Handle TypeOneBooleanFieldFiltered field
	toEvalTypeOneBooleanFieldFiltered := obj.TypeOneBooleanFieldFiltered
	if f.TypeOneBooleanFieldFiltered != nil && !f.TypeOneBooleanFieldFiltered.Eval(&toEvalTypeOneBooleanFieldFiltered) {
		return false
	}

	return true
}

// MinMax function for TypeOneNumberFieldFiltered
func (f *FilterTypeOne) MinMaxTypeOneNumberFieldFiltered() (min *int, max *int) {
	// Recursively handle And conditions
	if len(f.And) > 0 {
		for _, subFilter := range f.And {
			subMin, subMax := subFilter.MinMaxTypeOneNumberFieldFiltered()
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
			subMin, subMax := subFilter.MinMaxTypeOneNumberFieldFiltered()
			if subMin != nil && (min == nil || *subMin < *min) {
				min = subMin
			}
			if subMax != nil && (max == nil || *subMax > *max) {
				max = subMax
			}
		}
	}

	if f.TypeOneNumberFieldFiltered != nil {
		if f.TypeOneNumberFieldFiltered.Gt != nil {
			if min == nil || *f.TypeOneNumberFieldFiltered.Gt < *min {
				min = f.TypeOneNumberFieldFiltered.Gt
			}
			if max == nil || *f.TypeOneNumberFieldFiltered.Gt > *max {
				max = f.TypeOneNumberFieldFiltered.Gt
			}
		}

		if f.TypeOneNumberFieldFiltered.Lt != nil {
			if min == nil || *f.TypeOneNumberFieldFiltered.Lt < *min {
				min = f.TypeOneNumberFieldFiltered.Lt
			}
			if max == nil || *f.TypeOneNumberFieldFiltered.Lt > *max {
				max = f.TypeOneNumberFieldFiltered.Lt
			}
		}

		if f.TypeOneNumberFieldFiltered.Eq != nil {
			if min == nil || *f.TypeOneNumberFieldFiltered.Eq < *min {
				min = f.TypeOneNumberFieldFiltered.Eq
			}
			if max == nil || *f.TypeOneNumberFieldFiltered.Eq > *max {
				max = f.TypeOneNumberFieldFiltered.Eq
			}
		}
	}

	return min, max
}

func (f *FilterNestedTypeTwo) Eval(obj *NestedTypeTwo) bool {
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

	// Handle ValString field
	toEvalValString := obj.ValString
	if f.ValString != nil && !f.ValString.Eval(&toEvalValString) {
		return false
	}

	return true
}

func (f *FilterNestedType) Eval(obj *NestedType) bool {
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

	// Handle NestedOnNested slice
	if f.NestedOnNested != nil {
		for _, elem := range obj.NestedOnNested {
			if !f.NestedOnNested.Eval(elem) {
				return false
			}
		}
	}

	return true
}

func (f *FilterExternalType) Eval(obj *ExternalType) bool {
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

	// Handle TypeOne field
	toEvalTypeOne := obj.TypeOne()
	if f.TypeOne != nil && !f.TypeOne.Eval(toEvalTypeOne) {
		return false
	}

	// Handle NumberTwo field
	toEvalNumberTwo := toIntPtr(obj.NumberTwo())
	if f.NumberTwo != nil && !f.NumberTwo.Eval(toEvalNumberTwo) {
		return false
	}

	// Handle NumberThree field
	toEvalNumberThree := toIntPtr(obj.NumberThree())
	if f.NumberThree != nil && !f.NumberThree.Eval(toEvalNumberThree) {
		return false
	}

	// Handle NumberOne field
	toEvalNumberOne := toIntPtr(obj.NumberOne())
	if f.NumberOne != nil && !f.NumberOne.Eval(toEvalNumberOne) {
		return false
	}

	// Handle NumberList slice
	if f.NumberList != nil {
		for _, elem := range obj.NumberList() {
			if !f.NumberList.Eval(&elem) {
				return false
			}
		}
	}

	// Handle NumberFour field
	toEvalNumberFour := toIntPtr(obj.NumberFour())
	if f.NumberFour != nil && !f.NumberFour.Eval(toEvalNumberFour) {
		return false
	}

	// Handle NumberFive field
	toEvalNumberFive := toIntPtr(obj.NumberFive())
	if f.NumberFive != nil && !f.NumberFive.Eval(toEvalNumberFive) {
		return false
	}

	return true
}

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

	return rootEval(val, f.Exists, f.Eq, nil)
}

func (f *FilterInt) Eval(val *int) bool {
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
