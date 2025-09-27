package recu1

// Se tiene una cadena que contiene () y ningún otro caracter (considerar que un único caracter es de tipo rune). Un
// ejercicio típico es dada una cadena averiguar si está balanceada (es decir, todos los símbolos de apertura se cierran, y
// además respetan el orden en el que se abrieron. Ejemplos balanceados: "()()()", o "(())()". No balanceados: "(()",
// o ")(".
// Teniendo en cuenta esto, se tiene una cadena que se asegura que en caso de borrar algunos paréntesis la cadena será
// balanceada, se pide implementar una función func cantBorradosBalanceada(cadena string) int que dada una
// cadena de este tipo, devuelva la cantidad mínima de paréntesis que se deben borrar para que la cadena esté balanceada.
// Indicar y justificar la complejidad del algoritmo.
// Ejemplos:
// cadena: '()' -> 0
// cadena: ')(' -> 2
// cadena: '(()' -> 1
// cadena: ')(()' -> 2

type T any

type pila[T any] struct {
	cantidad int
}

func (p *pila[T]) Apilar(elemento T) {

}

func (p pila[T]) EstaVacia() bool {
	return false
}

func cantBorradosBalanceada(cadena string) int {

	pila := pila[T]{
		cantidad: 0,
	}

	borrar := 0

	for c := range cadena {

		if c == '(' {
			pila.Apilar(c)
		} else {
			if pila.EstaVacia() {
				borrar++
			} else {
				pila.Desapilar()
			}
		}

	}

	return borrar + pila.cantidad
}

// 2)
// Implementar una función func esCuadradoPerfecto(n int) bool que por División y Conquista determine si el
// número n (un positivo entero) es un cuadrado perfecto. Un número es cuadrado perfecto si existe un número entero x
// tal que x2 = n. Indicar y justificar la complejidad del algoritmo utilizando el Teorema Maestro.

func esCuadradoPerfecto(n int) bool {
	return esCuadradoPerfectoRecursivo(n, n)

}

func esCuadradoPerfectoRecursivo(n int, mid int) bool {
	if mid*mid == n {
		return true
	}

	mid = mid / 2

	if mid*mid == n {
		return true
	} else if mid*mid > n {
		return esCuadradoPerfectoRecursivo(n, mid)
	} else {
		return false
	}
}
