package main

import "fmt"

func main() {
	var nombre = "Juan"
	var edad = 32

	if nombre == "Juan" {
		fmt.Println(nombre)
	}
	if edad > 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

}
