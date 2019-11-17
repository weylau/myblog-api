CREATE TABLE `myblog`.`mb_articles` (
  `article_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `cate_id` int(10) NOT NULL DEFAULT '0' COMMENT '所属分类ID',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '标题',
  `description` varchar(512) NOT NULL DEFAULT '' COMMENT '简介、描叙',
  `keywords` varchar(256) NOT NULL DEFAULT '' COMMENT '关键词，用英文逗号隔开',
  `img_path` varchar(256) NOT NULL DEFAULT '' COMMENT '图片',
  `op_id` int(10) NOT NULL DEFAULT '0' COMMENT '操作人id',
  `op_user` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人显示帐号',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='文章表';

CREATE TABLE `myblog`.`mb_articles_contents` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `article_id` int(10) NOT NULL DEFAULT '0' COMMENT '文章ID',
  `show_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '内容展示类型：1-html、2-markdown',
  `contents` text NOT NULL COMMENT '文章内容',
  PRIMARY KEY (`id`),
  KEY `idx_article_id` (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='文章内容表';

CREATE TABLE `myblog`.`mb_articles_cate` (
  `cate_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `name` varchar (64) NOT NULL DEFAULT '' COMMENT '分类名',
  `c_name` varchar(64) NOT NULL DEFAULT '' COMMENT '分类中文名',
  `description` varchar(512) NOT NULL DEFAULT '' COMMENT '简介、描叙',
  `parent_id` int(10) NOT NULL DEFAULT '0' COMMENT '上级ID',
  `orderby` int(10) NOT NULL DEFAULT '0' COMMENT '排序',
  `op_id` int(10) NOT NULL DEFAULT '0' COMMENT '操作人id',
  `op_user` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人显示帐号',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`cate_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章类型表';

CREATE TABLE `myblog`.`mb_admins` (
  `admin_id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：1-正常，2-禁用',
  `op_id` int(10) NOT NULL DEFAULT '0' COMMENT '操作人id',
  `op_user` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人显示帐号',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`admin_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='管理员表';

insert into `myblog`.`mb_articles_cate`(cate_id,`name`, c_name) value (1,'php','PHP');
insert into `myblog`.`mb_articles_cate`(cate_id,`name`, c_name) value (2,'mysql','Mysql');
insert into `myblog`.`mb_articles_cate`(cate_id,`name`, c_name) value (3,'go','Go');
insert into `myblog`.`mb_articles_cate`(cate_id,`name`, c_name) value (4,'linux','Linux');
insert into `myblog`.`mb_articles_cate`(cate_id,`name`, c_name) value (5,'qianduan','前端');
insert into `myblog`.`mb_articles_cate`(cate_id,`name`, c_name) value (6,'other','其他');


insert into `myblog`.`mb_articles`(cate_id, title, description) value (1, '测试文章01','测试文章01');
insert into `myblog`.`mb_articles`(cate_id, title, description) value (1, '测试文章02','测试文章02');
insert into `myblog`.`mb_articles`(cate_id, title, description) value (1, '测试文章03','测试文章03');


insert into `myblog`.`mb_articles_contents`(article_id, contents) value (1,'测试文章01');
insert into `myblog`.`mb_articles_contents`(article_id, contents) value (2,'测试文章02');
insert into `myblog`.`mb_articles_contents`(article_id, contents) value (3,'测试文章03');

insert into `myblog`.`mb_admins` (username,password,status) value ('admin', '',1)

alter table `myblog`.`mb_articles_contents` ADD COLUMN `show_type` tinyint(2) NOT NULL DEFAULT 0 COMMENT '内容展示类型：1-html、2-markdown' AFTER `article_id`;
alter table `myblog`.`mb_articles` ADD COLUMN `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '状态：1-展示、2-隐藏' AFTER `title`;

