package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	require.Panics(t, func() { pila.VerTope() })

	require.Panics(t, func() { pila.Desapilar() })
}

func TestApilarYDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[any]()

	elementos := []any{1, "test", 3.14, true, "último"}

	for _, elem := range elementos {
		pila.Apilar(elem)
		require.Equal(t, elem, pila.VerTope())
		require.False(t, pila.EstaVacia())
	}

	for i := len(elementos) - 1; i >= 0; i-- {
		require.Equal(t, elementos[i], pila.VerTope())
		desapilado := pila.Desapilar()
		require.Equal(t, elementos[i], desapilado)
	}

	require.True(t, pila.EstaVacia())
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[any]()

	for i := 0; i < 100000; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}

	for i := 99999; i >= 0; i-- {
		elemento := pila.Desapilar()

		if i != 0 {
			require.Equal(t, i-1, pila.VerTope())
		}

		require.Equal(t, i, elemento)
	}

	require.Panics(t, func() { pila.VerTope() })
}

func TestPilaVaciaUsada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[any]()

	for i := 0; i < 100000; i++ {
		pila.Apilar(i)
	}

	for i := 99999; i >= 0; i-- {
		elemento := pila.Desapilar()

		if i != 0 {
			require.Equal(t, i-1, pila.VerTope())
		}

		require.Equal(t, i, elemento)
	}

	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.VerTope() })
	require.Panics(t, func() { pila.Desapilar() })
}

func TestTiposEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	enteros := []int{-5, 0, 10, 999, -1000}

	for _, num := range enteros {
		pila.Apilar(num)
		require.Equal(t, num, pila.VerTope())
	}

	for i := len(enteros) - 1; i >= 0; i-- {
		require.Equal(t, enteros[i], pila.Desapilar())
	}
}

func TestTiposCadenas(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	cadenas := []string{"", "hola", "mundo", "test string", "última cadena"}

	for _, str := range cadenas {
		pila.Apilar(str)
		require.Equal(t, str, pila.VerTope())
	}

	for i := len(cadenas) - 1; i >= 0; i-- {
		require.Equal(t, cadenas[i], pila.Desapilar())
	}
}

func TestTiposFlotantes(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()

	flotantes := []float64{0.0, -3.14, 2.71, 999.999, -0.001}

	for _, num := range flotantes {
		pila.Apilar(num)
		require.Equal(t, num, pila.VerTope())
	}

	for i := len(flotantes) - 1; i >= 0; i-- {
		require.Equal(t, flotantes[i], pila.Desapilar())
	}
}

func TestTiposBooleanos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()

	valores := []bool{true, false, true, true, false}

	for _, val := range valores {
		pila.Apilar(val)
		require.Equal(t, val, pila.VerTope())
	}

	for i := len(valores) - 1; i >= 0; i-- {
		require.Equal(t, valores[i], pila.Desapilar())
	}
}

func TestApilarVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	elementos := []string{"primero", "segundo", "tercero", "cuarto"}

	for i, elem := range elementos {
		pila.Apilar(elem)

		require.Equal(t, elem, pila.VerTope())
		require.False(t, pila.EstaVacia())

		for j := 0; j < 3; j++ {
			require.Equal(t, elem, pila.VerTope())
		}

		ultimoElem := pila.VerTope()
		require.Equal(t, elem, ultimoElem)
		require.Equal(t, i+1, len(elementos[:i+1]))
	}

	topeActual := pila.VerTope()
	require.Equal(t, "cuarto", topeActual)

	pila.Desapilar()
	require.Equal(t, "tercero", pila.VerTope())
}

func TestIntercalarOperaciones(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(1)
	require.Equal(t, 1, pila.VerTope())

	pila.Apilar(2)
	require.Equal(t, 2, pila.VerTope())

	desapilado1 := pila.Desapilar()
	require.Equal(t, 2, desapilado1)
	require.Equal(t, 1, pila.VerTope())

	pila.Apilar(3)
	require.Equal(t, 3, pila.VerTope())

	pila.Apilar(4)
	require.Equal(t, 4, pila.VerTope())

	desapilado2 := pila.Desapilar()
	require.Equal(t, 4, desapilado2)
	require.Equal(t, 3, pila.VerTope())

	desapilado3 := pila.Desapilar()
	require.Equal(t, 3, desapilado3)
	require.Equal(t, 1, pila.VerTope())

	pila.Apilar(5)
	require.Equal(t, 5, pila.VerTope())

	desapilado4 := pila.Desapilar()
	require.Equal(t, 5, desapilado4)
	require.Equal(t, 1, pila.VerTope())

	desapilado5 := pila.Desapilar()
	require.Equal(t, 1, desapilado5)
	require.True(t, pila.EstaVacia())
}
