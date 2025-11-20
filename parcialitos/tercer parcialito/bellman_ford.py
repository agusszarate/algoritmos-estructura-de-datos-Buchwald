from grafo import Grafo

def bellman_ford(grafo: Grafo, origen):
    """
    Calcula los caminos mínimos desde 'origen' usando Bellman-Ford.
    Soporta aristas con peso negativo.
    
    Devuelve:
    - un diccionario de 'padres'
    - un diccionario de 'distancias'
    - un booleano 'True' si detecta un ciclo negativo, 'False' si no.
    """
    
    # O(cant_v)
    distancias = {}
    padres = {}
    aristas = []
    cant_v = 0

    for v in grafo:
        distancias[v] = float('inf')
        padres[v] = None
        cant_v += 1

        for w in grafo.obtener_adyacentes(v):
            peso = grafo.peso_arista(v, w)
            aristas.append((v, w, peso)) # Tupla (origen, destino, peso)
    
    # Bellman-Ford itera sobre aristas, no vértices.
    # Esta operación es O(cant_v + E)
    
    # Seteamos el origen
    distancias[origen] = 0
    
    # Iteramos cant_v-1 veces. (range(cant_v) también funciona, 
    # cant_v-1 es suficiente si no hay ciclos neg).
    # Este bucle externo corre cant_v veces.
    for i in range(cant_v): 
        
        # El bucle interno itera sobre TODAS las E aristas
        # Costo interno: O(E)
        hubo_cambios = False # Optimización: si no hay cambios, podemos parar
        for u, v, peso in aristas:
            
            # Condición de Relajación:
            # Si la distancia a 'u' no es infinita Y
            # el camino a 'u' + el peso de (u,v) es mejor
            # que el camino que teníamos a 'v'...
            if distancias[u] != float('inf') and distancias[u] + peso < distancias[v]:
                # ...actualizamos 'v'
                distancias[v] = distancias[u] + peso
                padres[v] = u
                hubo_cambios = True
        
        # Optimización (corte temprano)
        if not hubo_cambios and i < cant_v-1:
            # Si en una pasada completa no relajamos ninguna arista,
            # ya encontramos todos los caminos mínimos.
            break

    # Hacemos UNA pasada más (la cant_v-ésima) sobre todas las aristas
    # Costo: O(E)
    for u, v, peso in aristas:
        if distancias[u] != float('inf') and distancias[u] + peso < distancias[v]:
            # Si AÚN podemos relajar una arista, significa que
            # estamos en un ciclo negativo.
            # No hay solución de camino mínimo (es -infinito).
            return None, None, True # True indica ciclo negativo

    # Si pasamos la cant_v-ésima iteración sin relajar, no hay ciclos
    # negativos alcanzables desde el origen.
    return padres, distancias, False # False indica no hay ciclo