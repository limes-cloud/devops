-- MySQL dump 10.13  Distrib 8.0.27, for macos11 (x86_64)
--
-- Host: localhost    Database: xk_com
-- ------------------------------------------------------
-- Server version	5.7.28

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
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT,
  `username` char(36) NOT NULL,
  `password` char(36) NOT NULL,
  `last_login` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'admin','3883b5e2ea8a4ef46856985ac6d3c5a1','2022-01-19 00:56:15');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admins`
--

DROP TABLE IF EXISTS `admins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admins` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT,
  `username` char(36) NOT NULL,
  `password` char(36) NOT NULL,
  `last_login` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admins`
--

LOCK TABLES `admins` WRITE;
/*!40000 ALTER TABLE `admins` DISABLE KEYS */;
INSERT INTO `admins` VALUES (1,'admins','3883b5e2ea8a4ef46856985ac6d3c5a1','2019-11-23 21:12:50');
/*!40000 ALTER TABLE `admins` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course`
--

DROP TABLE IF EXISTS `course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `teach_id` tinyint(4) NOT NULL,
  `name` char(36) NOT NULL,
  `desc` tinytext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `teach_id` (`teach_id`),
  CONSTRAINT `course_ibfk_1` FOREIGN KEY (`teach_id`) REFERENCES `teacher` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course`
--

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;
INSERT INTO `course` VALUES (24,1,'ss','sss'),(25,1,'sa1sa','sa1dsas1');
/*!40000 ALTER TABLE `course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `course_class`
--

DROP TABLE IF EXISTS `course_class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `course_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL,
  `sclass_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `course_id` (`course_id`,`sclass_id`)
) ENGINE=InnoDB AUTO_INCREMENT=262 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course_class`
--

LOCK TABLES `course_class` WRITE;
/*!40000 ALTER TABLE `course_class` DISABLE KEYS */;
INSERT INTO `course_class` VALUES (170,19,297),(171,19,305),(183,19,325),(173,19,328),(182,19,416),(175,19,417),(176,20,303),(177,20,311),(179,20,312),(178,20,327),(180,20,419),(185,21,296),(184,21,310),(186,21,366),(187,21,370),(213,21,443),(188,22,298),(191,22,313),(189,22,314),(190,22,315),(193,22,368),(192,22,374),(214,22,444),(194,23,306),(195,23,365),(196,23,372),(197,23,411),(198,23,414),(199,24,309),(202,24,326),(201,24,367),(200,24,373),(216,24,445),(215,24,446),(205,25,301),(206,25,302),(207,25,308),(203,25,324),(204,25,369),(209,28,304),(210,28,307),(208,28,323),(211,28,371),(212,28,413),(219,29,450),(217,29,467),(218,29,468),(222,29,478),(220,29,481),(221,29,485),(223,30,464),(224,30,471),(225,30,473),(227,30,479),(226,30,486),(228,30,488),(231,31,447),(230,31,451),(232,31,452),(229,31,465),(234,31,484),(233,31,491),(235,32,453),(236,32,458),(237,32,459),(238,32,463),(239,32,470),(241,32,472),(240,32,482),(242,33,454),(243,33,456),(244,33,457),(245,33,469),(248,33,476),(246,33,480),(247,33,483),(249,34,455),(250,34,466),(254,34,474),(252,34,475),(253,34,477),(251,34,489),(258,35,448),(257,35,449),(256,35,460),(255,35,461),(259,35,487),(260,35,490),(261,36,492);
/*!40000 ALTER TABLE `course_class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_teach`
--

DROP TABLE IF EXISTS `item_teach`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `item_teach` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_id` int(11) NOT NULL,
  `teach_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `item_id` (`item_id`),
  KEY `teach_id` (`teach_id`),
  CONSTRAINT `item_teach_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `sports_item` (`item_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `item_teach_ibfk_2` FOREIGN KEY (`teach_id`) REFERENCES `teach_admin` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_teach`
--

LOCK TABLES `item_teach` WRITE;
/*!40000 ALTER TABLE `item_teach` DISABLE KEYS */;
INSERT INTO `item_teach` VALUES (5,12,4),(6,13,3),(7,14,7),(8,15,6),(10,17,5),(11,18,3),(13,20,3),(14,21,8),(15,22,9);
/*!40000 ALTER TABLE `item_teach` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `join_in`
--

DROP TABLE IF EXISTS `join_in`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `join_in` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL,
  `student_id` tinyint(4) NOT NULL,
  `add_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `course_id` (`course_id`,`student_id`),
  KEY `student_id` (`student_id`),
  CONSTRAINT `join_in_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE CASCADE,
  CONSTRAINT `join_in_ibfk_2` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9024 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `join_in`
--

LOCK TABLES `join_in` WRITE;
/*!40000 ALTER TABLE `join_in` DISABLE KEYS */;
INSERT INTO `join_in` VALUES (9023,24,1,'2022-01-20 01:33:39');
/*!40000 ALTER TABLE `join_in` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL,
  `student_id` tinyint(4) NOT NULL,
  `title` varchar(128) NOT NULL,
  `content` text NOT NULL,
  `add_time` datetime NOT NULL,
  `student_name` varchar(15) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `course_id` (`course_id`),
  KEY `student_id` (`student_id`),
  CONSTRAINT `message_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE CASCADE,
  CONSTRAINT `message_ibfk_2` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
INSERT INTO `message` VALUES (1,24,1,'关于微积分正态分布的问题','关于微积分正态分布的问题关于微积分正态分布的问题关于微积分正态分布的问题关于微积分正态分布的问题','2021-12-12 12:12:12','方伟业'),(2,24,1,'sasa','ssssšsssssss','2022-01-20 01:43:03','方伟业');
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message_comment`
--

DROP TABLE IF EXISTS `message_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `msg_id` int(11) NOT NULL,
  `replay_name` varchar(15) NOT NULL,
  `type` varchar(15) NOT NULL,
  `content` text NOT NULL,
  `add_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `msg_id` (`msg_id`),
  CONSTRAINT `message_comment_ibfk_1` FOREIGN KEY (`msg_id`) REFERENCES `message` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message_comment`
--

LOCK TABLES `message_comment` WRITE;
/*!40000 ALTER TABLE `message_comment` DISABLE KEYS */;
INSERT INTO `message_comment` VALUES (1,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(2,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(3,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(4,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(5,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(6,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(7,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(8,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(9,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(10,1,'方伟业','student','这是一个正经的恢复','2020-12-14 12:12:12'),(11,1,'admin','teacher','我也来讨论了','2022-01-20 00:54:17'),(12,1,'方伟业','student','sassa','2022-01-20 01:33:50');
/*!40000 ALTER TABLE `message_comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `open_web`
--

DROP TABLE IF EXISTS `open_web`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `open_web` (
  `is_open` tinyint(1) NOT NULL,
  PRIMARY KEY (`is_open`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `open_web`
--

LOCK TABLES `open_web` WRITE;
/*!40000 ALTER TABLE `open_web` DISABLE KEYS */;
INSERT INTO `open_web` VALUES (1);
/*!40000 ALTER TABLE `open_web` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sclass`
--

DROP TABLE IF EXISTS `sclass`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sclass` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class_id` int(11) NOT NULL,
  `class_name` char(45) NOT NULL,
  `class_year` int(11) NOT NULL,
  `scxb` char(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `class_id` (`class_id`),
  UNIQUE KEY `class_name` (`class_name`)
) ENGINE=InnoDB AUTO_INCREMENT=493 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sclass`
--

LOCK TABLES `sclass` WRITE;
/*!40000 ALTER TABLE `sclass` DISABLE KEYS */;
INSERT INTO `sclass` VALUES (296,704118,'视觉181',2018,'艺术部'),(297,704128,'视觉182',2018,'艺术部'),(298,704132,'动画181',2018,'艺术部'),(299,704588,'18级舞蹈学1班',2018,'艺术部'),(301,705480,'中文181',2018,'文学部'),(302,705506,'中文182',2018,'文学部'),(303,705514,'中文183',2018,'文学部'),(304,705522,'中文184',2018,'文学部'),(305,705586,'新闻181',2018,'文学部'),(306,705598,'新闻182',2018,'文学部'),(307,705610,'英语181',2018,'文学部'),(308,705618,'英语182',2018,'文学部'),(309,705646,'英语183',2018,'文学部'),(310,708958,'环境设计181',2018,'艺术部'),(311,708962,'环境设计182',2018,'艺术部'),(312,708996,'绘画181',2018,'艺术部'),(313,709058,'音乐表演181',2018,'艺术部'),(314,709062,'环境设计183',2018,'艺术部'),(315,712996,'表演181',2018,'艺术部'),(323,757496,'计科181',2018,'工学部'),(324,757498,'计科182',2018,'工学部'),(325,757500,'电信181',2018,'工学部'),(326,757504,'电科181',2018,'工学部'),(327,757506,'物联181',2018,'工学部'),(328,757512,'通信181',2018,'工学部'),(365,757692,'法学183',2018,'法管部'),(366,757694,'法学182',2018,'法管部'),(367,757698,'法学184',2018,'法管部'),(368,757736,'财务管理181',2018,'商学部'),(369,757740,'财务管理182',2018,'商学部'),(370,757744,'财务管理183',2018,'商学部'),(371,757748,'财务管理184',2018,'商学部'),(372,757752,'工商管理181',2018,'商学部'),(373,757756,'工商管理182',2018,'商学部'),(374,757760,'国际经济与贸易181',2018,'商学部'),(411,757942,'金融学181',2018,'商学部'),(413,757948,'金融学182',2018,'商学部'),(414,757954,'金融学183',2018,'商学部'),(416,757964,'金融学184',2018,'商学部'),(417,757970,'旅游管理181',2018,'商学部'),(419,757976,'工程管理181',2018,'商学部'),(443,757526,'法学181班',2018,'法学与公共管理学部'),(444,757530,'行政管理181班',2018,'法学与公共管理学部'),(445,757534,'行政管理182班',2018,'法学与公共管理学部'),(446,757536,'公共管理181班',2018,'法学与公共管理学部'),(447,1004302,'计科科191班',2019,'工学部'),(448,1004304,'计科科192班',2019,'工学部'),(449,1004264,'电科科191班',2019,'工学部'),(450,1004268,'电信科191班',2019,'工学部'),(451,1004326,'物联科191班',2019,'工学部'),(452,1004320,'通信科191班',2019,'工学部'),(453,1004270,'动画科191班',2019,'艺术部'),(454,1004294,'环艺科191班',2019,'艺术部'),(455,1004296,'环艺科192班',2019,'艺术部'),(456,1004298,'环艺科193班',2019,'艺术部'),(457,1004250,'表演科191班',2019,'艺术部'),(458,1004352,'视觉科191班',2019,'艺术部'),(459,1004354,'视觉科192班',2019,'艺术部'),(460,1004336,'音乐科191班',2019,'艺术部'),(461,1004300,'绘画科191班',2019,'艺术部'),(462,1004322,'舞蹈科191班',2019,'艺术部'),(463,1004342,'中文科191班',2019,'文学部'),(464,1004346,'中文科192班',2019,'文学部'),(465,1004348,'中文科193班',2019,'文学部'),(466,1004350,'中文科194班',2019,'文学部'),(467,1004338,'英语科191班',2019,'文学部'),(468,1004340,'英语科192班',2019,'文学部'),(469,1005116,'英语科193班',2019,'文学部'),(470,1004330,'新闻科191班',2019,'文学部'),(471,1004334,'新闻科192班',2019,'文学部'),(472,1004286,'公管科191班',2019,'法学与公共管理学部'),(473,1004290,'行管科191班',2019,'法学与公共管理学部'),(474,1004292,'行管科192班',2019,'法学与公共管理学部'),(475,1004272,'法学科191班',2019,'法学与公共管理学部'),(476,1004274,'法学科192班',2019,'法学与公共管理学部'),(477,1004276,'法学科193班',2019,'法学与公共管理学部'),(478,1004278,'法学科194班',2019,'法学与公共管理学部'),(479,1004282,'工商科191班',2019,'商学部'),(480,1004284,'工商科192班',2019,'商学部'),(481,1004280,'工程科191班',2019,'商学部'),(482,1004254,'财管科191班',2019,'商学部'),(483,1004256,'财管科192班',2019,'商学部'),(484,1004258,'财管科193班',2019,'商学部'),(485,1004260,'财管科194班',2019,'商学部'),(486,1004288,'国贸科191班',2019,'商学部'),(487,1004306,'金融科191班',2019,'商学部'),(488,1004308,'金融科192班',2019,'商学部'),(489,1004310,'金融科193班',2019,'商学部'),(490,1004312,'金融科194班',2019,'商学部'),(491,1004314,'旅管科191班',2019,'商学部');
/*!40000 ALTER TABLE `sclass` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `score`
--

DROP TABLE IF EXISTS `score`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `score` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL,
  `student_id` tinyint(4) NOT NULL,
  `examination` tinyint(4) NOT NULL,
  `expression` tinyint(4) NOT NULL,
  `add_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `course_id` (`course_id`),
  KEY `student_id` (`student_id`)
) ENGINE=MyISAM AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `score`
--

LOCK TABLES `score` WRITE;
/*!40000 ALTER TABLE `score` DISABLE KEYS */;
INSERT INTO `score` VALUES (14,24,1,100,10,'2022-01-19 02:11:03');
/*!40000 ALTER TABLE `score` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sport_desc`
--

DROP TABLE IF EXISTS `sport_desc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sport_desc` (
  `id` int(11) NOT NULL,
  `notice` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sport_desc`
--

LOCK TABLES `sport_desc` WRITE;
/*!40000 ALTER TABLE `sport_desc` DISABLE KEYS */;
INSERT INTO `sport_desc` VALUES (1,'1：每个人只能选择一个项目|2：服从老师调剂|3：网上选课结束后24小时之内，如需调换项目的同学必须遵循“一进一出”的原则，两人持证件同时到场相互调换项目，调换项目的地点：桃园8号楼217室。');
/*!40000 ALTER TABLE `sport_desc` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sports_item`
--

DROP TABLE IF EXISTS `sports_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sports_item` (
  `item_id` int(11) NOT NULL AUTO_INCREMENT,
  `item_name` char(36) NOT NULL,
  `item_desc` tinytext NOT NULL,
  `join_sex` tinyint(1) NOT NULL,
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sports_item`
--

LOCK TABLES `sports_item` WRITE;
/*!40000 ALTER TABLE `sports_item` DISABLE KEYS */;
INSERT INTO `sports_item` VALUES (12,'乒乓球','乒乓球（table tennis），中国国球，是一种世界流行的球类体育项目，包括进攻、对抗和防守。比赛分团体、单打、双打、混双等数种；2001年9月1日前以21分为一局，现以11分为一局；采用五局三',2),(13,'足球','足球，英文：football，美国称为soccer，被誉为“世界第一运动”，全球体育界最具影响力的单项体育运动。标准的11人制足球比赛由两队各派10名球员与1名守门员，总共22人，在长方形的草地球场',2),(14,'排球','排球运动源于美国，由网球运动演变发展而成。球场是长方形，中间隔有高网，比赛双方(每方六人)各占球场的一方，球员用手把球从网上空打来打去。排球运动使用的球，用羊皮或人造革做壳',2),(15,'健美操','健美操中大量吸收了迪斯科舞、爵士舞、霹雳舞中的上下肢、躯干、头颈和足踩动作，特别是髋部动作，这给健美操增添了活力，同时也有利于减少臀部和腹部脂肪的堆积，有利于改善动作的协',2),(17,'羽毛球','羽毛球是一项室内、室外都可以进行的体育运动。依据参与的人数，可以分为单打与双打，及新兴的3打3。羽毛球拍由拍面、拍杆、拍柄及拍框与拍杆的接头构成。',2),(18,'武术1','武术是古代军事战争一种传承的技术。习武可以强身健体，亦可以防御敌人进攻。习武之人以“制止侵袭”为技术导向、引领修习者进入认识人与自然、社会客观规律的传统教化（武化）方式，',2),(20,'篮球1','篮球，英文（basketball），起源于美国马萨诸塞州，是1891年12月21日由詹姆斯·奈史密斯创造，是奥运会核心比赛项目，是以手为中心的身体对抗性体育运动 [1] 。',2),(21,'篮球2','篮球，英文（basketball），起源于美国马萨诸塞州，是1891年12月21日由詹姆斯·奈史密斯创造，是奥运会核心比赛项目，是以手为中心的身体对抗性体育运动 [1] 。',2),(22,'武术2','武术是古代军事战争一种传承的技术。习武可以强身健体，亦可以防御敌人进攻。习武之人以“制止侵袭”为技术导向、引领修习者进入认识人与自然、社会客观规律的传统教化（武化）方式，',2);
/*!40000 ALTER TABLE `sports_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student`
--

DROP TABLE IF EXISTS `student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `student` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT,
  `name` char(36) NOT NULL,
  `class` char(36) NOT NULL,
  `username` char(36) NOT NULL,
  `password` char(36) NOT NULL,
  `last_login` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student`
--

LOCK TABLES `student` WRITE;
/*!40000 ALTER TABLE `student` DISABLE KEYS */;
INSERT INTO `student` VALUES (1,'方伟业','即刻可171','1720041704','b7c870290c713891e6bcc02f91321d7e','2022-01-20 01:03:46'),(2,'方伟业','即刻可','17200','6222d5fd70abd4fe144a1df81fc2ea98','2022-01-19 00:27:04');
/*!40000 ALTER TABLE `student` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teach_admin`
--

DROP TABLE IF EXISTS `teach_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `teach_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `teach_name` char(18) NOT NULL,
  `username` char(36) NOT NULL,
  `password` char(36) NOT NULL,
  `last_login` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teach_admin`
--

LOCK TABLES `teach_admin` WRITE;
/*!40000 ALTER TABLE `teach_admin` DISABLE KEYS */;
INSERT INTO `teach_admin` VALUES (3,'吴昊','tyjys-wh','eade5a2fad9117903a138f3cbfbee0ec','2019-09-06 09:22:10'),(4,'陈志敏','tyjys-czm','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:03:46'),(5,'吴杰','tyjys-wj','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:04:21'),(6,'史小红','tyjys-sxh','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:04:43'),(7,'裴岚','tyjys-pl','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:05:08'),(8,'陈力','tyjys-cl','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:05:34'),(9,'唐松丽','tyjys-tsl','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:06:01'),(10,'袁凡','tyjys-yf','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:06:34'),(11,'陆沁蓉','tyjys-lqr','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:06:59'),(12,'朱应飞','tyjys-zyf','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:07:32'),(13,'杨保亚','tyjys-yby','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:07:58'),(14,'陈晓春','tyjys-cxc','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:08:22'),(15,'于海浩','tyjys-yhh','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:08:43'),(16,'包健','tyjys-bj','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:09:14'),(17,'于洪智','tyjys-yhz','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:09:38');
/*!40000 ALTER TABLE `teach_admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teacher`
--

DROP TABLE IF EXISTS `teacher`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `teacher` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT,
  `teach_name` char(36) NOT NULL,
  `username` char(36) NOT NULL,
  `password` char(36) NOT NULL,
  `last_login` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teacher`
--

LOCK TABLES `teacher` WRITE;
/*!40000 ALTER TABLE `teacher` DISABLE KEYS */;
INSERT INTO `teacher` VALUES (1,'老师','admin','3883b5e2ea8a4ef46856985ac6d3c5a1','2022-01-19 00:57:02'),(3,'吴昊','tyjys-wh','eade5a2fad9117903a138f3cbfbee0ec','2019-09-06 09:22:10'),(4,'陈志敏','tyjys-czm','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:03:46'),(5,'吴杰','tyjys-wj','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:04:21'),(6,'史小红','tyjys-sxh','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:04:43'),(7,'裴岚','tyjys-pl','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:05:08'),(8,'陈力','tyjys-cl','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:05:34'),(9,'唐松丽','tyjys-tsl','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:06:01'),(10,'袁凡','tyjys-yf','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:06:34'),(11,'陆沁蓉','tyjys-lqr','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:06:59'),(12,'朱应飞','tyjys-zyf','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:07:32'),(13,'杨保亚','tyjys-yby','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:07:58'),(14,'陈晓春','tyjys-cxc','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:08:22'),(15,'于海浩','tyjys-yhh','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:08:43'),(16,'包健','tyjys-bj','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:09:14'),(17,'于洪智','tyjys-yhz','eade5a2fad9117903a138f3cbfbee0ec','2019-09-24 09:09:38');
/*!40000 ALTER TABLE `teacher` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-20 22:54:18
