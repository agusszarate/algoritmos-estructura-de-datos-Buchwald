from heap import Heap

def dijkstra(grafo, origen, destino=None):
    """
    Calcula los caminos mínimos desde 'origen' usando Dijkstra,
    utilizando el TDA Heap (Max-Heap genérico) provisto.
    """

    # 1. Inicialización de estructuras
    distancias = {}
    padres = {}
    
    for v in grafo:
        distancias[v] = float('inf')
        padres[v] = None
        
    distancias[origen] = 0
    

    def cmp_min_heap(a, b):
        # a y b son tuplas (distancia, vertice)
        # Queremos que el de menor distancia (a[0])
        # tenga "mayor" prioridad.
        return b[0] - a[0] 

    # Bárbara crea el TDA Heap y le pasa la función
    heap = Heap(cmp_min_heap)
    
    # Bárbara encola la tupla (distancia, vertice)
    heap.encolar((0, origen)) 
    
    # --- FIN DEL CAMBIO ---
    
    visitados = set()
    
    # 2. Recorrido
    
    # Cambiamos 'while heap' por la primitiva de Alan
    while not heap.esta_vacia(): 
        
        dist_v, v = heap.desencolar() 

        if v in visitados:
            continue
            
        visitados.add(v)
        
        if v == destino:
            break
            
        for w in grafo.obtener_adyacentes(v):
            
            if w in visitados:
                continue
                
            peso_vw = grafo.peso_arista(v, w)
            dist_via_v = distancias[v] + peso_vw
            
            if dist_via_v < distancias[w]:
                distancias[w] = dist_via_v
                padres[w] = v 
                
                heap.encolar((distancias[w], w))

    return padres, distancias