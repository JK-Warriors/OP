/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : aiopms

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-03-28 17:23:47
*/

SET FOREIGN_KEY_CHECKS=0;




-- ----------------------------
-- Table structure for pms_roles
-- ----------------------------
DROP TABLE IF EXISTS `pms_roles`;
CREATE TABLE `pms_roles` (
  `id` bigint(20) NOT NULL,
  `name` varchar(30) DEFAULT NULL COMMENT '角色名称',
  `summary` varchar(500) DEFAULT NULL COMMENT '角色描述',
  `created` int(10) DEFAULT NULL,
  `changed` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `INDEX_NCC` (`name`,`created`,`changed`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Records of pms_roles
-- ----------------------------
INSERT INTO `pms_roles` VALUES ('1', '管理员', '系统管理员', UNIX_TIMESTAMP(now()), UNIX_TIMESTAMP(now()));
INSERT INTO `pms_roles` VALUES ('2', '业务员', '业务员', UNIX_TIMESTAMP(now()), UNIX_TIMESTAMP(now()));
INSERT INTO `pms_roles` VALUES ('3', '审计员', '审计员', UNIX_TIMESTAMP(now()), UNIX_TIMESTAMP(now()));


-- ----------------------------
-- Table structure for pms_permissions
-- ----------------------------
DROP TABLE IF EXISTS `pms_permissions`;
CREATE TABLE `pms_permissions` (
  `id` bigint(20) NOT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL COMMENT '中文名称',
  `ename` varchar(50) DEFAULT NULL COMMENT '英文名称',
  `url` varchar(255) DEFAULT '0' COMMENT 'URL地址',
  `icon` varchar(20) DEFAULT NULL,
  `is_nav` tinyint(1) DEFAULT '0' COMMENT '1是0否导航',
  `is_show` tinyint(1) DEFAULT '0' COMMENT '0不显示1显示',
  `sort` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `INDEX_PNETW` (`parent_id`,`name`,`ename`,`is_show`,`sort`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='权限表';

-- ----------------------------
-- Records of pms_permissions
-- ----------------------------
INSERT INTO `pms_permissions` VALUES ('0', '0', '根节点', 'root', '/root', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('2', '0', '配置中心', 'config', '/config', '', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('3', '0', 'Oracle', 'Oracle', '/oracle', '', '1', '1', '3');
INSERT INTO `pms_permissions` VALUES ('4', '0', 'MySQL', 'MySQL', '/mysql', '', '1', '1', '4');
INSERT INTO `pms_permissions` VALUES ('5', '0', 'SQLServer', 'SQLServer', '/mssql', '', '1', '1', '5');
INSERT INTO `pms_permissions` VALUES ('6', '0', 'OS', 'OS', '/os', '', '1', '1', '6');
INSERT INTO `pms_permissions` VALUES ('9', '0', '容灾操作', 'operation', '/operation', '', '1', '1', '9');
INSERT INTO `pms_permissions` VALUES ('98', '0', '告警管理', 'alarm', '/alarm', '', '1', '1', '98');
INSERT INTO `pms_permissions` VALUES ('99', '0', '系统管理', 'system', '/system', '', '1', '1', '99');

INSERT INTO `pms_permissions` VALUES ('2100', '2', '资产配置', 'config-db-manage', '/config/db/manage', '', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('2101', '2', '添加资产', 'config-db-add', '/config/db/add', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('2102', '2', '编辑资产', 'config-db-edit', '/config/db/edit', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('2103', '2', '删除资产', 'config-db-delete', '/config/db/delete', '', '0', '0', '0');

-- INSERT INTO `pms_permissions` VALUES ('2110', '2', '业务系统配置', 'config-business-manage', '/config/dr_business/manage', '', '1', '1', '2');
-- INSERT INTO `pms_permissions` VALUES ('2111', '2', '添加业务系统', 'config-business-add', '/config/dr_business/add', '', '0', '0', '0');
-- INSERT INTO `pms_permissions` VALUES ('2112', '2', '编辑业务系统', 'config-business-edit', '/config/dr_business/edit', '', '0', '0', '0');
-- INSERT INTO `pms_permissions` VALUES ('2113', '2', '删除业务系统', 'config-business-delete', '/config/dr_business/delete', '', '0', '0', '0');

INSERT INTO `pms_permissions` VALUES ('2120', '2', '容灾配置', 'config-dr-manage', '/config/dr_config/manage', '', '1', '1', '3');

INSERT INTO `pms_permissions` VALUES ('2130', '2', '全局配置', 'config-global-manage', '/config/config_global/manage', '', '1', '1', '3');

-- INSERT INTO `pms_permissions` VALUES ('2140', '2', '大屏配置', 'config-dr-manage', '/config/dr_config/manage', '', '1', '1', '4');

INSERT INTO `pms_permissions` VALUES ('2150', '2', '告警配置', 'config-alert-manage', '/config/config_alert/manage', '', '1', '1', '5');

INSERT INTO `pms_permissions` VALUES ('3100', '3', '实例状态', 'oracle-status-manage', '/oracle/status/manage', '', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('3110', '3', '表空间', 'oracle-tbs-manage', '/oracle/tbs/manage', '', '1', '1', '2');
INSERT INTO `pms_permissions` VALUES ('3120', '3', '磁盘组', 'oracle-asm-manage', '/oracle/asm/manage', '', '1', '1', '3');

INSERT INTO `pms_permissions` VALUES ('4100', '4', '实例状态', 'mysql-status-manage', '/mysql/status/manage', '', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('4110', '4', '资源', 'mysql-resource-manage', '/mysql/resource/manage', '', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('4120', '4', '键缓存', 'mysql-key-manage', '/mysql/key/manage', '', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('4130', '4', 'InnoDB', 'mysql-innodb-manage', '/mysql/innodb/manage', '', '1', '1', '1');
-- INSERT INTO `pms_permissions` VALUES ('4140', '4', '大表分析', 'mysql-bigtable-manage', '/mysql/bigtable/manage', '', '1', '1', '1');
-- INSERT INTO `pms_permissions` VALUES ('4150', '4', 'AWR报告', 'mysql-awr-manage', '/mysql/awr/manage', '', '1', '1', '1');

INSERT INTO `pms_permissions` VALUES ('5100', '5', '实例状态', 'mssql-status-manage', '/mssql/status/manage', '', '1', '1', '1');

INSERT INTO `pms_permissions` VALUES ('6100', '6', '健康状态', 'os-status-manage', '/os/status/manage', '', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('6110', '6', '磁盘', 'disk-status-manage', '/disk/status/manage', '', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('6120', '6', '磁盘IO', 'diskio-status-manage', '/diskio/status/manage', '', '1', '1', '1');

INSERT INTO `pms_permissions` VALUES ('9100', '9', '容灾切换', 'oper-switch-manage', '/operation/dr_switch/manage', '', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('9101', '9', '容灾切换', 'oper-switch-view', '/operation/dr_switch/view', '', '0', '0', '1');
INSERT INTO `pms_permissions` VALUES ('9102', '9', '容灾激活', 'oper-active-manage', '/operation/dr_active/manage', '', '1', '1', '2');
INSERT INTO `pms_permissions` VALUES ('9103', '9', '容灾同步', 'oper-sync-manage', '/operation/dr_sync/manage', '', '1', '1', '3');
INSERT INTO `pms_permissions` VALUES ('9104', '9', '容灾快照', 'oper-snapshot-manage', '/operation/dr_snapshot/manage', '', '1', '1', '4');
INSERT INTO `pms_permissions` VALUES ('9105', '9', '误删除恢复', 'oper-recover-manage', '/operation/dr_recover/manage', '', '1', '1', '5');


INSERT INTO `pms_permissions` VALUES ('9910', '99', '用户管理', 'user-manage', '/system/user/manage', 'fa-user', '1', '1', '1');
INSERT INTO `pms_permissions` VALUES ('9911', '99', '添加用户', 'user-add', '/system/user/add', null, '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('9912', '99', '编辑用户', 'user-edit', '/system/user/edit', null, '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('9913', '99', '删除用户', 'user-delete', '/system/user/delete', '', '0', '0', '0');

INSERT INTO `pms_permissions` VALUES ('9920', '99', '角色管理', 'role-manage', '/system/role/manage', '', '1', '1', '2');
INSERT INTO `pms_permissions` VALUES ('9921', '99', '添加角色', 'role-add', '/system/role/add', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('9922', '99', '编辑角色', 'role-edit', '/system/role/edit', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('9923', '99', '删除角色', 'role-delete', '/system/role/delete', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('9924', '99', '角色权限', 'role-permission', '/system/role/permission', '', '0', '0', '0');
-- INSERT INTO `pms_permissions` VALUES ('9925', '99', '角色成员', 'role-user', '/system/role/user', '', '0', '0', '0');
-- INSERT INTO `pms_permissions` VALUES ('9926', '99', '添加角色', 'role-user-add', '/system/role/useradd', '', '0', '0', '0');
-- INSERT INTO `pms_permissions` VALUES ('9927', '99', '删除角色', 'role-user-delete', '/system/role/userdelete', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('9930', '99', '权限管理', 'permission-manage', '/system/permission/manage', '', '0', '0', '3');
INSERT INTO `pms_permissions` VALUES ('9931', '99', '添加权限', 'permission-add', '/system/permission/add', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('9932', '99', '编辑权限', 'permission-edit', '/system/permission/edit', '', '0', '0', '0');
INSERT INTO `pms_permissions` VALUES ('9933', '99', '删除权限', 'permission-delete', '/system/permission/delete', '', '0', '0', '0');

INSERT INTO `pms_permissions` VALUES ('9940', '99', '日志管理', 'log-manage', '/system/log/manage', '', '1', '1', '4');
INSERT INTO `pms_permissions` VALUES ('9941', '99', '日志删除', 'log-delete', '/system/log/delete', '', '0', '0', '0');

INSERT INTO `pms_permissions` VALUES ('9950', '99', '消息管理', 'message-manage', '/system/message/manage', '', '0', '0', '5');
INSERT INTO `pms_permissions` VALUES ('9951', '99', '消息删除', 'message-delete', '/system/message/delete', '', '0', '0', '0');

-- ----------------------------
-- Table structure for pms_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `pms_role_permission`;
CREATE TABLE `pms_role_permission` (
  `id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  `permission_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`,`role_id`),
  KEY `INDEX_GP` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色权限表';

-- ----------------------------
-- Records of pms_groups_permission
-- ----------------------------
INSERT INTO `pms_role_permission` VALUES ('2', '1', '2');
INSERT INTO `pms_role_permission` VALUES ('3', '1', '3');
INSERT INTO `pms_role_permission` VALUES ('4', '1', '4');
INSERT INTO `pms_role_permission` VALUES ('5', '1', '5');
INSERT INTO `pms_role_permission` VALUES ('6', '1', '6');
INSERT INTO `pms_role_permission` VALUES ('9', '1', '9');
INSERT INTO `pms_role_permission` VALUES ('98', '1', '98');
INSERT INTO `pms_role_permission` VALUES ('99', '1', '99');
INSERT INTO `pms_role_permission` VALUES ('2100', '1', '2100');
INSERT INTO `pms_role_permission` VALUES ('2101', '1', '2101');
INSERT INTO `pms_role_permission` VALUES ('2102', '1', '2102');
INSERT INTO `pms_role_permission` VALUES ('2103', '1', '2103');
-- INSERT INTO `pms_role_permission` VALUES ('2110', '1', '2110');
-- INSERT INTO `pms_role_permission` VALUES ('2111', '1', '2111');
-- INSERT INTO `pms_role_permission` VALUES ('2112', '1', '2112');
-- INSERT INTO `pms_role_permission` VALUES ('2113', '1', '2113');
INSERT INTO `pms_role_permission` VALUES ('2120', '1', '2120');
INSERT INTO `pms_role_permission` VALUES ('2130', '1', '2130');
INSERT INTO `pms_role_permission` VALUES ('2140', '1', '2140');
INSERT INTO `pms_role_permission` VALUES ('2150', '1', '2150');
INSERT INTO `pms_role_permission` VALUES ('3100', '1', '3100');
INSERT INTO `pms_role_permission` VALUES ('3110', '1', '3110');
INSERT INTO `pms_role_permission` VALUES ('3120', '1', '3120');
INSERT INTO `pms_role_permission` VALUES ('4100', '1', '4100');
INSERT INTO `pms_role_permission` VALUES ('4110', '1', '4110');
INSERT INTO `pms_role_permission` VALUES ('4120', '1', '4120');
INSERT INTO `pms_role_permission` VALUES ('4130', '1', '4130');
-- INSERT INTO `pms_role_permission` VALUES ('4140', '1', '4140');
-- INSERT INTO `pms_role_permission` VALUES ('4150', '1', '4150');
INSERT INTO `pms_role_permission` VALUES ('5100', '1', '5100');
INSERT INTO `pms_role_permission` VALUES ('6100', '1', '6100');
INSERT INTO `pms_role_permission` VALUES ('6110', '1', '6110');
INSERT INTO `pms_role_permission` VALUES ('6120', '1', '6120');
INSERT INTO `pms_role_permission` VALUES ('9100', '1', '9100');
INSERT INTO `pms_role_permission` VALUES ('9101', '1', '9101');
INSERT INTO `pms_role_permission` VALUES ('9102', '1', '9102');
INSERT INTO `pms_role_permission` VALUES ('9103', '1', '9103');
INSERT INTO `pms_role_permission` VALUES ('9104', '1', '9104');
INSERT INTO `pms_role_permission` VALUES ('9105', '1', '9105');
INSERT INTO `pms_role_permission` VALUES ('9910', '1', '9910');
INSERT INTO `pms_role_permission` VALUES ('9911', '1', '9911');
INSERT INTO `pms_role_permission` VALUES ('9912', '1', '9912');
INSERT INTO `pms_role_permission` VALUES ('9913', '1', '9913');
INSERT INTO `pms_role_permission` VALUES ('9920', '1', '9920');
INSERT INTO `pms_role_permission` VALUES ('9921', '1', '9921');
INSERT INTO `pms_role_permission` VALUES ('9922', '1', '9922');
INSERT INTO `pms_role_permission` VALUES ('9923', '1', '9923');
INSERT INTO `pms_role_permission` VALUES ('9924', '1', '9924');
INSERT INTO `pms_role_permission` VALUES ('9930', '1', '9930');
INSERT INTO `pms_role_permission` VALUES ('9931', '1', '9931');
INSERT INTO `pms_role_permission` VALUES ('9932', '1', '9932');
INSERT INTO `pms_role_permission` VALUES ('9933', '1', '9933');
INSERT INTO `pms_role_permission` VALUES ('9940', '1', '9940');
INSERT INTO `pms_role_permission` VALUES ('9941', '1', '9941');
INSERT INTO `pms_role_permission` VALUES ('9950', '1', '9950');
INSERT INTO `pms_role_permission` VALUES ('9951', '1', '9951');



-- ----------------------------
-- Table structure for pms_users
-- ----------------------------
DROP TABLE IF EXISTS `pms_users`;
CREATE TABLE `pms_users` (
  `userid` bigint(20) NOT NULL,
  `profile_id` bigint(20) DEFAULT NULL,
  `username` varchar(15) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `avatar` varchar(100) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '1' COMMENT '状态1正常，2禁用',
  PRIMARY KEY (`userid`),
  KEY `INDEX_US` (`username`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户主表';

-- ----------------------------
-- Records of pms_users
-- ----------------------------
INSERT INTO `pms_users` VALUES ('1', '1', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '/static/img/avatar/1.jpg', '1');
INSERT INTO `pms_users` VALUES ('2', '2', 'audit', 'e10adc3949ba59abbe56e057f20f883e', '/static/img/avatar/1.jpg', '1');

-- ----------------------------
-- Table structure for pms_users_profile
-- ----------------------------
DROP TABLE IF EXISTS `pms_users_profile`;
CREATE TABLE `pms_users_profile` (
  `userid` bigint(20) NOT NULL auto_increment,
  `realname` varchar(15) DEFAULT NULL COMMENT '姓名',
  `sex` tinyint(1) DEFAULT '1' COMMENT '1男2女',
  `birth` varchar(15) DEFAULT NULL,
  `email` varchar(30) DEFAULT NULL COMMENT '邮箱',
  `webchat` varchar(15) DEFAULT NULL COMMENT '微信号',
  `qq` varchar(15) DEFAULT NULL COMMENT 'qq号',
  `phone` varchar(15) DEFAULT NULL COMMENT '手机',
  `tel` varchar(20) DEFAULT NULL COMMENT '电话',
  `address` varchar(100) DEFAULT NULL COMMENT '地址',
  `emercontact` varchar(15) DEFAULT NULL COMMENT '紧急联系人',
  `emerphone` varchar(15) DEFAULT NULL COMMENT '紧急电话',
  `lognum` int(10) DEFAULT '0' COMMENT '登录次数',
  `ip` varchar(15) DEFAULT NULL COMMENT '最近登录IP',
  `lasted` int(10) DEFAULT NULL COMMENT '最近登录时间',
  PRIMARY KEY (`userid`),
  KEY `INDEX_RSL` (`realname`,`sex`,`lasted`)
) ENGINE=InnoDB AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8 COMMENT='用户详情表';

-- ----------------------------
-- Records of pms_users_profile
-- ----------------------------
INSERT INTO `pms_users_profile` VALUES ('1', 'admin', '1', '1993-03-06', 'admin@tom.com', '', '', '13282176663', '', '', '', '',  '0', '', '0');
INSERT INTO `pms_users_profile` VALUES ('2', 'audit', '1', '1985-12-12', 'audit@163.com', '', '', '', '', '', '', '',  '0', '', '0');

-- ----------------------------
-- Table structure for pms_role_user
-- ----------------------------
DROP TABLE IF EXISTS `pms_role_user`;
CREATE TABLE `pms_role_user` (
  `id` bigint(20) NOT NULL auto_increment,
  `role_id` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `INDEX_GU` (`role_id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='角色成员';

-- ----------------------------
-- Records of pms_groups_user
-- ----------------------------
INSERT INTO `pms_role_user` VALUES ('1', '1', '1');
INSERT INTO `pms_role_user` VALUES ('2', '2', '2');


-- ----------------------------
-- Table structure for pms_messages
-- ----------------------------
DROP TABLE IF EXISTS `pms_messages`;
CREATE TABLE `pms_messages` (
  `msgid` bigint(20) NOT NULL,
  `userid` bigint(20) DEFAULT NULL,
  `touserid` bigint(20) DEFAULT NULL,
  `type` tinyint(2) DEFAULT NULL COMMENT '类型1评论2赞3审批',
  `subtype` tinyint(3) DEFAULT NULL COMMENT '11知识评论12相册评论21知识赞22相册赞31请假审批32加班33报销34出差35外出36物品',
  `title` varchar(200) DEFAULT NULL,
  `url` varchar(200) DEFAULT NULL,
  `view` tinyint(1) DEFAULT '1' COMMENT '1未看，2已看',
  `created` int(10) DEFAULT NULL,
  PRIMARY KEY (`msgid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT=' 消息表';

-- ----------------------------
-- Records of pms_messages
-- ----------------------------
INSERT INTO `pms_messages` VALUES ('66618325785907200', '1461312703628858832', '1469024587469707428', '4', '31', '去审批处理', '/leave/approval/66618286464307200', '1', '1490685934');
INSERT INTO `pms_messages` VALUES ('66626417378463744', '1461312703628858832', '1461312703628858832', '1', '11', 'OPMS 1.2 版本更新发布', '/knowledge/66618679508340736', '1', '1490687863');
INSERT INTO `pms_messages` VALUES ('66639445431947264', '1461312703628858832', '1461312703628858832', '1', '12', '油菜花', '/album/66621262012616704', '1', '1490690969');


-- ----------------------------
-- Table structure for pms_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `pms_admin_log`;
CREATE TABLE `pms_admin_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '用户名称',
  `url` varchar(1500) NOT NULL DEFAULT '' COMMENT '操作页面',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '日志标题',
  `content` text NOT NULL COMMENT '内容',
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'IP',
  `useragent` varchar(255) NOT NULL DEFAULT '' COMMENT 'User-Agent',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `name` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='系统日志表';


-- -----------------------------------------------------------------------------
-- Table structure for pms_asset_config
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_asset_config`;
CREATE TABLE `pms_asset_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `asset_type` tinyint(2) DEFAULT NULL,
  `host` varchar(30) NOT NULL DEFAULT '' COMMENT 'IP',
  `protocol` varchar(10) NOT NULL DEFAULT 0 COMMENT '协议',
  `port` int(10) NOT NULL DEFAULT 0 COMMENT '端口',
  `alias` varchar(255) DEFAULT '' COMMENT '别名',
  `instance_name` varchar(50) DEFAULT '' COMMENT '实例名',
  `db_name` varchar(50) DEFAULT '' COMMENT '数据库名',
  `username` varchar(30) DEFAULT '' COMMENT '用户名',
  `password` varchar(255) DEFAULT '' COMMENT '密码',
  `role` tinyint(2) DEFAULT 1 COMMENT '1：主; 2: 备',
  `os_type` tinyint(2) DEFAULT 1 COMMENT '1：Linux; 2: Windows; 3: AIX; 4: HP-Unix; 5: Solaris',
  `os_protocol` varchar(10) DEFAULT 0 COMMENT '主机协议',
  `os_port` varchar(10)  DEFAULT 0 COMMENT '主机端口',
  `os_username` varchar(30) DEFAULT '' COMMENT '主机用户名',
  `os_password` varchar(255) DEFAULT '' COMMENT '主机密码',
  `status` tinyint(2) DEFAULT 1 COMMENT '1: 激活；2：禁用',
  `is_delete` tinyint(2) DEFAULT 0 COMMENT '0：未删除; 1: 删除',
  `retention` int(10) NOT NULL DEFAULT 0 COMMENT '保留时间，默认单位为天',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  `updated` int(10) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `host` (`host`),
  KEY `alias` (`alias`)
)ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8 COMMENT='资产配置表';


-- -----------------------------------------------------------------------------
-- Table structure for pms_dr_business
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_dr_business`;
CREATE TABLE `pms_dr_business` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `bs_name` varchar(50) DEFAULT '' COMMENT '业务系统名',
  `is_delete` tinyint(2) DEFAULT 0 COMMENT '1: 删除；0：未删除',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  `updated` int(10) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `bs_name` (`bs_name`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='业务系统配置表';


-- -----------------------------------------------------------------------------
-- Table structure for pms_dr_config
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_dr_config`;
CREATE TABLE `pms_dr_config` (
  `bs_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Business Id',
  `bs_name` varchar(200) DEFAULT NULL,
  `asset_type` tinyint(2) DEFAULT NULL,
  `db_id_p` int(10) COMMENT 'primary db id',
  `db_dest_p` tinyint(2) COMMENT 'primary dest id',
  `db_id_s` int(10) COMMENT 'standby db id',
  `db_dest_s` tinyint(2) COMMENT 'standby dest id',
  `db_name` varchar(100) DEFAULT NULL,
  `fb_retention` int(10) COMMENT 'flashback retention',
  `is_shift` tinyint(1),
  `shift_vips` varchar(400),
  `network_p` varchar(100) COMMENT 'primary network card',
  `network_s` varchar(100) COMMENT 'standby network card',
  `is_switch` tinyint(1) DEFAULT 0,
  `status` tinyint(1) DEFAULT 1 COMMENT '1: 激活；2：禁用',
  `is_delete` tinyint(1) DEFAULT 0 COMMENT '0：未删除; 1: 删除',
  `on_process` tinyint(1) DEFAULT 0,
  `on_switchover` tinyint(1) DEFAULT 0,
  `on_failover` tinyint(1) DEFAULT 0,
  `on_startsync` tinyint(1) DEFAULT 0,
  `on_stopsync` tinyint(1) DEFAULT 0,
  `on_startread` tinyint(1) DEFAULT 0,
  `on_stopread` tinyint(1) DEFAULT 0,
  `on_startsnapshot` tinyint(1) DEFAULT 0,
  `on_stopsnapshot` tinyint(1) DEFAULT 0,
  `on_startflashback` tinyint(1) DEFAULT 0,
  `on_stopflashback` tinyint(1) DEFAULT 0,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  `updated` int(10) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`bs_id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='业务系统配置扩展表';

alter table pms_dr_config modify column on_process tinyint(1) DEFAULT 0 comment '值为1时，表明正在进行Switchover，或者Failover，或者开启停止MRP进程'; 
alter table pms_dr_config modify column on_switchover tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在进行Switchover切换'; 
alter table pms_dr_config modify column on_failover tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在进行Failover切换'; 
alter table pms_dr_config modify column on_startsync tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在开启同步进程'; 
alter table pms_dr_config modify column on_stopsync tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在停止同步进程'; 
alter table pms_dr_config modify column on_startread tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在开启可读'; 
alter table pms_dr_config modify column on_stopread tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在停止可读'; 
alter table pms_dr_config modify column on_startsnapshot tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在激活数据库快照'; 
alter table pms_dr_config modify column on_stopsnapshot tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在从快照恢复到物理备库'; 
alter table pms_dr_config modify column on_startflashback tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在进行数据库闪回'; 
alter table pms_dr_config modify column on_stopflashback tinyint(1) DEFAULT 0 comment '值为1时，表明当前正在从闪回恢复到同步状态'; 

-- -----------------------------------------------------------------------------
-- Table structure for pms_template
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_template`;
CREATE TABLE `pms_template` (
  `template_id` int(10) NOT NULL AUTO_INCREMENT,
  `asset_type`     varchar(50) DEFAULT NULL,
  `scraper_name` varchar(255)  DEFAULT NULL,
  `subsystem` varchar(255)  DEFAULT NULL,
  `metrix_name` varchar(255)  DEFAULT NULL,
  `label` varchar(255)  DEFAULT NULL,
  `value_type` tinyint(2) DEFAULT 0 COMMENT '1: Counter；2: Gauge；3：Histogram；4：Summary；5：Untyped',
  PRIMARY KEY (`template_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='模板表';


-- -----------------------------------------------------------------------------
-- Table structure for pms_items
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_items`;
CREATE TABLE `pms_items` (
	`item_id`                bigint unsigned                           NOT NULL         AUTO_INCREMENT,
	`type`                   integer         DEFAULT '0'               NOT NULL,
	`template_id`            bigint unsigned                           NULL,
	`obj_type`               varchar(50)                               NOT NULL,
	`obj_id`                 bigint unsigned                           NOT NULL,
	`name`                   varchar(255)    DEFAULT ''                NOT NULL,
	`key_`                   varchar(255)    DEFAULT ''                NOT NULL,
	`label`                  varchar(255)    DEFAULT ''                NOT NULL,
	`value_type`             integer         DEFAULT '1'               NOT NULL,
	`units`                  varchar(255)    DEFAULT ''                NOT NULL,
	`status`                 integer         DEFAULT '1'               NOT NULL,
  PRIMARY KEY (`item_id`),
  KEY `idx_items_1` (`obj_id`,`key_`,`label`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8 COMMENT='items表';

-- -----------------------------------------------------------------------------
-- Table structure for pms_item_data
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `pms_item_data`;
CREATE TABLE `pms_item_data` (
	`itemid`                 bigint unsigned                           NOT NULL,
	`time`                   int(10)         DEFAULT '0'               NOT NULL,
	`value`                  double(16,4)    DEFAULT '0.0000'          NOT NULL,
	`ns`                     integer         DEFAULT '0'               NOT NULL,
  PRIMARY KEY (`itemid`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='数据表';


-- ----------------------------
-- Table structure for pms_asset_status
-- ----------------------------
DROP TABLE IF EXISTS `pms_asset_status`;
CREATE TABLE `pms_asset_status` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `asset_id` int(10) NOT NULL DEFAULT '0',
  `asset_type` tinyint(2) DEFAULT NULL COMMENT '资产类型',
  `host` varchar(50) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL,
  `role`    varchar(30) DEFAULT NULL COMMENT '角色',
  `version` varchar(30) DEFAULT NULL COMMENT '版本',
  `connect` tinyint(2) DEFAULT NULL COMMENT '连接',
  `sessions` tinyint(2) NOT NULL DEFAULT '-1',
  `repl` tinyint(2) NOT NULL DEFAULT '-1',
  `repl_delay` tinyint(2) NOT NULL DEFAULT '-1',
  `tablespace` tinyint(2) NOT NULL DEFAULT '-1',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `pms_asset_status` ADD CONSTRAINT asset_id UNIQUE(`asset_id`);


-- ----------------------------
-- Table structure for pms_asset_status_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_asset_status_his`;
CREATE TABLE `pms_asset_status_his` (
  `id` int(10) unsigned NOT NULL COMMENT 'ID',
  `asset_id` int(10) NOT NULL DEFAULT '0',
  `asset_type` tinyint(2) DEFAULT NULL COMMENT '资产类型',
  `host` varchar(50) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL,
  `role`    varchar(30) DEFAULT NULL COMMENT '角色',
  `version` varchar(30) DEFAULT NULL COMMENT '版本',
  `connect` tinyint(2) DEFAULT NULL COMMENT '连接',
  `sessions` tinyint(2) NOT NULL DEFAULT '-1',
  `repl` tinyint(2) NOT NULL DEFAULT '-1',
  `repl_delay` tinyint(2) NOT NULL DEFAULT '-1',
  `tablespace` tinyint(2) NOT NULL DEFAULT '-1',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `idx_asset_id` (`asset_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_oracle_status
-- ----------------------------
DROP TABLE IF EXISTS `pms_oracle_status`;
CREATE TABLE `pms_oracle_status` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `db_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL,
  `connect` tinyint(2) DEFAULT NULL COMMENT '连接',
  `inst_num` tinyint(2) NOT NULL DEFAULT '-1',
  `inst_name` varchar(30) NOT NULL DEFAULT '-1',
  `inst_role` varchar(50) NOT NULL DEFAULT '-1',
  `inst_status` varchar(50) NOT NULL DEFAULT '-1',
  `version` varchar(50) NOT NULL DEFAULT '-1',
  `startup_time` varchar(50) NOT NULL DEFAULT '-1',
  `host_name` varchar(50) NOT NULL DEFAULT '-1',
  `archiver` varchar(50) NOT NULL DEFAULT '-1',
  `db_name` varchar(30) NOT NULL DEFAULT '-1',
  `db_role` varchar(50) NOT NULL DEFAULT '-1',
  `open_mode` varchar(30) NOT NULL DEFAULT '-1',
  `protection_mode` varchar(30) NOT NULL DEFAULT '-1',
  `session_total` int(10) NOT NULL DEFAULT '-1',
  `session_actives` int(4) NOT NULL DEFAULT '-1',
  `session_waits` int(4) NOT NULL DEFAULT '-1',
  `dg_stats` varchar(255) NOT NULL DEFAULT '-1',
  `dg_delay` int(10) NOT NULL DEFAULT '-1',
  `processes` int(10) NOT NULL DEFAULT '-1',
  `flashback_on` varchar(10) DEFAULT NULL COMMENT '闪回状态',
  `flashback_usage` varchar(10) DEFAULT NULL COMMENT '闪回空间使用率',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_oracle_status_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_oracle_status_his`;
CREATE TABLE `pms_oracle_status_his` (
  `id` int(10) unsigned NOT NULL COMMENT 'ID',
  `db_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL,
  `connect` tinyint(2) DEFAULT NULL COMMENT '连接',
  `inst_num` tinyint(2) NOT NULL DEFAULT '-1',
  `inst_name` varchar(30) NOT NULL DEFAULT '-1',
  `inst_role` varchar(50) NOT NULL DEFAULT '-1',
  `inst_status` varchar(50) NOT NULL DEFAULT '-1',
  `version` varchar(50) NOT NULL DEFAULT '-1',
  `startup_time` varchar(50) NOT NULL DEFAULT '-1',
  `host_name` varchar(50) NOT NULL DEFAULT '-1',
  `archiver` varchar(50) NOT NULL DEFAULT '-1',
  `db_name` varchar(30) NOT NULL DEFAULT '-1',
  `db_role` varchar(50) NOT NULL DEFAULT '-1',
  `open_mode` varchar(30) NOT NULL DEFAULT '-1',
  `protection_mode` varchar(30) NOT NULL DEFAULT '-1',
  `session_total` int(10) NOT NULL DEFAULT '-1',
  `session_actives` int(4) NOT NULL DEFAULT '-1',
  `session_waits` int(4) NOT NULL DEFAULT '-1',
  `dg_stats` varchar(255) NOT NULL DEFAULT '-1',
  `dg_delay` int(10) NOT NULL DEFAULT '-1',
  `processes` int(10) NOT NULL DEFAULT '-1',
  `flashback_on` varchar(10) DEFAULT NULL COMMENT '闪回状态',
  `flashback_usage` varchar(10) DEFAULT NULL COMMENT '闪回空间使用率',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `idx_db_id` (`db_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_oracle_tablespace
-- ----------------------------
DROP TABLE IF EXISTS `pms_oracle_tablespace`;
CREATE TABLE `pms_oracle_tablespace` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL,
  `tablespace_name` varchar(100) NOT NULL,
  `status` varchar(20) NOT NULL,
  `management` varchar(20) NOT NULL,
  `total_size` bigint(18) NOT NULL DEFAULT '0',
  `used_size` bigint(18) NOT NULL DEFAULT '0',
  `max_rate` float(10,2) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `idx_db_id` (`db_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_oracle_tablespace_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_oracle_tablespace_his`;
CREATE TABLE `pms_oracle_tablespace_his` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL,
  `tablespace_name` varchar(100) NOT NULL,
  `status` varchar(20) NOT NULL,
  `management` varchar(20) NOT NULL,
  `total_size` bigint(18) NOT NULL DEFAULT '0',
  `used_size` bigint(18) NOT NULL DEFAULT '0',
  `max_rate` float(10,2) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `idx_db_id` (`db_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_oracle_diskgroup
-- ----------------------------
DROP TABLE IF EXISTS `pms_oracle_diskgroup`;
CREATE TABLE `pms_oracle_diskgroup` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL,
  `diskgroup_name` varchar(100) NOT NULL,
  `state` varchar(20) NOT NULL,
  `type` varchar(20) NOT NULL,
  `total_mb` bigint(18) NOT NULL DEFAULT '0',
  `free_mb` bigint(18) NOT NULL DEFAULT '0',
  `used_rate` float(10,2) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `idx_db_id` (`db_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;



-- ----------------------------
-- Table structure for pms_oracle_diskgroup_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_oracle_diskgroup_his`;
CREATE TABLE `pms_oracle_diskgroup_his` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL,
  `diskgroup_name` varchar(100) NOT NULL,
  `state` varchar(20) NOT NULL,
  `type` varchar(20) NOT NULL,
  `total_mb` bigint(18) NOT NULL DEFAULT '0',
  `free_mb` bigint(18) NOT NULL DEFAULT '0',
  `used_rate` float(10,2) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `idx_db_id` (`db_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_opration
-- ----------------------------
DROP TABLE IF EXISTS `pms_opration`;
CREATE TABLE `pms_opration` (
  `id` bigint(20) NOT NULL,
  `bs_id` int(10) NOT NULL,
  `asset_type` varchar(50) NOT NULL,
  `op_type` varchar(20),
  `result` varchar(2),
  `reason` varchar(1000),
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_op_type` (`bs_id`,`asset_type`,`op_type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_opration_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_opration_his`;
CREATE TABLE `pms_opration_his` (
  `id` bigint(20) NOT NULL,
  `bs_id` int(10) NOT NULL,
  `asset_type` varchar(50) NOT NULL,
  `op_type` varchar(20),
  `result` varchar(2),
  `reason` varchar(1000),
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_op_type` (`bs_id`, `asset_type`, `op_type`,`created`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



-- ----------------------------
-- Table structure for pms_op_process
-- ----------------------------
DROP TABLE IF EXISTS `pms_op_process`;
CREATE TABLE `pms_op_process` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `op_id` bigint(20) NOT NULL,
  `bs_id` int(10) NOT NULL,
  `asset_type` varchar(50) NOT NULL,
  `process_type` varchar(20) COMMENT '2个类型：SWITCHOVER;FAILOVER;',
  `process_desc` varchar(1000),
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_op_type` (`asset_type`, `bs_id`, `process_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_op_process_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_op_process_his`;
CREATE TABLE `pms_op_process_his` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `op_id` bigint(20) NOT NULL,
  `bs_id` int(10) NOT NULL,
  `asset_type` varchar(50) NOT NULL,
  `process_type` varchar(20) COMMENT '2个类型：SWITCHOVER;FAILOVER;',
  `process_desc` varchar(1000),
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_op_type` (`asset_type`, `bs_id`, `process_type`,`created`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `pms_dr_pri_status`;
CREATE TABLE `pms_dr_pri_status` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `check_seq` smallint(4) NOT NULL DEFAULT '0',
  `dest_id` smallint(4) NOT NULL DEFAULT '0',
  `transmit_mode` varchar(20) DEFAULT NULL,
  `thread` smallint(4) NOT NULL DEFAULT '0',
  `sequence` int(10) DEFAULT NULL,
  `curr_scn` bigint(20) DEFAULT NULL,
  `curr_db_time` varchar(20) DEFAULT NULL,
  `archived_delay` int(10) DEFAULT NULL,
  `applied_delay` int(10) DEFAULT NULL,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `pms_dr_pri_status_his`;
CREATE TABLE `pms_dr_pri_status_his` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `check_seq` smallint(4) NOT NULL DEFAULT '0',
  `dest_id` smallint(4) NOT NULL DEFAULT '0',
  `transmit_mode` varchar(20) DEFAULT NULL,
  `thread` smallint(4) NOT NULL DEFAULT '0',
  `sequence` int(10) DEFAULT NULL,
  `curr_scn` bigint(20) DEFAULT NULL,
  `curr_db_time` varchar(20) DEFAULT NULL,
  `archived_delay` int(10) DEFAULT NULL,
  `applied_delay` int(10) DEFAULT NULL,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_db_id` (`db_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `pms_dr_sta_status`;
CREATE TABLE `pms_dr_sta_status` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `thread` smallint(4) NOT NULL,
  `sequence` int(20) DEFAULT NULL,
  `block` int(10) DEFAULT NULL,
  `delay_mins` int(10) DEFAULT NULL,
  `apply_rate` int(10) DEFAULT NULL,
  `curr_scn` bigint(20) DEFAULT NULL,
  `curr_db_time` varchar(20) DEFAULT NULL,
  `mrp_status` varchar(20) DEFAULT NULL,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `pms_dr_sta_status_his`;
CREATE TABLE `pms_dr_sta_status_his` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `thread` smallint(4) NOT NULL,
  `sequence` int(20) DEFAULT NULL,
  `block` int(10) DEFAULT NULL,
  `delay_mins` int(10) DEFAULT NULL,
  `apply_rate` int(10) DEFAULT NULL,
  `curr_scn` bigint(20) DEFAULT NULL,
  `curr_db_time` varchar(20) DEFAULT NULL,
  `mrp_status` varchar(20) DEFAULT NULL,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_db_id` (`db_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_mssql_status
-- ----------------------------
DROP TABLE IF EXISTS `pms_mssql_status`;
CREATE TABLE `pms_mssql_status` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `host` varchar(30) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL DEFAULT '',
  `connect` smallint(4) NOT NULL DEFAULT '0',
  `role` varchar(30) NOT NULL DEFAULT '-1',
  `uptime` varchar(50) NOT NULL DEFAULT '-1',
  `version` varchar(50) NOT NULL DEFAULT '-1',
  `lock_timeout` int(11) NOT NULL DEFAULT '-1',
  `trancount` int(11) NOT NULL DEFAULT '-1',
  `max_connections` int(11) NOT NULL DEFAULT '-1',
  `processes` int(11) NOT NULL DEFAULT '-1',
  `processes_running` int(11) NOT NULL DEFAULT '-1',
  `processes_waits` int(11) NOT NULL DEFAULT '-1',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_mssql_status_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_mssql_status_his`;
CREATE TABLE `pms_mssql_status_his` (
  `id` int(10) NOT NULL,
  `db_id` int(10) NOT NULL,
  `host` varchar(30) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL DEFAULT '',
  `connect` smallint(4) NOT NULL DEFAULT '0',
  `role` varchar(30) NOT NULL DEFAULT '-1',
  `uptime` varchar(50) NOT NULL DEFAULT '-1',
  `version` varchar(50) NOT NULL DEFAULT '-1',
  `lock_timeout` int(11) NOT NULL DEFAULT '-1',
  `trancount` int(11) NOT NULL DEFAULT '-1',
  `max_connections` int(11) NOT NULL DEFAULT '-1',
  `processes` int(11) NOT NULL DEFAULT '-1',
  `processes_running` int(11) NOT NULL DEFAULT '-1',
  `processes_waits` int(11) NOT NULL DEFAULT '-1',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_db_id` (`db_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_dr_mssql_p
-- ----------------------------
DROP TABLE IF EXISTS `pms_dr_mssql_p`;
CREATE TABLE `pms_dr_mssql_p` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dr_id` int(10) NOT NULL DEFAULT '-1',
  `db_id` int(10) NOT NULL DEFAULT '-1',
  `database_id` int(10) NOT NULL DEFAULT '-1',
  `db_name` varchar(30)  DEFAULT NULL,
  `role` tinyint(1) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL,
  `state_desc` varchar(60) DEFAULT NULL,
  `safety_level` tinyint(1) DEFAULT NULL,
  `partner_name` varchar(128) DEFAULT NULL,
  `partner_instance` varchar(128) DEFAULT NULL,
  `failover_lsn` numeric(25,0) DEFAULT NULL,
  `connection_timeout` int(10) DEFAULT NULL,
  `redo_queue` int(10) DEFAULT '-1' COMMENT '-1 表示对镜像服务器重做的最大日志量为UNLIMITED',
  `end_of_log_lsn` numeric(25,0) DEFAULT NULL,
  `replication_lsn` numeric(25,0) DEFAULT NULL,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_dr_mssql_p_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_dr_mssql_p_his`;
CREATE TABLE `pms_dr_mssql_p_his` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dr_id` int(10) NOT NULL DEFAULT '-1',
  `db_id` int(10) NOT NULL DEFAULT '-1',
  `database_id` int(10) NOT NULL DEFAULT '-1',
  `db_name` varchar(30)  DEFAULT NULL,
  `role` tinyint(1) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL,
  `state_desc` varchar(60) DEFAULT NULL,
  `safety_level` tinyint(1) DEFAULT NULL,
  `partner_name` varchar(128) DEFAULT NULL,
  `partner_instance` varchar(128) DEFAULT NULL,
  `failover_lsn` numeric(25,0) DEFAULT NULL,
  `connection_timeout` int(10) DEFAULT NULL,
  `redo_queue` int(10) DEFAULT '-1' COMMENT '-1 表示对镜像服务器重做的最大日志量为UNLIMITED',
  `end_of_log_lsn` numeric(25,0) DEFAULT NULL,
  `replication_lsn` numeric(25,0) DEFAULT NULL,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_dr_id` (`dr_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_dr_mssql_s
-- ----------------------------
DROP TABLE IF EXISTS `pms_dr_mssql_s`;
CREATE TABLE `pms_dr_mssql_s` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dr_id` int(10) NOT NULL DEFAULT '-1',
  `db_id` int(10) NOT NULL DEFAULT '-1',
  `database_id` int(10) NOT NULL DEFAULT '-1',
  `db_name` varchar(30)  DEFAULT NULL,
  `master_server` varchar(30) DEFAULT NULL,
  `master_port` varchar(20) DEFAULT NULL,
  `role` tinyint(1) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL,
  `state_desc` varchar(60) DEFAULT NULL,
  `safety_level` tinyint(1) DEFAULT NULL,
  `partner_name` varchar(128) DEFAULT NULL,
  `partner_instance` varchar(128) DEFAULT NULL,
  `failover_lsn` numeric(25,0) DEFAULT NULL,
  `connection_timeout` int(10) DEFAULT NULL,
  `redo_queue` int(10) DEFAULT '-1' COMMENT '-1 表示对镜像服务器重做的最大日志量为UNLIMITED',
  `end_of_log_lsn` numeric(25,0) DEFAULT NULL,
  `replication_lsn` numeric(25,0) DEFAULT NULL,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_dr_mssql_s_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_dr_mssql_s_his`;
CREATE TABLE `pms_dr_mssql_s_his` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dr_id` int(10) NOT NULL DEFAULT '-1',
  `db_id` int(10) NOT NULL DEFAULT '-1',
  `database_id` int(10) NOT NULL DEFAULT '-1',
  `db_name` varchar(30)  DEFAULT NULL,
  `master_server` varchar(30) DEFAULT NULL,
  `master_port` varchar(20) DEFAULT NULL,
  `role` tinyint(1) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL,
  `state_desc` varchar(60) DEFAULT NULL,
  `safety_level` tinyint(1) DEFAULT NULL,
  `partner_name` varchar(128) DEFAULT NULL,
  `partner_instance` varchar(128) DEFAULT NULL,
  `failover_lsn` numeric(25,0) DEFAULT NULL,
  `connection_timeout` int(10) DEFAULT NULL,
  `redo_queue` int(10) DEFAULT '-1' COMMENT '-1 表示对镜像服务器重做的最大日志量为UNLIMITED',
  `end_of_log_lsn` numeric(25,0) DEFAULT NULL,
  `replication_lsn` numeric(25,0) DEFAULT NULL,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_dr_id` (`dr_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_mysql_status
-- ----------------------------
DROP TABLE IF EXISTS `pms_mysql_status`;
CREATE TABLE `pms_mysql_status` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `db_id` int(10) NOT NULL,
  `host` varchar(30) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL DEFAULT '',
  `connect` smallint(4) NOT NULL DEFAULT '0',
  `role` varchar(30) NOT NULL DEFAULT '-1',
  `uptime` int(11) NOT NULL DEFAULT '-1',
  `version` varchar(50) NOT NULL DEFAULT '-1',
  `max_connections` smallint(4) NOT NULL DEFAULT '-1',
  `max_connect_errors` smallint(4) NOT NULL DEFAULT '-1',
  `open_files_limit` int(10) NOT NULL DEFAULT '-1',
  `open_files` smallint(4) NOT NULL DEFAULT '-1',
  `table_open_cache` smallint(4) NOT NULL DEFAULT '-1',
  `open_tables` smallint(4) NOT NULL DEFAULT '-1',
  `max_tmp_tables` smallint(4) NOT NULL DEFAULT '-1',
  `max_heap_table_size` int(10) NOT NULL DEFAULT '-1',
  `max_allowed_packet` int(10) NOT NULL DEFAULT '-1',
  `threads_connected` int(10) NOT NULL DEFAULT '-1',
  `threads_running` int(10) NOT NULL DEFAULT '-1',
  `threads_waits` int(10) NOT NULL DEFAULT '-1',
  `threads_created` int(10) NOT NULL DEFAULT '-1',
  `threads_cached` int(10) NOT NULL DEFAULT '-1',
  `connections` int(10) NOT NULL DEFAULT '-1',
  `aborted_clients` int(10) NOT NULL DEFAULT '-1',
  `aborted_connects` int(10) NOT NULL DEFAULT '-1',
  `connections_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `bytes_received_persecond` int(10) NOT NULL DEFAULT '-1',
  `bytes_sent_persecond` int(10) NOT NULL DEFAULT '-1',
  `com_select_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_insert_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_update_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_delete_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_commit_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_rollback_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `questions_persecond` int(10) NOT NULL DEFAULT '-1',
  `queries_persecond` int(10) NOT NULL DEFAULT '-1',
  `transaction_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `created_tmp_tables_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `created_tmp_disk_tables_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `created_tmp_files_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `table_locks_immediate_persecond` int(4) NOT NULL DEFAULT '-1',
  `table_locks_waited_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `key_buffer_size` bigint(18) NOT NULL DEFAULT '-1',
  `sort_buffer_size` int(10) NOT NULL DEFAULT '-1',
  `join_buffer_size` int(10) NOT NULL DEFAULT '-1',
  `key_blocks_not_flushed` int(10) NOT NULL DEFAULT '-1',
  `key_blocks_unused` int(10) NOT NULL DEFAULT '-1',
  `key_blocks_used` int(10) NOT NULL DEFAULT '-1',
  `key_read_requests_persecond` int(10) NOT NULL DEFAULT '-1',
  `key_reads_persecond` int(10) NOT NULL DEFAULT '-1',
  `key_write_requests_persecond` int(10) NOT NULL DEFAULT '-1',
  `key_writes_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_version` varchar(30) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_instances` smallint(4) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_size` bigint(18) NOT NULL DEFAULT '-1',
  `innodb_doublewrite` char(10) NOT NULL DEFAULT '-1',
  `innodb_file_per_table` char(10) NOT NULL DEFAULT '-1',
  `innodb_flush_log_at_trx_commit` tinyint(2) NOT NULL DEFAULT '-1',
  `innodb_flush_method` varchar(30) NOT NULL DEFAULT '-1',
  `innodb_force_recovery` tinyint(2) NOT NULL DEFAULT '-1',
  `innodb_io_capacity` int(10) NOT NULL DEFAULT '-1',
  `innodb_read_io_threads` tinyint(2) NOT NULL DEFAULT '-1',
  `innodb_write_io_threads` tinyint(2) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_pages_total` int(10) NOT NULL DEFAULT '-1' COMMENT '页总数目',
  `innodb_buffer_pool_pages_data` int(10) NOT NULL DEFAULT '-1' COMMENT '缓存池中包含数据的页的数目，包括脏页,单位page',
  `innodb_buffer_pool_pages_dirty` int(10) NOT NULL DEFAULT '-1' COMMENT '缓存池中脏页的数目-单位page',
  `innodb_buffer_pool_pages_flushed` bigint(18) NOT NULL DEFAULT '-1' COMMENT '缓存池中刷新页请求的数目-单位page',
  `innodb_buffer_pool_pages_free` int(10) NOT NULL DEFAULT '-1' COMMENT '剩余的页数目-单位是page',
  `innodb_buffer_pool_pages_misc` int(10) NOT NULL DEFAULT '-1' COMMENT '缓存池中当前已经被用作管理用途或hash index而不能用作为普通数据页的数目',
  `innodb_page_size` int(10) NOT NULL DEFAULT '-1',
  `innodb_pages_created` bigint(18) NOT NULL DEFAULT '-1',
  `innodb_pages_read` bigint(18) NOT NULL DEFAULT '-1',
  `innodb_pages_written` bigint(18) NOT NULL DEFAULT '-1',
  `innodb_row_lock_current_waits` varchar(100) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_pages_flushed_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_read_requests_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_reads_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_write_requests_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_rows_read_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_rows_inserted_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_rows_updated_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_rows_deleted_persecond` int(10) NOT NULL DEFAULT '-1',
  `query_cache_hitrate` varchar(10) NOT NULL DEFAULT '-1',
  `thread_cache_hitrate` varchar(10) NOT NULL DEFAULT '-1',
  `key_buffer_read_rate` varchar(10) NOT NULL DEFAULT '-1',
  `key_buffer_write_rate` varchar(10) NOT NULL DEFAULT '-1',
  `key_blocks_used_rate` varchar(10) NOT NULL DEFAULT '-1',
  `created_tmp_disk_tables_rate` varchar(10) NOT NULL DEFAULT '-1',
  `connections_usage_rate` varchar(10) NOT NULL DEFAULT '-1',
  `open_files_usage_rate` varchar(10) NOT NULL DEFAULT '-1',
  `open_tables_usage_rate` varchar(10) NOT NULL DEFAULT '-1',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_mysql_status_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_mysql_status_his`;
CREATE TABLE `pms_mysql_status_his` (
  `id` int(10) NOT NULL,
  `db_id` int(10) NOT NULL,
  `host` varchar(30) NOT NULL,
  `port` int(10) NOT NULL,
  `alias` varchar(50) NOT NULL DEFAULT '',
  `connect` smallint(4) NOT NULL DEFAULT '0',
  `role` varchar(30) NOT NULL DEFAULT '-1',
  `uptime` int(11) NOT NULL DEFAULT '-1',
  `version` varchar(50) NOT NULL DEFAULT '-1',
  `max_connections` smallint(4) NOT NULL DEFAULT '-1',
  `max_connect_errors` smallint(4) NOT NULL DEFAULT '-1',
  `open_files_limit` int(10) NOT NULL DEFAULT '-1',
  `open_files` smallint(4) NOT NULL DEFAULT '-1',
  `table_open_cache` smallint(4) NOT NULL DEFAULT '-1',
  `open_tables` smallint(4) NOT NULL DEFAULT '-1',
  `max_tmp_tables` smallint(4) NOT NULL DEFAULT '-1',
  `max_heap_table_size` int(10) NOT NULL DEFAULT '-1',
  `max_allowed_packet` int(10) NOT NULL DEFAULT '-1',
  `threads_connected` int(10) NOT NULL DEFAULT '-1',
  `threads_running` int(10) NOT NULL DEFAULT '-1',
  `threads_waits` int(10) NOT NULL DEFAULT '-1',
  `threads_created` int(10) NOT NULL DEFAULT '-1',
  `threads_cached` int(10) NOT NULL DEFAULT '-1',
  `connections` int(10) NOT NULL DEFAULT '-1',
  `aborted_clients` int(10) NOT NULL DEFAULT '-1',
  `aborted_connects` int(10) NOT NULL DEFAULT '-1',
  `connections_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `bytes_received_persecond` int(10) NOT NULL DEFAULT '-1',
  `bytes_sent_persecond` int(10) NOT NULL DEFAULT '-1',
  `com_select_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_insert_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_update_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_delete_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_commit_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `com_rollback_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `questions_persecond` int(10) NOT NULL DEFAULT '-1',
  `queries_persecond` int(10) NOT NULL DEFAULT '-1',
  `transaction_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `created_tmp_tables_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `created_tmp_disk_tables_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `created_tmp_files_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `table_locks_immediate_persecond` int(4) NOT NULL DEFAULT '-1',
  `table_locks_waited_persecond` smallint(4) NOT NULL DEFAULT '-1',
  `key_buffer_size` bigint(18) NOT NULL DEFAULT '-1',
  `sort_buffer_size` int(10) NOT NULL DEFAULT '-1',
  `join_buffer_size` int(10) NOT NULL DEFAULT '-1',
  `key_blocks_not_flushed` int(10) NOT NULL DEFAULT '-1',
  `key_blocks_unused` int(10) NOT NULL DEFAULT '-1',
  `key_blocks_used` int(10) NOT NULL DEFAULT '-1',
  `key_read_requests_persecond` int(10) NOT NULL DEFAULT '-1',
  `key_reads_persecond` int(10) NOT NULL DEFAULT '-1',
  `key_write_requests_persecond` int(10) NOT NULL DEFAULT '-1',
  `key_writes_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_version` varchar(30) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_instances` smallint(4) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_size` bigint(18) NOT NULL DEFAULT '-1',
  `innodb_doublewrite` char(10) NOT NULL DEFAULT '-1',
  `innodb_file_per_table` char(10) NOT NULL DEFAULT '-1',
  `innodb_flush_log_at_trx_commit` tinyint(2) NOT NULL DEFAULT '-1',
  `innodb_flush_method` varchar(30) NOT NULL DEFAULT '-1',
  `innodb_force_recovery` tinyint(2) NOT NULL DEFAULT '-1',
  `innodb_io_capacity` int(10) NOT NULL DEFAULT '-1',
  `innodb_read_io_threads` tinyint(2) NOT NULL DEFAULT '-1',
  `innodb_write_io_threads` tinyint(2) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_pages_total` int(10) NOT NULL DEFAULT '-1' COMMENT '页总数目',
  `innodb_buffer_pool_pages_data` int(10) NOT NULL DEFAULT '-1' COMMENT '缓存池中包含数据的页的数目，包括脏页,单位page',
  `innodb_buffer_pool_pages_dirty` int(10) NOT NULL DEFAULT '-1' COMMENT '缓存池中脏页的数目-单位page',
  `innodb_buffer_pool_pages_flushed` bigint(18) NOT NULL DEFAULT '-1' COMMENT '缓存池中刷新页请求的数目-单位page',
  `innodb_buffer_pool_pages_free` int(10) NOT NULL DEFAULT '-1' COMMENT '剩余的页数目-单位是page',
  `innodb_buffer_pool_pages_misc` int(10) NOT NULL DEFAULT '-1' COMMENT '缓存池中当前已经被用作管理用途或hash index而不能用作为普通数据页的数目',
  `innodb_page_size` int(10) NOT NULL DEFAULT '-1',
  `innodb_pages_created` bigint(18) NOT NULL DEFAULT '-1',
  `innodb_pages_read` bigint(18) NOT NULL DEFAULT '-1',
  `innodb_pages_written` bigint(18) NOT NULL DEFAULT '-1',
  `innodb_row_lock_current_waits` varchar(100) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_pages_flushed_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_read_requests_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_reads_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_buffer_pool_write_requests_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_rows_read_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_rows_inserted_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_rows_updated_persecond` int(10) NOT NULL DEFAULT '-1',
  `innodb_rows_deleted_persecond` int(10) NOT NULL DEFAULT '-1',
  `query_cache_hitrate` varchar(10) NOT NULL DEFAULT '-1',
  `thread_cache_hitrate` varchar(10) NOT NULL DEFAULT '-1',
  `key_buffer_read_rate` varchar(10) NOT NULL DEFAULT '-1',
  `key_buffer_write_rate` varchar(10) NOT NULL DEFAULT '-1',
  `key_blocks_used_rate` varchar(10) NOT NULL DEFAULT '-1',
  `created_tmp_disk_tables_rate` varchar(10) NOT NULL DEFAULT '-1',
  `connections_usage_rate` varchar(10) NOT NULL DEFAULT '-1',
  `open_files_usage_rate` varchar(10) NOT NULL DEFAULT '-1',
  `open_tables_usage_rate` varchar(10) NOT NULL DEFAULT '-1',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_db_id` (`db_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_os_status
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_status`;
CREATE TABLE `pms_os_status` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) NOT NULL,
  `connect` tinyint(2) NOT NULL DEFAULT '0',
  `hostname` varchar(100) NOT NULL DEFAULT '-1',
  `kernel` varchar(100) NOT NULL DEFAULT '-1',
  `system_date` varchar(50) NOT NULL DEFAULT '-1',
  `system_uptime` varchar(50) NOT NULL DEFAULT '-1',
  `process` int(10) NOT NULL DEFAULT '-1',
  `load_1` decimal(4,2) NOT NULL DEFAULT '-1.00',
  `load_5` decimal(4,2) NOT NULL DEFAULT '-1.00',
  `load_15` decimal(4,2) NOT NULL DEFAULT '-1.00',
  `cpu_user_time` int(4) NOT NULL DEFAULT '-1',
  `cpu_system_time` int(4) NOT NULL DEFAULT '-1',
  `cpu_idle_time` int(4) NOT NULL DEFAULT '-1',
  `swap_total` int(11) NOT NULL DEFAULT '-1',
  `swap_avail` int(11) NOT NULL DEFAULT '-1',
  `mem_total` int(11) NOT NULL DEFAULT '-1',
  `mem_avail` int(11) NOT NULL DEFAULT '-1',
  `mem_free` int(11) NOT NULL DEFAULT '-1',
  `mem_shared` int(11) NOT NULL DEFAULT '-1',
  `mem_buffered` int(11) NOT NULL DEFAULT '-1',
  `mem_cached` int(11) NOT NULL DEFAULT '-1',
  `mem_usage_rate` decimal(4,2) NOT NULL DEFAULT '-1.00',
  `mem_available` int(11) NOT NULL DEFAULT -1,
  `disk_io_reads_total` int(10) NOT NULL DEFAULT -1,
  `disk_io_writes_total` int(10) NOT NULL DEFAULT -1,
  `net_in_bytes_total` bigint(18) NOT NULL DEFAULT -1,
  `net_out_bytes_total` bigint(18) NOT NULL DEFAULT -1,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_os_status_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_status_his`;
CREATE TABLE `pms_os_status_his` (
  `id` int(10) NOT NULL,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) NOT NULL,
  `connect` tinyint(2) NOT NULL DEFAULT '0',
  `hostname` varchar(100) NOT NULL DEFAULT '-1',
  `kernel` varchar(100) NOT NULL DEFAULT '-1',
  `system_date` varchar(50) NOT NULL DEFAULT '-1',
  `system_uptime` varchar(50) NOT NULL DEFAULT '-1',
  `process` int(10) NOT NULL DEFAULT '-1',
  `load_1` decimal(4,2) NOT NULL DEFAULT '-1.00',
  `load_5` decimal(4,2) NOT NULL DEFAULT '-1.00',
  `load_15` decimal(4,2) NOT NULL DEFAULT '-1.00',
  `cpu_user_time` int(4) NOT NULL DEFAULT '-1',
  `cpu_system_time` int(4) NOT NULL DEFAULT '-1',
  `cpu_idle_time` int(4) NOT NULL DEFAULT '-1',
  `swap_total` int(11) NOT NULL DEFAULT '-1',
  `swap_avail` int(11) NOT NULL DEFAULT '-1',
  `mem_total` int(11) NOT NULL DEFAULT '-1',
  `mem_avail` int(11) NOT NULL DEFAULT '-1',
  `mem_free` int(11) NOT NULL DEFAULT '-1',
  `mem_shared` int(11) NOT NULL DEFAULT '-1',
  `mem_buffered` int(11) NOT NULL DEFAULT '-1',
  `mem_cached` int(11) NOT NULL DEFAULT '-1',
  `mem_usage_rate` decimal(4,2) NOT NULL DEFAULT '-1.00',
  `mem_available` int(11) NOT NULL DEFAULT -1,
  `disk_io_reads_total` int(10) NOT NULL DEFAULT -1,
  `disk_io_writes_total` int(10) NOT NULL DEFAULT -1,
  `net_in_bytes_total` bigint(18) NOT NULL DEFAULT -1,
  `net_out_bytes_total` bigint(18) NOT NULL DEFAULT -1,
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_os_id` (`os_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_os_disk
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_disk`;
CREATE TABLE `pms_os_disk` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) DEFAULT NULL,
  `mounted` varchar(50) NOT NULL DEFAULT '0',
  `total_size` bigint(18) NOT NULL DEFAULT '0',
  `used_size` bigint(18) NOT NULL DEFAULT '0',
  `avail_size` bigint(18) NOT NULL DEFAULT '0',
  `used_rate` varchar(10) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_os_disk_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_disk_his`;
CREATE TABLE `pms_os_disk_his` (
  `id` int(10) NOT NULL,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) DEFAULT NULL,
  `mounted` varchar(50) NOT NULL DEFAULT '0',
  `total_size` bigint(18) NOT NULL DEFAULT '0',
  `used_size` bigint(18) NOT NULL DEFAULT '0',
  `avail_size` bigint(18) NOT NULL DEFAULT '0',
  `used_rate` varchar(10) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_os_id` (`os_id`) USING BTREE,
  KEY `idx_created` (`created`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_os_diskio
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_diskio`;
CREATE TABLE `pms_os_diskio` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) DEFAULT NULL,
  `fdisk` varchar(50) NOT NULL DEFAULT '0',
  `disk_io_reads` bigint(18) NOT NULL DEFAULT '0',
  `disk_io_writes` bigint(18) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_os_diskio_tmp
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_diskio_tmp`;
CREATE TABLE `pms_os_diskio_tmp` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) DEFAULT NULL,
  `fdisk` varchar(50) NOT NULL DEFAULT '0',
  `disk_io_reads` bigint(18) NOT NULL DEFAULT '0',
  `disk_io_writes` bigint(18) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_os_diskio_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_diskio_his`;
CREATE TABLE `pms_os_diskio_his` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) DEFAULT NULL,
  `fdisk` varchar(50) NOT NULL DEFAULT '0',
  `disk_io_reads` bigint(18) NOT NULL DEFAULT '0',
  `disk_io_writes` bigint(18) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_os_id` (`os_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_os_net
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_net`;
CREATE TABLE `pms_os_net` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) DEFAULT NULL,
  `if_descr` varchar(50) NOT NULL DEFAULT '0',
  `in_bytes` bigint(18) NOT NULL DEFAULT '0',
  `out_bytes` bigint(18) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for pms_os_net_tmp
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_net_tmp`;
CREATE TABLE `pms_os_net_tmp` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) DEFAULT NULL,
  `if_descr` varchar(50) NOT NULL DEFAULT '0',
  `in_bytes` bigint(18) NOT NULL DEFAULT '0',
  `out_bytes` bigint(18) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for os_net_his
-- ----------------------------
DROP TABLE IF EXISTS `pms_os_net_his`;
CREATE TABLE `pms_os_net_his` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `os_id` int(10) NOT NULL,
  `host` varchar(50) NOT NULL,
  `alias` varchar(100) DEFAULT NULL,
  `if_descr` varchar(50) NOT NULL DEFAULT '0',
  `in_bytes` bigint(18) NOT NULL DEFAULT '0',
  `out_bytes` bigint(18) NOT NULL DEFAULT '0',
  `created` int(10) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_os_id` (`os_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for pms_global_options
-- ----------------------------
DROP TABLE IF EXISTS `pms_global_options`;
CREATE TABLE `pms_global_options` (
  `id`  varchar(50) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `description` varchar(500) DEFAULT NULL,
  KEY `idx_id` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pms_global_options
-- ----------------------------
INSERT INTO `pms_global_options` VALUES ('monitor', '监控', '1', '是否开启全局监控,此项如果关闭则所有项目都不会被监控，下面监控选项都失效');
INSERT INTO `pms_global_options` VALUES ('monitor_mysql', '监控MySQL', '1', '是否开启MySQL状态监控');
INSERT INTO `pms_global_options` VALUES ('monitor_oracle', '监控Oracle', '1', '是否监控Oracle');
INSERT INTO `pms_global_options` VALUES ('monitor_sqlserver', '监控SQLServer', '1', '是否开启SQLServer监控');
INSERT INTO `pms_global_options` VALUES ('monitor_os', '监控OS', '1', '是否开启OS监控');
INSERT INTO `pms_global_options` VALUES ('frequency_monitor', '监控频率', '60', '监控频率');
INSERT INTO `pms_global_options` VALUES ('alert', '告警', '1', '是否开启告警');
INSERT INTO `pms_global_options` VALUES ('send_alert_mail', '发送报警邮件', '1', '是否发送报警邮件');
INSERT INTO `pms_global_options` VALUES ('send_mail_to_list', '告警邮件收件人', '', '报警邮件通知人员');
INSERT INTO `pms_global_options` VALUES ('send_mail_sleep_time', '发送邮件休眠时间', '300', '告警休眠时间(秒)');
INSERT INTO `pms_global_options` VALUES ('mailtype', '邮件类型', 'html', '邮件发送配置:邮件类型');
INSERT INTO `pms_global_options` VALUES ('mailprotocol', '邮件协议', 'smtp', '邮件发送配置:邮件协议');
INSERT INTO `pms_global_options` VALUES ('smtp_host', 'SMTP主机', 'smtp.163.com', '邮件发送配置:邮件主机');
INSERT INTO `pms_global_options` VALUES ('smtp_port', 'SMTP端口', '25', '邮件发送配置:邮件端口');
INSERT INTO `pms_global_options` VALUES ('smtp_user', 'SMTP账号', 'wlblazers', '邮件发送配置:用户');
INSERT INTO `pms_global_options` VALUES ('smtp_pass', 'SMTP密码', '', '邮件发送配置:密码');
INSERT INTO `pms_global_options` VALUES ('smtp_timeout', 'SMTP超时时间', '10', '邮件发送配置:超时时间');
INSERT INTO `pms_global_options` VALUES ('mailfrom', '邮件发件人', 'wlblazers@163.com', '邮件发送配置:发件人');
INSERT INTO `pms_global_options` VALUES ('send_alert_sms', '发送告警短信', '0', '是否发送短信');
INSERT INTO `pms_global_options` VALUES ('send_sms_to_list', '告警短信收件人', '', '短信收件人列表');
INSERT INTO `pms_global_options` VALUES ('send_sms_sleep_time', '发送短信休眠时间', '300', '发送短信休眠时间(分钟)');
INSERT INTO `pms_global_options` VALUES ('send_alarm_wx', '发送告警微信', '0', '是否发送微信');
