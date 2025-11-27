package algogram

import (
	TDAColaPrioridad "tdas/cola_prioridad"
)

type Usuario struct {
	Nombre   string
	afinidad int
	feed     TDAColaPrioridad.ColaPrioridad[itemFeed]
}

type itemFeed struct {
	afinidad int
	idPost   int
}

func CrearUsuario(nombre string, afinidad int) *Usuario {
	return &Usuario{
		Nombre:   nombre,
		afinidad: afinidad,
		feed:     TDAColaPrioridad.CrearHeap(cmpFeed),
	}
}

func (usuario *Usuario) recibirPost(idPost int, afinidadAutor int) {
	afinidad := abs(usuario.afinidad - afinidadAutor)
	usuario.feed.Encolar(itemFeed{afinidad: afinidad, idPost: idPost})
}

func (usuario *Usuario) verSiguientePost() int {
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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
