package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(10)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())

	lista.InsertarPrimero(20)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
}

func TestInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(10)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())

	lista.InsertarUltimo(20)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 20, lista.VerUltimo())
}

func TestBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(10)
	lista.InsertarPrimero(20)
	lista.InsertarPrimero(30)

	elemento := lista.BorrarPrimero()
	require.Equal(t, 30, elemento)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 20, lista.VerPrimero())

	lista.BorrarPrimero()
	lista.BorrarPrimero()

	require.True(t, lista.EstaVacia())
}

func TestPanicoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.Panics(t, func() { lista.VerPrimero() })
}

func TestPanicoBorrarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.Panics(t, func() { lista.BorrarPrimero() })
}

func TestIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	contador := 0
	valor := 1

	for iter.HaySiguiente() {
		require.Equal(t, valor, iter.VerActual())
		iter.Siguiente()
		contador++
		valor++
	}

	require.Equal(t, 3, contador)
}

func TestIterar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	contador := 0
	suma := 0

	lista.Iterar(func(elemento int) bool {
		contador++
		suma += elemento
		return true
	})

	require.Equal(t, 3, contador)
	require.Equal(t, 6, suma)
}

func TestIterarCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	contador := 0

	lista.Iterar(func(elemento int) bool {
		contador++
		return elemento < 2
	})

	require.Equal(t, 2, contador)
}

func TestIteradorInsertar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(40)

	iter := lista.Iterador()
	iter.Insertar(10)
	require.Equal(t, 10, lista.VerPrimero())

	iter = lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	iter.Insertar(30)

	iter = lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(50)
	require.Equal(t, 50, lista.VerUltimo())

	arr := []int{10, 20, 30, 40, 50}
	valores := []int{}
	lista.Iterar(func(v int) bool {
		valores = append(valores, v)
		return true
	})
	for i := range arr {
		require.Equal(t, arr[i], valores[i])
	}
}

func TestIteradorBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for _, v := range []int{10, 20, 30, 40} {
		lista.InsertarUltimo(v)
	}

	iter := lista.Iterador()
	require.Equal(t, 10, iter.Borrar())

	iter = lista.Iterador()
	require.Equal(t, 20, iter.Borrar())
	require.Equal(t, 30, iter.Borrar())

	iter = lista.Iterador()
	require.Equal(t, 40, iter.Borrar())
	require.True(t, lista.EstaVacia())
}

func TestIteradorCasosBorde(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente())

	require.Panics(t, func() { iter.Borrar() })

	lista.InsertarUltimo(10)

	iter = lista.Iterador()

	iter.Siguiente()

	require.Equal(t, 10, iter.VerActual())

	iter.Borrar()

	require.True(t, lista.EstaVacia())

	iter = lista.Iterador()

	require.False(t, iter.HaySiguiente())
}
