from grafo import GrafoDirigido, Grafo
from cola import Cola
from heap import Heap
from pila import Pila

# Implementar un algoritmo que reciba un grafo dirigido y acíclico y determine si dicho grafo admite un único orden topológico. 
# Indicar y justificar la complejidad de la función. Pista: pensar qué condición puede darse para que exista más de un posible orden topológico.

def ejercicio1(grafo: GrafoDirigido):

    grados = obtenerGrados(grafo)

    cola = Cola()        
    for g in grados.keys():
        if (grados[g] == 0):
            cola.encolar(g)
    
    while not cola.esta_vacia():
        if (len(cola) > 1):
            return False
        elemento = cola.desencolar()
        for w in grafo.obtener_adyacentes(elemento):
            grados[w] -= 1

            if (grados[w] == 0):
                cola.encolar(w)

    return True

def obtenerGrados(grafo: GrafoDirigido):
    grados = {}

    for v in grafo:
        grados[v] = 0

    for v in grafo:
        for w in grafo.obtener_adyacentes(v):
            grados[w] += 1

    return grados





# Se tiene una app de tránsito que indica el mejor camino desde un punto A a un punto B.
# Esta app trabaja con un grafo dirigido de las calles de la Ciudad Autónoma de Buenos Aires, 
# en donde las esquinas son los vértices del Grafo, y las aristas son las calles que unen dichas esquinas.
# Las aristas tienen como peso: el tiempo (en segundos) que demora en recorrerse, y su nombre (por ejemplo, "Paseo Colón 800-900"). 
# Además, hay un informe respecto al estado de las calles que indica para cada uno 3 posibles estados:

# HABILITADA: Se puede transitar de forma normal.
# CONGESTIONADA: Se tarda el doble en transitar.
# CORTADA: Imposible transitar. 

# Se pide implementar un algoritmo que reciba el grafo, 
# el informe (un diccionario en el que las claves son los nombres de las aristas, y como valores su estado correspondiente), un punto A y un punto B,
# y determine la forma más rápida para llegar desde A hasta B, considerando el estado dado por el informe. Indicar y justificar la complejidad del algoritmo implementado.

CONGESTIONADA = 'CONGESTIONADA'

CORTADA = 'CORTADA'

def ejercicio2(grafo: GrafoDirigido, informe: dict, origen, destino):
    distancia = {}
    padre = {}

    for v in grafo:
        distancia[v] = float('inf')

    distancia[origen] = 0
    padre[origen] = None

    def cmp(a, b):
        return b[0] - a[0]

    cola = Heap(cmp)

    cola.encolar((0, origen))

    visitados = set()

    while not cola.esta_vacia():
        dist_v, v = cola.desencolar()

        if v == destino:
            return reconstruir_camino(padre, destino)
        
        if v in visitados:
            continue

        visitados.add(v)

        for w in grafo.obtener_adyacentes(v):
            if w in visitados:
                continue

            nombre, peso_vw = grafo.peso_arista(v, w)

            plus = 1

            if informe[nombre] == CONGESTIONADA:
                plus = 2
            elif informe[nombre] == CORTADA: 
                continue

            dist_via_v = distancia[v] + (peso_vw * plus)

            if dist_via_v < distancia[w]:
                distancia[w] = dist_via_v
                padre[w] = v 
                cola.encolar((dist_via_v, w))

    return None

        
def reconstruir_camino(padres, destino):
    camino = []

    destino = destino

    while destino is not None:
        camino.append(destino)

        destino = padres[destino]

    return camino[::-1]

    



# Implementar un algoritmo que reciba un grafo no dirigido y determine la cantidad máxima de aristas que se pueden agregar al mismo 
# de tal forma que no se reduzcan la cantidad de componentes conexas que hay en el mismo. 
# Indicar y justificar la complejidad del algoritmo implementado.

#La suma de los grados de todos los vértices es exactamente el doble de la cantidad de aristas.
#bfs siempre primero visitados.add antes de encolar

def max_aristas_igual_conexas(grafo: Grafo):
    visitados = set()

    total_aristas_posibles = 0

    for v_origen in grafo:
        if v_origen not in visitados:

            cantidad_grados_componente = 0
            cantidad_vertices_componente = 0
            
            cola = Cola()
            visitados.add(v_origen)
            cola.encolar(v_origen)

            while not cola.esta_vacia():
                v = cola.desencolar()

                cantidad_vertices_componente += 1

                adyacentes = grafo.obtener_adyacentes(v)

                cantidad_grados_componente += len(adyacentes)

                for w in adyacentes:
                    if w not in visitados:
                        visitados.add(w)
                        cola.encolar(w)

            total_aristas_componente = cantidad_grados_componente / 2

            total_aristas_posibles += (cantidad_vertices_componente * (cantidad_vertices_componente-1)) / 2 - total_aristas_componente
    
    return total_aristas_posibles
    

                




# Implementar un algoritmo que, 
# dado un grafo pesado (con pesos positivos), 
# un vértice v y otro w, 
# determine la cantidad de caminos mínimos que hay de v a w dentro del grafo. 
# Considerar que, justamente, podrían haber varios caminos de una misma distancia, que a su vez sean la mínima. 
# Indicar y justificar la complejidad del algoritmo implementado. 
# Por ejemplo, en el grafo de abajo hay 3 caminos mínimos del vértice A al vértice E: 
# A -> E, A -> B -> F -> D -> E, A -> H -> F -> D -> E, todos de costo 8.

def cant_caminos_minimos(grafo: GrafoDirigido, origen, destino):
    distancias = {}
    padres = {}

    for v in grafo:
        distancias[v] = float('inf')
        padres[v] = None

    def funcion_comp(a, b):
        return b[0] - a[0]

    cola = Heap(funcion_comp)

    distancias[origen] = 0
    cola.encolar((0, origen))
    contador_caminos = {origen: 1}

    while not cola.esta_vacia():
        _, v = cola.desencolar()

        for w in grafo.obtener_adyacentes(v):
            nueva_distancia = distancias[v] + grafo.peso_arista(v, w)

            if nueva_distancia < distancias[w]:
                distancias[w] = nueva_distancia
                cola.encolar((nueva_distancia, w))
                contador_caminos[w] = contador_caminos[v]
            elif nueva_distancia == distancias[w]:
                contador_caminos[w] += contador_caminos[v]

    return contador_caminos[destino]    
            

