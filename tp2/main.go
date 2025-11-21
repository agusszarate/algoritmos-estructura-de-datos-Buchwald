package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	ag, err := crearAlgoGram(os.Args[1])
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		if len(strings.TrimSpace(linea)) == 0 {
			continue
		}
		partes := strings.SplitN(linea, " ", 2)
		instruccion := partes[0]

		switch instruccion {
		case "login":
			ag.login(partes[1])
		case "logout":
			ag.logout()
		case "publicar":
			ag.publicarPost(partes[1])
		case "ver_siguiente_feed":
			ag.verSiguienteFeed()
		case "likear_post":
			ag.likearPost(partes[1])
		case "mostrar_likes":
			ag.mostrarLikes(partes[1])
		}
	}
}
