package number_to_words

import (
	"math"
	"strings"
)

// delimiters
var delimiter = " "
var delimiterForTenths = "-"

// number names
var minus = "minus"
var hundred = "hundred"
var thousand = "thousand"
var million = "million"
var billion = "billion"
var trillion = "trillion"
var quadrillion = "quadrillion"
var quintillion = "quintillion"

// digits
const (
	Ten            = 10
	OneHundred     = 100
	OneThousand    = 1000
	OneMillion     = 1000000
	OneBillion     = 1000000000
	OneTrillion    = 1000000000000
	OneQuadrillion = 1000000000000000    // 10**15
	OneQuintillion = 1000000000000000000 // 10**18
	MaxInt         = 9223372036854775807 // math.MaxInt64 (18)
)

var lessThenTwenty = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}
var tenths = []string{
	"ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

func IntToWords(number int) string {
	words := getWordsFromInt(number)
	return strings.Join(words, delimiter)
}

func getWordsFromInt(number int) []string {
	words := []string{}
	var reminder int

	if number == 0 {
		return append(words, lessThenTwenty[0])
	}

	if number < 0 {
		words = append(words, minus)
		number = -number
	}

	if number < 20 {
		words = append(words, lessThenTwenty[number])
	} else if number < OneHundred {
		reminder = number % Ten
		mathFloor := int(math.Floor(float64(number / Ten)))
		numWord := tenths[mathFloor-1]
		if reminder > 0 {
			numWord += delimiterForTenths + lessThenTwenty[reminder]
			reminder = 0
		}
		words = append(words, numWord)
	} else if number < OneThousand {
		calWords(number, OneHundred, hundred, &words)
	} else if number < OneMillion {
		calWords(number, OneThousand, thousand, &words)
	} else if number < OneBillion {
		calWords(number, OneMillion, million, &words)
	} else if number < OneTrillion {
		calWords(number, OneBillion, billion, &words)
	} else if number < OneQuadrillion {
		calWords(number, OneTrillion, trillion, &words)
	} else if number < OneQuintillion {
		calWords(number, OneQuadrillion, quadrillion, &words)
	} else if number <= MaxInt {
		calWords(number, OneQuintillion, quintillion, &words)
	}

	return words
}

func calWords(number, devider int, numberName string, words *[]string) {
	reminder := number % devider
	mathFloor := int(math.Floor(float64(number / devider)))
	hWords := getWordsFromInt(mathFloor)
	*words = append(*words, hWords...)
	*words = append(*words, numberName)
	if reminder > 0 {
		dWords := getWordsFromInt(reminder)
		*words = append(*words, dWords...)
	}
}
