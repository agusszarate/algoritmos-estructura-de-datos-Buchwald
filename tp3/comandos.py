#!/usr/bin/python3

import biblioteca

pagerank_cache = None
cfcs_cache = {}

def imprimir_camino_con_costo(camino):
    if camino is None:
        print("No se encontro recorrido")
    else:
        print(" -> ".join(camino))
        print(f"Costo: {len(camino) - 1}")

def listar_operaciones():
    print("camino")
    print("mas_importantes")
    print("conectados")
    print("ciclo")
    print("lectura")
    print("diametro")
    print("rango")
    print("comunidad")
    print("navegacion")
    print("clustering")


def camino(grafo, params):
    if len(params) != 2:
        return

    origen, destino = params[0], params[1]
    camino_resultado = biblioteca.camino_minimo(grafo, origen, destino)
    imprimir_camino_con_costo(camino_resultado)


def mas_importantes(grafo, params):
    global pagerank_cache

    if len(params) != 1:
        return

    try:
        n = int(params[0])
    except ValueError:
        return

    if pagerank_cache is None:
        pagerank_cache = biblioteca.pagerank(grafo)

    top_n = biblioteca.obtener_top_n(pagerank_cache, n)
    print(", ".join(top_n))


def conectados(grafo, params):
    global cfcs_cache

    if len(params) != 1:
        return

    pagina = params[0]

    if not grafo.vertice_pertenece(pagina):
        return

    if pagina in cfcs_cache:
        cfc = cfcs_cache[pagina]
    else:
        cfc = biblioteca.componente_fuertemente_conexa(grafo, pagina)
        for v in cfc:
            cfcs_cache[v] = cfc

    print(", ".join(sorted(cfc)))


def ciclo(grafo, params):
    if len(params) != 2:
        return

    pagina = params[0]
    try:
        n = int(params[1])
    except ValueError:
        return

    ciclo_resultado = biblioteca.ciclo_largo_n(grafo, pagina, n)

    if ciclo_resultado is None:
        print("No se encontro recorrido")
    else:
        print(" -> ".join(ciclo_resultado))


def lectura(grafo, params):
    if len(params) < 1:
        return

    orden = biblioteca.ordenamiento_topologico(grafo, params)

    if orden is None:
        print("No existe forma de leer las paginas en orden")
    else:
        print(", ".join(orden))


def diametro(grafo, params):
    camino_resultado = biblioteca.diametro(grafo)
    imprimir_camino_con_costo(camino_resultado)


def rango(grafo, params):
    if len(params) != 2:
        return

    pagina = params[0]
    try:
        n = int(params[1])
    except ValueError:
        return

    cantidad = biblioteca.bfs_distancia(grafo, pagina, n)
    print(cantidad)


def comunidad(grafo, params):
    if len(params) != 1:
        return

    pagina = params[0]
    comunidad_resultado = biblioteca.obtener_comunidad(grafo, pagina)
    print(", ".join(comunidad_resultado))


def navegacion(grafo, params):
    if len(params) != 1:
        return

    origen = params[0]
    camino_resultado = biblioteca.navegacion_primer_link(grafo, origen)

    if len(camino_resultado) == 1:
        print(camino_resultado[0])
    else:
        print(" -> ".join(camino_resultado))


def clustering(grafo, params):
    if len(params) == 0:
        c = biblioteca.coeficiente_clustering_promedio(grafo)
        print(f"{c:.3f}")
    elif len(params) == 1:
        pagina = params[0]
        if not grafo.vertice_pertenece(pagina):
            print("0.000")
            return
        c = biblioteca.coeficiente_clustering_vertice(grafo, pagina)
        print(f"{c:.3f}")


COMANDOS = {
    "listar_operaciones": listar_operaciones,
    "camino": camino,
    "mas_importantes": mas_importantes,
    "conectados": conectados,
    "ciclo": ciclo,
    "lectura": lectura,
    "diametro": diametro,
    "rango": rango,
    "comunidad": comunidad,
    "navegacion": navegacion,
    "clustering": clustering,
}


def ejecutar_comando(grafo, comando, params):
    if comando in COMANDOS:
        if comando == "listar_operaciones":
            COMANDOS[comando]()
        else:
            COMANDOS[comando](grafo, params)
