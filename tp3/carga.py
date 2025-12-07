#!/usr/bin/python3
import sys
from grafo import Grafo

def cargar_grafo(archivo):

    grafo = Grafo(es_dirigido=True)

    try:
        with open(archivo, 'r', encoding='utf-8') as f:
            for linea in f:

                campos = linea.strip().split('\t')
                
                if not campos:
                    continue

                origen = campos[0]
                grafo.agregar_vertice(origen)

                for destino in campos[1:]:
                    if not destino:
                        continue
                        
                    grafo.agregar_vertice(destino)
                    grafo.agregar_arista(origen, destino)

        return grafo

    except FileNotFoundError:
        print(f"Error: El archivo '{archivo}' no fue encontrado.", file=sys.stderr)
        sys.exit(1) 
    except Exception as e:
        print(f"Error inesperado al cargar el grafo: {e}", file=sys.stderr)
        sys.exit(1)