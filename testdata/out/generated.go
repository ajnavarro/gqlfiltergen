// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package out

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type UnionOne interface {
	IsUnionOne()
}

type UnionTwo interface {
	IsUnionTwo()
}

// Filter type for boolean fields. All added filters here are processed as AND operators.
type FilterBoolean struct {
	// Filter a boolean field checking if it exists or not.
	Exists *bool `json:"exists,omitempty"`
	// Filter a boolean field checking if it is equals to the specified value.
	Eq *bool `json:"eq,omitempty"`
}

// filter for ExternalType objects
type FilterExternalType struct {
	// logical operator for ExternalType that will combine two or more conditions, returning true if all of them are true.
	And []*FilterExternalType `json:"_and,omitempty"`
	// logical operator for ExternalType that will combine two or more conditions, returning true if at least one of them is true.
	Or []*FilterExternalType `json:"_or,omitempty"`
	// logical operator for ExternalType that will reverse conditions.
	Not *FilterExternalType `json:"_not,omitempty"`
	// filter for number_one field.
	NumberOne *FilterInt `json:"number_one,omitempty"`
	// filter for number_two field.
	NumberTwo *FilterInt `json:"number_two,omitempty"`
	// filter for number_three field.
	NumberThree *FilterInt `json:"number_three,omitempty"`
	// filter for number_four field.
	NumberFour *FilterInt `json:"number_four,omitempty"`
	// filter for number_five field.
	NumberFive *FilterInt `json:"number_five,omitempty"`
	// filter for number_list field.
	NumberList *FilterInt `json:"number_list,omitempty"`
	// filter for type_one field.
	TypeOne *NestedFilterTypeOne `json:"type_one,omitempty"`
}

// Filter type for number fields. All added filters here are processed as AND operators.
type FilterInt struct {
	// Filter a number field checking if it exists or not.
	Exists *bool `json:"exists,omitempty"`
	// Filter a number field checking if it is equals to the specified value.
	Eq *int `json:"eq,omitempty"`
	// Filter a number field checking if it is greater than the specified value.
	Gt *int `json:"gt,omitempty"`
	// Filter a number field checking if it is less than the specified value.
	Lt *int `json:"lt,omitempty"`
}

// filter for NestedType objects
type FilterNestedType struct {
	// logical operator for NestedType that will combine two or more conditions, returning true if all of them are true.
	And []*FilterNestedType `json:"_and,omitempty"`
	// logical operator for NestedType that will combine two or more conditions, returning true if at least one of them is true.
	Or []*FilterNestedType `json:"_or,omitempty"`
	// logical operator for NestedType that will reverse conditions.
	Not *FilterNestedType `json:"_not,omitempty"`
	// filter for nested_on_nested field.
	NestedOnNested *NestedFilterNestedTypeTwo `json:"nested_on_nested,omitempty"`
}

// filter for NestedTypeTwo objects
type FilterNestedTypeTwo struct {
	// logical operator for NestedTypeTwo that will combine two or more conditions, returning true if all of them are true.
	And []*FilterNestedTypeTwo `json:"_and,omitempty"`
	// logical operator for NestedTypeTwo that will combine two or more conditions, returning true if at least one of them is true.
	Or []*FilterNestedTypeTwo `json:"_or,omitempty"`
	// logical operator for NestedTypeTwo that will reverse conditions.
	Not *FilterNestedTypeTwo `json:"_not,omitempty"`
	// filter for val_string field.
	ValString *FilterString `json:"val_string,omitempty"`
}

// Filter type for string fields. It contains a variety of filter types for string types. All added filters here are processed as AND operators.
type FilterString struct {
	// Filter a string field checking if it exists or not.
	Exists *bool `json:"exists,omitempty"`
	// Filter a string field checking if it is equals to the specified value.
	Eq *string `json:"eq,omitempty"`
	// Filter a string field checking if it is like the specified value. You can use standard Go RegEx expressions here.
	Like *string `json:"like,omitempty"`
}

// Filter type for time fields. All added filters here are processed as AND operators.
type FilterTime struct {
	// Filter a time field checking if it exists or not.
	Exists *bool `json:"exists,omitempty"`
	// Filter a time field checking if it is equals to the specified value.
	Eq *time.Time `json:"eq,omitempty"`
	// Filter a time field checking if it is before than the specified value.
	Before *time.Time `json:"before,omitempty"`
	// Filter a time field checking if it is after the specified value.
	After *time.Time `json:"after,omitempty"`
}

// filter for TypeOne objects
type FilterTypeOne struct {
	// logical operator for TypeOne that will combine two or more conditions, returning true if all of them are true.
	And []*FilterTypeOne `json:"_and,omitempty"`
	// logical operator for TypeOne that will combine two or more conditions, returning true if at least one of them is true.
	Or []*FilterTypeOne `json:"_or,omitempty"`
	// logical operator for TypeOne that will reverse conditions.
	Not *FilterTypeOne `json:"_not,omitempty"`
	// filter for type_one_string_field_filtered field.
	TypeOneStringFieldFiltered *FilterString `json:"type_one_string_field_filtered,omitempty"`
	// filter for type_one_string_field_filtered_not_mandatory field.
	TypeOneStringFieldFilteredNotMandatory *FilterString `json:"type_one_string_field_filtered_not_mandatory,omitempty"`
	// filter for type_one_number_field_filtered field.
	TypeOneNumberFieldFiltered *FilterInt `json:"type_one_number_field_filtered,omitempty"`
	// filter for type_one_string_slice_filtered field.
	TypeOneStringSliceFiltered *FilterString `json:"type_one_string_slice_filtered,omitempty"`
	// filter for type_one_number_field_filtered_not_mandatory field.
	TypeOneNumberFieldFilteredNotMandatory *FilterInt `json:"type_one_number_field_filtered_not_mandatory,omitempty"`
	// filter for type_one_time_field_filtered field.
	TypeOneTimeFieldFiltered *FilterTime `json:"type_one_time_field_filtered,omitempty"`
	// filter for type_one_time_field_filtered_not_mandatory field.
	TypeOneTimeFieldFilteredNotMandatory *FilterTime `json:"type_one_time_field_filtered_not_mandatory,omitempty"`
	// filter for type_one_boolean_field_filtered field.
	TypeOneBooleanFieldFiltered *FilterBoolean `json:"type_one_boolean_field_filtered,omitempty"`
	// filter for type_one_boolean_field_filtered_not_mandatory field.
	TypeOneBooleanFieldFilteredNotMandatory *FilterBoolean `json:"type_one_boolean_field_filtered_not_mandatory,omitempty"`
	// filter for type_one_slice_with_type_twos field.
	TypeOneSliceWithTypeTwos *NestedFilterTypeTwo `json:"type_one_slice_with_type_twos,omitempty"`
}

// filter for TypeThree objects
type FilterTypeThree struct {
	// logical operator for TypeThree that will combine two or more conditions, returning true if all of them are true.
	And []*FilterTypeThree `json:"_and,omitempty"`
	// logical operator for TypeThree that will combine two or more conditions, returning true if at least one of them is true.
	Or []*FilterTypeThree `json:"_or,omitempty"`
	// logical operator for TypeThree that will reverse conditions.
	Not *FilterTypeThree `json:"_not,omitempty"`
	// filter for type_three_string_field_filtered field.
	TypeThreeStringFieldFiltered *FilterString `json:"type_three_string_field_filtered,omitempty"`
	// filter for type_three_number_field_filtered field.
	TypeThreeNumberFieldFiltered *FilterInt `json:"type_three_number_field_filtered,omitempty"`
	// filter for type_three_time_field_filtered field.
	TypeThreeTimeFieldFiltered *FilterTime `json:"type_three_time_field_filtered,omitempty"`
	// filter for type_three_boolean_field_filtered field.
	TypeThreeBooleanFieldFiltered *FilterBoolean `json:"type_three_boolean_field_filtered,omitempty"`
	// filter for type_union field.
	TypeUnion *NestedFilterUnionOne `json:"type_union,omitempty"`
	// filter for type_union_slice field.
	TypeUnionSlice *NestedFilterUnionOne `json:"type_union_slice,omitempty"`
	// filter for type_union_not_mandatory field.
	TypeUnionNotMandatory *NestedFilterUnionOne `json:"type_union_not_mandatory,omitempty"`
}

// filter for TypeTwo objects
type FilterTypeTwo struct {
	// logical operator for TypeTwo that will combine two or more conditions, returning true if all of them are true.
	And []*FilterTypeTwo `json:"_and,omitempty"`
	// logical operator for TypeTwo that will combine two or more conditions, returning true if at least one of them is true.
	Or []*FilterTypeTwo `json:"_or,omitempty"`
	// logical operator for TypeTwo that will reverse conditions.
	Not *FilterTypeTwo `json:"_not,omitempty"`
	// filter for type_two_string_field_filtered field.
	TypeTwoStringFieldFiltered *FilterString `json:"type_two_string_field_filtered,omitempty"`
	// filter for type_two_number_field_filtered field.
	TypeTwoNumberFieldFiltered *FilterInt `json:"type_two_number_field_filtered,omitempty"`
	// filter for type_two_time_field_filtered field.
	TypeTwoTimeFieldFiltered *FilterTime `json:"type_two_time_field_filtered,omitempty"`
	// filter for type_two_boolean_field_filtered field.
	TypeTwoBooleanFieldFiltered *FilterBoolean `json:"type_two_boolean_field_filtered,omitempty"`
	// filter for type_two_slice_with_type_twos field.
	TypeTwoSliceWithTypeTwos *NestedFilterTypeTwo `json:"type_two_slice_with_type_twos,omitempty"`
	// filter for type_two_with_type_three field.
	TypeTwoWithTypeThree *NestedFilterTypeThree `json:"type_two_with_type_three,omitempty"`
	// filter for type_two_with_type_three_not_mandatory field.
	TypeTwoWithTypeThreeNotMandatory *NestedFilterTypeThree `json:"type_two_with_type_three_not_mandatory,omitempty"`
}

// filter for UnionOne objects
type FilterUnionOne struct {
	// logical operator for UnionOne that will combine two or more conditions, returning true if all of them are true.
	And []*FilterUnionOne `json:"_and,omitempty"`
	// logical operator for UnionOne that will combine two or more conditions, returning true if at least one of them is true.
	Or []*FilterUnionOne `json:"_or,omitempty"`
	// logical operator for UnionOne that will reverse conditions.
	Not *FilterUnionOne `json:"_not,omitempty"`
	// filter for UnionTypeOne union type.
	UnionTypeOne *NestedFilterUnionTypeOne `json:"UnionTypeOne,omitempty"`
	// filter for UnionTypeTwo union type.
	UnionTypeTwo *NestedFilterUnionTypeTwo `json:"UnionTypeTwo,omitempty"`
}

// filter for UnionTypeOne objects
type FilterUnionTypeOne struct {
	// logical operator for UnionTypeOne that will combine two or more conditions, returning true if all of them are true.
	And []*FilterUnionTypeOne `json:"_and,omitempty"`
	// logical operator for UnionTypeOne that will combine two or more conditions, returning true if at least one of them is true.
	Or []*FilterUnionTypeOne `json:"_or,omitempty"`
	// logical operator for UnionTypeOne that will reverse conditions.
	Not *FilterUnionTypeOne `json:"_not,omitempty"`
	// filter for type_int_union_one field.
	TypeIntUnionOne *FilterInt `json:"type_int_union_one,omitempty"`
	// filter for type_string_union_one field.
	TypeStringUnionOne *FilterString `json:"type_string_union_one,omitempty"`
	// filter for type_time_union_one field.
	TypeTimeUnionOne *FilterTime `json:"type_time_union_one,omitempty"`
	// filter for type_nested field.
	TypeNested *NestedFilterNestedType `json:"type_nested,omitempty"`
}

// filter for UnionTypeTwo objects
type FilterUnionTypeTwo struct {
	// logical operator for UnionTypeTwo that will combine two or more conditions, returning true if all of them are true.
	And []*FilterUnionTypeTwo `json:"_and,omitempty"`
	// logical operator for UnionTypeTwo that will combine two or more conditions, returning true if at least one of them is true.
	Or []*FilterUnionTypeTwo `json:"_or,omitempty"`
	// logical operator for UnionTypeTwo that will reverse conditions.
	Not *FilterUnionTypeTwo `json:"_not,omitempty"`
	// filter for type_int_union_two field.
	TypeIntUnionTwo *FilterInt `json:"type_int_union_two,omitempty"`
	// filter for type_string_union_two field.
	TypeStringUnionTwo *FilterString `json:"type_string_union_two,omitempty"`
	// filter for type_time_union_two field.
	TypeTimeUnionTwo *FilterTime `json:"type_time_union_two,omitempty"`
	// filter for type_string_slice_union_two field.
	TypeStringSliceUnionTwo *FilterString `json:"type_string_slice_union_two,omitempty"`
}

type InputOne struct {
	TypeTwoStringFieldFiltered     string    `json:"type_two_string_field_filtered"`
	TypeTwoNumberFieldFiltered     int       `json:"type_two_number_field_filtered"`
	TypeTwoTimeFieldFiltered       time.Time `json:"type_two_time_field_filtered"`
	TypeTwoBooleanFieldFiltered    bool      `json:"type_two_boolean_field_filtered"`
	TypeTwoStringFieldWithNoFilter string    `json:"type_twoString_field_with_no_filter"`
	TypeTwoNumberFieldWithNoFilter int       `json:"type_twoNumber_field_with_no_filter"`
	TypeTwoTimeFieldWithNoFilter   time.Time `json:"type_twoTime_field_with_no_filter"`
}

// filter for NestedType objects
type NestedFilterNestedType struct {
	// logical operator for NestedType that will combine two or more conditions, returning true if all of them are true.
	And []*NestedFilterNestedType `json:"_and,omitempty"`
	// logical operator for NestedType that will combine two or more conditions, returning true if at least one of them is true.
	Or []*NestedFilterNestedType `json:"_or,omitempty"`
	// logical operator for NestedType that will reverse conditions.
	Not *NestedFilterNestedType `json:"_not,omitempty"`
	// filter for nested_on_nested field.
	NestedOnNested *NestedFilterNestedTypeTwo `json:"nested_on_nested,omitempty"`
}

// filter for NestedTypeTwo objects
type NestedFilterNestedTypeTwo struct {
	// logical operator for NestedTypeTwo that will combine two or more conditions, returning true if all of them are true.
	And []*NestedFilterNestedTypeTwo `json:"_and,omitempty"`
	// logical operator for NestedTypeTwo that will combine two or more conditions, returning true if at least one of them is true.
	Or []*NestedFilterNestedTypeTwo `json:"_or,omitempty"`
	// logical operator for NestedTypeTwo that will reverse conditions.
	Not *NestedFilterNestedTypeTwo `json:"_not,omitempty"`
	// filter for val_string field.
	ValString *FilterString `json:"val_string,omitempty"`
}

// filter for TypeOne objects
type NestedFilterTypeOne struct {
	// logical operator for TypeOne that will combine two or more conditions, returning true if all of them are true.
	And []*NestedFilterTypeOne `json:"_and,omitempty"`
	// logical operator for TypeOne that will combine two or more conditions, returning true if at least one of them is true.
	Or []*NestedFilterTypeOne `json:"_or,omitempty"`
	// logical operator for TypeOne that will reverse conditions.
	Not *NestedFilterTypeOne `json:"_not,omitempty"`
	// filter for type_one_string_field_filtered field.
	TypeOneStringFieldFiltered *FilterString `json:"type_one_string_field_filtered,omitempty"`
	// filter for type_one_string_field_filtered_not_mandatory field.
	TypeOneStringFieldFilteredNotMandatory *FilterString `json:"type_one_string_field_filtered_not_mandatory,omitempty"`
	// filter for type_one_number_field_filtered field.
	TypeOneNumberFieldFiltered *FilterInt `json:"type_one_number_field_filtered,omitempty"`
	// filter for type_one_string_slice_filtered field.
	TypeOneStringSliceFiltered *FilterString `json:"type_one_string_slice_filtered,omitempty"`
	// filter for type_one_number_field_filtered_not_mandatory field.
	TypeOneNumberFieldFilteredNotMandatory *FilterInt `json:"type_one_number_field_filtered_not_mandatory,omitempty"`
	// filter for type_one_time_field_filtered field.
	TypeOneTimeFieldFiltered *FilterTime `json:"type_one_time_field_filtered,omitempty"`
	// filter for type_one_time_field_filtered_not_mandatory field.
	TypeOneTimeFieldFilteredNotMandatory *FilterTime `json:"type_one_time_field_filtered_not_mandatory,omitempty"`
	// filter for type_one_boolean_field_filtered field.
	TypeOneBooleanFieldFiltered *FilterBoolean `json:"type_one_boolean_field_filtered,omitempty"`
	// filter for type_one_boolean_field_filtered_not_mandatory field.
	TypeOneBooleanFieldFilteredNotMandatory *FilterBoolean `json:"type_one_boolean_field_filtered_not_mandatory,omitempty"`
	// filter for type_one_slice_with_type_twos field.
	TypeOneSliceWithTypeTwos *NestedFilterTypeTwo `json:"type_one_slice_with_type_twos,omitempty"`
}

// filter for TypeThree objects
type NestedFilterTypeThree struct {
	// logical operator for TypeThree that will combine two or more conditions, returning true if all of them are true.
	And []*NestedFilterTypeThree `json:"_and,omitempty"`
	// logical operator for TypeThree that will combine two or more conditions, returning true if at least one of them is true.
	Or []*NestedFilterTypeThree `json:"_or,omitempty"`
	// logical operator for TypeThree that will reverse conditions.
	Not *NestedFilterTypeThree `json:"_not,omitempty"`
	// filter for type_three_string_field_filtered field.
	TypeThreeStringFieldFiltered *FilterString `json:"type_three_string_field_filtered,omitempty"`
	// filter for type_three_number_field_filtered field.
	TypeThreeNumberFieldFiltered *FilterInt `json:"type_three_number_field_filtered,omitempty"`
	// filter for type_three_time_field_filtered field.
	TypeThreeTimeFieldFiltered *FilterTime `json:"type_three_time_field_filtered,omitempty"`
	// filter for type_three_boolean_field_filtered field.
	TypeThreeBooleanFieldFiltered *FilterBoolean `json:"type_three_boolean_field_filtered,omitempty"`
	// filter for type_union field.
	TypeUnion *NestedFilterUnionOne `json:"type_union,omitempty"`
	// filter for type_union_slice field.
	TypeUnionSlice *NestedFilterUnionOne `json:"type_union_slice,omitempty"`
	// filter for type_union_not_mandatory field.
	TypeUnionNotMandatory *NestedFilterUnionOne `json:"type_union_not_mandatory,omitempty"`
}

// filter for TypeTwo objects
type NestedFilterTypeTwo struct {
	// logical operator for TypeTwo that will combine two or more conditions, returning true if all of them are true.
	And []*NestedFilterTypeTwo `json:"_and,omitempty"`
	// logical operator for TypeTwo that will combine two or more conditions, returning true if at least one of them is true.
	Or []*NestedFilterTypeTwo `json:"_or,omitempty"`
	// logical operator for TypeTwo that will reverse conditions.
	Not *NestedFilterTypeTwo `json:"_not,omitempty"`
	// filter for type_two_string_field_filtered field.
	TypeTwoStringFieldFiltered *FilterString `json:"type_two_string_field_filtered,omitempty"`
	// filter for type_two_number_field_filtered field.
	TypeTwoNumberFieldFiltered *FilterInt `json:"type_two_number_field_filtered,omitempty"`
	// filter for type_two_time_field_filtered field.
	TypeTwoTimeFieldFiltered *FilterTime `json:"type_two_time_field_filtered,omitempty"`
	// filter for type_two_boolean_field_filtered field.
	TypeTwoBooleanFieldFiltered *FilterBoolean `json:"type_two_boolean_field_filtered,omitempty"`
	// filter for type_two_slice_with_type_twos field.
	TypeTwoSliceWithTypeTwos *NestedFilterTypeTwo `json:"type_two_slice_with_type_twos,omitempty"`
	// filter for type_two_with_type_three field.
	TypeTwoWithTypeThree *NestedFilterTypeThree `json:"type_two_with_type_three,omitempty"`
	// filter for type_two_with_type_three_not_mandatory field.
	TypeTwoWithTypeThreeNotMandatory *NestedFilterTypeThree `json:"type_two_with_type_three_not_mandatory,omitempty"`
}

// filter for UnionOne objects
type NestedFilterUnionOne struct {
	// logical operator for UnionOne that will combine two or more conditions, returning true if all of them are true.
	And []*NestedFilterUnionOne `json:"_and,omitempty"`
	// logical operator for UnionOne that will combine two or more conditions, returning true if at least one of them is true.
	Or []*NestedFilterUnionOne `json:"_or,omitempty"`
	// logical operator for UnionOne that will reverse conditions.
	Not *NestedFilterUnionOne `json:"_not,omitempty"`
	// filter for UnionTypeOne union type.
	UnionTypeOne *NestedFilterUnionTypeOne `json:"UnionTypeOne,omitempty"`
	// filter for UnionTypeTwo union type.
	UnionTypeTwo *NestedFilterUnionTypeTwo `json:"UnionTypeTwo,omitempty"`
}

// filter for UnionTypeOne objects
type NestedFilterUnionTypeOne struct {
	// logical operator for UnionTypeOne that will combine two or more conditions, returning true if all of them are true.
	And []*NestedFilterUnionTypeOne `json:"_and,omitempty"`
	// logical operator for UnionTypeOne that will combine two or more conditions, returning true if at least one of them is true.
	Or []*NestedFilterUnionTypeOne `json:"_or,omitempty"`
	// logical operator for UnionTypeOne that will reverse conditions.
	Not *NestedFilterUnionTypeOne `json:"_not,omitempty"`
	// filter for type_int_union_one field.
	TypeIntUnionOne *FilterInt `json:"type_int_union_one,omitempty"`
	// filter for type_string_union_one field.
	TypeStringUnionOne *FilterString `json:"type_string_union_one,omitempty"`
	// filter for type_time_union_one field.
	TypeTimeUnionOne *FilterTime `json:"type_time_union_one,omitempty"`
	// filter for type_nested field.
	TypeNested *NestedFilterNestedType `json:"type_nested,omitempty"`
}

// filter for UnionTypeTwo objects
type NestedFilterUnionTypeTwo struct {
	// logical operator for UnionTypeTwo that will combine two or more conditions, returning true if all of them are true.
	And []*NestedFilterUnionTypeTwo `json:"_and,omitempty"`
	// logical operator for UnionTypeTwo that will combine two or more conditions, returning true if at least one of them is true.
	Or []*NestedFilterUnionTypeTwo `json:"_or,omitempty"`
	// logical operator for UnionTypeTwo that will reverse conditions.
	Not *NestedFilterUnionTypeTwo `json:"_not,omitempty"`
	// filter for type_int_union_two field.
	TypeIntUnionTwo *FilterInt `json:"type_int_union_two,omitempty"`
	// filter for type_string_union_two field.
	TypeStringUnionTwo *FilterString `json:"type_string_union_two,omitempty"`
	// filter for type_time_union_two field.
	TypeTimeUnionTwo *FilterTime `json:"type_time_union_two,omitempty"`
	// filter for type_string_slice_union_two field.
	TypeStringSliceUnionTwo *FilterString `json:"type_string_slice_union_two,omitempty"`
}

type NestedType struct {
	NestedOnNested []*NestedTypeTwo `json:"nested_on_nested,omitempty"`
}

type NestedTypeTwo struct {
	ValString string `json:"val_string"`
}

type Query struct {
}

type TypeOne struct {
	TypeOneStringFieldFiltered              string     `json:"type_one_string_field_filtered"`
	TypeOneStringFieldFilteredNotMandatory  *string    `json:"type_one_string_field_filtered_not_mandatory,omitempty"`
	TypeOneNumberFieldFiltered              int        `json:"type_one_number_field_filtered"`
	TypeOneStringSliceFiltered              []string   `json:"type_one_string_slice_filtered,omitempty"`
	TypeOneNumberFieldFilteredNotMandatory  *int       `json:"type_one_number_field_filtered_not_mandatory,omitempty"`
	TypeOneTimeFieldFiltered                time.Time  `json:"type_one_time_field_filtered"`
	TypeOneTimeFieldFilteredNotMandatory    *time.Time `json:"type_one_time_field_filtered_not_mandatory,omitempty"`
	TypeOneBooleanFieldFiltered             bool       `json:"type_one_boolean_field_filtered"`
	TypeOneBooleanFieldFilteredNotMandatory *bool      `json:"type_one_boolean_field_filtered_not_mandatory,omitempty"`
	TypeOneStringFieldWithNoFilter          string     `json:"type_one_string_field_with_no_filter"`
	TypeOneNumberFieldWithNoFilter          int        `json:"type_one_number_field_with_no_filter"`
	TypeOneTimeFieldWithNoFilter            time.Time  `json:"type_one_time_field_with_no_filter"`
	TypeOneSliceWithTypeTwos                []*TypeTwo `json:"type_one_slice_with_type_twos,omitempty"`
}

type TypeThree struct {
	TypeThreeStringFieldFiltered     string     `json:"type_three_string_field_filtered"`
	TypeThreeNumberFieldFiltered     int        `json:"type_three_number_field_filtered"`
	TypeThreeTimeFieldFiltered       time.Time  `json:"type_three_time_field_filtered"`
	TypeThreeBooleanFieldFiltered    bool       `json:"type_three_boolean_field_filtered"`
	TypeThreeStringFieldWithNoFilter string     `json:"type_three_string_field_with_no_filter"`
	TypeThreeNumberFieldWithNoFilter int        `json:"type_three_number_field_with_no_filter"`
	TypeThreeTimeFieldWithNoFilter   time.Time  `json:"type_three_time_field_with_no_filter"`
	TypeUnion                        UnionOne   `json:"type_union"`
	TypeUnionSlice                   []UnionOne `json:"type_union_slice,omitempty"`
	TypeUnionNotMandatory            UnionOne   `json:"type_union_not_mandatory,omitempty"`
}

type TypeTwo struct {
	TypeTwoStringFieldFiltered       string     `json:"type_two_string_field_filtered"`
	TypeTwoNumberFieldFiltered       int        `json:"type_two_number_field_filtered"`
	TypeTwoTimeFieldFiltered         time.Time  `json:"type_two_time_field_filtered"`
	TypeTwoBooleanFieldFiltered      bool       `json:"type_two_boolean_field_filtered"`
	TypeTwoStringFieldWithNoFilter   string     `json:"type_twoString_field_with_no_filter"`
	TypeTwoNumberFieldWithNoFilter   int        `json:"type_twoNumber_field_with_no_filter"`
	TypeTwoTimeFieldWithNoFilter     time.Time  `json:"type_twoTime_field_with_no_filter"`
	TypeTwoSliceWithTypeTwos         []*TypeTwo `json:"type_two_slice_with_type_twos,omitempty"`
	TypeTwoWithTypeThree             *TypeThree `json:"type_two_with_type_three"`
	TypeTwoWithTypeThreeNotMandatory *TypeThree `json:"type_two_with_type_three_not_mandatory,omitempty"`
}

type UnionTypeFour struct {
	TypeIntUnionTwo    *int       `json:"type_int_union_two,omitempty"`
	TypeStringUnionTwo *string    `json:"type_string_union_two,omitempty"`
	TypeTimeUnionTwo   *time.Time `json:"type_time_union_two,omitempty"`
}

func (UnionTypeFour) IsUnionTwo() {}

type UnionTypeOne struct {
	TypeIntUnionOne    *int        `json:"type_int_union_one,omitempty"`
	TypeStringUnionOne *string     `json:"type_string_union_one,omitempty"`
	TypeTimeUnionOne   *time.Time  `json:"type_time_union_one,omitempty"`
	TypeNested         *NestedType `json:"type_nested"`
}

func (UnionTypeOne) IsUnionOne() {}

type UnionTypeThree struct {
	TypeIntUnionOne    *int       `json:"type_int_union_one,omitempty"`
	TypeStringUnionOne *string    `json:"type_string_union_one,omitempty"`
	TypeTimeUnionOne   *time.Time `json:"type_time_union_one,omitempty"`
}

func (UnionTypeThree) IsUnionTwo() {}

type UnionTypeTwo struct {
	TypeIntUnionTwo         *int       `json:"type_int_union_two,omitempty"`
	TypeStringUnionTwo      *string    `json:"type_string_union_two,omitempty"`
	TypeTimeUnionTwo        *time.Time `json:"type_time_union_two,omitempty"`
	TypeStringSliceUnionTwo []string   `json:"type_string_slice_union_two,omitempty"`
}

func (UnionTypeTwo) IsUnionOne() {}

type UnionTypeTwoPrime struct {
	TypeIntUnionTwoPrime    *int       `json:"type_int_union_two_prime,omitempty"`
	TypeStringUnionTwoPrime *string    `json:"type_string_union_two_prime,omitempty"`
	TypeTimeUnionTwoPrime   *time.Time `json:"type_time_union_two_prime,omitempty"`
}

func (UnionTypeTwoPrime) IsUnionOne() {}

type FilterableExtra string

const (
	// Get minimum and maximum value used on all the filters for this field.
	// Useful when you need to do a range query for performance reasons.
	FilterableExtraMinmax FilterableExtra = "MINMAX"
)

var AllFilterableExtra = []FilterableExtra{
	FilterableExtraMinmax,
}

func (e FilterableExtra) IsValid() bool {
	switch e {
	case FilterableExtraMinmax:
		return true
	}
	return false
}

func (e FilterableExtra) String() string {
	return string(e)
}

func (e *FilterableExtra) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterableExtra(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterableExtra", str)
	}
	return nil
}

func (e FilterableExtra) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
