package dice

import (
	"regexp"
	"strings"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`[^a-z0-9âäàéêèëîïöôù]`)
}

// DiceString if a wrapper around a string containing bigrams and number of bigrams
type DiceString struct {
	Value   string
	Bigrams map[string]bool
	Size    int
}

// NewDiceString generates a DiceString from a simple String
func NewDiceString(value *string) *DiceString {
	newValue := re.ReplaceAllString(strings.ToLower(*value), "")
	size := 0
	bigrams := make(map[string]bool)

	for i := 0; i < len(newValue)-1; i++ {
		size++
		bigrams[newValue[i:i+2]] = true
	}

	return &DiceString{Value: newValue, Bigrams: bigrams, Size: size}
}

// CompareDiceString do the same as Compare but as a method for a DiceString
func (first *DiceString) CompareDiceString(second *DiceString) float64 {

	if first.Size == 0 && second.Size == 0 {
		return 1
	}

	common := 0.0

	for i := 0; i < len(second.Value)-1; i++ {
		if _, ok := first.Bigrams[second.Value[i:i+2]]; ok {
			common += 1.0
		}
	}

	return 2.0 * common / (float64(first.Size) + float64(second.Size))

}

// Compare() returns the dice coef of 2 strings
func Compare(first, second *string) float64 {

	//first lower the two strings and avoid special characters
	newFirst := re.ReplaceAllString(strings.ToLower(*first), "")
	newSecond := re.ReplaceAllString(strings.ToLower(*second), "")

	//rapid returns 1 (best score) if the two cleaned strings are the same
	if newFirst == newSecond {
		return 1
	}

	//rapid returns 0 (worst score) if one of the two cleaned strings is empty
	if newFirst == "" || newSecond == "" {
		return 0
	}

	mapping := make(map[string]bool) //init a map of existing bigrams of the first string
	commonBigrams := 0.0             //init a commun bigram score between the two strings
	totalBigrams := 0.0              //init the total bigrams of the two strings

	//generate the map of bigram of the first string, and increment the total bigram
	for i := 0; i < len(newFirst)-1; i++ {
		totalBigrams += 1.0
		mapping[newFirst[i:i+2]] = true
	}

	//check for all bigrams of the second string if they can be found in the map,
	//incrementing common bigram if found (and total bigram)
	for i := 0; i < len(newSecond)-1; i++ {
		totalBigrams += 1.0
		if _, ok := mapping[newSecond[i:i+2]]; ok {
			commonBigrams += 1.0
		}
	}

	return 2.0 * commonBigrams / totalBigrams
}
