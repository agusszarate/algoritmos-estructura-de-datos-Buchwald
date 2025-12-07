package algogram

import (
	TDAColaPrioridad "tdas/cola_prioridad"
)

type Usuario struct {
	nombre   string
	afinidad int
	feed     TDAColaPrioridad.ColaPrioridad[itemFeed]
}

type itemFeed struct {
	afinidad int
	idPost   int
}

func CrearUsuario(nombre string, afinidad int) *Usuario {
	return &Usuario{
		nombre:   nombre,
		afinidad: afinidad,
		feed:     TDAColaPrioridad.CrearHeap(cmpFeed),
	}
}

func (usuario *Usuario) ObtenerNombre() string {
	return usuario.nombre
}

func (usuario *Usuario) ObtenerPosicion() int {
	return usuario.afinidad
}

func (usuario *Usuario) RecibirPost(idPost int, posicionAutor int) {
	distancia := abs_usuario(usuario.afinidad - posicionAutor)
	usuario.feed.Encolar(itemFeed{afinidad: distancia, idPost: idPost})
}

func (usuario *Usuario) VerSiguientePost() int {
	if usuario.feed.EstaVacia() {
		return -1
	}
	return usuario.feed.Desencolar().idPost
}

func cmpFeed(a, b itemFeed) int {
	if a.afinidad != b.afinidad {
		return b.afinidad - a.afinidad
	}
	return b.idPost - a.idPost
}

func abs_usuario(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
