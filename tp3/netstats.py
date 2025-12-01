#!/usr/bin/python3

import sys
from carga import cargar_grafo
from comandos import ejecutar_comando

def parsear_comando(linea):
    linea = linea.strip()
    if not linea:
        return None, []

    partes = linea.split(None, 1)
    comando = partes[0]

    if len(partes) > 1:
        params = [p.strip() for p in partes[1].split(',')]
    else:
        params = []

    return comando, params

def procesar_comandos(grafo):
    for linea in sys.stdin:
        comando, params = parsear_comando(linea)
        if comando:
            ejecutar_comando(grafo, comando, params)

def main():
    if len(sys.argv) != 2:
        print("Uso: ./netstats <archivo>", file=sys.stderr)
        sys.exit(1)

    archivo = sys.argv[1]
    grafo = cargar_grafo(archivo)

    procesar_comandos(grafo)

if __name__ == "__main__":
    main()
