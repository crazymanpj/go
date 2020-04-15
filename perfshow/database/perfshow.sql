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
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL DEFAULT '',
  `gametype` varchar(255) NOT NULL DEFAULT '',
  `platform` varchar(255) NOT NULL DEFAULT '',
  `ver` varchar(255) NOT NULL DEFAULT '',
  `isenable` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  perfshow.gameinfo 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `gameinfo` DISABLE KEYS */;
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
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
REPLACE INTO `migrations` (`id_migration`, `name`, `created_at`, `statements`, `rollback_statements`, `status`) VALUES
	(1, 'User_20191121_163140', '2019-11-21 16:36:29', '', '', 'rollback');
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;


-- 导出  表 perfshow.perfdata 结构
CREATE TABLE IF NOT EXISTS `perfdata` (
  `id` int(11) NOT NULL,
  `taskid` int(11) NOT NULL DEFAULT '0',
  `fps_max` double NOT NULL DEFAULT '0',
  `fps_avr` double NOT NULL DEFAULT '0',
  `fps_min` double NOT NULL DEFAULT '0',
  `cpu_max` double NOT NULL DEFAULT '0',
  `cpu_avr` double NOT NULL DEFAULT '0',
  `cpu_min` double NOT NULL DEFAULT '0',
  `mem_max` double NOT NULL DEFAULT '0',
  `mem_avr` double NOT NULL DEFAULT '0',
  `mem_min` double NOT NULL DEFAULT '0',
  `netuse` double NOT NULL DEFAULT '0',
  `battery` varchar(255) NOT NULL DEFAULT '',
  `temperature` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  perfshow.perfdata 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `perfdata` DISABLE KEYS */;
/*!40000 ALTER TABLE `perfdata` ENABLE KEYS */;


-- 导出  表 perfshow.taskinfo 结构
CREATE TABLE IF NOT EXISTS `taskinfo` (
  `id` int(11) NOT NULL,
  `gameid` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `duration` datetime NOT NULL,
  `tester` int(11) NOT NULL DEFAULT '0',
  `time` datetime NOT NULL,
  `interval` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  perfshow.taskinfo 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `taskinfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `taskinfo` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
