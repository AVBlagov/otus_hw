package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ErrInvalidString - error.
var ErrInvalidString = errors.New("invalid string")

// ErrZeroRune - zero rune with int.
var ErrZeroRune = errors.New("zero rune with int")

// Unpack - func.
func Unpack(str string) (string, error) {
	var ustring strings.Builder
	var prev rune
	var temp = []rune(str)

	for i, cur := range temp {
		switch {
		case i == 0 && unicode.IsDigit(cur):
			return "", ErrInvalidString
		case i == 0 && !unicode.IsDigit(cur):
			prev = cur
			ustring.WriteString(string(prev))
		case i > 0 && unicode.IsDigit(cur) && !unicode.IsDigit(prev):
			num, _ := strconv.Atoi(string(cur))
			if num == 0 {
				runes := []rune(ustring.String())
				copy(runes[i-1:], runes[i:])
				runes[len(runes)-1] = ' '
				runes = runes[:len(runes)-1]
				ustring.Reset()
				ustring.WriteString(string(runes))
			} else {
				ustring.WriteString(strings.Repeat(string(prev), num-1))
				prev = cur
			}
		case i > 0 && unicode.IsDigit(cur) && unicode.IsDigit(prev):
			return "", ErrInvalidString
		case i > 0 && !unicode.IsDigit(cur):
			prev = cur
			ustring.WriteString(string(cur))
		}
	}

	return ustring.String(), nil
}
