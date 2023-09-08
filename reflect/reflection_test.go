package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"mach"},
			[]string{"mach"},
		}, {
			"Struct with two string field",
			struct {
				Name string
				City string
			}{
				"mach",
				"China",
			},
			[]string{"mach", "China"},
		}, {
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{
				"mach", 22,
			},
			[]string{"mach"},
		}, {
			"Nested fields",
			Person{
				"mach",
				Profile{22, "GuangDong"},
			},
			[]string{"mach", "GuangDong"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
