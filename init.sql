/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`gormtests` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `gormtests`;

/*Table structure for table `userinfos` */

DROP TABLE IF EXISTS `userinfos`;

CREATE TABLE `userinfos` (
  `class` varchar(10) NOT NULL COMMENT '班级',
  `no` int(11) NOT NULL COMMENT '学号',
  `name` varchar(10) NOT NULL COMMENT '名称',
  `sex` varchar(1) NOT NULL COMMENT '性别',
  `age` int(11) NOT NULL COMMENT '年龄',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`class`,`no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `userinfos` */

insert  into `userinfos`(`class`,`no`,`name`,`sex`,`age`,`create_time`,`update_time`) values ('一一班',1,'zhang','男',6,'2023-05-08 21:40:22','2023-05-08 21:40:24'),('一一班',2,'zhao','男',6,'2023-05-08 21:40:22','2023-05-08 21:40:24'),('一一班',3,'qian','男',7,'2023-05-08 21:40:22','2023-05-08 21:40:24'),('一一班',4,'sun','女',6,'2023-05-08 21:40:22','2023-05-08 21:40:24'),('一二班',1,'zhang','女',6,'2023-05-08 21:40:22','2023-05-08 21:40:24'),('一二班',2,'li','男',7,'2023-05-08 21:43:04','2023-05-08 21:43:06'),('一二班',3,'zhou','男',7,'2023-05-08 21:43:04','2023-05-08 21:43:06'),('一二班',4,'wu','男',7,'2023-05-08 21:43:04','2023-05-08 21:43:06');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
