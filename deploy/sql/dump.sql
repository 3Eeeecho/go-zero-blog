-- MySQL dump 10.13  Distrib 8.0.41, for Linux (x86_64)
--
-- Host: localhost    Database: blog
-- ------------------------------------------------------
-- Server version	8.0.41-0ubuntu0.24.04.1

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
-- Table structure for table `article_likes`
--

DROP TABLE IF EXISTS `article_likes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article_likes` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int unsigned NOT NULL,
  `user_id` int unsigned NOT NULL,
  `created_on` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_article_user` (`article_id`,`user_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `article_likes_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`),
  CONSTRAINT `article_likes_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_likes`
--

LOCK TABLES `article_likes` WRITE;
/*!40000 ALTER TABLE `article_likes` DISABLE KEYS */;
INSERT INTO `article_likes` VALUES (5,11,1,0);
/*!40000 ALTER TABLE `article_likes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `articles`
--

DROP TABLE IF EXISTS `articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `articles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int DEFAULT NULL,
  `created_by` bigint DEFAULT NULL COMMENT '创建人ID',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` bigint DEFAULT NULL COMMENT '修改人ID',
  `deleted_on` int unsigned DEFAULT '0',
  `state` tinyint unsigned DEFAULT '0' COMMENT '0: 草稿,1: 待审核,2:审核成功,3:审核失败',
  `views` int unsigned DEFAULT '0' COMMENT '浏览量',
  PRIMARY KEY (`id`),
  KEY `fk_tag_id` (`tag_id`),
  CONSTRAINT `fk_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb3 COMMENT='文章管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articles`
--

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;
INSERT INTO `articles` VALUES (11,1,'作为测试更改','esse fugiat occaecat ea irure','consequat',1742534805,1,1743077667,1,0,2,0),(12,4,'好些鸟笼对于从哎哟嗯如何箱子','in magna Excepteur ut','ea aliqua non in',1742554307,1,0,0,0,2,0),(13,1,'剪哎呀由穿过','dolore','in',1742783955,21,0,0,0,2,0),(15,4,'起司粉烤蛇肉','南去最员况划进候线。六号结内同。公该土反能类。','正石义己除志程准话争。重话情。发没角没使段。与团确认次分。车采本。',1742902218,2,0,0,0,2,0),(16,4,'奶油生菜和鸡肉饼','也教整今。义见也记需除学车形。方第系须历白。县只比儿。月率型快属这们见。关属后算拉建口。','发空在处。',1742902218,2,0,0,0,2,0),(17,4,'红烧肉','置格又报现养实海性细。细速问建明通前。划使育基装。结管备支世。最验运。内规热确此。','我车间消种技。',1742902304,2,0,0,0,2,0),(18,3,'柠檬和西瓜挞','约知他。动备车收那步东。','元米型求向。然中何上来管。厂细务几争己。',1742902304,2,0,0,0,2,0),(19,1,'芙蓉蟹','将包型更类。象由局铁层叫变产儿。住表想取便放此。','眼张商育响叫持圆。约直例两引先候。决式他际气调务。\n西回精较活志。形目而增受速精多适去。之查好北化直导划。\n性不小从支放月完。世而然支离际本特。只般门别一。',1742902348,2,0,0,0,2,0),(29,4,'瓜类沙拉','民支点叫查类物新。越文号。水受子红难阶三。们委西。','来不及存以又种响只。',1742902439,2,0,0,0,0,0),(30,3,'心雨','放克','情声则。',1742906377,1,0,0,0,0,0),(31,3,'米兰的小铁匠','非音乐','难铁各律。',1742906401,1,0,0,0,0,0),(32,3,'枫','蓝调','几便间利立。',1742906497,1,0,0,0,0,0),(33,12,'手写从前','世界','达周响县还多权。',1742907019,1,0,0,0,0,0),(34,1,'简单的爱','经典','水般义。',1742907073,1,0,0,0,0,0),(35,16,'阳光宅男','拉丁','想眼县。',1742907073,1,0,0,0,0,0),(36,12,'可爱女人','舞台与银幕','边规两知常规治构数却。',1742907073,1,0,0,0,0,0),(37,16,'心雨','嘻哈','领次铁才类品标土。',1742907073,1,0,0,0,0,0),(38,23,'周大侠','舞台与银幕','联原及万报同声资定。',1742907074,1,0,0,0,0,0),(39,21,'红颜如霜','重金属','者导断六信。',1742907074,1,0,0,0,0,0),(40,3,'断了的弦','经典','经米区。',1742907074,1,0,0,0,0,0),(41,21,'美人鱼','经典','给难备头圆各非近义且。',1742907074,1,0,0,0,0,0),(42,14,'听妈妈的话','世界','系前直土。',1742907074,1,0,0,0,0,0),(43,3,'开不了口','牙买加','不消农表太时开张。',1742907074,1,0,0,0,0,0);
/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int unsigned NOT NULL,
  `user_id` int unsigned NOT NULL,
  `content` text NOT NULL,
  `parent_id` int unsigned DEFAULT '0',
  `created_on` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `article_id` (`article_id`),
  KEY `user_id` (`user_id`),
  KEY `comments_ibfk_3` (`parent_id`),
  CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`),
  CONSTRAINT `comments_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (4,11,1,'不错哟',0,1742898941),(5,11,1,'thanks',4,1742899006);
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tags` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb3 COMMENT='文章标签管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES (1,'技术',1633072800,'admin',1633072800,'admin',0,1),(3,'旅行',1633072800,'admin',1633072800,'admin',0,1),(4,'美食',1633072800,'admin',1633072800,'admin',0,1),(12,'边境牧羊犬',1742813350,'laborum in ad',1742814650,'elit',0,59),(14,'肇天娇',1742902348,'aute dolor id labore',0,'',0,0),(15,'费玲',1742902438,'enim deserunt',0,'',0,1),(16,'狂呈轩',1742902438,'magna deserunt ullamco occaecat',0,'',0,0),(17,'謇浩晨',1742902438,'ex proident reprehenderit veniam deserunt',0,'',0,0),(18,'邢桂兰',1742902439,'in ullamco sunt anim occaecat',0,'',0,1),(19,'满国良',1742902439,'eu dolore nostrud ullamco',0,'',0,0),(20,'邶建华',1742902439,'ut in culpa aute',0,'',0,1),(21,'缑蒙',1742902439,'sunt mollit sed Lorem tempor',0,'',0,1),(22,'鲜颖',1742902439,'ad',0,'',0,0),(23,'犹颖',1742902439,'sed Excepteur',0,'',0,0),(24,'麴霞',1742902439,'commodo Ut ipsum eu',0,'',0,1);
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `role` enum('user','author','admin') DEFAULT 'user',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'wc','$2a$10$eWTDAIkuUjHurVZVWzySqeIPHJJcJWjYwuEp3hnHOPfoFGNFpNyVi','admin'),(2,'test1','$2a$10$IQL9tuHkGB9WXexaD/s0GOLBFTtii.UM2ZcCeFKoJnF1C/g6.vLpO','author');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-03-28 15:02:04
