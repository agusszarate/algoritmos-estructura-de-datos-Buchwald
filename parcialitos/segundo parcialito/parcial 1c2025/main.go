package main

import (
	. "tdas/diccionario"
	TDAPila "tdas/pila"
)

type nodoAbb[K any, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K any, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iterAbb[K any, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   func(K, K) int
}

func (arbol *abb[K, V]) panicNoPertenece() {
	panic("La clave no pertenece al diccionario")
}

func (arbol *abb[K, V]) crearNodo(clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, clave: clave, dato: dato}
}

func (arbol *abb[K, V]) buscarNodoConPadre(clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	var padre *nodoAbb[K, V] = nil
	actual := arbol.raiz
	for actual != nil {
		cmp := arbol.cmp(clave, actual.clave)
		if cmp == 0 {
			return actual, padre
		}
		padre = actual
		if cmp < 0 {
			actual = actual.izquierdo
		} else {
			actual = actual.derecho
		}
	}
	return nil, padre
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := arbol.buscarNodoConPadre(clave)
	return nodo != nil
}

func (arbol *abb[K, V]) Obtener(clave K) V {
	nodo, _ := arbol.buscarNodoConPadre(clave)
	if nodo == nil {
		arbol.panicNoPertenece()
	}
	return nodo.dato
}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	nodo, padre := arbol.buscarNodoConPadre(clave)

	if nodo != nil {
		nodo.dato = dato
		return
	}

	nuevo := arbol.crearNodo(clave, dato)
	arbol.cantidad++

	if padre == nil {
		arbol.raiz = nuevo
		return
	}

	if arbol.cmp(clave, padre.clave) < 0 {
		padre.izquierdo = nuevo
	} else {
		padre.derecho = nuevo
	}
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	nodo, padre := arbol.buscarNodoConPadre(clave)
	if nodo == nil {
		arbol.panicNoPertenece()
	}
	dato := nodo.dato
	arbol.cantidad--

	if nodo.izquierdo != nil && nodo.derecho != nil {
		reemplazo := nodo.derecho
		for reemplazo.izquierdo != nil {
			reemplazo = reemplazo.izquierdo
		}

		claveReemplazo := reemplazo.clave
		datoReemplazo := reemplazo.dato

		arbol.Borrar(claveReemplazo)

		nodo.clave = claveReemplazo
		nodo.dato = datoReemplazo

		arbol.cantidad++

		return dato
	}

	var hijo *nodoAbb[K, V]
	if nodo.izquierdo != nil {
		hijo = nodo.izquierdo
	} else {
		hijo = nodo.derecho
	}

	if padre == nil {
		arbol.raiz = hijo
	} else if padre.izquierdo == nodo {
		padre.izquierdo = hijo
	} else {
		padre.derecho = hijo
	}

	return dato
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	arbol.IterarRango(nil, nil, visitar)
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}

func iterarRango[K any, V any](nodo *nodoAbb[K, V], desde *K, hasta *K, cmp func(K, K) int, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}

	if desde != nil && cmp(nodo.clave, *desde) < 0 {
		return iterarRango(nodo.derecho, desde, hasta, cmp, visitar)
	}

	if hasta != nil && cmp(nodo.clave, *hasta) > 0 {
		return iterarRango(nodo.izquierdo, desde, hasta, cmp, visitar)
	}

	if !iterarRango(nodo.izquierdo, desde, hasta, cmp, visitar) {
		return false
	}

	if !visitar(nodo.clave, nodo.dato) {
		return false
	}

	return iterarRango(nodo.derecho, desde, hasta, cmp, visitar)
}

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRango(arbol.raiz, desde, hasta, arbol.cmp, visitar)
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{
		pila:  TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
		desde: desde,
		hasta: hasta,
		cmp:   arbol.cmp,
	}
	iter.apilarIzquierdos(arbol.raiz)
	return iter
}

func (iter *iterAbb[K, V]) apilarIzquierdos(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if iter.desde != nil && iter.cmp(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.derecho
			continue
		}
		if iter.hasta != nil && iter.cmp(nodo.clave, *iter.hasta) > 0 {
			nodo = nodo.izquierdo
			continue
		}
		iter.pila.Apilar(nodo)
		nodo = nodo.izquierdo
	}
}

func (iter *iterAbb[K, V]) panicFin() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	iter.panicFin()
	nodoActual := iter.pila.VerTope()
	return nodoActual.clave, nodoActual.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	iter.panicFin()
	nodoActual := iter.pila.Desapilar()
	iter.apilarIzquierdos(nodoActual.derecho)
}

// 1. Implementar una primitiva para el ABB func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K que reciba
// 2 claves y devuelva el último ancestro en común entre ambas claves. Dicho ancestro en común podría ser incluso alguna
// de estas claves. Si alguna clave pasada no se encuentra en el árbol, finalizar con panic. Indicar y justificar la complejidad
// de la primitiva implementada.

// Mostramos ejemplos de resultados esperados de invocar la primitiva al árbol del dorso:
//
// arbol.AncestroComun(1, 4) --> 2
// arbol.AncestroComun(2, 4) --> 2
// arbol.AncestroComun(9, 1) --> 5

func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K {
	if !arbol.Pertenece(clave1) || !arbol.Pertenece(clave2) {
		panic("Una de las claves no pertenece al arbol")
	} else {
		return arbol.ancestrosEnComun(clave1, clave2, arbol.raiz)
	}
}

func (arbol *abb[K, V]) ancestrosEnComun(clave1, clave2 K, nodo *nodoAbb[K, V]) K {
	comp1 := arbol.cmp(clave1, nodo.clave)
	comp2 := arbol.cmp(clave2, nodo.clave)

	if comp1 > 0 && comp2 > 0 {
		return arbol.ancestrosEnComun(clave1, clave2, nodo.derecho)
	} else if comp1 < 0 && comp2 < 0 {
		return arbol.ancestrosEnComun(clave1, clave2, nodo.izquierdo)
	} else {
		return nodo.clave
	}
}

// 2. Implementar una función func minimoExcluido(arr []int) int que dado un arreglo de valores enteros (mayores o
// iguales a 0), obtenga el mínimo valor que no se encuentre en el arreglo. Indicar y justificar la complejidad del algoritmo
// (explicar en detalle este paso, porque es fácil que se te puedan pasar detalles importantes a explicar).
//
// Por ejemplo:
//
// minimoExcluido([]int{0, 5, 1}) --> 2
// minimoExcluido([]int{3, 5, 1}) --> 0
// minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2}) --> 6
// minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2, 12345675433221345}) --> 6

func func_cmp(a int, b int) bool {
	return a == b
}

func minimoExcluido(arr []int) int {

	presentes := CrearHash[int, bool](func_cmp)

	for _, num := range arr {
		presentes.Guardar(num, true)
	}

	for i := range len(arr) + 1 {

		if !presentes.Pertenece(i) {
			return i
		}
	}

	return -1
}

// 3. Realizar el seguimiento de las siguientes operaciones sobre un heap de máximos:
//
// a. Construir el heap a partir del arreglo [5, 3, 6, 1, 4, 7, 8, 10]

// copia del arreglo

// el ultimo padre con propiedad de heap = pos 7(10)
// buscamos el padre (7 -1) / 2 = 3 (1)
// iteramos desde 3 hasta 0
// miramos sus hijos
// izq (3 * 2) + 1 = 7
// der (3 * 2) + 2 = 8 no existe
// 10 > 1
// swap pos 3(1) con 7(10)
// recursivamente buscamos si tiene hijos y hacemos downheap
// [5, 3, 6, 10, 4, 7, 8, 1]
// ahora pos 2(6)
// izq (2 * 2) + 1 = 5 (7)
// der (2 * 2) + 2 = 6 (8)
// max entre 7 y 8 = 8
// 8 > 6
// swap pos 2(6) con pos 6(8)
// [5, 3, 8, 10, 4, 7, 6, 1]
// pos 1(3)
// izq (1 * 2) + 1 = 3 (10)
// der (1 * 2) + 2 = 4 (4)
// max entre 10 y 4 = 10
// 10 > 3
// swap pos 1(3) con pos 3(10)
// [5, 10, 8, 3, 4, 7, 6, 1]
// pos 0(5)
// izq (0 * 2) + 1 = 1 (10)
// der (0 * 2) + 2 = 2 (8)
// max entre 10 y 8 = 10
// 10 > 5
// swap pos 0(5) con pos 1(10)
// [10, 5, 8, 3, 4, 7, 6, 1]

// b. Sobre el heap resultante del punto (a): Encolar 7, Desencolar, Desencolar, Encolar 13, Desencolar.
