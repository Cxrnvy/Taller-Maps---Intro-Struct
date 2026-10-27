package main

import (
	"fmt"
)

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d | %s | Precio: $%.2f | Stock: %d\n", p.ID, p.Nombre, p.Precio, p.Cantidad)
}

type Inventario struct {
	Productos map[int]Producto
}

func (i *Inventario) AgregarProducto(p Producto) {
	if i.Productos == nil {
		i.Productos = make(map[int]Producto)
	}
	i.Productos[p.ID] = p
}

func (i *Inventario) EliminarProducto(id int) {
	delete(i.Productos, id)
}

func (i *Inventario) BuscarProducto(id int) (Producto, bool) {
	prod, existe := i.Productos[id]
	return prod, existe
}

func (i *Inventario) ListarProductos() {
	if len(i.Productos) == 0 {
		fmt.Println("No hay productos")
		return
	}
	for _, prod := range i.Productos {
		prod.MostrarInfo()
	}
}

func main() {
	inv := Inventario{}

	p1 := NuevoProducto(1, "Mouse Logitech G305", 45.00, 15)
	p2 := NuevoProducto(2, "Audífonos Soundcore", 85.50, 8)

	inv.AgregarProducto(p1)
	inv.AgregarProducto(p2)

	fmt.Println("--- Inventario Inicial ---")
	inv.ListarProductos()

	if prod, existe := inv.BuscarProducto(1); existe {
		prod.ActualizarPrecio(40.00)
		prod.ActualizarCantidad(12)
		inv.AgregarProducto(prod)
	}

	inv.EliminarProducto(2)

	fmt.Println("\n--- Inventario Actualizado ---")
	inv.ListarProductos()
}
