-- MySQL dump 10.13  Distrib 8.0.31, for macos12 (x86_64)
--
-- Host: localhost    Database: internet_company
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP DATABASE IF EXISTS internet_company;
CREATE DATABASE internet_company;
USE internet_company;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
  `customer_id` int unsigned NOT NULL AUTO_INCREMENT,
  `dni` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `first_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `birth_date` datetime NOT NULL,
  `town` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `city` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `service` int unsigned DEFAULT NULL,
  PRIMARY KEY (`customer_id`),
  UNIQUE KEY `dni_UNIQUE` (`dni`),
  KEY `service_id_idx` (`service`),
  CONSTRAINT `service_id` FOREIGN KEY (`service`) REFERENCES `internet_services` (`service_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (1,'PEOS49','Victor','Martinez','1994-09-19 00:00:00','Benito Juarez','CDMX',1),(2,'HJ38IS','Fernando','Morales','1974-01-10 00:00:00','Cuahutemoc','CDMX',5),(3,'KOI340','Juan','Ramirez','1991-02-20 00:00:00','Cholula','Puebla',2),(4,'IU3449','Pedro','Pascal','1988-05-01 00:00:00','Santiago de Queretaro','Queretaro',5),(5,'03KSKD','Oscar','Juarez','1999-03-28 00:00:00','Pachuca','Hidalgo',1),(6,'94OKSD','Nicolas','Espinoza','1982-08-11 00:00:00','Cuernavaca','Morelos',3),(7,'JFH489','Pedro','Paramo','1976-02-18 00:00:00','Zapopan','Jalisco',4),(8,'GKFK80','Luis','Beltran','1980-12-23 00:00:00','Posarica','Veracruz',5),(9,'YKJD34','Camilo','Castillo','1996-11-10 00:00:00','Ciudad Juarez','Chihuahua',1),(10,'POWOE2','Carlos','Maldonado','1979-04-13 00:00:00','Juchitan','Oaxaca',5);
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `internet_services`
--

DROP TABLE IF EXISTS `internet_services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `internet_services` (
  `service_id` int unsigned NOT NULL AUTO_INCREMENT,
  `description` varchar(80) COLLATE utf8_unicode_ci NOT NULL,
  `megabits_rate` int NOT NULL,
  `price_per_month` decimal(4,2) NOT NULL,
  `discount` decimal(4,2) NOT NULL,
  PRIMARY KEY (`service_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `internet_services`
--

LOCK TABLES `internet_services` WRITE;
/*!40000 ALTER TABLE `internet_services` DISABLE KEYS */;
INSERT INTO `internet_services` VALUES (1,'plan dia de muertos',120,45.50,10.25),(2,'plan verano',100,49.90,9.90),(3,'plan invierno',160,50.00,0),(4,'basico',40,25.90,5.50),(5,'premium',280,99.99,9.99);
/*!40000 ALTER TABLE `internet_services` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-11-02 17:26:06
