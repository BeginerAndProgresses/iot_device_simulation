Create DATABASE IF NOT EXISTS `iot_device_simulate`;

use iot_device_simulate;

CREATE TABLE IF NOT EXISTS `device` (
                                        `id` int NOT NULL AUTO_INCREMENT,
                                        `plat_form` varchar(255) NOT NULL COMMENT '平台名称',
    `device_name` varchar(255) NOT NULL COMMENT '设备名称',
    `mqtt_parameter_id` int(10) unsigned zerofill DEFAULT NULL,
    `state` int(10) unsigned zerofill DEFAULT NULL COMMENT '状态，如果未启动为0，如果启动为1',
    `product_id` varchar(255) NOT NULL COMMENT '产品id',
    `user_id` int NOT NULL COMMENT '用户id',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `mqtt_parameter` (
                                                `id` int NOT NULL AUTO_INCREMENT,
                                                `client_id` varchar(255) NOT NULL COMMENT 'Client ID',
    `port` int NOT NULL COMMENT '端口号',
    `server_address` varchar(255) NOT NULL COMMENT '服务器地址',
    `username` varchar(255) NOT NULL COMMENT 'username',
    `password` varchar(255) NOT NULL COMMENT 'password',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `topic` (
                                       `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
                                       `plat_form` varchar(255) NOT NULL COMMENT '平台名',
    `topic` varchar(255) NOT NULL COMMENT '通信topic',
    `function_describe` varchar(255) NOT NULL COMMENT '功能描述',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `publish_info` (
                                              `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
                                              `json` varchar(255) NOT NULL COMMENT '信息json',
    `topic` varchar(255) NOT NULL COMMENT '通信topic',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `subscribe_info` (
                                                `id` int NOT NULL AUTO_INCREMENT,
                                                `sub_name` varchar(255) NOT NULL,
    `topic` varchar(255) NOT NULL,
    `info` varchar(255) NOT NULL COMMENT '返回的信息',
    `device_id` int NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE IF NOT EXISTS `user` (
                                      `id` int NOT NULL AUTO_INCREMENT,
                                      `username` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `nikename` varchar(255) DEFAULT NULL COMMENT '昵称',
    `avatar` varchar(255) DEFAULT NULL COMMENT '头像链接',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
