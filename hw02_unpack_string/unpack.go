package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/example/stringutil"
)

var ErrInvalidString = errors.New("invalid string")

const slash string = "slash"
const digit string = "digit"
const letter string = "letter"
const stop string = "stop"

func makeConf(confString string) []string {
	var confSlice []string

	for _, i := range confString {
		switch {
		case string(i) == `\`:
			confSlice = append(confSlice, slash)
		case unicode.IsDigit(i):
			confSlice = append(confSlice, digit)
		case unicode.IsLetter(i):
			confSlice = append(confSlice, letter)
		}
	}
	confSlice = append(confSlice, stop)
	return confSlice
}

func Unpack(oldString string) (string, error) {
	var result strings.Builder
	reverseString := stringutil.Reverse(oldString)
	sliceString := strings.Split(reverseString, "")
	confSlice := makeConf(reverseString)
	koff := 0

	for index := range sliceString {
		index += koff
		if index < len(sliceString) {
			switch {
			case confSlice[len(oldString)-1] == digit:
				return "", ErrInvalidString
			case confSlice[index] == digit && confSlice[index+1] == digit:
				return "", ErrInvalidString
			case confSlice[index] == digit && confSlice[index+1] == slash:
				result.WriteString(string(sliceString[index]))
			case confSlice[index] == slash:
				continue
			case confSlice[index] == letter && confSlice[index+1] == digit:
				result.WriteString(string(sliceString[index]))
			case confSlice[index] == digit && confSlice[index+1] == letter:
				count, _ := strconv.Atoi(sliceString[index])
				result.WriteString(strings.Repeat(string(sliceString[index+1]), count))
				koff++
			case (confSlice[index] == letter || confSlice[index] == digit) && confSlice[index+1] == stop:
				result.WriteString(string(sliceString[index]))
				break
			case confSlice[index] == letter && confSlice[index+1] == letter:
				result.WriteString(string(sliceString[index]))
			}
		}
	}
	return stringutil.Reverse(result.String()), nil
}
