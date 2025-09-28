package lista

import (
	"testing"
)

func TestListaVacia(t *testing.T) {
	lista := CrearListaEnlazada[int]()

	if !lista.EstaVacia() {
		t.Error("Lista nueva debe estar vacía")
	}

	if lista.Largo() != 0 {
		t.Error("Lista vacía debe tener largo 0")
	}
}

func TestInsertarPrimero(t *testing.T) {
	lista := CrearListaEnlazada[int]()

	lista.InsertarPrimero(10)
	if lista.EstaVacia() {
		t.Error("Lista no debe estar vacía después de insertar")
	}
	if lista.Largo() != 1 {
		t.Error("Lista debe tener largo 1")
	}
	if lista.VerPrimero() != 10 {
		t.Error("Primer elemento debe ser 10")
	}
	if lista.VerUltimo() != 10 {
		t.Error("Último elemento debe ser 10")
	}

	lista.InsertarPrimero(20)
	if lista.Largo() != 2 {
		t.Error("Lista debe tener largo 2")
	}
	if lista.VerPrimero() != 20 {
		t.Error("Primer elemento debe ser 20")
	}
	if lista.VerUltimo() != 10 {
		t.Error("Último elemento debe ser 10")
	}
}

func TestInsertarUltimo(t *testing.T) {
	lista := CrearListaEnlazada[int]()

	lista.InsertarUltimo(10)
	if lista.EstaVacia() {
		t.Error("Lista no debe estar vacía después de insertar")
	}
	if lista.Largo() != 1 {
		t.Error("Lista debe tener largo 1")
	}
	if lista.VerPrimero() != 10 {
		t.Error("Primer elemento debe ser 10")
	}
	if lista.VerUltimo() != 10 {
		t.Error("Último elemento debe ser 10")
	}

	lista.InsertarUltimo(20)
	if lista.Largo() != 2 {
		t.Error("Lista debe tener largo 2")
	}
	if lista.VerPrimero() != 10 {
		t.Error("Primer elemento debe ser 10")
	}
	if lista.VerUltimo() != 20 {
		t.Error("Último elemento debe ser 20")
	}
}

func TestBorrarPrimero(t *testing.T) {
	lista := CrearListaEnlazada[int]()

	lista.InsertarPrimero(10)
	lista.InsertarPrimero(20)
	lista.InsertarPrimero(30)

	elemento := lista.BorrarPrimero()
	if elemento != 30 {
		t.Error("Elemento borrado debe ser 30")
	}
	if lista.Largo() != 2 {
		t.Error("Lista debe tener largo 2")
	}
	if lista.VerPrimero() != 20 {
		t.Error("Primer elemento debe ser 20")
	}

	lista.BorrarPrimero()
	lista.BorrarPrimero()

	if !lista.EstaVacia() {
		t.Error("Lista debe estar vacía")
	}
}

func TestPanicoListaVacia(t *testing.T) {
	lista := CrearListaEnlazada[int]()

	defer func() {
		if r := recover(); r == nil {
			t.Error("VerPrimero debe entrar en pánico con lista vacía")
		}
	}()

	lista.VerPrimero()
}

func TestPanicoBorrarListaVacia(t *testing.T) {
	lista := CrearListaEnlazada[int]()

	defer func() {
		if r := recover(); r == nil {
			t.Error("BorrarPrimero debe entrar en pánico con lista vacía")
		}
	}()

	lista.BorrarPrimero()
}

// func TestIterador(t *testing.T) {
// 	lista := CrearListaEnlazada[int]()
// 	lista.InsertarUltimo(1)
// 	lista.InsertarUltimo(2)
// 	lista.InsertarUltimo(3)

// 	iter := lista.Iterador()
// 	contador := 0
// 	valor := 1

// 	for iter.HaySiguiente() {
// 		if iter.VerActual() != valor {
// 			t.Errorf("Valor actual debe ser %d, pero es %d", valor, iter.VerActual())
// 		}
// 		iter.Siguiente()
// 		contador++
// 		valor++
// 	}

// 	if contador != 3 {
// 		t.Error("Debe iterar 3 elementos")
// 	}
// }

// func TestIterar(t *testing.T) {
// 	lista := CrearListaEnlazada[int]()
// 	lista.InsertarUltimo(1)
// 	lista.InsertarUltimo(2)
// 	lista.InsertarUltimo(3)

// 	contador := 0
// 	suma := 0

// 	lista.Iterar(func(elemento int) bool {
// 		contador++
// 		suma += elemento
// 		return true
// 	})

// 	if contador != 3 {
// 		t.Error("Debe iterar 3 elementos")
// 	}
// 	if suma != 6 {
// 		t.Error("Suma debe ser 6")
// 	}
// }

// func TestIterarCorte(t *testing.T) {
// 	lista := CrearListaEnlazada[int]()
// 	lista.InsertarUltimo(1)
// 	lista.InsertarUltimo(2)
// 	lista.InsertarUltimo(3)

// 	contador := 0

// 	lista.Iterar(func(elemento int) bool {
// 		contador++
// 		return elemento < 2
// 	})

// 	if contador != 2 {
// 		t.Error("Debe iterar 2 elementos antes de cortar")
// 	}
// }
