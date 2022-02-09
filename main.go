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

	// recognize if a number is even

	if valor1%2 == 0 {
		fmt.Println("es par")
	} else {
		fmt.Println("no es par")
	}

	//types module declaration

	modulo := 3 % 2
	switch modulo {
	case 0:
		fmt.Println("es par")

	default:
		fmt.Println("no es par")
	}

	switch module := 4 % 2; module {
	case 0:
		fmt.Println("es par")

	default:
		fmt.Println("no es par")
	}

	//module without condition
	value := 200
	switch {
	case value > 100:
		fmt.Println("is greater than 100")
	case value < 0:
		fmt.Println(" is lesser than  0")
	default:
		fmt.Println(" no conditon")
	}

	//Keywords

	//Defer
	// defer fmt.Println("Hola")
	// fmt.Println("Mundo")

	//Continue & Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice, similar declaration to array.
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// In slice method
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
	// Append
	slice = append(slice, 7)
	fmt.Println(slice[7])

	// Append new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
