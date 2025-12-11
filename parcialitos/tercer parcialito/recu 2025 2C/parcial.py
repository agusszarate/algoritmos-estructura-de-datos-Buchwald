from tp3.grafo import Grafo
from cola import Cola
from heap import Heap, crear_heap_desde_lista

"""
ASIGNATURA: Algoritmos y Estructuras de Datos (CB100) - Curso Buchwald
EXAMEN: 3.er parcialito (1R) - 04/12/2025

CONSIGNA GENERAL:
Resolvé los siguientes problemas en forma clara y legible. Podés incluir tantas 
funciones auxiliares como creas necesarias. Los algoritmos a implementar deben 
ser de la mejor complejidad posible dadas las características del problema.
"""

# =============================================================================
# EJERCICIO 1
# =============================================================================
"""
1. Realizar un seguimiento del algoritmo de Tarjan para obtener los puntos de 
   articulación del siguiente grafo, comenzando desde el vértice G.
"""

# Representación del grafo de la imagen (Lista de Adyacencia):
grafo_punto_1 = {
    'A': ['E', 'B', 'F', 'G'],
    'B': ['E', 'A', 'F', 'D', 'C'],
    'C': ['B', 'D'],
    'D': ['C', 'B', 'F', 'H'],
    'E': ['B', 'A'],
    'F': ['A', 'B', 'D', 'H', 'G'],
    'G': ['A', 'F', 'K'],
    'H': ['D', 'F'],
    'K': ['G']
}

# empieza en G orden 0, mb = 0, padre none, visitado
# miro adyacente A, orden 1 mb = 1, padre = g, marco visitado
# miro adyacente de A, E orden 2 mb 2, padre A, marco visitado
# miro adyacente de E, B orden 3 mb 3, padre E, marco visitado 
# miro adyacente de B, E padre ignoro, miro el otro A, ya fue visitado mb[b] = min(mb[B], orden[A]) queda mb[b] = orden[A]
# miro adyacente de B, F orden 4, mb 4, padre B, marco visitado,
# miro adyacente de F, A ya fue visitado mb[f] = min(mb[f], orden[A]) queda orden[A] = 1
# miro adyacente de F, B padre ignoro
# miro adyacente de F, D orden 5, mb 5, padre F marco visitado
# miro adyacente de D, C orden 6, mb 6, padre D marco visitado
# miro adyacente de C, miro B ya visitado mb[C] = min(mb[c], orden[B]) mb[c] = orden[B] = 3
# miro adyacente de C, D padre ignoro vuelvo
# miro adyacente de D, B ya visitado mb[D] = min(mb[D], orden[B]) mb[D] = orden[B] = 3
# miro adyacente de D, F padre ignoro
# miro adyacente de D, H orden 6, mb 6, padre D marco visitado
# miro adyacente de H, D padre ignoro
# miro adyacente de H, F ya visitado, mb[H] = min(mb[H], orden[f]) mb[H] = orden[F] = 4
# vuelvo recursividad
# miro adyacente de F, H ya visitad, mb[F] = min(mb[F], orden[H]) mb[F] = mb[F]
# miro

def resolver_ejercicio_1(grafo):
    # Tu implementación del seguimiento de Tarjan aquí
    pass


# =============================================================================
# EJERCICIO 2
# =============================================================================
"""
2. Un famoso ladrón (ladrón en serio) apodado como "El Seca Nucas", se encuentra 
   planificando un gran robo en el barrio de Barracas. La Policía se encuentra 
   haciendo un gran operativo para atraparlo. Luego de tantos robos están 
   finalmente un paso adelante suyo. Tienen información de cuando va a ser el 
   próximo asalto, pero no saben exactamente en qué comercio o banco será. 

   Se tiene:
   - Un grafo no dirigido de la ciudad donde los vértices representan las esquinas 
     de la misma, y las aristas, las cuadras que conectan esas esquinas. Cada una 
     tiene un peso que indica la distancia entre ellas. Por extrañas razones, 
     ¡todos los comercios/bancos están sobre esquinas en este barrio!
   
   - Una lista de esquinas de fuga F por las cuales si pasa El Seca Nucas se sabe 
     que ya se puede dar a la fuga.

   CONSIGNA:
   Implementar un algoritmo que obtenga el top 3 de lugares en las cuales El Seca 
   Nucas podría ejecutar su asalto. Una esquina es la mejor para realizar un asalto 
   si la suma de las distancias de los caminos mínimos contra todas las esquinas 
   de Fuga es la menor. Inclusive, alguna de estas esquinas también podría ser 
   parte del top 3. Indicar y justificar la complejidad del mismo.
"""
def func_cmp(a, b):
    if a[0] > b[0]:
      return 1
    elif a[0] == b[0]:
      return 0
    else:
      return -1


def caminos_minimos(grafo: Grafo, origen):
    distancias = {}
    padres = {}
    visitados = set()

    for v in grafo:
        distancias[v] = float('inf')
        padres[v] = None
    
    distancias[origen] = 0
    
    heap = Heap(func_cmp)

    heap.encolar((0, origen))

    while not heap.esta_vacia():
      distancia, v = heap.desencolar()

      if v not in visitados:
        visitados.add(v)

        for w in grafo.adyacentes(v):
          if w not in visitados:
            nueva_dist = distancia + grafo.peso_arista(v, w)
            if nueva_dist < distancias[w]:
              distancias[w] = nueva_dist
              padres[w] = v
              heap.encolar((nueva_dist, w))

    return padres, distancias

def resolver_ejercicio_2(grafo, esquinas_fuga):
  suma_caminos = {}

  for v in grafo:
     suma_caminos[v] = 0

  for esquina in esquinas_fuga:
     _, distancias = caminos_minimos(grafo, esquina)
     for v in distancias.keys():
        suma_caminos[v] += distancias[v]
      
  ranking = []
  for v in suma_caminos.keys():
     ranking.append((suma_caminos[v], v))

  heap = crear_heap_desde_lista(ranking, func_cmp)

  top = []

  while not len(top) == 3 and not heap.esta_vacia():
     _, v = heap.desencolar()
     top.append(v)

  return top




# =============================================================================
# EJERCICIO 3
# =============================================================================
"""
3. Escribir una función que reciba un grafo y determine si el mismo es bipartito, 
   o no. Indicar y justificar la complejidad de la función implementada.
"""

def es_bipartito(grafo: Grafo):
  colores = {}

  for v in grafo:
      if v not in colores:
        colores[v] = 0
        cola = Cola()
        cola.encolar(v)

        while not cola.esta_vacia():

          actual = cola.desencolar()

          for w in grafo.adyacentes(actual):

            if w not in colores:
              colores[w] = 1 - colores[actual]
              cola.encolar(w)

            elif colores[w] == colores[actual]:
              return False

  return True
