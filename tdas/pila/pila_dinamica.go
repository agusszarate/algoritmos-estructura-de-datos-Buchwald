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

	p.datos = append(p.datos, elemento)
	p.cantidad += 1

	p.checkRedimension()
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
	if cap(p.datos) == p.cantidad {
		p.redimensionar(cap(p.datos) * FACTOR_REDIMENSION)
	} else if p.cantidad <= cap(p.datos)/FACTOR_ACHICAR {
		p.redimensionar(cap(p.datos) / FACTOR_REDIMENSION)
	}
}

func (p *pilaDinamica[T]) redimensionar(nuevoTamaño int) {
	nuevosDatos := make([]T, p.cantidad, nuevoTamaño)

	copy(nuevosDatos, p.datos[:p.cantidad])

	p.datos = nuevosDatos
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, 0, 2),
		cantidad: 0,
	}
}
