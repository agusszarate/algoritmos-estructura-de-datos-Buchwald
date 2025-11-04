package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tdas/diccionario"
)

type Usuario struct {
	nombre   string
	posicion int
}

func crearUsuario(nombre string, posicion int) *Usuario {

	return &Usuario{
		nombre:   nombre,
		posicion: posicion,
	}
}

type AlgoGram struct {
	usuarios        diccionario.Diccionario[string, *Usuario]
	usuarioLoggeado *Usuario
	listaUsuarios   []*Usuario
}

func crearAlgoGram() *AlgoGram {
	igualdadString := func(a, b string) bool { return a == b }
	return &AlgoGram{
		usuarios:        diccionario.CrearHash[string, *Usuario](igualdadString),
		usuarioLoggeado: nil,
		listaUsuarios:   make([]*Usuario, 0),
	}
}

func (ag *AlgoGram) cargarUsuarios(archivo string) error {
	file, err := os.Open(archivo)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	posicion := 0
	for scanner.Scan() {
		nombre := strings.TrimSpace(scanner.Text())
		if nombre == "" {
			continue
		}
		usuario := crearUsuario(nombre, posicion)
		ag.usuarios.Guardar(nombre, usuario)
		ag.listaUsuarios = append(ag.listaUsuarios, usuario)
		posicion++
	}

	return scanner.Err()
}

func (ag *AlgoGram) login(nombre string) string {
	if ag.usuarioLoggeado != nil {
		return "Error: Ya habia un usuario loggeado"
	}
	if !ag.usuarios.Pertenece(nombre) {
		return "Error: usuario no existente"
	}
	ag.usuarioLoggeado = ag.usuarios.Obtener(nombre)
	return "Hola " + nombre
}

func (ag *AlgoGram) logout() string {
	if ag.usuarioLoggeado == nil {
		return "Error: no habia usuario loggeado"
	}
	ag.usuarioLoggeado = nil
	return "Adios"
}

func (ag *AlgoGram) procesarComando(linea string) string {
	partes := strings.Fields(linea)
	if len(partes) == 0 {
		return ""
	}

	comando := partes[0]

	switch comando {
	case "login":
		if len(partes) < 2 {
			return "Error: comando invalido"
		}
		nombre := strings.TrimPrefix(linea, "login ")
		return ag.login(nombre)

	case "logout":
		return ag.logout()

	default:
		return ""
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s <archivo_usuarios>\n", os.Args[0])
		os.Exit(1)
	}

	ag := crearAlgoGram()
	if err := ag.cargarUsuarios(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Error al cargar usuarios: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		resultado := ag.procesarComando(linea)
		if resultado != "" {
			fmt.Println(resultado)
		}
	}
}
