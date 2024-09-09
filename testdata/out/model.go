package out

import "time"

type ExternalType struct {
}

func (t *ExternalType) NumberOne() int {
	return 42
}

func (t *ExternalType) NumberTwo() int64 {
	return 11
}

func (t *ExternalType) NumberThree() int32 {
	return 22
}

func (t *ExternalType) NumberFour() *int {
	v := 33
	return &v
}

func (t *ExternalType) NumberFive() *int32 {
	v := int32(33)
	return &v
}

func (t *ExternalType) NumberList() []int {
	return []int{1, 2, 3, 4}
}

func (t *ExternalType) TypeOne() *TypeOne {
	strNotMandatory := "fieldNotMandatory"
	nmrNotMandatory := 11
	return &TypeOne{
		TypeOneStringFieldFiltered:             "stringFieldFiltered",
		TypeOneStringFieldFilteredNotMandatory: &strNotMandatory,
		TypeOneNumberFieldFiltered:             42,
		TypeOneNumberFieldFilteredNotMandatory: &nmrNotMandatory,
		TypeOneTimeFieldFiltered:               time.Now(),
		TypeOneBooleanFiltered:                 true,
		TypeOneStringFieldWithNoFilter:         "no filter",
		TypeOneSliceWithTypeTwos:               nil,
	}
}
