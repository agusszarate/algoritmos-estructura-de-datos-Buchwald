from collections import deque

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

        if v == destino:
            camino = []
            actual = destino
            while actual is not None:
                camino.append(actual)
                actual = padres[actual]
            return camino[::-1] 

        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                cola.append(w)

    return None

def cantidad_en_rango(grafo, origen, n):
    if not grafo.vertice_pertenece(origen):
        return 0

    visitados = {origen}
    cola = deque([(origen, 0)]) 
    contador = 0

    while cola:
        v, dist = cola.popleft()

        if dist == n:
            contador += 1
            continue 
            
        if dist > n:
            break

        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                cola.append((w, dist + 1))
    
    return contador

def orden_topologico(grafo, lista_vertices):
    grados = _calcular_grados_entrada(grafo, lista_vertices)
    subconjunto_set = set(lista_vertices)

    cola = deque([v for v in lista_vertices if v in grados and grados[v] == 0])
    resultado = []

    while cola:
        v = cola.popleft()
        resultado.append(v)

        for w in grafo.adyacentes(v):
            if w in subconjunto_set:
                grados[w] -= 1
                if grados[w] == 0:
                    cola.append(w)

    if len(resultado) != len(lista_vertices):
        return None
    
    return resultado[::-1]

def clustering_individual(grafo, vertice):
    todos_adyacentes = grafo.adyacentes(vertice)
    
    adyacentes = [v for v in todos_adyacentes if v != vertice]
    
    cantidad_ady = len(adyacentes)

    if cantidad_ady < 2:
        return 0.0

    links_entre_vecinos = 0
    
    for i in range(cantidad_ady):
        for j in range(cantidad_ady):
            if i == j: continue
            
            v = adyacentes[i]
            w = adyacentes[j]
            
            if grafo.estan_unidos(v, w):
                links_entre_vecinos += 1

    max_links_posibles = cantidad_ady * (cantidad_ady - 1)
    return links_entre_vecinos / max_links_posibles

def clustering_promedio(grafo):
    total = 0.0
    cantidad = 0
   
    for v in grafo.obtener_vertices():
        total += clustering_individual(grafo, v)
        cantidad += 1
        
    if cantidad == 0: return 0.0
    return total / cantidad


def navegacion_primer_link(grafo, origen):
    camino = [origen]
    
    actual = origen
    
    for _ in range(20): 
        adyacentes = grafo.adyacentes(actual)
        
        if not adyacentes:
            break 
            
        siguiente = adyacentes[0] 
        camino.append(siguiente)
        actual = siguiente
        
    return camino


def _calcular_grados_entrada(grafo, lista_vertices):
    grados = {v: 0 for v in lista_vertices}
    subconjunto_set = set(lista_vertices)
    
    for v in lista_vertices:
        if not grafo.vertice_pertenece(v):
            continue
            
        for w in grafo.adyacentes(v):
            if w in subconjunto_set:
                grados[w] += 1
                
    return grados