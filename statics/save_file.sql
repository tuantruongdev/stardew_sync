-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th4 07, 2023 lúc 12:29 PM
-- Phiên bản máy phục vụ: 10.4.27-MariaDB
-- Phiên bản PHP: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `stardew_sync`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `save_file`
--

CREATE TABLE `save_file` (
  `id` int(11) NOT NULL,
  `owner_id` int(11) NOT NULL,
  `world_owner_id` int(11) NOT NULL,
  `description` varchar(255) NOT NULL,
  `path` varchar(255) NOT NULL,
  `favorite` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `image_url` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Đang đổ dữ liệu cho bảng `save_file`
--

INSERT INTO `save_file` (`id`, `owner_id`, `world_owner_id`, `description`, `path`, `favorite`, `created_at`, `updated_at`, `image_url`) VALUES
(1, 1, 9, '', 'statics/saves/2023/4/7/1680857954_jayfarmz_340438575.zip', 0, '2023-04-07 08:59:14', '2023-04-07 08:59:14', ''),
(2, 1, 9, '', 'statics/saves/2023/4/7/1680857955_jayfarmz_340438575.zip', 0, '2023-04-07 08:59:15', '2023-04-07 08:59:15', ''),
(3, 1, 9, '', 'statics/saves/2023/4/7/1680857957_jayfarmz_340438575.zip', 0, '2023-04-07 08:59:17', '2023-04-07 08:59:17', ''),
(9, 2, 15, '', 'statics/saves/2023/4/7/1680860147_jayfarmz_340438575.zip', 0, '2023-04-07 09:35:47', '2023-04-07 09:35:47', ''),
(10, 2, 16, '', 'statics/saves/2023/4/7/1680860498_test_340511369.zip', 0, '2023-04-07 09:41:38', '2023-04-07 09:41:38', '');

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `save_file`
--
ALTER TABLE `save_file`
  ADD PRIMARY KEY (`id`),
  ADD KEY `saveConst` (`owner_id`),
  ADD KEY `worldConst` (`world_owner_id`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `save_file`
--
ALTER TABLE `save_file`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Các ràng buộc cho các bảng đã đổ
--

--
-- Các ràng buộc cho bảng `save_file`
--
ALTER TABLE `save_file`
  ADD CONSTRAINT `saveConst` FOREIGN KEY (`owner_id`) REFERENCES `user` (`id`),
  ADD CONSTRAINT `worldConst` FOREIGN KEY (`world_owner_id`) REFERENCES `save_world` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
