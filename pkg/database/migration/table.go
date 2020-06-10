/**
 * Created by zc on 2020/6/3.
 */
package migration

func tables() map[string]string {
	return map[string]string{
		appTableName:              appTable,
		configTableName:           configTable,
		configRecordTableName:     configRecordTable,
		configRuleRelateTableName: configRuleRelateTable,
		ruleTableName:             ruleTable,
		ruleRecordTableName:       ruleRecordTable,
		spaceTableName:            spaceTable,
		spaceAccessTableName:      spaceAccessTable,
		spaceRuleTableName:        spaceRuleTable,
		userTableName:             userTable,
	}
}

const (
	appTableName              = "app"
	configTableName           = "config"
	configRecordTableName     = "config_record"
	configRuleRelateTableName = "config_rule_relate"
	ruleTableName             = "rule"
	ruleRecordTableName       = "rule_record"
	spaceTableName            = "space"
	spaceAccessTableName      = "space_access"
	spaceRuleTableName        = "space_rule"
	userTableName             = "user"
)

const (
	options      = ` ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	createdAt    = `created_at datetime DEFAULT NULL,` + "\n  "
	updatedAt    = `updated_at datetime DEFAULT NULL,` + "\n  "
	deletedAt    = `deleted_at datetime DEFAULT NULL,` + "\n"
	deletedAtKey = `KEY idx_app_deleted_at (deleted_at)` + "\n"
)

const (
	appTable = `CREATE TABLE ` + appTableName + ` (
  app_id varchar(32) NOT NULL COMMENT '应用唯一标识',
  sid varchar(32) NOT NULL COMMENT '空间唯一标识',
  title varchar(255) NOT NULL COMMENT '标题',
  description varchar(255) DEFAULT NULL COMMENT '描述',
  ` + createdAt + updatedAt + deletedAt + `
  PRIMARY KEY (app_id),
  UNIQUE KEY uix_app_app_id (app_id),
  ` + deletedAtKey + `)`

	configTable = `CREATE TABLE ` + configTableName + ` (
  config_id varchar(32) NOT NULL COMMENT '配置唯一标识',
  title varchar(255) NOT NULL COMMENT '标题',
  content text NOT NULL COMMENT '内容',
  version varchar(20) NOT NULL COMMENT '版本',
  ` + createdAt + updatedAt + deletedAt + `
  PRIMARY KEY (config_id),
  UNIQUE KEY uix_config_config_id (config_id),
  ` + deletedAtKey + `)`

	configRecordTable = `CREATE TABLE ` + configRecordTableName + ` (
  config_id varchar(32) NOT NULL COMMENT '配置唯一标识',
  content text NOT NULL COMMENT '内容',
  version varchar(20) NOT NULL COMMENT '版本',
  status tinyint(4) DEFAULT NULL COMMENT '状态',
  create_by varchar(32) NOT NULL,
  ` + createdAt + `
  PRIMARY KEY (config_id)` + "\n" + `)`

	configRuleRelateTable = `CREATE TABLE ` + configRuleRelateTableName + ` (
  config_id varchar(32) NOT NULL COMMENT '配置标识',
  space_rule_id varchar(32) NOT NULL COMMENT '空间规则标识'` + "\n" + `)`

	ruleTable = `CREATE TABLE ` + ruleTableName + ` (
  rule_id varchar(32) NOT NULL COMMENT '规则唯一标识',
  uid varchar(32) NOT NULL COMMENT '用户唯一标识',
  title varchar(255) NOT NULL COMMENT '标题',
  description varchar(255) DEFAULT NULL COMMENT '描述',
  content text NOT NULL COMMENT '内容',
  version varchar(20) NOT NULL COMMENT '版本',
  ` + createdAt + updatedAt + deletedAt + `
  PRIMARY KEY (rule_id),
  UNIQUE KEY uix_rule_rule_id (rule_id),
  ` + deletedAtKey + `)`

	ruleRecordTable = `CREATE TABLE ` + ruleRecordTableName + ` (
  rule_id varchar(32) NOT NULL COMMENT '规则唯一标识',
  content text NOT NULL COMMENT '内容',
  version varchar(20) NOT NULL COMMENT '版本',
  ` + createdAt + `
  PRIMARY KEY (rule_id)` + "\n" + `)`

	spaceTable = `CREATE TABLE ` + spaceTableName + ` (
  sid varchar(32) NOT NULL COMMENT '空间唯一标识',
  title varchar(255) NOT NULL COMMENT '标题',
  owner varchar(32) NOT NULL COMMENT '拥有者标识',
  ` + createdAt + updatedAt + deletedAt + `
  PRIMARY KEY (sid),
  UNIQUE KEY uix_space_sid (sid),
  ` + deletedAtKey + `)`

	spaceAccessTable = `CREATE TABLE ` + spaceAccessTableName + ` (
  sid varchar(32) NOT NULL COMMENT '空间唯一标识',
  uid varchar(32) NOT NULL COMMENT '用户唯一标识',
  access varchar(255) DEFAULT NULL COMMENT '权限',
  ` + createdAt + updatedAt + `
  PRIMARY KEY (sid)` + "\n" + `)`

	spaceRuleTable = `CREATE TABLE ` + spaceRuleTableName + ` (
  space_rule_id varchar(32) NOT NULL COMMENT '空间规则唯一标识',
  sid varchar(32) NOT NULL COMMENT '空间唯一标识',
  title varchar(255) NOT NULL COMMENT '标题',
  description varchar(255) DEFAULT NULL COMMENT '描述',
  content text NOT NULL COMMENT '内容',
  create_from varchar(32) NOT NULL COMMENT '来自用户标识',
  ` + createdAt + updatedAt + deletedAt + `
  PRIMARY KEY (space_rule_id),
  UNIQUE KEY uix_space_rule_space_rule_id (space_rule_id),
  ` + deletedAtKey + `)`

	userTable = `CREATE TABLE ` + userTableName + ` (
  uid varchar(32) NOT NULL COMMENT '用户唯一标识',
  username varchar(50) NOT NULL COMMENT '用户名',
  email varchar(100) NOT NULL COMMENT '邮箱',
  avatar varchar(255) DEFAULT NULL COMMENT '头像',
  pwd varchar(100) NOT NULL COMMENT '密码',
  salt varchar(10) NOT NULL COMMENT '盐值',
  ` + createdAt + updatedAt + deletedAt + `
  PRIMARY KEY (uid),
  UNIQUE KEY uix_user_uid (uid),
  ` + deletedAtKey + `)`
)
