package main

import "fmt"

func calcularValorInventario(p Producto) float64{
//multiplicar la cantidad por el precio
	return p.precio * float64(p.stock)
}

func (p Producto) NuevoStock(cantidad int){
	if cantidad > 0{
		p.stock:=cantidad
	}
}

func (p *Producto) Vender(cantidad int) bool{
	if cantidad < 0 || cantidad >= p.stock {
		return false
	}
}

func NewProduct(nombre string, precio float64, stock int) Producto{
	return Producto{
		Nombre : nombre,
		precio : precio,
		stock : 0
	}
}

func main() {
	type Producto struct {
		Nombre string
		precio float64
		stock  int
	}

	product := Producto{}

	product.Nombre = "Intel Core Ultra 7"
	product.precio = 850.90
	product.stock = 200

	product2 := Producto{
		Nombre: "Teclado",
		precio: 30.5,
		stock:  15,
	}
	fmt.Println(product)
	fmt.Println(product2)

	fmt.Println("El Producto: ", product2.Nombre, " tiene un valor del inventario de:", calcularValorInventario(product2))

if product.Vender(200) == true{
	fmt.Println("El producto", product.Nombre, "aun cuenta con stock disponible, el stock actual es de: ", product.stock)
} else {
	fmt.Println("El producto", product.Nombre, " ya no cuenta con stock disponible")
}

product3:=NewProduct("mouse", 12.81,50)
fmt.Println(product3)

}
