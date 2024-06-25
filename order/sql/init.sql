CREATE DATABASE `exchange` COLLATE 'utf8mb4_general_ci';
CREATE TABLE exchange.`ex_order` (
                                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                           `uid` bigint(20) unsigned NOT NULL COMMENT '用户ID',
                                           `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '入库时间',
                                           `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                           `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0初始化 1成功 2失败 3部分成交 4全部成交 5已撤单 ',
                                           `error_msg` varchar(255) NOT NULL DEFAULT '' COMMENT '失败原因',
                                           `orderType` varchar(255) NOT NULL DEFAULT '' COMMENT '订单类型',
                                           `symbol` varchar(255) NOT NULL DEFAULT '' COMMENT '币对',
                                           `amount` decimal(20,6) NOT NULL DEFAULT '0' COMMENT '成交额',
                                           `quantity` decimal(20,6) NOT NULL DEFAULT '0' COMMENT '数量',
                                           `bid` decimal(20,6) NOT NULL DEFAULT '0' COMMENT '出价',
                                           `buyOrSell` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0买 1卖',

                                           PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB ROW_FORMAT=DYNAMIC COMMENT='order';