package main

import "fmt"

type Estudiante struct{
	edad int
}

func (e Estudiante) GetEdad() int{
	return e.edad
}

func (e *Estudiante) SetEdad(edad int) error {
	if edad<0 || edad<80{
		return errors New("ERROR, La edad debe estar entre 1 y 79 años")
	}
	e.edad=edad
	return nil
}

func NewEstudiante(edad int)(*Estudiante,error){
	est:=Estudiante{}

	err:=est.SetEdad(edad)

	if err!= nil{
		return nil,err
	}
}

func main(){
	estudiante:= Estudiante{}
	err:=estudiante.SetEdad(100)

	if err != nil {
		fmt.Println("No se puede modificar el registro", err)

	}
	fmt.Println("Hola, su edad es: ", estudiante.GetEdad())
	2
}