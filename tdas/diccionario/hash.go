package diccionario

type estadoCelda int

const (
	vacio estadoCelda = iota
	ocupado
	borrado
	_TAMAÑO_HASH = 11
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoCelda
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

type iterHashCerrado[K comparable, V any] struct {
	hash *hashCerrado[K, V]
	pos  int
}

//-----------------------------------Iterador-----------------------------------//

func (iter *iterHashCerrado[K, V]) panicFin() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (h *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHashCerrado[K, V]{hash: h, pos: -1}
	iter.avanzar()
	return iter
}

func (iter *iterHashCerrado[K, V]) avanzar() {
	iter.pos++
	for iter.pos < len(iter.hash.tabla) {
		if iter.hash.tabla[iter.pos].estado == ocupado {
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

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &hashCerrado[K, V]{
		tabla: make([]celdaHash[K, V], _TAMAÑO_HASH),
		tam:   _TAMAÑO_HASH,
	}
	return hash
}
