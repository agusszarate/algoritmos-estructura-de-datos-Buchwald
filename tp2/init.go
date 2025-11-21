package main

import (
	"bufio"
	"os"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
)

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
