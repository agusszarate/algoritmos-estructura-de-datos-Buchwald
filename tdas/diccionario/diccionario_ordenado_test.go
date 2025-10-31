package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_ABB = []int{1000, 2000, 4000, 8000, 16000}

func comparacionStrings(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func comparacionInts(a, b int) int {
	return a - b
}

func TestABBVacio(t *testing.T) {
	t.Log("Comprueba que ABB vacio no tiene claves")
	abb := TDADiccionario.CrearABB[string, string](comparacionStrings)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestABBClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, sigue sin existir")
	abb := TDADiccionario.CrearABB[string, string](comparacionStrings)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("") })

	abbNum := TDADiccionario.CrearABB[int, string](comparacionInts)
	require.False(t, abbNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Borrar(0) })
}

func TestABBUnElemento(t *testing.T) {
	t.Log("Comprueba que ABB con un elemento tiene esa Clave, unicamente")
	abb := TDADiccionario.CrearABB[string, int](comparacionStrings)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("B"))
	require.EqualValues(t, 10, abb.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("B") })
}

func TestABBGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	abb := TDADiccionario.CrearABB[string, string](comparacionStrings)
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))

	require.False(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[1], valores[1])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))

	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], abb.Obtener(claves[2]))
}

func TestABBReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	abb := TDADiccionario.CrearABB[string, string](comparacionStrings)
	abb.Guardar(clave, "miau")
	abb.Guardar(clave2, "guau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, "miau", abb.Obtener(clave))
	require.EqualValues(t, "guau", abb.Obtener(clave2))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(clave, "miu")
	abb.Guardar(clave2, "baubau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener(clave))
	require.EqualValues(t, "baubau", abb.Obtener(clave2))
}

func TestABBBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se los borra, revisando que en todo momento el ABB se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](comparacionStrings)

	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[0]) })
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[0]) })

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[1]) })
}

func TestABBBorrarHoja(t *testing.T) {
	t.Log("Prueba borrar nodos hoja")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)
	abb.Guardar(10, "diez")
	abb.Guardar(5, "cinco")
	abb.Guardar(15, "quince")

	// Borro una hoja
	require.EqualValues(t, "cinco", abb.Borrar(5))
	require.False(t, abb.Pertenece(5))
	require.True(t, abb.Pertenece(10))
	require.True(t, abb.Pertenece(15))
	require.EqualValues(t, 2, abb.Cantidad())
}

func TestABBBorrarUnHijo(t *testing.T) {
	t.Log("Prueba borrar nodos con un solo hijo")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)
	abb.Guardar(10, "diez")
	abb.Guardar(5, "cinco")
	abb.Guardar(15, "quince")
	abb.Guardar(3, "tres")

	// Borro nodo con un hijo izquierdo
	require.EqualValues(t, "cinco", abb.Borrar(5))
	require.False(t, abb.Pertenece(5))
	require.True(t, abb.Pertenece(3))
	require.True(t, abb.Pertenece(10))
	require.True(t, abb.Pertenece(15))
	require.EqualValues(t, 3, abb.Cantidad())
}

func TestABBBorrarDosHijos(t *testing.T) {
	t.Log("Prueba borrar nodos con dos hijos")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)
	abb.Guardar(10, "diez")
	abb.Guardar(5, "cinco")
	abb.Guardar(15, "quince")
	abb.Guardar(3, "tres")
	abb.Guardar(7, "siete")
	abb.Guardar(12, "doce")
	abb.Guardar(17, "diecisiete")

	// Borro la raiz (tiene dos hijos)
	require.EqualValues(t, "diez", abb.Borrar(10))
	require.False(t, abb.Pertenece(10))
	require.True(t, abb.Pertenece(3))
	require.True(t, abb.Pertenece(5))
	require.True(t, abb.Pertenece(7))
	require.True(t, abb.Pertenece(12))
	require.True(t, abb.Pertenece(15))
	require.True(t, abb.Pertenece(17))
	require.EqualValues(t, 6, abb.Cantidad())
}

func TestABBReutilizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de reinserción de un elemento borrado")
	abb := TDADiccionario.CrearABB[string, string](comparacionStrings)
	clave := "hola"
	abb.Guardar(clave, "mundo!")
	abb.Borrar(clave)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(clave))
	abb.Guardar(clave, "mundooo!")
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, "mundooo!", abb.Obtener(clave))
}

func TestABBConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)
	clave := 10
	valor := "Gatito"

	abb.Guardar(clave, valor)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, valor, abb.Obtener(clave))
	require.EqualValues(t, valor, abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestABBClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	abb := TDADiccionario.CrearABB[string, string](comparacionStrings)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}

func TestABBValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	abb := TDADiccionario.CrearABB[string, *int](comparacionStrings)
	clave := "Pez"
	abb.Guardar(clave, nil)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, (*int)(nil), abb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestABBIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas en orden con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	abb := TDADiccionario.CrearABB[string, *int](comparacionStrings)
	abb.Guardar(clave2, nil)
	abb.Guardar(clave1, nil)
	abb.Guardar(clave3, nil)

	cs := make([]string, 3)
	cantidad := 0

	abb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		cantidad++
		return true
	})

	require.EqualValues(t, 3, cantidad)
	// Deben venir en orden alfabético
	require.EqualValues(t, "Gato", cs[0])
	require.EqualValues(t, "Perro", cs[1])
	require.EqualValues(t, "Vaca", cs[2])
}

func TestABBIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridos correctamente en orden con el iterador interno")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)
	abb.Guardar(5, "cinco")
	abb.Guardar(3, "tres")
	abb.Guardar(7, "siete")
	abb.Guardar(1, "uno")
	abb.Guardar(9, "nueve")

	resultado := ""
	abb.Iterar(func(_ int, dato string) bool {
		resultado += dato + " "
		return true
	})

	// Deben venir en orden numérico
	require.EqualValues(t, "uno tres cinco siete nueve ", resultado)
}

func TestABBIteradorInternoCorte(t *testing.T) {
	t.Log("Valida que el iterador interno se detenga cuando la función visitar retorna false")
	abb := TDADiccionario.CrearABB[int, int](comparacionInts)

	for i := 0; i < 100; i++ {
		abb.Guardar(i, i)
	}

	cantidad := 0
	abb.Iterar(func(clave int, dato int) bool {
		cantidad++
		return clave < 50 // Corta cuando llega a 50
	})

	require.EqualValues(t, 51, cantidad) // Debería haber iterado hasta el 50 inclusive
}

func TestABBIteradorExterno(t *testing.T) {
	t.Log("Prueba de iterador externo sin rango")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)
	abb.Guardar(5, "cinco")
	abb.Guardar(3, "tres")
	abb.Guardar(7, "siete")
	abb.Guardar(1, "uno")
	abb.Guardar(9, "nueve")

	iter := abb.Iterador()

	// Verificar orden
	require.True(t, iter.HaySiguiente())
	clave1, valor1 := iter.VerActual()
	require.EqualValues(t, 1, clave1)
	require.EqualValues(t, "uno", valor1)

	iter.Siguiente()
	clave2, valor2 := iter.VerActual()
	require.EqualValues(t, 3, clave2)
	require.EqualValues(t, "tres", valor2)

	iter.Siguiente()
	clave3, valor3 := iter.VerActual()
	require.EqualValues(t, 5, clave3)
	require.EqualValues(t, "cinco", valor3)

	iter.Siguiente()
	clave4, valor4 := iter.VerActual()
	require.EqualValues(t, 7, clave4)
	require.EqualValues(t, "siete", valor4)

	iter.Siguiente()
	clave5, valor5 := iter.VerActual()
	require.EqualValues(t, 9, clave5)
	require.EqualValues(t, "nueve", valor5)

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorExternoVacio(t *testing.T) {
	t.Log("Iterar sobre ABB vacio es simplemente tenerlo al final")
	abb := TDADiccionario.CrearABB[string, int](comparacionStrings)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIterarRango(t *testing.T) {
	t.Log("Prueba de IterarRango con diferentes rangos")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)

	// Inserto elementos desordenados
	abb.Guardar(10, "diez")
	abb.Guardar(5, "cinco")
	abb.Guardar(15, "quince")
	abb.Guardar(3, "tres")
	abb.Guardar(7, "siete")
	abb.Guardar(12, "doce")
	abb.Guardar(17, "diecisiete")
	abb.Guardar(1, "uno")

	// Test con rango [5, 12]
	desde := 5
	hasta := 12
	claves := []int{}
	abb.IterarRango(&desde, &hasta, func(clave int, _ string) bool {
		claves = append(claves, clave)
		return true
	})

	require.Equal(t, []int{5, 7, 10, 12}, claves)
}

func TestABBIterarRangoSinDesde(t *testing.T) {
	t.Log("Prueba de IterarRango sin límite inferior")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)

	for i := 0; i < 10; i++ {
		abb.Guardar(i, fmt.Sprintf("%d", i))
	}

	hasta := 5
	var claves []int
	abb.IterarRango(nil, &hasta, func(clave int, _ string) bool {
		claves = append(claves, clave)
		return true
	})

	require.Equal(t, []int{0, 1, 2, 3, 4, 5}, claves)
}

func TestABBIterarRangoSinHasta(t *testing.T) {
	t.Log("Prueba de IterarRango sin límite superior")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)

	for i := 0; i < 10; i++ {
		abb.Guardar(i, fmt.Sprintf("%d", i))
	}

	desde := 5
	var claves []int
	abb.IterarRango(&desde, nil, func(clave int, _ string) bool {
		claves = append(claves, clave)
		return true
	})

	require.Equal(t, []int{5, 6, 7, 8, 9}, claves)
}

func TestABBIterarRangoSinLimites(t *testing.T) {
	t.Log("Prueba de IterarRango sin límites (equivalente a Iterar)")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)

	for i := 0; i < 5; i++ {
		abb.Guardar(i, fmt.Sprintf("%d", i))
	}

	var claves []int
	abb.IterarRango(nil, nil, func(clave int, _ string) bool {
		claves = append(claves, clave)
		return true
	})

	require.Equal(t, []int{0, 1, 2, 3, 4}, claves)
}

func TestABBIteradorRango(t *testing.T) {
	t.Log("Prueba de IteradorRango con diferentes rangos")
	abb := TDADiccionario.CrearABB[string, int](comparacionStrings)

	abb.Guardar("d", 4)
	abb.Guardar("b", 2)
	abb.Guardar("f", 6)
	abb.Guardar("a", 1)
	abb.Guardar("c", 3)
	abb.Guardar("e", 5)
	abb.Guardar("g", 7)

	// Test con rango ["b", "e"]
	desde := "b"
	hasta := "e"
	iter := abb.IteradorRango(&desde, &hasta)

	require.True(t, iter.HaySiguiente())
	clave1, valor1 := iter.VerActual()
	require.EqualValues(t, "b", clave1)
	require.EqualValues(t, 2, valor1)

	iter.Siguiente()
	clave2, valor2 := iter.VerActual()
	require.EqualValues(t, "c", clave2)
	require.EqualValues(t, 3, valor2)

	iter.Siguiente()
	clave3, valor3 := iter.VerActual()
	require.EqualValues(t, "d", clave3)
	require.EqualValues(t, 4, valor3)

	iter.Siguiente()
	clave4, valor4 := iter.VerActual()
	require.EqualValues(t, "e", clave4)
	require.EqualValues(t, 5, valor4)

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestABBIteradorRangoVacio(t *testing.T) {
	t.Log("Prueba de IteradorRango con rango vacío")
	abb := TDADiccionario.CrearABB[int, string](comparacionInts)

	for i := 0; i < 10; i++ {
		abb.Guardar(i, fmt.Sprintf("%d", i))
	}

	// Rango imposible: desde > hasta
	desde := 10
	hasta := 5
	iter := abb.IteradorRango(&desde, &hasta)

	require.False(t, iter.HaySiguiente())
}

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	abb := TDADiccionario.CrearABB[string, int](comparacionStrings)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		abb.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = abb.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = abb.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = abb.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !abb.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, abb.Cantidad())
}

func BenchmarkABB(b *testing.B) {
	b.Log("Prueba de stress del ABB. Prueba guardando distinta cantidad de elementos, " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves generadas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenABB(b, n)
			}
		})
	}
}

func TestABBIterarVolumen(t *testing.T) {
	t.Log("Prueba de volumen de iterador, para validar que siempre itere en orden")

	abb := TDADiccionario.CrearABB[int, int](comparacionInts)

	// Inserto elementos en orden aleatorio
	elementos := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		elementos[i] = i
	}
	rand.Shuffle(len(elementos), func(i, j int) {
		elementos[i], elementos[j] = elementos[j], elementos[i]
	})

	for _, elem := range elementos {
		abb.Guardar(elem, elem)
	}

	// Verifico que el iterador los devuelva en orden
	anterior := -1
	ordenCorrecto := true
	abb.Iterar(func(clave int, _ int) bool {
		if clave <= anterior {
			ordenCorrecto = false
			return false
		}
		anterior = clave
		return true
	})

	require.True(t, ordenCorrecto, "El iterador no devolvió los elementos en orden")
}

func TestABBString(t *testing.T) {
	t.Log("Valida que el ABB funcione con strings y los ordene correctamente")
	abb := TDADiccionario.CrearABB[string, string](comparacionStrings)

	abb.Guardar("perro", "guau")
	abb.Guardar("gato", "miau")
	abb.Guardar("vaca", "muu")
	abb.Guardar("cerdo", "oink")
	abb.Guardar("pollo", "pio")

	// Verificar orden alfabético
	claves := []string{}
	abb.Iterar(func(clave string, _ string) bool {
		claves = append(claves, clave)
		return true
	})

	require.Equal(t, []string{"cerdo", "gato", "perro", "pollo", "vaca"}, claves)
}

func TestABBIteradorRangoString(t *testing.T) {
	t.Log("Valida el IteradorRango con strings")
	abb := TDADiccionario.CrearABB[string, int](comparacionStrings)

	palabras := []string{"casa", "barco", "dado", "elefante", "auto", "foca"}
	for i, palabra := range palabras {
		abb.Guardar(palabra, i)
	}

	desde := "barco"
	hasta := "elefante"
	iter := abb.IteradorRango(&desde, &hasta)

	var resultado []string
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
		iter.Siguiente()
	}

	require.Equal(t, []string{"barco", "casa", "dado", "elefante"}, resultado)
}

func TestABBClavesStructs(t *testing.T) {
	t.Log("Valida que también funcione con estructuras más complejas")
	type persona struct {
		nombre string
		edad   int
	}

	compararPersonas := func(a, b persona) int {
		if a.nombre != b.nombre {
			return strings.Compare(a.nombre, b.nombre)
		}
		return a.edad - b.edad
	}

	abb := TDADiccionario.CrearABB[persona, string](compararPersonas)

	p1 := persona{"Ana", 25}
	p2 := persona{"Carlos", 30}
	p3 := persona{"Ana", 20}
	p4 := persona{"Beatriz", 35}

	abb.Guardar(p1, "persona1")
	abb.Guardar(p2, "persona2")
	abb.Guardar(p3, "persona3")
	abb.Guardar(p4, "persona4")

	require.EqualValues(t, 4, abb.Cantidad())
	require.True(t, abb.Pertenece(p1))
	require.True(t, abb.Pertenece(p2))
	require.True(t, abb.Pertenece(p3))
	require.True(t, abb.Pertenece(p4))

	// Verificar orden
	var personas []persona
	abb.Iterar(func(clave persona, _ string) bool {
		personas = append(personas, clave)
		return true
	})

	// Deben estar ordenados por nombre, y si el nombre es igual, por edad
	require.Equal(t, p3, personas[0]) // Ana, 20
	require.Equal(t, p1, personas[1]) // Ana, 25
	require.Equal(t, p4, personas[2]) // Beatriz, 35
	require.Equal(t, p2, personas[3]) // Carlos, 30
}
