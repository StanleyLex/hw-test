package hw02unpackstring
import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var prev rune
	var count int
	lenString := len(s)

	for index, r := range s {
		if unicode.IsDigit(r) {
			if prev == 0 {
				return "", ErrInvalidString
			}
			if unicode.IsDigit(r) && unicode.IsDigit(prev) {
				return "", ErrInvalidString
			}
			if string(prev) != `\` {
				count, _ = strconv.Atoi(string(r))
				if count != 0 {
					result.WriteString(strings.Repeat(string(prev), count))
					prev = r
					count = 0
				}
				prev = r
			} else {
				result.WriteString(string(r))
				prev = r
			}
		} else {
			if unicode.IsDigit(prev) == false && index > 0 {
				if string(r) != `\` {
					result.WriteRune(prev)
					prev = r
				} else {
					if string(prev) == `\` && unicode.IsDigit(r) {
						result.WriteRune(r)
						prev = r
					} else {
						if unicode.IsDigit(prev) && string(r) == `\` {
							prev = r
						} else {
							result.WriteRune(prev)
							prev = r
						}
					}
				}
			} else {
				prev = r
			}
			if lenString == index+1 {
				result.WriteRune(r)
			}
		}
	}
	return result.String(), nil
}
