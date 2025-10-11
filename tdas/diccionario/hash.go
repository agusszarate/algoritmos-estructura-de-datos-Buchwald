package diccionario

import "fmt"

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

func CrearHash[K any, V any](igualdad func(K, K) bool) Diccionario[K, V] {
	return crearTabla[K, V](_TAMAÑO_HASH, igualdad)
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
