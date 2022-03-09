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

-- Dumping data for table shirt-shop.order: ~0 rows (approximately)
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
/*!40000 ALTER TABLE `order` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.order_item
CREATE TABLE IF NOT EXISTS `order_item` (
  `order_item_id` int NOT NULL AUTO_INCREMENT,
  `order_item_quantity` int DEFAULT NULL,
  `order_item_product_id` int DEFAULT NULL,
  `order_id` int DEFAULT NULL,
  PRIMARY KEY (`order_item_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3019 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.order_item: ~13 rows (approximately)
/*!40000 ALTER TABLE `order_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_item` ENABLE KEYS */;

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
  `name` varchar(4096) DEFAULT NULL,
  `catagory` bigint DEFAULT NULL,
  `size` bigint DEFAULT NULL,
  `gender` bigint DEFAULT NULL,
  `price` double unsigned zerofill DEFAULT NULL,
  `quantiry` float DEFAULT NULL,
  `ispadding` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.product: ~27 rows (approximately)
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
REPLACE INTO `product` (`id`, `name`, `catagory`, `size`, `gender`, `price`, `quantiry`, `ispadding`) VALUES
	(1, 'shirt', 1, 1, 2, 0000000000000000001111, 11111, 0),
	(2, 'shirt-2', 2, 1, 2, 0000000000000000002222, 22222, 0),
	(3, 'กางเกงขายาว', 2, 3, 1, 0000000000000000000005, 6, 0),
	(4, 'กางเกงขายาว', 2, 3, 1, 0000000000000000000005, 6, 0),
	(5, 'กางเกงขายาว', 2, 3, 1, 0000000000000000000005, 6, 0),
	(6, 'กางเกงขายาว', 2, 3, 1, 0000000000000000000005, 6, 0),
	(7, 'กางเกงขายาว', 2, 3, 1, 00000522.0999755859375, 6, 0),
	(8, 'กางเกงขายาว', 2, 3, 2, 0000000000000000000005, 6, 0),
	(9, 'กางเกงขายาว', 2, 3, 2, 0000000000000000000005, 6, 0),
	(10, 'เสื้อแขนยาว 1', 1, 1, 1, 0000000000000001125.54, 50, 0),
	(11, 'เสื้อแขนยาว 2', 1, 1, 1, 0000000000000001125.54, 50, 0),
	(12, 'เสื้อแขนยาว 3', 1, 4, 1, 0000000000000001125.54, 50, 0),
	(13, 'เสื้อแขนยาว 4', 1, 4, 1, 0000000000000001125.54, 50, 0),
	(14, 'เสื้อแขนยาว 5', 1, 4, 1, 0000000000000001125.54, 50, 0),
	(15, 'เสื้อแขนยาว 6', 2, 4, 2, 0000000000000001125.54, 50, 0),
	(16, 'เสื้อแขนยาว 7', 2, 4, 2, 0000000000000001125.54, 50, 0),
	(17, 'เสื้อแขนยาว 8', 2, 4, 2, 0000000000000001125.54, 50, 0),
	(18, 'เสื้อแขนยาว 9', 2, 4, 2, 0000000000000001125.54, 50, 0),
	(19, 'เสื้อแขนยาว 10', 2, 4, 2, 0000000000000001125.54, 50, 0),
	(20, 'เสื้อแขนสั้น 1', 2, 4, 2, 0000000000000001125.54, 50, 0),
	(21, 'เสื้อแขนสั้น 2', 2, 4, 2, 0000000000000001125.54, 50, 0),
	(22, 'เสื้อแขนสั้น 2', 2, 2, 2, 0000000000000001125.54, 50, 0),
	(23, 'เสื้อแขนสั้น 3', 2, 1, 2, 0000000000000001125.54, 50, 0),
	(24, 'เสื้อแขนสั้น 4', 2, 1, 2, 0000000000000001125.54, 50, 0),
	(25, 'เสื้อแขนสั้น 5', 2, 1, 2, 0000000000000001125.54, 50, 0),
	(26, 'เสื้อแขนสั้น 6', 2, 1, 2, 0000000000000001125.54, 50, 0),
	(27, 'เสื้อสีขาว 5', 2, 2, 2, 0000000000000000000100, 50, 1),
	(28, 'เสื้อแขนสั้น 7', 2, 1, 1, 0000000000000001125.54, 50, 0),
	(29, 'เสื้อแขนสั้น 7', 1, 1, 1, 0000000000000001125.54, 50, 0),
	(30, 'เสื้อแขนสั้น 8', 1, 1, 1, 0000000000000001125.54, 50, 0);
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

-- Dumping structure for table shirt-shop.user
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL,
  `surname` varchar(128) DEFAULT NULL,
  `username` varchar(128) DEFAULT NULL,
  `hash` varchar(2048) DEFAULT NULL,
  `role` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.user: ~3 rows (approximately)
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
REPLACE INTO `user` (`id`, `name`, `surname`, `username`, `hash`, `role`) VALUES
	(1, 'บัณฑิต', 'คุ้มสวัสดิ์', 'bandid12', '44s4455ddd', 1),
	(2, 'มานี', 'มีนา', 'mena', '$2a$08$hAm86bcUwEJOzcG8NKjh.OyOH3Emx2voCQltsZpLBsr1VZlCzP6mi', 2),
	(3, 'มานีมา', 'มีนามา', 'menana', '$2a$08$wqGi9id4.2SpDat4RRo1/e0SkRBWUozhBYkh.54Y8Z3Mj7D2OyHla', 2);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

-- Dumping structure for table shirt-shop.user_address
CREATE TABLE IF NOT EXISTS `user_address` (
  `addr_id` int NOT NULL AUTO_INCREMENT,
  `addr_address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `addr_user_id` int DEFAULT NULL,
  PRIMARY KEY (`addr_id`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table shirt-shop.user_address: ~15 rows (approximately)
/*!40000 ALTER TABLE `user_address` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_address` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
