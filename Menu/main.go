package main

import (
	"fmt"
)

func main() {
	opcion := -1
	for {
		fmt.Println("--------------------------------- MENU PRINCIPAL -------------------------------------")
		fmt.Println("----------------------USUARIOS ---------------------------")
		fmt.Println("0 - Salir")
		fmt.Println("1 - Nuevo  2 - Eliminar  3 - Buscar  4 - Actualizar")
		fmt.Print("Elige una opción: ")

		_, err := fmt.Scanln(&opcion)

		if err == nil {
			fmt.Println("Ha ocurrido un error", err)
		}
		if opcion == 0 {
			fmt.Println("Saliendo...")
			break
		}

		fmt.Printf("Elegiste la opción %d\n\n", opcion)

		switch opcion {
		case 1:
			fmt.Println("Agregando nuevo usuario")
			break

		case 2:
			fmt.Println("Eliminando un  usuario")
			break

		}

	}
}
