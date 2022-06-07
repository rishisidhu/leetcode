package main

import "fmt"

func firstUniqChar(s string) int {
	char_map := make(map[string]int)
	for i := 0; i < len(s); i++ {
		char_map[string(s[i])] = char_map[string(s[i])] + 1
	}
	fmt.Print(char_map, "\n")
	for i := 0; i < len(s); i++ {
		val, _ := char_map[string(s[i])]
		fmt.Print("[", i, "]: ", val, string(s[i]), "\n")
		if val == 1 {
			return i
		}

	}
	return -1
}

func main() {
	str := "aabbb"
	fmt.Print(firstUniqChar(str))
}
