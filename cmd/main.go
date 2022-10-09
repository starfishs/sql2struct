package main

import (
	"github.com/gangming/sql2struct/internal/infra"
	mysqlparser "github.com/gangming/sql2struct/internal/mysql"
)

func main() {

	infra.InitDBMysql()
	//ddl := "CREATE TABLE `project` (\n  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,\n  `name` varchar(255) NOT NULL COMMENT '项目名称',\n  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1:未上线, 2:已上线,3:上线失败',\n  `valid_from` datetime DEFAULT NULL COMMENT '项目上线时间',\n  `failed_msg` varchar(255) NOT NULL DEFAULT '' COMMENT '上线失败原因',\n  `top_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '置顶标记 :置顶',\n  `allow_empty_payment_setting` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否允许无付款设置 1: 不允许 2:允许',\n  `allow_empty_collection_setting` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否允许无收款设置 1: 不允许 2:允许',\n  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',\n  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',\n  `on_top_time` datetime(6) DEFAULT NULL COMMENT '项目置顶时间',\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `name` (`name`)\n) ENGINE=InnoDB AUTO_INCREMENT=155 DEFAULT CHARSET=utf8mb4 COMMENT='项目表'"

	ddl := "CREATE TABLE `payment_setting` (\n  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,\n  `project_id` bigint(20) unsigned NOT NULL COMMENT '项目 ID',\n  `type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '付款类型:  1:按用餐计划交易付款  2: 按散客交易付款 3:按充值付款',\n  `settle_option` tinyint(4) DEFAULT '0',\n  `bill_date_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '付款方式  1: 每月付款  2: 每周付款 3:每半月付款 4:指定具体天 ',\n  `bill_date_value` json NOT NULL COMMENT '付款日',\n  `payment_offset_day` tinyint(4) DEFAULT NULL,\n  `billoff_offset_day` tinyint(4) DEFAULT NULL,\n  `reconciliation_offset_day` tinyint(4) DEFAULT NULL,\n  `receiver` varchar(255) NOT NULL COMMENT '对方称呼',\n  `biz_type` tinyint(2) unsigned NOT NULL COMMENT '付款名目: 1: 餐费',\n  `biz_alias` varchar(255) NOT NULL COMMENT '付款名目别名',\n  `invoicing_info` tinyint(2) unsigned NOT NULL COMMENT '开票说明',\n  `split_bill` tinyint(2) unsigned NOT NULL COMMENT '拆分账单 1:不拆分 2:按商户拆分',\n  `receiver_mode` tinyint(2) NOT NULL COMMENT '付款方式 1: 对公转账(人工) 2: 收款中抵扣',\n  `receiver_account_config` tinyint(2) unsigned NOT NULL COMMENT '对方账户设置 1:跟随商户设置  2:设置账户',\n  `receiver_account` json NOT NULL COMMENT '对方账户',\n  `channel_rate` json NOT NULL COMMENT '支付渠道费率 千分比',\n  `commission_config` tinyint(1) unsigned NOT NULL COMMENT '佣金设置 1:跟随商户 2:千分比',\n  `commission_rate` int(10) unsigned NOT NULL COMMENT '佣金率 千分比',\n  `is_deduction` tinyint(2) unsigned NOT NULL COMMENT '是否用于抵扣收款 1:否   2:是',\n  `deductioned_collection_ids` text COMMENT '用于抵扣收款设置 ids',\n  `setting_details` json NOT NULL COMMENT '设置详情',\n  `valid_from` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '开始生效时间',\n  `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',\n  `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',\n  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '状态 1:有效  2:失效',\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB AUTO_INCREMENT=310 DEFAULT CHARSET=utf8mb4 COMMENT='收款设置表'"

	mysqlparser.GenerateFile(ddl)
}

// 快排
