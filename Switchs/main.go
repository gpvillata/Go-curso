package switchs

import "fmt"

func main() {

	fmt.Println("Ingrese un numero entero")
	var num uint
	fmt.Scan(&num)
	switch {
	case num > 0:
		fmt.Println("El numero es positivo")
	case num == 0:
		fmt.Println("El numero es cero")
	default:
		fmt.Println("El numero es negativo")
	}
}
