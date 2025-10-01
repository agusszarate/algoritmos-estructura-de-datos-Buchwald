package lista

//--------------------------------Lista Enlazada--------------------------------//

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}
type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func (lista listaEnlazada[T]) crearNodo(elemento T, siguiente *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: elemento, siguiente: siguiente}
}

func (lista listaEnlazada[T]) panicVacia() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0 && lista.primero == nil && lista.ultimo == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	var siguiente *nodoLista[T] = nil

	if !lista.EstaVacia() {
		siguiente = lista.primero
	}

	nodo := lista.crearNodo(elemento, siguiente)

	if lista.EstaVacia() {
		lista.ultimo = nodo
	}

	lista.primero = nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nodo := lista.crearNodo(elemento, nil)

	if lista.EstaVacia() {
		lista.primero = nodo
	} else {
		lista.ultimo.siguiente = nodo
	}

	lista.ultimo = nodo

	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.panicVacia()

	valor := lista.primero.dato

	lista.primero = lista.primero.siguiente

	lista.largo--

	if lista.largo == 0 {
		lista.ultimo = nil
	}

	return valor
}

func (lista listaEnlazada[T]) VerPrimero() T {
	lista.panicVacia()
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	lista.panicVacia()
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {

	nodoActual := lista.primero

	for nodoActual != nil {
		resultado := visitar(nodoActual.dato)

		nodoActual = nodoActual.siguiente

		if !resultado {
			nodoActual = nil
		}
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{
		lista:    lista,
		actual:   lista.primero,
		anterior: nil,
	}
}

//--------------------------------Iterador--------------------------------//

type iteradorListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func (iter *iteradorListaEnlazada[T]) panicFin() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iter *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorListaEnlazada[T]) VerActual() T {
	iter.panicFin()
	return iter.actual.dato
}

func (iter *iteradorListaEnlazada[T]) Siguiente() {
	iter.panicFin()

	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorListaEnlazada[T]) Insertar(elemento T) {
	nuevo := iter.lista.crearNodo(elemento, iter.actual)

	if iter.anterior == nil {
		iter.lista.primero = nuevo
	} else {
		iter.anterior.siguiente = nuevo
	}

	if iter.actual == nil {
		iter.lista.ultimo = nuevo
	}

	iter.actual = nuevo
	iter.lista.largo++
}

func (iter *iteradorListaEnlazada[T]) Borrar() T {
	iter.panicFin()

	valor := iter.actual.dato

	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
	}

	if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}

	iter.actual = iter.actual.siguiente
	iter.lista.largo--

	return valor
}