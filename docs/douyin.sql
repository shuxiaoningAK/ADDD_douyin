-- create database douyin;

-- ----------------------------
-- Table structure for
-- ----------------------------
DROP TABLE IF EXISTS ``;
CREATE TABLE ``
(
    `id`         int      NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted`    timestamp         DEFAULT NULL COMMENT '软删除标记',


    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 comment '';