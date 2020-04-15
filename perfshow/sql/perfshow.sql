-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.1.45-community - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win32
-- HeidiSQL 版本:                  8.0.0.4458
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出  表 perfshow.gameinfo 结构
CREATE TABLE IF NOT EXISTS `gameinfo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `gametype` varchar(255) NOT NULL DEFAULT '',
  `platform` varchar(255) NOT NULL DEFAULT '',
  `ver` varchar(255) NOT NULL DEFAULT '',
  `isenable` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=98 DEFAULT CHARSET=utf8;

-- 正在导出表  perfshow.gameinfo 的数据：~17 rows (大约)
DELETE FROM `gameinfo`;
/*!40000 ALTER TABLE `gameinfo` DISABLE KEYS */;
INSERT INTO `gameinfo` (`id`, `name`, `gametype`, `platform`, `ver`, `isenable`) VALUES
	(6, 'test006', 'wxgame', 'ios', '1.0.0', 1),
	(76, 'testa', 'test', 'test', '123', 1),
	(83, 'testa', 'test', 'test', '123', 1),
	(84, 'testa', 'test', 'test', '123', 1),
	(85, 'testa', 'test', 'test', '123', 1),
	(86, 'test007', 'h5game', 'wx', '', 0),
	(87, 'a', 'b', 'c', '', 0),
	(88, 'a', 'aa', 'c', '', 0),
	(89, 'c', 'd', 'e', '', 0),
	(90, '1', '2', '3', '4', 0),
	(91, 'a', 'b', 'c', 'd', 0),
	(92, 'aaa', 'bbb', 'ccc', 'ddd', 1),
	(93, '', '', '', '', 1),
	(94, '1', '2', '3', '4', 1),
	(95, '', '', '', '', 1),
	(96, 'acs', '休闲', '微信', '1.2', 1),
	(97, '压标', '休闲', '微信', '11111', 1);
/*!40000 ALTER TABLE `gameinfo` ENABLE KEYS */;


-- 导出  表 perfshow.migrations 结构
CREATE TABLE IF NOT EXISTS `migrations` (
  `id_migration` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'surrogate key',
  `name` varchar(255) DEFAULT NULL COMMENT 'migration name, unique',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'date migrated or rolled back',
  `statements` longtext COMMENT 'SQL statements for this migration',
  `rollback_statements` longtext COMMENT 'SQL statment for rolling back migration',
  `status` enum('update','rollback') DEFAULT NULL COMMENT 'update indicates it is a normal migration while rollback means this migration is rolled back',
  PRIMARY KEY (`id_migration`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- 正在导出表  perfshow.migrations 的数据：~1 rows (大约)
DELETE FROM `migrations`;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
INSERT INTO `migrations` (`id_migration`, `name`, `created_at`, `statements`, `rollback_statements`, `status`) VALUES
	(1, 'User_20191121_163140', '2019-11-21 16:36:29', '', '', 'rollback');
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;


-- 导出  表 perfshow.perfdata 结构
CREATE TABLE IF NOT EXISTS `perfdata` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `taskid` int(11) NOT NULL DEFAULT '0',
  `phone` varchar(50) NOT NULL DEFAULT '' COMMENT '测试机型',
  `fps_max` double NOT NULL DEFAULT '0',
  `fps_avr` double NOT NULL DEFAULT '0',
  `fps_min` double NOT NULL DEFAULT '0',
  `cpu_max` double NOT NULL DEFAULT '0',
  `cpu_avr` double NOT NULL DEFAULT '0',
  `cpu_min` double NOT NULL DEFAULT '0',
  `mem_max` double NOT NULL DEFAULT '0',
  `mem_avr` double NOT NULL DEFAULT '0',
  `mem_min` double NOT NULL DEFAULT '0',
  `net_max` double NOT NULL DEFAULT '0',
  `net_avr` double NOT NULL DEFAULT '0',
  `net_min` double NOT NULL DEFAULT '0',
  `battery_avr` double DEFAULT '0',
  `temp_avr` double DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

-- 正在导出表  perfshow.perfdata 的数据：~8 rows (大约)
DELETE FROM `perfdata`;
/*!40000 ALTER TABLE `perfdata` DISABLE KEYS */;
INSERT INTO `perfdata` (`id`, `taskid`, `phone`, `fps_max`, `fps_avr`, `fps_min`, `cpu_max`, `cpu_avr`, `cpu_min`, `mem_max`, `mem_avr`, `mem_min`, `net_max`, `net_avr`, `net_min`, `battery_avr`, `temp_avr`) VALUES
	(7, 2, '0', 60.6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
	(8, 2, '0', 60.6, 31.2, 21.2, 35.4, 25.4, 15.4, 700.533, 300.533, 100.633, 700.533, 300.533, 100.633, 300.533, 100.633),
	(9, 2, '0', 60.6, 31.2, 21.2, 35.4, 25.4, 15.4, 700.533, 300.533, 100.633, 700.533, 300.533, 100.633, 300.533, 100.633),
	(10, 3, '0', 60.6, 31.2, 21.2, 35.4, 25.4, 15.4, 700.533, 300.533, 100.633, 700.533, 300.533, 100.633, 300.533, 100.633),
	(11, 3, '0', 60.6, 31.2, 21.2, 35.4, 25.4, 15.4, 700.533, 300.533, 100.633, 700.533, 300.533, 100.633, 300.533, 100.633),
	(12, 3, '0', 60.6, 31.2, 21.2, 35.4, 25.4, 15.4, 700.533, 300.533, 100.633, 700.533, 300.533, 100.633, 300.533, 100.633),
	(13, 3, '0', 60.6, 31.2, 21.2, 35.4, 25.4, 15.4, 700.533, 300.533, 100.633, 700.533, 300.533, 100.633, 300.533, 100.633),
	(14, 3, 'xiaomi3', 60.6, 31.2, 21.2, 35.4, 25.4, 15.4, 700.533, 300.533, 100.633, 700.533, 300.533, 100.633, 300.533, 100.633);
/*!40000 ALTER TABLE `perfdata` ENABLE KEYS */;

-- 导出  表 perfshow.perfdata 结构
CREATE TABLE IF NOT EXISTS `perfdatashow` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `taskid` int(11) NOT NULL DEFAULT '0',
  `gameid` int(11) NOT NULL DEFAULT '0',
  `phone` varchar(50) NOT NULL DEFAULT '' COMMENT '测试机型',
  `fps_max` double NOT NULL DEFAULT '0',
  `fps_avr` double NOT NULL DEFAULT '0',
  `fps_min` double NOT NULL DEFAULT '0',
  `cpu_max` double NOT NULL DEFAULT '0',
  `cpu_avr` double NOT NULL DEFAULT '0',
  `cpu_min` double NOT NULL DEFAULT '0',
  `mem_max` double NOT NULL DEFAULT '0',
  `mem_avr` double NOT NULL DEFAULT '0',
  `mem_min` double NOT NULL DEFAULT '0',
  `net_max` double NOT NULL DEFAULT '0',
  `net_avr` double NOT NULL DEFAULT '0',
  `net_min` double NOT NULL DEFAULT '0',
  `battery_avr` double DEFAULT '0',
  `temp_avr` double DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;



-- 导出  表 perfshow.taskinfo 结构
CREATE TABLE IF NOT EXISTS `taskinfo` (
  `id` int(11) NOT NULL,
  `gameid` int(11) NOT NULL DEFAULT '0',
  `phone` varchar(255) NOT NULL COMMENT '手机设备号或IP地址',
  `duration` int(11) NOT NULL DEFAULT '0',
  `tester` int(11) NOT NULL DEFAULT '0',
  `time` datetime NOT NULL,
  `interval` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  perfshow.taskinfo 的数据：~6 rows (大约)
DELETE FROM `taskinfo`;
/*!40000 ALTER TABLE `taskinfo` DISABLE KEYS */;
INSERT INTO `taskinfo` (`id`, `gameid`, `phone`, `duration`, `tester`, `time`, `interval`) VALUES
	(3, 45, 'xiaomi', 60, 0, '2019-11-24 13:06:03', 1),
	(33, 45, 'xiaomi', 60, 0, '2019-11-25 03:46:46', 1),
	(34, 45, 'xiaomi', 60, 0, '2019-11-25 03:47:16', 1),
	(35, 45, 'xiaomi', 60, 0, '2019-11-25 07:18:10', 1),
	(37, 45, '10.12.130.131', 60, 0, '2019-11-26 07:49:33', 1),
	(38, 45, '7hrf347fhrefhd', 60, 0, '2019-11-26 07:49:57', 1);
/*!40000 ALTER TABLE `taskinfo` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
