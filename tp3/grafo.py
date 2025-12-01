#!/usr/bin/python3

class Grafo:
    def __init__(self, dirigido=True, pesado=False):
        self.vertices = {}
        self.cantidad_vertices = 0
        self.cantidad_aristas = 0
        self.dirigido = dirigido
        self.pesado = pesado

    def agregar_vertice(self, vertice):
        if vertice in self.vertices:
            return
        if self.pesado:
            self.vertices[vertice] = {}
        else:
            self.vertices[vertice] = set()
        self.cantidad_vertices += 1

    def agregar_arista(self, origen, destino, peso=1):
        if origen not in self.vertices or destino not in self.vertices:
            return

        if self.pesado:
            if destino in self.vertices[origen]:
                return
            self.vertices[origen][destino] = peso
        else:
            if destino in self.vertices[origen]:
                return
            self.vertices[origen].add(destino)

        self.cantidad_aristas += 1

        if not self.dirigido and origen != destino:
            if self.pesado:
                if origen not in self.vertices[destino]:
                    self.vertices[destino][origen] = peso
            else:
                if origen not in self.vertices[destino]:
                    self.vertices[destino].add(origen)

    def estan_unidos(self, origen, destino):
        if origen not in self.vertices or destino not in self.vertices:
            return False
        return destino in self.vertices[origen]

    def vertice_pertenece(self, vertice):
        return vertice in self.vertices

    def obtener_vertices(self):
        return list(self.vertices.keys())

    def adyacentes(self, vertice):
        if vertice not in self.vertices:
            return []
        if self.pesado:
            return list(self.vertices[vertice].keys())
        else:
            return list(self.vertices[vertice])

    def grado_salida(self, vertice):
        if vertice not in self.vertices:
            return 0
        return len(self.vertices[vertice])

    def __iter__(self):
        return iter(self.vertices)
