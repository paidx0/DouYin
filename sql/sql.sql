-- MySQL dump 10.13  Distrib 8.0.32, for macos12.6 (x86_64)
--
-- Host: localhost    Database: DouYin
-- ------------------------------------------------------
-- Server version	8.0.32

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
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (25,'2023-02-23 16:31:38','2023-02-23 16:31:38',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','e713ee21-0cbe-47d2-86a2-3f6073d3b4c0','牛逼哇'),(26,'2023-02-23 16:31:55','2023-02-23 16:31:55',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','e713ee21-0cbe-47d2-86a2-3f6073d3b4c0','666'),(27,'2023-02-24 00:38:01','2023-02-24 00:38:01',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','e6e8d2ab-e3cf-408a-9b56-9917559afcd4','666'),(28,'2023-02-24 00:38:34','2023-02-24 00:38:34',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','fc6511e6-d077-4cff-ba2a-eaa1fb51c394','2131'),(29,'2023-02-24 00:40:40','2023-02-24 00:40:40',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','fc6511e6-d077-4cff-ba2a-eaa1fb51c394','哈哈哈哈');
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
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (34,'2023-02-22 15:55:10','2023-02-22 15:55:10',NULL,'aaa@qq.com','e10adc3949ba59abbe56e057f20f883e',1,1,'015af569-bb7a-4c1f-8293-06d2440572d5',1),(35,'2023-02-22 15:55:31','2023-02-22 15:55:31',NULL,'bbb@qq.com','e10adc3949ba59abbe56e057f20f883e',0,1,'f4ec4593-2504-49ba-8700-3f482fd113fa',1),(36,'2023-02-22 15:55:37','2023-02-22 15:55:37',NULL,'ccc@qq.com','e10adc3949ba59abbe56e057f20f883e',0,0,'629cc846-3ec2-4798-8f84-9382a2494845',1),(37,'2023-02-22 15:55:44','2023-02-22 15:55:44',NULL,'ddd@qq.com','e10adc3949ba59abbe56e057f20f883e',0,1,'484c2c2e-0a48-4d40-9350-98f671523b1c',1),(38,'2023-02-22 15:55:58','2023-02-22 15:55:58',NULL,'paidx0@qq.com','e10adc3949ba59abbe56e057f20f883e',3,0,'9b0de624-549a-4b44-9f45-1bc3934afe03',1),(39,'2023-02-23 16:30:03','2023-02-23 16:30:03',NULL,'test@qq.com','e10adc3949ba59abbe56e057f20f883e',0,1,'37759cb4-3f1e-4a06-b60e-8b1df6afccda',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户关注列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userfocuson`
--

LOCK TABLES `userfocuson` WRITE;
/*!40000 ALTER TABLE `userfocuson` DISABLE KEYS */;
INSERT INTO `userfocuson` VALUES (23,'2023-02-23 16:19:43','2023-02-23 16:19:43',NULL,'015af569-bb7a-4c1f-8293-06d2440572d5','629cc846-3ec2-4798-8f84-9382a2494845',1),(24,'2023-02-23 16:33:23','2023-02-23 16:33:23',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','37759cb4-3f1e-4a06-b60e-8b1df6afccda',1),(25,'2023-02-24 00:38:26','2023-02-24 00:38:26',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','f4ec4593-2504-49ba-8700-3f482fd113fa',1),(26,'2023-02-24 00:41:14','2023-02-24 00:41:14','2023-02-24 00:41:21','9b0de624-549a-4b44-9f45-1bc3934afe03','484c2c2e-0a48-4d40-9350-98f671523b1c',1),(27,'2023-02-24 00:41:24','2023-02-24 00:41:24',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','484c2c2e-0a48-4d40-9350-98f671523b1c',1),(28,'2023-02-24 00:41:44','2023-02-24 00:41:44',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','015af569-bb7a-4c1f-8293-06d2440572d5',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户点赞视频绑定';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userlikevideo`
--

LOCK TABLES `userlikevideo` WRITE;
/*!40000 ALTER TABLE `userlikevideo` DISABLE KEYS */;
INSERT INTO `userlikevideo` VALUES (43,'2023-02-23 16:19:28','2023-02-23 16:19:28',NULL,'015af569-bb7a-4c1f-8293-06d2440572d5','a7c6d7d8-5924-45ef-b35a-cfc26749291b',1),(44,'2023-02-23 16:32:26','2023-02-23 16:32:26',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','e713ee21-0cbe-47d2-86a2-3f6073d3b4c0',1),(45,'2023-02-24 00:37:50','2023-02-24 00:37:50',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','e6e8d2ab-e3cf-408a-9b56-9917559afcd4',1),(46,'2023-02-24 00:38:08','2023-02-24 00:38:08',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','b84fecf0-a960-4de8-a017-95f82f294740',1),(47,'2023-02-24 00:38:28','2023-02-24 00:38:28','2023-02-24 00:40:32','9b0de624-549a-4b44-9f45-1bc3934afe03','fc6511e6-d077-4cff-ba2a-eaa1fb51c394',1),(48,'2023-02-24 00:38:46','2023-02-24 00:38:46',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','377eb87d-7c1b-4196-b56e-02ce67659bcd',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户发布视频绑定';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userpulishvideo`
--

LOCK TABLES `userpulishvideo` WRITE;
/*!40000 ALTER TABLE `userpulishvideo` DISABLE KEYS */;
INSERT INTO `userpulishvideo` VALUES (46,'2023-02-22 21:28:38','2023-02-22 21:28:38',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','c96d89d2-85a1-4411-bfa1-ebe29a5ddd67'),(47,'2023-02-22 21:30:11','2023-02-22 21:30:11',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','fff1495c-b49d-477d-ad73-b840df4305a6'),(48,'2023-02-22 21:30:54','2023-02-22 21:30:54',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','c6239aa2-084a-4c81-b9e0-af16eedcb498'),(49,'2023-02-22 21:31:09','2023-02-22 21:31:09',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','4ba5171b-4b2c-4288-8ced-dbf2a9ca318d'),(50,'2023-02-22 21:31:23','2023-02-22 21:31:23',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','b618f289-ec42-460f-be33-d7d26ec50aab'),(51,'2023-02-22 21:31:38','2023-02-22 21:31:38',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','eb9094c6-ed6c-492e-a229-605e6be8c92f'),(52,'2023-02-22 21:31:52','2023-02-22 21:31:52',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','ae965ff7-d9c1-44b5-ba1a-f9e51aad7bf5'),(53,'2023-02-22 21:32:06','2023-02-22 21:32:06',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','6db6dc86-36b6-4e03-bcdb-340bfc02b498'),(54,'2023-02-22 21:32:21','2023-02-22 21:32:21',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','897473ce-2390-4a49-8426-a93053827161'),(55,'2023-02-22 21:32:35','2023-02-22 21:32:35',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','eadde30b-5b0a-4b2a-9c66-003b4d126175'),(56,'2023-02-22 21:32:51','2023-02-22 21:32:51',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','22f000d1-b42d-420b-af6d-5f44246751c8'),(57,'2023-02-22 21:33:08','2023-02-22 21:33:08',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','98497b34-8393-46ea-a6ad-dc5f7227ae63'),(58,'2023-02-22 21:35:07','2023-02-22 21:35:07',NULL,'015af569-bb7a-4c1f-8293-06d2440572d5','11e2b538-1afb-4f23-a768-69314b9975d0'),(59,'2023-02-22 21:35:23','2023-02-22 21:35:23',NULL,'015af569-bb7a-4c1f-8293-06d2440572d5','a2f5e9ee-ad04-4886-97fc-4cce50157c8f'),(60,'2023-02-22 21:36:13','2023-02-22 21:36:13',NULL,'015af569-bb7a-4c1f-8293-06d2440572d5','977be835-52d5-4ef6-9038-afb08abb66ad'),(61,'2023-02-22 21:36:30','2023-02-22 21:36:30',NULL,'015af569-bb7a-4c1f-8293-06d2440572d5','7386002f-1e04-47b3-b081-9d66dc17ec0b'),(62,'2023-02-22 21:40:35','2023-02-22 21:40:35',NULL,'015af569-bb7a-4c1f-8293-06d2440572d5','d29fce69-1872-4055-8ad7-d74c6e0e1cc2'),(63,'2023-02-22 21:40:53','2023-02-22 21:40:53',NULL,'015af569-bb7a-4c1f-8293-06d2440572d5','4c712653-26b9-4b42-a4ed-d40c6397d628'),(64,'2023-02-22 21:42:51','2023-02-22 21:42:51',NULL,'f4ec4593-2504-49ba-8700-3f482fd113fa','5c8facec-baff-4cb8-a5a6-6fd230f283dc'),(65,'2023-02-22 21:43:45','2023-02-22 21:43:45',NULL,'f4ec4593-2504-49ba-8700-3f482fd113fa','383ef28b-17ee-4f9b-bc5a-d6486901cc88'),(66,'2023-02-22 21:44:01','2023-02-22 21:44:01',NULL,'f4ec4593-2504-49ba-8700-3f482fd113fa','091500fb-a481-43d0-b3db-53afa355ff43'),(67,'2023-02-22 21:44:19','2023-02-22 21:44:19',NULL,'f4ec4593-2504-49ba-8700-3f482fd113fa','7b167072-dce2-44a5-a4c3-0ac41875cae5'),(68,'2023-02-22 21:44:35','2023-02-22 21:44:35',NULL,'f4ec4593-2504-49ba-8700-3f482fd113fa','377eb87d-7c1b-4196-b56e-02ce67659bcd'),(69,'2023-02-22 21:44:56','2023-02-22 21:44:56',NULL,'f4ec4593-2504-49ba-8700-3f482fd113fa','fc6511e6-d077-4cff-ba2a-eaa1fb51c394'),(70,'2023-02-22 21:45:56','2023-02-22 21:45:56',NULL,'629cc846-3ec2-4798-8f84-9382a2494845','54a1559e-6c04-4a3e-b156-110f22a77948'),(71,'2023-02-22 21:46:10','2023-02-22 21:46:10',NULL,'629cc846-3ec2-4798-8f84-9382a2494845','26d205ee-4f1c-4a5c-9f36-a444243cfe25'),(72,'2023-02-22 21:46:42','2023-02-22 21:46:42',NULL,'629cc846-3ec2-4798-8f84-9382a2494845','b84fecf0-a960-4de8-a017-95f82f294740'),(73,'2023-02-22 21:46:56','2023-02-22 21:46:56',NULL,'629cc846-3ec2-4798-8f84-9382a2494845','a7c6d7d8-5924-45ef-b35a-cfc26749291b'),(74,'2023-02-22 21:47:28','2023-02-22 21:47:28',NULL,'484c2c2e-0a48-4d40-9350-98f671523b1c','e6e8d2ab-e3cf-408a-9b56-9917559afcd4'),(75,'2023-02-22 21:48:02','2023-02-22 21:48:02',NULL,'484c2c2e-0a48-4d40-9350-98f671523b1c','e713ee21-0cbe-47d2-86a2-3f6073d3b4c0'),(79,'2023-02-24 00:40:09','2023-02-24 00:40:09',NULL,'9b0de624-549a-4b44-9f45-1bc3934afe03','dc4e1232-8729-4fc6-999d-1a06b6106fd9');
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
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `video`
--

LOCK TABLES `video` WRITE;
/*!40000 ALTER TABLE `video` DISABLE KEYS */;
INSERT INTO `video` VALUES (46,'2023-02-22 21:28:38','2023-02-22 21:28:38',NULL,'视频1','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072517-1.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072517-1.mp4.png',0,0,'1.mp4','.mp4','c96d89d2-85a1-4411-bfa1-ebe29a5ddd67',1),(47,'2023-02-22 21:30:11','2023-02-22 21:30:11',NULL,'视频2','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072610-2.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072610-2.mp4.png',0,0,'2.mp4','.mp4','fff1495c-b49d-477d-ad73-b840df4305a6',1),(48,'2023-02-22 21:30:54','2023-02-22 21:30:54',NULL,'视频3','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072653-3.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072653-3.mp4.png',0,0,'3.mp4','.mp4','c6239aa2-084a-4c81-b9e0-af16eedcb498',1),(49,'2023-02-22 21:31:09','2023-02-22 21:31:09',NULL,'视频4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072668-4.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072668-4.mp4.png',0,0,'4.mp4','.mp4','4ba5171b-4b2c-4288-8ced-dbf2a9ca318d',1),(50,'2023-02-22 21:31:23','2023-02-22 21:31:23',NULL,'视频5','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072683-5.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072682-5.mp4.png',0,0,'5.mp4','.mp4','b618f289-ec42-460f-be33-d7d26ec50aab',1),(51,'2023-02-22 21:31:38','2023-02-22 21:31:38',NULL,'视频6','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072696-6.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072696-6.mp4.png',0,0,'6.mp4','.mp4','eb9094c6-ed6c-492e-a229-605e6be8c92f',1),(52,'2023-02-22 21:31:52','2023-02-22 21:31:52',NULL,'视频7','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072712-7.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072712-7.mp4.png',0,0,'7.mp4','.mp4','ae965ff7-d9c1-44b5-ba1a-f9e51aad7bf5',1),(53,'2023-02-22 21:32:06','2023-02-22 21:32:06',NULL,'视频8','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072726-8.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072725-8.mp4.png',0,0,'8.mp4','.mp4','6db6dc86-36b6-4e03-bcdb-340bfc02b498',1),(54,'2023-02-22 21:32:21','2023-02-22 21:32:21',NULL,'视频9','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072741-9.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072741-9.mp4.png',0,0,'9.mp4','.mp4','897473ce-2390-4a49-8426-a93053827161',1),(55,'2023-02-22 21:32:35','2023-02-22 21:32:35',NULL,'视频10','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072754-10.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072753-10.mp4.png',0,0,'10.mp4','.mp4','eadde30b-5b0a-4b2a-9c66-003b4d126175',1),(56,'2023-02-22 21:32:51','2023-02-22 21:32:51',NULL,'视频11','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072771-11.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072771-11.mp4.png',0,0,'11.mp4','.mp4','22f000d1-b42d-420b-af6d-5f44246751c8',1),(57,'2023-02-22 21:33:08','2023-02-22 21:33:08',NULL,'视频12','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072787-12.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072787-12.mp4.png',0,0,'12.mp4','.mp4','98497b34-8393-46ea-a6ad-dc5f7227ae63',1),(58,'2023-02-22 21:35:07','2023-02-22 21:35:07',NULL,'视频13','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072906-13.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072906-13.mp4.png',0,0,'13.mp4','.mp4','11e2b538-1afb-4f23-a768-69314b9975d0',1),(59,'2023-02-22 21:35:23','2023-02-22 21:35:23',NULL,'视频14','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072923-14.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072922-14.mp4.png',0,0,'14.mp4','.mp4','a2f5e9ee-ad04-4886-97fc-4cce50157c8f',1),(60,'2023-02-22 21:36:13','2023-02-22 21:36:13',NULL,'视频15','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072973-15.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072973-15.mp4.png',0,0,'15.mp4','.mp4','977be835-52d5-4ef6-9038-afb08abb66ad',1),(61,'2023-02-22 21:36:30','2023-02-22 21:36:30',NULL,'视频16','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072990-16.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677072989-16.mp4.png',0,0,'16.mp4','.mp4','7386002f-1e04-47b3-b081-9d66dc17ec0b',1),(62,'2023-02-22 21:40:35','2023-02-22 21:40:35',NULL,'视频17','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073234-17.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073233-17.mp4.png',0,0,'17.mp4','.mp4','d29fce69-1872-4055-8ad7-d74c6e0e1cc2',1),(63,'2023-02-22 21:40:53','2023-02-22 21:40:53',NULL,'视频18','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073252-18.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073252-18.mp4.png',0,0,'18.mp4','.mp4','4c712653-26b9-4b42-a4ed-d40c6397d628',1),(64,'2023-02-22 21:42:50','2023-02-22 21:42:50',NULL,'视频19','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073369-19.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073369-19.mp4.png',0,0,'19.mp4','.mp4','5c8facec-baff-4cb8-a5a6-6fd230f283dc',1),(65,'2023-02-22 21:43:45','2023-02-22 21:43:45',NULL,'视频20','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073424-20.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073423-20.mp4.png',0,0,'20.mp4','.mp4','383ef28b-17ee-4f9b-bc5a-d6486901cc88',1),(66,'2023-02-22 21:44:01','2023-02-22 21:44:01',NULL,'视频21','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073440-21.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073439-21.mp4.png',0,0,'21.mp4','.mp4','091500fb-a481-43d0-b3db-53afa355ff43',1),(67,'2023-02-22 21:44:19','2023-02-22 21:44:19',NULL,'视频22','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073459-22.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073459-22.mp4.png',0,0,'22.mp4','.mp4','7b167072-dce2-44a5-a4c3-0ac41875cae5',1),(68,'2023-02-22 21:44:35','2023-02-22 21:44:35',NULL,'视频23','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073474-23.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073474-23.mp4.png',1,0,'23.mp4','.mp4','377eb87d-7c1b-4196-b56e-02ce67659bcd',1),(69,'2023-02-22 21:44:56','2023-02-22 21:44:56',NULL,'视频24','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073493-24.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073492-24.mp4.png',0,2,'24.mp4','.mp4','fc6511e6-d077-4cff-ba2a-eaa1fb51c394',1),(70,'2023-02-22 21:45:56','2023-02-22 21:45:56',NULL,'视频25','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073555-25.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073554-25.mp4.png',0,0,'25.mp4','.mp4','54a1559e-6c04-4a3e-b156-110f22a77948',1),(71,'2023-02-22 21:46:10','2023-02-22 21:46:10',NULL,'视频26','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073569-27.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073569-27.mp4.png',0,0,'27.mp4','.mp4','26d205ee-4f1c-4a5c-9f36-a444243cfe25',1),(72,'2023-02-22 21:46:42','2023-02-22 21:46:42',NULL,'视频27','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073601-27.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073601-27.mp4.png',1,0,'27.mp4','.mp4','b84fecf0-a960-4de8-a017-95f82f294740',1),(73,'2023-02-22 21:46:56','2023-02-22 21:46:56',NULL,'视频28','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073616-28.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073616-28.mp4.png',1,0,'28.mp4','.mp4','a7c6d7d8-5924-45ef-b35a-cfc26749291b',1),(74,'2023-02-22 21:47:28','2023-02-22 21:47:28',NULL,'视频29','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073648-29.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073648-29.mp4.png',1,1,'29.mp4','.mp4','e6e8d2ab-e3cf-408a-9b56-9917559afcd4',1),(75,'2023-02-22 21:48:02','2023-02-22 21:48:02',NULL,'视频30','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073680-30.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677073680-30.mp4.png',1,2,'30.mp4','.mp4','e713ee21-0cbe-47d2-86a2-3f6073d3b4c0',1),(79,'2023-02-24 00:40:09','2023-02-24 00:40:09',NULL,'测试','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677170409-31.mp4','http://rqhhp7lsu.hn-bkt.clouddn.com/uploads/1677170408-31.mp4.png',0,0,'31.mp4','.mp4','dc4e1232-8729-4fc6-999d-1a06b6106fd9',1);
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

-- Dump completed on 2023-02-24  2:35:06
