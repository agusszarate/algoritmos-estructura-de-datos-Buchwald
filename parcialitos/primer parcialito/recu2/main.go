package recu2

import (
	. "tdas/cola"
	. "tdas/pila"
)

/*
================================================================================
1.er parcialito (2R) – 30/06/2025

Instrucciones:
Resolvé los siguientes problemas en forma clara y legible. Podés incluir tantas
funciones auxiliares como creas necesarias.
Los algoritmos a implementar deben ser de la mejor complejidad posible dadas
las características del problema.
================================================================================
*/

// -----------------------------------------------------------------------------
// EJERCICIO 1: Primitiva KUltimos para TDA Cola
// -----------------------------------------------------------------------------

// KUltimos es una primitiva para el TDA Cola.

// Consigna:
// Implementar una primitiva KUltimos[T any](k int) []T para el TDA Cola que dada una
// cola devuelva un arreglo con los últimos k elementos que saldrían de la cola
// (los del fondo), en el orden en que saldrían de esta.
//
// Condiciones:
//   - En el caso de tener menos de k elementos encolados, devolver los existentes en la cola.
//   - Al finalizar la ejecución de la función la cola debe quedar en el mismo estado que
//     antes de invocar a la primitiva.
//   - Indicar y justificar la complejidad del algoritmo.
//
// Ejemplo:
// La cola es [1, 2, 3, 4, 5] (primer elemento el 1) con k = 3,
// deberíamos obtener un arreglo [3, 4, 5].

func (c *Cola[T]) kUltimos(k int) []T {

	cant := c.cantidad

	colaAux := CrearColaEnlazada[T]()

	var tamaño int

	if cant < k {
		tamaño = cant
	} else {
		tamaño = k
	}

	ultimos := make([]T, tamaño)

	for i := 0; !c.EstaVacia(); i++ {
		elem := c.Desencolar()

		if i >= cant-tamaño {
			ultimos[i-(cant-tamaño)] = elem
		}

		colaAux.Encolar(elem)
	}

	for !colaAux.EstaVacia() {
		elem := colaAux.Desencolar()
		c.Encolar(elem)
	}

	return ultimos
}

//el costo del primer for es O(n) con N la cantidad de elementos de la cola
//el costo del segundo for es K, con K siendo constante
//el costo del tercer for es, en el peor caso, O(n)
//como K es constante y se puede despreciar junto con los 2 o(n)
//el algoritmo queda o(n)

// -----------------------------------------------------------------------------
// EJERCICIO 2: Verdadero o Falso (Teoría)
// -----------------------------------------------------------------------------

/*
2. Indicar Verdadero o Falso, justificando de forma concisa en cualquier caso.

   a. Se puede mejorar la complejidad de MergeSort si se cuenta con información extra.
      // Respuesta: ... Falso, mergesort siempre va a ser n log n

   b. Que un algoritmo de ordenamiento sea in-place implica que ordena respetando el
      orden original en el que aparecen los elementos de mismo valor.
      // Respuesta: ... Falso, que ordene respetando el orden original significa que es estable

   c. La complejidad de RadixSort depende únicamente del valor de d (cantidad de
      componentes) y de k (rango de cada componente).
      // Respuesta: ... falso d es la cantidad de digitos. la complejidad de radixsort es d * (k + n)
*/

// -----------------------------------------------------------------------------
// EJERCICIO 3: División y Conquista (Buscar duplicado)
// -----------------------------------------------------------------------------

// EncontrarDuplicado busca un número repetido en un arreglo ordenado.
//
// Consigna:
// Se cuenta con un arreglo de enteros ordenado de manera ascendente que contiene
// exactamente un número duplicado (es decir, todos los demás elementos son distintos).
// Implementar una función que encuentre dicho número utilizando División y Conquista.
// Indicar y justificar la complejidad del algoritmo, utilizando el Teorema Maestro.
func EncontrarDuplicado(arr []int) int {

	return -1
}

func EncontrarDuplicadoRecursivo(arr []int, min int, max int) int {
	if min == max {
		return arr[min]
	}

	mid := (min + max) / 2

	if arr[mid] == arr[0]+mid {
		return EncontrarDuplicadoRecursivo(arr, mid+1, max)
	} else {
		return EncontrarDuplicadoRecursivo(arr, min, mid)
	}

}
