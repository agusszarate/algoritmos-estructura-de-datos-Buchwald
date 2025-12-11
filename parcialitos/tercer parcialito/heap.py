class Heap:
    """Cola de prioridad implementada como heap"""
    
    def __init__(self, funcion_cmp):
        """Crea un heap vacío con la función de comparación dada"""
        raise NotImplementedError("no implementado")
    
    def esta_vacia(self):
        """Verifica si el heap está vacío"""
        raise NotImplementedError("no implementado")
    
    def encolar(self, elemento):
        """Encola un elemento en el heap"""
        raise NotImplementedError("no implementado")
    
    def ver_max(self):
        """Devuelve el elemento con mayor prioridad sin desencolarlo"""
        raise NotImplementedError("no implementado")
    
    def desencolar(self):
        """Desencola y devuelve el elemento con mayor prioridad"""
        raise NotImplementedError("no implementado")
    
    def cantidad(self):
        """Devuelve la cantidad de elementos en el heap"""
        raise NotImplementedError("no implementado")
    
    def __len__(self):
        """Devuelve la cantidad de elementos en el heap"""
        return self.cantidad()


def crear_heap(funcion_cmp):
    """Crea un heap vacío con la función de comparación dada"""
    raise NotImplementedError("no implementado")


def crear_heap_desde_lista(lista, funcion_cmp) -> Heap:
    """Crea un heap a partir de una lista con la función de comparación dada"""
    raise NotImplementedError("no implementado")


def heap_sort(elementos, funcion_cmp) -> list:
    """Ordena una lista usando heap sort"""
    raise NotImplementedError("no implementado")
