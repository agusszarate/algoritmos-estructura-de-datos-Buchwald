package main

import (
	"fmt"
	. "tdas/cola_prioridad"
	. "tdas/diccionario"
	. "tdas/lista"
)

type estadoCelda int

const (
	_VACIO estadoCelda = iota
	_OCUPADO
	_BORRADO
	_TAMAÑO_HASH           = 11
	_FACTOR_REDIMENSION    = 2
	_FACTOR_CARGA_AGRANDAR = 0.7
	_FACTOR_CARGA_ACHICAR  = 0.2
)

type celdaHash[K any, V any] struct {
	clave  K
	dato   V
	estado estadoCelda
}

type hashCerrado[K any, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
	igualdad func(K, K) bool
}

type iterHashCerrado[K any, V any] struct {
	hash *hashCerrado[K, V]
	pos  int
}

//-----------------------------------Iterador-----------------------------------//

func (iter *iterHashCerrado[K, V]) panicFin() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHashCerrado[K, V]{hash: hash, pos: -1}
	iter.avanzar()
	return iter
}

func (iter *iterHashCerrado[K, V]) avanzar() {
	iter.pos++
	for iter.pos < len(iter.hash.tabla) {
		if iter.hash.tabla[iter.pos].estado == _OCUPADO {
			return
		}
		iter.pos++
	}
}

func (iter *iterHashCerrado[K, V]) HaySiguiente() bool {
	return iter.pos < len(iter.hash.tabla)
}

func (iter *iterHashCerrado[K, V]) VerActual() (K, V) {
	iter.panicFin()
	celda := iter.hash.tabla[iter.pos]
	return celda.clave, celda.dato
}

func (iter *iterHashCerrado[K, V]) Siguiente() {
	iter.panicFin()
	iter.avanzar()
}

//-----------------------------------Hash-----------------------------------//

func crearTabla[K any, V any](tam int, igualdad func(K, K) bool) *hashCerrado[K, V] {
	return &hashCerrado[K, V]{
		tabla:    make([]celdaHash[K, V], tam),
		cantidad: 0,
		tam:      tam,
		borrados: 0,
		igualdad: igualdad,
	}
}

func (hash *hashCerrado[K, V]) panicNoPertenece() {
	panic("La clave no pertenece al diccionario")
}

// usamos fnv hashing como funcion de hash
func (hash *hashCerrado[K, V]) funcionHash(clave K) int {
	const _FNVM uint64 = 1099511628211
	const _FNVI uint64 = 14695981039346656037
	bytes := convertirABytes(clave)
	hashValue := _FNVI

	for _, b := range bytes {
		hashValue *= _FNVM
		hashValue ^= uint64(b)
	}

	return int(hashValue % uint64(hash.tam))
}

func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (hash *hashCerrado[K, V]) calcularFactorCarga() float64 {
	return float64(hash.cantidad+hash.borrados) / float64(hash.tam)
}

func (hash *hashCerrado[K, V]) debeAgrandar() bool {
	return hash.calcularFactorCarga() >= _FACTOR_CARGA_AGRANDAR
}

func (hash *hashCerrado[K, V]) debeAchicar() bool {
	factorCarga := float64(hash.cantidad) / float64(hash.tam)
	return hash.tam > _TAMAÑO_HASH && factorCarga <= _FACTOR_CARGA_ACHICAR
}

func (hash *hashCerrado[K, V]) buscarPosicion(clave K) (int, bool) {
	pos := hash.funcionHash(clave)
	inicio := pos
	primerBorrado := -1

	for {
		celda := &hash.tabla[pos]

		if celda.estado == _VACIO {
			if primerBorrado != -1 {
				return primerBorrado, false
			}
			return pos, false
		}

		if celda.estado == _BORRADO {
			if primerBorrado == -1 {
				primerBorrado = pos
			}
		} else if celda.estado == _OCUPADO && hash.igualdad(celda.clave, clave) {
			return pos, true
		}

		pos = (pos + 1) % hash.tam
		if pos == inicio {
			if primerBorrado != -1 {
				return primerBorrado, false
			}
			break
		}
	}
	return -1, false
}

func (hash *hashCerrado[K, V]) redimensionar(nuevoTam int) {
	tablaVieja := hash.tabla
	nuevaTabla := crearTabla[K, V](nuevoTam, hash.igualdad)

	for i := range len(tablaVieja) {
		if tablaVieja[i].estado == _OCUPADO {
			nuevaTabla.Guardar(tablaVieja[i].clave, tablaVieja[i].dato)
		}
	}

	hash.tabla = nuevaTabla.tabla
	hash.tam = nuevaTabla.tam
	hash.borrados = nuevaTabla.borrados
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if hash.debeAgrandar() {
		hash.redimensionar(hash.tam * _FACTOR_REDIMENSION)
	}

	pos, existe := hash.buscarPosicion(clave)
	if existe {
		hash.tabla[pos].dato = dato
		return
	}

	if hash.tabla[pos].estado == _BORRADO {
		hash.borrados--
	}
	hash.tabla[pos] = celdaHash[K, V]{
		clave:  clave,
		dato:   dato,
		estado: _OCUPADO,
	}
	hash.cantidad++
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, existe := hash.buscarPosicion(clave)
	return existe
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	pos, existe := hash.buscarPosicion(clave)
	if !existe {
		hash.panicNoPertenece()
	}
	return hash.tabla[pos].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	pos, existe := hash.buscarPosicion(clave)
	if !existe {
		hash.panicNoPertenece()
	}

	dato := hash.tabla[pos].dato
	hash.tabla[pos].estado = _BORRADO
	hash.cantidad--
	hash.borrados++

	if hash.debeAchicar() {
		hash.redimensionar(hash.tam / _FACTOR_REDIMENSION)
	}

	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := range len(hash.tabla) {
		if hash.tabla[i].estado == _OCUPADO {
			if !visitar(hash.tabla[i].clave, hash.tabla[i].dato) {
				return
			}
		}
	}
}

//-----------------------------------Hash Abierto-----------------------------------//

type parClaveValor[K any, V any] struct {
	clave K
	valor V
}

type hashAbierto[K any, V any] struct {
	tabla    []Lista[parClaveValor[K, V]]
	cantidad int
	tam      int
	igualdad func(K, K) bool
}

type iterHashAbierto[K any, V any] struct {
	hash      *hashAbierto[K, V]
	posTabla  int
	iterLista IteradorLista[parClaveValor[K, V]]
}

//-----------------------------------Iterador Hash Abierto-----------------------------------//

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHashAbierto[K, V]{hash: hash, posTabla: 0}
	iter.avanzarAListaNoVacia()
	return iter
}

func (iter *iterHashAbierto[K, V]) avanzarAListaNoVacia() {
	for iter.posTabla < len(iter.hash.tabla) {
		if iter.hash.tabla[iter.posTabla] != nil && !iter.hash.tabla[iter.posTabla].EstaVacia() {
			iter.iterLista = iter.hash.tabla[iter.posTabla].Iterador()
			return
		}
		iter.posTabla++
	}
}

func (iter *iterHashAbierto[K, V]) HaySiguiente() bool {
	return iter.posTabla < len(iter.hash.tabla) && iter.iterLista != nil && iter.iterLista.HaySiguiente()
}

func (iter *iterHashAbierto[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	par := iter.iterLista.VerActual()
	return par.clave, par.valor
}

func (iter *iterHashAbierto[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.iterLista.Siguiente()

	if !iter.iterLista.HaySiguiente() {
		iter.posTabla++
		iter.avanzarAListaNoVacia()
	}
}

//-----------------------------------Hash Abierto Primitivas-----------------------------------//

func crearTablaAbierta[K any, V any](tam int, igualdad func(K, K) bool) *hashAbierto[K, V] {
	return &hashAbierto[K, V]{
		tabla:    make([]Lista[parClaveValor[K, V]], tam),
		cantidad: 0,
		tam:      tam,
		igualdad: igualdad,
	}
}

func (hash *hashAbierto[K, V]) funcionHash(clave K) int {
	const _FNVM uint64 = 1099511628211
	const _FNVI uint64 = 14695981039346656037
	bytes := []byte(fmt.Sprintf("%v", clave))
	hashValue := _FNVI

	for _, b := range bytes {
		hashValue *= _FNVM
		hashValue ^= uint64(b)
	}

	return int(hashValue % uint64(hash.tam))
}

func (hash *hashAbierto[K, V]) calcularFactorCarga() float64 {
	return float64(hash.cantidad) / float64(hash.tam)
}

func (hash *hashAbierto[K, V]) debeAgrandar() bool {
	return hash.calcularFactorCarga() >= _FACTOR_CARGA_AGRANDAR
}

func (hash *hashAbierto[K, V]) debeAchicar() bool {
	return hash.tam > _TAMAÑO_HASH && hash.calcularFactorCarga() <= _FACTOR_CARGA_ACHICAR
}

func (hash *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	tablaVieja := hash.tabla
	nuevaTabla := crearTablaAbierta[K, V](nuevoTam, hash.igualdad)

	for i := range tablaVieja {
		if tablaVieja[i] != nil {
			iter := tablaVieja[i].Iterador()
			for iter.HaySiguiente() {
				par := iter.VerActual()
				nuevaTabla.Guardar(par.clave, par.valor)
				iter.Siguiente()
			}
		}
	}

	hash.tabla = nuevaTabla.tabla
	hash.tam = nuevaTabla.tam
}

func (hash *hashAbierto[K, V]) buscarEnLista(lista Lista[parClaveValor[K, V]], clave K) (IteradorLista[parClaveValor[K, V]], bool) {
	if lista == nil {
		return nil, false
	}

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		par := iter.VerActual()
		if hash.igualdad(par.clave, clave) {
			return iter, true
		}
		iter.Siguiente()
	}
	return nil, false
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	if hash.debeAgrandar() {
		hash.redimensionar(hash.tam * _FACTOR_REDIMENSION)
	}

	pos := hash.funcionHash(clave)

	if hash.tabla[pos] == nil {
		hash.tabla[pos] = CrearListaEnlazada[parClaveValor[K, V]]()
	}

	iter, existe := hash.buscarEnLista(hash.tabla[pos], clave)
	if existe {
		iter.Borrar()
		hash.tabla[pos].InsertarPrimero(parClaveValor[K, V]{clave, dato})
	} else {
		hash.tabla[pos].InsertarPrimero(parClaveValor[K, V]{clave, dato})
		hash.cantidad++
	}
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	pos := hash.funcionHash(clave)
	if hash.tabla[pos] == nil {
		return false
	}
	_, existe := hash.buscarEnLista(hash.tabla[pos], clave)
	return existe
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	pos := hash.funcionHash(clave)
	if hash.tabla[pos] == nil {
		panic("La clave no pertenece al diccionario")
	}

	iter, existe := hash.buscarEnLista(hash.tabla[pos], clave)
	if !existe {
		panic("La clave no pertenece al diccionario")
	}

	return iter.VerActual().valor
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	pos := hash.funcionHash(clave)
	if hash.tabla[pos] == nil {
		panic("La clave no pertenece al diccionario")
	}

	iter, existe := hash.buscarEnLista(hash.tabla[pos], clave)
	if !existe {
		panic("La clave no pertenece al diccionario")
	}

	valor := iter.VerActual().valor
	iter.Borrar()
	hash.cantidad--

	if hash.debeAchicar() {
		hash.redimensionar(hash.tam / _FACTOR_REDIMENSION)
	}

	return valor
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := range hash.tabla {
		if hash.tabla[i] != nil {
			iter := hash.tabla[i].Iterador()
			for iter.HaySiguiente() {
				par := iter.VerActual()
				if !visitar(par.clave, par.valor) {
					return
				}
				iter.Siguiente()
			}
		}
	}
}

// 1. Implementar una primitiva eliminarColisiones(clave K) []K para el hash, que elimine del hash todas las
// claves que colisionen con la clave pasada por parámetro en el estado actual (sin eliminar dicha clave del
// diccionario, si se encuentra), y devuelva dichas claves. Implementar tanto para el hash abierto como para el hash
// cerrado. Si no se implementa para alguno, el ejercicio no estará aprobable. Indicar y justificar la complejidad de
// la primitiva para ambos casos.

func (hash *hashAbierto[K, V]) eliminarColisiones(clave K) []K {
	pos := hash.funcionHash(clave)
	claves := make([]K, 0)

	if hash.tabla[pos] == nil {
		return claves
	}

	lista := hash.tabla[pos]

	iter := lista.Iterador()

	nuevaLista := CrearListaEnlazada[parClaveValor[K, V]]()

	for iter.HaySiguiente() {

		if hash.igualdad(iter.VerActual().clave, clave) {
			nuevaLista.InsertarUltimo(parClaveValor[K, V]{iter.VerActual().clave, iter.VerActual().valor})
		} else {
			claves = append(claves, iter.VerActual().clave)
			hash.cantidad--
		}

		iter.Siguiente()
	}

	hash.tabla[pos] = nuevaLista

	if hash.debeAchicar() {
		hash.redimensionar(hash.tam / _FACTOR_REDIMENSION)
	}

	return claves

}

func (hash *hashCerrado[K, V]) eliminarColisiones(clave K) []K {
	pos := hash.funcionHash(clave)
	claves := make([]K, 0)

	actual := pos
	estadoCelda := hash.tabla[pos].estado

	for estadoCelda != _VACIO {
		if hash.tabla[actual].estado == _OCUPADO && hash.funcionHash(hash.tabla[actual].clave) == pos {
			if !hash.igualdad(hash.tabla[actual].clave, clave) {

				claves = append(claves, hash.tabla[actual].clave)
				hash.tabla[actual].estado = _BORRADO
				hash.cantidad--
				hash.borrados++
			}
		}

		actual = (actual + 1) % hash.tam
		estadoCelda = hash.tabla[actual].estado
	}

	if hash.debeAchicar() {
		hash.redimensionar(hash.tam / _FACTOR_REDIMENSION)
	}

	return claves
}

// 2. Sobre un AVL cuyo estado inicial puede reconstruirse a partir del preorder: 40 - 10 - 3 - 17 - 15 - 64 -
// 47 - 74 - 92, realizar un seguimiento de insertar los siguientes elementos (incluyendo rotaciones intermedias):
// 20, 23, 13, 14, 16, 12.

// 3. Implementar una función mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno] que, dado un arreglo
// de Alumnos y un valor entero k, nos devuelva una lista de los k alumnos de mayor promedio (ordenada de
// mayor a menor). Indicar y justificar la complejidad del algoritmo implementado.
//
// Considerar que la estructura del alumno es:

type Alumno struct {
	nombre string
	padron int
	notas  []int
}

func comp(a1 Alumno, a2 Alumno) int {
	a1Total := 0
	a2Total := 0

	for _, nota := range a1.notas {
		a1Total += nota
	}

	for _, nota := range a2.notas {
		a2Total += nota
	}

	promedio1 := float64(a1Total) / float64(len(a1.notas))
	promedio2 := float64(a2Total) / float64(len(a2.notas))

	if promedio1 > promedio2 {
		return 1
	} else if promedio1 < promedio2 {
		return -1
	} else {
		return 0
	}
}

func mejoresPromedios(alumnos []Alumno, K int) Lista[Alumno] {
	cola := CrearHeapArr(alumnos, comp)

	var largo int

	if len(alumnos) >= K {
		largo = K
	} else {
		largo = len(alumnos)
	}

	lista := CrearListaEnlazada[Alumno]()

	for i := 0; i < largo; i++ {
		alumno := cola.Desencolar()
		lista.InsertarUltimo(alumno)
	}

	return lista
}
