# Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD” y 
# guardar en la misma los episodios de todas las temporadas de “The Walking Dead”.

CREATE TEMPORARY TABLE `TWD` 
SELECT e.* FROM episodes e
INNER JOIN seasons s
ON s.serie_id = (SELECT id FROM series WHERE title LIKE "%The Walking Dead%")
AND e.season_id = s.id;

SELECT * FROM TWD;

# En la base de datos “movies”, seleccionar una tabla donde crear un índice y luego chequear la creación del mismo. 

ALTER TABLE series
ADD INDEX idx_title (title);

SHOW INDEX FROM series;