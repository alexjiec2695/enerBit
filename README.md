
## Tabla de contenido 

* [Pre-requisitos ](#Pre-requisitos)
* [Instalación](#Instalación)


### Pre-requisitos 📋

  * Para poder utilizar este aplicativo es necesario instalar [Go.](https://golang.org/doc/install)
  
  * Instalar WIRE de forma global. Wire es una herramienta de generación de código que automatiza la conexión de componentes mediante la inyección de dependencias [Wire.](https://github.com/google/wire)

```
go get github.com/google/wire/cmd/wire
```

### Instalación

* Clonar el repositorio

````
git clone https://github.com/alexjiec2695/enerBit.git
````

* Debemos de navegar en nuestra aplicación utilizando la consola del equipo hasta donde se encuentra ubicado el archivo `Wire.go`

```
cd applications/rest_app/di

wire
```

* Con esta configuración completa estamos listos para invocar el comando `Wire` tener en cuenta que este comando genera el archivo `wire_gen.go`

* Para la ejecución del proyecto tenemos que navegar a la carpeta de la aplicación, _ejemplo_ `cd applications/rest_app` y correr el siguiente comando en la consola

```
go run .
```

## Construido con 🛠️

* [Go](https://golang.org/) - Lenguaje de programación base del proyecto Falcon. 
* [Gin ](https://github.com/gin-gonic/gin) - Librería web usada para la definición de los endpoints REST.
* [Wire](https://github.com/google/wire) - Librería de manejo de Inyección de dependencias.
* [Testify](https://github.com/stretchr/testify) - Librería que permite realizar pruebas unitarias.
* [Viper](https://github.com/spf13/viper) - Librería que sirve para la lectura de archivos de configuración de tipo JSON, YAML, TOML, entre otros.
