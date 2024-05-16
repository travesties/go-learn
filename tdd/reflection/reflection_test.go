package reflection

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Travis"},
			[]string{"Travis"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Travis", "Baltimore"},
			[]string{"Travis", "Baltimore"},
		},
		{
			"struct with one string field and one int field",
			struct {
				Name string
				Age  int
			}{"Travis", 32},
			[]string{"Travis"},
		},
		{
			"struct with nested fields",
			Person{
				"Travis",
				Profile{32, "Baltimore"},
			},
			[]string{"Travis", "Baltimore"},
		},
		{
			"pointers to things",
			&Person{
				"Travis",
				Profile{32, "Baltimore"},
			},
			[]string{"Travis", "Baltimore"},
		},
		{
			"slices",
			[]Profile{
				{32, "Baltimore"},
				{37, "San Francisco"},
			},
			[]string{"Baltimore", "San Francisco"},
		},
		{
			"arrays",
			[2]Profile{
				{32, "Baltimore"},
				{37, "San Francisco"},
			},
			[]string{"Baltimore", "San Francisco"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	// A separate test is necessary for maps because maps do not
	// guarantee order.
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Paris"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Paris"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Paris"}
		}

		var got []string
		want := []string{"Berlin", "Paris"}

		walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
