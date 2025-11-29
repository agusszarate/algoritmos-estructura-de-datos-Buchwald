package main

import (
	. "tdas/diccionario"
)

type ColaPrioridad[T any] interface {

	// EstaVacia devuelve true si la la cola se encuentra vacía, false en caso contrario.
	EstaVacia() bool

	// Encolar Agrega un elemento al heap.
	Encolar(T)

	// VerMax devuelve el elemento con máxima prioridad. Si está vacía, entra en pánico con un mensaje
	// "La cola esta vacia".
	VerMax() T

	// Desencolar elimina el elemento con máxima prioridad, y lo devuelve. Si está vacía, entra en pánico con un
	// mensaje "La cola esta vacia"
	Desencolar() T

	// Cantidad devuelve la cantidad de elementos que hay en la cola de prioridad.
	Cantidad() int
}

const (
	_CAPACIDAD_INICIAL  = 10
	_FACTOR_REDIMENSION = 2
	_FACTOR_REDUCCION   = 4
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func calcularPosicionPadre(posHijo int) int {
	return (posHijo - 1) / 2
}

func calcularPosicionHijoIzq(posPadre int) int {
	return 2*posPadre + 1
}

func calcularPosicionHijoDer(posPadre int) int {
	return 2*posPadre + 2
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (heap *heap[T]) upheap(pos int) {
	if pos == 0 {
		return
	}
	posPadre := calcularPosicionPadre(pos)
	if heap.cmp(heap.datos[pos], heap.datos[posPadre]) > 0 {
		swap(heap.datos, pos, posPadre)
		heap.upheap(posPadre)
	}
}

func downheapRecursivo[T any](arr []T, pos, tam int, cmp func(T, T) int) {
	posHijoIzq := calcularPosicionHijoIzq(pos)
	posHijoDer := calcularPosicionHijoDer(pos)

	if posHijoIzq >= tam {
		return
	}

	posMax := posHijoIzq
	if posHijoDer < tam && cmp(arr[posHijoDer], arr[posHijoIzq]) > 0 {
		posMax = posHijoDer
	}

	if cmp(arr[posMax], arr[pos]) > 0 {
		swap(arr, pos, posMax)
		downheapRecursivo(arr, posMax, tam, cmp)
	}
}

func (heap *heap[T]) downheap(pos int) {
	downheapRecursivo(heap.datos, pos, heap.cantidad, heap.cmp)
}

func (heap *heap[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, heap.datos[:heap.cantidad])
	heap.datos = nuevosDatos
}

func heapify[T any](arr []T, cmp func(T, T) int) {
	n := len(arr)
	for i := calcularPosicionPadre(n - 1); i >= 0; i-- {
		downheapRecursivo(arr, i, n, cmp)
	}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elem T) {
	if heap.cantidad == len(heap.datos) {
		heap.redimensionar(len(heap.datos) * _FACTOR_REDIMENSION)
	}
	heap.datos[heap.cantidad] = elem
	heap.upheap(heap.cantidad)
	heap.cantidad++
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	max := heap.datos[0]
	heap.cantidad--
	heap.datos[0] = heap.datos[heap.cantidad]
	heap.downheap(0)

	if heap.cantidad > 0 && heap.cantidad*_FACTOR_REDUCCION <= len(heap.datos) && len(heap.datos) > _CAPACIDAD_INICIAL {
		heap.redimensionar(len(heap.datos) / _FACTOR_REDIMENSION)
	}

	return max
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		datos:    make([]T, _CAPACIDAD_INICIAL),
		cantidad: 0,
		cmp:      funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	capacidad := len(arreglo)
	if capacidad < _CAPACIDAD_INICIAL {
		capacidad = _CAPACIDAD_INICIAL
	}
	datos := make([]T, capacidad)
	copy(datos, arreglo)
	heapify(datos, funcion_cmp)
	return &heap[T]{
		datos:    datos,
		cantidad: len(arreglo),
		cmp:      funcion_cmp,
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)
	tam := len(elementos)
	for i := tam - 1; i > 0; i-- {
		swap(elementos, 0, i)
		downheapRecursivo(elementos, 0, i, funcion_cmp)
	}
}

// 1. Implementar una primitiva Filter(f func(T) bool) para el Heap, que deje al heap únicamente con los elementos para
// los que la función devuelva true. La primitiva debe funcionar en O(n), siendo n la cantidad de elementos inicialmente
// en el heap. Por supuesto, luego de aplicar la operación, el heap debe quedar en un estado válido para poder seguir
// operando. Justificar la complejidad de la primitiva implementada.

func (heap *heap[T]) Filter(F func(T) bool) {

}

// 2. Implementar una función eliminarRepetidos(arreglo []int) []int que dado un arreglo de números, nos devuelva
// otro en el que estén los elementos del original sin repetidos. La primera aparición debe mantenerse, y las demás no ser
// consideradas. Indicar y justificar la complejidad del algoritmo implementado.

func func_cmp(a int, b int) bool {
	return a == b
}

func eliminarRepetidos(arreglo []int) []int {
	visitados := CrearHash[int, bool](func_cmp)
	nums := make([]int, 0, len(arreglo))

	for _, i := range arreglo {
		if !visitados.Pertenece(i) {
			visitados.Guardar(i, true)
			nums = append(nums, i)
		}
	}

	return nums
}

// 3. Se tiene un árbol binario que representa la fase eliminatoria del mundial. En cada nodo guarda el nombre del país, así
// como la cantidad de goles que convirtió en dicha fase (incluyendo la tanda de penales, si fuera necesario). El padre
// del nodo debe si o si tener al hijo que ganó (tuvo mayor cantidad de goles). Implementar una primitiva para el árbol
// donde solamente están los nombres de los equipos en las hojas (no en los internos), y deje el árbol completado con los
// ganadores en cada fase. Se puede asumir que el árbol es o bien completo, o que al menos todos los nodos internos tienen
// exactamente 2 hijos. La cantidad de goles en la raíz no es relevante. La estructura del árbol es:

type Arbol struct {
	pais  string
	goles int
	izq   *Arbol
	der   *Arbol
}

// Tomando el ejemplo del dorso, si invocamos para el árbol de la izquierda, debe quedar como el de la derecha.
