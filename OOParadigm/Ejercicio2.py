# Ejercicio 2 - Paradigma OO
# Fabian Jose Fernandez Fernandez

from datetime import datetime

# Definición de la clase Libro
class Libro:
    def __init__(self, codigo, titulo, autor):
        """
        Constructor de la clase Libro.

        Parameters:
        - codigo (str): Código único del libro.
        - titulo (str): Título del libro.
        - autor (str): Autor del libro.
        """
        self.codigo = codigo
        self.titulo = titulo
        self.autor = autor
        self.disponible = True

    def prestar(self):
        """
        Método para prestar el libro.

        Returns:
        - bool: True si el libro se prestó correctamente, False si no está disponible.
        """
        if self.disponible:
            self.disponible = False
            return True
        return False

    def devolver(self):
        """
        Método para devolver el libro.
        """
        if not self.disponible:
            self.disponible = True

    def esta_disponible(self):
        """
        Método para verificar si el libro está disponible.

        Returns:
        - bool: True si el libro está disponible, False si no lo está.
        """
        return self.disponible

    def obtener_info(self):
        """
        Método para obtener la información del libro.

        Returns:
        - str: Información formateada del libro.
        """
        return f"Código: {self.codigo}, Título: {self.titulo}, Autor: {self.autor}, Disponible: {self.disponible}"

# Definición de la clase Socio
class Socio:
    MAX_PRESTAMOS = 3

    def __init__(self, numero, nombre, direccion):
        """
        Constructor de la clase Socio.

        Parameters:
        - numero (int): Número único del socio.
        - nombre (str): Nombre del socio.
        - direccion (str): Dirección del socio.
        """
        self.numero = numero
        self.nombre = nombre
        self.direccion = direccion
        self.prestamos = []

    def tomar_prestado(self, libro):
        """
        Método para que el socio tome prestado un libro.

        Parameters:
        - libro (Libro): Libro que el socio desea tomar prestado.

        Returns:
        - bool: True si el libro se prestó correctamente, False si no se puede prestar.
        """
        if len(self.prestamos) < self.MAX_PRESTAMOS and libro.prestar():
            self.prestamos.append(libro)
            return True
        return False

    def devolver_libro(self, libro):
        """
        Método para que el socio devuelva un libro prestado.

        Parameters:
        - libro (Libro): Libro que el socio desea devolver.
        """
        if libro in self.prestamos:
            libro.devolver()
            self.prestamos.remove(libro)
        else:
            print(f"Error: El libro {libro.titulo} no está en los préstamos de {self.nombre}.")

    def libros_prestados(self):
        """
        Método para obtener la lista de libros prestados al socio.

        Returns:
        - list: Lista de libros prestados al socio.
        """
        return self.prestamos

    def tiene_mas_de_3_prestamos(self):
        """
        Método para verificar si el socio tiene más de 3 libros prestados.

        Returns:
        - bool: True si tiene más de 3 libros prestados, False si no.
        """
        return len(self.prestamos) > self.MAX_PRESTAMOS

    def obtener_info(self):
        """
        Método para obtener la información del socio.

        Returns:
        - str: Información formateada del socio.
        """
        return f"Número de Socio: {self.numero}, Nombre: {self.nombre}, Dirección: {self.direccion}"

# Crear instancias de libros con los nuevos títulos
libro1 = Libro("M111", "El Gran Gatsby", "F. Scott Fitzgerald")
libro2 = Libro("M222", "1984", "George Orwell")
libro3 = Libro("M333", "Matar un ruiseñor", "Harper Lee")
libro4 = Libro("CRC1", "Única mirando al mar", "Fernando Contreras Castro")
libro5 = Libro("CRC2", "Cocori", "Joaquín Gutiérrez")
libro6 = Libro("CRC3", "En una silla de ruedas", "Carmen Lyra")
libro7 = Libro("CRC4", "La isla de los hombres solos", "José León Sánchez")
libro8 = Libro("CRC5", "Cuentos de angustias y paisajes", "Carlos Salazar Herrera")
libro9 = Libro("ABC1", "Don Quijote de la Mancha", "Miguel de Cervantes")
libro10 = Libro("ABC2", "Cien años de soledad", "Gabriel García Márquez")
libro11 = Libro("ABC3", "El Señor de los Anillos", "J.R.R. Tolkien")

# Crear instancias de socios
socio1 = Socio(1, "Abel Mendez", "Calle Blancos")
socio2 = Socio(2, "María Salazar", "Avenida 54")
socio3 = Socio(3, "Leonardo Viquez", "Diagonal 200")
socio4 = Socio(4, "Esteban Ballestero", "Plaza 100")
socio5 = Socio(5, "Laura González", "Calle Principal")

# Realizar préstamos con los nuevos libros
socio1.tomar_prestado(libro4)
socio1.tomar_prestado(libro5)
socio2.tomar_prestado(libro6)
socio3.tomar_prestado(libro7)
socio4.tomar_prestado(libro8)
socio1.tomar_prestado(libro9)
socio2.tomar_prestado(libro10)
socio3.tomar_prestado(libro11)
socio4.tomar_prestado(libro9)
socio4.tomar_prestado(libro10)

# Hacer que el socio1 devuelva el libro4
socio1.devolver_libro(libro4)

# Realizar préstamos adicionales con los nuevos libros al socio3
socio3.tomar_prestado(libro1)
socio3.tomar_prestado(libro2)
socio3.tomar_prestado(libro3)

# Mostrar información actualizada de libros y socios
print("\nInformación actualizada de libros:")
for libro in [libro1, libro2, libro3, libro4, libro5, libro6, libro7, libro8, libro9, libro10, libro11]:
    print(libro.obtener_info())

print("\nLibros prestados a", socio4.nombre + ":")
for libro in socio4.libros_prestados():
    print(libro.obtener_info())

# Mostrar socios con más de 3 libros prestados
socios_con_mas_de_3_prestamos = filter(lambda s: s.tiene_mas_de_3_prestamos(), [socio1, socio2, socio3, socio4, socio5])
print("\nSocios con más de 3 libros prestados:")
for socio in socios_con_mas_de_3_prestamos:
    print(socio.obtener_info())