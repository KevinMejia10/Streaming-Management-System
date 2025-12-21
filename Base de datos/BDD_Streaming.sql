-- MySQL dump 10.13  Distrib 8.0.42, for Win64 (x86_64)
--
-- Host: localhost    Database: zenithpruebas
-- ------------------------------------------------------
-- Server version	9.3.0

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

--
-- Table structure for table `contenidos`
--

DROP TABLE IF EXISTS `contenidos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `contenidos` (
  `CONTENIDO_ID` bigint NOT NULL COMMENT 'Se registra ID del contenido (pelicula, seria, documental)',
  `TITULO` varchar(50) DEFAULT NULL COMMENT 'Se ingresa el titulo del contenido (serie, pelicula, documental)',
  `DESCRIPCION` varchar(200) DEFAULT NULL COMMENT 'Se ingresa la descripcion de la pelicula, serie, documental. ',
  `CLASIFICACION_EDAD` int DEFAULT NULL,
  `ES_ESTRENO` tinyint(1) DEFAULT NULL COMMENT 'Se define si el contenido esta en estreno y disponible para comprar',
  `PRECIO_COMPRA` int DEFAULT NULL COMMENT 'Se ingresa el precio de compra del coontenido',
  `PERFILES_PERFIL_ID` bigint NOT NULL,
  PRIMARY KEY (`CONTENIDO_ID`),
  KEY `CONTENIDOS_PERFILES_FK` (`PERFILES_PERFIL_ID`),
  CONSTRAINT `CONTENIDOS_PERFILES_FK` FOREIGN KEY (`PERFILES_PERFIL_ID`) REFERENCES `perfiles` (`PERFIL_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Tabla para gestionar los contenidos (películas, series, documentales)';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contenidos`
--

LOCK TABLES `contenidos` WRITE;
/*!40000 ALTER TABLE `contenidos` DISABLE KEYS */;
INSERT INTO `contenidos` VALUES (201,'El Viaje de Chihiro','Una niña se adentra en un mundo mágico habitado por espíritus, dioses y monstruos.',7,1,6,101),(202,'Stranger Things','Cuando un niño desaparece, sus amigos, la familia y la policía se ven envueltos en un misterio extraordinario en Hawkins, Indiana.',16,0,NULL,103),(203,'Planeta Tierra II','Documental que explora la vida salvaje en los ecosistemas más icónicos del planeta, narrado por David Attenborough.',0,1,7,101),(204,'La La Land: Ciudad de Sueños','Un pianista de jazz y una aspirante a actriz se enamoran en Los Ángeles mientras persiguen sus sueños.',13,0,NULL,104),(205,'Mad Max: Furia en la Carretera','En un futuro post-apocalíptico desértico, una mujer se rebela contra un tirano con la ayuda de un vagabundo solitario.',18,1,10,105),(206,'Coco','Un joven músico viaja a la Tierra de los Muertos para descubrir la verdadera historia de su familia y un secreto ancestral.',5,0,NULL,102),(207,'Parásitos','Una familia humilde se infiltra de manera progresiva en el sofisticado hogar de una familia adinerada, con consecuencias imprevisibles.',13,0,NULL,106),(208,'Blade Runner 2049','Un joven blade runner descubre un secreto enterrado hace mucho tiempo que podría sumir lo que queda de la sociedad en el caos.',16,1,6,107),(209,'Queen: Live at Wembley Stadium','Grabación de uno de los conciertos más icónicos de la legendaria banda británica Queen, con Freddie Mercury a la cabeza.',0,0,NULL,108),(210,'CantaJuego: ¡A divertirse!','Contenido musical y educativo diseñado para niños pequeños, con canciones pegadizas y coreografías divertidas para aprender jugando.',18,0,NULL,109);
/*!40000 ALTER TABLE `contenidos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `historial_visualizaciones`
--

DROP TABLE IF EXISTS `historial_visualizaciones`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `historial_visualizaciones` (
  `HISTORIAL_ID` bigint NOT NULL COMMENT 'Se registra el ID de historial',
  `PERFIL_ID` bigint DEFAULT NULL COMMENT 'Se registra el ID del perfil del usuario',
  `CONTENIDO_ID` bigint DEFAULT NULL COMMENT 'Se registra el ID del contenido (pelicula, serie, documental)',
  `TIEMPO_VISUALIZADO_SEGUNDOS` bigint DEFAULT NULL COMMENT 'Tiempo de visualizacion del contenido',
  `ULTIMA_VISUALIZACION` date DEFAULT NULL COMMENT 'Se registra la ultima fecha de visualizacion del contenido',
  `COMPLETADO` tinyint(1) DEFAULT NULL COMMENT 'Validacion para determinar si el contenido se visualizacion por completo',
  `PERFILES_PERFIL_ID` bigint NOT NULL,
  `CONTENIDOS_CONTENIDO_ID` bigint NOT NULL,
  PRIMARY KEY (`HISTORIAL_ID`),
  KEY `HISTORIAL_VISUALIZACIONES_CONTENIDOS_FK` (`CONTENIDOS_CONTENIDO_ID`),
  KEY `HISTORIAL_VISUALIZACIONES_PERFILES_FK` (`PERFILES_PERFIL_ID`),
  CONSTRAINT `HISTORIAL_VISUALIZACIONES_CONTENIDOS_FK` FOREIGN KEY (`CONTENIDOS_CONTENIDO_ID`) REFERENCES `contenidos` (`CONTENIDO_ID`),
  CONSTRAINT `HISTORIAL_VISUALIZACIONES_PERFILES_FK` FOREIGN KEY (`PERFILES_PERFIL_ID`) REFERENCES `perfiles` (`PERFIL_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Tabla para registrar el historial de visualizaciones';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `historial_visualizaciones`
--

LOCK TABLES `historial_visualizaciones` WRITE;
/*!40000 ALTER TABLE `historial_visualizaciones` DISABLE KEYS */;
INSERT INTO `historial_visualizaciones` VALUES (401,101,201,7200,'2025-01-15',1,101,201),(402,103,202,3600,'2025-03-10',0,103,202),(403,101,203,1800,'2025-04-05',0,101,203),(404,104,204,5400,'2025-05-20',1,104,204),(405,105,205,6000,'2025-01-25',0,105,205),(406,102,206,1200,'2025-03-22',1,102,206),(407,106,207,4800,'2025-04-18',0,106,207),(408,107,208,6500,'2025-05-01',1,107,208),(409,108,209,2400,'2025-01-01',0,108,209),(410,109,210,900,'2025-05-30',1,109,210);
/*!40000 ALTER TABLE `historial_visualizaciones` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `metodos_pago`
--

DROP TABLE IF EXISTS `metodos_pago`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `metodos_pago` (
  `METODO_PAGO_ID` bigint NOT NULL COMMENT 'Se registra el ID del metodo de pago',
  `USUARIO_ID` bigint DEFAULT NULL COMMENT 'Se registra el ID del usuario',
  `TIPO_PAGO` varchar(30) DEFAULT NULL COMMENT 'Se registra el tipo de pago que uso el usuario (tarjeta debito/credito, transferencia)',
  `ES_PREDETERMINADO` tinyint(1) DEFAULT NULL COMMENT 'Se valida si el pago esta registrado como predeterminado',
  `USUARIOS_USUARIO_ID` bigint NOT NULL,
  PRIMARY KEY (`METODO_PAGO_ID`),
  KEY `METODOS_PAGO_USUARIOS_FK` (`USUARIOS_USUARIO_ID`),
  CONSTRAINT `METODOS_PAGO_USUARIOS_FK` FOREIGN KEY (`USUARIOS_USUARIO_ID`) REFERENCES `usuarios` (`USUARIO_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Tabla para gestionar los métodos de pago';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `metodos_pago`
--

LOCK TABLES `metodos_pago` WRITE;
/*!40000 ALTER TABLE `metodos_pago` DISABLE KEYS */;
INSERT INTO `metodos_pago` VALUES (1,1,'Tarjeta Credito',1,1),(2,2,'PayPal',0,2),(3,3,'Transferencia',1,3);
/*!40000 ALTER TABLE `metodos_pago` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `perfiles`
--

DROP TABLE IF EXISTS `perfiles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `perfiles` (
  `PERFIL_ID` bigint NOT NULL COMMENT 'Registro del ID del perfil',
  `USUARIO_ID` bigint DEFAULT NULL COMMENT 'Se registra el ID del Usuario',
  `NOMBRE_PERFIL` varchar(10) DEFAULT NULL COMMENT 'Registro del nombre del perfil',
  `CLASIFICACION_EDAD_MAXIMA` decimal(28,0) DEFAULT NULL COMMENT 'Se define la clasificación para establecer control parental',
  `USUARIOS_USUARIO_ID` bigint NOT NULL,
  PRIMARY KEY (`PERFIL_ID`),
  KEY `PERFILES_USUARIOS_FK` (`USUARIOS_USUARIO_ID`),
  CONSTRAINT `PERFILES_USUARIOS_FK` FOREIGN KEY (`USUARIOS_USUARIO_ID`) REFERENCES `usuarios` (`USUARIO_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Tabla para gestionar los perfiles de usuario';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `perfiles`
--

LOCK TABLES `perfiles` WRITE;
/*!40000 ALTER TABLE `perfiles` DISABLE KEYS */;
INSERT INTO `perfiles` VALUES (101,1,'Principal',99,1),(102,1,'Niños',7,1),(103,2,'Mi Perfil',18,2),(104,3,'UsuarioA',13,3),(105,4,'UsuarioB',16,4),(106,5,'Main',99,5),(107,6,'Adolesc.',14,6),(108,7,'Secundario',99,7),(109,8,'Infantil',5,8),(110,9,'Premium',99,9);
/*!40000 ALTER TABLE `perfiles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `planes_suscripcion`
--

DROP TABLE IF EXISTS `planes_suscripcion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `planes_suscripcion` (
  `SUSCRIPCION_ID` bigint NOT NULL COMMENT 'Se registra el ID de la suscripción',
  `USUARIO_ID` bigint DEFAULT NULL,
  `METODO_PAGO_ID` bigint DEFAULT NULL,
  `TIPO_PLAN` varchar(30) DEFAULT NULL COMMENT 'Se registra el tipo de plan que contrata el usuario',
  `FECHA_INICIO` date DEFAULT NULL COMMENT 'Registro de la fecha de inicio del plan',
  `FECHA_FIN` date DEFAULT NULL COMMENT 'Registro fecha de finalizacion del plan',
  `ESTADO_SUSCRIPCION` varchar(10) DEFAULT NULL COMMENT 'Registro del estado de la suscripción (activa, cancelada, expirada)',
  `PRECIO_MENSUAL_ANUAL` decimal(10,2) DEFAULT NULL COMMENT 'Se ingresa el precio mensual o anual de la suscripcion',
  `FECHA_PROXIMO_PAGO` date DEFAULT NULL COMMENT 'Se define la fecha del proximo pago en funcion del precio anual o mensual',
  `METODOS_PAGO_METODO_PAGO_ID` bigint NOT NULL,
  `USUARIOS_USUARIO_ID` bigint NOT NULL,
  PRIMARY KEY (`SUSCRIPCION_ID`),
  KEY `PLANES_SUSCRIPCION_METODOS_PAGO_FK` (`METODOS_PAGO_METODO_PAGO_ID`),
  KEY `PLANES_SUSCRIPCION_USUARIOS_FK` (`USUARIOS_USUARIO_ID`),
  CONSTRAINT `PLANES_SUSCRIPCION_METODOS_PAGO_FK` FOREIGN KEY (`METODOS_PAGO_METODO_PAGO_ID`) REFERENCES `metodos_pago` (`METODO_PAGO_ID`),
  CONSTRAINT `PLANES_SUSCRIPCION_USUARIOS_FK` FOREIGN KEY (`USUARIOS_USUARIO_ID`) REFERENCES `usuarios` (`USUARIO_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Tabla para gestionar los planes de suscripción';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `planes_suscripcion`
--

LOCK TABLES `planes_suscripcion` WRITE;
/*!40000 ALTER TABLE `planes_suscripcion` DISABLE KEYS */;
INSERT INTO `planes_suscripcion` VALUES (301,1,1,'Premium Mensual','2024-05-01','2024-06-01','Activa',10.99,'2024-06-01',1,1),(302,2,2,'Básico Anual','2023-10-10','2024-10-10','Activa',99.99,'2024-10-10',2,2),(303,3,3,'Estandar Mensual','2024-04-15','2024-05-15','Expirada',7.99,'2024-05-15',3,3);
/*!40000 ALTER TABLE `planes_suscripcion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transacciones`
--

DROP TABLE IF EXISTS `transacciones`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transacciones` (
  `TRANSACCION_ID` bigint NOT NULL COMMENT 'Registro del ID de la transaccion',
  `USUARIO_ID` bigint DEFAULT NULL,
  `METODO_PAGO_ID` bigint DEFAULT NULL,
  `CONTENIDO_ID` bigint DEFAULT NULL,
  `SUSCRIPCION_ID` bigint DEFAULT NULL,
  `FECHA_TRANSACCION` date DEFAULT NULL COMMENT 'Se registra la fecha de transacion de compra de contenido o pago suscripción',
  `MONTO` decimal(10,2) DEFAULT NULL COMMENT 'Se registra el valor de la transacción',
  `ESTADO_TRANSACCION` varchar(30) DEFAULT NULL COMMENT 'Registro del estado de la transacción (exitosa, fallida)',
  `PLANES_SUSCRIPCION_SUSCRIPCION_ID` bigint NOT NULL,
  `USUARIOS_USUARIO_ID` bigint NOT NULL,
  `CONTENIDOS_CONTENIDO_ID` bigint NOT NULL,
  `METODOS_PAGO_METODO_PAGO_ID` bigint NOT NULL,
  PRIMARY KEY (`TRANSACCION_ID`),
  KEY `TRANSACCIONES_CONTENIDOS_FK` (`CONTENIDOS_CONTENIDO_ID`),
  KEY `TRANSACCIONES_METODOS_PAGO_FK` (`METODOS_PAGO_METODO_PAGO_ID`),
  KEY `TRANSACCIONES_PLANES_SUSCRIPCION_FK` (`PLANES_SUSCRIPCION_SUSCRIPCION_ID`),
  KEY `TRANSACCIONES_USUARIOS_FK` (`USUARIOS_USUARIO_ID`),
  CONSTRAINT `TRANSACCIONES_CONTENIDOS_FK` FOREIGN KEY (`CONTENIDOS_CONTENIDO_ID`) REFERENCES `contenidos` (`CONTENIDO_ID`),
  CONSTRAINT `TRANSACCIONES_METODOS_PAGO_FK` FOREIGN KEY (`METODOS_PAGO_METODO_PAGO_ID`) REFERENCES `metodos_pago` (`METODO_PAGO_ID`),
  CONSTRAINT `TRANSACCIONES_PLANES_SUSCRIPCION_FK` FOREIGN KEY (`PLANES_SUSCRIPCION_SUSCRIPCION_ID`) REFERENCES `planes_suscripcion` (`SUSCRIPCION_ID`),
  CONSTRAINT `TRANSACCIONES_USUARIOS_FK` FOREIGN KEY (`USUARIOS_USUARIO_ID`) REFERENCES `usuarios` (`USUARIO_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Tabla para gestionar las transacciones';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transacciones`
--

LOCK TABLES `transacciones` WRITE;
/*!40000 ALTER TABLE `transacciones` DISABLE KEYS */;
INSERT INTO `transacciones` VALUES (501,1,1,201,NULL,'2024-05-15',15.00,'exitosa',301,1,201,1),(502,1,1,NULL,301,'2024-05-01',10.99,'exitosa',301,1,201,1),(503,2,2,NULL,302,'2023-10-10',99.99,'exitosa',302,2,202,2),(504,3,3,203,NULL,'2024-04-10',18.00,'exitosa',303,3,203,3),(505,4,1,NULL,301,'2023-11-20',119.99,'exitosa',301,4,204,1),(506,5,2,205,NULL,'2024-03-05',20.00,'fallida',302,5,205,2),(507,6,3,NULL,303,'2024-05-10',10.99,'exitosa',303,6,206,3),(508,7,1,207,NULL,'2024-05-28',0.00,'exitosa',301,7,207,1),(509,8,2,NULL,302,'2024-02-25',5.99,'exitosa',302,8,208,2),(510,9,3,209,NULL,'2024-05-30',0.00,'exitosa',303,9,209,3);
/*!40000 ALTER TABLE `transacciones` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuarios`
--

DROP TABLE IF EXISTS `usuarios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usuarios` (
  `USUARIO_ID` bigint NOT NULL AUTO_INCREMENT COMMENT 'Registro del ID del usuario',
  `NOMBRE_USUARIO` varchar(10) DEFAULT NULL COMMENT 'Registro de nombre del usuario',
  `EMAIL` varchar(30) NOT NULL COMMENT 'Usuario registra su correo para iniciar sesion',
  `PASSWORD_HASH` varchar(30) DEFAULT NULL COMMENT 'Usuario registra su contraseña',
  `FECHA_REGISTRO` date DEFAULT NULL COMMENT 'Se registra la fecha de creacion del usuario',
  `METODOS_PAGO_METODO_PAGO_ID` bigint NOT NULL,
  PRIMARY KEY (`USUARIO_ID`),
  UNIQUE KEY `UK_EMAIL` (`EMAIL`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Tabla para gestionar los usuarios';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuarios`
--

LOCK TABLES `usuarios` WRITE;
/*!40000 ALTER TABLE `usuarios` DISABLE KEYS */;
INSERT INTO `usuarios` VALUES (1,'juanperez','juan.perez@example.com','pass123','2023-01-01',1),(2,'mariag','maria.g@example.com','pass456','2023-01-10',2),(3,'carlosr','carlos.r@example.com','pass789','2023-01-15',3),(4,'laurav','laura.v@example.com','passabc','2023-01-20',1),(5,'miguelt','miguel.t@example.com','passdef','2023-01-25',2),(6,'anam','ana.m@example.com','passghi','2023-02-01',3),(7,'davidl','david.l@example.com','passjkl','2023-02-05',1),(8,'sofiah','sofia.h@example.com','passmno','2023-02-10',2),(9,'javiers','javier.s@example.com','passpqr','2023-02-15',3),(10,'elenaq','elena.q@example.com','passtuv','2023-02-20',1);
/*!40000 ALTER TABLE `usuarios` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-06-20 18:28:46
