package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	isDigit := regexp.MustCompile(`[0-9][0-9]`)
	containsDigit := regexp.MustCompile(`[0-9]`)
	startsWithDigit := regexp.MustCompile(`^[0-9]`)
	var result strings.Builder
	var repeat bool

	if isDigit.MatchString(input) || startsWithDigit.MatchString(input) {
		return "", ErrInvalidString
	}

	runes := []rune(input)

	if !containsDigit.MatchString(input) || len(runes) == 1 {
		return input, nil
	}

	for i := 0; i < len(runes)-1; i++ {
		repeat = false

		if unicode.IsDigit(runes[i+1]) {
			repeat = true
			max, _ := strconv.Atoi(string(runes[i+1]))
			for j := 0; j < max; j++ {
				result.WriteRune(runes[i])
			}
			i++
		} else {
			result.WriteRune(runes[i])
		}
	}
	if repeat && !containsDigit.MatchString(string(runes[len(runes)-1])) {
		result.WriteRune(runes[len(runes)-1])
	}

	return result.String(), nil
}
