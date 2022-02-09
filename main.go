package main

import "fmt"

func main() {
	valor1 := 2
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")

	} else {
		fmt.Println("No es 1")
	}

	//with and

	if valor1 == 1 && valor2 == 3 {

		fmt.Println("I'ts true")

	} else {
		fmt.Println("false")
	}
	//with or

	if valor1 == 1 || valor2 == 3 {

		fmt.Println("I'ts true")

	}

	if valor1%2 == 0 {
		fmt.Println("es par")
	} else {
		fmt.Println("no es par")
	}
}
