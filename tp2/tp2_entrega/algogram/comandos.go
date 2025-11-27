package algogram

import (
	"bufio"
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

		switch instruccion {
		case "login":
			ag.Login(parametro)
		case "logout":
			ag.Logout()
		case "publicar":
			ag.PublicarPost(parametro)
		case "ver_siguiente_feed":
			ag.VerSiguienteFeed()
		case "likear_post":
			ag.LikearPost(parametro)
		case "mostrar_likes":
			ag.MostrarLikes(parametro)
		}
	}
}
