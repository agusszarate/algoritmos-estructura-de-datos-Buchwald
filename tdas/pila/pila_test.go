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

	pila.Apilar(1)
	require.Equal(t, 1, pila.Desapilar())

	pila.Apilar("test")
	require.Equal(t, "test", pila.Desapilar())

	pila.Apilar(1)
	pila.Apilar("1")
	require.Equal(t, "1", pila.Desapilar())
	require.Equal(t, 1, pila.Desapilar())
}
