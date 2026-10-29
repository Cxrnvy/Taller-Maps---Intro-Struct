package main

import "fmt"

func main(){
	nombre string
}

func (P Persona) GetNombre() string{
	return p.P.nombre

}

func (p *Persona) SetNombre(nombre string){
	p.nombre=nombre
}

func main(){
	persona := Persona()

	persona.SetNombre("Cesar")
	fmt.Println("Su nombre es: ", persona.GetNombre())

	persona.SetNombre("Ariel")

	fmt.Println("Ahora su nombre es: ", persona.GetNombre())
	
}

