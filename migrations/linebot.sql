-- MySQL dump 10.13  Distrib 5.6.31, for Linux (x86_64)
--
-- Host: localhost    Database: motivating_menu
-- ------------------------------------------------------
-- Server version	5.7.14-google

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `motivating_menu`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `motivating_menu` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `motivating_menu`;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8mb4 NOT NULL,
  `imageUrl` text CHARACTER SET utf8mb4 NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'ご飯もの',''),(2,'どんぶり・寿司',''),(3,'パン・スナック',''),(4,'めん料理',''),(5,'肉料理',''),(6,'魚料理',''),(7,'卵料理',''),(8,'豆腐料理',''),(9,'野菜・いも料理',''),(10,'サラダ','');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recipe`
--

DROP TABLE IF EXISTS `recipe`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recipe` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `category_id` bigint(20) unsigned NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 NOT NULL,
  `imageUrl` text CHARACTER SET utf8mb4 NOT NULL,
  `count1` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count2` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count3` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count4` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count5` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count6` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count7` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count8` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count9` bigint(20) unsigned NOT NULL DEFAULT '0',
  `count10` bigint(20) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `recipe_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recipe`
--

LOCK TABLES `recipe` WRITE;
/*!40000 ALTER TABLE `recipe` DISABLE KEYS */;
INSERT INTO `recipe` VALUES (1,1,'お茶づけ','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/1_1.jpg',15,10,2,5,9,30,40,10,90,20),(2,1,'鮭飯','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/1_2.jpg',20,15,30,8,1,10,20,2,5,2),(3,2,'いくら丼','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/2_1.jpg',2,5,30,20,40,10,80,90,20,10),(4,2,'うに丼','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/2_2.jpe',0,30,30,10,20,0,10,20,20,80),(5,3,'ガーリックトースト','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/3_1.jpg',22,20,83,22,32,12,20,80,20,0),(6,3,'ハンバーガー','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/3_2.jpg',34,42,32,22,20,21,23,25,26,28),(7,4,'みそラーメン','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/4_1.jpg',10,12,21,23,88,43,22,53,93,32),(8,4,'醤油ラーメン','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/4_2.jpeg',21,23,10,23,11,23,33,42,22,21),(9,5,'肉の塩焼き','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/5_1.jpg',10,32,22,22,23,82,92,42,32,82),(10,5,'鶏肉チーズフライ','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/5_2.jpg',32,12,32,11,32,43,32,76,42,42),(11,6,'かつおのたたき','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/6_1.jpg',83,43,32,21,43,34,21,34,84,27),(12,6,'いわしの梅肉煮','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/6_2.jpg',23,56,32,21,32,47,92,62,85,22),(13,7,'目玉焼き','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/7_1.jpg',32,21,75,54,43,32,1,23,94,64),(14,7,'ハムエッグ','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/7_2.jpeg',63,86,32,43,32,87,21,43,93,23),(15,8,'揚げ出し豆腐','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/8_1.jpg',12,23,44,88,32,35,56,72,93,32),(16,8,'鶏挽肉のゆば包み','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/8_2.jpg',43,37,32,4,5,83,32,4,43,32),(17,9,'ゴーヤチャンプルー','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/9_1.jpg',32,4,6,72,32,43,51,67,82,32),(18,9,'野菜とじゃこの煮物','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/9_2.jpg',32,3,10,0,20,44,32,44,77,32),(19,10,'青菜と貝のゴママヨネーズ','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/10_1.jpg',22,1,32,2,32,4,6,43,93,2),(20,10,'ごぼうのサラダ','https://raw.githubusercontent.com/line-bot/motivating_menu/master/image/category/10_2.jpeg',32,44,55,11,32,22,53,11,0,32);
/*!40000 ALTER TABLE `recipe` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-02-04 20:13:55
