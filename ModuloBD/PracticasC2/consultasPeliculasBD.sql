SELECT * FROM movies;


# Agregar una película a la tabla movies.

INSERT INTO `movies`
(`title`,
`rating`,
`awards`,
`release_date`,
`length`,
`genre_id`)
VALUES
("Avengers", 9.5, 2, NOW(), 130, 5);


# Agregar un género a la tabla genres.

INSERT INTO `genres`
(`name`,
`ranking`,
`active`)
VALUES
("Super Heroes",13,1);

# Asociar a la película del punto 1. con el género creado en el punto 2.

UPDATE `movies` SET genre_id = 14 WHERE id = 22;

# Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.

UPDATE `actors` SET favorite_movie_id = 22 WHERE id=1;

# Crear una tabla temporal copia de la tabla movies.

CREATE TEMPORARY TABLE `temp_movies` SELECT * FROM `movies`;

# Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

SET SQL_SAFE_UPDATES = 0;
DELETE FROM `temp_movies` WHERE awards < 5;

# Obtener la lista de todos los géneros que tengan al menos una película.

SELECT g.name FROM `genres` AS g
INNER JOIN `movies` AS m
ON g.id = m.genre_id
GROUP BY g.name
HAVING COUNT(m.id) >= 1;


# Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.

SELECT a.* FROM `actors` a 
INNER JOIN `movies` m 
ON m.id = a.favorite_movie_id
WHERE m.awards > 3;

# Crear un índice sobre el nombre en la tabla movies.

ALTER TABLE `movies`
ADD INDEX idx_title (title);

# Chequee que el índice fue creado correctamente.

SHOW INDEX FROM `movies`;

# En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.
-- R= Si, por lo menos en las tablas donde se crean los indices, el rendimiento sera mejor debido a la cantidad de consultas
-- que se ejecutan en este script de sql y la forma en que estan estrucutradas dichas sentencias

# ¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
-- R= Creario otro indice en la tabla de actors debido a que es una tabla bastante concurrente en las sentencias