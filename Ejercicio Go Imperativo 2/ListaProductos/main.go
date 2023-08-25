package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	// Modificar el código para que cuando se agregue un producto y este ya existe, incremente la cantidad
	// y, si es necesario, actualice el precio.
	for i := range *l {
		if (*l)[i].nombre == nombre {
			(*l)[i].cantidad += cantidad
			if precio != (*l)[i].precio {
				(*l)[i].precio = precio
			}
			return
		}
	}
	*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
}

// Función para agregar una cantidad potencialmente infinita de productos previamente creados
func (l *listaProductos) agregarProductos(productos ...producto) {
	for _, p := range productos {
		l.agregarProducto(p.nombre, p.cantidad, p.precio)
	}
}

func (l *listaProductos) buscarProducto(nombre string) (producto, int) {
	for i, p := range *l {
		if p.nombre == nombre {
			return p, i
		}
	}
	return producto{}, -1
}

func (l *listaProductos) venderProducto(nombre string, cantidad int) error {
	p, idx := l.buscarProducto(nombre)
	if idx == -1 {
		return fmt.Errorf("producto no encontrado")
	}

	if p.cantidad < cantidad {
		return fmt.Errorf("no hay suficiente existencia")
	}

	p.cantidad -= cantidad

	if p.cantidad == 0 {
		// Eliminar el producto si la cantidad llega a cero
		*l = append((*l)[:idx], (*l)[idx+1:]...)
	}

	return nil
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	var productosMinimos listaProductos
	for _, p := range *l {
		if p.cantidad < existenciaMinima {
			productosMinimos = append(productosMinimos, p)
		}
	}
	return productosMinimos
}

func (l *listaProductos) aumentarInventarioDeMinimos() {
	productosMinimos := l.listarProductosMínimos()
	for _, p := range productosMinimos {
		cantidadFaltante := existenciaMinima - p.cantidad
		if cantidadFaltante > 0 {
			l.agregarProducto(p.nombre, cantidadFaltante, p.precio)
			fmt.Printf("Se agregaron %d unidades de %s al inventario.\n", cantidadFaltante, p.nombre)
		}
	}
}

func (l *listaProductos) ordenarProductosPorPrecio() {
	sort.Slice(*l, func(i, j int) bool {
		return (*l)[i].precio < (*l)[j].precio
	})
}

func main() {
	llenarDatos()
	fmt.Println("Lista de productos original:")
	fmt.Println(lProductos)

	err := lProductos.venderProducto("arroz", 5)
	if err != nil {
		fmt.Println("Error al vender arroz:", err)
	}

	fmt.Println("\nProductos con existencia mínima:")
	productosMinimos := lProductos.listarProductosMínimos()
	fmt.Println(productosMinimos)

	lProductos.aumentarInventarioDeMinimos()
	fmt.Println("\nLista de productos después de aumentar inventario:")
	fmt.Println(lProductos)

	lProductos.ordenarProductosPorPrecio()
	fmt.Println("\nLista de productos ordenada por precio:")
	fmt.Println(lProductos)
}

func llenarDatos() {
	lProductos.agregarProductos(
		producto{nombre: "arroz", cantidad: 15, precio: 2500},
		producto{nombre: "frijoles", cantidad: 4, precio: 2000},
		producto{nombre: "leche", cantidad: 8, precio: 1200},
		producto{nombre: "café", cantidad: 12, precio: 4500},
		producto{nombre: "pan", cantidad: 20, precio: 1500},
		producto{nombre: "azúcar", cantidad: 10, precio: 1800},
		producto{nombre: "aceite", cantidad: 5, precio: 3000},
		producto{nombre: "sal", cantidad: 15, precio: 900},
	)
}
