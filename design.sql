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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表';

CREATE TABLE `myblog`.`mb_articles_contents` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `article_id` int(10) NOT NULL DEFAULT '0' COMMENT '文章ID',
  `contents` text NOT NULL  COMMENT '文章内容',
  PRIMARY KEY (`id`),
  KEY `idx_article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章内容表';

CREATE TABLE `myblog`.`mb_articles_cate` (
  `cate_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `name` varchar (64) NOT NULL DEFAULT '' COMMENT '分类名',
  `c_name` varchar(64) NOT NULL DEFAULT '' COMMENT '分类中文名',
  `description` varchar(512) NOT NULL DEFAULT '' COMMENT '简介、描叙',
  `parent_id` int(10) NOT NULL DEFAULT '0' COMMENT '上级ID',
  `op_id` int(10) NOT NULL DEFAULT '0' COMMENT '操作人id',
  `op_user` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人显示帐号',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`cate_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表';

CREATE TABLE `myblog`.`mb_admins` (
  `admin_id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar (64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：1-正常，2-禁用',
  `op_id` int(10) NOT NULL DEFAULT '0' COMMENT '操作人id',
  `op_user` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人显示帐号',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员表';

insert into `myblog`.`mb_articles_cate`(`name`, c_name) value ('php','php');
insert into `myblog`.`mb_articles_cate`(`name`, c_name) value ('golang','golang');
insert into `myblog`.`mb_articles_cate`(`name`, c_name) value ('linux','linux');
insert into `myblog`.`mb_articles_cate`(`name`, c_name) value ('qianduan','前端');


insert into `myblog`.`mb_articles`(cate_id, title, description) value (1, '测试文章01','测试文章01');
insert into `myblog`.`mb_articles`(cate_id, title, description) value (1, '测试文章02','测试文章02');
insert into `myblog`.`mb_articles`(cate_id, title, description) value (1, '测试文章03','测试文章03');


insert into `myblog`.`mb_articles_contents`(article_id, contents) value (1,'测试文章01');
insert into `myblog`.`mb_articles_contents`(article_id, contents) value (2,'测试文章02');
insert into `myblog`.`mb_articles_contents`(article_id, contents) value (3,'测试文章03');

insert into `myblog`.`mb_admins` (username,password,status) value ('admin', '',1)