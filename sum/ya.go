package main

import (
	"errors"
	"unicode/utf8"
)

var ErrInvalidUT"invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}
