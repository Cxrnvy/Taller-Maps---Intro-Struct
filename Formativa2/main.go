package main

import (
	"fmt"
	"sistema_farmacia/farmacia"
	"time"
)

func main() {
	fechaFab := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

	t1, _ := farmacia.NewTablet("Paracetamol", "Genfar", fechaFab, 24, 500, false)
	t2, _ := farmacia.NewTablet("Ibuprofeno", "Bayer", fechaFab, 36, 400, false)

	s1, _ := farmacia.NewSyrup("Amoxicilina", "Mk", fechaFab, 12, 120, "Fresa")
	s2, _ := farmacia.NewSyrup("Loratadina", "Genfar", fechaFab, 18, 100, "Cereza")

	tablets := []farmacia.Tablet{t1, t2}
	syrups := []farmacia.Syrup{s1, s2}

	fmt.Println("--- INVENTARIO DE MEDICAMENTOS ---")
	for _, t := range tablets {
		t.MostrarDetalles()
		fmt.Printf("Caducidad: %s\n\n", t.CalcularCaducidad().Format("2006-01-02"))
	}

	for _, s := range syrups {
		s.MostrarDetalles()
		fmt.Printf("Caducidad: %s\n\n", s.CalcularCaducidad().Format("2006-01-02"))
	}

	fmt.Println("--- ACTUALIZANDO MEDICAMENTO ---")
	fmt.Println("Nombre actual:", tablets[0].GetName())
	tablets[0].SetName("Paracetamol Forte")
	fmt.Println("Nuevo nombre con getter:", tablets[0].GetName())
}
