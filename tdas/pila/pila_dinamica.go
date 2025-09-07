package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	FACTOR_REDIMENSION = 2
	FACTOR_ACHICAR     = 4
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
	p.checkRedimension()

	p.datos = append(p.datos, elemento)
	p.cantidad += 1
}

func (p *pilaDinamica[T]) Desapilar() T {
	p.panicVacia()

	elemento := p.datos[p.cantidad-1]
	p.datos = p.datos[:p.cantidad-1]
	p.cantidad -= 1

	p.checkRedimension()

	return elemento
}

func (p *pilaDinamica[T]) panicVacia() {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
}

func (p *pilaDinamica[T]) checkRedimension() {
	if len(p.datos) == p.cantidad {
		p.redimensionar(FACTOR_REDIMENSION)
	} else if p.cantidad <= cap(p.datos)/FACTOR_ACHICAR {
		p.redimensionar(1 / FACTOR_REDIMENSION)
	}
}

func (p *pilaDinamica[T]) redimensionar(factor int) {
	nuevosDatos := make([]T, cap(p.datos)*factor)

	copy(nuevosDatos, p.datos[:p.cantidad])

	p.datos = nuevosDatos
}

func CrearPilaDinamica[T any]() Pila[T] {
	return new(pilaDinamica[T])
}
