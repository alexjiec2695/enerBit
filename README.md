
## Tabla de contenido 

* [Pre-requisitos ](#Pre-requisitos)
* [Instalaci贸n](#Instalaci贸n)


### Pre-requisitos 馃搵

  * Para poder utilizar este aplicativo es necesario instalar [Go.](https://golang.org/doc/install)
  
  * Instalar WIRE de forma global. Wire es una herramienta de generaci贸n de c贸digo que automatiza la conexi贸n de componentes mediante la inyecci贸n de dependencias [Wire.](https://github.com/google/wire)

```
go get github.com/google/wire/cmd/wire
```

### Instalaci贸n

* Clonar el repositorio

````
git clone https://github.com/alexjiec2695/enerBit.git
````

* Debemos de navegar en nuestra aplicaci贸n utilizando la consola del equipo hasta donde se encuentra ubicado el archivo `Wire.go`

```
cd applications/rest_app/di

wire
```

* Con esta configuraci贸n completa estamos listos para invocar el comando `Wire` tener en cuenta que este comando genera el archivo `wire_gen.go`

* Para la ejecuci贸n del proyecto tenemos que navegar a la carpeta de la aplicaci贸n, _ejemplo_ `cd applications/rest_app` y correr el siguiente comando en la consola

```
go run .
```

## Construido con 馃洜锔?

* [Go](https://golang.org/) - Lenguaje de programaci贸n base del proyecto Falcon. 
* [Gin ](https://github.com/gin-gonic/gin) - Librer铆a web usada para la definici贸n de los endpoints REST.
* [Wire](https://github.com/google/wire) - Librer铆a de manejo de Inyecci贸n de dependencias.
* [Testify](https://github.com/stretchr/testify) - Librer铆a que permite realizar pruebas unitarias.
* [Viper](https://github.com/spf13/viper) - Librer铆a que sirve para la lectura de archivos de configuraci贸n de tipo JSON, YAML, TOML, entre otros.
