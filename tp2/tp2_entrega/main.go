package main

import (
	"bufio"
	"os"
	"strings"
	"tp2/algogram"
)

func cargarUsuarios(ruta string) ([]*algogram.Usuario, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	var lista []*algogram.Usuario
	scanner := bufio.NewScanner(archivo)
	idx := 0

	for scanner.Scan() {
		nombre := strings.TrimSpace(scanner.Text())
		if nombre == "" {
			continue
		}

		usuario := algogram.CrearUsuario(nombre, idx)
		lista = append(lista, usuario)
		idx++
	}
	return lista, scanner.Err()
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	usuarios, err := cargarUsuarios(os.Args[1])
	if err != nil {
		return
	}

	sistema := algogram.CrearAlgoGram(usuarios)
	sistema.Ejecutar()
}
