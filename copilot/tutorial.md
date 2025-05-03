
Tener en cuenta que en general se van agregando nuevas funcionalidades a la herramienta, por lo que es recomendable revisar la documentacion oficial de copilot para estar al tanto de las novedades.   




Este es un conjunto de tips/tutoriales para el uso de copilot.

Modelo.



1. Autocomplete

Como primer paso, asegurense que copilot este configurado para poder usar el autocomplete. 
En el editor hacer click en el icono de copilot.

![Copilot icon](copilot.png)

Este el primer feature que tuvo github copilto cuando salio en beta en 2021,  probablemente uno de mis favoritos y el que mas utilize a diario.


Basicamente, a medida que vamos tipeado en el editor, copilot va sugiriendo lineas de codigo,  teniedo en cuenta los archivos 
que estan abiertos en el editor, preferencias de convenciones y teniendo en cuenta el estilo de codigo que el proyecto tiene.
como todo no es algo infalible y a veces no sugiere lo que uno espera, pero en general esta bastante cerca de lo que uno 
espera en general.

Por ejemplo en este ejemplo,  nosotros tenemos un metodo llamado "Register" que se encargar de registrar un usuario en una base de datos,
si nosotros empezamos a escribir un nuevo metodo, este caso login vamos a ver que copilot nos sugiere un bloque de codigo tomando en cuenta 
las practicas y convenciones que tiene el proyecto, como vemos en la imagen.

![Copilot icon](autocomplete01.png)

Tambien se puede ver en la imagen que tenemos la posiblidad de ver otras sugerencias, esto va a aparecer al hacer hover con el cursor del mouse sobre el bloque de codigo sugerido y en caso de querer ir mas paso a paso, poder ir aceptando las sugerencias de a poco ( esto pude ser util si copilot sugiere un codigo muy largo y demasiado complejo, o si no queremos aceptar todo de una vez).

2. Editor inline
Otro feature muy util es el editor inline,  el cual nos permite poder decirle a copilot que intente soluciar un error de sintaxis o de compilacion en una linea en particuular.

Por ejemplo veamos este ejemplo.

![Copilot icon](inline01.png)

En este caso tenemos un error de complicion/transpilacion en typescript,  al hacer click en esea linea, a la izquierda va a aparecer un icono, el cual al hacer click nos va a permitir
indicarle a copilot que intente solucionarlo o al menos nos comente porque aparece el error.

![Copilot icon](inline02.png)

Aca tenemos varias opciones, podemos aceptar el cambio,  volver a pedirle que lo itente,  probar con otro modelo de LLM en caso de no encontrar una solucion, etc.

Otro caso de uso del editor inline es poder seleccinar un bloque de codigo y pedirle a copilot un cambio en especifico, 
por ejemplo en este caso le vamos a pedir a copilot un ajuste en una funcion de Registro.

![Copilot icon](inline03.png)

Este el codigo generado 
![Copilot icon](inline04.png)














3. Usar el context select para modificar codigo.

2. Modo Ask 

3. Modo edit 

4. Modo agente

5. Modo reglas.