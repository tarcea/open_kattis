package main

import (
	"fmt"
	"strings"
)

func main() {
	var adventures int
	var results []string
	fmt.Scanln(&adventures)

	var c string

	for i := 0; i < adventures; i++ {
		fmt.Scanln(&c)
		result := readAdventure(c)
		results = append(results, result)
	}

	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
}

func readAdventure(str string) string {
	money := strings.Count(str, "$")
	inceses := strings.Count(str, "|")
	gems := strings.Count(str, "*")
	traders := strings.Count(str, "t")
	jewelers := strings.Count(str, "j")
	bankers := strings.Count(str, "b")

	if money-bankers != 0 || inceses-traders != 0 || gems-jewelers != 0 {
		return "NO"
	} else {

		return "YES"
	}
}
