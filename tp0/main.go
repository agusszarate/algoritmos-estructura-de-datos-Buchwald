package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const ruta1 = "archivo1.in"
const ruta2 = "archivo2.in"

func leer(ruta string) []int {
	archivo, err := os.Open(ruta)

	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", err, ruta)
	}

	defer archivo.Close()

	lista := []int{}

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea, err := strconv.Atoi(s.Text())

		if err != nil {
			fmt.Printf("Error %v al convertir linea %d", err, linea)
		}

		lista = append(lista, linea)
	}

	return lista
}

func imprimirOrdenado(lista []int) {

	ejercicios.Seleccion(lista)

	for i := 0; i < len(lista); i++ {
		num := lista[i]

		fmt.Println(num)
	}
}

func main() {

	archivo1 := leer(ruta1)
	archivo2 := leer(ruta2)
	resultado := ejercicios.Comparar(archivo1, archivo2)

	switch resultado {
	case -1:
		imprimirOrdenado(archivo2)
	case 0:
		imprimirOrdenado(archivo1)
	case 1:
		imprimirOrdenado(archivo1)
	}

}
