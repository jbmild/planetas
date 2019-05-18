# Documentacion

La api, posee 4 servicios principales

1. http://[URL]/
	* Mensaje de bienvenida al sistema
	> {"message":"Bienvenido al pronóstico de la universo. ¿Como podemos ayudarte?"}
1. http://[URL]/clima?dia=[NRO]
	* Devuelve un json indicando el dia y el clima esperado
	> {"clima":"Sequía","dia":1}
1. http://[URL]/procesar10anios
	* Inicia el proceso que calcula el clima de los proximos 10 años y los guarda en la base de datos. 
	* Si ya fue calculado, no hace nada. 
	* Este proceso se corre en un thread separado
	> {"message":"La tarea se inicio correctamente."}
1. http://[URL]/vaciarbasededatos
	* Vacia la base de datos para que el job que calcula los proximos 10 años pueda volverse a correr.
	> {"message":"La información se borro correctamente."}

# Posibles mejoras

Estas son algunas ideas mediante las cuales se podria mejorar el codigo y funcionalidades.

* Crear archivo de configuracion y unficiar la conexion a la base de datos
* Mejorar la recuperacion ante fallos y reusar codigo
* Cambiar el numero de dia, por la fecha
* Permitir el ingreso de configuracion ideal como por ejemplo: vector de inicio de planetas, fecha de inicio de promostico, como se fraccionan los dias (para mejora de calculo)
* Cuando se inicia el job que calcula los 10 años, si el mismo ya fue ejecutado informar que lo fue.
* Usar el patron factory para crear los planetas
* Usar un ORM para el guardado del pronostico (en este momento no se justifica por la dimesion pero para que sea escalable)
