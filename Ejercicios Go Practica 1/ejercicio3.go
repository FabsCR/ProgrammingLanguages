// Descripción: Este programa en Go implementa una función que permite rotar una secuencia de elementos N posiciones
// a la izquierda o a la derecha, según se especifique. A partir de esta función, el programa realiza varias rotaciones
// en una secuencia previamente creada que es inmutable. Al final, se imprime el resultado de cada rotación junto con la
// secuencia original. Cada iteración de rotación mueve el elemento más a la izquierda al final de la secuencia (en caso
// de rotación a la izquierda) o el elemento más a la derecha al principio (en caso de rotación a la derecha).
// Autor: Fabián Fernández

package main

import (
	"fmt"
)

func rotateSlice(slice []interface{}, shift int, direction int) {
	length := len(slice)

	if shift < 0 {
		shift = length - ((-shift) % length)
	} else {
		shift = shift % length
	}

	if direction == 0 { // Rotación a la izquierda
		copy(slice, append(slice[shift:], slice[:shift]...))
	} else if direction == 1 { // Rotación a la derecha
		copy(slice, append(slice[length-shift:], slice[:length-shift]...))
	}
}

func main() {
	originalSequence := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h"}

	rotations := []struct {
		shift     int
		direction int
	}{
		{3, 0}, // Rotación a la izquierda
		{2, 1}, // Rotación a la derecha
		{5, 0}, // Rotación a la izquierda
	}

	fmt.Println("Secuencia Original:", originalSequence)

	for i, rotation := range rotations {
		// Clonamos la secuencia original para mantenerla inmutable
		sequenceCopy := make([]interface{}, len(originalSequence))
		copy(sequenceCopy, originalSequence)

		rotateSlice(sequenceCopy, rotation.shift, rotation.direction)

		direction := "izq"
		if rotation.direction == 1 {
			direction = "der"
		}

		fmt.Printf("Rotación %d: Cantidad = %d, Dirección = %s, Resultado = %v\n", i+1, rotation.shift, direction, sequenceCopy)
	}
}
