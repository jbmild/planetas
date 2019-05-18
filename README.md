# Documentación

La api, posee 4 servicios principales

1. http://[URL]/
	* Mensaje de bienvenida al sistema
	> {"message":"Bienvenido al pronóstico del universo. ¿Como podemos ayudarte?"}
1. http://[URL]/clima?dia=[NRO]
	* Devuelve un json indicando el dia y el clima esperado
	> {"clima":"Sequía","dia":1}
1. http://[URL]/procesar10anios
	* Inicia el proceso que calcula el clima de los próximos 10 años y los guarda en la base de datos.
	* Si ya fue calculado, no hace nada.
	* Este proceso se corre en un thread separado
	> {"message":"La tarea se inicio correctamente."}
1. http://[URL]/vaciarbasededatos
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
* Usar un ORM para el guardar el pronostico (en este momento no se justifica por la dimensión pero para que sea escalable)