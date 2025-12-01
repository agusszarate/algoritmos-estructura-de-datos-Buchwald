package algogram

import (
	"fmt"
	"strconv"
	TDADiccionario "tdas/diccionario"
)

type AlgoGram struct {
	usuarios        TDADiccionario.Diccionario[string, *Usuario]
	posts           []*Post
	usuarioLoggeado *Usuario
	listaUsuarios   []*Usuario
}

func CrearAlgoGram(listaUsuarios []*Usuario) *AlgoGram {
	ag := &AlgoGram{
		usuarios:      TDADiccionario.CrearHash[string, *Usuario](func(a, b string) bool { return a == b }),
		posts:         make([]*Post, 0),
		listaUsuarios: listaUsuarios,
	}
	for _, u := range listaUsuarios {
		ag.usuarios.Guardar(u.Nombre, u)
	}
	return ag
}

func (ag *AlgoGram) Login(nombre string) string {
	if ag.usuarioLoggeado != nil {
		return ERR_USUARIO_YA_LOGGEADO
	}
	if !ag.usuarios.Pertenece(nombre) {
		return ERR_USUARIO_NO_EXISTENTE
	}
	ag.usuarioLoggeado = ag.usuarios.Obtener(nombre)
	return "Hola " + nombre
}

func (ag *AlgoGram) Logout() string {
	if ag.usuarioLoggeado == nil {
		return ERR_USUARIO_NO_LOGGEADO
	}
	ag.usuarioLoggeado = nil
	return "Adios"
}

func (ag *AlgoGram) PublicarPost(mensaje string) string {
	if ag.usuarioLoggeado == nil {
		return ERR_USUARIO_NO_LOGGEADO
	}
	nuevoPost := crearPost(len(ag.posts), ag.usuarioLoggeado.Nombre, mensaje)
	ag.posts = append(ag.posts, nuevoPost)

	for _, usuario := range ag.listaUsuarios {
		if usuario.Nombre != ag.usuarioLoggeado.Nombre {
			usuario.recibirPost(nuevoPost.Id, ag.usuarioLoggeado.afinidad)
		}
	}
	return "Post publicado"
}

func (ag *AlgoGram) VerSiguienteFeed() string {
	if ag.usuarioLoggeado == nil {
		return ERR_USUARIO_NO_LOGGEADO_O_POSTS_VISTOS
	}

	idPost := ag.usuarioLoggeado.verSiguientePost()
	if idPost == -1 {
		return ERR_USUARIO_NO_LOGGEADO_O_POSTS_VISTOS
	}

	post := ag.posts[idPost]
	return fmt.Sprintf("Post ID %d\n%s dijo: %s\nLikes: %d", post.Id, post.Autor, post.Mensaje, post.likes.Cantidad())
}

func (ag *AlgoGram) LikearPost(idStr string) string {
	id, err := strconv.Atoi(idStr)
	if err != nil || ag.usuarioLoggeado == nil || id < 0 || id >= len(ag.posts) {
		return ERR_USUARIO_NO_LOGGEADO_O_POST_INEXISTENTE
	}
	ag.posts[id].likear(ag.usuarioLoggeado.Nombre)
	return "Post likeado"
}

func (ag *AlgoGram) MostrarLikes(idStr string) string {
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(ag.posts) {
		return ERR_POST_INEXISTENTE_O_SIN_LIKES
	}
	return ag.posts[id].obtenerLikes()
}
