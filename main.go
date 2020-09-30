package main

import "fmt"

func main() {

	var e float64 = 0.0
	var ciclo int

	fmt.Scanln(&ciclo)

	for i := 0; i < ciclo; i++ {

		e += float64(1) / float64(factorial(i))

	}

	fmt.Println(e)
}

func factorial(numero int) int {

	factorial := 1

	for i := 1; i <= numero; i++ {
		factorial = factorial * i
	}

	return factorial
}
