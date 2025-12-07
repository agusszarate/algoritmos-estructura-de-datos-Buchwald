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
		ag.usuarios.Guardar(u.ObtenerNombre(), u)
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
	return ADIOS
}

func (ag *AlgoGram) PublicarPost(mensaje string) string {
	if ag.usuarioLoggeado == nil {
		return ERR_USUARIO_NO_LOGGEADO
	}

	nuevoPost := crearPost(len(ag.posts), ag.usuarioLoggeado.ObtenerNombre(), mensaje)
	ag.posts = append(ag.posts, nuevoPost)

	posPublicador := ag.usuarioLoggeado.ObtenerPosicion()

	for _, usuario := range ag.listaUsuarios {
		if usuario.ObtenerNombre() != ag.usuarioLoggeado.ObtenerNombre() {
			usuario.RecibirPost(nuevoPost.ObtenerId(), posPublicador)
		}
	}
	return POST_PUBLICADO
}

func (ag *AlgoGram) VerSiguienteFeed() string {
	if ag.usuarioLoggeado == nil {
		return ERR_USUARIO_NO_LOGGEADO_O_POSTS_VISTOS
	}

	idPost := ag.usuarioLoggeado.VerSiguientePost()
	if idPost == -1 {
		return ERR_USUARIO_NO_LOGGEADO_O_POSTS_VISTOS
	}

	post := ag.posts[idPost]

	return fmt.Sprintf("Post ID %d\n%s dijo: %s\nLikes: %d",
		post.ObtenerId(),
		post.ObtenerAutor(),
		post.ObtenerMensaje(),
		post.ObtenerCantidadLikes())
}

func (ag *AlgoGram) LikearPost(idStr string) string {
	id, err := strconv.Atoi(idStr)
	if err != nil || ag.usuarioLoggeado == nil || id < 0 || id >= len(ag.posts) {
		return ERR_USUARIO_NO_LOGGEADO_O_POST_INEXISTENTE
	}

	ag.posts[id].Likear(ag.usuarioLoggeado.ObtenerNombre())
	return POST_LIKEADO
}

func (ag *AlgoGram) MostrarLikes(idStr string) string {
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(ag.posts) {
		return ERR_POST_INEXISTENTE_O_SIN_LIKES
	}

	return ag.posts[id].ObtenerLikesFormato()
}
