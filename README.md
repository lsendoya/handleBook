# HandleBook: Gestor de Libros con Arquitectura Hexagonal
El repositorio HandleBook implementa un gestor de libros utilizando la arquitectura hexagonal. Este proyecto, desarrollado en Go, integra tecnologías y bibliotecas clave como JWT para
autenticación, Echo como framework web, GORM para la interacción con la base de datos, y PostgreSQL como sistema de gestión de base de datos.

## Estructura del Proyecto

### Directorio cmd:

Contiene los archivos de configuración de la base de datos, del servidor y de las variables de entorno.
main.go inicia el servidor y carga la configuración necesaria para el funcionamiento de la aplicación.

### Directorio Domain:

Incluye toda la lógica del negocio de todas las entidades. Define interfaces UseCase y Storage para abstraer las operaciones de negocio y de almacenamiento de datos, respectivamente, siguiendo los principios de la arquitectura hexagonal.

### Directorio Infrastructure:

Incluye las rutas, los handlers y los middleware de la aplicación. Contiene la implementación de la capa de almacenamiento para las entidades, integrando GORM para la interacción con PostgreSQL.

### Directorio Model:

Incluye la definición de todas las entidades con sus campos.
Define estructuras como Book que representan las entidades del dominio y sus relaciones.

## Resumen del Proyecto
HandleBook es un sistema de gestión de libros que aplica la arquitectura hexagonal para separar claramente las preocupaciones de la lógica de negocio de las interfaces externas y la infraestructura. Utiliza JWT para la autenticación, Echo como framework web, GORM para la interacción ORM con PostgreSQL, y está estructurado para facilitar la escalabilidad y el mantenimiento.

## Características Clave
### Autenticación JWT: 
Seguridad robusta mediante tokens JWT para la autenticación de usuarios.
### Framework Echo: 
Uso eficiente de Echo, un framework web de alto rendimiento para Go, facilitando la creación de APIs RESTful.
### ORM con GORM: 
Manejo eficiente de operaciones de base de datos con GORM, un ORM popular para Go.
### PostgreSQL: 
Uso de PostgreSQL, una base de datos relacional poderosa y de código abierto.
### Arquitectura Hexagonal: 
Diseño modular y desacoplado que promueve la mantenibilidad y la facilidad de prueba.
### Separación de Responsabilidades: 
La arquitectura hexagonal en HandleBook permite una clara separación entre la lógica de negocio (dominio) y las interfaces externas (infraestructura). Esto facilita la mantenibilidad y escalabilidad del proyecto.

Adaptadores y Puertos: Los adaptadores en infrastructure interactúan con los puertos definidos en domain, siguiendo los principios de la arquitectura hexagonal. Esto asegura una baja dependencia entre la lógica de negocio y las tecnologías externas.

## Cómo Contribuir
Si estás interesado en contribuir al proyecto HandleBook, puedes empezar por revisar los issues abiertos en el repositorio de GitHub, clonar el repositorio, y enviar tus pull requests. Toda contribución es bienvenida para mejorar y expandir las capacidades de este gestor de libros.
