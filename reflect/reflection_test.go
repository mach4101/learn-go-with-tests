package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

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
		}, {
			"Pointer fields",
			&Person{
				"mach",
				Profile{22, "GuangDong"},
			},
			[]string{"mach", "GuangDong"},
		}, {
			"slice",
			[]Profile{
				{23, "GuangDong"},
				{21, "LiaoNing"},
			},
			[]string{"GuangDong", "LiaoNing"},
		},
		{
			"Array",
			[2]Profile{
				{23, "GuangDong"},
				{21, "LiaoNing"},
			},
			[]string{"GuangDong", "LiaoNing"},
		}, {
			"Map",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
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

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didn't", haystack, needle)
	}
}
