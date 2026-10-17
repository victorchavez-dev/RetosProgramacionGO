package main

import "fmt"

func main() {
	contador := make(map[int]int)

	for i := 1; i <= 100; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				contador[i]++
			}
		}
		if contador[i] == 2 { //Divisores de 1 y mismo numero
			fmt.Println("Es Primo:", i)
		}
	}
	//Muestra Clave(Numero)|Valor(Contador)
	fmt.Println(contador)
}
