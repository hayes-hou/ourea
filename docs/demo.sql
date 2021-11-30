-- phpMyAdmin SQL Dump
-- version 4.4.10
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: 2021-08-25 05:47:43
-- 服务器版本： 8.0.25
-- PHP Version: 7.3.24-(to be removed in future macOS)

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `young`
--

-- --------------------------------------------------------

--
-- 表的结构 `upload_img`
--

CREATE TABLE IF NOT EXISTS `upload_img` (
  `id` bigint unsigned NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `path` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=COMPACT;

--
-- 转存表中的数据 `upload_img`
--

INSERT INTO `upload_img` (`id`, `name`, `path`, `created_at`, `updated_at`, `deleted_at`, `url`) VALUES
(1, '1680723ca91aa57d719d5cdbc1d910a1.1680723c (1).jpg', '/Users/houhuiyang/go/src/yuanqi/public/img_187122900.jpg', '2021-08-03 17:51:40', '2021-08-03 17:51:40', '2021-08-09 17:37:54', 'http://127.0.0.1:1818/pic/img_187122900.jpg'),
(2, '1680723ca91aa57d719d5cdbc1d910a1.1680723c (1).jpg', '/Users/houhuiyang/go/src/yuanqi/public/img_411730249.jpg', '2021-08-03 18:43:44', '2021-08-03 18:43:44', NULL, 'http://127.0.0.1:1818/pic/img_411730249.jpg'),
(3, '1680723ca91aa57d719d5cdbc1d910a1.1680723c (1).jpg', '/Users/houhuiyang/go/src/yuanqi/public/img_625838662.jpg', '2021-08-03 18:48:57', '2021-08-03 18:48:57', NULL, 'http://127.0.0.1:1818/pic/img_625838662.jpg');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `upload_img`
--
ALTER TABLE `upload_img`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `upload_img`
--
ALTER TABLE `upload_img`
  MODIFY `id` bigint unsigned NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=4;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
