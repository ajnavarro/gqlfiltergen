package gqlfiltergen

import (
	_ "embed"
	"errors"
	"os/exec"
	"testing"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/stretchr/testify/require"

	"github.com/ajnavarro/gqlfiltergen/testdata/out"
)

//go:embed testdata/injectSchema.graphql
var injectedSchema string

func TestFiltersGeneration(t *testing.T) {
	cfg, err := config.LoadConfig("testdata/gqlgen.yml")
	require.NoError(t, err)

	p := NewPlugin(&Options{
		InjectCodeAfter: injectedSchema,
	})

	err = api.Generate(cfg,
		api.AddPlugin(p),
	)
	require.NoError(t, err)

	require.NoError(t, goBuild(t, "./testdata/out/"))

	nexists := false

	eqlValue := "eqValue"
	likeValue := "^t.*$"
	eqlValueInt := 42
	minmax1 := 2222

	f1 := &out.FilterTypeOne{
		TypeOneStringFieldFiltered: &out.FilterString{
			Eq: &eqlValue,
		},
		Or: []*out.FilterTypeOne{
			{
				Or: []*out.FilterTypeOne{
					{
						TypeOneNumberFieldFiltered: &out.FilterInt{
							Eq: &eqlValueInt,
						},
					},
					{
						TypeOneNumberFieldFiltered: &out.FilterInt{
							Exists: &nexists,
						},
					},
					{
						TypeOneNumberFieldFiltered: &out.FilterInt{
							Eq: &minmax1,
						},
					},
				},
			},
			{
				TypeOneStringFieldFilteredNotMandatory: &out.FilterString{
					Like: &likeValue,
				},
			},
		},
	}

	likeVal := "this starts with t"
	nlikeVal := "but this doesn't"

	ts := []*out.TypeOne{
		{
			TypeOneStringFieldFiltered:             eqlValue,
			TypeOneStringFieldFilteredNotMandatory: &likeVal,
		},
		{
			TypeOneStringFieldFiltered: eqlValue,
			TypeOneNumberFieldFiltered: 23,
		},
		{
			TypeOneStringFieldFiltered: eqlValue,
			TypeOneNumberFieldFiltered: 42,
		},
		{
			TypeOneStringFieldFiltered:             eqlValue,
			TypeOneStringFieldFilteredNotMandatory: &nlikeVal,
		},
		{
			TypeOneStringFieldFiltered: "othervalue",
		},
		{
			TypeOneStringFieldFiltered: eqlValue,
		},
	}

	results := []bool{
		true, true, true, false, false, true,
	}

	for i, typ := range ts {
		require.Equal(t, results[i], f1.Eval(typ), "position %d", i)
	}

	min, max := f1.MinMaxTypeOneNumberFieldFiltered()

	require.Equal(t, &eqlValueInt, min)
	require.Equal(t, &minmax1, max)

	// test union

	valOne := "valOne"
	valTwo := "valTwo"

	f2 := &out.FilterTypeThree{
		TypeUnionSlice: &out.NestedFilterUnionOne{
			UnionTypeOne: &out.NestedFilterUnionTypeOne{
				TypeStringUnionOne: &out.FilterString{
					Eq: &valOne,
				},
			},
			UnionTypeTwo: &out.NestedFilterUnionTypeTwo{
				TypeStringUnionTwo: &out.FilterString{
					Eq: &valTwo,
				},
			},
		},
	}

	ignored := "val MUST BE IGNORED"

	tus := &out.TypeThree{
		TypeUnionSlice: []out.UnionOne{
			out.UnionTypeTwo{
				TypeStringUnionTwo: &valTwo,
			},
			out.UnionTypeOne{
				TypeStringUnionOne: &valOne,
			},
			out.UnionTypeTwoPrime{
				TypeStringUnionTwoPrime: &ignored,
			},
		},
	}

	require.Equal(t, true, f2.Eval(tus))
}

func goBuild(t *testing.T, path string) error {
	t.Helper()
	cmd := exec.Command("go", "build", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}

	return nil
}
