package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var aCardOne, aCardTwo int
	var pos, res string
	str := "123456789"

	fmt.Scanln(&aCardOne, &aCardTwo)
	fmt.Scanln(&pos)

	switch pos {
	case "AABB":
		res = str[aCardTwo:9]
	case "ABAB":
		res = str[aCardOne:aCardTwo-1] + str[aCardTwo:9]
	case "ABBA":
		res = str[aCardOne : aCardTwo-1]
	case "BBAA":
		res = str[0 : aCardOne-1]
	case "BABA":
		res = str[0:aCardOne-1] + str[aCardOne:aCardTwo-1]
	case "BAAB":
		res = str[0:aCardOne-1] + str[aCardTwo:9]
	default:
		res = str
	}

	if len(res) == 2 {
		r := strings.Split(res, "")
		n1, _ := strconv.Atoi(r[0])
		n2, _ := strconv.Atoi(r[1])
		fmt.Println(n1, n2)
	} else {
		fmt.Println(-1)
	}

}
