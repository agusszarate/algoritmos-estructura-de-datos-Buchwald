import random

class Grafo:
    def __init__(self, es_dirigido=True):
        self.vertices = {}
        self.es_dirigido = es_dirigido

    def agregar_vertice(self, vertice):
        if vertice not in self.vertices:
            self.vertices[vertice] = {}

    def borrar_vertice(self, vertice):
        if vertice not in self.vertices:
            return None

        self.vertices.pop(vertice)

        for v in self.vertices:
            if vertice in self.vertices[v]:
                self.vertices[v].pop(vertice)

    def agregar_arista(self, origen, destino):
        if origen not in self.vertices or destino not in self.vertices:
            return False
        
        self._agregar_conexion(origen, destino)
        
        if not self.es_dirigido:
            self._agregar_conexion(destino, origen)       
        return True

    def borrar_arista(self, origen, destino):
        if not self.estan_unidos(origen, destino):
            return None
        
        peso = self.vertices[origen].pop(destino)
        
        if not self.es_dirigido:
            self.vertices[destino].pop(origen)
            
        return peso

    def _agregar_conexion(self, desde, hasta):
        self.vertices[desde][hasta] = 1

    def estan_unidos(self, origen, destino):
        if origen not in self.vertices: return False
        return destino in self.vertices[origen]

    def peso_arista(self, origen, destino):
        if not self.estan_unidos(origen, destino):
            return None 
        return self.vertices[origen][destino]

    def obtener_vertices(self):
        return self.vertices.keys()
    
    def vertice_pertenece(self, vertice):
        return vertice in self.vertices
    
    def vertice_aleatorio(self):
        if not self.vertices: return None
        return random.choice(list(self.vertices.keys()))

    def adyacentes(self, vertice):
        if vertice not in self.vertices: return []
        return list(self.vertices[vertice].keys())

    def __len__(self):
        return len(self.vertices)

    def __iter__(self):
        return iter(self.vertices)
    
    def __str__(self):
        return f"Grafo(Vertices={len(self)}, Dirigido={self.es_dirigido}"