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

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0 && lista.primero == nil && lista.ultimo == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemeneto T) {

}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {

}

func (lista *listaEnlazada[T]) BorrarPrimero() T {

}

func (lista listaEnlazada[T]) VerPrimero() T {

}

func (lista listaEnlazada[T]) VerUltimo() T {

}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}
