package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age int
	City string
}

type Person struct {
	Name string
	Profile Profile
}

func TestWalk(t *testing.T) {
	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{string}{"Oskar"},
			[]string{"Oskar"},
		},
		{
			"struct with two string fields",
			struct{
				Name string
				City string
			}{"Oskar", "Manchester"},
			[]string{"Oskar", "Manchester"},
		},
		{
			"struct with a string and number",
			struct{
				Name string
				Age int
			}{"Oskar", 30},
			[]string{"Oskar"},
		},
		{
			"nested fields",
			Person{
				"Oskar",
				Profile{30, "Manchester"},
			},
			[]string{"Oskar", "Manchester"},
		},
		{
			"pointers to things",
			&Person{
				"Oskar",
				Profile{30, "Manchester"},
			},
			[]string{"Oskar", "Manchester"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Manchester"},
			},
			[]string{"London", "Manchester"},
		},
	}

	for _,test := range cases {
		t.Run(test.Name, func (t *testing.T) {
			var got []string

			walk(test.Input, func(val string) {
				got = append(got, val)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}


		})
	}
}