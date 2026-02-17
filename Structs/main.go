package structs

import "fmt"

type Persona struct {
	ID              int
	Nombre          string
	Edad            int
	FechaNacimiento string
	Genero          string
	Altura          float64
	Peso            float64
	DNI             string
}

func (p Persona) Imprimir() {
	fmt.Println("ID: ", p.ID)
	fmt.Println("Nombre: ", p.Nombre)
	fmt.Println("Edad: ", p.Edad)
	fmt.Println("Fecha de nacimiento: ", p.FechaNacimiento)
	fmt.Println("Genero: ", p.Genero)
	fmt.Println("Altura: ", p.Altura)
	fmt.Println("Peso: ", p.Peso)
	fmt.Println("DNI: ", p.DNI)
}

func main() {
	fmt.Println("Nueva persona")

	var persona Persona = Persona{
		ID:              1,
		Nombre:          "Piero",
		Edad:            22,
		FechaNacimiento: "17/02/2004",
		Genero:          "Masculino",
		Altura:          1.75,
		Peso:            70,
		DNI:             "12345678",
	}
	persona.Imprimir()
}
