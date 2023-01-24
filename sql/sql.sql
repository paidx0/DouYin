-- MySQL dump 10.13  Distrib 8.0.31, for macos12.6 (x86_64)
--
-- Host: 127.0.0.1    Database: DouYin
-- ------------------------------------------------------
-- Server version	8.0.31

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
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment` (
  `cid` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `user_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户唯一编号',
  `video_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '视频唯一编号',
  `comment_text` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '评论内容',
  PRIMARY KEY (`cid`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (8,'2023-01-24 14:06:27','2023-01-24 14:06:27','2023-01-24 14:07:36','2cb6307b-8fd8-401d-aa89-f515dc45077a','3e1fee56-f903-45c7-88ba-8ab6ac8a2b3e','帅！！！'),(9,'2023-01-24 14:06:46','2023-01-24 14:06:46','2023-01-24 14:11:35','2cb6307b-8fd8-401d-aa89-f515dc45077a','3e1fee56-f903-45c7-88ba-8ab6ac8a2b3e','牛逼'),(10,'2023-01-24 14:10:49','2023-01-24 14:10:49',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','3e1fee56-f903-45c7-88ba-8ab6ac8a2b3e','牛逼'),(11,'2023-01-24 14:12:27','2023-01-24 14:12:27',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','3e1fee56-f903-45c7-88ba-8ab6ac8a2b3e','牛逼'),(12,'2023-01-24 14:12:27','2023-01-24 14:12:27',NULL,'a8278467-187a-4598-9806-c1308d76ebae','3e1fee56-f903-45c7-88ba-8ab6ac8a2b3e','牛逼'),(13,'2023-01-24 14:12:27','2023-01-24 14:12:27',NULL,'778d02e1-a26e-4a53-8913-fc225eea9317','3e1fee56-f903-45c7-88ba-8ab6ac8a2b3e','牛逼');
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `uid` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `username` varchar(50) DEFAULT NULL COMMENT '用户名',
  `password` varchar(50) DEFAULT NULL COMMENT '密码',
  `follow_count` int DEFAULT '0' COMMENT '关注总数',
  `follower_count` int DEFAULT '0' COMMENT '粉丝总数',
  `like_total` int DEFAULT '0' COMMENT '点赞总数',
  `user_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '唯一编号',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'2023-01-16 13:32:37','2023-01-16 13:32:37',NULL,'test','cc03e747a6afbbcbf8be7668acfebee5',0,0,0,'778d02e1-a26e-4a53-8913-fc225eea9317'),(2,'2023-01-16 13:41:15','2023-01-16 13:41:15',NULL,'paidx0','bc575fe0557c046f433d556612b3e564',20,10,5,'2cb6307b-8fd8-401d-aa89-f515dc45077a'),(3,'2023-01-16 14:23:27','2023-01-16 14:23:27',NULL,'paidx','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'a8278467-187a-4598-9806-c1308d76ebae'),(4,'2023-01-18 13:05:57','2023-01-18 13:05:57',NULL,'1213','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'bfbe12e3-5100-45cd-a445-df44d263b6a0'),(5,'2023-01-18 13:08:05','2023-01-18 13:08:05',NULL,'12133','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'61e5291b-6bad-4659-8fcd-b0117d5868a0'),(6,'2023-01-18 13:08:52','2023-01-18 13:08:52',NULL,'121334','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'f0b27d50-c52e-44eb-ba50-dc563c045788'),(7,'2023-01-18 13:35:40','2023-01-18 13:35:40',NULL,'12133445','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'dce57a10-b20f-4914-82c7-455f5a52570b'),(8,'2023-01-18 13:37:26','2023-01-18 13:37:26',NULL,'121334451','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'70e51eb8-3b25-4ce7-be2d-d57556685313'),(9,'2023-01-18 13:40:27','2023-01-18 13:40:27',NULL,'1213344511','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'2ea203ad-08b9-4ef6-8914-0726f4583380'),(10,'2023-01-18 13:41:19','2023-01-18 13:41:19',NULL,'12133445111','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'8c634e90-9d48-47ea-85df-f6dc7d784985'),(11,'2023-01-18 13:43:34','2023-01-18 13:43:34',NULL,'121334451112','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'9fa1628f-033b-459c-af3c-143f31d155fa'),(12,'2023-01-18 13:44:28','2023-01-18 13:44:28',NULL,'121334451112a','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'9b2c19e3-8417-437b-8058-d39665ffd1e6'),(13,'2023-01-18 13:45:36','2023-01-18 13:45:36',NULL,'12a','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'93276185-84dd-4359-9202-10056d932948'),(14,'2023-01-18 13:46:25','2023-01-18 13:46:25',NULL,'12a2','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'6f9bc499-9566-471e-b3b6-b1e2aaf8502c'),(15,'2023-01-18 13:48:33','2023-01-18 13:48:33',NULL,'12a23','03daa4871a1dd32048201ccb574f4ea6',0,0,0,'c3c8af18-cc53-4ac2-a49c-73bf6efa33f2'),(16,'2023-01-18 13:59:43','2023-01-18 13:59:43',NULL,'12a234','c4ca4238a0b923820dcc509a6f75849b',0,0,0,'5c7ee1f5-05e0-4b05-84d8-5ceea1cbfadd'),(17,'2023-01-18 14:46:28','2023-01-18 14:46:28',NULL,'12a2342','c4ca4238a0b923820dcc509a6f75849b',0,0,0,'e782e564-56ee-4677-a20e-91b18236d33f'),(18,'2023-01-19 14:11:14','2023-01-19 14:11:14',NULL,'123','c4ca4238a0b923820dcc509a6f75849b',0,0,0,'c45b7378-d82c-4855-9549-309ec18f8e26'),(19,'2023-01-23 15:53:05','2023-01-23 15:53:05',NULL,'qwe','c4ca4238a0b923820dcc509a6f75849b',0,0,0,'d5380b35-f6d9-4922-8659-52a78c8a38d7'),(20,'2023-01-23 15:53:21','2023-01-23 15:53:21',NULL,'qwes','c4ca4238a0b923820dcc509a6f75849b',0,0,0,'77754860-dc86-4031-a445-f2a791756ded'),(21,'2023-01-23 17:39:00','2023-01-23 17:39:00',NULL,'asd','c4ca4238a0b923820dcc509a6f75849b',0,0,0,'624d292e-cf8d-4339-b45d-95a625dd9d77');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userfocuson`
--

DROP TABLE IF EXISTS `userfocuson`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userfocuson` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `user_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户唯一编号',
  `to_user_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '被关注用户唯一编号',
  `is_follow` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户关注列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userfocuson`
--

LOCK TABLES `userfocuson` WRITE;
/*!40000 ALTER TABLE `userfocuson` DISABLE KEYS */;
INSERT INTO `userfocuson` VALUES (1,NULL,NULL,NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','a8278467-187a-4598-9806-c1308d76ebae',1),(2,NULL,NULL,NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','778d02e1-a26e-4a53-8913-fc225eea9317',1);
/*!40000 ALTER TABLE `userfocuson` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userlikevideo`
--

DROP TABLE IF EXISTS `userlikevideo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userlikevideo` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `user_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户唯一编号',
  `video_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '视频唯一编号',
  `is_favorite` tinyint(1) DEFAULT NULL COMMENT '点赞',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户点赞视频绑定';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userlikevideo`
--

LOCK TABLES `userlikevideo` WRITE;
/*!40000 ALTER TABLE `userlikevideo` DISABLE KEYS */;
INSERT INTO `userlikevideo` VALUES (3,'2023-01-18 21:30:12','2023-01-18 21:30:12',NULL,'a8278467-187a-4598-9806-c1308d76ebae','test-2',1),(8,'2023-01-18 21:30:12','2023-01-18 21:30:12',NULL,'a8278467-187a-4598-9806-c1308d76ebae','test-888',1),(17,'2023-01-23 17:24:04','2023-01-23 17:24:04',NULL,'a8278467-187a-4598-9806-c1308d76ebae','test-3',1),(18,'2023-01-23 17:25:04','2023-01-23 17:25:04',NULL,'a8278467-187a-4598-9806-c1308d76ebae','test-123',1),(20,'2023-01-23 17:27:26','2023-01-23 17:27:26',NULL,'a8278467-187a-4598-9806-c1308d76ebae','test-789',1),(22,'2023-01-23 17:27:26','2023-01-23 17:27:26',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','test-3',1),(24,'2023-01-23 22:41:33','2023-01-23 22:41:33',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','test-789',1),(25,'2023-01-23 22:56:49','2023-01-23 22:56:49','2023-01-23 22:58:53','2cb6307b-8fd8-401d-aa89-f515dc45077a','92b24b06-e611-40c1-9c0e-ba3e74ac6164',1),(26,'2023-01-23 22:59:28','2023-01-23 22:59:28',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','3e1fee56-f903-45c7-88ba-8ab6ac8a2b3e',1);
/*!40000 ALTER TABLE `userlikevideo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userpulishvideo`
--

DROP TABLE IF EXISTS `userpulishvideo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userpulishvideo` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `user_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户唯一编号',
  `video_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '视频唯一编号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户发布视频绑定';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userpulishvideo`
--

LOCK TABLES `userpulishvideo` WRITE;
/*!40000 ALTER TABLE `userpulishvideo` DISABLE KEYS */;
INSERT INTO `userpulishvideo` VALUES (1,NULL,NULL,NULL,'61e5291b-6bad-4659-8fcd-b0117d5868a0','test-123'),(2,NULL,NULL,NULL,'778d02e1-a26e-4a53-8913-fc225eea9317','test-456'),(3,NULL,NULL,NULL,'bfbe12e3-5100-45cd-a445-df44d263b6a0','test-789'),(4,NULL,NULL,NULL,'c45b7378-d82c-4855-9549-309ec18f8e26','test-888'),(5,NULL,NULL,NULL,'a8278467-187a-4598-9806-c1308d76ebae','test-000'),(6,NULL,NULL,NULL,'70e51eb8-3b25-4ce7-be2d-d57556685313','test-2'),(7,NULL,NULL,NULL,'a8278467-187a-4598-9806-c1308d76ebae','test-3'),(8,'2023-01-23 21:49:26','2023-01-23 21:49:26',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','be9038ae-5d14-402b-b0e4-ce9ddb967539'),(9,'2023-01-23 21:54:03','2023-01-23 21:54:03',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','3e1fee56-f903-45c7-88ba-8ab6ac8a2b3e'),(10,'2023-01-23 22:41:58','2023-01-23 22:41:58',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','92b24b06-e611-40c1-9c0e-ba3e74ac6164'),(11,'2023-01-24 15:38:21','2023-01-24 15:38:21',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','d7720a86-edd5-46a8-97c7-96fa507d8c13'),(12,'2023-01-24 17:48:39','2023-01-24 17:48:39',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','e965ce5d-7450-4cb5-86f5-9eac80026bd4'),(13,'2023-01-24 17:50:45','2023-01-24 17:50:45',NULL,'2cb6307b-8fd8-401d-aa89-f515dc45077a','124fb6e2-657d-4efd-8725-582b9b082b17');
/*!40000 ALTER TABLE `userpulishvideo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `video`
--

DROP TABLE IF EXISTS `video`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `video` (
  `vid` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `title` varchar(100) DEFAULT NULL COMMENT '视频标题',
  `play_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '视频播放地址',
  `cover_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '视频封面地址',
  `favorite_count` int DEFAULT '0' COMMENT '视频点赞总数',
  `comment_count` int DEFAULT '0' COMMENT '视频评论总数',
  `name` varchar(100) DEFAULT NULL COMMENT '文件名',
  `tag` varchar(100) DEFAULT NULL COMMENT '文件标签',
  `video_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '唯一编号',
  PRIMARY KEY (`vid`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `video`
--

LOCK TABLES `video` WRITE;
/*!40000 ALTER TABLE `video` DISABLE KEYS */;
INSERT INTO `video` VALUES (1,NULL,NULL,NULL,'test','123.qq.com','456.qq.com',10,20,'测试1','demo','test-123'),(2,NULL,NULL,NULL,'test','123','456',10,20,'测试7','demo','test-3'),(3,NULL,NULL,NULL,'test','123.qq.com','456.qq.com',10,20,'测试2','demo','test-456'),(4,NULL,NULL,NULL,'test','123.qq.com','456.qq.com',10,20,'测试3','demo','test-789'),(5,NULL,NULL,NULL,'test','123.qq.com','456.qq.com',10,20,'测试4','demo','test-000'),(6,NULL,NULL,NULL,'test','123.qq.com','456.qq.com',10,20,'测试5','demo','test-888'),(7,NULL,NULL,NULL,'test','123.qq.com','456.qq.com',10,20,'测试6','demo','test-2'),(8,'2023-01-23 21:49:26','2023-01-23 21:49:26',NULL,'测试投稿视频','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1674481764-a.jpg','',0,0,'a.jpg','.jpg','be9038ae-5d14-402b-b0e4-ce9ddb967539'),(9,'2023-01-23 21:54:03','2023-01-23 21:54:03',NULL,'测试投稿视频','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1674482043-00.png','',0,0,'00.png','.png','3e1fee56-f903-45c7-88ba-8ab6ac8a2b3e'),(10,'2023-01-23 22:41:57','2023-01-23 22:41:57',NULL,'测试投稿视频','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1674484916-00.png','',0,0,'00.png','.png','92b24b06-e611-40c1-9c0e-ba3e74ac6164'),(11,'2023-01-24 15:38:21','2023-01-24 15:38:21',NULL,'分块投稿视频','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1674545900-c.jpg','',0,0,'c.jpg','.jpg','d7720a86-edd5-46a8-97c7-96fa507d8c13'),(12,'2023-01-24 17:48:39','2023-01-24 17:48:39',NULL,'分块投稿视频','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1674553718-c.jpg','',0,0,'c.jpg','.jpg','e965ce5d-7450-4cb5-86f5-9eac80026bd4'),(13,'2023-01-24 17:50:45','2023-01-24 17:50:45',NULL,'分块投稿视频','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1674553845-c.jpg','',0,0,'c.jpg','.jpg','124fb6e2-657d-4efd-8725-582b9b082b17');
/*!40000 ALTER TABLE `video` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'DouYin'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-01-24 18:15:15
