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
	backslash := "\\"
	escape := false

	r := []rune(str)

	for i := 0; i < len(r); i++ {
		if string(r[i]) == backslash {
			if escape {
				if i+1 != len(r) {
					if digit, err := strconv.Atoi(string(r[i+1])); err == nil {
						b.WriteString(strings.Repeat(string(r[i]), digit))
					} else {
						b.WriteRune(r[i])
					}
				} else {
					b.WriteRune(r[i])
				}
				escape = false
			} else {
				escape = true
				continue
			}
		}
		if unicode.IsDigit(r[i]) {
			if escape {
				if i+1 != len(r) {
					if digit, err := strconv.Atoi(string(r[i+1])); err == nil {
						b.WriteString(strings.Repeat(string(r[i]), digit))
					} else {
						b.WriteRune(r[i])
					}
				} else {
					b.WriteRune(r[i])
				}
				escape = false
			} else {
				if i+1 != len(r) {
					if i == 0 || unicode.IsDigit(r[i+1]) {
						return "", ErrInvalidString
					}
				}
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
	if escape {
		return "", ErrInvalidString
	}
	return b.String(), nil
}
