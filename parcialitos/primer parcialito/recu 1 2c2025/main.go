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
// La firma de la función debe ser:
// func Ordenar(cadenas []string, valoresHash []int64, K int64) []string
//
// valoresHash ya tiene cada valor posterior a haber hecho % K.
// Recomendamos recordar las propiedades de las funciones de hashing.
// Si necesitás un algoritmo de ordenamiento auxiliar al que estés implementando,
// podés considerarlo ya implementado.
// Justificar brevemente por qué es correcta la aplicación del algoritmo que implementaste.
// Justificar la complejidad del algoritmo implementado.
//
// Ejemplo: si queremos ordenar las cadenas: gato, perro, elefante, comadreja
// y los resultados de aplicarles la función de hashing (y % K) son 19, 703, 9872, 37,
// respectivamente, las cadenas deben de quedar: gato, comadreja, perro, elefante.

func Ordenar(cadenas []string, valoresHash []int64, K int64) []string {
	// TODO: Implementar solución
	return nil
}

// =============================================================================
// EJERCICIO 2
// =============================================================================
// Implementar un algoritmo que reciba un arreglo de enteros desordenado y un
// número elem que, por división y conquista determine si elem se encuentra en
// el arreglo.
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

// =============================================================================
// EJERCICIO 3
// =============================================================================
// Implementar una función que reciba una pila de enteros y devuelva la suma de
// todos los elementos. Al finalizar la ejecución de la función, la pila debe
// quedar en el mismo estado que tenía antes de ejecutar la misma.
// La función no puede utilizar estructuras auxiliares (incluyendo arreglos).
// La firma de la función debe ser pilaSumar(pila Pila[int]) int.
// Indicar y justificar la complejidad de la función implementada.

// Definición de la interfaz necesaria para la firma de la función

func pilaSumar(pila Pila[int]) int {

	if pila.EstaVacia() {
		return 0
	}

	elem := pila.Desapilar()

	total := pilaSumar(pila)

	pila.Apilar(elem)

	return total + elem
}
