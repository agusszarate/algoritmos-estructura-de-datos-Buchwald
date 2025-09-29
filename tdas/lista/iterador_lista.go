package lista

type IteradorLista[T any] interface {

	// VerActual obtiene el valor del elemento actual. Si no hay siguiente elemento, entra en pánico con un mensaje
	// "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve verdadero si hay más elementos por iterar, false en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza al siguiente elemento y devuelve su valor. Si no hay siguiente elemento, entra en pánico con un mensaje
	// "El iterador termino de iterar".
	Siguiente() T

	// Insertar agrega un nuevo elemento en la posición actual del iterador.
	Insertar(T)

	// Borrar borra el elemento actual del iterador y devuelve su valor. Si no hay elemento actual, entra en pánico.
	Borrar() T
}
