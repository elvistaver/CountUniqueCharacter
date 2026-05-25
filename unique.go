package main

import (
	"fmt"
	"strings"
)

func CountUnique(text string) int {
	if len(text) == 0 {
		return 0
	}

	input := strings.ToLower(text)
	count := make(map[rune]int)
	final:=0

	for _, char := range input {
		count[char]++
		fmt.Println(count)
	}
	for _, value := range count {
		if value != 1 {
			continue
		}
		final++
	}
	return final
}
func main() {
	fmt.Println(CountUnique("gOoD"))
}