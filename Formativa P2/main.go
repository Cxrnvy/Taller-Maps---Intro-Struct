package main

import (
	"fmt"
	"time"
)

func main() {
	fechaFab := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

	t1, _ := NewTablet("Paracetamol", "Genfar", fechaFab, 24, 500, false)
	t2, _ := NewTablet("Ibuprofeno", "Bayer", fechaFab, 36, 400, false)

	s1, _ := NewSyrup("Amoxicilina", "Mk", fechaFab, 12, 120, "Fresa")
	s2, _ := NewSyrup("Loratadina", "Genfar", fechaFab, 18, 100, "Cereza")

	tablets := []Tablet{t1, t2}
	syrups := []Syrup{s1, s2}

	fmt.Println("INVENTARIO DE MEDICAMENTOS")
	for _, t := range tablets {
		t.MostrarDetalles()
		fmt.Printf("Caducidad: %s\n\n", t.CalcularCaducidad().Format("2006-01-02"))
	}

	for _, s := range syrups {
		s.MostrarDetalles()
		fmt.Printf("Caducidad: %s\n\n", s.CalcularCaducidad().Format("2006-01-02"))
	}

	fmt.Println("ACTUALIZANDO MEDICAMENTO")
	fmt.Println("Nombre actual:", tablets[0].GetName())
	tablets[0].SetName("Paracetamol Forte")
	fmt.Println("Nuevo nombre con getter:", tablets[0].GetName())
}

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}
