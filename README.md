# Dogs API con Fiber, GORM & Docker (MySQL)
Creacion de API con Fiber, Gorm y utilizo Docker para usar MySQL en mi proyecto.
Este proyecto fue creado con el fin de aprender a utilizar Gorm y Docker de forma básica y rápida.

| Métodos | URL | Funciones |
| -------- | -------- | --------- |
| `POST` |  /dog | CreateDog |
| `GET` |  /dogs | GetDogs |
| `GET` |  /get/`:id` | GetDog |
| `PUT` |  /update/`:id` | UpdateDog |
| `POST` | /delete/`:id`| DeleteDog |

Cada endpoint retorna un `JSON`
