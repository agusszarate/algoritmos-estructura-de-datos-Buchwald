package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	_FACTOR_REDIMENSION = 2
	_FACTOR_ACHICAR     = 4
	_LARGO_INICIAL      = 5
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	p.panicVacia()
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elemento T) {

	if len(p.datos) == p.cantidad {
		p.redimensionar(len(p.datos) * _FACTOR_REDIMENSION)
	}

	p.datos[p.cantidad] = elemento
	p.cantidad++

}

func (p *pilaDinamica[T]) Desapilar() T {
	p.panicVacia()

	p.cantidad--
	elemento := p.datos[p.cantidad]

	nuevoTamaño := cap(p.datos) / _FACTOR_REDIMENSION

	if p.cantidad <= cap(p.datos)/_FACTOR_ACHICAR && nuevoTamaño >= _LARGO_INICIAL {
		p.redimensionar(nuevoTamaño)
	}

	return elemento
}

func (p *pilaDinamica[T]) panicVacia() {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
}

func (p *pilaDinamica[T]) redimensionar(nuevoTamaño int) {
	nuevosDatos := make([]T, nuevoTamaño)

	copy(nuevosDatos, p.datos[:p.cantidad])

	p.datos = nuevosDatos
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, _LARGO_INICIAL),
		cantidad: 0,
	}
}
