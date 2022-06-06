package main

import "fmt"

func sourceOfTruth(s1 string) map[string]int {
	rune_map := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		val, _ := rune_map[string(s1[i])]
		rune_map[string(s1[i])] = val + 1
	}
	fmt.Print("SOT: ", rune_map, "\n****************\n")
	return rune_map
}

func makeZero(rune_map map[string]int) map[string]int {
	for key, _ := range rune_map {
		rune_map[key] = 0
	}
	return rune_map
}

func checkInclusion(s1 string, s2 string) bool {
	rune_map := make(map[string]int)
	source := sourceOfTruth(s2)
	for i := 0; i < len(s2); i++ {
		rune_map[string(s2[i])] = 0
	}

	counter := 0
	deletion_index := -1
	for i := 0; i < len(s1); i++ {
		val, ok := rune_map[string(s1[i])]
		fmt.Print("Pre: ", i, val, ok, rune_map, string(s1[i]), " | ", counter, "\n")
		if ok {
			if val < source[string(s1[i])] {
				if deletion_index == -1 {
					deletion_index = i
				}
				rune_map[string(s1[i])] = val + 1
				counter++
				if counter == len(s2) {
					return true
				}
			} else {
				fmt.Print("mistake section: ", deletion_index, counter, "\n")
				//make the mistake
				rune_map[string(s1[i])] = val + 1
				counter++
				//correct the mistake
				for {
					fmt.Print("mistake section pre: ", deletion_index, counter, "\n")
					rune_map[string(s1[deletion_index])] = rune_map[string(s1[deletion_index])] - 1
					deletion_index++
					counter--
					fmt.Print("mistake section post: ", deletion_index, counter, "\n")
					if rune_map[string(s1[i])] == val {
						break
					}
				}
			}
		} else {
			if counter == 0 {
				continue
			} else if counter != 0 {
				counter = 0
				rune_map = makeZero(rune_map)
				deletion_index = -1
				fmt.Print("counter is 0 now\n", deletion_index)
			}
		}
		fmt.Print("Post: ", i, val, ok, rune_map, string(s1[i]), " | ", counter, "\n\n")
	}
	fmt.Print(rune_map, "\n")
	return false
}

func main() {

	str1 := "indianloggersaresitajs"
	str2 := "sitares"
	fmt.Print(checkInclusion(str1, str2))
}
