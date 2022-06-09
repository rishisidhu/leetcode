package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	char_map := make(map[string]int)
	for i := 0; i < len(magazine); i++ {
		char_map[string(magazine[i])] = char_map[string(magazine[i])] + 1
	}
	fmt.Print(char_map, "\n")

	for i := 0; i < len(ransomNote); i++ {
		val, ok := char_map[string(ransomNote[i])]
		fmt.Print(i, ": ", string(ransomNote[i]), " | ", val, ok, "\n")
		if !ok {
			return false
		} else {
			if val == 0 {
				return false
			} else {
				char_map[string(ransomNote[i])] = val - 1
				fmt.Print(char_map)
			}
		}
	}

	return true
}

func main() {
	ransom := "aa"
	mag := "ab"
	fmt.Print(canConstruct(ransom, mag))
}
