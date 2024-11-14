package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array
	fmt.Println("Array")

	var array1 [5]string
	array1[0] = "possicao 1"
	fmt.Println(array1)

	array2 := [5]string{"possicao 1", "possicao 2", "possicao 3", "possicao 4", "possicao 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 34, 5, 6, 78}
	fmt.Println(array3)

	// slice
	fmt.Println("Slices")

	slice := []int{12, 1, 24, 5, 737, 53, 2, 123, 74, 537}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 76)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "possicao alterada"
	fmt.Println(slice2)

	// arrays internos
	fmt.Println("----------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	// quando bater o maximo ele dobra o array interno
	slice3 = append(slice3, 7)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 7)
	fmt.Println(len(slice4)) // tamanho
	fmt.Println(cap(slice4)) // capacidade

}
