package number_to_words_test

import (
	"testing"

	"github.com/ashlinchak/number_to_words"
	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	assert.Equal(t, "zero", number_to_words.IntToWords(0))
}
func TestMinus(t *testing.T) {
	assert.Equal(t, "minus one", number_to_words.IntToWords(-1))
}

func TestPositive(t *testing.T) {
	assert.Equal(t, "one", number_to_words.IntToWords(1))
}

func TestLessThenTwenty(t *testing.T) {
	assert.Equal(t, "eleven", number_to_words.IntToWords(11))
	assert.Equal(t, "twelve", number_to_words.IntToWords(12))
	assert.Equal(t, "thirteen", number_to_words.IntToWords(13))
	assert.Equal(t, "fourteen", number_to_words.IntToWords(14))
	assert.Equal(t, "fifteen", number_to_words.IntToWords(15))
	assert.Equal(t, "sixteen", number_to_words.IntToWords(16))
	assert.Equal(t, "seventeen", number_to_words.IntToWords(17))
	assert.Equal(t, "eighteen", number_to_words.IntToWords(18))
	assert.Equal(t, "nineteen", number_to_words.IntToWords(19))
}

func TestTenths(t *testing.T) {
	assert.Equal(t, "ten", number_to_words.IntToWords(10))
	assert.Equal(t, "twenty", number_to_words.IntToWords(20))
	assert.Equal(t, "thirty", number_to_words.IntToWords(30))
	assert.Equal(t, "forty", number_to_words.IntToWords(40))
	assert.Equal(t, "fifty", number_to_words.IntToWords(50))
	assert.Equal(t, "sixty", number_to_words.IntToWords(60))
	assert.Equal(t, "seventy", number_to_words.IntToWords(70))
	assert.Equal(t, "eighty", number_to_words.IntToWords(80))
	assert.Equal(t, "ninety", number_to_words.IntToWords(90))
	assert.Equal(t, "forty-four", number_to_words.IntToWords(44))
	assert.Equal(t, "fifty-seven", number_to_words.IntToWords(57))
	assert.Equal(t, "ninety-nine", number_to_words.IntToWords(99))
}

func TestLessThenThousand(t *testing.T) {
	assert.Equal(t, "one hundred", number_to_words.IntToWords(100))
	assert.Equal(t, "one hundred twenty-nine", number_to_words.IntToWords(129))
	assert.Equal(t, "two hundred", number_to_words.IntToWords(200))
	assert.Equal(t, "four hundred", number_to_words.IntToWords(400))
	assert.Equal(t, "six hundred one", number_to_words.IntToWords(601))
	assert.Equal(t, "nine hundred eighty", number_to_words.IntToWords(980))
	assert.Equal(t, "nine hundred ninety-eight", number_to_words.IntToWords(998))
}

func TestLessThenMillion(t *testing.T) {
	assert.Equal(t, "one thousand", number_to_words.IntToWords(1000))
	assert.Equal(t, "ten thousand", number_to_words.IntToWords(10000))
	assert.Equal(t, "one hundred ten thousand", number_to_words.IntToWords(110000))
	assert.Equal(t, "one hundred twenty-three thousand", number_to_words.IntToWords(123000))
	assert.Equal(t, "two hundred forty-five thousand six hundred", number_to_words.IntToWords(245600))
	assert.Equal(t, "three hundred twelve thousand eight hundred ninety", number_to_words.IntToWords(312890))
	assert.Equal(t, "nine hundred ninety-nine thousand nine hundred ninety-nine", number_to_words.IntToWords(999999))
	assert.Equal(t, "eight hundred four thousand two", number_to_words.IntToWords(804002))
}

func TestLessThenBillion(t *testing.T) {
	assert.Equal(t, "one million", number_to_words.IntToWords(1000000))
	assert.Equal(t, "ten million", number_to_words.IntToWords(10000000))
	assert.Equal(t, "one hundred ten million", number_to_words.IntToWords(110000000))
	assert.Equal(t, "one hundred twenty-three million", number_to_words.IntToWords(123000000))
	assert.Equal(t, "two hundred forty-five million six hundred thousand", number_to_words.IntToWords(245600000))
	assert.Equal(t, "three hundred twelve million eight hundred ninety thousand", number_to_words.IntToWords(312890000))
	assert.Equal(t, "three hundred twelve million eight hundred ninety-four thousand", number_to_words.IntToWords(312894000))
	assert.Equal(t, "seven hundred fifteen million eight hundred ninety-four thousand eight hundred",
		number_to_words.IntToWords(715894800))
	assert.Equal(t, "two hundred twenty-five million nine hundred ninety-five thousand eight hundred thirty",
		number_to_words.IntToWords(225995830))
	assert.Equal(t, "five hundred sixty-eight million five hundred ninety-five thousand eight hundred forty-two",
		number_to_words.IntToWords(568595842))
	assert.Equal(t, "five hundred million one hundred thousand five hundred two", number_to_words.IntToWords(500100502))
}

func TestLessThenTrillion(t *testing.T) {
	assert.Equal(t, "one billion", number_to_words.IntToWords(1000000000))
	assert.Equal(t, "ten billion", number_to_words.IntToWords(10000000000))
	assert.Equal(t, "one hundred billion", number_to_words.IntToWords(100000000000))
	assert.Equal(t, "two hundred thirty billion", number_to_words.IntToWords(230000000000))
	assert.Equal(t, "four hundred fifty-five billion", number_to_words.IntToWords(455000000000))
	assert.Equal(t, "nine hundred fifty-five billion six hundred million", number_to_words.IntToWords(955600000000))
	assert.Equal(t, "six hundred seventy-six billion six hundred seventy million", number_to_words.IntToWords(676670000000))
	assert.Equal(t, "three hundred thirty-six billion six hundred seventy-eight million", number_to_words.IntToWords(336678000000))
	assert.Equal(t, "three hundred thirty-six billion six hundred seventy-eight million two hundred thousand",
		number_to_words.IntToWords(336678200000))
	assert.Equal(t, "one hundred thirty-six billion six hundred seventy-eight million two hundred eighty thousand",
		number_to_words.IntToWords(136678280000))
	assert.Equal(t, "one hundred thirty-six billion six hundred seventy-eight million two hundred eighty-five thousand",
		number_to_words.IntToWords(136678285000))
	assert.Equal(t, "six hundred sixty-six billion six hundred seventy-eight million two hundred eighty-five thousand four hundred",
		number_to_words.IntToWords(666678285400))
	assert.Equal(t, "four hundred fifty-five billion six hundred seventy-eight million two hundred eighty-five thousand four hundred eighty",
		number_to_words.IntToWords(455678285480))
	assert.Equal(t, "four hundred fifty-five billion six hundred seventy-eight million two hundred eighty-five thousand four hundred eighty-one",
		number_to_words.IntToWords(455678285481))
	assert.Equal(t, "one billion one", number_to_words.IntToWords(1000000001))

}

func TestLessThenQuadrillion(t *testing.T) {
	assert.Equal(t, "one trillion", number_to_words.IntToWords(1000000000000))
	assert.Equal(t, "one hundred trillion one", number_to_words.IntToWords(100000000000001))
	assert.Equal(t, "nine hundred ninety-nine trillion nine hundred ninety-nine billion "+
		"nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		number_to_words.IntToWords(999999999999999))
}

func TestLessThenQuintillion(t *testing.T) {
	assert.Equal(t, "one quadrillion", number_to_words.IntToWords(1000000000000000))
	assert.Equal(t, "one quadrillion ten million one", number_to_words.IntToWords(1000000010000001))
	assert.Equal(t, "nine hundred ninety-nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion "+
		"nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		number_to_words.IntToWords(999999999999999999))
}

func TestLessThenMaxUint32(t *testing.T) {
	assert.Equal(t, "one quintillion", number_to_words.IntToWords(1000000000000000000))
	assert.Equal(t, "nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion "+
		"thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven",
		number_to_words.IntToWords(9223372036854775807))
}
