-- MySQL dump 10.13  Distrib 8.0.27, for macos11 (x86_64)
--
-- Host: localhost    Database: devops_service
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
-- Table structure for table `code_registry`
--

DROP DATABASE  IF EXISTS `devops_service`;
CREATE DATABASE `devops_service`;
USE `devops_service`;

DROP TABLE IF EXISTS `code_registry`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `code_registry` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL COMMENT '仓库名称',
  `desc` varchar(256) NOT NULL COMMENT '仓库描述',
  `type` varchar(128) NOT NULL COMMENT '仓库类型',
  `host` varchar(256) NOT NULL COMMENT '仓库地址',
  `token` varchar(128) NOT NULL COMMENT '授权token',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `clone_type` varchar(128) NOT NULL COMMENT '下载代码方式',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `code_registry`
--

LOCK TABLES `code_registry` WRITE;
/*!40000 ALTER TABLE `code_registry` DISABLE KEYS */;
INSERT INTO `code_registry` VALUES (1,'测试仓库','测试仓库测试仓库测试仓库','gitee','https://gitee.com/api','62ce05eed1df2f076ecc83196ff63043','方伟业',1,1670072232,1670151111,'htmlUrl'),(2,'测试仓库-gitlab','测试仓库-gitlab测试仓库-gitlab','gitlab','http://121.5.102.204:8929/api/v4','GX***Fc','方伟业',1,1670072674,NULL,'htmlUrl'),(3,'测试仓库-github','测试仓库-github测试仓库-github','github','https://github.com/api','gh***1l','方伟业',1,1670072720,NULL,'htmlUrl'),(4,'gitee-ps-go','gitee-ps-go','gitee','https://gitee.com/api','d389ff1ff444b671c730a9337deda92b','方伟业',1,1670490530,1670490530,'htmlUrl');
/*!40000 ALTER TABLE `code_registry` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dockerfile_template`
--

DROP TABLE IF EXISTS `dockerfile_template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dockerfile_template` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL COMMENT 'dockerfile名称',
  `desc` varchar(256) NOT NULL COMMENT 'dockerfile描述',
  `template` text NOT NULL COMMENT 'dockerfile模板',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dockerfile_template`
--

LOCK TABLES `dockerfile_template` WRITE;
/*!40000 ALTER TABLE `dockerfile_template` DISABLE KEYS */;
INSERT INTO `dockerfile_template` VALUES (1,'测试','sss','FROM golang:alpine AS build\nENV GOPROXY=https://goproxy.cn,direct\nENV GO111MODULE on\nWORKDIR /go/cache\nADD go.mod .\nADD go.sum .\nRUN go mod download\n\nWORKDIR /go/build\nADD . .\nRUN GOOS=linux CGO_ENABLED=0 go build -ldflags=\"-s -w\" -installsuffix cgo -o entry main.go\nFROM alpine\nEXPOSE {ListenPort}\nWORKDIR /go/build\nCOPY --from=build /go/build/entry /go/build/entry\nADD ./config /go/build/config\nCMD [\"./entry\"]','方伟业',1,1670081785,1671370334);
/*!40000 ALTER TABLE `dockerfile_template` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `environment`
--

DROP TABLE IF EXISTS `environment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `environment` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '环境关键字',
  `name` varchar(128) DEFAULT NULL COMMENT '环境名',
  `description` varchar(128) NOT NULL COMMENT '备注信息',
  `status` tinyint(1) DEFAULT '0' COMMENT '环境状态',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `type` varchar(128) NOT NULL COMMENT '环境名',
  `host` varchar(128) NOT NULL COMMENT '连接地址',
  `namespace` varchar(128) NOT NULL COMMENT '命名空间',
  `token` text NOT NULL COMMENT '连接token',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `image_registry`
--

DROP TABLE IF EXISTS `image_registry`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `image_registry` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL COMMENT '仓库名称',
  `desc` varchar(256) NOT NULL COMMENT '仓库描述',
  `host` varchar(256) NOT NULL COMMENT '仓库地址',
  `username` varchar(128) NOT NULL COMMENT '仓库账号',
  `password` varchar(128) NOT NULL COMMENT '仓库密码',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `history_count` int(11) NOT NULL COMMENT '历史记录数量',
  `protocol` varchar(128) NOT NULL COMMENT '仓库协议',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `network`
--

DROP TABLE IF EXISTS `network`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `network` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `srv_id` int(11) NOT NULL COMMENT '服务id',
  `env_id` int(11) NOT NULL COMMENT '环境id',
  `host` varchar(256) NOT NULL COMMENT 'url',
  `cert` text COMMENT 'tls公钥',
  `key` text COMMENT 'tls密钥',
  `redirect` tinyint(1) DEFAULT '0' COMMENT '是否强制跳转到https',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `host` (`host`),
  KEY `srv_id` (`srv_id`),
  CONSTRAINT `network_ibfk_1` FOREIGN KEY (`srv_id`) REFERENCES `service` (`id`),
  CONSTRAINT `network_ibfk_2` FOREIGN KEY (`srv_id`) REFERENCES `environment` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `pack_log`
--

DROP TABLE IF EXISTS `pack_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pack_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `service_keyword` varchar(128) NOT NULL COMMENT '服务关键字',
  `service_name` varchar(128) NOT NULL COMMENT '服务名称',
  `dockerfile_name` varchar(128) NOT NULL COMMENT 'dockerfile名称',
  `code_registry_name` varchar(128) NOT NULL COMMENT '代码仓库名称',
  `image_registry_id` int(11) NOT NULL COMMENT '镜像仓库id',
  `image_registry_name` varchar(128) NOT NULL COMMENT '镜像仓库名称',
  `clone_type` varchar(128) NOT NULL COMMENT '下载分支还是标签',
  `clone_value` varchar(128) NOT NULL COMMENT '分支或者标签具体值',
  `commit_id` varchar(128) NOT NULL COMMENT '提交代码的commit_id',
  `image_name` varchar(256) NOT NULL COMMENT '生成的镜像id',
  `use_time` int(11) NOT NULL COMMENT '构建时长',
  `desc` text NOT NULL COMMENT '构建详情',
  `is_clear` tinyint(1) DEFAULT '0' COMMENT '是否清理镜像',
  `is_finish` tinyint(1) DEFAULT '0' COMMENT '是否完成构建',
  `status` tinyint(1) DEFAULT NULL COMMENT '构建结果',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `service_keyword` (`service_keyword`,`image_registry_id`,`clone_value`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `release_log`
--

DROP TABLE IF EXISTS `release_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `release_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `service_keyword` varchar(128) NOT NULL COMMENT '服务关键字',
  `service_name` varchar(128) NOT NULL COMMENT '服务名称',
  `image_registry_name` varchar(128) NOT NULL COMMENT '镜像仓库名',
  `image_name` varchar(256) NOT NULL COMMENT '需要构建的镜像名称',
  `use_time` int(11) DEFAULT '0' COMMENT '使用时间',
  `desc` text NOT NULL COMMENT '构建详情',
  `env_keyword` varchar(128) NOT NULL COMMENT '环境keyword',
  `env_name` varchar(128) NOT NULL COMMENT '环境名称',
  `is_finish` tinyint(1) DEFAULT '0' COMMENT '是否完成构建',
  `status` varchar(128) DEFAULT NULL COMMENT '构建状态 true 成功，false 失败',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `service_keyword` (`service_keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `release_template`
--

DROP TABLE IF EXISTS `release_template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `release_template` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL COMMENT '清单名称',
  `type` varchar(128) NOT NULL COMMENT '清单类型',
  `desc` varchar(256) NOT NULL COMMENT '清单描述',
  `template` text NOT NULL COMMENT '清单模板',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `service`
--

DROP TABLE IF EXISTS `service`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `keyword` varchar(128) NOT NULL COMMENT '服务关键字',
  `name` varchar(128) DEFAULT NULL COMMENT '服务名',
  `is_private` tinyint(1) DEFAULT '0' COMMENT '是否私有服务',
  `team_id` int(11) DEFAULT NULL COMMENT '所属部门',
  `release_id` int(11) NOT NULL COMMENT '发布清单模板id',
  `dockerfile_id` int(11) NOT NULL COMMENT 'dockerfile模板id',
  `code_registry_id` int(11) NOT NULL COMMENT '代码仓库id',
  `image_registry_id` int(11) NOT NULL COMMENT '代码仓库id',
  `run_port` int(11) NOT NULL COMMENT '运行端口',
  `listen_port` int(11) NOT NULL COMMENT '监听端口',
  `owner` varchar(256) NOT NULL COMMENT '代码仓库空间',
  `repo` varchar(256) NOT NULL COMMENT '代码仓库名称',
  `description` varchar(256) NOT NULL COMMENT '备注信息',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `replicas` int(11) NOT NULL COMMENT '副本数量',
  `probe_init_delay` int(11) NOT NULL COMMENT '延迟时间',
  `probe_type` varchar(128) NOT NULL COMMENT '副本数量',
  `probe_value` varchar(128) NOT NULL COMMENT '副本数量',
  `run_type` varchar(256) NOT NULL COMMENT '运行模式',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  KEY `release_id` (`release_id`),
  KEY `dockerfile_id` (`dockerfile_id`),
  KEY `code_registry_id` (`code_registry_id`),
  KEY `image_registry_id` (`image_registry_id`),
  CONSTRAINT `service_ibfk_1` FOREIGN KEY (`release_id`) REFERENCES `release_template` (`id`),
  CONSTRAINT `service_ibfk_2` FOREIGN KEY (`dockerfile_id`) REFERENCES `dockerfile_template` (`id`),
  CONSTRAINT `service_ibfk_3` FOREIGN KEY (`code_registry_id`) REFERENCES `code_registry` (`id`),
  CONSTRAINT `service_ibfk_4` FOREIGN KEY (`image_registry_id`) REFERENCES `image_registry` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `service_env`
--

DROP TABLE IF EXISTS `service_env`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_env` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `env_id` int(11) NOT NULL COMMENT '环境id',
  `srv_id` int(11) NOT NULL COMMENT '服务id',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `env_id` (`env_id`,`srv_id`),
  KEY `srv_id` (`srv_id`),
  CONSTRAINT `service_env_ibfk_1` FOREIGN KEY (`env_id`) REFERENCES `environment` (`id`) ON DELETE CASCADE,
  CONSTRAINT `service_env_ibfk_2` FOREIGN KEY (`srv_id`) REFERENCES `service` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_env`
--

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-28 16:33:59
-- MySQL dump 10.13  Distrib 8.0.27, for macos11 (x86_64)
--
-- Host: localhost    Database: devops_configure
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
-- Table structure for table `environment`
--

DROP DATABASE  IF EXISTS `devops_configure`;
CREATE DATABASE `devops_configure`;
USE `devops_configure`;

DROP TABLE IF EXISTS `environment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `environment` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `env_keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '环境关键字',
  `drive` varchar(128) NOT NULL COMMENT '中间件',
  `config` text NOT NULL COMMENT '配置信息',
  `prefix` varchar(128) NOT NULL COMMENT '存储目录',
  `token` varchar(128) NOT NULL COMMENT '获取连接认证token',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `env_keyword` (`env_keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `environment`
--


--
-- Table structure for table `field`
--

DROP TABLE IF EXISTS `field`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `field` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `service_keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `field` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '服务字段',
  `description` varchar(156) NOT NULL COMMENT '字段简介',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_keyword` (`service_keyword`,`field`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;



--
-- Table structure for table `field_value`
--

DROP TABLE IF EXISTS `field_value`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `field_value` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `field_id` int(11) NOT NULL,
  `env_keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `value` text NOT NULL,
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `env_keyword` (`env_keyword`,`field_id`),
  KEY `field_id` (`field_id`),
  CONSTRAINT `field_value_ibfk_1` FOREIGN KEY (`field_id`) REFERENCES `field` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `resource`
--

DROP TABLE IF EXISTS `resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `field` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '系统字段',
  `description` varchar(256) NOT NULL COMMENT '字段简介',
  `child_field` text NOT NULL COMMENT '子字段',
  `type` varchar(128) NOT NULL COMMENT '配置类型',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `field` (`field`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `resource_value`
--

DROP TABLE IF EXISTS `resource_value`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `resource_value` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `resource_id` int(11) NOT NULL,
  `env_keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `value` text NOT NULL,
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `resource_id` (`resource_id`,`env_keyword`),
  CONSTRAINT `resource_value_ibfk_1` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `service_resource`
--

DROP TABLE IF EXISTS `service_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `service_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `service_keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `resource_id` int(11) NOT NULL,
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_keyword` (`service_keyword`,`resource_id`),
  KEY `resource_id` (`resource_id`),
  CONSTRAINT `service_resource_ibfk_1` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `template`
--

DROP TABLE IF EXISTS `template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `template` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `service_keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '服务名称',
  `content` text NOT NULL COMMENT '模板内容',
  `version` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '版本号',
  `is_use` tinyint(1) NOT NULL COMMENT '是否使用',
  `description` varchar(128) NOT NULL COMMENT '版本描述',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `version` (`version`),
  KEY `service_keyword` (`service_keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `template_log`
--

DROP TABLE IF EXISTS `template_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `template_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `service_keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '服务id',
  `env_keyword` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '环境id',
  `config` text NOT NULL COMMENT '配置内容',
  `description` text NOT NULL COMMENT '配置内容',
  `operator` varchar(128) NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `template_log`
--

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-28 16:33:59
-- MySQL dump 10.13  Distrib 8.0.27, for macos11 (x86_64)
--
-- Host: localhost    Database: devops_ums
-- ------------------------------------------------------
-- Server version	5.7.28

--
-- Table structure for table `casbin_rule`
--

DROP DATABASE  IF EXISTS `devops_ums`;
CREATE DATABASE `devops_ums`;
USE `devops_ums`;

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v6` varchar(25) COLLATE utf8_unicode_ci DEFAULT NULL,
  `v7` varchar(25) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`)
) ENGINE=InnoDB AUTO_INCREMENT=270 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (269,'p','test','/ums/user','DELETE','','','','',''),(264,'p','test','/ums/user','POST','','','','',''),(265,'p','test','/ums/user','PUT','','','','',''),(263,'p','test','/ums/user/page','GET','','','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `login_log`
--

DROP TABLE IF EXISTS `login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `login_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户账号',
  `ip` char(32) NOT NULL COMMENT 'IP地址',
  `address` varchar(256) NOT NULL COMMENT '登陆地址',
  `browser` varchar(128) NOT NULL COMMENT '浏览器',
  `device` varchar(128) NOT NULL COMMENT '登录设备',
  `status` tinyint(1) NOT NULL COMMENT '登录状态',
  `code` int(11) NOT NULL COMMENT '错误码',
  `description` varchar(256) NOT NULL COMMENT '登录备注',
  `created_at` int(11) DEFAULT NULL COMMENT '登陆时间',
  PRIMARY KEY (`id`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(300) COLLATE utf8_unicode_ci NOT NULL COMMENT '标题',
  `icon` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '图标',
  `path` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '路径',
  `name` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '菜单名',
  `type` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '菜单类型',
  `permission` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '指令',
  `method` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '接口请求方式',
  `component` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '组件地址',
  `parent_id` int(11) NOT NULL COMMENT '父级菜单ID',
  `weight` int(11) DEFAULT '0' COMMENT '菜单权重',
  `hidden` tinyint(1) DEFAULT '0' COMMENT '是否隐藏',
  `is_frame` tinyint(1) DEFAULT '0' COMMENT '是否新开窗口',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `redirect` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '跳转地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=152 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (1,'系统菜单','menu',NULL,NULL,'R',' ',NULL,NULL,0,0,0,0,'系统创建',0,NULL,NULL,NULL),(2,'首页','s-home','/','P','M',' ',NULL,'Layout',1,99,0,0,'方伟业',1,NULL,1669456611,'/home'),(3,'用户中心','user','system','System','M','','','Layout',1,98,0,0,'方伟业',1,NULL,1669456630,NULL),(4,'用户管理','user','/user','User','M','','','system/user/index',3,0,0,0,'方伟业',1,NULL,1669391679,''),(5,'角色管理','bangzhu','/role','Role','M','','','system/role/index',3,0,0,0,'方伟业',1,NULL,1669391696,''),(6,'部门管理','s-operation','/team','Team','M','','','system/team/index',3,0,0,0,'系统创建',0,NULL,NULL,NULL),(7,'菜单管理','menu','/menu','Menu','M','','','system/menu/index',3,0,0,0,'系统创建',0,NULL,NULL,NULL),(8,'首页','s-home','/home','Home','M','','','home/index',2,0,0,0,'系统创建',0,NULL,NULL,NULL),(9,'获取菜单树','','/ums/menu','','A','system:menu:tree','GET','',7,0,0,0,'系统创建',0,NULL,NULL,NULL),(10,'获取当前用户基本信息','','/ums/user','','A','baseApi','GET','',92,0,0,0,'方伟业',1,NULL,1669029790,NULL),(11,'获取当前用户的菜单树','','/ums/role/menu','','A','baseApi','GET','',92,0,0,0,'方伟业',1,NULL,1669029816,NULL),(12,'新增菜单','','/ums/menu','','A','system:menu:add','POST','',7,0,0,0,'系统创建',0,NULL,NULL,NULL),(13,'修改菜单','','/ums/menu','','A','system:menu:update','PUT','',7,0,0,0,'系统创建',0,NULL,NULL,NULL),(14,'删除菜单','','/ums/menu','','A','system:menu:delete','DELETE','',7,0,0,0,'系统创建',0,NULL,NULL,NULL),(15,'查看分页用户数据','','/ums/user/page','','A','system:user:page','GET','',4,0,0,0,'系统创建',0,NULL,NULL,NULL),(16,'新增用户数据','','/ums/user','','A','system:user:add','POST','',4,0,0,0,'系统创建',0,NULL,NULL,NULL),(17,'修改用户信息','','/ums/user','','A','system:user:update','PUT','',4,0,0,0,'系统创建',0,NULL,NULL,NULL),(18,'删除用户数据','','/ums/user','','A','system:user:delete','DELETE','',4,0,0,0,'方伟业',1,NULL,1669310944,NULL),(19,'基本接口','setting','/baseApi','baseApi','M','baseApi','','',1,100,1,0,'方伟业',1,NULL,1669456640,NULL),(20,'获取角色列表','','/ums/role','','A','system:role:list','GET','',5,0,0,0,'系统创建',0,NULL,NULL,NULL),(21,'新增角色','','/ums/role','','A','system:role:add','POST','',5,0,0,0,'系统创建',0,NULL,NULL,NULL),(22,'修改角色','','/ums/role','','A','system:role:update','PUT','',5,0,0,0,'系统创建',0,NULL,NULL,NULL),(23,'删除角色','','/ums/role','','A','system:role:delete','DELETE','',5,0,0,0,'系统创建',0,NULL,NULL,NULL),(24,'修改角色菜单','','/ums/role/menu','','A','system:role:menu','POST','',63,0,0,0,'方伟业',1,NULL,1661748305,NULL),(25,'获取指定角色的菜单ID','','/ums/role/menu_ids','','A','baseApi','GET','',63,0,0,0,'方伟业',1,NULL,1669028269,NULL),(26,'获取部门树','','/ums/team','','A','baseApi','GET','',92,0,0,0,'方伟业',1,NULL,1669029806,NULL),(27,'新增部门','','/ums/team','','A','system:team:add','POST','',6,0,0,0,'系统创建',0,NULL,NULL,NULL),(28,'更新部门信息','','/ums/team','','A','system:team:update','PUT','',6,0,0,0,'系统创建',0,NULL,NULL,NULL),(29,'删除指定部门','','/ums/team','','A','system:team:delete','DELETE','',6,0,0,0,'系统创建',0,NULL,NULL,NULL),(30,'登陆日志','files','/log','Log','M','','','system/log/index',3,0,0,0,'系统创建',0,NULL,NULL,NULL),(31,'查询登陆日志','','/ums/login/log','','A','system:log:page','GET','',30,0,0,0,'系统创建',0,NULL,NULL,NULL),(33,'配置中心','document','/configure','Configure','M','','','Layout',1,0,0,0,'方伟业',1,1661180576,1661181468,'/configure/environment'),(34,'环境管理','monitor','/service/environment','EnvironmentMgr','M','','','service/environment/index',95,0,0,0,'方伟业',1,1661180760,1670034714,''),(35,'查看环境列表详细信息','','/service/environments','','A','service:environment:all','GET','',34,0,0,0,'方伟业',1,1661181119,1669451961,''),(36,'新增环境','','/service/environment','','A','service:environment:add','POST','',34,0,0,0,'方伟业',1,1661181330,1669451969,''),(37,'修改环境','','/service/environment','','A','service:environment:update','PUT','',34,0,0,0,'方伟业',1,1661181392,1669451977,''),(38,'删除环境','','/service/environment','','A','service:environment:delete','DELETE','',34,0,0,0,'方伟业',1,1661181419,1669451987,''),(39,'服务管理','coin','/service/service','ServiceMgr','M','','','service/service/index',95,0,0,0,'方伟业',1,1661184428,1670034704,''),(41,'新增服务','','/service/service','','A','service:service:add','POST','',39,0,0,0,'方伟业',1,1661184582,1669452049,''),(42,'修改服务','','/service/service','','A','service:service:update','PUT','',39,0,0,0,'方伟业',1,1661184616,1669452056,''),(43,'删除服务','','/service/service','','A','service:service:delete','DELETE','',39,0,0,0,'方伟业',1,1661184640,1669452062,''),(44,'资源管理','set-up','/configure/resource','Resource','M','','','configure/resource/index',33,0,0,0,'方伟业',1,1661180577,1669442409,''),(45,'查看资源字段','','/configure/resource/page','','A','configure:resource:page','GET','',44,0,0,0,'方伟业',1,1661190757,1669443984,''),(46,'新增资源字段','','/configure/resource','','A','configure:resource:add','POST','',44,0,0,0,'方伟业',1,1661190788,1669443997,''),(47,'修改资源字段','','/configure/resource','','A','configure:resource:update','PUT','',44,0,0,0,'方伟业',1,1661190823,1669444047,''),(48,'删除资源字段','','/configure/resource','','A','configure:resource:delete','DELETE','',44,0,0,0,'方伟业',1,1661190856,1669444062,''),(49,'业务字段','c-scale-to-original','/configure/field','ConfigureField','M','','','configure/field/index',33,0,0,0,'方伟业',1,1661190988,1669473763,''),(50,'查看服务字段','','/configure/field/page','','A','configure:field:page','GET','',49,0,0,0,'方伟业',1,1661191032,1669473778,''),(51,'新增服务字段','','/configure/field','','A','configure:field:add','POST','',49,0,0,0,'方伟业',1,1661191063,1669473786,''),(52,'修改服务字段','','/configure/field','','A','configure:field:update','PUT','',49,0,0,0,'方伟业',1,1661191147,1669473794,''),(53,'删除服务字段','','/configure/field','','A','configure:field:delete','DELETE','',49,0,0,0,'方伟业',1,1661191203,1669473803,''),(55,'查询系统字段值配置','','/configure/resource/value','','A','configure:resource_value:query','GET','',44,0,0,0,'方伟业',1,1661311151,1669444164,''),(56,'更新系统字段值配置','','/configure/resource/value','','A','configure:resource_value:update','POST','',44,0,0,0,'方伟业',1,1661311187,1669444169,''),(57,'查询服务值配置','','/configure/field_value','','A','configure:field_value:query','GET','',49,0,0,0,'方伟业',1,1661311513,1669473813,''),(58,'更新服务值配置','','/configure/field_value','','A','configure:field_value:update','POST','',49,0,0,0,'方伟业',1,1661311551,1669473821,''),(60,'查看服务资源','','/service/environment/service','','A','service:service:resource:all','GET','',39,0,0,0,'方伟业',1,1661347980,1669452069,''),(61,'更新服务资源','','/service/environment/service','','A','service:service:resource:update','POST','',39,0,0,0,'方伟业',1,1661348029,1669452079,''),(62,'配置模板','document','/configure/template','ConfigureTemplate','M','','','configure/template/index',33,0,0,0,'方伟业',1,1661357663,1661357663,''),(63,'修改角色菜单','','','','G','','','',5,0,0,0,'方伟业',1,1661748154,1669028281,''),(65,'变更模板配置','','/configure/template','','A','configure:template:add','POST','',62,0,0,0,'方伟业',1,1661761309,1661761309,''),(66,'切换模板版本','','/configure/template','','A','configure:template:update','PUT','',62,0,0,0,'方伟业',1,1661763623,1662198240,''),(68,'获取模板列表','','/configure/templates','','A','configure:template:all','GET','',62,0,0,0,'方伟业',1,1662198286,1662198342,''),(69,'获取指定模板详细配置','','/configure/template','','A','configure:template:info','GET','',62,0,0,0,'方伟业',1,1662198332,1662198352,''),(70,'渲染测试','','/configure/template/parse','','A','configure:template:parse','GET','',62,0,0,0,'方伟业',1,1662198383,1662198383,''),(71,'同步配置','','','','G','configure:config:sync','','',62,0,0,0,'方伟业',1,1662198485,1662199148,''),(72,'对比模板配置','','/configure/config/compare','','A','configure:config:compare','GET','',71,0,0,0,'方伟业',1,1662198522,1662198522,''),(73,'同步配置','','/configure/config/sync','','A','configure:config:sync','POST','',71,0,0,0,'方伟业',1,1662198557,1662199138,''),(74,'配置回归','','','','G','configure:config:rollback','','',62,0,0,0,'方伟业',1,1662198653,1662199178,''),(75,'查看历史版本配置列表','','/configure/config/logs','','A','configure:config:logs','GET','',74,0,0,0,'方伟业',1,1662198701,1662198701,''),(76,'查看详细配置内容','','/configure/config/log','','A','configure:config:log','GET','',74,0,0,0,'方伟业',1,1662198738,1662198738,''),(77,'回归配置','','/configure/config/rollback','','A','configure:config:rollback','POST','',74,0,0,0,'方伟业',1,1662198773,1662198773,''),(78,'查看配置','','/configure/config/driver','','A','configure:config:driver','GET','',62,0,0,0,'方伟业',1,1662198810,1662198810,''),(79,'告警中心','warning-outline','/notice','Notice','M','','','Layout',1,0,0,0,'方伟业',1,1663483539,1663483539,'/notice/notice'),(80,'通道配置','s-operation','/notice/channel','NoticeChannel','M','','','notice/channel/index',79,0,0,0,'方伟业',1,1663483671,1663483697,''),(81,'通知配置','bell','/notice/notice','NoticeConfig','M','','','notice/notice/index',79,0,0,0,'方伟业',1,1663483787,1663483787,''),(82,'通知日志','tickets','/notice/log','NoticeLog','M','','','notice/log/index',79,0,0,0,'方伟业',1,1663483837,1663496113,''),(83,'查询通道','','/notice/channels','','A','notice:channel:query','GET','',80,0,0,0,'方伟业',1,1663485499,1663485499,''),(84,'新增通道','','/notice/channel','','A','notice:channel:add','POST','',80,0,0,0,'方伟业',1,1663485522,1663485522,''),(85,'修改通道','','/notice/channel','','A','notice:channel:update','PUT','',80,0,0,0,'方伟业',1,1663485548,1663485548,''),(86,'删除通道','','/notice/channel','','A','notice:channel:delete','DELETE','',80,0,0,0,'方伟业',1,1663485569,1663485577,''),(87,'查新通知','','/notice/notice/page','','A','notice:notice:query','GET','',81,0,0,0,'方伟业',1,1663489033,1663489046,''),(88,'新增通知','','/notice/notice','','A','notice:notice:add','POST','',81,0,0,0,'方伟业',1,1663489070,1663489077,''),(89,'修改通知','','/notice/notice','','A','notice:notice:update','PUT','',81,0,0,0,'方伟业',1,1663489099,1663489099,''),(90,'删除通知','','/notice/notice','','A','notice:notice:delete','DELETE','',81,0,0,0,'方伟业',1,1663489131,1663489143,''),(91,'查询日志','','/notice/log/page','','A','notice:log:query','GET','',82,0,0,0,'方伟业',1,1663489172,1663489172,''),(92,'用户中心基本接口','','','','M','','','',19,0,1,0,'方伟业',1,1669029774,1669029774,''),(93,'配置中心基本接口','','','','M','','','',19,0,1,0,'方伟业',1,1669395165,1669395165,''),(94,'获取环境列表','','/configure/environment/filter','','A','baseApi','GET','',93,0,0,0,'方伟业',1,1669395225,1669395225,''),(95,'服务中心','s-grid','/service','Service','M','','','Layout',1,97,0,0,'方伟业',1,1669451234,1670034341,'/service/environment'),(96,'环境配置','s-tools','/configure/environment','ConfigureEnv','M','','','configure/environment/index',33,1,0,0,'方伟业',1,1669469991,1669470299,''),(97,'查询环境配置','','/configure/environment','','A','configure:environment:all','GET','',96,0,0,0,'方伟业',1,1669470078,1669470078,''),(98,'新增环境配置','','/configure/environment','','A','configure:environment:add','POST','',96,0,0,0,'方伟业',1,1669470160,1669470160,''),(99,'更新环境配置','','/configure/environment','','A','configure:environment:update','PUT','',96,0,0,0,'方伟业',1,1669470209,1669470209,''),(100,'删除环境配置','','/configure/environment','','A','configure:environment:delete','DELETE','',96,0,0,0,'方伟业',1,1669470240,1669470240,''),(101,'连接测试','','/configure/environment/connect','','A','configure:environment:connect','POST','',96,0,0,0,'方伟业',1,1669471623,1669471623,''),(102,'查询资源挂载的服务列表','','/configure/resource/service','','A','configure:resource:service:query','GET','',44,0,0,0,'方伟业',1,1669479262,1669479262,''),(103,'修改资源挂载的服务列表','','/configure/resource/service','','A','configure:resource:service:update','POST','',44,0,0,0,'方伟业',1,1669479305,1669479305,''),(104,'服务中心基本接口','','','','M','','','',19,0,0,0,'方伟业',1,1670034356,1670034356,''),(105,'获取服务环境列表','','/service/environment/filter','','A','baseApi','GET','',104,0,0,0,'方伟业',1,1670034394,1670034394,''),(106,'代码仓库','more','/service/code_registry','CodeRegistry','M','','','service/code_registry/index',95,1,0,0,'方伟业',1,1670070229,1670133364,''),(107,'获取代码仓库列表','','/service/code_registry/filter','','A','baseApi','GET','',104,0,0,0,'方伟业',1,1670070512,1670070512,''),(108,'获取代码仓库配置列表','','/service/code_registries','','A','service:code_registry:all','GET','',106,0,0,0,'方伟业',1,1670070713,1670075312,''),(109,'新增代码仓库配置','','/service/code_registry','','A','service:code_registry:add','POST','',106,0,0,0,'方伟业',1,1670070741,1670070839,''),(110,'修改代码仓库配置','','/service/code_registry','','A','service:code_registry:update','PUT','',106,0,0,0,'方伟业',1,1670070772,1670070819,''),(111,'删除代码仓库配置','','/service/code_registry','','A','service:code_registry:delete','DELETE','',106,0,0,0,'方伟业',1,1670070810,1670070810,''),(112,'连接代码仓库','','/service/code_registry/connect','','A','service:code_registry:connect','POST','',106,0,0,0,'方伟业',1,1670070905,1670070905,''),(113,'获取代码仓库类型','','/service/code_registry/types','','A','baseApi','GET','',104,0,0,0,'方伟业',1,1670070953,1670070953,''),(114,'镜像仓库','info','/service/image_registry','ImageRegistry','M','','','service/image_registry/index',95,1,0,0,'方伟业',1,1670075221,1670133373,''),(115,'获取镜像仓库配置列表','','/service/image_registries','','A','service:image_registry:all','GET','',114,0,0,0,'方伟业',1,1670075272,1670075319,''),(116,'新增镜像仓库配置','','/service/image_registry','','A','service:image_registry:add','POST','',114,0,0,0,'方伟业',1,1670075301,1670075301,''),(117,'修改镜像仓库配置','','/service/image_registry','','A','service:image_registry:update','PUT','',114,0,0,0,'方伟业',1,1670075344,1670075344,''),(118,'删除镜像仓库配置','','/service/image_registry','','A','service:image_registry:delete','DELETE','',114,0,0,0,'方伟业',1,1670075376,1670075423,''),(119,'镜像仓库连接配置','','/service/image_registry/connect','','A','service:image_registry:connect','POST','',114,0,0,0,'方伟业',1,1670075416,1670075416,''),(120,'获取镜像仓库列表','','/service/image_registry/filter','','A','baseApi','GET','',104,0,0,0,'方伟业',1,1670075454,1670075454,''),(121,'获取dockerfile模板列表','','/service/dockerfile/filter','','A','baseApi','GET','',104,0,0,0,'方伟业',1,1670078300,1670078300,''),(122,'dockerfile管理','circle-check','/service/dockerfile','Dockerfile','M','','','service/dockerfile/index',95,1,0,0,'方伟业',1,1670078363,1670133382,''),(123,'获取dockerfile模板列表','','/service/dockerfile/page','','A','service:dockerfile:all','GET','',122,0,0,0,'方伟业',1,1670078427,1670682291,''),(124,'新增dockerfile模板','','/service/dockerfile','','A','service:dockerfile:add','POST','',122,0,0,0,'方伟业',1,1670078478,1670078478,''),(125,'修改dockerfile模板','','/service/dockerfile','','A','service:dockerfile:update','PUT','',122,0,0,0,'方伟业',1,1670078526,1670078526,''),(126,'删除dockerfile模板','','/service/dockerfile','','A','service:dockerfile:delete','DELETE','',122,0,0,0,'方伟业',1,1670078558,1670078558,''),(127,'获取服务列表','','/service/service/filter','','A','baseApi','GET','',104,0,0,0,'方伟业',1,1670132397,1670132397,''),(128,'获取代码仓库下载类型','','/service/code_registry/clone_types','','A','baseApi','GET','',104,0,0,0,'方伟业',1,1670132953,1670132953,''),(129,'获取代码仓库指定项目','','/service/code_registry/project','','A','service:code_registry:project','GET','',106,0,0,0,'方伟业',1,1670170156,1670170221,''),(130,'服务构建','upload2','/service/pack_log','ServicePack','M','','','service/pack_log/index',95,0,0,0,'方伟业',1,1670170265,1670170471,''),(131,'获取构建日志','','/service/pack_log/page','','A','service:pack_log:page','GET','',130,0,0,0,'方伟业',1,1670170302,1670170302,''),(132,'进行服务构建','','/service/pack','','A','service:pack:add','POST','',130,0,0,0,'方伟业',1,1670170340,1670170340,''),(133,'获取指定服务全部分支','','/service/code_registry/branches','','A','service:code_registry:branch','GET','',130,0,0,0,'方伟业',1,1670170397,1670170397,''),(134,'获取指定服务全部标签','','/service/code_registry/tags','','A','service:code_registry:tag','GET','',130,0,0,0,'方伟业',1,1670170434,1670170434,''),(135,'发布管理','warning-outline','service/release','Release','M','','','service/release/index',95,1,0,0,'方伟业',1,1670682215,1670682483,''),(136,'获取发布模板列表','','/service/release/page','','A','service:release:page','GET','',135,0,0,0,'方伟业',1,1670682278,1670682278,''),(137,'新增发布模板','','/service/release','','A','service:release:add','POST','',135,0,0,0,'方伟业',1,1670682318,1670682349,''),(138,'修改发布模板','','/service/release','','A','service:release:update','PUT','',135,0,0,0,'方伟业',1,1670682345,1670682345,''),(139,'获取发布模板类型','','/service/release/types','','A','baseApi','GET','',104,0,0,0,'方伟业',1,1670682395,1670682417,''),(140,'删除发布模板','','/service/release','','A','service:release:delete','DELETE','',135,0,0,0,'方伟业',1,1670682447,1670682447,''),(141,'服务发布','s-promotion','/service/release_log','ServiceRelease','M','','','service/release_log/index',95,0,0,0,'方伟业',1,1670760156,1670760691,''),(142,'获取服务发布日志','','/service/release_log/page','','A','service:release_log:page','GET','',141,0,0,0,'方伟业',1,1670760214,1670760214,''),(143,'发布服务','','/service/release_log','','A','service:release_log:add','POST','',141,0,0,0,'方伟业',1,1670760598,1670760598,''),(144,'获取发布状态','','/service/release/status','','A','baseApi','GET','',104,0,0,0,'方伟业',1,1670761055,1670761055,''),(145,'获取指定服务的可发布镜像','','/service/release/images','','A','service:release:images','GET','',141,0,0,0,'方伟业',1,1670763372,1670763372,''),(146,'网络管理','set-up','/service/network','ServiceNetwork','M','','','service/network/index',95,0,0,0,'方伟业',1,1671636187,1671636187,''),(147,'获取网络管理列表','','/service/network/page','','A','service:network:page','GET','',146,0,0,0,'方伟业',1,1671636233,1671636233,''),(148,'新增网络管理','','/service/network','','A','service:network:add','POST','',146,0,0,0,'方伟业',1,1671636263,1671636263,''),(149,'修改网络','','/service/network','','A','service:network:update','PUT','',146,0,0,0,'方伟业',1,1671636285,1671636285,''),(150,'删除网络','','/service/network','','A','service:network:delete','DELETE','',146,0,0,0,'方伟业',1,1671636311,1671636311,''),(151,'获取服务运行类型','','/service/run_types','','A','baseApi','GET','',105,0,0,0,'方伟业',1,1672069723,1672069723,'');
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '角色名称',
  `keyword` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '角色关键字',
  `status` tinyint(1) NOT NULL COMMENT '角色状态',
  `weight` int(11) DEFAULT '0' COMMENT '角色权重',
  `description` varchar(300) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '角色备注',
  `data_scope` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '数据权限',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `team_ids` text COLLATE utf8_unicode_ci COMMENT '自定义权限部门id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'超级管理','super_admin',1,0,'超级管理员','ALLTEAM','系统创建',0,1659631587,1659631587,NULL),(2,'测试','test',1,0,'ss','ALLTEAM','方伟业',1,1669027963,1669221467,'[5]');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_menu`
--

DROP TABLE IF EXISTS `role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `menu_id` int(11) NOT NULL COMMENT '菜单ID',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`),
  KEY `menu_id` (`menu_id`),
  CONSTRAINT `role_menu_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `role_menu_ibfk_2` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=254 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_menu`
--

LOCK TABLES `role_menu` WRITE;
/*!40000 ALTER TABLE `role_menu` DISABLE KEYS */;
INSERT INTO `role_menu` VALUES (247,2,4,'方伟业',1,1669222734,1669222734),(248,2,15,'方伟业',1,1669222734,1669222734),(249,2,16,'方伟业',1,1669222734,1669222734),(250,2,17,'方伟业',1,1669222734,1669222734),(251,2,18,'方伟业',1,1669222734,1669222734),(252,2,1,'方伟业',1,1669222734,1669222734),(253,2,3,'方伟业',1,1669222734,1669222734);
/*!40000 ALTER TABLE `role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `team`
--

DROP TABLE IF EXISTS `team`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `team` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '部门名称',
  `description` varchar(300) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '部门备注',
  `avatar` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '部门头像',
  `parent_id` int(11) NOT NULL COMMENT '上级ID',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `team`
--

LOCK TABLES `team` WRITE;
/*!40000 ALTER TABLE `team` DISABLE KEYS */;
INSERT INTO `team` VALUES (1,'贵州青橙科技','科技、创造','/static/logo.png',0,'system',0,NULL,NULL),(3,'商务部','科技、创造','/static/logo.png',1,'',0,1659873887,1659873887),(4,'人事部','科技、创造','/static/logo.png',1,'方伟业',1,1659873897,1669302753),(5,'财务部','科技、创造','/static/logo.png',1,'方伟业',1,1659873902,1669302679),(20,'1','1','/static/logo.png',5,'方伟业',1,1669305302,1669305302),(24,'112','1','/static/logo.png',20,'方伟业',1,1669305486,1669305486),(26,'1111','1','/static/logo.png',20,'方伟业',1,1669306509,1669306509),(28,'11111','1111','/static/logo.png',26,'方伟业',1,1669306529,1669306529),(30,'1223','','/static/logo.png',28,'方伟业',1,1669306641,1669306641);
/*!40000 ALTER TABLE `team` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `team_id` int(11) NOT NULL COMMENT '部门ID',
  `nickname` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户昵称',
  `name` varchar(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户姓名',
  `phone` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户电话',
  `avatar` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户头像',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户邮箱',
  `sex` tinyint(1) NOT NULL COMMENT '用户性别',
  `password` varchar(300) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户密码',
  `last_login` int(11) DEFAULT NULL COMMENT '最后登陆时间',
  `status` tinyint(1) NOT NULL COMMENT '用户状态',
  `operator` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人员',
  `operator_id` int(11) NOT NULL COMMENT '操作人员ID',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  UNIQUE KEY `email` (`email`),
  KEY `role_id` (`role_id`),
  KEY `team_id` (`team_id`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_ibfk_2` FOREIGN KEY (`team_id`) REFERENCES `team` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,1,1,'测试','方伟业','18200000000','/user/login.png','1280291001@qq.com',1,'$2a$10$2bVK6bmc/BUkYOmLPYaytucEw0Tf9/l3H8lVHY2gg97WCXsYFrZYO',NULL,1,'方伟业',1,1659631587,1669213587);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-28 16:34:00
