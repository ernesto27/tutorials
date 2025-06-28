# Tutorial: Claude Code y Container Use

Este es un tutorial en el cual vamos a ver un ejemplo paso a paso de cómo utiliza Claude Code en conjunto con contenedores (container-use), para poder ejecutar tareas en background que se estén ejecutando en un contenedor en un nuevo branch en git.

- Setup Claude Code
- Crear proyecto API Go
- Uso de Claude Code 
- Instalación de Container Use
- Ejemplo de tarea en background container

## Setup Claude Code

Claude es una herramienta que nos permite interactuar con un agente IA, usando los modelos de Anthropic Claude (Sonnet 4, Sonnet 3.7, etc) desde la terminal. De esta manera no dependemos de utilizar un editor en particular (VSCode, Cursor, Windsurf) aunque de todas maneras se puede conectar Claude Code a VSCode, Cursor en caso de que lo necesitemos.

### Requisitos Previos

Como primer paso necesitamos tener Node.js instalado en nuestra máquina y ejecutar este comando:

```bash
npm install -g @anthropic-ai/claude-code
```

### Configuración Inicial

Una vez instalado,  Claude Code debe ser ejecutado dentro de la carpeta donde se encuentre nuestro proyecto, para esto vamos a crear una carpeta de esta manera:

```bash
mkdir claude-code-example && cd claude-code-example
```

Ejecutamos Claude:

```bash
claude
```

Nos va a pedir permisos, presionamos en **Yes**.

### Opciones de Autenticación

Claude Code nos va a mostrar estas dos opciones:

1. **Suscripción en Claude**
2. **Anthropic Account Console**

En caso de tener una cuenta en Claude ya estariamos listos al elegir la primera opcion. 
De otro, tienen que seleccionar la segunda opción, en la cual Claude se va a encargar de configurar un acceso ApiKey automaticamente en nuestra cuenta.

### Configuración Manual de API Key (Opcional)

En caso de tener un ApiKey ya generado previamente, existe la posibilidad de utilizarlo haciendo lo siguiente:

1. Ir a la siguiente carpeta:
   ```bash
   cd ~/.claude
   ```

2. Crear o modificar el archivo `settings.json` con lo siguiente:
   ```json
   {
     "apiKeyHelper": "~/.claude/anthropic_key.sh"
   }
   ```

3. Crear el script:
   ```bash
   nano anthropic_key.sh
   ```

4. Agregar la API Key:
   ```bash
   echo "sk-..."
   ```

Una vez configurado nuestro ApiKey o cuenta, podemos acceder a Claude.


![Claude Code Setup](image1.png)

### Crear proyecto API go.

Para probar claude-code en conjunto con container-use ,  vamos a crear un proyecto de una API en Go.
https://go.dev/

En la carpeta creada anteriormente, claude-code-example, agregamos estos dos archivos.

#### go.mod

```go
module goapi

go 1.23
```

#### main.go

```go
package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello World"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
```

Ejecutar el servicio de esta manera:

```bash
go run main.go
```

Probamos que este todo Ok usando Curl via terminal:

```bash
curl http://localhost:8080/hello
```   

Inicializar git (necesario para usar container-use)
```bash
git init
git add .
git commit -m "Initial commit"
```


## Uso de Claude Code  

Antes de arrancar con el uso de claude code, debemos crear un archivo de configuracion CLAUDE.md para indicarle cuales son las instrucciones que debe seguir Claude en nuestras interacciones con el proyecto.

Dentro de claude, ejecutar lo siguiente.

```sh 
/init
```

Durante la ejecución de este comando, Claude les va a pedir permiso para ejecutar diferentes herramientas.
![](image2.png)

Archivo generado por Claude 

https://github.com/ernesto27/tutorials/commit/975140d21d0f116a68dfea28c0e55feae95bdc57


Vamos a ver un ejemplo de como seria el uso de Claude Code desde la terminal, realizando cambios en nuestro branch actual.

Le pedimos claude que nos genere un endpoint que retorne la version actual de la API.

```bash
claude "Crea un endpoint que retorne la version  de la API, la version se define en un archivo llamado version.json"
```

A medida que se va generado el codigo, claude nos va mostrando los cambios que va realizando y la opcion de aceptar o modificar los cambios.

Cambios generados por claude:

https://github.com/ernesto27/tutorials/commit/5f80e12ea9aad4a51fece7a34de5db72c921ba4f


Este sería un flujo de uso básico de Claude Code, en donde vamos viendo los cambios de manera incremental y podemos aceptar o modificar los cambios que nos propone en el momento, lo cual sería algo similar a lo que se puede hacer con Copilot, Cursor o Windsurf pero desde la terminal.

En el proximo paso vamos a ver como podemos usar Claude Code para ejecutar tareas en background.


## Claude code con Container use.

Al momento si quisiéramos usar Claude Code para ejecutar alguna otra tarea fuera del branch en el que estamos trabajando, podríamos tener otra copia de proyecto en otra ruta y ejecutar Claude o ejecutar en un servidor o VM.  
Pero de que manera podriamos delegar una tarea a Claude Code en nuestar maquina sin que afecte a nuestro branch o carpeta acutal en la que necesitemos trabajar en otras tareas?

Para esto podemos utilizar contenedores, los cuales nos permiten tener un entorno aislado para ejecutar tareas, probar ideas, librerías, refactors, etc.  
vamos a utilizar un proyecto llamado container-use,  el cual esta desarrollado por el equipo creador de Docker y posteriormente de Dagger.

https://github.com/dagger/container-use

### Instalación de Container Use

Ejecutar el siguiente comando 
```bash
curl -fsSL https://raw.githubusercontent.com/dagger/container-use/main/install.sh | bash
```

Una vez instalado cu,  tenemos que configurar claude para que utilice el servicio MCP de cu.

https://docs.anthropic.com/en/docs/mcp

en la carpeta del proyecto ejecutamos el siguiente comando:

```bash
claude mcp add container-use -- cu stdio
```

Agregar reglas de container use a CLAUDE.md 

```bash
curl https://raw.githubusercontent.com/dagger/container-use/main/rules/agent.md >> CLAUDE.md
```

### Crear tarea en background con container use

Una vez configurado container-user vamos a probar lo siguiente:

- Crear tarea en background con container use, claude code 
- Revisar cambios realizados por claude code
- Merger los cambios al branch principal


Vamos a pedirle a claude que genere tests para los endpoints de nuestra API y ademas agregue una carpeta para los handlers/controladores de la API.

```bash
claude "Crea tests para los endpoints de la API y agrega una carpeta llamada controllers para los handlers de los endpoints,  usa container-use MCP"
```

Claude va a empezar una nueva sesión, con la diferencia de que esta vez va a utilizar el MCP cu (container-use) para ejecutar los cambios en un contenedor.
A medida que va ir avanzando, nos va a pedir permismo para ejecutar diferentes tools.
![](image3.png)

Selecciones la segunda opcion "Yes, and don't ask again for container-use:environment_create",  de esta menera en las proximas tareas ya no nos va a aparecer este mensaje.

Podemos ver cuales son los branchs creados por container user, con el siguient comando:

```bash 
cu list 

```
```bash 
ID                TITLE                           CREATED         UPDATED
devoted-squirrel  Go API Testing and Refactoring  5 minutes ago   11 seconds ago
```
Podemos ver los cambios realizados via terminal

```bash
cu logs 
```

Debido a que container-use genere un nuevo branch,  tambien podemos ver los cambios haciendo un checkout al branch correspondiente:

![](image4.png)

A medida que se va generado el codigo, claude nos va mostrando los cambios que va realizando y la opcion de aceptar o modificar los cambios,  aca debemos elegir la segunda opcion para que no vuelva a preguntar en las proximas tareas.

```bash

Una vez que estamos conformes con los cambios podemos hacer un merge al branch principal,  de esta manera:

```bash
cu merge devoted-squirrel main
```

Cambios generados 


https://github.com/ernesto27/tutorials/commit/75bb0b1c7323a5ef236d668a4dfb44093728ec6e

## Resumen

En este tutorial hemos aprendido a combinar Claude Code con Container Use para ejecutar tareas en paralelo y un entorno aislado sin afectar nuestro branch principal. A continuación, los puntos clave:

- Configuración inicial de Claude Code, incluyendo autenticación y uso de `claude` desde la terminal.
- Creación de una API en Go y uso de Claude Code para agregar endpoints y pruebas de manera incremental.
- Instalación y configuración de Container Use (`cu`) para delegar tareas en contenedores aislados.
- Ejecución de tareas en background con `cu`, incluyendo gestión de branches, visualización de logs y merges.
- Beneficios de esta integración: entorno reproducible, aislamiento de cambios y flujo de trabajo paralelo.



















