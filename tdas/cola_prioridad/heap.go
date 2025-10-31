package cola_prioridad

const (
	_CAPACIDAD_INICIAL  = 10
	_FACTOR_REDIMENSION = 2
	_FACTOR_REDUCCION   = 4
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func calcularPosicionPadre(posHijo int) int {
	return (posHijo - 1) / 2
}

func calcularPosicionHijoIzq(posPadre int) int {
	return 2*posPadre + 1
}

func calcularPosicionHijoDer(posPadre int) int {
	return 2*posPadre + 2
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (h *heap[T]) upheap(pos int) {
	if pos == 0 {
		return
	}
	posPadre := calcularPosicionPadre(pos)
	if h.cmp(h.datos[pos], h.datos[posPadre]) > 0 {
		swap(h.datos, pos, posPadre)
		h.upheap(posPadre)
	}
}

func downheapGenerico[T any](arr []T, pos, tam int, cmp func(T, T) int) {
	posHijoIzq := calcularPosicionHijoIzq(pos)
	posHijoDer := calcularPosicionHijoDer(pos)

	if posHijoIzq >= tam {
		return
	}

	posMax := posHijoIzq
	if posHijoDer < tam && cmp(arr[posHijoDer], arr[posHijoIzq]) > 0 {
		posMax = posHijoDer
	}

	if cmp(arr[posMax], arr[pos]) > 0 {
		swap(arr, pos, posMax)
		downheapGenerico(arr, posMax, tam, cmp)
	}
}

func (h *heap[T]) downheap(pos int) {
	downheapGenerico(h.datos, pos, h.cantidad, h.cmp)
}

func (h *heap[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, h.datos[:h.cantidad])
	h.datos = nuevosDatos
}

func heapify[T any](arr []T, cmp func(T, T) int) {
	n := len(arr)
	for i := calcularPosicionPadre(n - 1); i >= 0; i-- {
		downheapGenerico(arr, i, n, cmp)
	}
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == len(h.datos) {
		h.redimensionar(len(h.datos) * _FACTOR_REDIMENSION)
	}
	h.datos[h.cantidad] = elem
	h.upheap(h.cantidad)
	h.cantidad++
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	max := h.datos[0]
	h.cantidad--
	h.datos[0] = h.datos[h.cantidad]
	h.downheap(0)

	if h.cantidad > 0 && h.cantidad*_FACTOR_REDUCCION <= len(h.datos) && len(h.datos) > _CAPACIDAD_INICIAL {
		h.redimensionar(len(h.datos) / _FACTOR_REDIMENSION)
	}

	return max
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		datos:    make([]T, _CAPACIDAD_INICIAL),
		cantidad: 0,
		cmp:      funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, len(arreglo))
	copy(datos, arreglo)
	heapify(datos, funcion_cmp)
	return &heap[T]{
		datos:    datos,
		cantidad: len(arreglo),
		cmp:      funcion_cmp,
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)
	tam := len(elementos)
	for i := tam - 1; i > 0; i-- {
		swap(elementos, 0, i)
		downheapGenerico(elementos, 0, i, funcion_cmp)
	}
}
