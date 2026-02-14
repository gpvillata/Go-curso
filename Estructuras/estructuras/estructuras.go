package estructuras

import "fmt"

func Arreglos() {
	fmt.Println("Arreglos")

	var arreglo = [...]int{11, 22, 32, 82727}

	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])
	}
}
