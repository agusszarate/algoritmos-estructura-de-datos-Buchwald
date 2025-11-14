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

func (heap *heap[T]) upheap(pos int) {
	if pos == 0 {
		return
	}
	posPadre := calcularPosicionPadre(pos)
	if heap.cmp(heap.datos[pos], heap.datos[posPadre]) > 0 {
		swap(heap.datos, pos, posPadre)
		heap.upheap(posPadre)
	}
}

func downheapRecursivo[T any](arr []T, pos, tam int, cmp func(T, T) int) {
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
		downheapRecursivo(arr, posMax, tam, cmp)
	}
}

func (heap *heap[T]) downheap(pos int) {
	downheapRecursivo(heap.datos, pos, heap.cantidad, heap.cmp)
}

func (heap *heap[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, heap.datos[:heap.cantidad])
	heap.datos = nuevosDatos
}

func heapify[T any](arr []T, cmp func(T, T) int) {
	n := len(arr)
	for i := calcularPosicionPadre(n - 1); i >= 0; i-- {
		downheapRecursivo(arr, i, n, cmp)
	}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elem T) {
	if heap.cantidad == len(heap.datos) {
		heap.redimensionar(len(heap.datos) * _FACTOR_REDIMENSION)
	}
	heap.datos[heap.cantidad] = elem
	heap.upheap(heap.cantidad)
	heap.cantidad++
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	max := heap.datos[0]
	heap.cantidad--
	heap.datos[0] = heap.datos[heap.cantidad]
	heap.downheap(0)

	if heap.cantidad > 0 && heap.cantidad*_FACTOR_REDUCCION <= len(heap.datos) && len(heap.datos) > _CAPACIDAD_INICIAL {
		heap.redimensionar(len(heap.datos) / _FACTOR_REDIMENSION)
	}

	return max
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		datos:    make([]T, _CAPACIDAD_INICIAL),
		cantidad: 0,
		cmp:      funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	capacidad := len(arreglo)
	if capacidad < _CAPACIDAD_INICIAL {
		capacidad = _CAPACIDAD_INICIAL
	}
	datos := make([]T, capacidad)
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
		downheapRecursivo(elementos, 0, i, funcion_cmp)
	}
}
