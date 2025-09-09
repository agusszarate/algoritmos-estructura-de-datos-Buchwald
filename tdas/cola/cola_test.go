package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())

	require.Panics(t, func() { cola.VerPrimero() })

	require.Panics(t, func() { cola.Desencolar() })
}

func TestEncolarYDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()

	elementos := []any{1, "test", 3.14, true, "último"}

	for _, elem := range elementos {
		cola.Encolar(elem)
		require.Equal(t, elementos[0], cola.VerPrimero())
		require.False(t, cola.EstaVacia())
	}

	for i := 0; i < len(elementos); i++ {
		require.Equal(t, elementos[i], cola.VerPrimero())
		desencolado := cola.Desencolar()
		require.Equal(t, elementos[i], desencolado)
	}

	require.True(t, cola.EstaVacia())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()

	for i := 0; i < 100000; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}

	for i := 0; i < 100000; i++ {
		elemento := cola.Desencolar()

		if i != 99999 {
			require.Equal(t, i+1, cola.VerPrimero())
		}

		require.Equal(t, i, elemento)
	}

	require.Panics(t, func() { cola.VerPrimero() })
}

func TestColaVaciaUsada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()

	for i := 0; i < 100000; i++ {
		cola.Encolar(i)
	}

	for i := 0; i < 100000; i++ {
		elemento := cola.Desencolar()

		if i != 99999 {
			require.Equal(t, i+1, cola.VerPrimero())
		}

		require.Equal(t, i, elemento)
	}

	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() })
	require.Panics(t, func() { cola.Desencolar() })
}

func TestTiposEnteros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	enteros := []int{-5, 0, 10, 999, -1000}

	for _, num := range enteros {
		cola.Encolar(num)
		require.Equal(t, enteros[0], cola.VerPrimero())
	}

	for i := 0; i < len(enteros); i++ {
		require.Equal(t, enteros[i], cola.Desencolar())
	}
}

func TestTiposCadenas(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()

	cadenas := []string{"", "hola", "mundo", "test string", "última cadena"}

	for _, str := range cadenas {
		cola.Encolar(str)
		require.Equal(t, cadenas[0], cola.VerPrimero())
	}

	for i := 0; i < len(cadenas); i++ {
		require.Equal(t, cadenas[i], cola.Desencolar())
	}
}

func TestTiposFlotantes(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()

	flotantes := []float64{0.0, -3.14, 2.71, 999.999, -0.001}

	for _, num := range flotantes {
		cola.Encolar(num)
		require.Equal(t, flotantes[0], cola.VerPrimero())
	}

	for i := 0; i < len(flotantes); i++ {
		require.Equal(t, flotantes[i], cola.Desencolar())
	}
}

func TestTiposBooleanos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[bool]()

	valores := []bool{true, false, true, true, false}

	for _, val := range valores {
		cola.Encolar(val)
		require.Equal(t, valores[0], cola.VerPrimero())
	}

	for i := 0; i < len(valores); i++ {
		require.Equal(t, valores[i], cola.Desencolar())
	}
}

func TestEncolarVerPrimero(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()

	elementos := []string{"primero", "segundo", "tercero", "cuarto"}

	for i, elem := range elementos {
		cola.Encolar(elem)

		require.Equal(t, elementos[0], cola.VerPrimero())
		require.False(t, cola.EstaVacia())

		for j := 0; j < 3; j++ {
			require.Equal(t, elementos[0], cola.VerPrimero())
		}

		primerElem := cola.VerPrimero()
		require.Equal(t, elementos[0], primerElem)
		require.Equal(t, i+1, len(elementos[:i+1]))
	}

	primeroActual := cola.VerPrimero()
	require.Equal(t, "primero", primeroActual)

	cola.Desencolar()
	require.Equal(t, "segundo", cola.VerPrimero())
}

func TestIntercalarOperaciones(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(1)
	require.Equal(t, 1, cola.VerPrimero())

	cola.Encolar(2)
	require.Equal(t, 1, cola.VerPrimero())

	desencolado1 := cola.Desencolar()
	require.Equal(t, 1, desencolado1)
	require.Equal(t, 2, cola.VerPrimero())

	cola.Encolar(3)
	require.Equal(t, 2, cola.VerPrimero())

	cola.Encolar(4)
	require.Equal(t, 2, cola.VerPrimero())

	desencolado2 := cola.Desencolar()
	require.Equal(t, 2, desencolado2)
	require.Equal(t, 3, cola.VerPrimero())

	desencolado3 := cola.Desencolar()
	require.Equal(t, 3, desencolado3)
	require.Equal(t, 4, cola.VerPrimero())

	cola.Encolar(5)
	require.Equal(t, 4, cola.VerPrimero())

	desencolado4 := cola.Desencolar()
	require.Equal(t, 4, desencolado4)
	require.Equal(t, 5, cola.VerPrimero())

	desencolado5 := cola.Desencolar()
	require.Equal(t, 5, desencolado5)
	require.True(t, cola.EstaVacia())
}
