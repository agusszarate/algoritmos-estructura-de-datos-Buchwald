#!/usr/bin/python3

from collections import deque
import random

def camino_minimo(grafo, origen, destino):
    if not grafo.vertice_pertenece(origen) or not grafo.vertice_pertenece(destino):
        return None

    if origen == destino:
        return [origen]

    visitados = {origen}
    padres = {origen: None}
    cola = deque([origen])

    while cola:
        v = cola.popleft()

        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                cola.append(w)

                if w == destino:
                    camino = []
                    actual = destino
                    while actual is not None:
                        camino.append(actual)
                        actual = padres[actual]
                    return list(reversed(camino))

    return None


def pagerank(grafo, iteraciones=20, d=0.85):
    vertices = grafo.obtener_vertices()
    n = len(vertices)

    if n == 0:
        return {}

    pr = {v: 1.0 / n for v in vertices}
    grados = {v: grafo.grado_salida(v) for v in vertices}

    for _ in range(iteraciones):
        nuevo_pr = {}

        for v in vertices:
            suma = 0.0
            for w in vertices:
                if grafo.estan_unidos(w, v):
                    if grados[w] > 0:
                        suma += pr[w] / grados[w]

            nuevo_pr[v] = (1 - d) / n + d * suma

        pr = nuevo_pr

    return pr


def obtener_top_n(ranking, n):
    items = sorted(ranking.items(), key=lambda x: x[1], reverse=True)
    return [item[0] for item in items[:n]]


def tarjan(grafo):
    indice = [0]
    stack = []
    lowlink = {}
    index = {}
    en_stack = set()
    cfcs = []

    def strongconnect(v):
        index[v] = indice[0]
        lowlink[v] = indice[0]
        indice[0] += 1
        stack.append(v)
        en_stack.add(v)

        for w in grafo.adyacentes(v):
            if w not in index:
                strongconnect(w)
                lowlink[v] = min(lowlink[v], lowlink[w])
            elif w in en_stack:
                lowlink[v] = min(lowlink[v], index[w])

        if lowlink[v] == index[v]:
            cfc = set()
            while True:
                w = stack.pop()
                en_stack.remove(w)
                cfc.add(w)
                if w == v:
                    break
            cfcs.append(cfc)

    for v in grafo:
        if v not in index:
            strongconnect(v)

    return cfcs


def componente_fuertemente_conexa(grafo, vertice):
    cfcs = tarjan(grafo)

    for cfc in cfcs:
        if vertice in cfc:
            return cfc

    return set()


def ciclo_largo_n(grafo, origen, n):
    if not grafo.vertice_pertenece(origen):
        return None

    if n < 2:
        return None

    def backtrack(actual, camino, visitados):
        if len(camino) == n:
            if grafo.estan_unidos(actual, origen):
                return camino + [origen]
            return None

        for w in grafo.adyacentes(actual):
            if w not in visitados:
                visitados.add(w)
                resultado = backtrack(w, camino + [w], visitados)
                if resultado:
                    return resultado
                visitados.remove(w)

        return None

    visitados = {origen}
    return backtrack(origen, [origen], visitados)


def ordenamiento_topologico(grafo, vertices_subconjunto):
    grado_entrada = {v: 0 for v in vertices_subconjunto}
    vertices_set = set(vertices_subconjunto)

    for v in vertices_subconjunto:
        for w in grafo.adyacentes(v):
            if w in vertices_set:
                grado_entrada[w] += 1

    cola = deque([v for v in vertices_subconjunto if grado_entrada[v] == 0])
    resultado = []

    while cola:
        v = cola.popleft()
        resultado.append(v)

        for w in grafo.adyacentes(v):
            if w in vertices_set:
                grado_entrada[w] -= 1
                if grado_entrada[w] == 0:
                    cola.append(w)

    if len(resultado) != len(vertices_subconjunto):
        return None

    return resultado


def bfs_distancia(grafo, origen, distancia_objetivo):
    if not grafo.vertice_pertenece(origen):
        return 0

    visitados = {origen}
    distancias = {origen: 0}
    cola = deque([origen])

    while cola:
        v = cola.popleft()

        if distancias[v] >= distancia_objetivo:
            continue

        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                distancias[w] = distancias[v] + 1
                cola.append(w)

    return sum(1 for v in distancias if distancias[v] == distancia_objetivo)


def label_propagation(grafo, max_iteraciones=10):
    etiquetas = {v: i for i, v in enumerate(grafo)}
    vertices = grafo.obtener_vertices()

    for _ in range(max_iteraciones):
        cambio = False
        random.shuffle(vertices)

        for v in vertices:
            contador = {}
            for w in grafo.adyacentes(v):
                etiq = etiquetas[w]
                contador[etiq] = contador.get(etiq, 0) + 1

            if not contador:
                continue

            max_count = max(contador.values())
            etiquetas_max = [e for e, c in contador.items() if c == max_count]
            nueva_etiqueta = min(etiquetas_max)

            if etiquetas[v] != nueva_etiqueta:
                etiquetas[v] = nueva_etiqueta
                cambio = True

        if not cambio:
            break

    return etiquetas


def obtener_comunidad(grafo, vertice):
    etiquetas = label_propagation(grafo)

    if vertice not in etiquetas:
        return []

    etiqueta_objetivo = etiquetas[vertice]
    return [v for v, e in etiquetas.items() if e == etiqueta_objetivo]


def navegacion_primer_link(grafo, origen, max_pasos=20):
    if not grafo.vertice_pertenece(origen):
        return [origen]

    camino = [origen]
    actual = origen
    visitados = {origen}

    for _ in range(max_pasos - 1):
        adyacentes = grafo.adyacentes(actual)

        if not adyacentes:
            break

        siguiente = adyacentes[0]
        camino.append(siguiente)

        if siguiente in visitados:
            actual = siguiente
        else:
            visitados.add(siguiente)
            actual = siguiente

    return camino


def coeficiente_clustering_vertice(grafo, vertice):
    adyacentes = grafo.adyacentes(vertice)
    k = len(adyacentes)

    if k < 2:
        return 0.0

    aristas_entre_adyacentes = 0
    for v in adyacentes:
        if v == vertice:
            continue
        for w in adyacentes:
            if w == vertice:
                continue
            if v != w and grafo.estan_unidos(v, w):
                aristas_entre_adyacentes += 1

    return aristas_entre_adyacentes / (k * (k - 1))


def coeficiente_clustering_promedio(grafo):
    suma = 0.0
    vertices = grafo.obtener_vertices()
    n = len(vertices)

    if n == 0:
        return 0.0

    for v in vertices:
        suma += coeficiente_clustering_vertice(grafo, v)

    return suma / n


def diametro(grafo):
    vertices = grafo.obtener_vertices()
    max_distancia = -1
    max_camino = None

    for origen in vertices:
        distancias = {}
        padres = {}
        visitados = {origen}
        distancias[origen] = 0
        padres[origen] = None
        cola = deque([origen])

        while cola:
            v = cola.popleft()

            for w in grafo.adyacentes(v):
                if w not in visitados:
                    visitados.add(w)
                    distancias[w] = distancias[v] + 1
                    padres[w] = v
                    cola.append(w)

        for destino in distancias:
            if distancias[destino] > max_distancia:
                max_distancia = distancias[destino]
                camino = []
                actual = destino
                while actual is not None:
                    camino.append(actual)
                    actual = padres[actual]
                max_camino = list(reversed(camino))

    return max_camino if max_camino else None
