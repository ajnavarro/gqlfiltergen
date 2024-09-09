package gqlfiltergen

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/stretchr/testify/require"

	"github.com/ajnavarro/gqlfiltergen/testdata/out"
)

func TestFiltersGeneration(t *testing.T) {
	cfg, err := config.LoadConfig("testdata/gqlgen.yml")
	require.NoError(t, err)

	p := NewPlugin(&Options{
		Queries: []string{
			"testFilter(filter: FilterTypeOne!): [TypeOne!]",
		},
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

	f1 := &out.FilterTypeOne{
		TypeOneStringFieldFiltered: &out.FilterString{
			Eq: &eqlValue,
		},
		Or: []*out.FilterTypeOne{
			{
				Or: []*out.FilterTypeOne{
					{
						TypeOneNumberFieldFiltered: &out.FilterNumber{
							Eq: &eqlValueInt,
						},
					},
					{
						TypeOneNumberFieldFiltered: &out.FilterNumber{
							Exists: &nexists,
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
