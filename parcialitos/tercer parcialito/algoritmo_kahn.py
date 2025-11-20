from grafo import GrafoDirigido
from cola import Cola

def algoritmo_kahn(grafo: GrafoDirigido):
    """
    Calcula un orden topológico usando el algoritmo de Kahn.
    """
    
    # 1. Calcular grados de entrada (usando tu función)
    grados = obtenerGrados(grafo)

    # 2. Inicializar la cola con todos los nodos de grado 0
    cola = Cola()        
    for v in grados.keys():
        if (grados[v] == 0):
            cola.encolar(v)
    
    # Lista para guardar el resultado
    orden_topologico = Cola()
    
    # Contador para detectar ciclos
    vertices_procesados = 0

    # 3. Procesar la cola
    while not cola.esta_vacia():
        # Desencolamos un vértice y lo agregamos al orden
        u = cola.desencolar()
        orden_topologico.encolar(u)
        vertices_procesados += 1
        
        # Recorremos sus adyacentes
        for v in grafo.obtener_adyacentes(u):
            # "Removemos" la arista (u, v)
            grados[v] -= 1

            # Si el vecino 'v' quedó con grado 0, a la cola
            if (grados[v] == 0):
                cola.encolar(v)

    # 4. Detección de ciclos
    # Si procesamos menos vértices que el total, había un ciclo
    if vertices_procesados != len(grados):
        # Puedes manejar esto como prefieras:
        # return None
        # return "El grafo tiene un ciclo"
        raise ValueError("El grafo contiene un ciclo, no se puede ordenar topológicamente.")
        
    return orden_topologico

def obtenerGrados(grafo: GrafoDirigido):
    """
    Calcula el grado de entrada de cada vértice del grafo.
    """
    grados = {}

    # Inicializa todos los vértices con grado 0
    for v in grafo:
        grados[v] = 0

    # Recorre las aristas para sumar los grados de entrada
    for v in grafo:
        for w in grafo.obtener_adyacentes(v):
            grados[w] += 1

    return grados