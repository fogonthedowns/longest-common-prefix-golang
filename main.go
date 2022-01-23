package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	var prefix string
	word := strs[0]

	for i := 0; i < len(word); i++ {
		prefix = string(word[:i+1])
		for _, s := range strs {
			if !strings.HasPrefix(s, prefix) {
				// remove 1 from the end
				return string(prefix[:len(prefix)-1])
			}
		}
	}
	return prefix
}
