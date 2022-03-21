/*
 Navicat Premium Data Transfer

 Source Server         : gow-support
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 172.21.2.59:3306
 Source Schema         : support

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 01/12/2021 18:01:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
                            `id` bigint NOT NULL COMMENT 'id',
                            `lang` enum('cn','en') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '语言',
                            `status` int NOT NULL DEFAULT '0' COMMENT '状态,0未公开,1公开',
                            `cate_id` bigint NOT NULL COMMENT '分类id',
                            `summary` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '摘要',
                            `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '封面',
                            `markdown` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT 'markdown内容',
                            `rich_text` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '富文本内容',
                            `hot` int DEFAULT NULL COMMENT '热门,0否,1是',
                            `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `game_id` bigint DEFAULT NULL,
                            PRIMARY KEY (`id`,`lang`) USING BTREE,
                            KEY `idx_created_at` (`created`),
                            KEY `idx_updated_at` (`updated`),
                            KEY `index` (`id`,`status`,`cate_id`,`lang`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for games
-- ----------------------------
DROP TABLE IF EXISTS `games`;
CREATE TABLE `games` (
                         `id` bigint NOT NULL COMMENT 'id',
                         `game_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏名称',
                         `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '封面',
                         `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
                         `summary` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '摘要',
                         `status` int NOT NULL DEFAULT '0' COMMENT '状态，1测试中，2正式运营，3开发中，4预售中',
                         `lang` enum('cn','en') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '语言',
                         `telegram` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT 'Telegram',
                         `facebook` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT 'Facebook',
                         `youtube` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT 'Youtube',
                         `twitter` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT 'Twitter',
                         `game_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '开始游戏',
                         `guide` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  COMMENT '新手引导',
                         `created` timestamp default CURRENT_TIMESTAMP,
                         `updated` timestamp not null on update CURRENT_TIMESTAMP default CURRENT_TIMESTAMP,
                         `release` int NOT NULL DEFAULT '0' COMMENT '状态,0未公开,1公开',
                         `about_games` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT ' AboutGames',
                         `stragegy` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT 'Stragey',
                         `revenue_analysis` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT 'RevenueAnalysis',
                         index idx_updated_at(updated),
                         PRIMARY KEY (`id`) USING BTREE,
                         KEY `index` (`id`,`status`,`game_name`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------------
-- Table structure for game_parameters
-- ----------------------------------
DROP TABLE IF EXISTS `game_parameters`;
CREATE TABLE `game_parameters` (
                                   `id` bigint NOT NULL COMMENT 'id',
                                   `coin` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Token',
                                   `game_fi` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  COMMENT '游戏名称',
                                   `price` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  COMMENT 'Price',
                                   `one_day` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  COMMENT 'OneDay',
                                   `one_week` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  COMMENT 'OneWeek',
                                   `day_volume` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  COMMENT 'DayVolume',
                                   `mkt_cap` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  COMMENT 'MktCap',
                                   `status` int NOT NULL DEFAULT '0' COMMENT '状态,0未发布,1发布',
                                   `created` timestamp not null default CURRENT_TIMESTAMP,
                                   `updated` timestamp not null on update CURRENT_TIMESTAMP default CURRENT_TIMESTAMP,
                                   index idx_created_at(created),
                                   index idx_updated_at(updated),
                                   PRIMARY KEY (`id`) USING BTREE,
                                   KEY `index` (`id`,`coin`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for currency_data
-- ----------------------------
DROP TABLE IF EXISTS `currencies`;
CREATE TABLE `currencies` (
                              `id` bigint NOT NULL COMMENT 'id',
                              `logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'logo',
                              `currency_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
                              `max_amount` bigint(20) unsigned not null default 0 comment '最大供应量',
                              `value` bigint(20) unsigned not null default 0 comment '市值',
                              `flow_amount` bigint(20) unsigned not null default 0 comment '流通量',
                              `issue_at` datetime comment '发行时间',
                              `address` varchar(100) not null default '' NOT NULL comment '合约地址',
                              `created` timestamp not null default CURRENT_TIMESTAMP,
                              `updated` timestamp not null on update CURRENT_TIMESTAMP default CURRENT_TIMESTAMP,
                              index idx_created_at(created),
                              index idx_updated_at(updated),
                              PRIMARY KEY (`id`) USING BTREE,
                              KEY `index` (`id`,`logo`,`currency_name`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
                              `id` bigint NOT NULL COMMENT 'unix时间戳ID',
                              `lang` enum('cn','en') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '语言',
                              `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类名',
                              `intro` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
                              `parent_id` bigint NOT NULL DEFAULT '1',
                              PRIMARY KEY (`id`) USING BTREE,
                              KEY `index` (`id`,`lang`,`parent_id`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------
-- Table structure for label
-- --------------------------
DROP TABLE IF EXISTS `labels`;
CREATE TABLE `labels` (
                          `id` bigint NOT NULL COMMENT 'id',
                          `word` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '词条',
                          `lang` enum('cn','en') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '语言',
                          PRIMARY KEY (`id`) USING BTREE,
                          KEY `index` (`id`,`word`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------
-- Table structure for class
-- --------------------------
DROP TABLE IF EXISTS `classes`;
CREATE TABLE `classes` (
                         `id` bigint NOT NULL COMMENT 'id',
                         `class` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型',
                         PRIMARY KEY (`id`) USING BTREE,
                         KEY `index` (`id`,`class`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------
-- Table structure for banner
-- --------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
                          `id` bigint NOT NULL COMMENT 'id',
                          `chain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '链接',
                          `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '封面',
                          PRIMARY KEY (`id`) USING BTREE,
                          KEY `index` (`id`,`chain`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- --------------------------
-- Table structure for chain
-- --------------------------
DROP TABLE IF EXISTS `chains`;
CREATE TABLE `chains` (
                          `id` bigint NOT NULL COMMENT 'id',
                          `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '链名',
                          `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'ICON',
                          PRIMARY KEY (`id`) USING BTREE,
                          KEY `index` (`id`,`name`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- -------------------------------------------------
-- Table of relationship between game and label
-- -------------------------------------------------
DROP TABLE IF EXISTS `game_label`;
CREATE TABLE `game_label` (
                              `game_id` bigint(20) unsigned not null,
                              `label_id` bigint(20) unsigned not null,
                              primary key(`game_id`, `label_id`)
)engine=InnoDB default charset=utf8mb4 comment='游戏标签关联表';

-- -------------------------------------------------
-- Table of relationship between game and class
-- -------------------------------------------------
DROP TABLE IF EXISTS `game_class`;
CREATE TABLE `game_class` (
                              `game_id` bigint(20) unsigned not null,
                              `class_id` bigint(20) unsigned not null,
                              primary key(`game_id`, `class_id`)
)engine=InnoDB default charset=utf8mb4 comment='游戏类型关联表';

-- -------------------------------------------------
-- Table of relationship between games and chain
-- -------------------------------------------------
DROP TABLE IF EXISTS `game_chain`;
CREATE TABLE `game_chain` (
                              `game_id` bigint(20) unsigned not null,
                              `chain_id` bigint(20) unsigned not null,
                              primary key(`game_id`, `chain_id`)
)engine=InnoDB default charset=utf8mb4 comment='游戏链关联表';

-- -------------------------------------------------
-- Table of relationship between games and currency
-- -------------------------------------------------
DROP TABLE IF EXISTS `game_currency`;
CREATE TABLE `game_currency` (
                                 `game_id` bigint(20) unsigned not null,
                                 `currency_id` bigint(20) unsigned not null,
                                 primary key(`game_id`, `currency_id`)
)engine=InnoDB default charset=utf8mb4 comment='游戏代币关联表';

-- -------------------------------------------------
-- Table of relationship between article and label
-- -------------------------------------------------
DROP TABLE IF EXISTS `article_label`;
CREATE TABLE `article_label` (
                                 `article_id` bigint(20) unsigned not null,
                                 `label_id` bigint(20) unsigned not null,
                                 primary key(`article_id`, `label_id`)
)engine=InnoDB default charset=utf8mb4 comment='文章标签关联表';

-- -------------------------------------------------
-- Table of relationship between games and article
-- -------------------------------------------------
DROP TABLE IF EXISTS `game_article`;
CREATE TABLE `game_article` (
                                `game_id` bigint(20) unsigned not null,
                                `article_id` bigint(20) unsigned not null,
                                primary key(`game_id`, `article_id`)
)engine=InnoDB default charset=utf8mb4 comment='游戏文章关联表';

-- ----------------------------------
-- Table structure for game_parameters
-- ----------------------------------
DROP TABLE IF EXISTS `ip_records`;
CREATE TABLE `ip_records` (
                                `id` bigint NOT NULL COMMENT 'id',
                                `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'ip地址',
                                `country` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '国家',
                                `created` timestamp NULL DEFAULT CURRENT_DATE,
                                PRIMARY KEY (`id`) USING BTREE,
                                KEY `index` (`id`) COMMENT '复合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;