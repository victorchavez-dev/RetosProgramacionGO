package main

import (
	"fmt"
)

func main() {
	var j, k int = 0, 1

	for i := range 50 {
		fmt.Printf("F(%d) = %d\n", i, j)
		j, k = k, j+k
	}
}
