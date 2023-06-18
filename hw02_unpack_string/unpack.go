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
	// fmt.Println(s)
	lenString := len(s)

	for index, r := range s {
		
		// Проверяем - является ли первый символ в строке цифрой, если да - возвращаем ошибку
		if unicode.IsDigit(r) {
			if prev == 0 {
				return "", ErrInvalidString
			} else {
				// Если идут две цифры подряд - возвращам ошибку
				if unicode.IsDigit(r) && unicode.IsDigit(prev) {
					return "", ErrInvalidString
				}
				if string(prev) != `\` {
					count, _ = strconv.Atoi(string(r))
					// обрабатываем случай появления цифры 0 в строке
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

			}

		} else {
			// Записываем литералы в новую строку. Чтобы избежать появления нечитаемых символов в строке вводим условие с index > 0
			// т.к. при index == 0 условие unicode.IsDigit(prev) == false будет верным, но prev в этом случае будет в состоянии ZeroValue
			if unicode.IsDigit(prev) == false && index > 0 {
				if string(r) != `\` {
					result.WriteRune(prev)
					prev = r
				} else {
					if string(prev) == `\` && unicode.IsDigit(r) {
						result.WriteRune(r)
						prev = r
					} else {
						if unicode.IsDigit(prev) && string(r) == `\`{
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
			// Запись последней руны в строке
			if lenString == index+1 {
				result.WriteRune(r)
			}

		}
	}
	return result.String(), nil
}	
