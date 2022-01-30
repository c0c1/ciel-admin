-- MySQL dump 10.13  Distrib 8.0.24, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: ciel_admin
-- ------------------------------------------------------
-- Server version	8.0.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `s_admin`
--

DROP TABLE IF EXISTS `s_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_admin` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `rid` bigint unsigned DEFAULT NULL,
  `uname` varchar(64) DEFAULT NULL,
  `icon` varchar(64) DEFAULT NULL,
  `pwd` varchar(255) DEFAULT NULL,
  `status` int DEFAULT '1',
  `desc` varchar(255) DEFAULT NULL,
  `ex1` varchar(32) DEFAULT NULL COMMENT 'other info',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `rid` (`rid`),
  CONSTRAINT `s_admin_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_admin`
--

LOCK TABLES `s_admin` WRITE;
/*!40000 ALTER TABLE `s_admin` DISABLE KEYS */;
INSERT INTO `s_admin` (`id`, `rid`, `uname`, `icon`, `pwd`, `status`, `desc`, `ex1`, `created_at`, `updated_at`) VALUES (18,8,'admin','icon/2021/11/In1vuW.png','$2a$10$Fv4wZ4XS3HR8UbKwE8AUyuXX3ZjETilumZbQP7nB9YyVj33LthxxO',1,'','true','2021-11-23 12:18:21','2022-01-29 15:57:48'),(46,9,'test2','','$2a$10$Ny0c.Yslfm4F2kl.dh7Z4edZQC.DxD.CFP.qYRD5VSrZNXl7o2bmm',1,'22','0','2022-01-04 16:10:33','2022-01-29 15:07:23');
/*!40000 ALTER TABLE `s_admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_api`
--

DROP TABLE IF EXISTS `s_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(64) DEFAULT NULL,
  `method` varchar(64) DEFAULT NULL,
  `group` varchar(64) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `status` int DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_api`
--

LOCK TABLES `s_api` WRITE;
/*!40000 ALTER TABLE `s_api` DISABLE KEYS */;
INSERT INTO `s_api` (`id`, `url`, `method`, `group`, `desc`, `status`, `created_at`, `updated_at`) VALUES (2,'/menu/add','POST','menu','22',1,'2021-11-19 19:02:52','2021-11-19 19:02:52'),(3,'/menu','DELETE','menu','',1,'2021-11-19 19:02:55','2021-11-19 19:02:55'),(4,'/menu/update','PUT','menu','',1,'2021-11-19 19:02:32','2021-11-19 19:02:32'),(5,'/api/add','POST','api','',1,'2021-11-19 19:03:47','2021-11-19 19:03:47'),(6,'/api','DELETE','api','',1,'2021-11-19 19:03:58','2021-11-19 19:03:58'),(7,'/api/update','PUT','api','',1,'2021-11-19 19:04:06','2021-11-19 19:04:06'),(8,'/role/add','POST','role','',1,'2021-11-20 16:39:50','2021-11-20 16:39:50'),(9,'/role','DELETE','role','',1,'2021-11-20 16:40:01','2021-11-20 16:40:01'),(10,'/role','DELETE','role','',1,'2021-11-20 16:40:08','2021-11-20 16:40:08'),(11,'/admin/add','POST','admin','',1,'2021-11-20 16:40:26','2022-01-25 15:39:01'),(12,'/admin','DELETE','admin','',1,'2021-11-23 11:42:00','2022-01-25 15:39:05'),(13,'/admin/update','PUT','admin','2',1,'2021-12-04 11:01:40','2022-01-27 16:05:14'),(14,'/dict/add','POST','dict','',1,'2021-11-20 16:40:58','2021-11-20 16:40:58'),(15,'/dict','DELETE','dict','',1,'2021-11-20 16:41:07','2021-11-20 16:41:07'),(16,'/dict/update','PUT','dict','',1,'2021-11-20 16:41:12','2022-01-28 18:18:06'),(17,'/file/add','POST','file','',1,'2021-11-20 16:41:20','2022-01-28 18:18:01'),(18,'/file','DELETE','file','',1,'2021-11-20 16:41:27','2022-01-28 15:20:45'),(19,'/file/update','PUT','file','file',1,'2021-11-20 16:41:35','2022-01-28 18:49:39'),(20,'/file/upload','POST','file','222',1,'2021-11-20 16:42:35','2022-01-30 15:21:38');
/*!40000 ALTER TABLE `s_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_dict`
--

DROP TABLE IF EXISTS `s_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_dict` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `k` varchar(64) NOT NULL,
  `v` varchar(255) NOT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `group` varchar(64) NOT NULL DEFAULT 'sys',
  `status` int DEFAULT NULL,
  `type` int NOT NULL DEFAULT '1' COMMENT '0 文本，1 img',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_dict`
--

LOCK TABLES `s_dict` WRITE;
/*!40000 ALTER TABLE `s_dict` DISABLE KEYS */;
INSERT INTO `s_dict` (`id`, `k`, `v`, `desc`, `group`, `status`, `type`, `created_at`, `updated_at`) VALUES (14,'img_prefix','http://localhost:2022/resource/upload/','','',1,4,'2022-01-25 15:46:27','2022-01-29 15:29:02');
/*!40000 ALTER TABLE `s_dict` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_file`
--

DROP TABLE IF EXISTS `s_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `group` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=233 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_file`
--

LOCK TABLES `s_file` WRITE;
/*!40000 ALTER TABLE `s_file` DISABLE KEYS */;
INSERT INTO `s_file` (`id`, `name`, `group`, `status`, `created_at`, `updated_at`) VALUES (224,'file/2022/01/WLB12M.png',3,1,'2022-01-17 16:32:20','2022-01-25 16:12:52'),(226,'file/2022/01/AAelFN.png',3,1,'2022-01-17 16:32:20','2022-01-25 16:10:42'),(227,'file/2022/01/fGg9RV.png',3,1,'2022-01-17 16:32:20','2022-01-28 19:35:57');
/*!40000 ALTER TABLE `s_file` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_icon`
--

DROP TABLE IF EXISTS `s_icon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_icon` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` int DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Sys icon';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_icon`
--

LOCK TABLES `s_icon` WRITE;
/*!40000 ALTER TABLE `s_icon` DISABLE KEYS */;
INSERT INTO `s_icon` (`id`, `content`, `status`, `created_at`, `updated_at`) VALUES (1,'FaceRetouchingNa',1,'2022-01-05 13:56:01','2022-01-05 13:58:04'),(2,'FaceIcon',1,'2022-01-05 13:58:47','2022-01-05 13:58:47'),(3,'ListAltIcon',1,'2022-01-05 13:58:54','2022-01-05 13:58:54'),(4,'LiquorIcon',1,'2022-01-05 13:59:39','2022-01-05 13:59:39'),(5,'BrandingWatermarkIcon',1,'2022-01-05 13:59:55','2022-01-05 14:00:59'),(6,'BadgeIcon',1,'2022-01-05 14:01:07','2022-01-05 14:01:07'),(7,'AutoFixHighIcon',1,'2022-01-05 14:01:15','2022-01-05 14:01:15'),(8,'AddRoadIcon',1,'2022-01-05 14:02:35','2022-01-05 14:02:35'),(9,'FormatListNumberedIcon',1,'2022-01-05 14:02:44','2022-01-05 14:02:44'),(10,'AlignHorizontalLeftIcon',1,'2022-01-05 14:02:54','2022-01-05 14:02:54'),(11,'SavingsIcon',1,'2022-01-05 14:03:04','2022-01-05 14:03:04'),(12,'AttachMoneyIcon',1,'2022-01-05 14:03:16','2022-01-05 14:03:16'),(13,'PaidIcon',1,'2022-01-05 14:03:25','2022-01-05 14:03:25'),(14,'CreditCardIcon',1,'2022-01-05 14:03:32','2022-01-05 14:03:32'),(15,'AccountBalanceIcon',1,'2022-01-05 14:03:39','2022-01-05 14:03:39'),(16,'FormatListNumberedRtlIcon',1,'2022-01-05 14:03:44','2022-01-05 14:03:44'),(17,'PeopleIcon',1,'2022-01-05 14:03:56','2022-01-05 14:03:56'),(18,'QuestionAnswerIcon',1,'2022-01-05 14:04:00','2022-01-05 14:04:00'),(19,'LocalAtmIcon',1,'2022-01-05 14:04:05','2022-01-05 14:04:05'),(20,'AccountBalanceWalletIcon',1,'2022-01-05 14:04:11','2022-01-05 14:04:11'),(21,'AttachFileIcon',1,'2022-01-05 14:04:18','2022-01-05 14:04:18'),(22,'MenuBookIcon',1,'2022-01-05 14:04:24','2022-01-05 14:04:24'),(23,'SupervisorAccountIcon',1,'2022-01-05 14:04:33','2022-01-05 14:04:33'),(24,'ContactPageIcon',1,'2022-01-05 14:04:40','2022-01-05 14:04:40'),(25,'ListIcon',1,'2022-01-05 14:04:47','2022-01-05 14:04:47'),(26,'MenuIcon',1,'2022-01-05 14:04:53','2022-01-05 14:04:53'),(27,'SettingsIcon',1,'2022-01-05 14:04:58','2022-01-05 14:04:58');
/*!40000 ALTER TABLE `s_icon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_menu`
--

DROP TABLE IF EXISTS `s_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `pid` bigint DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `path` varchar(64) DEFAULT NULL,
  `icon` varchar(54) DEFAULT NULL,
  `type` int unsigned NOT NULL DEFAULT '0' COMMENT '0 normal,1 group,2 divide',
  `sort` decimal(7,2) DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1' COMMENT '-1 off 1 ok',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=360 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_menu`
--

LOCK TABLES `s_menu` WRITE;
/*!40000 ALTER TABLE `s_menu` DISABLE KEYS */;
INSERT INTO `s_menu` (`id`, `pid`, `name`, `path`, `icon`, `type`, `sort`, `status`, `created_at`, `updated_at`) VALUES (2,34,'Menu','/sys/menu','MenuIcon',2,1.10,1,'2021-12-03 10:19:33','2022-01-26 16:31:31'),(3,34,'API','/sys/api','ListIcon',2,1.20,1,'2021-12-03 10:19:36','2021-11-23 10:13:47'),(4,34,'Role','/sys/role','ContactPageIcon',2,1.30,1,'2021-12-03 10:19:36','2022-01-26 16:31:24'),(5,34,'Admin','/sys/admin','SupervisorAccountIcon',2,1.40,1,'2021-12-03 10:19:36','2022-01-26 16:31:05'),(6,34,'Dict','/sys/dict','MenuBookIcon',2,1.50,1,'2021-12-03 10:19:36','2022-01-26 16:31:10'),(7,34,'Files','/sys/file','AttachFileIcon',2,1.60,1,'2021-12-03 10:19:36','2022-01-26 16:31:16'),(34,-1,'System','','SettingsIcon',1,1.00,1,'2021-12-03 10:19:36','2022-01-26 16:31:36'),(359,0,'','','',0,0.00,0,'2022-01-30 15:22:36','2022-01-30 15:22:36');
/*!40000 ALTER TABLE `s_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_role`
--

DROP TABLE IF EXISTS `s_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL,
  `status` int DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role`
--

LOCK TABLES `s_role` WRITE;
/*!40000 ALTER TABLE `s_role` DISABLE KEYS */;
INSERT INTO `s_role` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES (8,'超级管理员',1,'2021-11-19 07:50:25','2022-01-25 15:40:31'),(9,'管理员',1,'2021-11-19 19:15:50','2022-01-24 14:34:43');
/*!40000 ALTER TABLE `s_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_role_api`
--

DROP TABLE IF EXISTS `s_role_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_role_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `rid` bigint unsigned DEFAULT NULL,
  `aid` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `rid` (`rid`,`aid`),
  CONSTRAINT `s_role_api_ibfk_1` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=147 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_api`
--

LOCK TABLES `s_role_api` WRITE;
/*!40000 ALTER TABLE `s_role_api` DISABLE KEYS */;
/*!40000 ALTER TABLE `s_role_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `s_role_menu`
--

DROP TABLE IF EXISTS `s_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `s_role_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `rid` bigint unsigned DEFAULT NULL,
  `mid` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `rid_2` (`rid`,`mid`),
  KEY `mid` (`mid`),
  KEY `rid` (`rid`,`mid`),
  CONSTRAINT `s_role_menu_ibfk_1` FOREIGN KEY (`mid`) REFERENCES `s_menu` (`id`) ON DELETE CASCADE,
  CONSTRAINT `s_role_menu_ibfk_2` FOREIGN KEY (`rid`) REFERENCES `s_role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=307 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `s_role_menu`
--

LOCK TABLES `s_role_menu` WRITE;
/*!40000 ALTER TABLE `s_role_menu` DISABLE KEYS */;
INSERT INTO `s_role_menu` (`id`, `rid`, `mid`) VALUES (121,8,2),(122,8,3),(123,8,4),(124,8,5),(125,8,6),(182,8,7),(141,8,34),(284,9,2),(285,9,3),(286,9,4),(287,9,5),(288,9,6),(289,9,7),(290,9,34);
/*!40000 ALTER TABLE `s_role_menu` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-30 16:09:53
