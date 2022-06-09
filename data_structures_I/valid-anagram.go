package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	char_map := make(map[string]int)
	for i := 0; i < len(s); i++ {
		char_map[string(s[i])] = char_map[string(s[i])] + 1
	}
	fmt.Print(char_map, "\n")

	for i := 0; i < len(t); i++ {
		val, ok := char_map[string(t[i])]
		//fmt.Print(i, ": ", string(ransomNote[i]), " | ", val, ok, "\n")
		if !ok {
			return false
		} else {
			if val == 0 {
				return false
			} else {
				char_map[string(t[i])] = val - 1
				//fmt.Print(char_map)
			}
		}
	}

	remains := 0
	for i := 0; i < len(s); i++ {
		remains = remains + char_map[string(s[i])]
	}
	if remains != 0 {
		return false
	}
	return true
}

func main() {
	s := "anatoly"
	t := "lonyatab"
	fmt.Print(isAnagram(s, t))
}
