package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	lenS := len(s)
	freq := make([]int, 26)
	for i := 0; i < lenS; i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}
	for _, count := range freq {
		if count != 0 {
			return false
		}
	}
	return true
}
