/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : could_customer_server

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 10/04/2024 20:15:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密后的密码',
  `nickname` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '头像地址',
  `last_login_at` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `email`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for chat_supports
-- ----------------------------
DROP TABLE IF EXISTS `chat_supports`;
CREATE TABLE `chat_supports`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `email` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密后的密码',
  `nickname` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '头像地址',
  `last_login_at` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `is_online` int(11) NOT NULL DEFAULT 0 COMMENT '是否在线 1-在线 0-不在线',
  `status` int(255) NOT NULL DEFAULT 1 COMMENT '账户状态 0-冻结 1-正常',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `email`(`email`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for clients
-- ----------------------------
DROP TABLE IF EXISTS `clients`;
CREATE TABLE `clients`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `email` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '客人邮箱',
  `dial_code` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '电话区号',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '电话号码',
  `merchant_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商户id',
  `domain` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '域名',
  `last_login_time` timestamp(6) NULL DEFAULT NULL COMMENT '最后登录时间',
  `last_send_message_time` int(11) NULL DEFAULT NULL COMMENT '最后发送消息时间',
  `ip` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录ip',
  `browser_fingerprint` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器指纹',
  `area` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '国家',
  `user_agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '浏览器UA',
  `lang` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '浏览器语言',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '账号状态 0-正常 1-被冻结',
  `is_online` tinyint(4) NOT NULL DEFAULT 0 COMMENT '在线状态 0-不在线 1-在线',
  `iso2` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '国家区号',
  `status_change_time` timestamp NULL DEFAULT NULL COMMENT '账号状态修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for merchant_group
-- ----------------------------
DROP TABLE IF EXISTS `merchant_group`;
CREATE TABLE `merchant_group`  (
  `chat_support_id` int(11) NOT NULL,
  `merchant_id` int(11) NOT NULL,
  PRIMARY KEY (`chat_support_id`, `merchant_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for merchants
-- ----------------------------
DROP TABLE IF EXISTS `merchants`;
CREATE TABLE `merchants`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `merchant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `merchant_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '消息时间',
  `ticket_id` int(11) NOT NULL COMMENT '工单id',
  `sender_type` smallint(11) NOT NULL COMMENT '发送类型:1-客人发送给客服，2-客服发送给客人',
  `client_id` int(11) NOT NULL COMMENT '顾客的id',
  `chat_support_id` int(11) NOT NULL COMMENT '客服的id',
  `msg_type` smallint(6) NOT NULL COMMENT '消息类型 1-文本消息',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息内容',
  `is_read` smallint(6) NULL DEFAULT 0 COMMENT '消息是否被对方读取:0-未读，1-已读',
  `read_time` timestamp NULL DEFAULT NULL COMMENT '读取消息的时间',
  `referer_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '客户发送消息时候所在的页面',
  PRIMARY KEY (`id`, `ticket_id`) USING BTREE,
  INDEX `ticket_id`(`ticket_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 362 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tickets
-- ----------------------------
DROP TABLE IF EXISTS `tickets`;
CREATE TABLE `tickets`  (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `client_id` int(11) NOT NULL COMMENT '客人id',
  `account` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户的账号名',
  `chat_support_id` int(11) NOT NULL COMMENT '客服id',
  `last_message_time` timestamp NULL DEFAULT NULL COMMENT '最后发送消息时间',
  `cs_unread_msg_count` int(11) NOT NULL DEFAULT 0 COMMENT '客服未读消息数',
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`, `cs_unread_msg_count`) USING BTREE,
  UNIQUE INDEX `client_support`(`client_id`, `chat_support_id`) USING BTREE,
  INDEX `account`(`account`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
