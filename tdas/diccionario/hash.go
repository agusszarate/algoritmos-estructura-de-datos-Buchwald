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
	_FACTOR_CARGA_ACHICAR  = 0.3
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

func (h *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHashCerrado[K, V]{hash: h, pos: -1}
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

func CrearHash[K any, V any](igualdad func(K, K) bool) Diccionario[K, V] {
	return &hashCerrado[K, V]{
		tabla:    make([]celdaHash[K, V], _TAMAÑO_HASH),
		cantidad: 0,
		tam:      _TAMAÑO_HASH,
		borrados: 0,
		igualdad: igualdad,
	}
}

func (h *hashCerrado[K, V]) panicNoPertenece() {
	panic("La clave no pertenece al diccionario")
}

// usamos fnv hashing como funcion de hash
func (h *hashCerrado[K, V]) funcionHash(clave K) int {
	const _FNVM uint64 = 1099511628211
	const _FNVI uint64 = 14695981039346656037
	bytes := convertirABytes(clave)
	hash := _FNVI

	for _, b := range bytes {
		hash *= _FNVM
		hash ^= uint64(b)
	}

	return int(hash % uint64(h.tam))
}

func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h *hashCerrado[K, V]) calcularFactorCarga() float64 {
	return float64(h.cantidad+h.borrados) / float64(h.tam)
}

func (h *hashCerrado[K, V]) debeAgrandar() bool {
	return h.calcularFactorCarga() >= _FACTOR_CARGA_AGRANDAR
}

func (h *hashCerrado[K, V]) debeAchicar() bool {
	factorCarga := float64(h.cantidad) / float64(h.tam)
	return h.tam > _TAMAÑO_HASH && factorCarga <= _FACTOR_CARGA_ACHICAR
}

func (h *hashCerrado[K, V]) buscarPosicion(clave K) (int, bool) {
	pos := h.funcionHash(clave)
	inicio := pos
	primerBorrado := -1

	for {
		celda := &h.tabla[pos]

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
		} else if celda.estado == _OCUPADO && h.igualdad(celda.clave, clave) {
			return pos, true
		}

		pos = (pos + 1) % h.tam
		if pos == inicio {
			if primerBorrado != -1 {
				return primerBorrado, false
			}
			break
		}
	}
	return -1, false
}

func (h *hashCerrado[K, V]) redimensionar(nuevoTam int) {
	tablaVieja := h.tabla
	h.tabla = make([]celdaHash[K, V], nuevoTam)
	h.tam = nuevoTam
	cantidadVieja := h.cantidad
	h.cantidad = 0
	h.borrados = 0

	for i := range len(tablaVieja) {
		if tablaVieja[i].estado == _OCUPADO {
			h.Guardar(tablaVieja[i].clave, tablaVieja[i].dato)
		}
	}
	h.cantidad = cantidadVieja
}

func (h *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if h.debeAgrandar() {
		h.redimensionar(h.tam * _FACTOR_REDIMENSION)
	}

	pos, existe := h.buscarPosicion(clave)
	if existe {
		h.tabla[pos].dato = dato
		return
	}

	if h.tabla[pos].estado == _BORRADO {
		h.borrados--
	}
	h.tabla[pos] = celdaHash[K, V]{
		clave:  clave,
		dato:   dato,
		estado: _OCUPADO,
	}
	h.cantidad++
}

func (h *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, existe := h.buscarPosicion(clave)
	return existe
}

func (h *hashCerrado[K, V]) Obtener(clave K) V {
	pos, existe := h.buscarPosicion(clave)
	if !existe {
		h.panicNoPertenece()
	}
	return h.tabla[pos].dato
}

func (h *hashCerrado[K, V]) Borrar(clave K) V {
	pos, existe := h.buscarPosicion(clave)
	if !existe {
		h.panicNoPertenece()
	}

	dato := h.tabla[pos].dato
	h.tabla[pos].estado = _BORRADO
	h.cantidad--
	h.borrados++

	if h.debeAchicar() {
		h.redimensionar(h.tam / _FACTOR_REDIMENSION)
	}

	return dato
}

func (h *hashCerrado[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := range len(h.tabla) {
		if h.tabla[i].estado == _OCUPADO {
			if !visitar(h.tabla[i].clave, h.tabla[i].dato) {
				return
			}
		}
	}
}
