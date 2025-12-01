#!/usr/bin/python3

import sys
from grafo import Grafo

def cargar_grafo(archivo):
    grafo = Grafo()

    try:
        with open(archivo, 'r', encoding='utf-8') as f:
            for linea in f:
                partes = linea.strip().split('\t')
                if not partes:
                    continue

                pagina = partes[0]
                grafo.agregar_vertice(pagina)

                for i in range(1, len(partes)):
                    link = partes[i]
                    if link:
                        grafo.agregar_vertice(link)
                        grafo.agregar_arista(pagina, link)

        return grafo
    except FileNotFoundError:
        print(f"Error: no se pudo abrir el archivo {archivo}", file=sys.stderr)
        sys.exit(1)
    except Exception as e:
        print(f"Error al cargar el grafo: {e}", file=sys.stderr)
        sys.exit(1)
