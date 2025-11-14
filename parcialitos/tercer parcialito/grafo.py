class Grafo:
    """Grafo no dirigido"""
    
    def __init__(self):
        raise NotImplementedError("no implementado")
    
    def agregar_vertice(self, clave, valor):
        """Agrega un vértice al grafo"""
        raise NotImplementedError("no implementado")
    
    def borrar_vertice(self, clave):
        """Elimina un vértice del grafo"""
        raise NotImplementedError("no implementado")
    
    def agregar_arista(self, origen, destino, peso=1):
        """Agrega una arista entre dos vértices"""
        raise NotImplementedError("no implementado")
    
    def borrar_arista(self, origen, destino):
        """Elimina una arista entre dos vértices"""
        raise NotImplementedError("no implementado")
    
    def existe_vertice(self, clave):
        """Verifica si existe un vértice"""
        raise NotImplementedError("no implementado")
    
    def existe_arista(self, origen, destino):
        """Verifica si existe una arista entre dos vértices"""
        raise NotImplementedError("no implementado")
    
    def obtener_vertice(self, clave):
        """Obtiene el valor de un vértice"""
        raise NotImplementedError("no implementado")
    
    def modificar_vertice(self, clave, valor):
        """Modifica el valor de un vértice"""
        raise NotImplementedError("no implementado")
    
    def obtener_vertices(self):
        """Obtiene todos los vértices del grafo"""
        raise NotImplementedError("no implementado")
    
    def obtener_adyacentes(self, clave):
        """Obtiene los vértices adyacentes a un vértice"""
        raise NotImplementedError("no implementado")
    
    def cantidad(self):
        """Devuelve la cantidad de vértices"""
        raise NotImplementedError("no implementado")
    
    def __iter__(self):
        """Itera sobre todos los vértices del grafo"""
        raise NotImplementedError("no implementado")
    
    def __len__(self):
        """Devuelve la cantidad de vértices"""
        return self.cantidad()


class GrafoDirigido:
    """Grafo dirigido"""
    
    def __init__(self):
        raise NotImplementedError("no implementado")
    
    def agregar_vertice(self, clave, valor):
        """Agrega un vértice al grafo"""
        raise NotImplementedError("no implementado")
    
    def borrar_vertice(self, clave):
        """Elimina un vértice del grafo"""
        raise NotImplementedError("no implementado")
    
    def agregar_arista(self, origen, destino, peso=1):
        """Agrega una arista entre dos vértices"""
        raise NotImplementedError("no implementado")
    
    def borrar_arista(self, origen, destino):
        """Elimina una arista entre dos vértices"""
        raise NotImplementedError("no implementado")
    
    def existe_vertice(self, clave):
        """Verifica si existe un vértice"""
        raise NotImplementedError("no implementado")
    
    def existe_arista(self, origen, destino):
        """Verifica si existe una arista entre dos vértices"""
        raise NotImplementedError("no implementado")
    
    def obtener_vertice(self, clave):
        """Obtiene el valor de un vértice"""
        raise NotImplementedError("no implementado")
    
    def modificar_vertice(self, clave, valor):
        """Modifica el valor de un vértice"""
        raise NotImplementedError("no implementado")
    
    def obtener_vertices(self):
        """Obtiene todos los vértices del grafo"""
        raise NotImplementedError("no implementado")
    
    def obtener_adyacentes(self, clave):
        """Obtiene los vértices adyacentes a un vértice"""
        raise NotImplementedError("no implementado")
    
    def cantidad(self):
        """Devuelve la cantidad de vértices"""
        raise NotImplementedError("no implementado")
    
    def __iter__(self):
        """Itera sobre todos los vértices del grafo"""
        raise NotImplementedError("no implementado")
    
    def __len__(self):
        """Devuelve la cantidad de vértices"""
        return self.cantidad()
