package lista

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

func (iter *iteradorListaEnlazada[T]) Siguiente() T {
	iter.panicFin()

	valor := iter.actual.dato
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
	return valor
}

func (iter *iteradorListaEnlazada[T]) Insertar(elemento T) {
	nuevo := iter.lista.crearNodo(elemento, iter.actual)

	if iter.anterior == nil {
		iter.lista.primero = nuevo
	} else {
		iter.anterior.siguiente = nuevo
	}

	if iter.actual.siguiente == nil {
		iter.lista.ultimo = nuevo
	}

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
