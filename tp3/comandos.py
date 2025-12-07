#!/usr/bin/python3
import biblioteca

# --- Función Auxiliar ---
def imprimir_resultado_camino(camino):
    if camino is None:
        print("No se encontro recorrido")
    else:
        print(" -> ".join(camino))
        print(f"Costo: {len(camino) - 1}")

# --- Comandos del TP ---
def listar_operaciones():

    print("camino")
    print("rango")
    print("navegacion")
    print("clustering")
    print("lectura")

def camino(grafo, params):
    if len(params) != 2:
        return

    origen, destino = params[0], params[1]
    resultado = biblioteca.camino_minimo(grafo, origen, destino)
    imprimir_resultado_camino(resultado)

def rango(grafo, params):
    if len(params) != 2:
        return

    pagina = params[0]
    try:
        n = int(params[1])
    except ValueError:
        return

    # Llamamos a la función que cuenta nodos a distancia exacta n
    cantidad = biblioteca.cantidad_en_rango(grafo, pagina, n)
    print(cantidad)

def navegacion(grafo, params):
    if len(params) != 1:
        return

    origen = params[0]
    resultado = biblioteca.navegacion_primer_link(grafo, origen)

    # El formato de salida es simplemente la lista unida por flechas
    # Si solo hay 1 elemento (el origen sin links), se imprime solo.
    print(" -> ".join(resultado))

def clustering(grafo, params):
    # Caso 1: Promedio de toda la red
    if len(params) == 0:
        c = biblioteca.clustering_promedio(grafo)
        print(f"{c:.3f}")
    
    # Caso 2: Clustering de una página específica
    elif len(params) == 1:
        pagina = params[0]
        # Si la página no existe, por definición devolvemos 0.000 (o validamos antes)
        if not grafo.vertice_pertenece(pagina):
            print("0.000")
            return
            
        c = biblioteca.clustering_individual(grafo, pagina)
        print(f"{c:.3f}")

def lectura(grafo, params):
    # params ya es una lista de páginas ["p1", "p2", ...]
    if not params:
        return

    orden = biblioteca.orden_topologico(grafo, params)

    if orden is None:
        print("No existe forma de leer las paginas en orden")
    else:
        print(", ".join(orden))

# --- Diccionario de despacho ---

COMANDOS = {
    "listar_operaciones": listar_operaciones,
    "camino": camino,
    "rango": rango,
    "navegacion": navegacion,
    "clustering": clustering,
    "lectura": lectura,
}

def ejecutar_comando(grafo, comando, params):
    if comando not in COMANDOS:
        return

    # Caso especial: listar_operaciones no recibe parámetros
    if comando == "listar_operaciones":
        COMANDOS[comando]()
    else:
        COMANDOS[comando](grafo, params)