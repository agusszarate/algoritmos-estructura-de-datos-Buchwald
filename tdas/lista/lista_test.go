package lista_test

import (
	TDALista "tdas/lista"
	"testing"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	if !lista.EstaVacia() {
		t.Error("Lista nueva debe estar vacía")
	}

	if lista.Largo() != 0 {
		t.Error("Lista vacía debe tener largo 0")
	}
}

func TestInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

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
	lista := TDALista.CrearListaEnlazada[int]()

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
	lista := TDALista.CrearListaEnlazada[int]()

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
	lista := TDALista.CrearListaEnlazada[int]()

	defer func() {
		if r := recover(); r == nil {
			t.Error("VerPrimero debe entrar en pánico con lista vacía")
		}
	}()

	lista.VerPrimero()
}

func TestPanicoBorrarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	defer func() {
		if r := recover(); r == nil {
			t.Error("BorrarPrimero debe entrar en pánico con lista vacía")
		}
	}()

	lista.BorrarPrimero()
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
		if iter.VerActual() != valor {
			t.Errorf("Valor actual debe ser %d, pero es %d", valor, iter.VerActual())
		}
		iter.Siguiente()
		contador++
		valor++
	}

	if contador != 3 {
		t.Error("Debe iterar 3 elementos")
	}
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

	if contador != 3 {
		t.Error("Debe iterar 3 elementos")
	}
	if suma != 6 {
		t.Error("Suma debe ser 6")
	}
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

	if contador != 2 {
		t.Error("Debe iterar 2 elementos antes de cortar")
	}
}

func TestIteradorInsertar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(40)

	iter := lista.Iterador()
	iter.Insertar(10)
	if lista.VerPrimero() != 10 {
		t.Errorf("Primero debe ser 10 pero es %d", lista.VerPrimero())
	}

	iter = lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	iter.Insertar(30)

	iter = lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(50)
	if lista.VerUltimo() != 50 {
		t.Errorf("Ultimo debe ser 50 pero es %d", lista.VerUltimo())
	}

	arr := []int{10, 20, 30, 40, 50}
	valores := []int{}
	lista.Iterar(func(v int) bool {
		valores = append(valores, v)
		return true
	})
	for i := range arr {
		if valores[i] != arr[i] {
			t.Errorf("Se esperaba %d pero se obtuvo %d", arr[i], valores[i])
		}
	}
}

func TestIteradorBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for _, v := range []int{10, 20, 30, 40} {
		lista.InsertarUltimo(v)
	}

	iter := lista.Iterador()
	if iter.Borrar() != 10 {
		t.Error("Se esperaba borrar 10")
	}

	iter = lista.Iterador()
	if iter.Borrar() != 20 {
		t.Error("Se esperaba borrar 20")
	}
	if iter.Borrar() != 30 {
		t.Error("Se esperaba borrar 30")
	}

	iter = lista.Iterador()
	if iter.Borrar() != 40 {
		t.Error("Se esperaba borrar 40")
	}
	if !lista.EstaVacia() {
		t.Error("Lista debe quedar vacía")
	}

}

func TestIteradorCasosBorde(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	if iter.HaySiguiente() {
		t.Error("En lista vacía no debería haber siguiente")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("Se esperaba pánico al borrar en lista vacía")
		}
	}()

	iter.Borrar()

	lista.InsertarUltimo(10)

	iter = lista.Iterador()

	if iter.Siguiente() != 10 {
		t.Errorf("Se esperaba 10 como primer elemento")
	}

	iter.Borrar()

	if !lista.EstaVacia() {
		t.Error("Lista debe quedar vacía")
	}

	iter = lista.Iterador()

	if iter.HaySiguiente() {
		t.Error("Iterador no debe tener siguiente en lista vacía")
	}
}
