package cola

type nodo[T any] struct {
	dato    T
	proximo *nodo[T]
}

type colaEnlazada[T any] struct {
	cantidad   int
	primerNodo *nodo[T]
	ultimoNodo *nodo[T]
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.cantidad == 0
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	cola.panicVacia()

	return cola.primerNodo.dato
}

func (cola *colaEnlazada[T]) Encolar(elemento T) {

	nuevoNodo := &nodo[T]{dato: elemento, proximo: nil}

	if cola.cantidad == 0 {
		cola.primerNodo = nuevoNodo
		cola.ultimoNodo = nuevoNodo
	} else {
		cola.ultimoNodo.proximo = nuevoNodo
		cola.ultimoNodo = nuevoNodo
	}

	cola.cantidad++

}

func (cola *colaEnlazada[T]) Desencolar() T {
	cola.panicVacia()

	elemento := cola.primerNodo
	cola.primerNodo = cola.primerNodo.proximo
	cola.cantidad--

	return elemento.dato
}

func (cola colaEnlazada[T]) panicVacia() {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func CrearColaEnlazada[T any]() Cola[T] {

	return &colaEnlazada[T]{
		cantidad:   0,
		primerNodo: nil,
		ultimoNodo: nil,
	}
}
