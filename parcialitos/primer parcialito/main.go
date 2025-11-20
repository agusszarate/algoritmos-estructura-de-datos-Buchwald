package main

import (
	. "tdas/cola"
)

// 1)Implementar una primitiva del TDA cola que devuelva dos colas:
// una con los elementos de las posiciones pares y otra con los elementos de las posiciones impares
// (el primero de la cola puede considerarse como elemento en la posicion 0).
// La cola original debe quedar en el mismo estado inicial.
// Indicar y justificar la complejidad de la primitiva.

func (c *Cola[T]) Dividir() (Cola[T], Cola[T]) {
	cPar := CrearColaEnlazada[T]()
	cImpar := CrearColaEnlazada[T]()

	for i := 0; i < c.cantidad-1; i++ {

		elem := c.Desencolar()

		if i%2 == 0 {
			cPar.Encolar(elem)
		} else {
			cImpar.Encolar(elem)
		}

		c.Encolar(elem)
	}

	return cPar, cImpar
}

func (c *Cola[T]) Dividir() (Cola[T], Cola[T]) {
	cPar := CrearColaEnlazada[T]()
	cImpar := CrearColaEnlazada[T]()

	for i := range c.cantidad {
		elem := c.Desencolar()

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
// Implementar una funcion que, dado un arreglo ordenado
// y sin elementos repetidos de valores enteros no negativos
// obtenga el minimo valor que no se encuentre en el arreglo.
// Indicar y justificar adecuadamente la complejidad del algoritmo

func MinimoExcluidoAsi(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	mid := len(arr) / 2

	if arr[mid] == arr[0]+mid {
		return MinimoExcluidoAsi(arr[mid:])
	} else {
		return MinimoExcluidoAsi(arr[:mid+1])
	}
}

func minimoExcluidoRecursivo(arr []int, inicio int, fin int) int {
	if inicio > fin {
		return arr[0] + inicio
	}

	mid := (inicio + fin) / 2

	if arr[mid] == arr[0]+mid {
		return minimoExcluidoRecursivo(arr, mid+1, fin)
	} else {
		return minimoExcluidoRecursivo(arr, inicio, mid-1)
	}

}

func MinimoExcluido(arr []int) int {
	return minimoExcluidoRecursivo(arr, 0, len(arr)-1)
}

// 3)
// Implementar una funcion masGrandePosible(digitos []int) int que,
// dado un arreglo de digitos (0 a 9) devuelva el numero mas grande posible
// que se puede formar con esos digitos. Por ejemplo, si recibe [1, 3, 1, 5, 9] debe devolver 95311.
// Indicar y justificar adecuadamente la complejidad del algoritmo

func masGrandePosiblee(digitos []int) int {
	lista := make([]int, 10)

	for _, i := range digitos {
		lista[i]++
	}

	var numero int

	for i := len(lista) - 1; i >= 0; i-- {

		for cant := 0; cant < lista[i]; cant++ {
			numero = numero*10 + i
		}

	}

	return numero
}

func masGrandePosible(digitos []int) int { //el algoritmo es O(N) con N = len de digitos
	countingSort := make([]int, 10)

	for _, digito := range digitos { //esto se ejecuta exactamente len de digitos
		countingSort[digito]++
	}

	numero := 0

	for digito := len(countingSort) - 1; digito >= 0; digito-- { //esto se ejecuta exactamente 10 veces
		for a := 0; a < countingSort[digito]; a++ { //esto se ejecuta exactamente len de digitos
			numero = (numero * 10) + digito
		}
	}

	return numero
}

// complejidad O(N)
