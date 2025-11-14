from grafo import GrafoDirigido
from cola import Cola

# Implementar un algoritmo que reciba un grafo dirigido y acíclico y determine si dicho grafo admite un único orden topológico. 
# Indicar y justificar la complejidad de la función. Pista: pensar qué condición puede darse para que exista más de un posible orden topológico.

def ejercicio1(grafo: GrafoDirigido):

    grados = {}
    
    for v in grafo:
        grados[v] = 0

        for w in grafo.obtener_adyacentes(v):
            if w not in grados:
                grados[w] = 1
            else:
                grados[w] += 1
    cola = Cola()



def unico_orden_topologico(grafo: GrafoDirigido, grados, cola):
    
    for g in grados.keys():
        if grados[g] == 0:
            cola.encolar(g)
            if len(cola) > 1:
                return False
        
    item = cola.desencolar()
    for w in grafo.obtener_adyacentes(item):
        grados[w] -= 1
    
    return unico_orden_topologico(grafo, grados, cola)



# Se tiene una app de tránsito que indica el mejor camino desde un punto A a un punto B. Esta app trabaja con un grafo dirigido de las calles de la Ciudad Autónoma de Buenos Aires, en donde las esquinas son los vértices del Grafo, y las aristas son las calles que unen dichas esquinas. Las aristas tienen como peso: el tiempo (en segundos) que demora en recorrerse, y su nombre (por ejemplo, "Paseo Colón 800-900"). Además, hay un informe respecto al estado de las calles que indica para cada uno 3 posibles estados:

# HABILITADA: Se puede transitar de forma normal.

# CONGESTIONADA: Se tarda el doble en transitar.

# CORTADA: Imposible transitar. Se pide implementar un algoritmo que reciba el grafo, el informe (un diccionario en el que las claves son los nombres de las aristas, y como valores su estado correspondiente), un punto A y un punto B, y determine la forma más rápida para llegar desde A hasta B, considerando el estado dado por el informe. Indicar y justificar la complejidad del algoritmo implementado.

# Implementar un algoritmo que reciba un grafo no dirigido y determine la cantidad máxima de aristas que se pueden agregar al mismo de tal forma que no se reduzcan la cantidad de componentes conexas que hay en el mismo. Indicar y justificar la complejidad del algoritmo implementado.

# Implementar un algoritmo que, dado un grafo pesado (con pesos positivos), un vértice v y otro w, determine la cantidad de caminos mínimos que hay de v a w dentro del grafo. Considerar que, justamente, podrían haber varios caminos de una misma distancia, que a su vez sean la mínima. Indicar y justificar la complejidad del algoritmo implementado. Por ejemplo, en el grafo de abajo hay 3 caminos mínimos del vértice A al vértice E: A -> E, A -> B -> F -> D -> E, A -> H -> F -> D -> E, todos de costo 8.