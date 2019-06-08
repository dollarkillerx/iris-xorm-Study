CREATE TABLE `star_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name_zh` varchar(50) NOT NULL DEFAULT '' COMMENT '中文名',
  `name_en` varchar(50) NOT NULL DEFAULT '' COMMENT '英文名',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `birthday` varchar(50) NOT NULL DEFAULT '' COMMENT '出生日期',
  `height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '身高，单位cm',
  `weight` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '体重，单位g',
  `club` varchar(50) NOT NULL DEFAULT '' COMMENT '俱乐部',
  `jersy` varchar(50) NOT NULL DEFAULT '' COMMENT '球衣号码以及主打位置',
  `country` varchar(50) NOT NULL DEFAULT '' COMMENT '国籍',
  `birthaddress` varchar(255) NOT NULL DEFAULT '' COMMENT '出生地',
  `feature` varchar(255) NOT NULL DEFAULT '' COMMENT '个人特点',
  `moreinfo` text COMMENT '更多介绍',
  `sys_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，默认值 0 正常，1 删除',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;