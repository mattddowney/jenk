package cmd

import "unicode"

func capitalize(input string) string {
	inputRune := []rune(input)
	inputRune[0] = unicode.ToUpper(inputRune[0])

	return string(inputRune)
}
