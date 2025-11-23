package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
)

type AlgoGram struct {
	usuarios        TDADiccionario.Diccionario[string, *Usuario]
	posts           []*Post
	usuarioLoggeado *Usuario
	listaUsuarios   []*Usuario
}

//-------------------------------------

func crearAlgoGram(ruta string) (*AlgoGram, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	ag := &AlgoGram{
		usuarios: TDADiccionario.CrearHash[string, *Usuario](func(a, b string) bool { return a == b }),
		posts:    make([]*Post, 0),
	}

	scanner := bufio.NewScanner(archivo)
	idx := 0
	for scanner.Scan() {
		nombre := strings.TrimSpace(scanner.Text())
		if nombre == "" {
			continue
		}
		usuario := &Usuario{
			nombre:   nombre,
			afinidad: idx,
			feed:     TDAColaPrioridad.CrearHeap(cmpFeed),
		}
		ag.usuarios.Guardar(nombre, usuario)
		ag.listaUsuarios = append(ag.listaUsuarios, usuario)
		idx++
	}
	return ag, scanner.Err()
}

//-------------------------------------

func (ag *AlgoGram) login(nombre string) {
	if ag.usuarioLoggeado != nil {
		fmt.Println(_ERR_USUARIO_YA_LOGGEADO)
		return
	}
	if !ag.usuarios.Pertenece(nombre) {
		fmt.Println(_ERR_USUARIO_NO_EXISTENTE)
		return
	}
	ag.usuarioLoggeado = ag.usuarios.Obtener(nombre)
	fmt.Printf("Hola %s\n", nombre)
}

func (ag *AlgoGram) logout() {
	if ag.usuarioLoggeado == nil {
		fmt.Println(_ERR_USUARIO_NO_LOGGEADO)
		return
	}
	ag.usuarioLoggeado = nil
	fmt.Println("Adios")
}

func (ag *AlgoGram) publicarPost(mensaje string) {
	if ag.usuarioLoggeado == nil {
		fmt.Println(_ERR_USUARIO_NO_LOGGEADO)
		return
	}
	nuevoPost := &Post{
		id:      len(ag.posts),
		autor:   ag.usuarioLoggeado.nombre,
		mensaje: mensaje,
		likes:   TDADiccionario.CrearABB[string, bool](cmpString),
	}
	ag.posts = append(ag.posts, nuevoPost)

	for _, usuario := range ag.listaUsuarios {
		if usuario.nombre == ag.usuarioLoggeado.nombre {
			continue
		}
		afinidad := abs(ag.usuarioLoggeado.afinidad - usuario.afinidad)
		usuario.feed.Encolar(itemFeed{afinidad: afinidad, idPost: nuevoPost.id})
	}
	fmt.Println("Post publicado")
}

func (ag *AlgoGram) verSiguienteFeed() {
	if ag.usuarioLoggeado == nil || ag.usuarioLoggeado.feed.EstaVacia() {
		fmt.Println(_ERR_USUARIO_NO_LOGGEADO_O_POSTS_VISTOS)
		return
	}
	item := ag.usuarioLoggeado.feed.Desencolar()
	post := ag.posts[item.idPost]
	fmt.Printf("Post ID %d\n%s dijo: %s\nLikes: %d\n", post.id, post.autor, post.mensaje, post.likes.Cantidad())
}

func (ag *AlgoGram) likearPost(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil || ag.usuarioLoggeado == nil || id < 0 || id >= len(ag.posts) {
		fmt.Println(_ERR_USUARIO_NO_LOGGEADO_O_POST_INEXISTENTE)
		return
	}
	ag.posts[id].likes.Guardar(ag.usuarioLoggeado.nombre, true)
	fmt.Println("Post likeado")
}

func (ag *AlgoGram) mostrarLikes(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(ag.posts) {
		fmt.Println(_ERR_POST_INEXISTENTE_O_SIN_LIKES)
		return
	}
	post := ag.posts[id]
	if post.likes.Cantidad() == 0 {
		fmt.Println(_ERR_POST_INEXISTENTE_O_SIN_LIKES)
		return
	}
	fmt.Printf("El post tiene %d likes:\n", post.likes.Cantidad())
	post.likes.Iterar(func(nombre string, _ bool) bool {
		fmt.Printf("\t%s\n", nombre)
		return true
	})
}
