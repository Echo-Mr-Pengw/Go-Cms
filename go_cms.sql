-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2021-01-06 11:47:37
-- 服务器版本： 5.7.28
-- PHP 版本： 7.3.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `go_cms`
--

-- --------------------------------------------------------

--
-- 表的结构 `tbl_admin`
--

CREATE TABLE `tbl_admin` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '主键',
  `name` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '管理员姓名',
  `password` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `grade` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '管理员等级 1:超管 2:普通',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态 1:正常 2:冻结',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `tbl_admin`
--

INSERT INTO `tbl_admin` (`id`, `name`, `password`, `grade`, `status`, `create_time`) VALUES
(1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 2, 1, '2020-12-24 22:54:31');

-- --------------------------------------------------------

--
-- 表的结构 `tbl_article`
--

CREATE TABLE `tbl_article` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '主键',
  `tag_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '标签id',
  `title` varchar(648) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '文章标题',
  `abstract` varchar(128) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '文章摘要',
  `content` longtext COLLATE utf8_unicode_ci NOT NULL COMMENT '文章内容',
  `author` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '文章作者',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态 1:正常 2:冻结',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `tbl_article_read_num`
--

CREATE TABLE `tbl_article_read_num` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '主键',
  `article_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '文章id',
  `read_num` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '阅读量',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态 1:正常 2:冻结',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `tbl_article_tag`
--

CREATE TABLE `tbl_article_tag` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '主键',
  `name` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '标签名',
  `total` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '标签下文章的总数',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态 1:正常 2:冻结',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `tbl_links`
--

CREATE TABLE `tbl_links` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '主键',
  `name` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '友情名',
  `url` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '链接地址',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态 1:正常 2:冻结',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `tbl_profile`
--

CREATE TABLE `tbl_profile` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '主键',
  `nickname` varchar(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `motto` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '座右铭',
  `wx_qq_img` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '个人微信/qq图片',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态 1:正常 2:冻结',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `tbl_profile`
--

INSERT INTO `tbl_profile` (`id`, `nickname`, `motto`, `wx_qq_img`, `status`, `create_time`) VALUES
(1, 'Gophper', '加油！！', '', 1, '2020-12-29 09:26:28');

--
-- 转储表的索引
--

--
-- 表的索引 `tbl_admin`
--
ALTER TABLE `tbl_admin`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `tbl_article`
--
ALTER TABLE `tbl_article`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `tbl_article_read_num`
--
ALTER TABLE `tbl_article_read_num`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `article_id` (`article_id`);

--
-- 表的索引 `tbl_article_tag`
--
ALTER TABLE `tbl_article_tag`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `tbl_links`
--
ALTER TABLE `tbl_links`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `tbl_profile`
--
ALTER TABLE `tbl_profile`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `tbl_admin`
--
ALTER TABLE `tbl_admin`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `tbl_article`
--
ALTER TABLE `tbl_article`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `tbl_article_read_num`
--
ALTER TABLE `tbl_article_read_num`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `tbl_article_tag`
--
ALTER TABLE `tbl_article_tag`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `tbl_links`
--
ALTER TABLE `tbl_links`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `tbl_profile`
--
ALTER TABLE `tbl_profile`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
