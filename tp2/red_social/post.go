package algogram

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
)

type Post struct {
	id      int
	autor   string
	mensaje string
	likes   TDADiccionario.DiccionarioOrdenado[string, bool]
}

func crearPost(id int, autor, mensaje string) *Post {
	return &Post{
		id:      id,
		autor:   autor,
		mensaje: mensaje,
		likes:   TDADiccionario.CrearABB[string, bool](cmpString),
	}
}

func (post *Post) ObtenerId() int {
	return post.id
}

func (post *Post) ObtenerAutor() string {
	return post.autor
}

func (post *Post) ObtenerMensaje() string {
	return post.mensaje
}

func (post *Post) ObtenerCantidadLikes() int {
	return post.likes.Cantidad()
}

func (post *Post) Likear(usuario string) {
	post.likes.Guardar(usuario, true)
}

func (post *Post) ObtenerLikesFormato() string {
	if post.likes.Cantidad() == 0 {
		return ERR_POST_INEXISTENTE_O_SIN_LIKES
	}
	resultado := fmt.Sprintf("El post tiene %d likes:", post.likes.Cantidad())
	post.likes.Iterar(func(nombre string, _ bool) bool {
		resultado += "\n\t" + nombre
		return true
	})
	return resultado
}

func cmpString(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
