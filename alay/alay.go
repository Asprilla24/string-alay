package alay

import (
	"strings"
)

var m = map[rune]rune{
	'a': '4',
	'i': '1',
	'e': '3',
	'o': '0',
}

func ToAlayMapper(r rune) rune {
	if v, ok := m[r]; ok {
		return v
	}
	return r
}

func toAlay(s string) string {
	return strings.Map(ToAlayMapper, s)
}
