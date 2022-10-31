package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inp string

	fmt.Scan(&inp)

	a := strings.Split(inp, "")

	h1 := dToB(a[0])
	h1 = strings.Replace(h1, "0", ".", -1)
	h1 = strings.Replace(h1, "1", "*", -1)
	h2 := dToB(a[1])
	h2 = strings.Replace(h2, "0", ".", -1)
	h2 = strings.Replace(h2, "1", "*", -1)
	m1 := dToB(a[2])
	m1 = strings.Replace(m1, "0", ".", -1)
	m1 = strings.Replace(m1, "1", "*", -1)
	m2 := dToB(a[3])
	m2 = strings.Replace(m2, "0", ".", -1)
	m2 = strings.Replace(m2, "1", "*", -1)

	fmt.Println(h1[0:1], h2[0:1], " ", m1[0:1], m2[0:1])
	fmt.Println(h1[1:2], h2[1:2], " ", m1[1:2], m2[1:2])
	fmt.Println(h1[2:3], h2[2:3], " ", m1[2:3], m2[2:3])
	fmt.Println(h1[3:4], h2[3:4], " ", m1[3:4], m2[3:4])
}

func dToB(a string) string {
	b, _ := strconv.ParseInt(a, 10, 64)
	r := strconv.FormatInt(b, 2)

	for i := len(r); i < 4; i++ {
		r = "0" + r
	}
	return r
}
