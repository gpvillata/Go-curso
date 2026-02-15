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

		fmt.Scanln(&opcion)
		if opcion < 0 || opcion > 4 {
			fmt.Println("Opción no válida")
			continue
		}

		if opcion == 0 {
			fmt.Println("Saliendo...")
			break
		}

		fmt.Printf("Elegiste la opción %d\n\n", opcion)

		switch opcion {
		case 1:
			fmt.Println("Agregando nuevo usuario")

		case 2:
			fmt.Println("Eliminando un  usuario")

		case 3:
			fmt.Println("Buscando un usuario")

		case 4:
			fmt.Println("Actualizando un usuario")

		default:
			fmt.Println("Opción no válida")

		}

	}
}
