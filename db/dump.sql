-- MySQL dump 10.13  Distrib 8.0.13, for Win64 (x86_64)
--
-- Host: localhost    Database: ht
-- ------------------------------------------------------
-- Server version	8.0.13

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `campaign`
--

DROP TABLE IF EXISTS `campaign`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `campaign` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `total_received` varchar(45) DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  `course_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_campaign_user1` (`user_id`),
  KEY `fk_campaign_course1` (`course_id`),
  CONSTRAINT `fk_campaign_course1` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`),
  CONSTRAINT `fk_campaign_user1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `campaign`
--

LOCK TABLES `campaign` WRITE;
/*!40000 ALTER TABLE `campaign` DISABLE KEYS */;
INSERT INTO `campaign` VALUES (1,'300',1,1),(2,'4500',2,2),(3,'5980',3,3),(4,'49000',4,4),(5,'0',5,1);
/*!40000 ALTER TABLE `campaign` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course`
--

DROP TABLE IF EXISTS `course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `description` varchar(500) DEFAULT NULL,
  `cost` int(11) DEFAULT NULL,
  `length` varchar(45) DEFAULT NULL,
  `requirements` varchar(500) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `school_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_course_school1` (`school_id`),
  CONSTRAINT `fk_course_school1` FOREIGN KEY (`school_id`) REFERENCES `school` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course`
--

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;
INSERT INTO `course` VALUES (1,'Payments Everywhere','Course to understand more the visa business',500,'5 months','Being a Visa employee','paymentseverywhere@visa.com',1),(2,'Computer Science','Learn about computer science and application of technology in the modern world',27000,'3 years','ABB A-Levels','uor_admissions@reading.ac.uk',2),(3,'Computer Science','Learn about computer science and application of technology in the modern world',27000,'3 years','AAA A-Levels','ucl_admissions@ucl.ac.uk',3),(4,'History','Learn about past, present and future of human history and how it affected, affects and will affect society',50000,'5 years','AAA A-Levels','harvard_admissions@harvard.com',4);
/*!40000 ALTER TABLE `course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post`
--

DROP TABLE IF EXISTS `post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `post` (
  `id` int(11) NOT NULL,
  `timestamp` datetime DEFAULT NULL,
  `description` varchar(500) DEFAULT NULL,
  `total_campaign_snapshot` int(11) DEFAULT NULL,
  `campaign_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_post_campaign_12` (`campaign_id`),
  CONSTRAINT `fk_post_campaign_12` FOREIGN KEY (`campaign_id`) REFERENCES `campaign` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post`
--

LOCK TABLES `post` WRITE;
/*!40000 ALTER TABLE `post` DISABLE KEYS */;
INSERT INTO `post` VALUES (1,'2019-01-05 14:12:32','Hi, my name is Sam and I\'ve just joined Visa because I\'m very passionate about payments and I would like to start this course to learn even more about payments.',0,1),(2,'2018-10-05 19:12:32','Hi, my name is Mikey and in recent years I ot passionate about computer science and its modern applications.',0,2),(3,'2018-11-05 19:12:32','Hi, my name is Peter and I would love to learn more about machine learning and blockchain, this course will help me achieve these goals. ',0,3),(4,'2018-11-05 19:12:32','Hi, my name James and I would love to study in the US to study computer science.',0,4),(5,'2018-11-05 19:12:32','Hi, my name is Tahnik and I would like to learn more about payments and the Visa business in general.',0,5);
/*!40000 ALTER TABLE `post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `school`
--

DROP TABLE IF EXISTS `school`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `school` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `description` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `school`
--

LOCK TABLES `school` WRITE;
/*!40000 ALTER TABLE `school` DISABLE KEYS */;
INSERT INTO `school` VALUES (1,'Visa University','Visa Corporate university'),(2,'UoR','Reading town university'),(3,'UCL','University of central london'),(4,'Harvard','Harvard University');
/*!40000 ALTER TABLE `school` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone` int(11) DEFAULT NULL,
  `country` varchar(100) DEFAULT NULL,
  `usertype_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_usertype1` (`usertype_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'samsilva','123','sam','sam@gmail.com',7846212,'UK',1),(2,'mikeyw','123','mikey','mikey@gmail.com',7846212,'UK',1),(3,'peterd','123','peter','peter@gmail.com',7846212,'UK',1),(4,'jameb','123','james','james@gmail.com',7846212,'US',1),(5,'tahnik','123','tahnik','tahnik@gmail.com',7846212,'UK',1),(6,'Visa','123','Visa europe','visa@visa.com',7846213,'UK',2);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usertype`
--

DROP TABLE IF EXISTS `usertype`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `usertype` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `description` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usertype`
--

LOCK TABLES `usertype` WRITE;
/*!40000 ALTER TABLE `usertype` DISABLE KEYS */;
INSERT INTO `usertype` VALUES (1,'individual','profile intende for individual clients'),(2,'organisation','profile intended for organisational type of bodies such as companies');
/*!40000 ALTER TABLE `usertype` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-01-20 16:34:47
