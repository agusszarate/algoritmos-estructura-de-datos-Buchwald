package diccionario

import TDAPila "tdas/pila"

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

func CrearABB[K any, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: funcionCmp}
}

func (arbol *abb[K, V]) panicNoPertenece() {
	panic("La clave no pertenece al diccionario")
}

func (arbol *abb[K, V]) crearNodo(clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, clave: clave, dato: dato}
}

func (arbol *abb[K, V]) buscarNodo(nodo *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}
	comparacion := arbol.cmp(clave, nodo.clave)
	if comparacion == 0 {
		return nodo
	}
	if comparacion < 0 {
		return arbol.buscarNodo(nodo.izquierdo, clave)
	}
	return arbol.buscarNodo(nodo.derecho, clave)
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	return arbol.buscarNodo(arbol.raiz, clave) != nil
}

func (arbol *abb[K, V]) Obtener(clave K) V {
	nodo := arbol.buscarNodo(arbol.raiz, clave)
	if nodo == nil {
		arbol.panicNoPertenece()
	}
	return nodo.dato
}

func (arbol *abb[K, V]) guardarNodo(nodo *nodoAbb[K, V], clave K, dato V) (*nodoAbb[K, V], bool) {
	if nodo == nil {
		return arbol.crearNodo(clave, dato), true
	}
	comparacion := arbol.cmp(clave, nodo.clave)
	if comparacion == 0 {
		nodo.dato = dato
		return nodo, false
	}
	if comparacion < 0 {
		nuevoIzquierdo, esNuevo := arbol.guardarNodo(nodo.izquierdo, clave, dato)
		nodo.izquierdo = nuevoIzquierdo
		return nodo, esNuevo
	}
	nuevoDerecho, esNuevo := arbol.guardarNodo(nodo.derecho, clave, dato)
	nodo.derecho = nuevoDerecho
	return nodo, esNuevo
}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	nuevoNodo, esNuevo := arbol.guardarNodo(arbol.raiz, clave, dato)
	arbol.raiz = nuevoNodo
	if esNuevo {
		arbol.cantidad++
	}
}

func buscarReemplazante[K any, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.izquierdo == nil {
		return nodo
	}
	return buscarReemplazante(nodo.izquierdo)
}

func (arbol *abb[K, V]) borrarNodo(nodo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], bool) {
	if nodo == nil {
		return nil, false
	}

	comparacion := arbol.cmp(clave, nodo.clave)
	if comparacion < 0 {
		nuevoIzquierdo, encontrado := arbol.borrarNodo(nodo.izquierdo, clave)
		nodo.izquierdo = nuevoIzquierdo
		return nodo, encontrado
	}
	if comparacion > 0 {
		nuevoDerecho, encontrado := arbol.borrarNodo(nodo.derecho, clave)
		nodo.derecho = nuevoDerecho
		return nodo, encontrado
	}

	if nodo.izquierdo == nil && nodo.derecho == nil {
		return nil, true
	}

	if nodo.izquierdo == nil {
		return nodo.derecho, true
	}
	if nodo.derecho == nil {
		return nodo.izquierdo, true
	}

	reemplazante := buscarReemplazante(nodo.derecho)
	nodo.clave = reemplazante.clave
	nodo.dato = reemplazante.dato
	nuevoDerecho, _ := arbol.borrarNodo(nodo.derecho, reemplazante.clave)
	nodo.derecho = nuevoDerecho
	return nodo, true
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	nodo := arbol.buscarNodo(arbol.raiz, clave)
	if nodo == nil {
		arbol.panicNoPertenece()
	}
	dato := nodo.dato
	nuevoNodo, _ := arbol.borrarNodo(arbol.raiz, clave)
	arbol.raiz = nuevoNodo
	arbol.cantidad--
	return dato
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func iterarInOrder[K any, V any](nodo *nodoAbb[K, V], visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if !iterarInOrder(nodo.izquierdo, visitar) {
		return false
	}
	if !visitar(nodo.clave, nodo.dato) {
		return false
	}
	return iterarInOrder(nodo.derecho, visitar)
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterarInOrder(arbol.raiz, visitar)
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

	enRango := true
	if desde != nil && cmp(nodo.clave, *desde) < 0 {
		enRango = false
	}
	if hasta != nil && cmp(nodo.clave, *hasta) > 0 {
		enRango = false
	}

	if enRango && !visitar(nodo.clave, nodo.dato) {
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