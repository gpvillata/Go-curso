package control

import (
	"fmt"
)

func ImprimirFor() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

}
func ImprimirDo(n int) {
	var i int
	i = 1
	for i < n {
		i = i + 1
		println("Hola mundo")
	}
}
func Imprimir() {

	for i := range 10 {
		fmt.Println(i)
	}
}
func ImprimirRange() {
	for range 10 {
		var i int = 0
		fmt.Println("Bienvenidos", i)
		i++
	}
}
