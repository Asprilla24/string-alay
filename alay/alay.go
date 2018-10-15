package alay

import (
	"math/rand"
	"strings"
	"unicode"
)

var m = map[rune]rune{
	'a': '4',
	'i': '1',
	'e': '3',
	'o': '0',
}

func ToAlayMapper(r rune) rune {
	if i := rand.Intn(10); i == 0 {
		if v, ok := m[r]; ok {
			return v
		}
		return unicode.ToUpper(r)
	}

	return r
}

//ToAlay is a function to translate string become alay language
func ToAlay(s string) string {
	return strings.Map(ToAlayMapper, s)
}
