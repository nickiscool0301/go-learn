package main

import (
	"strings"
	"unicode"
)

func calculate(s string) int {
	n := len(s)
	stack := []int{}
	var num int
	op := '+'

	for i, ch := range s {
		if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
		}
		if strings.ContainsRune("+-*/", ch) || i == n-1 {
			if op == '+' {
				stack = append(stack, num)
			} else if op == '-' {
				stack = append(stack, -num)
			} else if op == '*' {
				head := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, head*num)
			} else if op == '/' {
				head := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, head/num)
			}
			op = ch
			num = 0
		}
	}
	var res int
	for _, num := range stack {
		res += num
	}
	return res
}
