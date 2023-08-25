// Descripción: Este programa en Go utiliza una estructura y un slice para administrar un inventario de una tienda de zapatos.
// Cada zapato en el inventario tiene información sobre su modelo (marca), precio y talla (entre 34 y 44). El programa permite
// realizar ventas de zapatos, actualizando el stock en cada venta. Además, se incluyen marcas de zapatos adicionales: Vans, Gucci,
// Givenchy y Converse. El programa muestra el inventario inicial y realiza ventas de ejemplo para demostrar su funcionalidad.
// Autor: Fabián Fernández

package main

import (
	"fmt"
)

type Calzado struct {
	Modelo string
	Precio float64
	Talla  int
	Stock  int
}

func venderZapato(inventario []*Calzado, indice int) {
	if indice < 0 || indice >= len(inventario) {
		fmt.Println("Índice de zapato inválido")
		return
	}

	zapato := inventario[indice]
	if zapato.Stock > 0 {
		zapato.Stock--
		fmt.Println("Venta exitosa. Stock restante:", zapato.Stock)
	} else {
		fmt.Println("No hay stock disponible para este zapato.")
	}
}

func main() {
	inventario := []*Calzado{
		{"Nike", 59000, 40, 5},
		{"Adidas", 67200, 40, 8},
		{"Puma", 29000, 36, 3},
		{"Vans", 35775, 41, 7},
		{"Gucci", 200000, 43, 2},
		{"Givenchy", 350000, 37, 4},
		{"Converse", 40000, 35, 6},
		{"Reebok", 48000, 39, 10},
		{"New Balance", 62000, 38, 12},
		{"Skechers", 39000, 37, 5},
	}

	fmt.Println("Inventario de Zapatos:")
	for i, zapato := range inventario {
		fmt.Printf("%d. Modelo: %s, Talla: %d, Precio: %.2f colones, Stock: %d\n", i+1, zapato.Modelo, zapato.Talla, zapato.Precio, zapato.Stock)
	}

	// Realizar algunas ventas de ejemplo
	fmt.Println("\nRealizando ventas:")
	venderZapato(inventario, 0) // Venta exitosa
	venderZapato(inventario, 0) // Venta exitosa
	venderZapato(inventario, 0) // No hay stock
	venderZapato(inventario, 3) // Venta exitosa
	venderZapato(inventario, 3) // Venta exitosa
	venderZapato(inventario, 3) // Venta exitosa
	venderZapato(inventario, 3) // Venta exitosa
	venderZapato(inventario, 3) // Venta exitosa
	venderZapato(inventario, 3) // Venta exitosa
	venderZapato(inventario, 3) // No hay stock
}
