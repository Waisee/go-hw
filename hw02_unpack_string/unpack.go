package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var b strings.Builder
	r := []rune(str)

	for i := 0; i < len(r); i++ {
		if unicode.IsDigit(r[i]) {
			if i == 0 || unicode.IsDigit(r[i+1]) {
				return "", ErrInvalidString
			}
		}
		if unicode.IsLetter(r[i]) || unicode.IsSpace(r[i]) {
			if i+1 != len(r) {
				if digit, err := strconv.Atoi(string(r[i+1])); err == nil {
					b.WriteString(strings.Repeat(string(r[i]), digit))
				} else {
					b.WriteString(string(r[i]))
				}
			} else {
				b.WriteString(string(r[i]))
			}
		}
	}
	return b.String(), nil
}
