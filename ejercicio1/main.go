package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese una frase: ")
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')

	texto = strings.ToLower(texto)
	palabras := strings.Fields(texto)

	conteo := make(map[string]int)
	masRepetida := ""
	maximo := 0

	for _, p := range palabras {
		conteo[p]++

		if conteo[p] > maximo {
			maximo = conteo[p]
			masRepetida = p
		}
	}

	for p, c := range conteo {
		fmt.Println(p, "-", c)
	}

	fmt.Println("La palabra mas repetida es:", masRepetida, "y aparece", maximo, "veces")
}
