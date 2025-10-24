package diccionario

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
	pila  []*nodoAbb[K, V]
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

func (arbol *abb[K, V]) insertarNodo(nodo **nodoAbb[K, V], clave K, dato V) {
	*nodo = arbol.crearNodo(clave, dato)
	arbol.cantidad++
}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	if arbol.raiz == nil {
		arbol.insertarNodo(&arbol.raiz, clave, dato)
		return
	}
	arbol.guardarNodo(arbol.raiz, clave, dato)
}

func (arbol *abb[K, V]) guardarNodo(nodo *nodoAbb[K, V], clave K, dato V) {
	comparacion := arbol.cmp(clave, nodo.clave)

	if comparacion == 0 {
		nodo.dato = dato
		return
	}

	var hijo **nodoAbb[K, V]
	if comparacion < 0 {
		hijo = &nodo.izquierdo
	} else {
		hijo = &nodo.derecho
	}

	if *hijo == nil {
		arbol.insertarNodo(hijo, clave, dato)
		return
	}

	arbol.guardarNodo(*hijo, clave, dato)
}

func buscarReemplazante[K any, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {

	if nodo.izquierdo == nil {
		return nodo
	}

	return buscarReemplazante(nodo.izquierdo)
}

func (arbol *abb[K, V]) borrarNodo(nodo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], V) {
	if nodo == nil {
		arbol.panicNoPertenece()
	}

	comparacion := arbol.cmp(clave, nodo.clave)

	if comparacion != 0 {
		if comparacion < 0 {
			nuevoIzq, dato := arbol.borrarNodo(nodo.izquierdo, clave)
			nodo.izquierdo = nuevoIzq
			return nodo, dato
		}
		nuevoDer, dato := arbol.borrarNodo(nodo.derecho, clave)
		nodo.derecho = nuevoDer
		return nodo, dato
	}

	dato := nodo.dato

	if nodo.izquierdo == nil && nodo.derecho == nil {
		return nil, dato
	}

	if nodo.izquierdo == nil {
		return nodo.derecho, dato
	}

	if nodo.derecho == nil {
		return nodo.izquierdo, dato
	}

	reemplazante := buscarReemplazante(nodo.derecho)
	nodo.clave = reemplazante.clave
	nodo.dato = reemplazante.dato
	nodo.derecho, _ = arbol.borrarNodo(nodo.derecho, reemplazante.clave)
	return nodo, dato
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	nuevoRaiz, dato := arbol.borrarNodo(arbol.raiz, clave)
	arbol.raiz = nuevoRaiz
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
		pila:  make([]*nodoAbb[K, V], 0),
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
		iter.pila = append(iter.pila, nodo)
		nodo = nodo.izquierdo
	}
}

func (iter *iterAbb[K, V]) panicFin() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return len(iter.pila) > 0
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	iter.panicFin()
	nodoActual := iter.pila[len(iter.pila)-1]
	return nodoActual.clave, nodoActual.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	iter.panicFin()
	nodoActual := iter.pila[len(iter.pila)-1]
	iter.pila = iter.pila[:len(iter.pila)-1]
	iter.apilarIzquierdos(nodoActual.derecho)
}
