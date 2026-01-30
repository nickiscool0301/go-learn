package main

import (
	"strconv"
	"unicode"
)

func validWordAbbreviation(word string, abbr string) bool {
	wLen, aLen := len(word), len(abbr)
	p1, p2 := 0, 0

	for p1 < wLen && p2 < aLen {
		if word[p1] == abbr[p2] {
			p1++
			p2++
		} else if abbr[p2] == '0' {
			return false
		} else if unicode.IsDigit(rune(abbr[p2])) {
			k := p2
			for k < aLen && unicode.IsDigit(rune(abbr[k])) {
				k++
			}
			numStr := abbr[p2:k]
			num, _ := strconv.Atoi(numStr)
			p1 += num
			p2 = k
		} else {
			return false
		}
	}

	return p1 == wLen && p2 == aLen
}
