from grafo import Grafo
from cola import Cola

def bfs_caminos_minimos(grafo: Grafo, origen):
    """
    Realiza un recorrido BFS para encontrar caminos mínimos
    en un grafo no pesado (distancia = cantidad de aristas).
    
    Recibe el grafo y el vértice origen.
    
    Devuelve:
    - un diccionario de 'padres' (para reconstruir caminos)
    - un diccionario de 'distancias' (distancia mínima desde el origen)
    """

    # 1. Inicialización de estructuras
    # Vamos a necesitar diccionarios para guardar los resultados.
    distancias = {}
    padres = {}
    
    # También necesitamos un registro de vértices ya visitados.
    # Usamos un 'set' de Python, que tiene complejidad O(1)
    # para la operación 'in' (pertenencia).
    visitados = set()

    # Asumimos que el grafo es iterable (nos da todos sus vértices)
    for v in grafo:
        distancias[v] = float('inf') # Infinito
        padres[v] = None

    # 2. Seteamos el origen
    distancias[origen] = 0
    # No tiene padre
    
    cola = Cola() # Creamos la cola
    cola.encolar(origen)
    visitados.add(origen) # Marcamos el origen como visitado

    # 3. Recorrido BFS
    while not cola.esta_vacia():
        v = cola.desencolar()

        # (estamos usando la interfaz del TDA Grafo)
        for w in grafo.obtener_adyacentes(v):
            
            # Si 'w' no fue visitado, significa que estamos
            # llegando a él por primera vez (y por ende, por
            # el camino más corto desde 'origen').
            if w not in visitados:
                # Lo descubrimos
                visitados.add(w)
                padres[w] = v
                distancias[w] = distancias[v] + 1
                cola.encolar(w)

    # Devolvemos los diccionarios con los resultados
    return padres, distancias