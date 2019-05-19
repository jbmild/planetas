# Instalación

Para poder correr el código localmente, es necesario tener instalado go, git y las siguientes dependencias:
> go get -u github.com/go-sql-driver/mysql

> go get -u github.com/gin-gonic/gin

# Ejecución

## Parte base

La parte principal del ejercicio, en el cual se corre mediante consola tiene como punto de ingreso el archivo 10yearanalysis.go que se encuentra en la raíz del proyecto. Dicho archivo, al ejecutarse generara una salida con la información solicitada en el enunciado.
> go run [ARCHIVO 10yearanalysis.go]

## Bonus

El bonus solicita generar una base de datos con un job que calcule el clima para los próximos 10 años, exponerlo en una api y subirlo a un servicio de hosteo. Esta ultima parte, de subirlo, no esta realizada sin embargo puede ejecutarse localmente.
Para ejecutarlo localmente, es necesario tener adicionalmente una base de datos mysql y luego seguir estos pasos:
* Crear un schema llamado "planets"
* Ejecutar el script sql ubicado en [RAIZ]/migrations/forecast.sql en el schema recién creado
* Modificar conexión con la base de datos.
	* En las linea 47 y 81 se encuentran dos conexiones con la base de datos. En las mismas hay que reemplazar el usuario, contraseña, host y puerto de ser necesario según la configuración local
	>db := Utils.GetConnection("[USUARIO]", "[CONTRASEÑA]", "[HOST]", "[PUERTO]", "planets")

Una vez realizado esto, puede ejecutarse el servicio:
> go run [ARCHIVO api.go]


# Documentación

La api, posee 4 servicios principales

1. http://localhost:8080/
	* Mensaje de bienvenida al sistema
	> {"message":"Bienvenido al pronóstico del universo. ¿Como podemos ayudarte?"}
1. http://localhost:8080/clima?dia=[NRO]
	* Devuelve un json indicando el dia y el clima esperado
	> {"clima":"Sequía","dia":1}
1. http://localhost:8080/procesar10anios
	* Inicia el proceso que calcula el clima de los próximos 10 años y los guarda en la base de datos.
	* Si ya fue calculado, no hace nada.
	* Este proceso se corre en un thread separado
	> {"message":"La tarea se inicio correctamente."}
1. http://localhost:8080/vaciarbasededatos
	* Vacía la base de datos para que el job que calcula los próximos 10 años pueda volverse a correr.
	> {"message":"La información se borro correctamente."}

Dentro del repositorio, se encuentra un archivo llamado "10yearanalysis.go", dicho archivo, puede utilizarse para ejecutar la primera parte de este ejercicio desde la consola.

# Posibles mejoras

Estas son algunas ideas mediante las cuales se podría mejorar el código y funcionalidades.

* Crear archivo de configuración e iniciar la conexión a la base de datos
* Mejorar la recuperación ante fallos y re usar código
* Cambiar el numero de día, por la fecha
* Permitir el ingreso de configuración ideal como por ejemplo: vector de inicio de planetas, fecha de inicio de pronóstico, como se fraccionan los días (para mejora de cálculos)
* Cuando se inicia el job que calcula los 10 años, si el mismo ya fue ejecutado informar que lo fue.
* Usar el patrón factory para crear los planetas
* Implementar el patron strategy para calcular los pronosticos
* Usar un ORM para el guardar el pronostico (en este momento no se justifica por la dimensión pero para que sea escalable)
* Subirlo al hosting
