from grafo import GrafoDirigido, Grafo
from cola import Cola
from heap import Heap

# Implementar un algoritmo que reciba un grafo dirigido y acíclico y determine si dicho grafo admite un único orden topológico. 
# Indicar y justificar la complejidad de la función. 
# Pista: pensar qué condición puede darse para que exista más de un posible orden topológico.

def obtener_grados(grafo: GrafoDirigido): #Total O(2V + E) == O(V + E)
    grados = {} #O(1)

    for v in grafo: #O(V)
        grados[v] = 0#O(1)

    for v in grafo: # O(V)
        for w in grafo.obtener_adyacentes(v): #O(E)
            grados[w] += 1 #O(1)

    return grados #O(1)


def admite_unico_order_topologico(grafo: GrafoDirigido): #Total O(V + E)
    grados = obtener_grados(grafo) #O(V + E) por cada vertice miro sus adyacentes

    cola = Cola() #O(1)

    for v in grados.keys(): #O(V)
        if grados[v] == 0: #O(1)
            cola.encolar(v) #O(1)

    while not cola.esta_vacia(): # O(V + E) por cada vertice miro sus adyacentes

        if len(cola) > 1: #O(1)
            return False #O(1)

        v = cola.desencolar() #O(1)

        for w in grafo.obtener_adyacentes(v): #O(E)
            grados[w] -= 1 #O(1)
            if grados[w] == 0: #O(1)
                cola.encolar(w) #O(1)

    return True #O(1)

# Se tiene una app de tránsito que indica el mejor camino desde un punto A a un punto B.
# Esta app trabaja con un grafo dirigido de las calles de la Ciudad Autónoma de Buenos Aires, 
# en donde las esquinas son los vértices del Grafo, y las aristas son las calles que unen dichas esquinas.
# Las aristas tienen como peso: el tiempo (en segundos) que demora en recorrerse, 
# y su nombre (por ejemplo, "Paseo Colón 800-900"). 
# Además, hay un informe respecto al estado de las calles que indica para cada uno 3 posibles estados:

# HABILITADA: Se puede transitar de forma normal.
# CONGESTIONADA: Se tarda el doble en transitar.
# CORTADA: Imposible transitar. 

# Se pide implementar un algoritmo que reciba el grafo, 
# el informe (un diccionario en el que las claves son los nombres de las aristas, 
# y como valores su estado correspondiente), un punto A y un punto B,
# y determine la forma más rápida para llegar desde A hasta B, considerando el estado dado por el informe. 
# Indicar y justificar la complejidad del algoritmo implementado.

CONGESTIONADA = 'CONGESTIONADA'

CORTADA = 'CORTADA'

#En Dijkstra, se marca como visitado después de desencolar y antes de procesar, no al encolar.
def mapa(grafo: Grafo, informe, origen, destino):
    visitados = set()
    padres = {}
    distancias = {}

    for v in grafo:
        distancias[v] = float('Inf')
        padres[v] = None

    def funcion_comp(a, b):
        return b[0] - a[0]

    cola = Heap(funcion_comp)

    distancias[origen] = 0
    cola.encolar(origen)

    while not cola.esta_vacia():
        v = cola.desencolar()

        if v == destino:
            return

        if v not in visitados:
            visitados.add(v)
            for w in grafo.obtener_adyacentes(v):
                if w not in visitados:
                    tiempo, nombre = grafo.peso_arista(v, w)
                    if informe[nombre] == CONGESTIONADA:
                        tiempo *= 2
                    elif informe[nombre] == CORTADA:
                        continue

                    dist_v = distancias[v] + tiempo
                    
                    if dist_v < distancias[w]:
                        distancias[w] = dist_v
                        cola.encolar(w)

def recontruir_camino(padres, destino):
    inicio = destino
    camino = [destino]

    while inicio is not None:
        camino.append(inicio)
        inicio = padres[destino]

                
                



