package nato

import (
	"testing"
)

func BenchmarkPhonetic(b *testing.B) {
	benches := []struct {
		Name string
		F    func(string) string
	}{
		{Name: "StringMap", F: PhoneticStringMap},
		{Name: "RuneMap", F: PhoneticRuneMap},
		{Name: "Slice", F: PhoneticSlice},
		{Name: "Array", F: PhoneticArray},
	}

	testData := "WTF wtf ABCDEFG testing testing, one, two, three."

	for _, bench := range benches {
		b.Run(bench.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bench.F(testData)
			}
		})
	}
}

func TestPhoenticImplementations(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyz0123456789"
	expected := "alfa bravo charlie delta echo foxtrot golf hotel india juliett kilo lima mike november oscar papa quebec romeo sierra tango uniform victor whiskey x-ray yankee zulu zero one two three four five six seven eight nine"

	funcs := []struct {
		Name string
		F    func(string) string
	}{
		{Name: "StringMap", F: PhoneticStringMap},
		{Name: "RuneMap", F: PhoneticRuneMap},
		{Name: "Slice", F: PhoneticSlice},
		{Name: "Array", F: PhoneticArray},
	}

	for _, test := range funcs {
		t.Run(test.Name, func(t *testing.T) {
			if result := test.F(input); result != expected {
				t.Fatalf("%s returned %q, expected %q", test.Name, result, expected)
			}
		})
	}
}
