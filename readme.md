# Proyecto de Votaciones API

Este proyecto es una API RESTful construida con Go y GORM para gestionar un sistema de votaciones.

## Tabla de Contenidos

- [Instalación](#instalación)
- [Uso](#uso)
- [Endpoints](#endpoints)
- [Contribuir](#contribuir)
- [Licencia](#licencia)

## Instalación

Para instalar el proyecto, sigue estos pasos:

1. Clona el repositorio: `git clone HTTP`
2. Navega al directorio del proyecto: `cd tu_repositorio`
3. Instala las dependencias: `go get`
4. Ejecuta el servidor: `go run main.go`

## Uso

Para usar la API, puedes hacer peticiones a los siguientes endpoints:

- `GET /usuarios`: Obtiene todos los usuarios.
- `POST /usuarios`: Crea un nuevo usuario.
- `PUT /usuarios/:id`: Actualiza un usuario existente.
- `POST /usuarios/login`: Autentica a un usuario.

## Endpoints

### Usuarios

- `GET /usuarios`: Obtiene todos los usuarios.
- `POST /usuarios`: Crea un nuevo usuario.
- `PUT /usuarios/:id`: Actualiza un usuario existente.
- `POST /usuarios/login`: Autentica a un usuario.

## Contribuir

¡Contribuciones son bienvenidas! Si quieres contribuir a este proyecto, por favor sigue estos pasos:

1. Haz un fork del proyecto
2. Crea una nueva rama (`git checkout -b feature/feature_name`)
3. Haz commit de tus cambios (`git commit -am 'Añade una nueva función'`)
4. Sube la rama (`git push origin feature/feature_name`)
5. Abre un pull request

## Licencia

Este proyecto está bajo la Licencia [Nombre de la Licencia]. Para más detalles, lee el archivo [LICENSE](LICENSE).
