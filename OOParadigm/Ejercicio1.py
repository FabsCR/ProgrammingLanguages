# Ejercicio 1 - Paradigma OO
# Fabian Jose Fernandez Fernandez

class ObjetoRepresentable:
    def __init__(self, nombre):
        """
        Constructor de la clase ObjetoRepresentable.

        Parámetros:
        - nombre: str, nombre del objeto representable.
        """
        self._nombre = nombre

    def get_nombre(self):
        """
        Obtener el nombre del objeto representable.

        Retorna:
        - str, nombre del objeto representable.
        """
        return self._nombre

    def representar(self):
        """
        Método abstracto para representar el objeto.
        Debe ser implementado por las clases hijas.

        Retorna:
        - str, representación del objeto.
        """
        pass

class Texto(ObjetoRepresentable):
    def __init__(self, nombre, contenido):
        """
        Constructor de la clase Texto.

        Parámetros:
        - nombre: str, nombre del texto.
        - contenido: str, contenido del texto.
        """
        super().__init__(nombre)
        self._contenido = contenido

    def get_contenido(self):
        """
        Obtener el contenido del texto.

        Retorna:
        - str, contenido del texto.
        """
        return self._contenido

    def representar(self):
        """
        Representa el objeto Texto.

        Retorna:
        - str, representación del objeto Texto.
        """
        return f"Texto: {self.get_nombre()}, Contenido: {self.get_contenido()}"

class ObjetoGeometrico(ObjetoRepresentable):
    def __init__(self, nombre, tipo):
        """
        Constructor de la clase ObjetoGeometrico.

        Parámetros:
        - nombre: str, nombre del objeto geométrico.
        - tipo: str, tipo de objeto geométrico.
        """
        super().__init__(nombre)
        self._tipo = tipo

    def get_tipo(self):
        """
        Obtener el tipo del objeto geométrico.

        Retorna:
        - str, tipo del objeto geométrico.
        """
        return self._tipo

    def representar(self):
        """
        Representa el objeto geométrico.

        Retorna:
        - str, representación del objeto geométrico.
        """
        return f"Objeto Geométrico: {self.get_nombre()}, Tipo: {self.get_tipo()}"

class Grupo(ObjetoRepresentable):
    def __init__(self, nombre):
        """
        Constructor de la clase Grupo.

        Parámetros:
        - nombre: str, nombre del grupo.
        """
        super().__init__(nombre)
        self._elementos = []

    def agregar_elemento(self, elemento):
        """
        Agrega un elemento al grupo.

        Parámetros:
        - elemento: ObjetoRepresentable, objeto representable a agregar al grupo.
        """
        self._elementos.append(elemento)

    def get_elementos(self):
        """
        Obtiene la lista de elementos del grupo.

        Retorna:
        - list, lista de objetos representables en el grupo.
        """
        return self._elementos

    def representar(self):
        """
        Representa el objeto Grupo y sus elementos.

        Retorna:
        - str, representación del objeto Grupo.
        """
        representacion = f"Grupo: {self.get_nombre()}, Elementos:\n"
        for elemento in self.get_elementos():
            representacion += f"  - {elemento.representar()}\n"
        return representacion

if __name__ == "__main__":
    # Crear instancias de objetos representables
    texto_intro = Texto("Introducción", "Bienvenido al Editor Gráfico de Documentos.")
    circulo_rojo = ObjetoGeometrico("Círculo Rojo", "Círculo")
    rectangulo_azul = ObjetoGeometrico("Rectángulo Azul", "Rectángulo")

    # Crear un grupo para organizar elementos
    grupo_figuras = Grupo("Grupo de Figuras")

    # Agregar objetos al grupo de figuras
    grupo_figuras.agregar_elemento(circulo_rojo)
    grupo_figuras.agregar_elemento(rectangulo_azul)

    # Crear otro grupo anidado
    grupo_anidado = Grupo("Grupo Anidado")
    grupo_anidado.agregar_elemento(Texto("Nota", "Este es un grupo anidado."))

    # Agregar el grupo anidado al grupo de figuras
    grupo_figuras.agregar_elemento(grupo_anidado)

    # Crear el documento principal y agregar elementos
    documento_principal = Grupo("Documento Principal")
    documento_principal.agregar_elemento(texto_intro)
    documento_principal.agregar_elemento(grupo_figuras)

    # Representar el documento principal
    print("Representación del Documento Principal:")
    print(documento_principal.representar())
