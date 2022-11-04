
# Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.

SELECT concat(e.nombres," ",e.apellidos) AS "nombre completo", e.puesto, d.localidad 
FROM empleados e 
INNER JOIN departamentos d 
ON e.depto_id = d.id;


# Visualizar los departamentos con más de cinco empleados.

SELECT d.id, d.nombre, d.localidad FROM departamentos d
INNER JOIN empleados e
ON e.depto_id = d.id
GROUP BY d.id
HAVING COUNT(d.id) > 2;


# Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.

SELECT concat(e.nombres," ",e.apellidos) AS "nombre completo", e.salario, d.nombre AS "departamento"
FROM empleados e 
INNER JOIN departamentos d 
ON e.depto_id = d.id 
AND e.puesto = (select puesto from empleados WHERE nombres = "Mito" AND apellidos = "Barchuk");


# Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.

SELECT e.* FROM empleados e
INNER JOIN departamentos d
ON e.depto_id = d.id
AND e.depto_id = (select id from departamentos where nombre = "Contabilidad")
ORDER BY e.nombres;


# Mostrar el nombre del empleado que tiene el salario más bajo.

SELECT concat(nombres," ",apellidos) AS "nombre completo" FROM empleados 
WHERE salario = (select min(salario) from empleados);


# Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.

SELECT * FROM empleados 
WHERE salario = 
(SELECT max(e.salario) AS "salario mas alto" FROM empleados e
INNER JOIN departamentos d
ON e.depto_id = d.id
AND d.nombre = "Ventas");