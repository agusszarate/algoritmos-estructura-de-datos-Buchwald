package main

// 1)Implementar una primitiva del TDA cola que devuelva dos colas:
// una con los elementos de las posiciones pares y otra con los elementos de las posiciones impares
// (el primero de la cola puede considerarse como elemento en la posicion 0).
// La cola original debe quedar en el mismo estado inicial.
// Indicar y justificar la complejidad de la primitiva.

func (c *Cola[T]) Dividir() (Cola[T], Cola[T]) {
	cPar := NuevaCola[T]()
	cImpar := NuevaCola[T]()

	for i := 0; i < c.cantidad; i++ {
		elem := c.Desacolar()
		if i%2 == 0 {
			cPar.Encolar(elem)
		} else {
			cImpar.Encolar(elem)
		}
		c.Encolar(elem)
	}
	return cPar, cImpar
}

// 2)
// Implementar una funcion que, dado un arreglo ordenado y sin elementos repetidos de valores enteros no negativos
// obtenga el minimo valor que no se encuentre en el arreglo.
// Indicar y justificar adecuadamente la complejidad del algoritmo

func MinimoExcluido(arr []int) int {
	return minimoExcluidoRecursivo(arr, 0, len(arr)-1)
}

func minimoExcluidoRecursivo(arr []int, inicio int, fin int) int {

	if inicio > fin {
		return inicio
	}

	mid := (inicio + fin) / 2

	if arr[mid] > mid {
		return minimoExcluidoRecursivo(arr, inicio, mid-1)
	} else {
		return minimoExcluidoRecursivo(arr, mid+1, fin)
	}
}

// 3)
//Implementar una funciona masGrandePOsible/diguitos []int) int que, dado un arreglo de digitos (0 a 9) devuelva el numero mas grande posible
// que se puede formar con esos digitos. Por ejemplo, si recibe [1, 3, 1, 5, 9] debe devolver 95311.
// Indicar y justificar adecuadamente la complejidad del algoritmo

func masGrandePosible(digitos []int) int {
	// Ordenamiento manual usando bubble sort (descendente)
	for i := 0; i < len(digitos); i++ {
		for j := 0; j < len(digitos)-1-i; j++ {
			if digitos[j] < digitos[j+1] {
				digitos[j], digitos[j+1] = digitos[j+1], digitos[j]
			}
		}
	}

	numero := 0
	for i := 0; i < len(digitos); i++ {
		numero = numero*10 + digitos[i]
	}
	return numero
}

// Complejidad: O(n log n) debido al ordenamiento, donde n es la cantidad de digitos en el arreglo.
