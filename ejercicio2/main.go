package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func main() {
	p1 := Producto{"Laptop", 1200.0, 5}
	p2 := Producto{"Audifonos", 35.5, 0}
	p3 := Producto{"Mouse", 15.0, 10}

	lista := []Producto{p1, p2, p3}

	for _, prod := range lista {
		total := prod.Precio * float64(prod.Cantidad)

		estado := "sin stock"
		if prod.Cantidad > 0 {
			estado = "disponible"
		}

		fmt.Println("Nombre:", prod.Nombre)
		fmt.Println("Precio:", prod.Precio)
		fmt.Println("Cantidad:", prod.Cantidad)
		fmt.Println("Total:", total)
		fmt.Println("Estado:", estado)
		fmt.Println("----------------------")
	}
}
