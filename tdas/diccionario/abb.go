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
