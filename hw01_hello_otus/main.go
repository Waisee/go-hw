package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func main() {
	// fmt.Println(reverse.String("Hello, OTUS!"))
	var b strings.Builder
	backslash := "\\"
	r := []rune(`qwe\4\5`)

	for i := 0; i < len(r); i++ {
		if string(r[i]) == backslash {
			if string(r[i-1]) == backslash {
				if i+1 != len(r) {
					if digit, err := strconv.Atoi(string(r[i+1])); err == nil {
						b.WriteString(strings.Repeat(string(r[i]), digit))
					} else {
						b.WriteRune(r[i])
					}
				} else {
					b.WriteRune(r[i])
				}
			}
		}
		if unicode.IsDigit(r[i]) {
			if string(r[i-1]) == backslash {
				if i+1 != len(r) {
					if digit, err := strconv.Atoi(string(r[i+1])); err == nil {
						b.WriteString(strings.Repeat(string(r[i]), digit))
					} else {
						b.WriteRune(r[i])
					}
				} else {
					b.WriteRune(r[i])
				}
			} else {
				if i == 0 || unicode.IsDigit(r[i+1]) {
					fmt.Println(ErrInvalidString)
					return
				}
			}
		}
		if unicode.IsLetter(r[i]) || unicode.IsSpace(r[i]) {
			if i+1 != len(r) {
				if digit, err := strconv.Atoi(string(r[i+1])); err == nil {
					b.WriteString(strings.Repeat(string(r[i]), digit))
				} else {
					b.WriteRune(r[i])
				}
			} else {
				b.WriteRune(r[i])
			}
		}
	}
	fmt.Println(b.String())
}
