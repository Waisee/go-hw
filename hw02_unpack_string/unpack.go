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
	escape := false

	r := []rune(str)

	for i := 0; i < len(r); i++ {
		switch getType(r[i]) {
		case "backslash":
			if escape {
				if i+1 != len(r) {
					checkAndWrite(&b, r[i], r[i+1])
				} else {
					b.WriteRune(r[i])
				}
				escape = false
			} else {
				escape = true
				continue
			}
		case "digit":
			if escape {
				if i+1 != len(r) {
					checkAndWrite(&b, r[i], r[i+1])
				} else {
					b.WriteRune(r[i])
				}
				escape = false
			} else if i+1 != len(r) && (i == 0 || unicode.IsDigit(r[i+1])) {
				return "", ErrInvalidString
			}
		case "char":
			if escape {
				return "", ErrInvalidString
			}
			if i+1 != len(r) {
				checkAndWrite(&b, r[i], r[i+1])
			} else {
				b.WriteRune(r[i])
			}
		}
	}
	return b.String(), nil
}

func getType(r rune) string {
	switch {
	case string(r) == "\\":
		return "backslash"
	case unicode.IsDigit(r):
		return "digit"
	case unicode.IsLetter(r) || unicode.IsSpace(r):
		return "char"
	default:
		return ""
	}
}

func checkAndWrite(b *strings.Builder, currentRune rune, nextRune rune) {
	if digit, err := strconv.Atoi(string(nextRune)); err == nil {
		b.WriteString(strings.Repeat(string(currentRune), digit))
	} else {
		b.WriteRune(currentRune)
	}
}
