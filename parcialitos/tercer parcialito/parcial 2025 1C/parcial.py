from grafo import Grafo
from cola import Cola
from heap import heap_sort
# =============================================================================
# DATOS DEL GRAFO (IMAGEN 1)
# =============================================================================
# Grafo no dirigido y pesado. Representado como lista de adyacencia.
# Se incluyen ambas direcciones para cada arista (u->v y v->u).

grafo_dorso = {
    'A': {'E': 1, 'D': 7},
    'C': {'D': 6, 'I': 1, 'J': 0},
    'D': {'A': 7, 'E': 5, 'G': 7, 'I': 6, 'C': 6},
    'E': {'A': 1, 'F': 3, 'G': 2, 'D': 5},
    'F': {'E': 3, 'G': 4},
    'G': {'E': 2, 'F': 4, 'D': 7, 'I': 6},
    'I': {'C': 1, 'D': 6, 'G': 6, 'J': 2},
    'J': {'C': 0, 'I': 2}
}

# =============================================================================
# EJERCICIO 1: Árbol de Tendido Mínimo (Kruskal)
# =============================================================================
"""
1. a. Realizar un seguimiento de aplicar el algoritmo de Kruskal para obtener el 
      árbol de tendido mínimo del grafo del dorso.

   b. Encontrar, si es posible, una arista en este grafo al que se le pueda asignar 
      un valor más bajo, pero aún así sea imposible que esa arista sea parte de un 
      árbol de tendido mínimo válido del grafo. En ese caso, indicar cuál es dicha 
      arista y qué valor se le asignaría para cumplir la condición.
"""

def seguimiento_kruskal(grafo):
    # Espacio para tu implementación o comentarios del seguimiento
    # Orden de aristas: (C,J,0), (A,E,1), (C,I,1), (E,G,2), (I,J,2)...
    pass

# Respuesta 1b:
# (Escribe aquí tu justificación)


# =============================================================================
# EJERCICIO 2: AlgoMaps y Puntos Importantes
# =============================================================================
"""
2. En una pequeña ciudad costera, todos los habitantes usan el AlgoMaps, una aplicación 
   que les ayuda a decidir sus rutas de viaje diarias. Conociendo su funcionamiento, 
   sabemos que siempre ofrece la ruta del camino más corto de un punto a otro dadas 
   las rutas existentes, sabiendo que cuenta con la información de la distancia cubierta 
   por cada camino entre punto y punto. 
   
   En el futuro se planifican ciertas obras en diferentes puntos de la ciudad que 
   seguramente afecten a la red de transporte. Para evitar afectar en demasía a los 
   ciudadanos, se busca analizar la información disponible para planificar las obras. 
   Sería deseable evitar bloquear el transporte de los puntos más importantes de la ciudad.

   a. Explicar cómo la información disponible se puede modelar con un grafo, y explicar 
      sus características. ¿Qué información debemos solicitarle a AlgoMaps?

   b. Desarrollar una función puntosImportantes(grafo, k) que reciba el grafo definido 
      en el punto (a) y que devuelva los k puntos más importantes en la ciudad, según 
      lo modelado. Indicar y justificar su complejidad.
"""

# Respuesta 2a:
# El grafo se modela con Intersecciones/Puntos como Vértices y Calles como Aristas (Pesadas por distancia).
# Información a solicitar: Centralidad...

def puntosImportantes(grafo, k):
    # Tip: "Puntos más importantes" en rutas más cortas suele referirse a 
    # Centralidad de Intermediación (Betweenness Centrality).
    pass


# =============================================================================
# EJERCICIO 3: Zanahorias Podridas (Difusión)
# =============================================================================
"""
3. Se tiene un grafo no dirigido sin pesos, que representa una huerta de zanahorias. 
   Cada vértice representa a una zanahoria y cada arista (v, w) indica que v y w se 
   encuentran lo suficientemente cerca como para que si una se pudre, entonces podrá 
   "contagiar a la otra". Luego, esta nueva zanahoria contagiada infectará a las otras 
   que se encuentren cercanas. 
   
   En una unidad de tiempo, una zanahoria podrida infecta a todas las que tiene a su 
   alcance. Suponiendo que existe una única componente conexa en dicho grafo, implementar 
   una función que reciba un grafo de dichas características y una zanahoria podrida 
   inicial, e indique cuál o cuáles son las últimas zanahorias en pudrirse. 
   Indicar y justificar la complejidad de la función.

   Por ejemplo, en el grafo del ejercicio 1, pero quitando los pesos:
    - Si empieza podrida la zanahoria D, las últimas zanahorias en pudrirse serían F y J.
    - Si empieza podrida la zanahoria E, la última en pudrirse sería la J.
"""
def func_cmp(a, b):
    if a[0] > b[0]:
        return 1
    elif a[0] == b[0]:
        return 0
    else:
        return -1

def bfs(grafo: Grafo, origen):
    distancias = {}
    visitados = set()

    for v in grafo:
        distancias[v] = float('inf')

    cola = Cola()
    distancias[origen] = 0
    cola.encolar((0, origen))
    visitados.add(origen)

    while not cola.esta_vacia(): 
        dist, v = cola.desencolar()

        for w in grafo.obtener_adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                distancias[w] = dist + 1
                cola.encolar((dist + 1, w))
    
    return distancias



def ultimas_zanahorias(grafo: Grafo, inicio):
    distancias = bfs(grafo, inicio) #O(V + E)

    max_dist = -1
    for v in distancias:
        if distancias[v] > max_dist:
            max_dist = distancias[v]

    ultimas = []
    for v in distancias:
        if distancias[v] == max_dist:
            ultimas.append(v)

    return ultimas




    

# Test con el ejemplo del ejercicio (usando grafo_dorso ignorando pesos):
# print(ultimas_zanahorias(grafo_dorso, 'D')) # Debería dar ['F', 'J']
# print(ultimas_zanahorias(grafo_dorso, 'E')) # Debería dar ['J']