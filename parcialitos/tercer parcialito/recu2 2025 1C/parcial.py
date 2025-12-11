from grafo import Grafo
from cola import Cola
# =============================================================================
# EJERCICIO 1: Rociadores e Insecticidas
# =============================================================================
"""
1. En nuestra huerta tenemos unos rociadores de insecticidas automáticos. Cada rociador 
   cuenta con la dosis apropiada para cubrir hasta un máximo de 5 plantaciones a su alrededor. 
   Es necesario averiguar si algún rociador tiene más de 5 plantaciones alrededor ya que de 
   tener una mayor cantidad la dosis sería insuficiente. Se tiene un grafo en donde los vértices 
   son los rociadores y plantas, es no pesado y dirigido (el origen de una arista es el rociador 
   y el destino es una planta en su rango).
   
   Implementar una función que reciba este grafo y devuelva, en caso que un rociador tenga más 
   de 5 plantaciones alrededor, el conjunto de plantaciones alrededor de dicho rociador (si hay 
   más de un rociador que cumpla esta condición, devolver la información correspondiente a 
   cualquiera de estos). En caso contrario, devolver None. 
   
   Indicar y justificar la complejidad de la función.
"""
def obtener_grados(grafo: Grafo):
    grados = {}

    for v in grafo:
        grados[v] = 0

    for v in grafo:
        for w in grafo.obtener_adyacentes(v):
            grados[w] += 1

    return grados

def verificar_rociadores(grafo: Grafo):

    grados = obtener_grados(grafo) #O(V + E)

    for v in grados.keys():
        if grados[v] == 0:
            contador = 0
            plantas = []
            for w in grafo.obtener_adyacentes(v):
                contador += 1
                plantas.append(w)
            
            if len(plantas) > 5: 
                return plantas

            


# =============================================================================
# EJERCICIO 2: AlgoConnect (AlgoFriends vs AlgoBuy)
# =============================================================================
"""
2. La exitosa empresa AlgoConnect está desarrollando dos aplicaciones:
   
   - AlgoFriends: Es una red social en donde los usuarios pueden subir imágenes y entablar 
     relaciones. Cualquier usuario puede conectarse con cualquier otro.
   - AlgoBuy: Es una plataforma de ventas, en donde cada usuario se registra como comprador 
     o vendedor. El único motivo de conexión posible entre usuarios es hacer consultas sobre 
     un producto, es por esto que los compradores sólo pueden interactuar con vendedores, 
     y viceversa. 
     
   Cada aplicación cuenta con su propio grafo, en ambos casos los grafos son no dirigidos, 
   no pesados, en donde los vértices son los usuarios y las aristas representan que dos 
   usuarios se han conectado.

   Lamentablemente hubo un error en el software que borró las etiquetas de estos grafos, por 
   lo que debemos re-etiquetarlos para saber si cada grafo corresponde a AlgoFriends o a AlgoBuy. 
   
   Implementar una función que dado DOS GRAFOS nos devuelva dos Strings que correspondan a la 
   asignación de cada uno de los grafos. Debe devolver:
    - "AlgoFriends", "AlgoBuy" si el primero correspondería al de AlgoFriends y el segundo al de AlgoBuy;
    - "AlgoBuy", "AlgoFriends", si es lo contrario;
    - None, None si dado cómo son los grafos, no es posible determinar cuál aplicación es cada uno. 
   
   Indicar y justificar la complejidad de la función.
"""

def es_bipartito(grafo: Grafo):
    codigos = {}

    for v in grafo:
        if v not in codigos:
            codigos[v] = 0
            cola = Cola()
            cola.encolar(v)

            while not cola.esta_vacia():
                actual = cola.desencolar()
                
                for w in grafo.obtener_adyacentes(actual):
                    if w not in codigos:
                        codigos[w] = 1 - codigos[actual]
                        cola.encolar(w)
                    elif codigos[w] == codigos[actual]:
                        return False
    
    return True


def identificar_aplicaciones(grafo_1, grafo_2):
    bipartito1 = es_bipartito(grafo_1)
    bipartito2 = es_bipartito(grafo_2)

    if bipartito1 == bipartito2:
        return None, None
    elif bipartito1:
        return "AlgoBuy", "AlgoFriends"
    else: 
        return "AlgoFriends", "AlgoBuy"


# =============================================================================
# EJERCICIO 3: Dijkstra
# =============================================================================
"""
3. a. Hacer un seguimiento del Algoritmo de Dijkstra sobre el grafo del dorso para llegar 
      desde el vértice A hacia todos los demás vértices del grafo. Asegurarse de mostrar 
      los pasos del seguimientos y las estructuras internas que se utilizan.
   
   b. Si tuviéramos que cambiar el peso de la arista de B a D, tal que dicha arista no sea 
      utilizada en ningún camino mínimo desde A, ¿cuál es el valor más bajo que se le debería 
      asignar? Justifique brevemente.
"""

# Datos del grafo (transcritos de la imagen)
grafo_dijkstra = {
    'G': {'A': 2},
    'A': {'B': 5, 'F': 10},
    'B': {'E': 15, 'F': 4, 'D': 1, 'C': 8},
    'C': {'E': 4, 'D': 4},
    'D': {'F': 2, 'C': 5},
    'E': {'A': 10, 'B': 3},
    'F': {} 
}

def seguimiento_dijkstra(grafo, origen):
    # Escribe tu código para el seguimiento aquí
    pass

# Respuestas teóricas (puedes escribirlas como comentarios):
# 3a. (Ver output de la función)
# 3b. Respuesta justificada: ...