// Descripción: Este programa en Go imprime en pantalla una figura compuesta de asteriscos (*) con una cantidad
// de elementos en la línea del centro proporcionada por el usuario. La cantidad de elementos debe ser un número impar positivo.
// Autor: Fabián Fernández

package main

import "fmt"

func main() {
	fmt.Println("Ingrese la cantidad de elementos en la línea del centro (número impar):")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n%2 == 0 || n <= 0 {
		fmt.Println("Por favor, ingrese un número impar positivo.")
		return
	}

	printFigure(n)
}

func printFigure(n int) {
	spaces := n / 2

	for i := 0; i < n; i += 2 {
		for j := 0; j < spaces; j++ {
			fmt.Print(" ") // Imprimir espacios
		}

		for j := 0; j <= i; j++ {
			fmt.Print("* ") // Imprimir asteriscos y espacios
		}

		fmt.Println()
		spaces--
	}

	spaces = 1

	for i := n - 2; i > 0; i -= 2 {
		for j := 0; j < spaces; j++ {
			fmt.Print(" ") // Imprimir espacios
		}

		for j := 0; j < i; j++ {
			fmt.Print("* ") // Imprimir asteriscos y espacios
		}

		fmt.Println()
		spaces++
	}
}
