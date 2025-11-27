package algogram

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
)

type Post struct {
	Id      int
	Autor   string
	Mensaje string
	likes   TDADiccionario.DiccionarioOrdenado[string, bool]
}

func crearPost(id int, autor, mensaje string) *Post {
	return &Post{
		Id:      id,
		Autor:   autor,
		Mensaje: mensaje,
		likes:   TDADiccionario.CrearABB[string, bool](cmpString),
	}
}

func (post *Post) likear(usuario string) {
	post.likes.Guardar(usuario, true)
}

func (post *Post) mostrarLikes() {
	if post.likes.Cantidad() == 0 {
		fmt.Println(ERR_POST_INEXISTENTE_O_SIN_LIKES)
		return
	}
	fmt.Printf("El post tiene %d likes:\n", post.likes.Cantidad())
	post.likes.Iterar(func(nombre string, _ bool) bool {
		fmt.Printf("\t%s\n", nombre)
		return true
	})
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
