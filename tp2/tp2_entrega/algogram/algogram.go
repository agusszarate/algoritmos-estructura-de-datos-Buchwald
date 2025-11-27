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

func (ag *AlgoGram) Login(nombre string) {
	if ag.usuarioLoggeado != nil {
		fmt.Println(ERR_USUARIO_YA_LOGGEADO)
		return
	}
	if !ag.usuarios.Pertenece(nombre) {
		fmt.Println(ERR_USUARIO_NO_EXISTENTE)
		return
	}
	ag.usuarioLoggeado = ag.usuarios.Obtener(nombre)
	fmt.Printf("Hola %s\n", nombre)
}

func (ag *AlgoGram) Logout() {
	if ag.usuarioLoggeado == nil {
		fmt.Println(ERR_USUARIO_NO_LOGGEADO)
		return
	}
	ag.usuarioLoggeado = nil
	fmt.Println("Adios")
}

func (ag *AlgoGram) PublicarPost(mensaje string) {
	if ag.usuarioLoggeado == nil {
		fmt.Println(ERR_USUARIO_NO_LOGGEADO)
		return
	}
	nuevoPost := crearPost(len(ag.posts), ag.usuarioLoggeado.Nombre, mensaje)
	ag.posts = append(ag.posts, nuevoPost)

	for _, usuario := range ag.listaUsuarios {
		if usuario.Nombre != ag.usuarioLoggeado.Nombre {
			usuario.recibirPost(nuevoPost.Id, ag.usuarioLoggeado.afinidad)
		}
	}
	fmt.Println("Post publicado")
}

func (ag *AlgoGram) VerSiguienteFeed() {
	if ag.usuarioLoggeado == nil {
		fmt.Println(ERR_USUARIO_NO_LOGGEADO_O_POSTS_VISTOS)
		return
	}

	idPost := ag.usuarioLoggeado.verSiguientePost()
	if idPost == -1 {
		fmt.Println(ERR_USUARIO_NO_LOGGEADO_O_POSTS_VISTOS)
		return
	}

	post := ag.posts[idPost]
	fmt.Printf("Post ID %d\n%s dijo: %s\nLikes: %d\n", post.Id, post.Autor, post.Mensaje, post.likes.Cantidad())
}

func (ag *AlgoGram) LikearPost(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil || ag.usuarioLoggeado == nil || id < 0 || id >= len(ag.posts) {
		fmt.Println(ERR_USUARIO_NO_LOGGEADO_O_POST_INEXISTENTE)
		return
	}
	ag.posts[id].likear(ag.usuarioLoggeado.Nombre)
	fmt.Println("Post likeado")
}

func (ag *AlgoGram) MostrarLikes(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(ag.posts) {
		fmt.Println(ERR_POST_INEXISTENTE_O_SIN_LIKES)
		return
	}
	ag.posts[id].mostrarLikes()
}
