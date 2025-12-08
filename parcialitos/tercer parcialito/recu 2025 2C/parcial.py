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

def resolver_ejercicio_2(grafo, esquinas_fuga):
    # Tu implementación para encontrar el top 3 aquí
    pass


# =============================================================================
# EJERCICIO 3
# =============================================================================
"""
3. Escribir una función que reciba un grafo y determine si el mismo es bipartito, 
   o no. Indicar y justificar la complejidad de la función implementada.
"""

def es_bipartito(grafo):
    # Tu implementación aquí
    pass