package main

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	for _, str := range strs {
		freq := make([]int, 26)
		for _, c := range str {
			freq[c-'a']++
		}
		key := ""
		for _, count := range freq {
			key += "#" + string(count)
		}
		if _, exist := groupMap[key]; !exist {
			groupMap[key] = []string{str}
		} else {
			groupMap[key] = append(groupMap[key], str)
		}
	}
	result := make([][]string, 0, len(groupMap))
	for _, group := range groupMap {
		result = append(result, group)
	}
	return result
}
