package main

import (
	. "tdas/pila"
)

// =============================================================================
// EJERCICIO 1
// =============================================================================
// Se tiene un arreglo de n cadenas y una buena función de hashing. Se quiere
// ordenar dichas cadenas por su resultado en la función de hashing, habiéndole
// hecho previamente % K a cada uno de los resultados (donde K es un valor enorme,
// muchísimo más grande que n).
//
// Implementar un algoritmo que ordene las cadenas por dicho criterio en O(n).
// Justificar brevemente por qué es correcta la aplicación del algoritmo que
// implementaste. Justificar la complejidad del algoritmo implementado.
func Ordenar(cadenas []string, valoresHash []int64, K int64) []string {
	// TODO: Implementar aquí
	return nil
}

// =============================================================================
// EJERCICIO 2
// =============================================================================
// Implementar un algoritmo que reciba un arreglo de enteros desordenado y un
// número elem que, por división y conquista determine si elem se encuentra
// en el arreglo.
//
// Indicar y justificar adecuadamente la complejidad del algoritmo implementado.
func existe(arr []int, elem int) bool {
	if len(arr) == 0 {
		return false
	}

	mid := len(arr) / 2

	if arr[mid] == elem {
		return true
	}

	der := existe(arr[mid:], elem)

	if der {
		return true
	} else {
		return existe(arr[:mid], elem)
	}
}

// T(o) = 2 * T(n/2) + O(n^0)
// log2 2 = 1 > 0
// complejidad: O(n)

// =============================================================================
// EJERCICIO 3
// =============================================================================
// Implementar una función que reciba una pila de enteros y devuelva la suma de
// todos los elementos. Al finalizar la ejecución de la función, la pila debe
// quedar en el mismo estado que tenía antes de ejecutar la misma.
//
// La función no puede utilizar estructuras auxiliares (incluyendo arreglos).
// Indicar y justificar la complejidad de la función implementada.

func pilaSumar(pila Pila[int]) int {

	if pila.EstaVacia() {
		return 0
	}

	elem := pila.Desapilar()

	total := pilaSumar(pila)

	pila.Apilar(elem)

	return total + elem
}
