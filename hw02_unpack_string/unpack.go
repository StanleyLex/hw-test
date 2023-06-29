package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func ErrCheck(chars string) error {
	var prev rune

	for index, char := range chars {
		switch {
		case unicode.IsDigit(char) && index == 0:
			return ErrInvalidString
		case unicode.IsDigit(char) && unicode.IsDigit(prev):
			return ErrInvalidString
		default:
			prev = char
		}
	}
	return nil
}

func Unpack(chars string) (string, error) {
	var result strings.Builder
	var prev rune

	if err := ErrCheck(chars); err != nil {
		return "", err
	}

	for _, char := range chars {
		switch {
		case unicode.IsLetter(char) && prev == 0:
			result.WriteRune(char)
			prev = char
		case unicode.IsDigit(char) && unicode.IsLetter(prev):
			count, _ := strconv.Atoi(string(char))
			if count != 0 {
				result.WriteString(strings.Repeat(string(prev), count-1))
				prev = char
			} else {
				bytes := []byte(result.String())
				result.Reset()
				result.Write(bytes[:len(bytes)-1])
				prev = char
			}
		case unicode.IsLetter(char) && unicode.IsLetter(prev):
			result.WriteRune(char)
			prev = char
		case unicode.IsLetter(char) && unicode.IsDigit(prev):
			result.WriteRune(char)
			prev = char
		case unicode.IsDigit(char) && string(prev) == `\`:
			result.WriteRune(char)
		case string(char) == `\` && unicode.IsDigit(prev):
			prev = char
		case string(char) == `\` && unicode.IsLetter(prev):
			prev = char
		}
	}

	return result.String(), nil
}
