package main

import (
	"fmt"
	"strings"
)

func esAnagrama(p1 string, p2 string) bool {
	p1 = strings.ToLower(p1)
	p2 = strings.ToLower(p2)

	runa1 := []rune(p1)
	runa2 := []rune(p2)

	if len(runa1) != len(runa2) {
		return false
	}

	contador := make(map[rune]int)

	for _, letra := range runa1 {
		contador[letra]++
	}

	for _, letra := range runa2 {
		contador[letra]--
		if contador[letra] < 0 {
			return false
		}
	}

	return true
}

func main() {
	var palabra1 string = "Sergio"
	var palabra2 string = "riesgo"

	fmt.Println("Palabras a Revisar")
	fmt.Println("Palabra1:", palabra1)
	fmt.Println("Palabra2:", palabra2)

	if esAnagrama(palabra1, palabra2) {
		fmt.Println("Es Anagrama: SI")
	} else {
		fmt.Println("Es Anagrama: NO")
	}

}
