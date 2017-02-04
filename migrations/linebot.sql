-- phpMyAdmin SQL Dump
-- version 4.5.2
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: 2017 年 2 月 05 日 03:25
-- サーバのバージョン： 10.1.19-MariaDB
-- PHP Version: 5.6.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `motivating_menu`
--

-- --------------------------------------------------------

--
-- テーブルの構造 `category`
--

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 NOT NULL,
  `imageUrl` text CHARACTER SET utf8mb4 NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- テーブルのデータのダンプ `category`
--

INSERT INTO `category` (`id`, `name`, `imageUrl`) VALUES
(1, 'ご飯もの', ''),
(2, 'どんぶり・寿司', ''),
(3, 'パン・スナック', ''),
(4, 'めん料理', ''),
(5, '肉料理', ''),
(6, '魚料理', ''),
(7, '卵料理', ''),
(8, '豆腐料理', ''),
(9, '野菜・いも料理', ''),
(10, 'サラダ', '');

-- --------------------------------------------------------

--
-- テーブルの構造 `recipe`
--

DROP TABLE IF EXISTS `recipe`;
CREATE TABLE `recipe` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `category_id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 NOT NULL,
  `imageUrl` text CHARACTER SET utf8mb4 NOT NULL,
  `count1` bigint(20) UNSIGNED NOT NULL DEFAULT '0',
  `count2` bigint(20) UNSIGNED NOT NULL DEFAULT '0',
  `count3` bigint(20) UNSIGNED NOT NULL DEFAULT '0',
  `count4` bigint(20) UNSIGNED NOT NULL DEFAULT '0',
  `count5` bigint(20) UNSIGNED NOT NULL DEFAULT '0',
  `count6` bigint(20) UNSIGNED NOT NULL DEFAULT '0',
  `count7` bigint(20) UNSIGNED NOT NULL DEFAULT '0',
  `count8` bigint(20) UNSIGNED NOT NULL DEFAULT '0',
  `count9` bigint(20) UNSIGNED NOT NULL DEFAULT '0',
  `count10` bigint(20) UNSIGNED NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- テーブルのデータのダンプ `recipe`
--

INSERT INTO `recipe` (`id`, `category_id`, `name`, `imageUrl`, `count1`, `count2`, `count3`, `count4`, `count5`, `count6`, `count7`, `count8`, `count9`, `count10`) VALUES
(1, 1, '001', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 0, 10, 0, 0, 0, 0, 0, 0, 0, 0),
(2, 2, '002', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 20, 0, 30, 0, 0, 10, 20, 0, 0, 0),
(3, 3, '003', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 0, 0, 0, 20, 40, 0, 0, 0, 0, 0),
(4, 4, '004', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 0, 30, 30, 0, 0, 0, 0, 20, 0, 0),
(5, 5, '005', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 0, 20, 0, 0, 0, 0, 0, 0, 0, 0),
(6, 6, '006', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 0, 0, 0, 0, 20, 0, 0, 0, 0, 0),
(7, 7, '007', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(8, 8, '008', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(9, 9, '009', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 10, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(10, 10, '010', 'https://cloud.githubusercontent.com/assets/25542931/22619362/782bd6e4-eb36-11e6-8d4e-e7ef2c228817.jpeg', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `recipe`
--
ALTER TABLE `recipe`
  ADD PRIMARY KEY (`id`),
  ADD KEY `category_id` (`category_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `category`
--
ALTER TABLE `category`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
--
-- AUTO_INCREMENT for table `recipe`
--
ALTER TABLE `recipe`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
--
-- ダンプしたテーブルの制約
--

--
-- テーブルの制約 `recipe`
--
ALTER TABLE `recipe`
  ADD CONSTRAINT `recipe_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
