package dice

import "testing"

var testCases = []struct {
	description string
	first       string
	second      string
	expected    float64
}{
	{
		"empty strings gives 1",
		"",
		"",
		1.0,
	},
	{
		"only one empty string gives 0",
		"",
		"foo",
		0.0,
	},
	{
		"only one empty string gives 0",
		"foo",
		"",
		0.0,
	},
	{
		"dice doesn't care with spaces",
		"      ",
		"",
		1.0,
	},
	{
		"dice doesn't care with spaces",
		"      ",
		"    ",
		1.0,
	},
	{
		"dice doesn't care with special characters",
		"    **_`'",
		"@@@##-_",
		1.0,
	},
	{
		"sames strings give 1",
		"foo",
		"foo",
		1.0,
	},
	{
		"dice doesn't care with case sensitivity",
		"foo",
		"fOO",
		1.0,
	},
	{
		"dice returns coef of commun bigrams",
		"foo",
		"foa",
		0.5,
	},
	{
		"dice returns coef of commun bigrams",
		"fOo",
		"foA",
		0.5,
	},
	{
		"dice returns coef of commun bigrams",
		"nacht",
		"nicht",
		0.5,
	},
	{
		"dice returns coef of commun bigrams",
		"one",
		"two",
		0.0,
	},
	{
		"dice returns coef of commun bigrams with different lengths of words",
		"a letter",
		"one letter",
		10.0 / 14.0,
	},
	{
		"dice returns coef of commun bigrams with different lengths of words",
		"more complex word",
		"more big word",
		12.0 / 24.0,
	},
}

func TestDice(t *testing.T) {
	for _, test := range testCases {
		if result := Compare(&test.first, &test.second); result != test.expected {
			t.Fatalf("%s failed. get %v want %v", test.description, result, test.expected)
		}
	}
}

func TestDiceWithStruct(t *testing.T) {
	for _, test := range testCases {
		first := NewDiceString(&test.first)
		second := NewDiceString(&test.second)
		if result := first.CompareDiceString(second); result != test.expected {
			t.Fatalf("%s failed. get %v want %v", test.description, result, test.expected)
		}
	}
}
