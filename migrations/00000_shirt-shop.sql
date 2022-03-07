-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.28 - MySQL Community Server - GPL
-- Server OS:                    Linux
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for shirt-shop
CREATE DATABASE IF NOT EXISTS `shirt-shop` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `shirt-shop`;

-- Dumping structure for table shirt-shop.order
CREATE TABLE IF NOT EXISTS `order` (
  `order_id` int NOT NULL AUTO_INCREMENT,
  `order_datetime` datetime DEFAULT NULL,
  `order_status` int DEFAULT NULL,
  `order_user_id` int DEFAULT NULL,
  `order_addr` int DEFAULT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=317 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.order: ~7 rows (approximately)
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
/*!40000 ALTER TABLE `order` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.order_status
CREATE TABLE IF NOT EXISTS `order_status` (
  `order_status_id` int NOT NULL AUTO_INCREMENT,
  `order_status_status` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`order_status_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.order_status: ~4 rows (approximately)
/*!40000 ALTER TABLE `order_status` DISABLE KEYS */;
REPLACE INTO `order_status` (`order_status_id`, `order_status_status`) VALUES
	(1, 'กำลังตรวจสอบการชำระเงิน'),
	(2, 'จัดเตรียมสินค้าเรียบร้อย'),
	(3, 'สินค้าอยู่ขั้นตอนจัดส่ง'),
	(4, 'ส่งค้าส่งเรียบร้อยแล้ว');
/*!40000 ALTER TABLE `order_status` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.product
CREATE TABLE IF NOT EXISTS `product` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `catagory` bigint DEFAULT NULL,
  `size` bigint DEFAULT NULL,
  `gender` bigint DEFAULT NULL,
  `price` float unsigned zerofill DEFAULT NULL,
  `quantiry` float DEFAULT NULL,
  `ispadding` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.product: ~0 rows (approximately)
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
REPLACE INTO `product` (`id`, `name`, `catagory`, `size`, `gender`, `price`, `quantiry`, `ispadding`) VALUES
	(1, 'shirt', 1, 11, 111, 000000001111, 11111, NULL),
	(2, 'shirt-2', 2, 22, 222, 000000002222, 22222, NULL),
	(3, 'กางเกงขายาว', 2, 3, 4, 000000000005, 6, NULL),
	(4, 'กางเกงขายาว', 2, 3, 4, 000000000005, 6, NULL),
	(5, 'กางเกงขายาว', 2, 3, 4, 000000000005, 6, NULL),
	(6, 'กางเกงขายาว', 2, 3, 4, 000000000005, 6, NULL),
	(7, 'กางเกงขายาว', 2, 3, 4, 000000000005, 6, NULL),
	(8, 'กางเกงขายาว', 2, 3, 4, 000000000005, 6, NULL),
	(9, 'กางเกงขายาว', 2, 3, 4, 000000000005, 6, NULL);
/*!40000 ALTER TABLE `product` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.product_catagory
CREATE TABLE IF NOT EXISTS `product_catagory` (
  `product_catagory_id` int NOT NULL AUTO_INCREMENT,
  `product_catagory_name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`product_catagory_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.product_catagory: ~3 rows (approximately)
/*!40000 ALTER TABLE `product_catagory` DISABLE KEYS */;
REPLACE INTO `product_catagory` (`product_catagory_id`, `product_catagory_name`) VALUES
	(1, 'Plain Color'),
	(2, 'Pattern'),
	(3, 'Figure');
/*!40000 ALTER TABLE `product_catagory` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.product_gender
CREATE TABLE IF NOT EXISTS `product_gender` (
  `product_gender_id` int NOT NULL AUTO_INCREMENT,
  `product_gender_name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`product_gender_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.product_gender: ~2 rows (approximately)
/*!40000 ALTER TABLE `product_gender` DISABLE KEYS */;
REPLACE INTO `product_gender` (`product_gender_id`, `product_gender_name`) VALUES
	(1, 'ชาย'),
	(2, 'หญิง');
/*!40000 ALTER TABLE `product_gender` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.product_size
CREATE TABLE IF NOT EXISTS `product_size` (
  `product_size_id` int NOT NULL AUTO_INCREMENT,
  `product_size_name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`product_size_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.product_size: ~8 rows (approximately)
/*!40000 ALTER TABLE `product_size` DISABLE KEYS */;
REPLACE INTO `product_size` (`product_size_id`, `product_size_name`) VALUES
	(1, 'S'),
	(2, 'M'),
	(3, 'L'),
	(4, 'XL'),
	(5, '2XL'),
	(6, '3XL'),
	(7, '4XL'),
	(8, '5XL');
/*!40000 ALTER TABLE `product_size` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.role
CREATE TABLE IF NOT EXISTS `role` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.role: ~2 rows (approximately)
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
REPLACE INTO `role` (`id`, `name`) VALUES
	(1, 'admin'),
	(2, 'user');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.roles
CREATE TABLE IF NOT EXISTS `roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.roles: ~0 rows (approximately)
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.user
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL,
  `surname` varchar(128) DEFAULT NULL,
  `username` varchar(128) DEFAULT NULL,
  `hash` varchar(2048) DEFAULT NULL,
  `rold` bigint DEFAULT NULL,
  `role` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.user: ~0 rows (approximately)
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
REPLACE INTO `user` (`id`, `name`, `surname`, `username`, `hash`, `rold`, `role`) VALUES
	(1, 'บัณฑิต', 'คุ้มสวัสดิ์', 'bandid12', '44s4455ddd', 1, NULL),
	(2, 'มานี', 'มีนา', 'mena', '$2a$08$hAm86bcUwEJOzcG8NKjh.OyOH3Emx2voCQltsZpLBsr1VZlCzP6mi', 2, NULL),
	(3, 'มานีมา', 'มีนามา', 'menana', '$2a$08$wqGi9id4.2SpDat4RRo1/e0SkRBWUozhBYkh.54Y8Z3Mj7D2OyHla', 2, NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
