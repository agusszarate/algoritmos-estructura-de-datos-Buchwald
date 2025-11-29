package algogram

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (ag *AlgoGram) Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		if len(strings.TrimSpace(linea)) == 0 {
			continue
		}
		partes := strings.SplitN(linea, " ", 2)
		instruccion := partes[0]
		parametro := ""
		if len(partes) > 1 {
			parametro = partes[1]
		}

		var mensaje string
		switch instruccion {
		case "login":
			mensaje = ag.Login(parametro)
		case "logout":
			mensaje = ag.Logout()
		case "publicar":
			mensaje = ag.PublicarPost(parametro)
		case "ver_siguiente_feed":
			mensaje = ag.VerSiguienteFeed()
		case "likear_post":
			mensaje = ag.LikearPost(parametro)
		case "mostrar_likes":
			mensaje = ag.MostrarLikes(parametro)
		}
		if mensaje != "" {
			fmt.Println(mensaje)
		}
	}
}
