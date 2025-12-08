import random

class Grafo:
    def __init__(self, es_dirigido=True, es_pesado=False):
        self.vertices = {}
        self.es_dirigido = es_dirigido
        self.es_pesado = es_pesado

    def agregar_vertice(self, vertice):
        if vertice not in self.vertices:
            self.vertices[vertice] = {} if self.es_pesado else set()

    def agregar_arista(self, origen, destino, peso=1):
        if origen not in self.vertices or destino not in self.vertices:
            return False
        peso_real = peso if self.es_pesado else 1
        self._agregar_conexion(origen, destino, peso_real)
        if not self.es_dirigido:
            self._agregar_conexion(destino, origen, peso_real)       
        return True

    def _agregar_conexion(self, desde, hasta, peso):
        if self.es_pesado:
            self.vertices[desde][hasta] = peso
        else:
            self.vertices[desde].add(hasta)

    def estan_unidos(self, origen, destino):
        if origen not in self.vertices: return False
        return destino in self.vertices[origen]

    def peso_arista(self, origen, destino):
        if not self.estan_unidos(origen, destino):
            return None 
        
        if self.es_pesado:
            return self.vertices[origen][destino]
        else:
            return 1 

    def obtener_vertices(self):
        return self.vertices.keys()
    
    def vertice_pertenece(self, vertice):
        return vertice in self.vertices
    
    def vertice_aleatorio(self):
        if not self.vertices: return None
        return random.choice(list(self.vertices.keys()))

    def adyacentes(self, vertice):
        if vertice not in self.vertices: return []
        
        if self.es_pesado:
            return list(self.vertices[vertice].keys())
        return list(self.vertices[vertice])

    def __len__(self):
        return len(self.vertices)

    def __iter__(self):
        return iter(self.vertices)
    
    def __str__(self):
        return f"Grafo(Vertices={len(self)}, Dirigido={self.es_dirigido}, Pesado={self.es_pesado})"