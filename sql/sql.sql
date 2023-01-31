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
  PRIMARY KEY (`cid`),
  KEY `comment_user_key_IDX` (`user_key`) USING BTREE,
  KEY `comment_video_key_IDX` (`video_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
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
  `user_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '唯一编号',
  `version` int DEFAULT NULL,
  PRIMARY KEY (`uid`),
  KEY `user_user_key_IDX` (`user_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (28,'2023-01-30 15:54:17','2023-01-30 15:54:17',NULL,'paidx0@qq.com','bc575fe0557c046f433d556612b3e564',0,0,'f56cd82b-9a7d-489e-8636-49abe939b75b',1);
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
  PRIMARY KEY (`id`),
  KEY `userfocuson_user_key_IDX` (`user_key`) USING BTREE,
  KEY `userfocuson_to_user_key_IDX` (`to_user_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户关注列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userfocuson`
--

LOCK TABLES `userfocuson` WRITE;
/*!40000 ALTER TABLE `userfocuson` DISABLE KEYS */;
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
  PRIMARY KEY (`id`),
  KEY `userlikevideo_user_key_IDX` (`user_key`) USING BTREE,
  KEY `userlikevideo_video_key_IDX` (`video_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户点赞视频绑定';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userlikevideo`
--

LOCK TABLES `userlikevideo` WRITE;
/*!40000 ALTER TABLE `userlikevideo` DISABLE KEYS */;
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
  PRIMARY KEY (`id`),
  KEY `userpulishvideo_user_key_IDX` (`user_key`) USING BTREE,
  KEY `userpulishvideo_video_key_IDX` (`video_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户发布视频绑定';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userpulishvideo`
--

LOCK TABLES `userpulishvideo` WRITE;
/*!40000 ALTER TABLE `userpulishvideo` DISABLE KEYS */;
INSERT INTO `userpulishvideo` VALUES (39,'2023-01-31 14:39:46','2023-01-31 14:39:46',NULL,'f56cd82b-9a7d-489e-8636-49abe939b75b','98d84bda-daea-4e35-98cf-a83b7be62d08'),(40,'2023-01-31 14:43:42','2023-01-31 14:43:42',NULL,'f56cd82b-9a7d-489e-8636-49abe939b75b','346169a0-ac8e-42ae-ad08-2c851a90fdf7'),(41,'2023-01-31 14:45:11','2023-01-31 14:45:11',NULL,'f56cd82b-9a7d-489e-8636-49abe939b75b','f6896623-226e-4621-8cf6-e5630ce8afdd'),(42,'2023-01-31 14:55:52','2023-01-31 14:55:52',NULL,'f56cd82b-9a7d-489e-8636-49abe939b75b','0b01f56b-274e-4547-a7a4-b42f891df7fc'),(43,'2023-01-31 14:59:05','2023-01-31 14:59:05',NULL,'f56cd82b-9a7d-489e-8636-49abe939b75b','b3f3f3d5-64ed-4b6a-bff3-967a6630cbdd');
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
  `version` int DEFAULT NULL COMMENT '乐观锁',
  PRIMARY KEY (`vid`),
  KEY `video_video_key_IDX` (`video_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `video`
--

LOCK TABLES `video` WRITE;
/*!40000 ALTER TABLE `video` DISABLE KEYS */;
INSERT INTO `video` VALUES (41,'2023-01-31 14:45:11','2023-01-31 14:45:11',NULL,'3','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1675147510-bear.mp4','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1675147510-bear.mp4.png',0,0,'bear.mp4','.mp4','f6896623-226e-4621-8cf6-e5630ce8afdd',1),(42,'2023-01-31 14:55:52','2023-01-31 14:55:52',NULL,'4','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1675148151-bear.mp4','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1675148150-bear.mp4.png',0,0,'bear.mp4','.mp4','0b01f56b-274e-4547-a7a4-b42f891df7fc',1),(43,'2023-01-31 14:59:05','2023-01-31 14:59:05',NULL,'6','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1675148344-bear.mp4','http://rorwnk4kz.hn-bkt.clouddn.com/uploads/1675148343-bear.mp4.png',0,0,'bear.mp4','.mp4','b3f3f3d5-64ed-4b6a-bff3-967a6630cbdd',1);
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

-- Dump completed on 2023-01-31 15:03:23
