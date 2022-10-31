package main

import (
	"fmt"
	"math"
)

func main() {
	var i float64
	fmt.Scan(&i)
	fmt.Println(1 + math.Ceil(math.Log(i)/math.Log(2)))
}
