class Cola:
    """Cola enlazada"""
    
    def __init__(self):
        raise NotImplementedError("no implementado")
    
    def encolar(self, elemento):
        """Encola un elemento"""
        raise NotImplementedError("no implementado")
    
    def desencolar(self):
        """Desencola y devuelve el primer elemento"""
        raise NotImplementedError("no implementado")
    
    def ver_primero(self):
        """Devuelve el primer elemento sin desencolarlo"""
        raise NotImplementedError("no implementado")
    
    def esta_vacia(self) -> bool:
        """Verifica si la cola está vacía"""
        return True
    
    def __len__(self):
        """Devuelve la cantidad de elementos en la cola"""
        raise NotImplementedError("no implementado")
