package parcialito

// 1)
// Sincero Piladibujo, un famoso corredor de autos, se está yendo a la costa a disfrutar sus vacaciones de Semana Santa.
// Resulta que viene manejando por la ruta y tiene una fila de autos por delante, bastante extensa, que quiere sobrepasar.
// Al ser una ruta doble mano, puede ser bastante peligroso intentar sobrepasar toda la fila y que venga un auto de
// frente, por lo cual nos pidió ayuda. Se sabe que la fila se representa como una Cola[Auto] y que se tiene una función
// ObtenerTiempoDeSobrepaso(auto Auto) int, que se ejecuta en O(1) y que brinda el tiempo que tardará Sincero en
// sobrepasar a un auto en particular.
// Implementar una función PuedeSobrepasar(filaAutos Cola[Auto], tiempoMaximoDeManiobra int) bool siendo
// tiempoMaximoDeManiobra el tiempo en el que, por el paso del auto de la mano contraria, genera que la maniobra sea
// imposible y demasiado peligrosa. Al finalizar la ejecución de la función, la cola debe quedar en el estado original que
// tenía antes de ser ejecutada. Indicar y justificar la complejidad del algoritmo.

type Auto any

type Cola[Auto any] struct {
	nodos    any
	cantidad int
}

func (c *Cola[Auto]) Desencolar() Auto {
	// Implementation needed - this is a placeholder
	var zero Auto
	if c.cantidad > 0 {
		c.cantidad--
	}
	return zero
}

func (c *Cola[Auto]) Encolar(auto Auto) {
	c.cantidad++
}

func (c Cola[Auto]) EstaVacia() bool {
	return false
}

func ObtenerTiempoSobrepaso(auto Auto) int {
	return 1
}

func PuedeSobrepasar(filaAutos Cola[Auto], tiempoMaximoDeManiobra int) bool {

	colaAux := Cola[Auto]{
		cantidad: 0,
	}

	for filaAutos.EstaVacia() { //O(n)
		auto := filaAutos.Desencolar() // O(1)

		tiempo := ObtenerTiempoSobrepaso(auto) // O(1)
		tiempoMaximoDeManiobra -= tiempo       // O(1)

		colaAux.Encolar(auto) // O(1)

	}

	for colaAux.EstaVacia() {
		filaAutos.Encolar(colaAux.Desencolar())
	}

	if tiempoMaximoDeManiobra > 0 { // O(1)
		return true // O(1)
	}

	return false // O(1)
}

//complejidad del algoritmo es:
// O(n) + O(1) ...
//que es igual a
// O(n + 1 ...)
//Las constantes son despreciables con N porque siempre vamos a poder acotar por algo K * C
//complejidad queda O(n) con N igual a la cantidad de autos que tenga la fila

// 2)
// Se tiene un arreglo ordenado ascendentemente el cual ha sufrido k rotaciones (el cual es desconocido)
// y se quiere hallar el menor elemento del mismo.
// Implementar una función hallarMenor(array []int) int que lo retorne, utilizando División y Conquista.
// ¿Cuál es la complejidad del algoritmo? Justificar utilizando el Teorema Maestro.

func HallarMenor(array []int) int {

	if len(array) == 1 { // O(1)
		return array[0] // O(1)
	}

	mid := len(array) / 2 // O(1)

	if array[0] <= array[mid] { // O(1)
		return HallarMenor(array[:mid]) // O(n / 2)
	} else {
		return HallarMenor(array[mid:]) // O(n / 2)
	}

}

// complejidad:
// T[n] = T(n/2) + O(1)
// A = 1
// B = 2
// C = 0
// Log2(1) = 0
// T(n) = O((n ^ 0) * log2(2 * n))
// T(n) = O(log(n))
// se cambia de base el logaritmo con constantes
// Luego esas constantes pueden ser despreciadas
