package main

import (
	TDAColaPrioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
)

const (
	_ERR_USUARIO_YA_LOGGEADO                    = "Error: Ya habia un usuario loggeado"
	_ERR_USUARIO_NO_EXISTENTE                   = "Error: usuario no existente"
	_ERR_USUARIO_NO_LOGGEADO                    = "Error: no habia usuario loggeado"
	_ERR_USUARIO_NO_LOGGEADO_O_POSTS_VISTOS     = "Usuario no loggeado o no hay mas posts para ver"
	_ERR_USUARIO_NO_LOGGEADO_O_POST_INEXISTENTE = "Error: Usuario no loggeado o Post inexistente"
	_ERR_POST_INEXISTENTE_O_SIN_LIKES           = "Error: Post inexistente o sin likes"
)

type Usuario struct {
	nombre   string
	afinidad int
	feed     TDAColaPrioridad.ColaPrioridad[itemFeed]
}

type Post struct {
	id      int
	autor   string
	mensaje string
	likes   TDADiccionario.DiccionarioOrdenado[string, bool]
}

type itemFeed struct {
	afinidad int
	idPost   int
}

type AlgoGram struct {
	usuarios        TDADiccionario.Diccionario[string, *Usuario]
	posts           []*Post
	usuarioLoggeado *Usuario
	listaUsuarios   []*Usuario
}
