package main

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isValid(s, l+1, r) || isValid(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isValid(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
