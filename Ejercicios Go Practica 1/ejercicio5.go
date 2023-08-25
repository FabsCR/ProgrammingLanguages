// Descripción: Este programa en Go implementa un evaluador de expresiones matemáticas expresadas en notación infija.
// Utiliza pilas basadas en slices para gestionar la evaluación de las expresiones. Se proporcionan ejemplos de expresiones
// matemáticas para demostrar el funcionamiento del evaluador. Las expresiones deben estar correctamente formateadas y
// separadas por espacios. El programa mostrará el resultado de cada expresión evaluada.
// Autor: Fabián Fernández

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) Push(val string) {
	*s = append(*s, val)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func evaluateExpression(expression string) float64 {
	operators := Stack{}
	values := Stack{}

	tokens := strings.Fields(expression)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			operators.Push(token)
		default:
			val, err := strconv.ParseFloat(token, 64)
			if err != nil {
				fmt.Println("Error de sintaxis en la expresión:", expression)
				return 0
			}

			values.Push(strconv.FormatFloat(val, 'f', -1, 64))
		}

		for len(operators) >= 2 && len(values) >= 2 {
			op := operators.Pop()
			if op == "" {
				break
			}

			b, _ := strconv.ParseFloat(values.Pop(), 64)
			a, _ := strconv.ParseFloat(values.Pop(), 64)

			switch op {
			case "+":
				values.Push(strconv.FormatFloat(a+b, 'f', -1, 64))
			case "-":
				values.Push(strconv.FormatFloat(a-b, 'f', -1, 64))
			case "*":
				values.Push(strconv.FormatFloat(a*b, 'f', -1, 64))
			case "/":
				if b != 0 {
					values.Push(strconv.FormatFloat(a/b, 'f', -1, 64))
				} else {
					fmt.Println("División por cero en la expresión:", expression)
					return 0
				}
			}
		}
	}

	if len(values) != 1 || len(operators) != 0 {
		fmt.Println("Error de sintaxis en la expresión:", expression)
		return 0
	}

	result, _ := strconv.ParseFloat(values.Pop(), 64)
	return result
}

func main() {
	expressions := []string{
		"3 + 4 * 2 / ( 1 - 5 ) ^ 2", // Ejemplo 1
		"2 * ( 3 + 5 ) - 8 / 4",     // Ejemplo 2
		"( 4 + 8 ) * 2 - 9 / 3",     // Ejemplo 3
	}

	for i, expression := range expressions {
		result := evaluateExpression(expression)
		fmt.Printf("Resultado de la expresión %d: %s = %.2f\n", i+1, expression, result)
	}
}
