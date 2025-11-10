package cola_prioridad_test

import (
	"math/rand"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmpEnteros(a, b int) int {
	return a - b
}

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())

	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

func TestEncolarYDesencolar(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)

	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())

	heap.Encolar(20)
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 20, heap.VerMax())

	heap.Encolar(5)
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 20, heap.VerMax())

	desencolado1 := heap.Desencolar()
	require.Equal(t, 20, desencolado1)
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())

	desencolado2 := heap.Desencolar()
	require.Equal(t, 10, desencolado2)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	desencolado3 := heap.Desencolar()
	require.Equal(t, 5, desencolado3)
	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
}

func TestPropiedadDeHeap(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)

	elementos := []int{15, 10, 20, 8, 25, 3, 30, 12}

	for _, elem := range elementos {
		heap.Encolar(elem)
		require.False(t, heap.EstaVacia())
	}

	require.Equal(t, len(elementos), heap.Cantidad())

	// El máximo debe ser 30
	require.Equal(t, 30, heap.VerMax())

	// Desencolamos todos y verificamos que salen en orden descendente
	anterior := heap.Desencolar()
	require.Equal(t, 30, anterior)

	for !heap.EstaVacia() {
		actual := heap.Desencolar()
		require.True(t, anterior >= actual, "Los elementos deben salir en orden descendente")
		anterior = actual
	}

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
}

func TestVolumen(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)

	n := 10000
	elementos := make([]int, n)
	for i := 0; i < n; i++ {
		elementos[i] = i
	}

	rand.Shuffle(n, func(i, j int) {
		elementos[i], elementos[j] = elementos[j], elementos[i]
	})

	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	require.Equal(t, n, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

func TestHeapVacioUsado(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)

	n := 1000
	elementos := make([]int, n)
	for i := 0; i < n; i++ {
		elementos[i] = i
	}

	rand.Shuffle(n, func(i, j int) {
		elementos[i], elementos[j] = elementos[j], elementos[i]
	})

	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })

	heap.Encolar(42)
	require.False(t, heap.EstaVacia())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 42, heap.VerMax())
}

func TestTiposEnteros(t *testing.T) {
	heap := TDAHeap.CrearHeap(cmpEnteros)

	enteros := []int{-5, 0, 10, 999, -1000, 50, -50}

	for _, num := range enteros {
		heap.Encolar(num)
	}

	require.Equal(t, len(enteros), heap.Cantidad())

	// Deben salir en orden descendente
	anterior := heap.Desencolar()
	for !heap.EstaVacia() {
		actual := heap.Desencolar()
		require.True(t, anterior >= actual)
		anterior = actual
	}
}

func TestTiposCadenas(t *testing.T) {
	cmpStr := func(a, b string) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	heap := TDAHeap.CrearHeap[string](cmpStr)

	cadenas := []string{"hola", "mundo", "algoritmos", "estructuras", "datos"}

	for _, str := range cadenas {
		heap.Encolar(str)
	}

	require.Equal(t, len(cadenas), heap.Cantidad())

	// Deben salir en orden descendente alfabético
	anterior := heap.Desencolar()
	for !heap.EstaVacia() {
		actual := heap.Desencolar()
		require.True(t, anterior >= actual)
		anterior = actual
	}
}

func TestIntercalarOperaciones(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)

	heap.Encolar(10)
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 1, heap.Cantidad())

	heap.Encolar(20)
	require.Equal(t, 20, heap.VerMax())
	require.Equal(t, 2, heap.Cantidad())

	desencolado1 := heap.Desencolar()
	require.Equal(t, 20, desencolado1)
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 1, heap.Cantidad())

	heap.Encolar(30)
	require.Equal(t, 30, heap.VerMax())
	require.Equal(t, 2, heap.Cantidad())

	heap.Encolar(5)
	require.Equal(t, 30, heap.VerMax())
	require.Equal(t, 3, heap.Cantidad())

	desencolado2 := heap.Desencolar()
	require.Equal(t, 30, desencolado2)
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 2, heap.Cantidad())

	heap.Encolar(15)
	require.Equal(t, 15, heap.VerMax())
	require.Equal(t, 3, heap.Cantidad())

	desencolado3 := heap.Desencolar()
	require.Equal(t, 15, desencolado3)
	require.Equal(t, 2, heap.Cantidad())

	desencolado4 := heap.Desencolar()
	require.Equal(t, 10, desencolado4)
	require.Equal(t, 1, heap.Cantidad())

	desencolado5 := heap.Desencolar()
	require.Equal(t, 5, desencolado5)
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
}

func TestCrearHeapArrVacio(t *testing.T) {
	arr := []int{}
	heap := TDAHeap.CrearHeapArr(arr, cmpEnteros)

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.Panics(t, func() { heap.VerMax() })
}

func TestCrearHeapArrUnElemento(t *testing.T) {
	arr := []int{42}
	heap := TDAHeap.CrearHeapArr(arr, cmpEnteros)

	require.False(t, heap.EstaVacia())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 42, heap.VerMax())
	require.Equal(t, 42, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestCrearHeapArrVariosElementos(t *testing.T) {
	arr := []int{15, 10, 20, 8, 25, 3, 30, 12}
	heap := TDAHeap.CrearHeapArr(arr, cmpEnteros)

	require.Equal(t, len(arr), heap.Cantidad())
	require.False(t, heap.EstaVacia())

	// El máximo debe ser 30
	require.Equal(t, 30, heap.VerMax())

	// Desencolamos todos y verificamos que salen en orden descendente
	anterior := heap.Desencolar()
	for !heap.EstaVacia() {
		actual := heap.Desencolar()
		require.True(t, anterior >= actual)
		anterior = actual
	}
}

func TestCrearHeapArrNoModificaArregloOriginal(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	arrCopia := make([]int, len(arr))
	copy(arrCopia, arr)

	heap := TDAHeap.CrearHeapArr(arr, cmpEnteros)

	// El arreglo original no debe ser modificado
	require.Equal(t, arrCopia, arr)

	// El heap debe funcionar correctamente
	require.Equal(t, 9, heap.VerMax())
}

func TestCrearHeapArrVolumen(t *testing.T) {
	n := 10000
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	rand.Shuffle(n, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	heap := TDAHeap.CrearHeapArr(arr, cmpEnteros)

	require.Equal(t, n, heap.Cantidad())
	require.Equal(t, n-1, heap.VerMax())

	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
}

func TestHeapSortVacio(t *testing.T) {
	arr := []int{}
	TDAHeap.HeapSort(arr, cmpEnteros)
	require.Equal(t, []int{}, arr)
}

func TestHeapSortUnElemento(t *testing.T) {
	arr := []int{42}
	TDAHeap.HeapSort(arr, cmpEnteros)
	require.Equal(t, []int{42}, arr)
}

func TestHeapSortVariosElementos(t *testing.T) {
	arr := []int{15, 10, 20, 8, 25, 3, 30, 12}
	TDAHeap.HeapSort(arr, cmpEnteros)

	esperado := []int{3, 8, 10, 12, 15, 20, 25, 30}
	require.Equal(t, esperado, arr)
}

func TestHeapSortOrdenado(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	TDAHeap.HeapSort(arr, cmpEnteros)

	esperado := []int{1, 2, 3, 4, 5}
	require.Equal(t, esperado, arr)
}

func TestHeapSortInverso(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	TDAHeap.HeapSort(arr, cmpEnteros)

	esperado := []int{1, 2, 3, 4, 5}
	require.Equal(t, esperado, arr)
}

func TestHeapSortDuplicados(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	TDAHeap.HeapSort(arr, cmpEnteros)

	esperado := []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9}
	require.Equal(t, esperado, arr)
}

func TestHeapSortNegativos(t *testing.T) {
	arr := []int{-5, 10, -3, 0, 8, -1}
	TDAHeap.HeapSort(arr, cmpEnteros)

	esperado := []int{-5, -3, -1, 0, 8, 10}
	require.Equal(t, esperado, arr)
}

func TestHeapSortVolumen(t *testing.T) {
	n := 10000
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	rand.Shuffle(n, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	TDAHeap.HeapSort(arr, cmpEnteros)

	for i := 0; i < n-1; i++ {
		require.True(t, arr[i] <= arr[i+1])
	}
}

func TestHeapSortStrings(t *testing.T) {
	cmpStr := func(a, b string) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	arr := []string{"perro", "gato", "elefante", "avestruz", "zebra"}
	TDAHeap.HeapSort(arr, cmpStr)

	esperado := []string{"avestruz", "elefante", "gato", "perro", "zebra"}
	require.Equal(t, esperado, arr)
}

func TestRedimensionamientoAutomatico(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)

	// Agregamos muchos elementos para forzar redimensionamiento
	for i := 0; i < 1000; i++ {
		heap.Encolar(i)
	}

	require.Equal(t, 1000, heap.Cantidad())

	// Quitamos la mayoría para forzar reducción
	for i := 0; i < 950; i++ {
		heap.Desencolar()
	}

	require.Equal(t, 50, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	// Verificamos que sigue funcionando correctamente
	anterior := heap.VerMax()
	for !heap.EstaVacia() {
		actual := heap.Desencolar()
		require.True(t, anterior >= actual)
		anterior = actual
	}
}

func TestEncolarVerMaxMultiple(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)

	heap.Encolar(10)
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 10, heap.VerMax())

	heap.Encolar(20)
	require.Equal(t, 20, heap.VerMax())
	require.Equal(t, 20, heap.VerMax())

	heap.Encolar(5)
	require.Equal(t, 20, heap.VerMax())
	require.Equal(t, 20, heap.VerMax())

	require.Equal(t, 3, heap.Cantidad())
}

func TestElementosIguales(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpEnteros)

	heap.Encolar(5)
	heap.Encolar(5)
	heap.Encolar(5)
	heap.Encolar(5)

	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	for i := 0; i < 4; i++ {
		require.Equal(t, 5, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
}
