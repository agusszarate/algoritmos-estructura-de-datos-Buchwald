package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}
type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorListaEnlazada[T any] struct{}

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

func (lista listaEnlazada[T]) Iterador() IteradorLista[T] {
}

func (lista listaEnlazada[T]) CrearListaEnlazada() *listaEnlazada[T] {
	return &listaEnlazada[T]{largo: 0, primero: nil, ultimo: nil}
}
