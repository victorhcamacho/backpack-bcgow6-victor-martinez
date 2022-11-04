
### Respuestas a las preguntas del ejercicio 2

a. ¿Cuál es la primary key para la tabla de clientes?.
> La PK de la tabla customers debe ser el identificador (customer_id), debido a que debe ser unico para cada registro y un medio para diferenciarlos con otros registros que poseean datos similares.

b. ¿Cuál es la primary key para la tabla de planes de internet?.
> La PK de la tabla internet services debe ser el identificador (service_id), debido a que debe ser unico para cada registro y un medio para diferenciarlos con otros registros que poseean datos similares.

c. ¿Cómo serían las relaciones entre tablas? ¿En qué tabla debería haber foreign key? ¿A qué campo de qué tabla hace referencia dicha foreign key?
> La relacion entre las tablas debe ser N(customers) - 1(services), un customer puede tener unicamente un service y un service puede tener muchos clientes. El FK debe estar en la tabla de customers y hacer referencia al service id de la tabla services.

