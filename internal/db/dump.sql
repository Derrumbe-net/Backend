-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Jan 30, 2026 at 07:48 PM
-- Server version: 11.8.3-MariaDB-log
-- PHP Version: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `u573582047_derrumbe`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `admin_id` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `isAuthorized` tinyint(1) DEFAULT 0,
  `is_email_verified` tinyint(1) NOT NULL,
  `verification_token` varchar(64) DEFAULT NULL,
  `token_expires_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`admin_id`, `email`, `password`, `isAuthorized`, `is_email_verified`, `verification_token`, `token_expires_at`) VALUES
(1, 'jose.irizarry24@upr.edu', 'temp_pw', 1, 0, NULL, NULL),
(5, 'alanis.negroni@upr.edu', '$2y$10$J4rBV15NgYeNO8dAxwSuiuqwe5sF2cd99hFhsVQzKfZSiTCONi9zi', 1, 1, NULL, NULL),
(9, 'pedro.matos4@upr.edu', '$2y$10$PDemGvPenC8eHxA0QpfRU.ehVq5o3PFN2QJWyGMiqB9NfW3wjf57.', 1, 0, NULL, NULL),
(10, 'gabriel.colon32@upr.edu', '$2y$10$RHL5H5wS3Qm1qKCdzAnv6OKtZKMeadM5l/Hlq4faXd63YhCIlRRpa', 1, 0, NULL, NULL),
(11, 'slidespr@gmail.com', '$2y$10$I.dlgpgHnR.9hQ84vsMfGuaP4PIoOElZB7UU6kXF8dorpOJUY1Jo6', 1, 1, NULL, NULL),
(26, 'glerysbethserrano@gmail.com', '$2y$12$8wiamJV5/IFvEFxcH5JhUuKX7K80i/GO2qNMDpf24rXhtCiKoa9wC', 1, 0, NULL, NULL),
(35, 'jose.rivera471@upr.edu', '$2y$10$wx7kNezH8uxSgyTBjwF1Ku6vg9UKmzH358RQqKoRSJRA6aJJzI7rq', 0, 0, '9365825190ef589b1df97231253684a0d679bc9eb035ab46d9b19a1b445c6d42', '2025-12-12 09:51:37'),
(40, 'capstone.derrumbes@gmail.com', '$2y$10$00w4LaBd66z97jG1vzEZkec4cyklL77IKxb1tR/QS959Cw1mVosnq', 1, 1, NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `landslide`
--

CREATE TABLE `landslide` (
  `landslide_id` int(11) NOT NULL,
  `admin_id` int(11) NOT NULL,
  `landslide_date` datetime NOT NULL,
  `latitude` decimal(9,6) NOT NULL,
  `longitude` decimal(9,6) NOT NULL,
  `image_url` varchar(512) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `landslide`
--

INSERT INTO `landslide` (`landslide_id`, `admin_id`, `landslide_date`, `latitude`, `longitude`, `image_url`) VALUES
(2, 1, '2020-08-08 00:00:00', 18.116138, -66.293337, '2020-08-08-1'),
(3, 1, '2021-07-30 00:00:00', 18.318784, -67.104539, '2021-07-30-1'),
(4, 1, '2021-08-25 00:00:00', 18.258775, -67.069342, '2021-08-25-1'),
(5, 1, '2021-09-21 00:00:00', 18.291664, -66.838589, '2021-09-21-1'),
(6, 1, '2021-09-22 00:00:00', 18.280264, -66.746939, '2021-09-22-1'),
(7, 1, '2021-09-22 00:00:00', 18.215194, -66.673456, '2021-09-22-2'),
(8, 1, '2021-09-28 00:00:00', 18.286406, -66.480567, '2021-09-28-1'),
(9, 1, '2021-09-29 00:00:00', 18.205897, -66.552931, '2021-09-29-1'),
(10, 1, '2021-10-17 00:00:00', 18.214444, -66.515722, '2021-10-17-1'),
(11, 1, '2021-10-22 00:00:00', 18.143047, -66.458444, '2021-10-22-1'),
(12, 1, '2022-02-03 00:00:00', 18.210139, -66.981492, '2022-02-03-1'),
(13, 1, '2022-02-03 00:00:00', 18.267489, -66.988928, '2022-02-03-2'),
(14, 1, '2022-02-06 00:00:00', 18.319558, -66.139428, '2022-02-06-1'),
(15, 1, '2022-02-06 00:00:00', 18.279336, -66.217081, '2022-02-06-2'),
(16, 1, '2022-02-06 00:00:00', 18.251531, -66.134421, '2022-02-06-3'),
(17, 1, '2022-02-06 00:00:00', 18.299843, -65.947106, '2022-02-06-4'),
(18, 1, '2022-02-06 00:00:00', 18.417699, -66.523026, '2022-02-06-5'),
(19, 1, '2022-02-06 00:00:00', 18.281130, -66.082704, '2022-02-06-6'),
(20, 1, '2022-02-08 00:00:00', 18.350535, -66.016244, '2022-02-08-1'),
(21, 1, '2022-02-08 00:00:00', 18.388111, -66.176500, '2022-02-08-2'),
(22, 1, '2022-03-06 00:00:00', 18.319403, -65.988184, '2022-03-06-1'),
(23, 1, '2022-03-07 00:00:00', 18.266606, -66.138312, '2022-03-07-1'),
(24, 1, '2022-04-11 00:00:00', 18.359916, -67.000534, '2022-04-11-1'),
(25, 1, '2022-05-27 00:00:00', 18.269413, -66.541450, '2022-05-27-1'),
(26, 1, '2022-06-11 00:00:00', 18.282595, -66.673732, '2022-06-11-1'),
(27, 1, '2022-06-12 00:00:00', 18.210538, -66.987291, '2022-06-12-1'),
(28, 1, '2022-06-30 00:00:00', 18.229513, -67.163367, '2022-06-30-1'),
(29, 1, '2022-07-15 00:00:00', 18.247579, -66.312979, '2022-07-15-1'),
(30, 1, '2022-08-20 00:00:00', 18.269177, -67.093162, '2022-08-20-1'),
(31, 1, '2022-08-20 00:00:00', 18.096225, -66.736383, '2022-08-20-2'),
(32, 1, '2022-08-22 00:00:00', 18.312019, -66.473743, '2022-08-22-1'),
(33, 1, '2022-08-23 00:00:00', 18.292940, -67.116486, '2022-08-23-2'),
(34, 1, '2022-08-23 00:00:00', 18.257734, -66.388611, '2022-08-23-3'),
(35, 1, '2022-08-23 00:00:00', 18.097889, -66.641939, '2022-08-23-4'),
(36, 1, '2022-08-23 00:00:00', 18.262567, -66.902765, '2022-08-23-5'),
(37, 1, '2022-08-23 00:00:00', 18.348770, -66.973103, '2022-08-23-6'),
(38, 1, '2022-08-25 00:00:00', 18.285875, -67.049455, '2022-08-25-1'),
(39, 1, '2022-08-25 00:00:00', 18.265430, -66.697294, '2022-08-25-2'),
(40, 1, '2022-08-28 00:00:00', 18.411184, -66.246292, '2022-08-28-1'),
(41, 1, '2022-08-30 00:00:00', 18.276853, -66.479575, '2022-08-30-1'),
(42, 1, '2022-09-01 00:00:00', 18.135478, -67.038694, '2022-09-01-1'),
(43, 1, '2022-09-05 00:00:00', 18.285988, -66.251241, '2022-09-05-1'),
(44, 1, '2022-09-05 00:00:00', 18.267571, -66.249584, '2022-09-05-2'),
(45, 1, '2022-09-06 00:00:00', 18.274048, -66.206785, '2022-09-06-1'),
(46, 1, '2022-09-06 00:00:00', 18.272883, -66.207061, '2022-09-06-2'),
(47, 1, '2022-09-06 00:00:00', 18.267680, -66.206707, '2022-09-06-3'),
(48, 1, '2022-09-07 00:00:00', 18.311156, -66.430472, '2022-09-07-1'),
(49, 1, '2022-09-11 00:00:00', 18.111470, -66.639737, '2022-09-11-1'),
(50, 1, '2022-09-14 00:00:00', 18.277845, -66.194779, '2022-09-14-1'),
(51, 1, '2022-09-14 00:00:00', 18.088829, -66.740959, '2022-09-14-2'),
(52, 1, '2022-09-17 00:00:00', 18.215095, -65.720648, '2022-09-17-1'),
(53, 1, '2022-09-17 00:00:00', 18.368465, -66.107942, '2022-09-17-2'),
(54, 1, '2022-09-18 00:00:00', 18.166744, -66.318049, '2022-09-18-1'),
(55, 1, '2022-09-18 00:00:00', 18.226805, -66.336418, '2022-09-18-2'),
(56, 1, '2022-09-18 00:00:00', 18.383686, -66.613488, '2022-09-18-3'),
(57, 1, '2022-09-18 00:00:00', 18.101910, -66.508050, '2022-09-18-4'),
(58, 1, '2022-09-18 00:00:00', 18.241391, -65.891086, '2022-09-18-5'),
(59, 1, '2022-09-18 00:00:00', 18.200546, -66.070521, '2022-09-18-6'),
(60, 1, '2022-09-18 00:00:00', 18.411030, -66.247929, '2022-09-18-7'),
(61, 1, '2022-09-18 00:00:00', 18.289175, -66.271847, '2022-09-18-8'),
(62, 1, '2022-09-19 00:00:00', 18.344771, -66.384079, '2022-09-19-1'),
(63, 1, '2022-09-19 00:00:00', 18.250304, -66.368983, '2022-09-19-2'),
(64, 1, '2022-09-19 00:00:00', 18.185408, -66.480570, '2022-09-19-3'),
(65, 1, '2022-09-19 00:00:00', 18.195239, -66.473574, '2022-09-19-4'),
(66, 1, '2022-09-19 00:00:00', 18.236631, -66.467228, '2022-09-19-5'),
(67, 1, '2022-09-19 00:00:00', 18.213694, -66.480112, '2022-09-19-6'),
(68, 1, '2022-09-19 00:00:00', 18.211183, -66.478848, '2022-09-19-7'),
(69, 1, '2022-09-19 00:00:00', 18.158015, -66.515061, '2022-09-19-8'),
(70, 1, '2022-09-19 00:00:00', 18.082258, -66.531066, '2022-09-19-9'),
(71, 1, '2022-09-19 00:00:00', 18.107775, -66.237816, '2022-09-19-10'),
(72, 1, '2022-09-19 00:00:00', 18.254075, -65.795067, '2022-09-19-11'),
(73, 1, '2022-09-19 00:00:00', 18.204612, -67.112674, '2022-09-19-12'),
(74, 1, '2022-09-19 00:00:00', 18.148839, -66.940964, '2022-09-19-13'),
(75, 1, '2022-09-19 00:00:00', 18.243424, -66.124676, '2022-09-19-14'),
(76, 1, '2022-09-19 00:00:00', 18.182604, -66.483883, '2022-09-19-15'),
(77, 1, '2022-09-19 00:00:00', 18.216561, -66.591972, '2022-09-19-16'),
(78, 1, '2022-09-19 00:00:00', 18.352452, -67.230511, '2022-09-19-17'),
(79, 1, '2022-09-19 00:00:00', 18.203104, -66.068984, '2022-09-19-18'),
(80, 1, '2022-09-19 00:00:00', 18.178049, -66.462893, '2022-09-19-19'),
(81, 1, '2022-09-19 00:00:00', 18.148921, -66.097571, '2022-09-19-20'),
(82, 1, '2022-09-19 00:00:00', 18.262206, -66.400449, '2022-09-19-21'),
(83, 1, '2022-09-20 00:00:00', 18.207272, -66.396450, '2022-09-20-1'),
(84, 1, '2022-09-20 00:00:00', 18.036516, -65.882385, '2022-09-20-2'),
(85, 1, '2022-09-20 00:00:00', 18.194721, -66.093522, '2022-09-20-3'),
(86, 1, '2022-09-20 00:00:00', 18.231582, -67.091767, '2022-09-20-4'),
(87, 1, '2022-09-20 00:00:00', 18.193076, -67.039395, '2022-09-20-5'),
(88, 1, '2022-09-20 00:00:00', 18.132657, -66.249345, '2022-09-20-6'),
(89, 1, '2022-09-20 00:00:00', 18.089513, -66.159734, '2022-09-20-7'),
(90, 1, '2022-09-20 00:00:00', 18.325942, -65.765109, '2022-09-20-8'),
(91, 1, '2022-09-21 00:00:00', 18.227415, -67.054343, '2022-09-21-1'),
(92, 1, '2022-09-21 00:00:00', 18.264691, -65.794913, '2022-09-21-2'),
(93, 1, '2022-09-21 00:00:00', 18.210886, -66.631186, '2022-09-21-3'),
(94, 1, '2022-09-21 00:00:00', 18.241570, -67.091816, '2022-09-21-4'),
(95, 1, '2022-09-21 00:00:00', 18.265012, -66.714082, '2022-09-21-5'),
(96, 1, '2022-09-21 00:00:00', 18.192826, -67.039131, '2022-09-21-6'),
(97, 1, '2022-09-21 00:00:00', 18.089234, -66.179082, '2022-09-21-7'),
(98, 1, '2022-09-21 00:00:00', 18.085610, -66.177617, '2022-09-21-8'),
(99, 1, '2022-09-21 00:00:00', 18.269685, -66.677927, '2022-09-21-9'),
(100, 1, '2022-09-21 00:00:00', 18.184465, -67.056049, '2022-09-21-10'),
(101, 1, '2022-09-22 00:00:00', 18.249175, -66.859538, '2022-09-22-1'),
(102, 1, '2022-09-22 00:00:00', 18.404520, -66.146599, '2022-09-22-2'),
(103, 1, '2022-09-22 00:00:00', 18.348111, -66.176845, '2022-09-22-3'),
(104, 1, '2022-09-22 00:00:00', 18.186074, -66.763214, '2022-09-22-4'),
(105, 1, '2022-09-22 00:00:00', 18.137337, -66.047567, '2022-09-22-5'),
(106, 1, '2022-09-22 00:00:00', 18.269505, -66.397532, '2022-09-22-6'),
(107, 1, '2022-09-22 00:00:00', 18.112128, -66.231791, '2022-09-22-17'),
(108, 1, '2022-09-23 00:00:00', 18.257553, -65.675813, '2022-09-23-1'),
(109, 1, '2022-09-23 00:00:00', 18.123932, -65.985886, '2022-09-23-2'),
(110, 1, '2022-09-23 00:00:00', 18.351446, -67.225884, '2022-09-23-3'),
(111, 1, '2022-09-23 00:00:00', 18.090699, -65.945647, '2022-09-23-4'),
(112, 1, '2022-09-23 00:00:00', 18.086123, -67.015000, '2022-09-23-5'),
(113, 1, '2022-09-23 00:00:00', 18.191354, -66.592058, '2022-09-23-6'),
(114, 1, '2022-09-23 00:00:00', 18.207687, -66.973800, '2022-09-23-7'),
(115, 1, '2022-09-24 00:00:00', 18.211483, -67.114012, '2022-09-24-1'),
(116, 1, '2022-09-25 00:00:00', 18.245845, -66.865486, '2022-09-25-1'),
(117, 1, '2022-09-25 00:00:00', 18.132317, -66.056669, '2022-09-25-2'),
(118, 1, '2022-09-25 00:00:00', 18.261548, -65.914346, '2022-09-25-3'),
(119, 1, '2022-09-25 00:00:00', 18.208165, -65.916251, '2022-09-25-4'),
(120, 1, '2022-09-25 00:00:00', 18.162333, -67.032022, '2022-09-25-5'),
(121, 1, '2022-09-25 00:00:00', 18.151080, -67.035764, '2022-09-25-6'),
(122, 1, '2022-09-26 00:00:00', 18.129633, -66.054755, '2022-09-26-1'),
(123, 1, '2022-09-26 00:00:00', 18.295447, -66.807421, '2022-09-26-2'),
(124, 1, '2022-09-27 00:00:00', 18.230834, -66.635405, '2022-09-27-1'),
(125, 1, '2022-09-27 00:00:00', 18.209541, -66.645600, '2022-09-27-2'),
(126, 1, '2022-09-27 00:00:00', 18.169181, -66.894906, '2022-09-27-3'),
(127, 1, '2022-09-27 00:00:00', 18.149327, -66.946619, '2022-09-27-4'),
(128, 1, '2022-09-27 00:00:00', 18.328248, -67.153310, '2022-09-27-5'),
(129, 1, '2022-09-27 00:00:00', 18.174009, -66.828201, '2022-09-27-6'),
(130, 1, '2022-09-27 00:00:00', 18.301008, -66.156269, '2022-09-27-7'),
(131, 1, '2022-09-28 00:00:00', 18.053473, -66.100371, '2022-09-28-1'),
(132, 1, '2022-09-28 00:00:00', 18.304725, -66.174983, '2022-09-28-2'),
(133, 1, '2022-09-28 00:00:00', 18.084359, -66.725561, '2022-09-28-3'),
(134, 1, '2022-09-29 00:00:00', 18.180441, -66.896539, '2022-09-29-1'),
(135, 1, '2022-09-30 00:00:00', 18.231333, -66.491726, '2022-09-30-1'),
(136, 1, '2022-10-03 00:00:00', 18.176211, -67.026205, '2022-10-03-1'),
(137, 1, '2022-10-04 00:00:00', 18.047284, -66.722776, '2022-10-04-1'),
(138, 1, '2022-10-04 00:00:00', 18.260147, -66.751339, '2022-10-04-2'),
(139, 1, '2022-10-05 00:00:00', 18.328645, -67.189798, '2022-10-05-1'),
(140, 1, '2022-10-07 00:00:00', 18.078328, -66.038183, '2022-10-07-1'),
(141, 1, '2022-10-07 00:00:00', 18.205166, -67.113526, '2022-10-07-2'),
(142, 1, '2022-10-10 00:00:00', 18.300248, -66.828999, '2022-10-10-1'),
(143, 1, '2022-10-10 00:00:00', 18.379219, -66.108539, '2022-10-10-2'),
(144, 1, '2022-10-11 00:00:00', 18.322000, -67.125780, '2022-10-11-1'),
(145, 1, '2022-10-11 00:00:00', 18.205997, -66.956616, '2022-10-11-2'),
(146, 1, '2022-10-14 00:00:00', 18.180650, -66.486656, '2022-10-14-1'),
(147, 1, '2022-10-15 00:00:00', 18.379379, -67.081958, '2022-10-15-1'),
(148, 1, '2022-10-16 00:00:00', 18.296215, -66.890967, '2022-10-16-1'),
(149, 1, '2022-10-16 00:00:00', 18.208495, -66.724401, '2022-10-16-2'),
(150, 1, '2022-10-17 00:00:00', 18.120354, -66.573503, '2022-10-17-1'),
(151, 1, '2022-10-17 00:00:00', 18.338087, -67.173355, '2022-10-17-2'),
(152, 1, '2022-10-20 00:00:00', 18.343764, -67.104404, '2022-10-20-1'),
(153, 1, '2022-10-21 00:00:00', 18.121103, -65.966038, '2022-10-21-1'),
(154, 1, '2022-10-27 00:00:00', 18.075596, -67.106156, '2022-10-27-1'),
(155, 1, '2022-10-27 00:00:00', 18.228708, -66.971393, '2022-10-27-2'),
(156, 1, '2022-10-27 00:00:00', 18.334012, -67.224688, '2022-10-27-3'),
(157, 1, '2022-10-27 00:00:00', 18.039785, -65.893135, '2022-10-27-4'),
(158, 1, '2022-10-27 00:00:00', 18.210046, -66.976051, '2022-10-27-5'),
(159, 1, '2022-10-27 00:00:00', 18.180402, -67.121461, '2022-10-27-6'),
(160, 1, '2022-10-28 00:00:00', 18.115708, -66.253344, '2022-10-28-1'),
(161, 1, '2022-10-29 00:00:00', 18.038592, -65.888886, '2022-10-29-1'),
(162, 1, '2022-11-02 00:00:00', 18.098153, -66.926602, '2022-11-02-1'),
(163, 1, '2022-11-03 00:00:00', 18.263102, -66.642446, '2022-11-03-1'),
(164, 1, '2022-11-04 00:00:00', 18.422064, -66.682135, '2022-11-04-1'),
(165, 1, '2022-11-05 00:00:00', 18.204195, -66.589817, '2022-11-05-1'),
(166, 1, '2022-11-05 00:00:00', 18.118579, -66.571066, '2022-11-05-2'),
(167, 1, '2022-11-05 00:00:00', 18.255000, -66.394934, '2022-11-05-3'),
(168, 1, '2022-11-05 00:00:00', 18.142136, -67.039699, '2022-11-05-4'),
(169, 1, '2022-11-05 00:00:00', 18.077231, -66.680798, '2022-11-05-5'),
(170, 1, '2022-11-05 00:00:00', 18.077641, -66.689148, '2022-11-05-6'),
(171, 1, '2022-11-05 00:00:00', 18.022730, -66.612649, '2022-11-05-7'),
(172, 1, '2022-11-05 00:00:00', 18.108775, -66.852220, '2022-11-05-8'),
(173, 1, '2022-11-05 00:00:00', 18.122060, -66.557609, '2022-11-05-9'),
(174, 1, '2022-11-05 00:00:00', 18.090033, -66.753912, '2022-11-05-10'),
(175, 1, '2022-11-05 00:00:00', 18.075462, -67.106173, '2022-11-05-11'),
(176, 1, '2022-11-05 00:00:00', 18.109860, -66.152807, '2022-11-05-12'),
(177, 1, '2022-11-06 00:00:00', 18.070745, -66.217254, '2022-11-06-1'),
(178, 1, '2022-11-08 00:00:00', 18.081696, -66.710725, '2022-11-08-1'),
(179, 1, '2022-11-09 00:00:00', 18.223659, -66.225318, '2022-11-09-1'),
(180, 1, '2022-11-09 00:00:00', 18.052328, -65.998619, '2022-11-09-2'),
(181, 1, '2022-11-29 00:00:00', 18.061542, -66.813634, '2022-11-29-1'),
(182, 1, '2023-02-08 00:00:00', 18.279490, -66.256630, '2023-02-08-1'),
(183, 1, '2023-04-08 00:00:00', 18.301920, -65.986964, '2023-04-08-2'),
(184, 1, '2023-04-08 00:00:00', 18.251628, -65.917651, '2023-04-08-3'),
(185, 1, '2023-04-08 00:00:00', 18.034194, -65.833778, '2023-04-08-4'),
(186, 1, '2023-04-08 00:00:00', 18.236749, -65.968466, '2023-04-08-5'),
(187, 1, '2023-04-18 00:00:00', 18.084183, -66.817757, '2023-04-18-1'),
(188, 1, '2023-04-28 00:00:00', 18.245868, -66.818615, '2023-04-28-1'),
(189, 1, '2023-05-01 00:00:00', 18.250513, -66.208242, '2023-05-01-1'),
(190, 1, '2023-05-01 00:00:00', 18.252310, -66.208626, '2023-05-01-2'),
(191, 1, '2023-06-19 00:00:00', 18.298365, -67.218990, '2023-06-19-1'),
(192, 1, '2023-06-28 00:00:00', 18.358447, -67.089270, '2023-06-28-1'),
(193, 1, '2023-07-26 00:00:00', 18.207693, -67.009707, '2023-07-26-1'),
(194, 1, '2023-08-02 00:00:00', 18.243490, -66.278909, '2023-08-02-1'),
(195, 1, '2023-08-26 00:00:00', 18.282967, -65.972050, '2023-08-26-1'),
(196, 1, '2023-08-26 00:00:00', 18.302763, -65.988512, '2023-08-26-2'),
(197, 1, '2023-09-19 00:00:00', 18.107777, -66.237886, '2023-09-19-1'),
(198, 1, '2023-10-04 00:00:00', 18.333518, -66.949689, '2023-10-04-1'),
(199, 1, '2023-10-27 00:00:00', 18.359789, -67.000580, '2023-10-27-1'),
(200, 1, '2023-11-08 00:00:00', 18.339396, -65.692110, '2023-11-08-1'),
(201, 1, '2023-11-27 00:00:00', 18.052933, -67.117467, '2023-11-27-1');

-- --------------------------------------------------------

--
-- Table structure for table `project`
--

CREATE TABLE `project` (
  `project_id` int(11) NOT NULL,
  `admin_id` int(11) NOT NULL,
  `title` varchar(150) NOT NULL,
  `start_year` smallint(5) UNSIGNED DEFAULT NULL,
  `end_year` smallint(5) UNSIGNED DEFAULT NULL,
  `project_status` enum('planned','active','paused','completed','archived') DEFAULT NULL,
  `description` text DEFAULT NULL,
  `image_url` varchar(512) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `project`
--

INSERT INTO `project` (`project_id`, `admin_id`, `title`, `start_year`, `end_year`, `project_status`, `description`, `image_url`) VALUES
(1, 1, 'Evaluation of the Soil Mass Movement Risk Rating in Puerto Rico using the SLIDES-PR Hurricane María Slope Failure Inventory.', 2019, 2021, 'completed', '($50,098) Acuerdo de cooperación entre el USDA y el NRCS NR20F3520001C001: \"Evaluación de la calificación de riesgo de movimiento de masas de suelo en Puerto Rico utilizando el inventario de fallas de taludes del huracán María de SLIDES-PR\". Este proyecto condujo a una comprensión mucho mejor del papel de la composición y las características del suelo en los deslizamientos de tierra superficiales en Puerto Rico.', NULL),
(2, 1, 'Collaborative Research: Quantifying controls on weathering of volcanic arc rocks.', 2020, 2022, 'completed', '($117,469) Premio NSF de Ciencias de la Tierra #2011358: \"Investigación colaborativa: cuantificación de los controles sobre la erosión de las rocas de arco volcánico\". El objetivo de este proyecto era medir cómo la erosión química de las rocas volcánicas e ígneas ricas en hierro depende del suministro de minerales frescos por erosión física. Esta relación es fundamental para comprender el papel de los procesos tectónicos en el control del ciclo global del carbono y el clima a través del tiempo geológico. Colaborador: Universidad de Purdue', 'proyecto_pasado_2.webp'),
(3, 1, 'Track I Center Catalyst: Collaborative Center for Landslides and Ground Failure Geohazards.', 2022, 2024, 'completed', '($89,700) Premio NSF de Ciencias de la Tierra #2224973: \"Track I Center Catalyst: Collaborative Center for Landslides and Ground Failure Geohazards\". Este proyecto se centró en la investigación relacionada con las causas fundamentales y los mecanismos desencadenantes de los deslizamientos de tierra, así como en el desarrollo de una comprensión de los peligros que se generan a partir de los derrumbes del terreno. Este proyecto de Track I utilizó a Puerto Rico como un laboratorio viviente para estudiar los deslizamientos de tierra y su impacto en la comunidad. Colaboradores: Georgia Tech, Universidad de Colorado.', 'proyecto_pasado_3.webp'),
(4, 1, 'Collaborative Research: RAPID: The fate of landslide-derived sediment following tropical cyclones: a case study of Hurricane Fiona in Puerto Rico.', 2022, 2023, 'completed', '($21,750) Premio NSF de Ciencias de la Tierra #2301379: \"Investigación colaborativa: RAPID: El destino de los sedimentos derivados de deslizamientos de tierra después de ciclones tropicales: un estudio de caso del huracán Fiona en Puerto Rico\". Este proyecto respaldó la recopilación de datos perecederos y sensibles al tiempo de las cuencas fluviales montañosas de Puerto Rico fuertemente afectadas por el vertido masivo después del huracán Fiona. Colaboradores: Georgia Tech', 'proyecto_pasado_4.webp'),
(6, 1, 'LandslideReady community engagement program in Puerto Rico 2024-2026.', 2024, 2026, 'active', '($473,895) Acuerdo de cooperación del USGS n.º G24AC00484: \"Programa de participación comunitaria LandslideReady en Puerto Rico 2024-2026\". El propósito de este acuerdo es apoyar la investigación y la recopilación de datos relacionados con la implementación y el análisis de un programa de certificación municipal estructurado LandslideReady en Puerto Rico (PR). LandslideReady está modelado a partir de los exitosos esfuerzos StormReady y TsunamiReady del NWS. En los últimos años, la Oficina de Mitigación de Peligros de Deslizamientos de Tierra de PR (PRLHMO) ha codiseñado una versión piloto de LandslideReady con aportes de científicos físicos, científicos sociales, el gobierno federal/estatal/local, la industria privada, líderes comunitarios, grupos sin fines de lucro y otros socios ciudadanos. Las oficinas de gestión de emergencias municipales son los grupos objetivo que se certificarán como LandslideReady.', 'proyecto_actual_1.webp'),
(7, 1, 'Climate Adaptation Partnerships: Caribbean Climate Adaptation Network: Building equitable adaptive capacities of the USVI and Puerto Rico.', 2022, 2027, 'active', '($462,505) NOAA Climate Program Office Award #NA22OAR4310545: \"Asociaciones de adaptación climática: Red de adaptación climática del Caribe: creación de capacidades de adaptación equitativas de las Islas Vírgenes de los Estados Unidos y Puerto Rico\". Este esfuerzo busca mejorar y expandir las asociaciones mediante el desarrollo y la convocatoria de partes interesadas en Puerto Rico y las Islas Vírgenes de los Estados Unidos. La red de conocimiento y acción propuesta está diseñada para ayudar a desarrollar capacidades de adaptación para futuros extremos climáticos, planificar respuestas a peligros climáticos en cascada y crisis de gobernanza. Colaboradores: UPR Ciencias Médicas, Universidad de las Islas Vírgenes, Universidad de Texas, Universidad de Nueva York, Instituto Politécnico de Worcester', 'proyecto_actual_2.webp'),
(8, 1, 'Puerto Rico Landslide Hazard Reduction Project 2023-2025.', 2023, 2026, 'active', '($499,956) Acuerdo de cooperación del USGS n.° G23AC00479: \"Proyecto de reducción del riesgo de deslizamientos de tierra en Puerto Rico 2023-2025\". Esta adjudicación amplía un acuerdo vigente entre el Programa de riesgo de deslizamientos de tierra del USGS y la Universidad de Puerto Rico en Mayagüez para establecer y operar una red de monitoreo hidrológico del suelo casi en tiempo real. El nuevo acuerdo ayudará a expandir la red de monitoreo hidrológico actual (de 15 estaciones actuales a al menos 20 estaciones) y brindará un medio para mantener la red funcional, complementando así una nueva \"Oficina de mitigación del riesgo de deslizamientos de tierra en Puerto Rico\" en el campus de la UPRM. El objetivo del esfuerzo de investigación y recopilación de datos es desarrollar métricas de pronóstico de deslizamientos de tierra en todo el territorio de la isla para usarlas en un sistema operativo.', 'proyecto_actual_3.webp'),
(10, 1, 'Landslide Hazard Science and Risk Communication in Puerto Rico.', 2024, 2026, 'active', '($149,998) Premio #2024-00188 del Fideicomiso de Ciencia, Tecnología e Investigación de Puerto Rico: \"Ciencia y comunicación de riesgos de deslizamientos de tierra en Puerto Rico\". Este proyecto apoyará el desarrollo e implementación de un sistema operativo y en tiempo real de pronóstico de deslizamientos de tierra en Puerto Rico.', 'proyecto_actual_4.webp'),
(12, 1, 'Collaborative Research: Testing Critical Zone Controls on Mountain-Scale Relief in a Tropical Climate.', 2022, 2025, 'active', '($284,503) Premio NSF de Ciencias de la Tierra #2139895: \"Investigación colaborativa: prueba de controles de zonas críticas en relieve a escala montañosa en un clima tropical\". Este proyecto examina cómo las diferencias en los procesos de la zona crítica influyen en la topografía a través de un experimento comparativo de dos unidades de lecho rocoso diferentes en la isla tropical de Puerto Rico. Colaborador: Universidad Estatal de Colorado', 'proyecto_actual_5.webp');

-- --------------------------------------------------------

--
-- Table structure for table `publication`
--

CREATE TABLE `publication` (
  `publication_id` int(11) NOT NULL,
  `admin_id` int(11) NOT NULL,
  `title` varchar(150) NOT NULL,
  `publication_url` varchar(512) NOT NULL,
  `image_url` varchar(512) DEFAULT NULL,
  `description` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `publication`
--

INSERT INTO `publication` (`publication_id`, `admin_id`, `title`, `publication_url`, `image_url`, `description`) VALUES
(3, 1, 'Tracking a limestone bedrock landslide on an urbanized hillslope in Guanica, Puerto Rico', 'https://drive.google.com/file/d/14cj37pysCrrXOQkm8ZWmBoEodR6vGTyR/view?usp=sharing', 'pub1.webp', 'Rodríguez-Feliciano, César A.; Hughes, K. Stephen; Vélez-Santiago, Freddie; Department of Geology, University of Puerto Rico, Mayagüez, PR; Guánica, PR. (2025). Tracking a limestone bedrock landslide on an urbanized hillslope in Guanica, Puerto Rico.'),
(4, 1, 'Chemical Weathering and Physical Erosion Fluxes From Serpentinite in Puerto Rico', 'https://agupubs.onlinelibrary.wiley.com/doi/full/10.1029/2024JF007776', 'pub2.webp', 'Angus K. Moore, Kimberly Méndez Méndez, K. Stephen Hughes, Darryl E. Granger. (2025). Chemical Weathering and Physical Erosion Fluxes From Serpentinite in Puerto Rico,https://agupubs.onlinelibrary.wiley.com/doi/full/10.1029/2024JF007776'),
(5, 1, 'Dynamic Landslide Susceptibility for Extreme Rainfall Events Using an Optimized Convolutional Neural Network Approach', 'https://link.springer.com/content/pdf/10.1007/s11069-025-07396-9.pdf', 'pub3.webp', 'Mejia-Manrique, S.A., Ramos-Scharrón, C.E., Hughes, K.S., Gonzalez-Cruz, J.E., and Khanbilvardi, R.M., 2025, Dynamic Landslide Susceptibility for Extreme Rainfall Events Using an Optimized Convolutional Neural Network Approach'),
(6, 1, 'Neotectonic Mapping of Puerto Rico', 'https://seismica.library.mcgill.ca/article/view/1102', 'pub4.webp', 'Jessica A. Thompson Jobe, Richard Briggs, K. Stephen Hughes, James Joyce, Ryan Gold, Shannon Mahan, Harrison Gray, Laura Strickland, U.S. Geological Survey, Geologic Hazard Science Center, Golden, CO, USA, Department of Geology, University of Puerto Rico Mayagüez, Mayagüez, Puerto Rico, USA, U.S. Geological Survey, Geoscience and Environmental Change Science Center, Denver, CO, USA. (2024). Neotectonic Mapping of Puerto Rico, https://seismica.library.mcgill.ca/article/view/1102'),
(7, 1, 'Volcanic arc weathering rates in the humid tropics controlled by the interplay between physical erosion and precipitation', 'https://agupubs.onlinelibrary.wiley.com/doi/full/10.1029/2023AV001066', 'pub5.webp', 'Moore, A. K., Méndez Méndez, K., Hughes, K. S., & Granger, D. E. (2024). Volcanic arc weathering rates in the humid tropics controlled by the interplay between physical erosion and precipitation. AGU Advances, 5, e2023AV001066. https://doi. org/10.1029/2023AV001066'),
(8, 1, 'Pseudo-Three-Dimensional Back-Analysis Of Rainfall-Induced Landslides In Utuado, Puerto Rico', 'https://mst.elsevierpure.com/ws/portalfiles/portal/41586365/Pseudo-Three-Dimensional%20Back-Analysis%20Of%20Rainfall-Induced%20Landsl.pdf', 'pub6.webp', 'Mirna Kassem; Weibing Gong; Dimitrios Zekkos; Marin Clark; et. al;  Missouri University of Science and Technology. (2024). Pseudo-Three-Dimensional Back-Analysis Of Rainfall-Induced Landslides In Utuado, Puerto Rico, https://mst.elsevierpure.com/ws/portalfiles/portal/41586365/Pseudo-Three-Dimensional%20Back-Analysis%20Of%20Rainfall-Induced%20Landsl.pdf'),
(9, 1, 'Assessing Social Vulnerability to Landslides in Rural Puerto Rico', 'https://hazards.colorado.edu/public-health-disaster-research/steep-risks', 'pub7.webp', 'West, J., Rodríguez-Cruz, L. A., & Hughes, K. S. (2023). Steep Risks: Assessing Social Vulnerability to Landslides in Rural Puerto Rico (Natural Hazards Center Public Health Disaster Research Report Series, Report 34). Natural Hazards Center, University of Colorado Boulder. https://hazards.colorado.edu/public-health-disaster-research/steep-risks'),
(10, 1, 'Climato-tectonic evolution of siliciclastic sandstones on Puerto Rico: from lithic arenites to quartz-arenitic sands in an oceanic island-arc setting ', 'https://pubs.geoscienceworld.org/sepm/jsedres/article/93/11/857/625560/Climato-tectonic-evolution-of-siliciclastic', 'pub8.webp', 'David K. Larue; Kimberly Mendez Mendez; José L. Corchado Albelo; Lauryn N. Martinez; K. Stephen Hughes; Thomas Hudgins; Hernan Santos; Alan L. Smith; Chris Osterberg. (2023). Climato-tectonic evolution of siliciclastic sandstones on Puerto Rico: from lithic arenites to quartz-arenitic sands in an oceanic island-arc setting, https://pubs.geoscienceworld.org/sepm/jsedres/article/93/11/857/625560/Climato-tectonic-evolution-of-siliciclastic'),
(11, 1, 'Geotechnical Impacts of Hurricane Fiona in Puerto Rico', 'https://www.geerassociation.org/components/com_geer_reports/geerfiles/GEER_HurricaneFiona_report.pdf', 'pub9.webp', 'Morales, A.R., Hughes, S.K., Lang, K.D., Rivera-Hernandez, F.M., Vargas Vargas, P.V., Lozano, J.M., Karantanellis, E.P., Kassem, M.L., Gomberg, D.A., Plescher, R.T., Irizarry, E.J., Vicens, E.G., Figueroa, T.L., Friedman, C.M., Cunillera, K.N., Ruiz, A.L., and Ortega, V.J., University of Puerto Rico at Mayagüez, Georgia Institute of Technology, University of Michigan, University of California at Berkeley, Mar. 2023, Geotechnical Impacts of Hurricane Fiona in Puerto Rico, DOI:10.18118/G6Z38B'),
(12, 1, 'Principles for collaborative risk communication: Reducing landslide losses in Puerto Rico, Journal of Emergency Management', 'https://hazards.colorado.edu/uploads/documents/principles-of-collaborative-risk-communication-puertorico-west-et-al-2021.pdf', 'pub10.webp', 'West, J., Davis, L., Lugo Bendezú, R., Álvarez Gandía, Y.D., Hughes, K.S., Godt, J., and Peek, L., 2021, Principles for collaborative risk communication: Reducing landslide losses in Puerto Rico, Journal of Emergency Management, v. 19, no. 2, p. (pdf) https://hazards.colorado.edu/uploads/documents/principles-of-collaborative-risk-communication-puertorico-west-et-al-2021.pdf'),
(13, 1, 'WIDESPREAD SHALLOW MASS WASTING DURING HURRICANE MARIA: LONG-TERM SIGNIFICANCE OF SEDIMENTATION IN THE TROPICS', 'https://www.scipedia.com/public/Irizarry-Brugman_et_al_2021a', 'pub11.webp', 'Edwin O. Irizarry-Brugman, Desiree Bayouth-García, Kenneth S. Hughes. (2021). WIDESPREAD SHALLOW MASS WASTING DURING HURRICANE MARIA: LONG-TERM SIGNIFICANCE OF SEDIMENTATION IN THE TROPICS'),
(14, 1, 'Geotechnical Reconnaissance of the January 7, 2020 M6.4 Southwest Puerto Rico Earthquake and Associated Seismic Sequence, Geotechnical Extreme Events ', 'https://www.geerassociation.org/components/com_geer_reports/geerfiles/GEER_PuertoRico_Report.pdf', 'pub12.webp', 'Morales-Vélez, A.C., Bernal, J., Hughes, K.S., Pando, M., Pérez, J., and Rodríguez, L.A., 2020, Geotechnical Reconnaissance of the January 7, 2020 M6.4 Southwest Puerto Rico Earthquake and Associated Seismic Sequence, Geotechnical Extreme Events Reconnaissance Report No. 066, 55p. (link) http://www.geerassociation.org/administrator/components/com_geer_reports/geerfiles/GEER_PuertoRico_Report.pdf'),
(15, 1, 'Landslide Science in Puerto Rico: Past, Present, and Future; Revista Internacional de Desastres Naturales, Accidentes e Infraestructura Civil', 'https://drive.google.com/file/d/1NelGaHVMEOPSRbBWbDx4KqZiEKVz-vBe/view?usp=sharing', 'pub13.webp', 'Hughes, K.S., and Morales Vélez, A.C., 2020, Landslide Science in Puerto Rico: Past, Present, and Future; Revista Internacional de Desastres Naturales, Accidentes e Infraestructura Civil, v. 19-20, no. 1, p. 175-187. (pdf) https://drive.google.com/file/d/1NelGaHVMEOPSRbBWbDx4KqZiEKVz-vBe/view?usp=sharing'),
(16, 1, 'Map depicting susceptibility to landslides triggered by intense rainfall, Puerto Rico', 'https://doi.org/10.3133/ofr20201022', 'pub14.webp', 'Hughes, K.S., and Schulz, W.H., 2020, Map depicting susceptibility to landslides triggered by intense rainfall, Puerto Rico: U.S. Geological Survey Open-File Report 2020-1022, 91 p., 1 plate, scale 1:150,000, https://doi.org/10.3133/ofr20201022'),
(17, 1, 'Landslides triggered by Hurricane Maria : Assessment of an extreme event in Puerto Rico', 'https://www.geosociety.org/gsatoday/science/G383A/article.htm', 'pub15.webp', 'Bessette-Kirton, E.K., Cerovski-Darriau, C., Schulz, W.H., Coe, J.A., Kean, J.W., Godt, J.W., Thomas, M.A., and Hughes, K.S., 2019, Landslides triggered by Hurricane Maria : Assessment of an extreme event in Puerto Rico, GSA Today v. 29, no. 6, p. 4-10. (https://www.geosociety.org/gsatoday/science/G383A/article.htm)'),
(18, 1, 'Map of slope-failure locations in Puerto Rico after Hurricane María', 'https://doi.org/10.5066/P9BVMD74', 'pub16.webp', 'Hughes, K.S., Bayouth García, D., Martínez Milian, G.O., Schulz, W.H., and Baum, R.L., 2019, Map of slope-failure locations in Puerto Rico after Hurricane María: U.S. Geological Survey data release, https://doi.org/10.5066/P9BVMD74.'),
(19, 1, 'Multi-Decadal Earth Dam Deformation Monitoring using Airborne LiDAR and Structure from Motion at Lago Guajataca, Puerto Rico', 'https://drive.google.com/file/d/1FAWv0tWDF84OkZpxUiUN8z8ZVsjIN7_2/view?usp=sharing', 'pub17.webp', 'Villareal Arango, A.F., Hughes, K.S., and Morales-Vélez, A.C., 2019, Multi-Decadal Earth Dam Deformation Monitoring using Airborne LiDAR and Structure from Motion at Lago Guajataca, Puerto Rico : Geocongress, American Society of Civil Engineers, (link). https://drive.google.com/file/d/1FAWv0tWDF84OkZpxUiUN8z8ZVsjIN7_2/view?usp=sharing'),
(20, 1, 'Geotechnical Impacts of Hurricane Maria in Puerto Rico : Geotechnical Extreme Events Reconnaissance', 'https://geerassociation.org/components/com_geer_reports/geerfiles/180629_GEER_PR_Report_No_GEER-057.pdf', 'pub18.webp', 'Silva-Tulla, F., Pando, M.A., Soto, A.E., Morales, A.C., Pradel, D., Inci, G., Sasanakul, I., Bernal, J.R., Kayen, R., Hughes, K.S., Adams, T., and Park, Y., 2018, Geotechnical Impacts of Hurricane Maria in Puerto Rico : Geotechnical Extreme Events Reconnaissance Report No. 057, 234 p. (link) http://www.geerassociation.org/administrator/components/com_geer_reports/geerfiles/180629_GEER_PR_Report_No_GEER-057.pdf'),
(21, 1, 'Comprehensive Hurricane María Mass Wasting Inventory and Improved Frequency Ratio Landslide Hazard Mapping: Status Update From the University of Puert', 'https://drive.google.com/file/d/1wf6xLHbGjE38L3hGB8s2_R9jEWcXsZBa/view?usp=drivesdk', 'pub19.webp', 'Morales-Vélez, A.C., and Hughes, K.S., 2018, Comprehensive Hurricane María Mass Wasting Inventory and Improved Frequency Ratio Landslide Hazard Mapping: Status Update From the University of Puerto Rico at Mayagüez : Dimension, Colegio de Inginieros y Agrimensores de Puerto Rico, v. 1, p. 23-26. (link) https://drive.google.com/file/d/1wf6xLHbGjE38L3hGB8s2_R9jEWcXsZBa/view?usp=drivesdk');

-- --------------------------------------------------------

--
-- Table structure for table `report`
--

CREATE TABLE `report` (
  `report_id` int(11) NOT NULL,
  `landslide_id` int(11) DEFAULT NULL,
  `reported_at` datetime DEFAULT NULL,
  `description` text DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  `image_url` varchar(512) DEFAULT NULL,
  `latitude` decimal(9,6) DEFAULT NULL,
  `longitude` decimal(9,6) DEFAULT NULL,
  `reporter_name` varchar(100) DEFAULT NULL,
  `reporter_phone` varchar(30) DEFAULT NULL,
  `reporter_email` varchar(255) DEFAULT NULL,
  `physical_address` varchar(512) DEFAULT NULL,
  `is_validated` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `report`
--

INSERT INTO `report` (`report_id`, `landslide_id`, `reported_at`, `description`, `city`, `image_url`, `latitude`, `longitude`, `reporter_name`, `reporter_phone`, `reporter_email`, `physical_address`, `is_validated`) VALUES
(239, NULL, '2023-11-29 00:00:00', 'https://www.facebook.com/photo/?fbid=662570289399254&set=pcb.662570396065910', 'Km 2.6 PR-162, Aibonito', '2023-11-29_1', NULL, NULL, NULL, NULL, NULL, 'Contratistas de la ACT trabajando en la reparación del deslizamiento', 0),
(240, NULL, '2023-12-01 00:00:00', 'https://www.facebook.com/photo/?fbid=923369042479358&set=a.562956738520592', 'Carr PR-900, Km 6.5, Yabucoa', '2023-12-01_1', NULL, NULL, NULL, NULL, NULL, 'Brigadas del Departamento de Transportación y Obras Públicas trabajan en la construcción de un muro de contención en gaviones por deslizamiento ocurrido', 0),
(241, NULL, '2024-02-06 00:00:00', 'https://www.facebook.com/SlidesPR\nCarr 834, barrio Río hacia el sector El Laberinto de Hato Nuevo, en Guaynabo', 'Guaynabo', '2024-02-06_241', NULL, NULL, '', '', '', 'Las autoridades reportan el cierre al tránsito por un deslizamiento', 0),
(242, NULL, '2024-02-07 00:00:00', NULL, 'Bo. San Antonio en Caguas', '2024-02-07_1', NULL, NULL, NULL, NULL, NULL, 'Notificado por Estefanía whatsapp', 0),
(243, NULL, '2024-02-10 00:00:00', NULL, 'Cayey', '2024-02-10_1', 18.090720, -66.181375, NULL, NULL, NULL, '18.0907199, -66.1813751', 0),
(244, NULL, '2024-03-25 00:00:00', 'https://m.facebook.com/story.php?story_fbid=397032166419424&id=100083379304955&sfnsn=wa&mibextid=VhDh1V', 'Ciales', '2024-03-25_1', NULL, NULL, NULL, NULL, NULL, 'Derrumbe en la carretera La Aldea PR-649 en Ciales', 0),
(245, NULL, '2024-04-08 00:00:00', 'Wapa.tv', 'Ciales', '2024-04-08_1', NULL, NULL, NULL, NULL, NULL, 'desprendimiento de terreno esta afectando la carretera PR-149 del sector Casa blanca del barrio Cialitos, en Ciales. Personal De Manejo de Emergencias Estatal en la zona Investigando.', 0),
(246, NULL, '2024-04-09 00:00:00', 'https://www.facebook.com/SlidesPR', 'Naranjito', '', NULL, NULL, NULL, NULL, NULL, 'CARRETERA CERRADA | Debido a un deslizamiento de terreno la Carr. 164 KM 2.7 (7 Curvas) permanecerá CERRADA hasta que personal de Obras Públicas Municipal, que ya se encuentra en el lugar, remuevan el mismo.', 0),
(247, NULL, '2024-04-12 00:00:00', 'https://www.facebook.com/SlidesPR', 'Barranquitas', '', NULL, NULL, NULL, NULL, NULL, 'Sectores La Torres y Barrancas', 0),
(248, NULL, '2024-04-18 00:00:00', 'https://www.facebook.com/SlidesPR', 'Orocovis', '', NULL, NULL, NULL, NULL, NULL, 'Bo. El Gato, Carretera 155 a las 6:00pm', 0),
(249, NULL, '2024-04-19 00:00:00', 'https://www.facebook.com/SlidesPR', 'Naranjito', '', NULL, NULL, NULL, NULL, NULL, 'Camino Corozal a Naranjito', 0),
(250, NULL, '2024-04-20 00:00:00', 'https://www.facebook.com/SlidesPR', 'Maunabo', '', NULL, NULL, NULL, NULL, NULL, 'Carr #759 En dirección del Barrio Lias hacia Matuyas', 0),
(251, NULL, '2024-04-22 00:00:00', 'https://www.facebook.com/SlidesPR', 'Aguas Buenas', '', NULL, NULL, NULL, NULL, NULL, 'barrio Jagüeyes entrando antiguo Pito Radiadores', 0),
(252, NULL, '2024-04-24 00:00:00', 'https://www.facebook.com/SlidesPR', 'Cidra', '', NULL, NULL, NULL, NULL, NULL, 'Carr 172 Km 9.3', 0),
(253, NULL, '2024-04-25 00:00:00', 'https://www.facebook.com/SlidesPR', 'Utuado', '', NULL, NULL, NULL, NULL, NULL, NULL, 0),
(254, NULL, '2024-04-26 00:00:00', 'https://www.facebook.com/SlidesPR', 'Yauco', '', NULL, NULL, NULL, NULL, NULL, NULL, 0),
(255, NULL, '2024-05-03 00:00:00', 'https://www.facebook.com/share/DFjtKC59uMtiTqMq/?mibextid=xfxF2i', 'Orocovis', '2024-05-03_1', NULL, NULL, NULL, NULL, NULL, 'Sector Cometa, Bermejales, ruta Panorámica Carr PR-143', 0),
(256, NULL, '2024-05-06 00:00:00', 'https://www.facebook.com/SlidesPR', 'Orocovis', '', NULL, NULL, NULL, NULL, NULL, 'Lunes 5:00pm, Carr 157', 0),
(257, NULL, '2024-05-06 00:00:00', 'https://www.facebook.com/SlidesPR', 'Coamo', '', NULL, NULL, NULL, NULL, NULL, 'Sector Santa Ana', 0),
(258, NULL, '2024-05-06 00:00:00', 'https://www.facebook.com/SlidesPR', 'Cidra', '', NULL, NULL, NULL, NULL, NULL, 'Carr 171, km 2.3 bajando hacia Bo. Rincon cerca de la comunidad Los Torres', 0),
(259, NULL, '2024-05-07 00:00:00', 'https://www.facebook.com/SlidesPR', 'Morovis', '', NULL, NULL, NULL, NULL, NULL, 'Carr 567', 0),
(260, NULL, '2024-05-07 00:00:00', 'https://www.facebook.com/SlidesPR', 'Ciales', '', NULL, NULL, NULL, NULL, NULL, 'Carr 149 de Ciales, sector Casa Blanca', 0),
(261, NULL, '2024-05-07 00:00:00', 'https://www.facebook.com/SlidesPR', 'Lares', '', NULL, NULL, NULL, NULL, NULL, 'Bo. Bartolo', 0),
(262, NULL, '2024-05-08 00:00:00', 'https://www.facebook.com/SlidesPR', 'San Sebastian', '', NULL, NULL, NULL, NULL, NULL, 'Salto Collazo', 0),
(263, NULL, '2024-05-08 00:00:00', 'https://www.facebook.com/share/qH2McNizS3GZxtpR/?mibextid=xfxF2i', 'Ponce', '2024-05-08_2', NULL, NULL, '', '', '', '', 0),
(264, NULL, '2024-05-09 00:00:00', 'https://www.facebook.com/SlidesPR', 'Yauco', '', NULL, NULL, NULL, NULL, NULL, 'Carr 128, BO Rubias ', 0),
(265, NULL, '2024-05-09 00:00:00', 'Via whatsapp', 'Yauco', '2024-05-09_2', NULL, NULL, NULL, NULL, NULL, 'PR 128 km 24.9', 0),
(266, NULL, '2024-05-09 00:00:00', 'https://www.facebook.com/photo/?fbid=868389611984882&set=pcb.868389711984872&__cft__[0]=AZWoD2GhZ5PvOTrjG9zNW2mCIq0nZzKesjlfhl9EPGGMgSsSUYUXJ8HK7y-jhUrlccqYNExRPJPnrnEPfSyq4WrlqCFWnxs61-avXKfsH4rv-8-rLa4mD-s9-vx-UBNT5tS3c2YUqn9oVQ9TqCDxIKKkVWLp25Ll3vrq-4_2xWGpO1K-uWgTwQOQYURo1V36DHVZ21Zrp0_iAd6OACLLGbLT&__tn__=*bH-R', 'Barranquitas ', '2024-05-09_3', NULL, NULL, NULL, NULL, NULL, 'Carretera 152 frente al restaurante El Genesis Bo. Quebradillas ', 0),
(267, NULL, '2024-05-10 00:00:00', 'https://www.facebook.com/share/dFxdiCDuby2f4QbC/?mibextid=xfxF2i', 'Vega Baja', '2024-05-10_1', NULL, NULL, NULL, NULL, NULL, 'Rio Cibuco fuera de cauce y derrumbes, Carr Candelaria Palmita', 0),
(268, NULL, '2024-05-11 00:00:00', 'https://www.facebook.com/share/Egkwh7UNCdV2Ghwg/?mibextid=xfxF2i', 'San sebastian', '2024-05-11_1', NULL, NULL, NULL, NULL, NULL, 'Carr 119 km 43.4 cerca del centro comunal en el barrio guacho de San Sebastián ', 0),
(269, NULL, '2024-05-17 00:00:00', 'https://www.facebook.com/photo/?fbid=786034777004560&set=a.424178616523513&__cft__[0]=AZUIk1Zi1QeWJsSg2T2t5Ojh2mQNSewp33xdfrlpoSm2jneZ3dwO5RI_T9XSRtbq6ncljkgXP8DmiLe66RzBgmQ1pd8gCuMGSISjr4nzRYedGyhQN1IyI1Pixd-Dr_cVzOJ8BxFf2_INayENXJmGDtn7IjUpoHh2oRcwZQXTIcOou1rPJxVps9yysfuGmXZD28C8a4OslDkUsXX1ibvHzhwn&__tn__=EH-R', 'Lares ', '2024-05-17_1', NULL, NULL, NULL, NULL, NULL, 'Pezuela ', 0),
(270, NULL, '2024-05-18 00:00:00', 'https://www.facebook.com/photo/?fbid=760626546255422&set=pcb.760626836255393&__cft__[0]=AZVlBtOCL66zSNJksOyGiTij9byRhaR11TWDtQtQ9vT-tVrwGP9QZglCNLDVMIjxtCtGpyp47ZDoRlr0yFu9zHpdMrAwrxlrpZWx2T-OeuTBOoaFQQiR6fjEOzQW9xYcpmL2obps-GLOEfaTOUfHp5WYRgF5ZUyaycls7rG8epSd8YJ74hacKBGOMlVDqNX3nvrbFXVOm633JtjdqTY7TboB&__tn__=*bH-R', 'Lares', '2024-05-18_1', NULL, NULL, NULL, NULL, NULL, NULL, 0),
(271, NULL, '2024-05-21 00:00:00', 'https://www.facebook.com/photo/?fbid=846931340785170&set=pcb.846929267452044', 'Guánica', '2024-05-21_1', NULL, NULL, NULL, NULL, NULL, 'Alturas de Bélgica, Calle Las Rosas', 0),
(272, NULL, '2024-06-02 00:00:00', 'https://www.facebook.com/photo/?fbid=763036719323154&set=a.302259388734225', 'Yabucoa ', '2024-06-02_1', NULL, NULL, NULL, NULL, NULL, 'Carretera 182, km 11.2, cerca del negocio La Gallera en el barrio Guayabota', 0),
(273, NULL, '2024-06-02 00:00:00', 'https://www.facebook.com/photo/?fbid=759752516329412&set=a.297521622552506&__cft__[0]=AZU8yk5_v-CI_Fcg2qkdBmidD0dKysEya1tuKe2eghjnjEbUsOV8AeFSJj04zR52XkDSN9kVB7dw8IwM4HH5lp5Gwwe0hv4Z0pS0L1npO5WGVhkDozQ3JV-FtZ5wbsb54KuMl8QJ1nRRq4ZGIbKUyDy7hjqq9Fibk6nB58q4TERUdszSfWTi0-_v2DQWA7KtAuCptCEhXy5ex9qvZ-Z4snFaTsNa49u76CJIafMY8BFAB1Yw65L9BSGDiLNb_gn-_6E&__tn__=EH-y-R', 'Naguabo', '2024-06-02_2', NULL, NULL, NULL, NULL, NULL, 'Barrio Higüerillo ', 0),
(274, NULL, '2024-06-02 00:00:00', 'https://www.elvocero.com/el-tiempo/reportan-deslizamientos-de-terreno-en-varios-municipios-de-la-isla/article_4a1bccba-211a-11ef-82bd-f368349c3c1d.html', 'Aguas Buenas ', '2024-06-02_3', NULL, NULL, NULL, NULL, NULL, 'Carretera 173 Barrio Sumidero', 0),
(275, NULL, '2024-06-02 00:00:00', 'https://www.elvocero.com/el-tiempo/reportan-deslizamientos-de-terreno-en-varios-municipios-de-la-isla/article_4a1bccba-211a-11ef-82bd-f368349c3c1d.html', 'Gurabo', '2024-06-02_4', NULL, NULL, NULL, NULL, NULL, 'Carretera 932 Barrrio Rincón', 0),
(276, NULL, '2024-06-06 00:00:00', 'https://www.facebook.com/photo/?fbid=761821066122557&set=pcb.761821429455854&__cft__[0]=AZWX-CNaIsE9FMSIaTliFdxifMOWtZJsDijZPrHn6IYjxYlbmfbeflB5N_iy3IqsS1GFg8Mf_VNQBZPbWy35r3CpOs3udHZkZaLknbxIjpojmkOoZ6VO_0iJpb8QIU46EbB0XpTyS-1O25xnQNSU75KV5VvMbSi6NfJMnaoompMfFW8K326cxnJWQf1RRg2ZtZRWbciHNcdL-ANxKk83b7PiFqpNb6g6hqZYxIWxWauGOjcnI3HJ8fe8h7smNbOxKbg&__tn__=*bH-y-R', 'Naguabo', '2024-06-06_1', NULL, NULL, NULL, NULL, NULL, 'Carretera Estatal PR #191', 0),
(277, NULL, '2024-06-20 00:00:00', 'https://www.facebook.com/photo/?fbid=805933118348059&set=pcb.805933885014649&__cft__[0]=AZWnRe1IWn4Q9DBV2pIFlPnCxqFF0yUjBeO_1vqjgFj0tQPTEJNxemXDGaw0NPoOAr5IZZprOUTp98_4ALDa6RguFUhzAiaMb_4gIb1cvLM1dDj_kjziV998cjfVd7oe9cnV3BiDZKK51cNMhe-dH3kv1i9XtjD58-BA6y-1979H0KtZWV7IvvpfwfXl2ykMi5e-jmR5feA3qQy5gIvNiZ5XvXjPxJL_r7rWYp7iPmiJWTWG_IdgEM4K0C0OhOQM-68&__tn__=*bH-y-R', 'Lares ', '2024-06-20_1', NULL, NULL, NULL, NULL, NULL, 'Antes del puente de Rio Prieto', 0),
(278, NULL, '2024-06-27 00:00:00', 'https://www.telemundopr.com/noticias/puerto-rico/se-desprende-parte-de-la-carretera-en-la-pr-1/2620559/', 'Caguas hacia San Juan Carretera PR-1', '2024-06-27_1', NULL, NULL, NULL, NULL, NULL, 'Carril de la PR-1', 0),
(279, NULL, '2024-07-29 00:00:00', 'https://www.facebook.com/photo/?fbid=1082529626578340&set=a.200362694795042&__cft__[0]=AZWUG9Z5I484zdV2rWK-Zh9cAvvhYo0PRc9oTvCd9cTZVJM5ZPsRwLT-kiZiJ-KS7fNOd6MtjDWCNnw-7DMSV15kgKAK7ppLKjGE1bNPiPTJ-E5bsb-hxlYWmsyrdoHDDY_EnF8B4cyN-6lKmpxA9fHpeeOt690fpYmEUvDUqLLw54GFfgAcqLtmW7Kwuc-2vcF0cuc8fMsSwADaLbg6g54uhJdNoIE3se3-MJH_y1zWt2wzo6tXkdjSgt_HNAIY0Vg&__tn__=EH-y-R', 'Aguada ', '2024-07-29_1', NULL, NULL, NULL, NULL, NULL, NULL, 0),
(280, NULL, '2024-07-30 00:00:00', 'https://www.facebook.com/photo/?fbid=887799900049727&set=a.354021900094199&__cft__[0]=AZWFhi7uN_B2WZ61924CnzaiT95nQuG87t_z-TnEDJxCHPout0tkzCGMba12NQm6KFXS7uy4a723korKcotbypRJluNA0R9ABmr6ZujXoHn3UhijmxGfXvoHTG2hzC4a8ni0Vv9DIafXsVqYO0-dC6swxERaUSzqgINrkyMjFNkrc2W7RYULCm9SMNamtl-eirWGc1WCPzowdLYcHXwFyGIC2Dy1tJcuJYkQvhm6Oi3gzgG7YkGi0jCEV39eRFNkvd8&__tn__=EH-y-R', 'Guaynabo ', '2024-07-30_1', NULL, NULL, NULL, NULL, NULL, 'Carmen Hills, La calle Sunset Boulevard ', 0),
(281, NULL, '2024-08-13 00:00:00', 'https://www.facebook.com/photo/?fbid=931798472310662&set=pcb.931798612310648&__cft__[0]=AZVynntYQ6PR2842Qi0Z1dJdxZ_iIGAqI7DsE9Xuk9TDJaDu6DI-gHzpe9z_0ib3hpgJRfejYS9VuQPBmKkqo1Eee4AO1-nPPch3j-NIieKPK4wL2Gj7dnKnk2YKic9Ra4kG_ty-tQoZ3nUgQFGNM0lsQ_MXD6IjOzQXGSsM9BJIlkySTteP2thAvlcc7QvhgV5ON_tiSXJ1GHy86uf5gpGVo8kABKIz5eDrgw3Ti8LxkcGToowDXPyyToQEJdE2f1A&__tn__=H-y-R', 'Barranquitas ', '2024-08-13_1', NULL, NULL, NULL, NULL, NULL, 'Carr. 771 cerca de la estación de gasolina Tavín', 0),
(282, NULL, '2024-08-14 00:00:00', 'https://www.facebook.com/photo/?fbid=519152110508345&set=a.218252060598353&__cft__[0]=AZUs3hMpTQH1mpB36ILf2cn6dQmNuss4Dewq_UGeP4MkChAuJT26zlH8JhqHL85ZzZU6-q-ypahxrkNeeh0Gvy0xKmmg6fbF63A1sG7J9KIyNLrtnUr1uWdRUWJteF8wNM0S5AxgyZZwBZxmYVvGGNBHWWOQHXUIMu9-82v_HKy2Rm8RL0SCLWRGhJt4M49LDnxKWW89_ps8FQB-gjWWO6_qWG4igtbp2PQmGgwMuCb0FHcuOYTuxW9NV-skVi9w2GY&__tn__=EH-y-R', 'Mayagüez ', '2024-08-14_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-105 en el kilómetro 6.4', 0),
(283, NULL, '2024-08-14 00:00:00', 'https://www.primerahora.com/noticias/policia-tribunales/notas/no-hay-paso-en-cuatro-carreteras-en-san-sebastian/?utm_source=Social&utm_medium=Facebook&utm_campaign=primera_hora', 'San Sebastian', '2024-08-14_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-445 Barrio Saltos, Carretera PR-119 Barrio Guacio y Culebrinas, Carretera PR-446 Barrio Robles ', 0),
(284, NULL, '2024-08-14 00:00:00', 'https://www.facebook.com/photo/?fbid=805075751801368&set=pcb.805075321801411&__cft__[0]=AZVi0XjuwZk9TlKFlM3rF0Lx7P_ba0lMpRk6Vv89oQXWJSF4Fakhw8QpsfJnVeoQuFolhlPLZ3dlxbdQC4-M1ar-9dOarlxt3RnShFlq_6a6ODBIxe6IQ27tUAuU69p9NAnl0p1eluwb-QzmhWrulGMzQj5mMXbwPRz4cvgKu87p9ug-3EQBCJttCmBxdrPCnPyeLbkqGbA9duhVYhBc2P5Lh46PtFWTMFd7NyDcxYhOWNGJmG5EkSlqy9fqARq0zOw&__tn__=*bH-y-R', 'Aguada ', '2024-08-14_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-2 Barrio Malpaso ', 0),
(285, NULL, '2024-08-14 00:00:00', 'https://www.facebook.com/photo/?fbid=1882835165583984&set=a.588053345062179&__cft__[0]=AZVqv68Gd3gaXPQc8nG-2LZ4a_Dd67pECGej07pJxcTl_GtdXM_2gvlUp6Y_v_rER_aUp0M98qTb2K4E8ob7P8f0wT-mamBcKQo19x5AakWCrneeLZMzaZ1_c83-OIwBdP29i1uOjuVMQUVKPwLQqG6PCa-7wXuSc0yVbkBGFyAEK0KqCoenEl7qRMQbFFhHBZk20BSXlacb__VW90xbectiFa_6CSBhZzILzmh_AJ-mZAVo3YjqF5Edxnrg4mCKwJSVbPyC1uiGGyuFFmas7xK82_bKQ-BquuKyTYnrYkrSmA&__tn__=EH-y-R', 'Cidra ', '2024-08-14_4', NULL, NULL, NULL, NULL, NULL, 'Barrio Rabanal Sector La Frontera', 0),
(286, NULL, '2024-08-14 00:00:00', 'https://www.facebook.com/photo/?fbid=1022174969911763&set=a.460809269381672&__cft__[0]=AZXDrUml2sI_fRGDOl06P4QphvuSinTnVLyYVazm-pSqdXunAPvvmDpqezvz3_YxuZA4pS2JuYW9cjeLrMudr3WifOPbLi0oGG0eBW4x30CzS4VAYpY0wuThGIiNeqmiyX4xFQzGK3NlAWrVxY1lBLGZOkWGAc6S_266BnskMvRYkevNyp5hn9eEOxBTat6iDd-rkfLTyb99yMsRChVcVw8a&__tn__=EH-R', 'Naguabo ', '2024-08-14_5', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-191 Barrio Cubuy', 0),
(287, NULL, '2024-08-14 00:00:00', 'https://www.facebook.com/photo/?fbid=1048764277253238&set=pcb.1048764367253229&__cft__[0]=AZUucOCB5CcZ9AQZswTG71it5HvL7UsQsVPlDemty5zWI581smzRKO57ay8rIGRlS-hKwbs-g-2Ikz8vY1QNpoDnZKpk_hvz3PiT8wxNbQ0tWMV2KCjZEtuUY1nLBip45vV3q2h21xEVgHIe5jHpuvOFJcj4mORaHhzV2vYM67pK14j-M0mJ5VAVuFj12rrJlvZcVM42DG9DfTQjaDMJ28f9wwy-6d_ADFGmtvdZjatQAgxnsNzj1Al4phUKQtzQPGU&__tn__=*bH-y-R', 'Coamo', '2024-08-14_6', NULL, NULL, NULL, NULL, NULL, 'Carretera 155 en el kilómetro 13.8, en la vía que conecta a Coamo con Orocovis', 0),
(288, NULL, '2024-08-15 00:00:00', 'https://www.facebook.com/photo/?fbid=1017139670422185&set=pcb.1017129390423213&__cft__[0]=AZUV4e_3NTdbZk5s8XBGNfdcFzvDVsDv9lXm7cVpOB6btNqqTIbNVJit_WgXkkAcm-o76qZtrq3ywjPJaw-em8a6DJjIXjOF6GuMtrvwNQm0EIPV8U2YNxazsFX6oc7BiPwVL_fixbbR5U1TdiYeOMUGnjg4WJCJMyXpneoq-HvAKQyy7cO6C8pPEXl41TvgrJyuk60pxKZcZy03AO0UZ1Cm&__tn__=*bH-R', 'Ponce', '2024-08-15_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-507 hacia el Parque Wito Morales', 0),
(289, NULL, '2024-08-16 00:00:00', 'https://www.facebook.com/photo/?fbid=917721430399644&set=a.361783352660124&__cft__[0]=AZUEvag7hjZhIH7ehcvvcoYMDRzmgo-nH6x8qxHE0E-0pIpD_4Pz6QjbcbyNzfthtGv69P8grQbkrCmaohtpZ36GFQHA8mfbjhOb2YzImOFLR2Z2qdbWOjVwiqw6P3LVFWbzeBy2ybxQSLENJmjKrI106KaBH-3bz9U1VR63Jxh8BLeRTC8Pnqek4SbUzoZaW08IQdt9aGSMIW_9eeEjrR0Yx0wf_PAr4YY5zwYtaoysYdS49CNf13ipEsgiBjfF8-g&__tn__=EH-y-R', 'Lares y San Sebastian ', '2024-08-16_1', NULL, NULL, NULL, NULL, NULL, 'Lares ramal 111 luego del antiguo salón El Imperial. Abonados en Seburuquillo, Bo. Pueblo, Palmas del Sol, Juncal', 0),
(290, NULL, '2024-08-21 00:00:00', 'https://www.facebook.com/photo/?fbid=908151278012940&set=pcb.908151358012932&__cft__[0]=AZW5OzerleykHgirAip72oXJyb5RjtaIX5tf-I1q33qK_qMlXLGnV3s4knSfgecvmz-L0y6QAXUfqCR5wY_g0zS9y9HdrTjWQWuehfQFuqtvd1IbpjzgEEJbailrsHZZCKmKxv2Q-ZIaAl0woLIbyoq-gEvqXeZ2nqqTtnpUHuMTDHBYadbSmTnkRIl4DxcCujUC5OMtQap349MRoelnz2Gm57AciMO5L0SqtOW6mfBCoj3AaJoeiSbm8VrpmHeK15c&__tn__=*bH-y-R', 'Ponce-Adjuntas ', '2024-08-21_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-10 kilómetro 16.1 Carril de Adjuntas hacia Ponce', 0),
(291, NULL, '2024-08-24 00:00:00', NULL, 'Mayagüez ', '2024-08-24_1', NULL, NULL, NULL, NULL, NULL, 'Km 4.3 Carretera 105 Barrio El Limón', 0),
(292, NULL, '2024-08-25 00:00:00', 'https://www.facebook.com/photo/?fbid=923282419843545&set=a.361783355993457&__cft__[0]=AZW-gA106WkvHbP2VFXleRcCpG_EfBRUg3xComNzRmFbi5xNb76gyzE7d0AvRl8SzxoK7eBFUcSS3NieYTq7q0hMROiEkGEX_SE_A6gZYka0h_g1-griaDD8V31aNQJd0yL0Zq7q7tDKr0DcSjO_9jBx9i_7USkbJZomt7bLBZHmqLZrtEfYc6cPIBHCyKLylgBBz00oITUBVNZy_JkfGllsOFZ38xDb9NZLpOPF3xy_kGQJwWXfT1_qKLD57W4jI3c&__tn__=EH-y-R', 'Orocovis ', '2024-08-25_1', NULL, NULL, NULL, NULL, NULL, 'PR-569 Km 1.5', 0),
(293, NULL, '2024-08-27 00:00:00', 'https://www.facebook.com/photo/?fbid=1050472479768979&set=pcb.1050472526435641&__cft__[0]=AZUfh-LLYRPF91c0MNUtU1VlrG_nw9gMZTs0vaeJsQmckh3RjhG8c5ABiU82AoQy38gsg26kBBbOpdUwCx7mTgZ5SsuDK9VVHo5GO6F0W1xRCLBBih-yHk08W4tgomL4YumlTzThDP-xLbJAAY2BJFq2_6SMfdy6evVtvI3qGlJmNEr2syhYmrRmq-C9AQvFzBFuBE-STybpTOVTUc_eC43oPSFW9clZmX_eZHwMlOFGX9uxHNKr8PBuVc1K0i3uBLc&__tn__=*bH-y-R', 'San Sebastian', '2024-08-27_1', NULL, NULL, NULL, NULL, NULL, 'Salto Collazo ', 0),
(294, NULL, '2024-09-06 00:00:00', 'https://www.facebook.com/photo/?fbid=1056749679141259&set=a.650070479809183&__cft__[0]=AZX4NLtsR2LA1uJ8LXonr1MXYxlTFAu-DfCIWQmgnBOnf62B3vKE-b38T95zMT9J-myArZGdLpXMaPYJdaZk81rvkjCOtseRbHx6Vyf3meMMYXB9wPMDKSKGaHfXX75h8PE5IUBBFvAMfHUsTzVCplkiUPqzD_JmiWK8rt3yL_kMxaw9HtxJZTLFQq6IE75M98FmkI4N8jeYoZHEvYKHREbHL9jw_-BMwLIOCZyEAfc8sQ&__tn__=EH-y-R', 'San Sebastian', '2024-09-06_1', NULL, NULL, NULL, NULL, NULL, 'Sector Bartolo Cordero en el barrio Guacio de San Sebastian ', 0),
(295, NULL, '2024-09-09 00:00:00', 'https://www.facebook.com/photo/?fbid=1085369402947797&set=a.287132782771467&__cft__[0]=AZWFccDu9plpvjhjLm52Gb9s2loAkqzLYpboleuhKY8TgzCR6uA0R4qG26M97jaV6zEUXM0_Cy0TOoqAECKf553_wyiXhgFHaB8ec9A0gmt6K8-XCiwMiOpfM0UAiKQpR3nez7b6WtO6joVIkWssatoUkhaaKu4-0wQQXbbAbP0yfl-ipNcJHYN0rIKJxdmfvZB6aKrERH9N_E2RvKvFcYw2wjt4P8q75KKZMIYKa52xgg&__tn__=EH-y-R', 'San German ', '2024-09-09_1', NULL, NULL, NULL, NULL, NULL, 'Carretera 119 San German ', 0),
(296, NULL, '2024-09-09 00:00:00', 'https://www.facebook.com/photo/?fbid=854244640143891&set=pcb.854236933477995&__cft__[0]=AZXrev2gL58xh6lxb1wOLAq-uTJmKGc1kmz0TDwKG_5AeWn2b5bbrsp2D1i-TMqDcdzJ8n3YlB5RmAFijDGExepfJ6AAmP4RWzPhwNiK6a9A2OxEYpJBL1rutqcSeCxA6gmcXiQclMZDaxkpxh4btImmYupXZrB2LAn44RqEEVbx7uX74Yjzq90LBiNxr0s_SS2CH3IhsBDyQn3MZR4WrD2ratsXwxcj3hEHvULVqK-BTg&__tn__=*bH-y-R', 'Jayuya', '2024-09-09_2', NULL, NULL, NULL, NULL, NULL, 'Carretera #140 km 4.7 Bo. Jauca, Jayuya', 0),
(297, NULL, '2024-09-13 00:00:00', 'https://www.facebook.com/@actpuertorico/?__cft__[0]=AZWwu5GvxGnD5XPwjLKgzuHH-u7mGIcxJekgu4MH66edBUrEeFRTKQEPT55_EnRIXlo6h15ZX0XRnbWQjetmjPQ7BPivuxdu1NH9QJaG9BfWzbbQdcOVCDVAIsuaE1pvMR-777_fidnhqn7LY5Q-Yr74YCECVqaamxm_tosgKfLEmXjt3SC-IQDT1N9-z-2IUrSwxqWiq3fJp-xD2lYS55xOvmdSZdWoLqtYfOgMhEJeGw&__tn__=-UC%2CP-y-R', 'Toa Alta ', '2024-09-13_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-861 ', 0),
(298, NULL, '2024-09-14 00:00:00', 'https://www.facebook.com/photo/?fbid=900827528741369&set=pcb.900827695408019&__cft__[0]=AZWYGyiyvBp_YdAZSPmdGYgHjgRuZoju2SwkoXEvIpumJXUaDCmwcx_CS-AGtET9LtzovU_pn3K7jUf8lXNbLQM3PEuVYSjIQsuXQasm2GgwcYhL1ba7s-z5YNoz7U8qcfPW7S8XeVwS9AjiFnDTcYlSoSZkNUKg6cuAH_TnsP8Xf9IxdBF7uOjsNKqiju_rtnm-1qtpppL-GE5IX52c9MGp9t6SLSUvptd5IVCO4rmW9A&__tn__=*bH-y-R', 'Utuado ', '2024-09-14_1', NULL, NULL, NULL, NULL, NULL, 'Carretera #123 km 42.6', 0),
(299, NULL, '2024-09-20 00:00:00', 'https://www.facebook.com/photo/?fbid=1066128928203334&set=a.650070479809183&__cft__[0]=AZX23fVYvFIQne-85uUb-EP_TE28YVoaFBuSdplQ5z--MjxMDy55-tcKr4RnLUnnTG7kw9SFLrBaUZQNBvCqoJGkHGeWXefHwQjfUPg2oYvON0cBaWLxNdlOq7f03lGLavxe6cXOatfmxPFhxfKjUyQK8DTPrZ1sY6xGKcdImKRkQyp8Btmqsleg4VSFEeDgZEOUw5FG4hf_Kp3XmyvsoHimruG5XUi3m8GukgWoDRh7gw&__tn__=EH-y-R', 'San Sebastian ', '2024-09-20_1', NULL, NULL, NULL, NULL, NULL, 'Carretera 119 cerca de la Iglesia Pentecostal en el Barrio Guacio San Sebastian', 0),
(300, NULL, '2024-09-20 00:00:00', 'https://www.facebook.com/photo/?fbid=1066146114868282&set=a.650070479809183&__cft__[0]=AZWgapPjv9Qrc4nTTa-gapAxSH0ndLPR3ur0Wemat1FF0BdSW6js6d4p1c3KPZHs0-TyAiIUpGVhBDDqW_ElPKDPddwzGlYgHHx_7ZfV0fd8DsZ_tQQ7nsRnZuMR_-bXbmcfkzPT3Uafod8bETW29bp1lOoXy7d34spZrX-zrz3YYuIcHdFBxbj4CAVO4vVWjCYX6mQ1PyQKcI9k_pvs71OBE185_DIQTH_xHGAk7DCxyQ&__tn__=EH-y-R', 'San Sebastian ', '2024-09-20_2', NULL, NULL, NULL, NULL, NULL, 'Carretera 435 cerca del cruce de los barrios Mirabales y Perchas San Sebastian ', 0),
(301, NULL, '2024-09-21 00:00:00', 'https://www.facebook.com/photo/?fbid=1117558199742149&set=a.200362694795042&__cft__[0]=AZXbPHFSxoM6Kj-BrThehNOlwaGOhR_evpGMbU_mR0IrNvUHZv2D7hQT7eBtxfn_is372KzdGwJeHVyTpHejSvCCQDWnTzeU50gwf_OAAhmGZLnnf1VoVWAmNMGH8TKY7XIuWoVL8Snn6wgNCdTAmwwG0QkUj1WYYXCfddA0t0pZRAYperwGUlFmXFPFbZ4kJ5pcjLeH5ezMOzwLhT72v5dcmfouVXFgSl0RMzxfXntOUw&__tn__=EH-y-R', 'Adjuntas hacia Utuado ', '2024-09-21_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-10 de Adjuntas hacia Utuado ', 0),
(302, NULL, '2024-09-21 00:00:00', 'https://www.facebook.com/photo/?fbid=1093650802119657&set=a.287132782771467&__cft__[0]=AZXfH5Y8NIx2EhushoWeD34KShsf9bBYqOozk2EkpbNDv7vx0wH3vpGpjWKi4gJNHtbCVOeSki-TNW5s6OGcKDhAadfmtr1bFmEU6HPrhiFCEkDdrFaEJHbXFdpynAHw5rvnJYuvGF-Thxk-N665B-NrFIkHRbVcCYSm57yaH-XCzZD06ZyIBF3kvCryRoiWLQy1cFgl6m1cKLbUO1_zSwBc8bMi5fTiDhfkSLSSouxTQA&__tn__=EH-y-R', 'Adjuntas ', '2024-09-21_2', NULL, NULL, NULL, NULL, NULL, 'Carretera 123 km 44.5 Barrio Pelljas  cerca de Tony\'s Place de Utuado hacia Adjuntas', 0),
(303, NULL, '2024-11-01 00:00:00', 'https://www.facebook.com/share/1CgAxFbvFh/', 'Las Piedras ', '2024-11-01_1', NULL, NULL, NULL, NULL, NULL, 'Sector Palestina, Barrio Quebrada Grande Las Piedras ', 0),
(304, NULL, '2024-11-01 00:00:00', 'https://www.facebook.com/share/17xd5TPfTk/', 'Maunabo', '2024-11-01_2', NULL, NULL, NULL, NULL, NULL, 'Barrio Lizas, Maunabo ', 0),
(305, NULL, '2024-11-10 00:00:00', 'https://www.facebook.com/share/15cxQFQYqC/', 'Orocovis ', '2024-11-10_1', NULL, NULL, NULL, NULL, NULL, 'PR-157 cerca del Puente de Cacao, Orocovis ', 0),
(306, NULL, '2024-11-10 00:00:00', NULL, 'Coamo', '2024-11-10_2', NULL, NULL, NULL, NULL, NULL, 'Carretera 155 que va da Coamo hacia Orocovis kilometros 5.4, 7.0, 7.4 Barrio Coamo Arriba ', 0),
(307, NULL, '2024-11-11 00:00:00', NULL, 'Caguas', '2024-11-11_1', NULL, NULL, NULL, NULL, NULL, 'Carretera 172 km. 22.3 Caguas ', 0),
(308, NULL, '2024-11-05 00:00:00', 'https://www.facebook.com/share/18Hh9BjDCU/', 'Rio Grande, Yunque ', '2024-11-05_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-191 ', 0),
(309, NULL, '2024-11-10 00:00:00', 'https://www.facebook.com/share/15Jid4fwfY/', 'Maunabo', '2024-11-10_3', NULL, NULL, NULL, NULL, NULL, 'Barrio Emajaguas carretera 901 y sector los Chinos en barrio Palo Seco', 0),
(310, NULL, '2024-11-10 00:00:00', 'https://www.facebook.com/share/1JkDbCsDMr/', 'Naguabo', '2024-11-10_4', NULL, NULL, NULL, NULL, NULL, 'Sierra El Duque, casa colapsa por derrumbe', 0),
(311, NULL, '2024-11-10 00:00:00', 'https://www.facebook.com/share/19NVDDfUep/', 'Caguas', '2024-11-10_5', NULL, NULL, NULL, NULL, NULL, 'Carretera 172 km. 22.3 Caguas ', 0),
(312, NULL, '2024-11-11 00:00:00', 'https://www.facebook.com/share/15aUCb25PY/', 'Yabucoa', '2024-11-11_2', NULL, NULL, NULL, NULL, NULL, 'Carretera 901 Bo. Camino Nuevo, cercano a Cafe Don Mario', 0),
(313, NULL, '2024-11-16 00:00:00', 'https://www.facebook.com/share/p/1Ac5swvNea/', 'Naguabo', '2024-11-16_1', NULL, NULL, NULL, NULL, NULL, 'Carrretera Pr-53, Km 12.3 Direccion Sur hacia Humacao, carril derecho quedo cerrado', 0),
(314, NULL, '2024-11-17 00:00:00', 'https://www.facebook.com/share/p/167wCYKJEn/', 'Humacao', '2024-11-17_1', NULL, NULL, NULL, NULL, NULL, 'Bo. Mariana, Sector Juan Rodriguez, Rockslide, boulder cae al borde de la carretera, se ve una casa a lo lejos en la s fotos, no hay informacion sobre carretera', 0),
(315, NULL, '2024-11-17 00:00:00', 'https://www.facebook.com/share/r/191zb4uQjN/', 'Las Piedras', '2024-11-17_2', NULL, NULL, NULL, NULL, NULL, 'Urbanizacion Extension Oriente, Calle Luis Mu~oz Marin #89, cercano a una residencia, el enlace lleva a un video', 0),
(316, NULL, '2025-02-02 00:00:00', 'https://www.facebook.com/share/p/19C9y6njS9/', 'Utuado', '2025-02-02_1', NULL, NULL, NULL, NULL, NULL, 'Barrio Arenas, Carretera 5523, Km 1.3, se ve que el derrumbe llega a la pared de una residencia con rejas altas, otra residencia de dos pisos aleda~a, se adjuntara un segundo enlace', 0),
(317, NULL, '2025-02-05 00:00:00', 'https://www.facebook.com/share/p/1Akvj3T1Bu/', 'Hatillo', '2025-02-05_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-129 direccion de Hatillo hacia Lares, esto no fue por un evento de lluvia', 0),
(318, NULL, '2025-03-05 00:00:00', 'https://www.facebook.com/share/p/15TaGb1BoS/', 'Ciales', '2025-03-05_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-149, Km 22.2, entre la entrada a las Parcelas Maria y el negocio de Enyo, un ciudadano reclama que ya ese evento llevaba tiempo pero fue rellenado por la administracion pasada', 0),
(319, NULL, '2025-04-14 00:00:00', 'https://www.facebook.com/share/p/1JaCcbWYTu/', 'Penuelas', '2025-04-14_1', NULL, NULL, NULL, NULL, NULL, 'El municipio reporta estar limpiando derrumbes en distintos puntos del pueblo pero no se especifica en donde son dichas localizaciones, se menciona que fue por el evento de lluvia en ese mismo dia', 0),
(320, NULL, '2025-04-19 00:00:00', 'https://www.facebook.com/share/p/19CzxgPXPC/', 'Naranjito', '2025-04-19_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-152, derrumbe en un Centro Comercial (Naranjito Shopping Village), se ven unos cuantos carros aplastados por un muro que cayo, parece ser mas efecto de construccion que deslizamiento', 0),
(321, NULL, '2025-04-19 00:00:00', 'https://www.facebook.com/share/p/1BE5MLTfvD/', 'Ciales', '2025-04-19_2', NULL, NULL, NULL, NULL, NULL, 'No se encuentra informacion de donde exactamente ocurrio el deslizamiento, se ve que se derrumbo el costado de una residencia', 0),
(322, NULL, '2025-04-19 00:00:00', 'https://www.facebook.com/share/p/16DzzwSysg/', 'Barranquitas', '2025-04-19_3', NULL, NULL, NULL, NULL, NULL, 'Barrio Ca~abon, en direccion hacia la Torre', 0),
(323, NULL, '2025-04-19 00:00:00', 'https://www.facebook.com/share/p/1E7ZWX14yZ/', 'Orocovis', '2025-04-19_4', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-155, no hay mas informacion', 0),
(324, NULL, '2025-04-19 00:00:00', 'https://www.facebook.com/share/p/1E7ZWX14yZ/', 'Caguas', '2025-04-19_5', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-172, Km 22.4, no hay ms informacion', 0),
(325, NULL, '2025-04-19 00:00:00', 'https://www.facebook.com/share/p/1H7VPRu4a5/', 'Morovis', '2025-04-19_6', NULL, NULL, NULL, NULL, NULL, 'Carretera que conduce al Barrio Pasto, Km 6.9', 0),
(326, NULL, '2025-04-19 00:00:00', 'https://www.telemundopr.com/noticias/telenoticias/impresionante-casas-llenas-de-lodo-por-copiosas-lluvias/2709946/?fbclid=IwY2xjawKEQG9leHRuA2FlbQIxMQBicmlkETFsYVh6YnlrdDRJcDd3dzdJAR7nDyCObGQ8zk_TTO0BmZ7u6S0JrfnGKoTt7uDuYkNycU7Aw3hKWI8Nmm3c2w_aem_2ft1XVFZwshwbcwgPv8e2Q', 'Barranquitas', '2025-04-19_7', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-770, Barrio Ca~abon Sector La Colina, Se reportan 6 deslizamientos en esa comunidad', 0),
(327, NULL, '2025-04-19 00:00:00', 'https://www.telemundopr.com/noticias/telenoticias/impresionante-casas-llenas-de-lodo-por-copiosas-lluvias/2709946/?fbclid=IwY2xjawKEQG9leHRuA2FlbQIxMQBicmlkETFsYVh6YnlrdDRJcDd3dzdJAR7nDyCObGQ8zk_TTO0BmZ7u6S0JrfnGKoTt7uDuYkNycU7Aw3hKWI8Nmm3c2w_aem_2ft1XVFZwshwbcwgPv8e2Q', 'Naranjito', '2025-04-19_8', NULL, NULL, NULL, NULL, NULL, 'Carretera PR -772, Barrio Nuevo, Sector Mulita, Socavon (no seguro si cuenta como deslizamiento), a las 11 de la noche una explosion y era la carretera, Se dice que se tapo de sedimento la tuberia, no aguanto la presion y exploto', 0),
(328, NULL, '2025-04-19 00:00:00', 'https://wapa.tv/noticias/locales/deslizamientos-y-familias-incomunicadas-en-naranjito-tras-fuertes-lluvias/article_92144f64-31a2-43b2-b946-ce25334e9c18.html', 'Naranjito', '2025-04-19_9', NULL, NULL, NULL, NULL, NULL, 'Barrio Cerro Abajo, Sector El Pilon, la due~a de la residencia informa que la parte trasera de la casa cedio, menciona que la escorrentia de la carretera baja por su solar, cae en el pozo muro el cual se lleno y cedio', 0),
(329, NULL, '2025-04-19 00:00:00', 'https://wapa.tv/noticias/locales/familias-quedan-incomunicadas-por-deslizamientos-provocados-por-lluvias/article_85b70f7c-edcf-453f-976d-2f617e63b2d0.html', 'Manati', '2025-04-19_10', NULL, NULL, NULL, NULL, NULL, 'Barrio Pugnado, Sector Los Burgos, el deslizamiento bloqueo la unica carreterra que lleva a una fila de residencias', 0),
(330, NULL, '2025-04-20 00:00:00', 'https://www.facebook.com/share/p/155onuKykek/', 'Aibonito', '2025-04-20_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-14, en direccion de Aibonito hacia Coamo (en las curvas)', 0),
(331, NULL, '2025-04-20 00:00:00', 'https://www.facebook.com/share/p/1Ga91C4XLL/', 'Morovis', '2025-04-20_2', NULL, NULL, NULL, NULL, NULL, 'Barrio Unibon, Sector Patron, deslizamiento en la entrada de una residencia, da~o visible en vehiculos', 0),
(332, NULL, '2025-04-20 00:00:00', 'https://www.facebook.com/share/p/18m4CYMjR1/', 'Morovis', '2025-04-20_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-155, Km 37.1, Barrio Perchas, Rockslide', 0),
(333, NULL, '2025-04-20 00:00:00', 'https://www.facebook.com/share/r/16ZzD7Gyks/', 'Cayey', '2025-04-20_4', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-184, Km 20.8, region de Guavate, NO ES UN DESLIZAMIENTO, es una area en la que se esta trabajando y los residentes estan preocupados de un POSIBLE deslizamiento.', 0),
(334, NULL, '2025-04-21 00:00:00', 'https://www.facebook.com/share/v/1FxtmHkZpo/', 'Guayama', '2025-04-21_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-179, Km 10.8, Barrio Carite, no se ve a simple viste pero dos barricadas de hormigon cedieron con el suelo (levemente)', 0),
(335, NULL, '2025-04-25 00:00:00', 'https://www.facebook.com/share/v/1BJkHSpsPP/', 'Cidra', '2025-04-25_1', NULL, NULL, NULL, NULL, NULL, 'Barrio Ceiba, Comunidad Hernandez, colapso de terreno donde se realiza obra municipal, el viernes 25 colapso y se rompio un tubo, con el tuvo y las fuertes lluvias el terreno cedio aun mas al dia siguiente', 0),
(336, NULL, '2025-04-26 00:00:00', 'https://www.facebook.com/share/p/18yyQmHtBr/', 'Cayey', '2025-04-26_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-1, Km 62, deslizamiento por paso de vaguada', 0),
(337, NULL, '2025-04-26 00:00:00', 'https://www.facebook.com/share/p/1BuWxp3KDB/', 'Utuado', '2025-04-26_2', NULL, NULL, NULL, NULL, NULL, 'No se brinda ninguna informacion al respecto. Solo se mencionan varios deslizamientos en el municipio. En uno de los comentarios se un residente del municipio muestra una imagen de un peque~o deslizamiento que dice haber sido en la Carretera PR-6103', 0),
(338, NULL, '2025-04-28 00:00:00', 'https://www.facebook.com/MunGuaynabo/posts/pfbid02H9bW45iubV2e65K8n1rK958bgfQ93fs1EnKmQzPxCaDrqEtANKgNTuQGMaVw8mJEl', 'Guaynabo', '2025-04-28_1', NULL, NULL, NULL, NULL, NULL, 'Barrio Hato Nuevo, Sector Las Lomas, obstaculiza el camino a 3 residencias, el director de la oficina de manejo de emergencias dijo que han atendido varios caos de deslizamientos en el municipio a causa de las recientes lluvias', 0),
(339, NULL, '2025-04-28 00:00:00', 'https://www.facebook.com/share/p/1BjFnG4pXf/', 'Comerio', '2025-04-28_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156, Km 28.7, donde era la planta de gas del Barrio Rio Hondo, Sector La Juncia', 0),
(340, NULL, '2025-04-28 00:00:00', 'https://www.facebook.com/share/v/14J17jX1c6w/', 'San Juan', '2025-04-28_3', NULL, NULL, NULL, NULL, NULL, 'Barrio Caimito, no hay mas informacion al respecto', 0),
(341, NULL, '2025-04-28 00:00:00', 'https://www.facebook.com/share/p/1AQN8pNqyc/', 'Maricao', '2025-04-28_4', NULL, NULL, NULL, NULL, NULL, 'No hay detalles del barrio o carretera. El post se hizo el dia 29 pero un comentario de una ciudadana dice que ese deslizamiento ocurrio el dia anterior.', 0),
(342, NULL, '2025-04-29 00:00:00', 'https://www.facebook.com/share/p/19FttPyUii/', 'Ciales', '2025-04-29_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-615, Km 8 entre el Taller de Hojalateria Shelo y el Puente de Pozas (sin servicio electrico la comunidad)', 0),
(343, NULL, '2025-04-29 00:00:00', 'https://www.facebook.com/share/p/12Jf4mnCr68/', 'Yauco', '2025-04-29_2', NULL, NULL, NULL, NULL, NULL, 'Barrio Rubias', 0),
(344, NULL, '2025-04-29 00:00:00', 'https://www.facebook.com/share/p/1AFQTy2Ra8/', 'Comerio', '2025-04-29_3', NULL, NULL, NULL, NULL, NULL, 'Subien Sector La Prieta', 0),
(345, NULL, '2025-04-29 00:00:00', 'https://www.facebook.com/share/p/15XrmBvpRP/', 'Naranjito', '2025-04-29_4', NULL, NULL, NULL, NULL, NULL, 'Barrio Anones, Sector Cayito Rios, se aprecia que si el derrumbe continuo afectara gravemente una residencia en la colina', 0),
(346, NULL, '2025-04-29 00:00:00', 'https://www.facebook.com/share/p/18zBsVjD4k/', 'Orocovis', '2025-04-29_5', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-157, Barrio Damian Arriba (los comentarios dan informacion diferente a la \"\"posteada\"\"', 0),
(347, NULL, '2025-04-30 00:00:00', 'https://www.facebook.com/share/p/1Ltph39KwD/', 'Orocovis', '2025-04-30_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-568, Sector El Puente', 0),
(348, NULL, '2025-04-30 00:00:00', 'https://www.facebook.com/share/p/16UK7oF6gV/', 'Yauco', '2025-04-30_2', NULL, NULL, NULL, NULL, NULL, 'Sectores Collores y Rancheras, hay 4 imagenes en la que se ven dos eventos diferentes, no se puede distinguir cual es cual pero uno esta justo en un letrero que dice Km 12 y el otro deslizamiento en un letrero de Km 11.5', 0),
(349, NULL, '2025-04-30 00:00:00', 'https://www.facebook.com/share/p/16UK7oF6gV/', 'Yauco', '2025-04-30_3', NULL, NULL, NULL, NULL, NULL, 'Sectores Collores y Rancheras, hay 4 imagenes en la que se ven dos eventos diferentes, no se puede distinguir cual es cual pero uno esta justo en un letrero que dice Km 12 y el otro deslizamiento en un letrero de Km 11.5', 0),
(350, NULL, '2025-04-30 00:00:00', '-', 'Sabana Grande', '2025-04-30_4', 18.060000, -66.940000, NULL, NULL, NULL, 'Lat: 18.06, Lon: -66.94, datos enviados por imagen en nuestro grupo de WhatsApp, 1 N Liborio Negron', 0),
(351, NULL, '2025-04-30 00:00:00', '-', 'Sabana Grande', '2025-04-30_5', 18.060000, -66.940000, NULL, NULL, NULL, 'Lat: 18.06, Lon: -66.94, datos enviados por imagen en nuestro grupo de WhatsApp, 1 N Liborio Negron', 0),
(352, NULL, '2025-04-30 00:00:00', '-', 'Barranquitas', '2025-04-30_6', 18.180000, -66.360000, NULL, NULL, NULL, 'Lat: 18.18, Lon: -66.36, datos enviados por imagen en nuestro grupo de WhatsApp, 4 WSW Barranquitas', 0),
(353, NULL, '2025-04-30 00:00:00', 'https://www.facebook.com/share/p/1F6wZzNRUd/', 'Barranquitas', '2025-04-30_7', NULL, NULL, NULL, NULL, NULL, 'Barrio Ca~abon, Las Pinas', 0),
(354, NULL, '2025-04-30 00:00:00', '-', 'Ponce ', '2025-04-30_8', 18.140000, -66.610000, NULL, NULL, NULL, 'Lat: 18.14, Lon: -66.61, datos enviados por imagen en nuestro grupo de WhatsApp, 5 S Jayuya', 0),
(355, NULL, '2025-04-30 00:00:00', '-', 'Salinas', '2025-04-30_9', 18.070000, -66.220000, NULL, NULL, NULL, 'Lat: 18.07, Lon: -66.22, datos enviados por imagen en nuestro grupo de WhatsApp, 1 E Vazquez', 0),
(356, NULL, '2025-04-30 00:00:00', '-', 'Salinas', '2025-04-30_10', 18.040000, -66.220000, NULL, NULL, NULL, 'Lat: 18.04, Lon: -66.22, datos enviados por imagen en nuestro grupo de WhatsApp, 1 WSW La Plena', 0),
(357, NULL, '2025-05-01 00:00:00', 'https://www.facebook.com/share/p/19iXNZpTBu/', 'Morovis', '2025-05-01_1', NULL, NULL, NULL, NULL, NULL, 'Barrio Rio Grande, Sector El Cerro, desprendimiento de carretera cerca de La Capilla, el derrumbe provoco se rompiera tuberia de AAA', 0),
(358, NULL, '2025-05-01 00:00:00', '-', 'Morovis', '2025-05-01_2', 18.300000, -66.410000, NULL, NULL, NULL, 'Lat: 18.30, Lon: -66.41, datos enviados por imagen en nuestro grupo de WhatsApp, 2 S Morovis', 0),
(359, NULL, '2025-05-01 00:00:00', '-', 'Pe~uelas', '2025-05-01_3', 18.090000, -66.720000, NULL, NULL, NULL, 'Lat: 18.09, Lon: -66.72, datos enviados por imagen en nuestro grupo de WhatsApp, 2 N Pe~uelas', 0),
(360, NULL, '2025-05-01 00:00:00', '-', 'Cayey', '2025-05-01_4', 18.100000, -66.200000, NULL, NULL, NULL, 'Lat: 18.10, Lon: -66.20, datos enviados por imagen en nuestro grupo de WhatsApp, 3 NE Vazquez', 0),
(361, NULL, '2025-05-01 00:00:00', 'https://www.facebook.com/share/p/15qumbgNMA/', 'Ciales', '2025-05-01_5', 18.292408, -66.488598, NULL, NULL, NULL, 'Carretera PR-615, Km 3.2, Barrio Pozas, sin servicio electrico la comunidad (N 18.292408, W 66.488598)', 0),
(362, NULL, '2025-05-01 00:00:00', 'https://www.facebook.com/share/p/1Ck6F84FH1/', 'Cayey', '2025-05-01_6', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-15, Km 17.6', 0),
(363, NULL, '2025-05-01 00:00:00', 'https://www.facebook.com/share/p/16QUUxM7VW/', 'Yauco', '2025-05-01_7', NULL, NULL, NULL, NULL, NULL, 'Barrio Collores, Sector Tortugo', 0),
(364, NULL, '2025-05-01 00:00:00', 'https://www.facebook.com/share/p/16YzDizm2L/', 'Sabana Grande', '2025-05-01_8', NULL, NULL, NULL, NULL, NULL, 'Comunidad La Pica en el Camino Los Sanchez', 0),
(365, NULL, '2025-05-02 00:00:00', 'https://www.facebook.com/share/r/1CP2BGZjqM/', 'Orocovis', '2025-05-02_1', NULL, NULL, NULL, NULL, NULL, 'Barrio Saltos, En el patio de una residencia', 0),
(366, NULL, '2025-05-02 00:00:00', 'https://www.facebook.com/share/p/1DhGT8pan1/', 'Utuado', '2025-05-02_2', NULL, NULL, NULL, NULL, NULL, 'PR-10', 0),
(367, NULL, '2025-05-02 00:00:00', 'https://www.facebook.com/share/p/1ATaPsyHFb/', 'Cayey', '2025-05-02_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-1, Direccion de Cayey a Salinas', 0),
(368, NULL, '2025-05-02 00:00:00', 'https://www.facebook.com/share/p/18rdnzQVDK/', 'Guayanilla', '2025-05-02_4', NULL, NULL, NULL, NULL, NULL, 'Barrio Llano', 0),
(369, NULL, '2025-05-02 00:00:00', 'https://www.facebook.com/share/p/19FjQgj9T9/', 'Comerio', '2025-05-02_5', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-781, Km 4, Barrio Cedrito, Sector La Prieta', 0),
(370, NULL, '2025-05-02 00:00:00', 'https://www.facebook.com/share/p/1FJhQxGa8S/', 'Comerio', '2025-05-02_6', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-172, Km 4.2, Barrio Vega Redonda', 0),
(371, NULL, '2025-05-02 00:00:00', 'https://www.facebook.com/share/p/1FWUFPx8pE/', 'Guaynabo', '2025-05-02_7', NULL, NULL, NULL, NULL, NULL, 'Varios deslizamientos reportados. En los comentarios un ciudadano comento la localizacion de uno de los deslizamientos: Carretera PR-8834 Km 6 Hm 0, Barrio Rio, Sector Ferreira', 0),
(372, NULL, '2025-05-02 00:00:00', 'https://www.facebook.com/share/p/1FWQoupyv4/', 'Lares', '2025-05-02_8', NULL, NULL, NULL, NULL, NULL, 'Casco Urbano de Lares. No hay informacion adicional', 0),
(373, NULL, '2025-05-02 00:00:00', 'https://www.facebook.com/share/p/1Yh42xGDt4/', 'Utuado', '2025-05-02_9', NULL, NULL, NULL, NULL, NULL, 'Barrio Rio Abajo', 0),
(374, NULL, '2025-05-02 00:00:00', 'https://wapa.tv/noticias/locales/varios-deslizamientos-afectan-carreteras-del-campo-de-guaynabo/article_e2a4d4f4-d36e-4f8d-813c-fcd62063c489.html', 'Guaynabo', '2025-05-02_10', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-833, Sector Guaragua, se reportaron *4* deslizamientos en dicha carreterra (dicho por Luis Iván Díaz, director de Manejo de Emergencias Municipal)', 0),
(375, NULL, '2025-05-02 00:00:00', 'https://wapa.tv/noticias/locales/varios-deslizamientos-afectan-carreteras-del-campo-de-guaynabo/article_e2a4d4f4-d36e-4f8d-813c-fcd62063c489.html', 'Guaynabo', '2025-05-02_11', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-833, Sector Guaragua, se reportaron *4* deslizamientos en dicha carreterra', 0),
(376, NULL, '2025-05-02 00:00:00', 'https://wapa.tv/noticias/locales/varios-deslizamientos-afectan-carreteras-del-campo-de-guaynabo/article_e2a4d4f4-d36e-4f8d-813c-fcd62063c489.html', 'Guaynabo', '2025-05-02_12', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-833, Sector Guaragua, se reportaron *4* deslizamientos en dicha carreterra', 0),
(377, NULL, '2025-05-02 00:00:00', 'https://wapa.tv/noticias/locales/varios-deslizamientos-afectan-carreteras-del-campo-de-guaynabo/article_e2a4d4f4-d36e-4f8d-813c-fcd62063c489.html', 'Guaynabo', '2025-05-02_13', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-833, Sector Guaragua, se reportaron *4* deslizamientos en dicha carreterra', 0),
(378, NULL, '2025-05-02 00:00:00', 'No hay enlace', 'Utuado', '2025-05-02_14', NULL, NULL, NULL, NULL, NULL, 'Hay 11 imagenes de un total de 8 deslizamientos enviados por la OMME a Isabella, se conoce que los deslizamiento ocurrienron en el dia 2 y 3 de mayo de 2025.', 0),
(379, NULL, '2025-05-02 00:00:00', 'No hay enlace', 'Utuado', '2025-05-02_15', NULL, NULL, NULL, NULL, NULL, 'Hay 11 imagenes de un total de 8 deslizamientos enviados por la OMME a Isabella, se conoce que los deslizamiento ocurrienron en el dia 2 y 3 de mayo de 2025.', 0),
(380, NULL, '2025-05-02 00:00:00', 'No hay enlace', 'Utuado', '2025-05-02_16', NULL, NULL, NULL, NULL, NULL, 'Hay 11 imagenes de un total de 8 deslizamientos enviados por la OMME a Isabella, se conoce que los deslizamiento ocurrienron en el dia 2 y 3 de mayo de 2025.', 0),
(381, NULL, '2025-05-02 00:00:00', 'No hay enlace', 'Utuado', '2025-05-02_17', NULL, NULL, NULL, NULL, NULL, 'Hay 11 imagenes de un total de 8 deslizamientos enviados por la OMME a Isabella, se conoce que los deslizamiento ocurrienron en el dia 2 y 3 de mayo de 2025.', 0),
(382, NULL, '2025-05-03 00:00:00', 'https://www.facebook.com/share/p/15gxCnZjPk/', 'Vega Baja', '2025-05-03_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-155, en los comentarios una ciudadana menciona un boulder cayo en el patio de su residencia y coloca la foto cmo evidencia', 0),
(383, NULL, '2025-05-03 00:00:00', 'https://www.facebook.com/share/p/1FYqY8yajh/', 'Yabucoa', '2025-05-03_2', 18.026141, -65.848598, NULL, NULL, NULL, 'Sector Nuevo Cu~o, Camino Nuevo, incluyeron enlace de google maps: https://www.google.com/maps/place/18%C2%B001\'34.1%22N+65%C2%B050\'45.7%22W/@18.0261409,-65.8485978,826m/data=!3m2!1e3!4b1!4m4!3m3!8m2!3d18.0261409!4d-65.8460229?entry=ttu&g_ep=EgoyMDI1MDUwNS4wIKXMDSoASAFQAw%3D%3D', 0),
(384, NULL, '2025-05-03 00:00:00', 'https://www.facebook.com/share/p/16Qz2Z8mbJ/', 'Ceiba', '2025-05-03_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-978, Km 2.4', 0),
(385, NULL, '2025-05-03 00:00:00', 'https://www.facebook.com/share/r/1Wc6Uyb8ft/', 'Toal Alta', '2025-05-03_4', 18.401056, -66.231333, NULL, NULL, NULL, 'Urbanizacion Estancias de la Fuente, 18°24\'03.8\"N 66°13\'52.8\"W, https://www.youtube.com/watch?v=9TZPtpzouw4', 0),
(386, NULL, '2025-05-03 00:00:00', 'No hay enlace', 'Caguas', '2025-05-03_5', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-761, Alturas de Borinquen (mencionado por Tania en el chat de la oficina)', 0),
(387, NULL, '2025-05-03 00:00:00', 'No hay enlace', 'Utuado', '2025-05-03_6', NULL, NULL, NULL, NULL, NULL, 'Barrio Tetuan, mencionado por Isabella en el chat de la oficina, hay foto. Fecha tentativa, se entiende que fue el dia 2 de mayo o el 3 de mayo de 2025. ', 0),
(388, NULL, '2025-05-03 00:00:00', 'https://www.facebook.com/share/p/16D3jA8eJC/', 'Yauco', '2025-05-03_7', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-374, Km 0.4, Sector Carrizalez (na residencia se encuentra en riesgo en el Municipio de #Yauco Un derrumbe está obstruyendo la carretera 374, kilómetro 0.4. El material vegetativo descio la escorrentía hasta una vivienda)', 0),
(389, NULL, '2025-05-03 00:00:00', 'https://www.facebook.com/share/p/15CjXvgfaM/', 'Comerio', '2025-05-03_8', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-167, en el Salto, cerca de las letras de Comerio, segun los comentarios del post despues de la limpieza municipal, horas despues, el terreno sigui desprendiendose', 0),
(390, NULL, '2025-05-03 00:00:00', 'https://www.facebook.com/share/r/14HuxH5sThg/', 'Naranjito', '2025-05-03_9', NULL, NULL, NULL, NULL, NULL, 'Sector Morales', 0),
(391, NULL, '2025-05-03 00:00:00', 'No hay enlace', 'Utuado', '2025-05-03_10', NULL, NULL, NULL, NULL, NULL, 'Hay 11 imagenes de un total de 8 deslizamientos enviados por la OMME a Isabella, se conoce que los deslizamiento ocurrienron en el dia 2 y 3 de mayo de 2025.', 0),
(392, NULL, '2025-05-03 00:00:00', 'No hay enlace', 'Utuado', '2025-05-03_11', NULL, NULL, NULL, NULL, NULL, 'Hay 11 imagenes de un total de 8 deslizamientos enviados por la OMME a Isabella, se conoce que los deslizamiento ocurrienron en el dia 2 y 3 de mayo de 2025.', 0),
(393, NULL, '2025-05-03 00:00:00', 'No hay enlace', 'Utuado', '2025-05-03_12', NULL, NULL, NULL, NULL, NULL, 'Hay 11 imagenes de un total de 8 deslizamientos enviados por la OMME a Isabella, se conoce que los deslizamiento ocurrienron en el dia 2 y 3 de mayo de 2025.', 0),
(394, NULL, '2025-05-03 00:00:00', 'No hay enlace', 'Utuado', '2025-05-03_13', NULL, NULL, NULL, NULL, NULL, 'Hay 11 imagenes de un total de 8 deslizamientos enviados por la OMME a Isabella, se conoce que los deslizamiento ocurrienron en el dia 2 y 3 de mayo de 2025.', 0),
(395, NULL, '2025-05-04 00:00:00', 'https://www.facebook.com/share/v/14kEczuikQ/', 'Orocovis', '2025-05-04_1', NULL, NULL, NULL, NULL, NULL, 'Barrio Pellejas 2', 0),
(396, NULL, '2025-05-04 00:00:00', '                          ', 'Naranjito', '2025-05-04_2', NULL, NULL, NULL, NULL, NULL, 'En el patio de una residencia en dicho pueblo, Barrio Cedro Arriba, Sector Feijoo por la academia Santa Teresita', 0),
(397, NULL, '2025-05-04 00:00:00', 'https://www.facebook.com/share/p/16YZC1WW6E/', 'Comerio', '2025-05-04_3', NULL, NULL, NULL, NULL, NULL, 'Barrio Cedrito, 2 residencias afectadas. Hay un seundo enlace: https://www.facebook.com/IrvinRiveraAlcalde/posts/pfbid02WPfw4PLGBH45RbhouMQaDCsJWXmFYZ1xvgn5jG6KXG7XC7C4t726iwaXJ1tv9zNdl', 0),
(398, NULL, '2025-05-04 00:00:00', 'https://www.facebook.com/share/r/1B3Fp1eqW4/', 'Arecibo', '2025-05-04_4', NULL, NULL, NULL, NULL, NULL, 'Barrio Rio Arriba', 0),
(399, NULL, '2025-05-04 00:00:00', 'https://www.facebook.com/share/p/1BoExm1X1v/', 'Comerio', '2025-05-04_5', NULL, NULL, NULL, NULL, NULL, 'Barrio Pi~as', 0),
(400, NULL, '2025-05-04 00:00:00', 'https://www.facebook.com/share/p/1CBj9Kgcyb/', 'Comerio', '2025-05-04_6', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-781, Km 4, Barrio El Cedrito', 0),
(401, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/1ZFUDbS4SA/', 'Utuado', '2025-05-05_1', NULL, NULL, NULL, NULL, NULL, 'Discrepancia en la Carretera, algunos residentes dicen,  PR-621, mientras otros dicen PR-123, el kilometro se ve en una de las imagenes, Km 67.5, son dos deslizamientos en la misma carretera. ', 0),
(402, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/1ZFUDbS4SA/', 'Utuado', '2025-05-05_2', NULL, NULL, NULL, NULL, NULL, 'Discrepancia en la Carretera, algunos residentes dicen,  PR-621, mientras otros dicen PR-123, el kilometro se ve en una de las imagenes, Km 67.5, son dos deslizamientos en la misma carretera. ', 0),
(403, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/16ZvQPbSkW/', 'Utuado', '2025-05-05_3', NULL, NULL, NULL, NULL, NULL, 'No se especifica carretera ni barrio', 0),
(404, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/16AQ2AoyMc/', 'Comerio', '2025-05-05_4', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-779, Km 6, Barrio Palomas (frente a una estructura abandonada)', 0),
(405, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/18tdMehdor/', 'Comerio', '2025-05-05_5', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156, Km 30.2, Barrio Río Hondo', 0),
(406, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/12KG3w33MKa/', 'Comerio', '2025-05-05_6', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156, Km 30.2, Barrio Río Hondo', 0),
(407, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/16GrUJfokX/', 'Comerio', '2025-05-05_7', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156, Km 29.9, Barrio Río Hondo', 0),
(408, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/14hx2z6Jds/', 'Comerio', '2025-05-05_8', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156, Km 29.5, Barrio Río Hondo', 0),
(409, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/16QwwcdeHy/', 'Comerio', '2025-05-05_9', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156 Km 23.8, Barrio La Juncia', 0),
(410, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/16QwwcdeHy/', 'Comerio', '2025-05-05_10', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156 Km 25.6, Barrio La Juncia', 0),
(411, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/16QwwcdeHy/', 'Comerio', '2025-05-05_11', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156 Km 26.4, Barrio La Juncia', 0),
(412, NULL, '2025-05-05 00:00:00', 'https://www.facebook.com/share/p/165wgsqHxC/', 'Cabo Rojo', '2025-05-05_12', NULL, NULL, NULL, NULL, NULL, 'Deslizamiento de tierra en la carretera #100 en Cabo Rojo en el semaforo del Panapen.', 0),
(413, NULL, '2025-05-06 00:00:00', 'https://www.facebook.com/share/p/1LDAwuT3Qn/', 'Comerio', '2025-05-06_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-167, Km 4.2, Barrio El Salto.', 0),
(414, NULL, '2025-05-07 00:00:00', 'https://www.facebook.com/share/p/197rCP78Tm/', 'Utuado', '2025-05-07_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-10, Km 45.3', 0),
(415, NULL, '2025-05-07 00:00:00', 'https://www.facebook.com/share/p/16EwVijLud/', 'Comerio', '2025-05-07_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-774, Sector La Mora', 0),
(416, NULL, '2025-05-07 00:00:00', 'https://www.facebook.com/share/p/16EwVijLud/', 'Comerio', '2025-05-07_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-791, Frente a la Gomera Santos, este ciudadano menciono esta informacion en un comentario respecto a una publicacion del Municipio. No se sabe con exactitud la fecha, se estima el 6 o 7 de mayo de 2025', 0),
(417, NULL, '2025-05-07 00:00:00', 'https://www.facebook.com/share/p/16EwVijLud/', 'Comerio', '2025-05-07_4', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-791, Frente a la Gomera Santos, este ciudadano menciono esta informacion en un comentario respecto a una publicacion del Municipio. No se sabe con exactitud la fecha, se estima el 6 o 7 de mayo de 2026', 0);
INSERT INTO `report` (`report_id`, `landslide_id`, `reported_at`, `description`, `city`, `image_url`, `latitude`, `longitude`, `reporter_name`, `reporter_phone`, `reporter_email`, `physical_address`, `is_validated`) VALUES
(418, NULL, '2025-05-07 00:00:00', 'https://www.facebook.com/share/p/1BhV42t3Jw/', 'Maricao', '2025-05-07_5', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-128, Se menciona que es limpieza de los deslizamieno ocurridos por las lluvias de los pasados dias, por lo que el deslizamiento no necesariamente fue el dia 7', 0),
(419, NULL, '2025-05-07 00:00:00', 'https://www.facebook.com/share/p/1Z9jFjfoZG/', 'Comerio', '2025-05-07_6', NULL, NULL, NULL, NULL, NULL, 'Camino “Los González”, Barrio Doña Elena.', 0),
(420, NULL, '2025-05-08 00:00:00', 'https://www.facebook.com/share/p/1FxCmv5igk/', 'Comerio', '2025-05-08_1', NULL, NULL, NULL, NULL, NULL, 'Sector La Mora, entrada a una residencia', 0),
(421, NULL, '2025-05-08 00:00:00', 'https://www.facebook.com/share/p/1XaXHZjRZo/', 'Comerio', '2025-05-08_2', NULL, NULL, NULL, NULL, NULL, 'Barrio Pi~as', 0),
(422, NULL, '2025-05-08 00:00:00', 'https://www.facebook.com/share/p/1AXxuxiPaU/', 'Comerio', '2025-05-08_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-791, Barrio Cejas, peque~os deslizamiento a lo largo de la carreteras (se ven dos)', 0),
(423, NULL, '2025-05-08 00:00:00', 'https://www.facebook.com/share/p/1AXxuxiPaU/', 'Comerio', '2025-05-08_4', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-791, Barrio Cejas, peque~os deslizamiento a lo largo de la carreteras (se ven dos)', 0),
(424, NULL, '2025-05-08 00:00:00', 'https://www.facebook.com/share/p/1DgttKdpj6/', 'Ciales', '2025-05-08_5', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-615, 18.290098, -66.482743', 0),
(425, NULL, '2025-05-12 00:00:00', 'No hay enlace', 'Adjuntas', '2025-05-12_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-131, Km 57.8, Carretea con fango recientemente despejada. Skarp estaba fresco. No pudimos tirar foto por la incomodidad en la carretera, 18°11\'15.8\"N 66°50\'37.1\"W, la fecha adjunta no es seguramente cuando ocurrio si no cuando fue descubierto por nuestro grupo(informacion en el chat de la oficina)', 0),
(426, NULL, '2025-05-12 00:00:00', 'No hay enlace', 'Maricao', '2025-05-12_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-128, Km 30.1, 18°11\'15.8\"N 66°50\'37.1\"W, la fecha adjunta no es seguramente cuando ocurrio si no cuando fue descubierto por nuestro grupo, (informacion del chat de la oficina)', 0),
(427, NULL, '2025-05-13 00:00:00', 'https://www.facebook.com/share/p/194m7i5oPF/', 'Comerio', '2025-05-13_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-775, Interior, Barrio Pi~as', 0),
(428, NULL, '2025-05-13 00:00:00', 'https://www.facebook.com/share/p/196NGw6jrc/', 'Comerio', '2025-05-13_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156, Km 29, antes del puente del Barrio Rio Hondo', 0),
(429, NULL, '2025-05-14 00:00:00', 'https://www.facebook.com/share/p/1AjwwmXbWe/', 'Yauco', '2025-05-14_1', NULL, NULL, NULL, NULL, NULL, 'Barrio Collores', 0),
(430, NULL, '2025-05-14 00:00:00', 'No hay enlace', 'Orocovis', '2025-05-14_2', 18.218446, -66.397224, NULL, NULL, NULL, '18.2184457, -66.3972238', 0),
(431, NULL, '2025-05-14 00:00:00', 'No hay enlace', 'Orocovis', '2025-05-14_3', 18.210453, -66.398004, NULL, NULL, NULL, '18.2104532, -66.3980037', 0),
(432, NULL, '2025-05-14 00:00:00', 'No hay enlace', 'Orocovis', '2025-05-14_4', 18.208567, -66.396826, NULL, NULL, NULL, '18.2085671, -663968259', 0),
(433, NULL, '2025-05-15 00:00:00', 'https://www.facebook.com/share/p/16bWBy71nv/', 'Comerio', '2025-05-15_1', 18.326000, -67.159400, NULL, NULL, NULL, 'Carretera PR-775, Barrio Pi~as', 0),
(434, NULL, '2025-05-17 00:00:00', 'No hay enlace', 'Aguada', '2025-05-17_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-419, Km 1.1, Barrio Naranjo, 18.326, -67.1594', 0),
(435, NULL, '2025-05-17 00:00:00', 'https://www.facebook.com/share/p/1DjjSo3rku/', 'Villalba', '2025-05-17_2', NULL, NULL, NULL, NULL, NULL, 'Barrio Camarones', 0),
(436, NULL, '2025-05-17 00:00:00', 'https://www.facebook.com/share/p/16Vv3fLe9G/', 'Ponce ', '2025-05-17_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-505, Km 11.1, Camino La Vega', 0),
(437, NULL, '2025-05-18 00:00:00', 'https://www.facebook.com/share/p/18ySUGpLMU/', 'Comerio', '2025-05-18_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156, La Guitarra', 0),
(438, NULL, '2025-05-18 00:00:00', 'https://www.facebook.com/share/p/1J4V3MzHoa/', 'Morovis', '2025-05-18_2', NULL, NULL, NULL, NULL, NULL, 'PR-155, Barrio Perchas, Sector Bavaria', 0),
(439, NULL, '2025-05-18 00:00:00', 'https://www.facebook.com/share/p/16jbDKGv5d/', 'Morovis', '2025-05-18_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-155, Barrio Perchas, Sector Bavaria. En los mismos comentarios los ciudadanos muestran fotos de otros deslizamientos pero no se menciona fecha o area impactada.', 0),
(440, NULL, '2025-05-19 00:00:00', 'https://www.facebook.com/share/p/1CYE3qFjAV/', 'Rio Grande', '2025-05-19_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-9966, Km 7.0, escombros de derrumbe en el pavimento', 0),
(441, NULL, '2025-05-19 00:00:00', 'https://www.facebook.com/share/p/1CYE3qFjAV/', 'Rio Grande', '2025-05-19_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-9966, Km 4.3, Barrio Jimenez, carretera socavada y hundida 18 pulgadas (frente a varias residencias), Se reportan el dia 19 pero parece haber ocurrido antes.', 0),
(442, NULL, '2025-05-20 00:00:00', 'https://www.facebook.com/share/p/12Jp5NFPj88/', 'Utuado', '2025-05-20_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-111, Km 57.3, Barrio Caguana', 0),
(443, NULL, '2025-05-21 00:00:00', 'https://www.facebook.com/share/p/19SqZgwLuo/', 'Comerio', '2025-05-21_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-156, Km 39.3, Barrio Naranjo', 0),
(444, NULL, '2025-05-21 00:00:00', 'https://www.facebook.com/share/p/1Axf5bW4Uq/', 'Comerio', '2025-05-21_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-791, Km 1.5, Barrio El Verde', 0),
(445, NULL, '2025-05-22 00:00:00', 'https://www.facebook.com/share/p/16hB3ViBs1/', 'Comerio', '2025-05-22_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-167 Km 4.9, Barrio El Salto, un vehiculo quedo aplastado por el deslizamiento (Tacoma color azul)', 0),
(446, NULL, '2025-05-27 00:00:00', 'https://www.facebook.com/share/p/18wscHAbWQ/', 'San Sebastian', '2025-05-27_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-119, Barrio Guacio', 0),
(447, NULL, '2025-06-19 00:00:00', 'https://www.facebook.com/share/p/16dzMZf6o4/', 'Guayama', '2025-06-19_1', 18.047452, -66.096943, NULL, NULL, NULL, 'Carretera PR-179, Km 9.3, direccion a Carite, Lat: N 18,2,50.826 Lon: W 66,5,48.996', 0),
(448, NULL, '2025-07-28 00:00:00', 'https://www.facebook.com/share/p/1BdkwRFgKz/', 'Ciales', '2025-07-28_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-146, Km 16.6, Barrio frontón, Sector el hoyo, Es el comentario de una seccion de septiembre del alcalde de Ciales', 0),
(449, NULL, '2025-08-08 00:00:00', 'No hay enlace', 'Utuado', '2025-08-08_1', NULL, NULL, NULL, NULL, NULL, ' Barrio Arenas, 3:15pm con las fuertes lluvias', 0),
(450, NULL, '2025-08-11 00:00:00', 'No hay enlace', 'Moca', '2025-08-11_1', 18.382083, -67.092861, NULL, NULL, NULL, 'Carretera PR-111 Km 7.2, Cerca de gasolinera total, 18°22\'55.5\"N 67°05\'34.3\"W, Tania dijo no haber podido tomar fotos pero ella fue quien lo reporto', 0),
(451, NULL, '2025-08-11 00:00:00', 'No hay enlace', 'Utuado', '2025-08-11_2', NULL, NULL, NULL, NULL, NULL, 'Carretera 1PR-23, Km 65, Salto Abajo, Isabella envio las fotos al chat', 0),
(452, NULL, '2025-08-17 00:00:00', 'https://www.facebook.com/share/p/1BZjbhT61y/', 'San German', '2025-08-17_1', NULL, NULL, NULL, NULL, NULL, 'Carretera 119, cerca del Puente Cayo', 0),
(453, NULL, '2025-08-17 00:00:00', 'https://www.facebook.com/share/p/16nwpB8KFa/', 'San German', '2025-08-17_2', NULL, NULL, NULL, NULL, NULL, 'Esta fecha es incorrecta. En los comentarios los residentes dice que ese delizamiento lleva dias, posiblemnte mas de una semana de haber ocurrido. En la marginal de la entrada a Plaza del Oeste (entrada principal del pueblo) ', 0),
(454, NULL, '2025-08-31 00:00:00', 'https://www.facebook.com/share/p/1DE5uHHnjT/?mibextid=wwXIfr', 'Mayaguez', '2025-08-31_1', NULL, NULL, NULL, NULL, NULL, 'Camino Manantiales ', 0),
(455, NULL, '2025-09-04 00:00:00', 'https://www.facebook.com/share/p/1A6H7ThNzi/', 'San Sebastian', '2025-09-04_1', NULL, NULL, NULL, NULL, NULL, 'Sector Los Rosario en el barrio Cidral', 0),
(456, NULL, '2025-09-05 00:00:00', 'https://www.facebook.com/share/p/1Bcu2Tb9st/', 'Utuado', '2025-09-05_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-140, km 39.1, Barrio Mameyes', 0),
(457, NULL, '2025-09-06 00:00:00', 'https://www.facebook.com/share/p/16g3dvb78k/', 'Moca', '2025-09-06_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-4444, Km 1.0, Barrio Cuchillas, Sector Loperena', 0),
(458, NULL, '2025-09-06 00:00:00', 'https://www.facebook.com/share/p/1BeDvH9vZr/', 'Ciales', '2025-09-06_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-149,  justo en la planta de la luz en Jaguas', 0),
(459, NULL, '2025-09-06 00:00:00', 'https://www.facebook.com/share/p/19Q52zmp9p/', 'Utuado', '2025-09-06_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-140, km 39.1, barrio Don Alonso', 0),
(460, NULL, '2025-09-11 00:00:00', 'https://www.facebook.com/share/p/16ZVJWZDis/', 'Utuado', '2025-09-11_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-10, kilómetro 45.4, donde previamente había ocurrido un deslizamiento, nuevamente se ha registrado otro desprendimiento de terreno que continúa activo.', 0),
(461, NULL, '2025-09-12 00:00:00', 'https://www.facebook.com/share/p/17KbVBsDoU/', 'Utuado', '2025-09-12_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-111, km 55.2, \"Llegando al negocio de Papin despues de Jacana\"', 0),
(462, NULL, '2025-09-14 00:00:00', 'https://www.facebook.com/share/p/14LLomLShYY/', 'Utuado', '2025-09-14_1', NULL, NULL, NULL, NULL, NULL, 'La Altura, Parece haber ocurrido la noche antes y la limpieza comenzo el dia 14', 0),
(463, NULL, '2025-09-15 00:00:00', 'No hay enlace', 'Mayaguez', '2025-09-15_1', NULL, NULL, NULL, NULL, NULL, 'Enviado por Stephen al chat de la oficina y dije \"Eso fue la semana pasada en la urb quintas de santa maría en miradero.  Dentro del área con control aceso.  Compartido por el director de OMME mayaguez, David Rivera.   For the inventory, we can put it on the day of that week with the most rain.\"', 0),
(464, NULL, '2025-09-15 00:00:00', 'https://www.facebook.com/share/p/1GwWgtmNFf/', 'Mayaguez', '2025-09-15_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-106, Km 8.8, desprendimiento de carretera', 0),
(465, NULL, '2025-09-16 00:00:00', 'https://www.facebook.com/share/p/14LUHxjLvdM/', 'San German', '2025-09-16_1', NULL, NULL, NULL, NULL, NULL, 'PR-119, Km 70.1', 0),
(466, NULL, '2025-09-19 00:00:00', 'https://www.facebook.com/share/p/1DYbhtMEsp/', 'San German', '2025-09-19_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-119', 0),
(467, NULL, '2025-09-20 00:00:00', 'https://www.facebook.com/share/p/1FivfBmaDi/', 'San German', '2025-09-20_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-119, Km 70.7', 0),
(468, NULL, '2025-09-22 00:00:00', 'https://www.facebook.com/permalink.php?story_fbid=pfbid028aEhfRzXjpmXVaFD8bP9SFi9Dsqfzjnp1yiyqMj8sdxZAYSWDmLKqcaPw53ocVLWl&id=100064423912654', 'Utuado', '2025-09-22_1', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-613, Km 6.2, Barrio Tetuán', 0),
(469, NULL, '2025-09-24 00:00:00', 'https://www.facebook.com/share/p/17ayDmphUM/', 'Mayaguez', '2025-09-24_1', NULL, NULL, NULL, NULL, NULL, 'PR-106, Km 5.4', 0),
(470, NULL, '2025-09-24 00:00:00', 'https://www.facebook.com/share/p/1ERLar2mV5/', 'Trujillo Alto', '2025-09-24_2', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-851, Km 0.5, Barrio La Gloria', 0),
(471, NULL, '2025-09-24 00:00:00', 'https://www.facebook.com/share/v/19r9xu14Sn/', 'Coamo', '2025-09-24_3', NULL, NULL, NULL, NULL, NULL, 'Carretera PR-155, Km 5.9, Barrio Pasto', 0),
(472, NULL, '2025-09-25 00:00:00', 'https://www.facebook.com/share/p/1AQ5HzCwp2/', 'Guayanilla', '2025-09-25_1', NULL, NULL, NULL, NULL, NULL, 'Barrio Quebrada Honda, Sector Paganes, dejó sin salida a 5 familias residentes del lugar', 0),
(473, NULL, '2025-09-25 00:00:00', 'https://www.facebook.com/share/p/1CPPK2gh67/', 'Coamo', '2025-09-25_2', 18.150394, -66.404476, NULL, NULL, NULL, 'Carretera hacia Pedro García', 0),
(474, NULL, '2025-09-25 00:00:00', 'https://www.facebook.com/share/p/1QS5Q2f8vc/', 'Villalba', '2025-09-25_3', 18.152834, -66.450002, NULL, NULL, NULL, 'Carretera PR-151, Barrio El Limón', 0),
(475, NULL, '2025-10-01 00:00:00', 'No hay enlace', 'Sabana Grande', '2025-10-01_1', 18.110595, -66.925921, NULL, NULL, NULL, ', Stephen dijo no tener fecha exacta y que habia sido aproximadamente hace una semana desde que lo anuncion (7 de octubre lo dijo)', 0),
(494, NULL, '2025-12-11 00:00:00', 'bgyujnbgyuikjnbgyu', 'Aguada', '2025-12-11_494', 18.209847, -66.327422, 'Anonymous', '', '', '', 0),
(495, NULL, '2025-12-11 00:00:00', 'test', 'Adjuntas', '2025-12-11_495', 18.210311, -67.139887, 'Anonymous', '', '', '', 0),
(496, NULL, '2025-12-11 00:00:00', 'Test', 'Adjuntas', '2025-12-11_496', 18.210252, -67.139884, 'Jose', '7873828964', 'jose.irizarry24@upr.edu', 'Test', 0),
(497, NULL, '2025-12-11 00:00:00', 'Deslizamiento en la calle 123 del barrio Quenepas.', 'Aguadilla', '2025-12-11_497', 18.301239, -67.170166, 'Jose Rivera', '787 555 5555', 'capstone.derrumbes@gmail.com', '', 0),
(498, NULL, '2025-12-11 00:00:00', 'Deslizamiento en la calle 123 detrás del barrio Quenepas', 'Aguadilla', '2025-12-11_498', 18.457913, -67.109741, 'Jose Rivera', '787 555 5555', 'capstone.derrumbes@gmail.com', '', 0),
(499, NULL, '2025-12-11 00:00:00', 'Deslizamiento en la calle 123 detras del barrio Quenepas', 'Adjuntas', '2025-12-11_499', 18.209455, -67.139947, 'Jose Rivera', '787 555 5555', 'capstone.derrumbes@gmail.com', '', 0),
(500, NULL, '2025-12-12 00:00:00', 'Derrumbe en la calle 172.', 'Caguas', '2025-12-12_500', 18.210167, -67.139828, 'Glerysbeth Serrano', '787-505-8836', 'glerysbeth.serrano@upr.edu', '156 Villas del Bosque', 0),
(501, NULL, '2025-12-12 00:00:00', 'Deslizamiento en la carretera 172.', 'Caguas', '2025-12-12_501', 18.210427, -67.139774, 'Glerysbeth Serrano', '7875058836', 'glerysserrano@gmail.com', 'PR-172', 0),
(502, NULL, '2025-12-12 00:00:00', 'Deslizamiento en carretera 172.', 'Caguas', '2025-12-12_502', 18.210243, -67.139818, 'Anonymous', '', '', 'PR-172', 0);

-- --------------------------------------------------------

--
-- Table structure for table `station_info`
--

CREATE TABLE `station_info` (
  `station_id` int(11) NOT NULL,
  `admin_id` int(11) NOT NULL,
  `soil_saturation` int(10) UNSIGNED NOT NULL,
  `precipitation` decimal(6,2) DEFAULT NULL,
  `sensor_image_url` varchar(512) DEFAULT NULL,
  `data_image_url` varchar(512) DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  `is_available` tinyint(1) NOT NULL DEFAULT 0,
  `last_updated` datetime DEFAULT NULL,
  `latitude` decimal(9,6) DEFAULT NULL,
  `longitude` decimal(9,6) DEFAULT NULL,
  `ftp_file_path` text DEFAULT NULL,
  `wc1` decimal(9,6) DEFAULT NULL,
  `wc2` decimal(9,6) DEFAULT NULL,
  `wc3` decimal(9,6) DEFAULT NULL,
  `wc4` decimal(9,6) DEFAULT NULL,
  `land_unit` varchar(50) DEFAULT NULL,
  `geological_unit` varchar(50) DEFAULT NULL,
  `slope` decimal(10,2) DEFAULT NULL,
  `elevation` int(11) DEFAULT NULL,
  `susceptibility` varchar(50) DEFAULT NULL,
  `depth` varchar(50) DEFAULT NULL,
  `history_data_url` varchar(512) DEFAULT NULL,
  `station_installation_date` datetime DEFAULT NULL,
  `collaborator` varchar(50) DEFAULT NULL,
  `landslide_forecast` decimal(6,2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `station_info`
--

INSERT INTO `station_info` (`station_id`, `admin_id`, `soil_saturation`, `precipitation`, `sensor_image_url`, `data_image_url`, `city`, `is_available`, `last_updated`, `latitude`, `longitude`, `ftp_file_path`, `wc1`, `wc2`, `wc3`, `wc4`, `land_unit`, `geological_unit`, `slope`, `elevation`, `susceptibility`, `depth`, `history_data_url`, `station_installation_date`, `collaborator`, `landslide_forecast`) VALUES
(2, 1, 36, 0.00, 'stations/mayaguez.jpg', 'network/plots/Mayaguez.jpeg', 'MAYA', 1, '2026-01-29 14:12:30', 18.220011, -67.144000, 'network/data/latest/mayaguez_t60min.dat', 0.457000, 0.472000, 0.450000, 0.512000, 'Arcilla mucara', 'Formación Maricao', 40.00, 33, 'Muy alta', '20 cm, 40 cm, 60 cm, 80 cm', 'stations_history/mayaguez_merged_60min.csv', '2022-03-01 00:00:00', 'UPR Mayagüez', 37.73),
(3, 1, 85, 0.00, 'stations/adjuntas.jpg', 'network/plots/Adjuntas.jpeg', 'ADJU', 1, '2026-01-29 14:12:30', 18.147066, -66.765009, 'network/data/latest/adjuntas_t60min.dat', 0.531000, 0.499000, 0.526000, 0.522000, 'Arcilla Los Guineos', 'Formación Yauco', 45.00, 1020, 'Muy alta', '20 cm, 40 cm, 60 cm, 80 cm', 'stations_history/adjuntas_merged_60min.csv', '2022-03-01 00:00:00', 'Departamento de Recursos Naturales y Ambientales', 18.55),
(4, 1, 81, 0.00, 'stations/aguada.jpg', 'network/plots/Aguada.jpeg', 'AGUA', 1, '2026-01-29 14:12:30', 18.317900, -67.163900, 'network/data/latest/aguada_t60min.dat', 0.498000, 0.544000, 0.502000, 0.503000, NULL, NULL, NULL, 189, NULL, '25 cm, 50 cm, 75 cm, 100 cm', 'stations_history/aguada_merged_60min.csv', '2025-05-01 00:00:00', NULL, 29.11),
(5, 1, 78, 0.00, 'stations/anasco.jpg', 'network/plots/Anasco.jpeg', 'AÑAS', 1, '2026-01-29 14:12:30', 18.294006, -67.051014, 'network/data/latest/anasco_t60min.dat', 0.510000, 0.496000, 0.482000, 0.502000, 'Arcilla de Consumo', 'Formación Río Blanco', 35.00, 185, 'Muy alta', '25 cm, 50 cm, 75 cm, 100 cm', 'stations_history/anasco_merged_60min.csv', '2023-09-01 00:00:00', 'Propietario privado', 30.54),
(6, 1, 24, 0.00, 'stations/barranquitas.jpg', 'network/plots/Barranquitas.jpeg', 'BARR', 1, '2026-01-29 14:12:30', 18.166013, -66.299999, 'network/data/latest/barranquitas_t60min.dat', 0.318000, 0.341000, 0.283000, 0.263000, 'Franco arcilloso Caguabo', 'Formación Robles', 45.00, 540, 'Muy alta', '25 cm, 53 cm, 82 cm, 100 cm', 'stations_history/barranquitas_merged_60min.csv', '2022-02-01 00:00:00', 'Para la Naturaleza', 0.00),
(7, 1, 69, 0.00, 'stations/cayey.jpg', 'network/plots/Cayey.jpeg', 'CAYE', 1, '2026-01-29 14:12:30', 18.109066, -66.150998, 'network/data/latest/cayey_t60min.dat', 0.462000, 0.491000, 0.498000, 0.492000, 'Arcilla mucara', 'Formación A', 35.00, 480, 'Alta', '21 cm, 42 cm, 63 cm, 84 cm', 'stations_history/cayey_merged_60min.csv', '2021-12-01 00:00:00', 'Observatorio geomagnético del USGS', 54.06),
(8, 1, 61, 0.00, 'stations/ciales.jpg', 'network/plots/Ciales.jpeg', 'CIAL', 1, '2026-01-29 14:12:30', 18.312029, -66.468999, 'network/data/latest/ciales_t60min.dat', 0.489000, 0.463000, 0.499000, 0.550000, 'Arcilla Mucara', 'Formación Río Orocovis, miembro de lava Avispa', 30.00, 320, 'Muy alta', '22,5 cm, 45 cm, 67,5 cm, 90 cm', 'stations_history/ciales_merged_60min.csv', '2022-02-01 00:00:00', 'Propietario privado', 70.33),
(9, 1, 51, 0.00, 'stations/lares.jpg', 'network/plots/Lares.jpeg', 'LARE', 1, '2026-01-29 14:12:30', 18.248041, -66.880999, 'network/data/latest/lares_t60min.dat', 0.495000, 0.511000, 0.502000, 0.494000, 'Franco arcilloso Morado', 'Miembro de brecha de la Formación Milagros', 35.00, 250, 'Muy alta', '22 cm, 44 cm, 66 cm, 88 cm', 'stations_history/lares_merged_60min.csv', '2022-03-01 00:00:00', 'Propietario privado', 25.37),
(10, 1, 65, 0.20, 'stations/maricao.jpg', 'network/plots/Maricao.jpeg', 'MARI', 1, '2026-01-29 14:12:30', 18.173651, -67.031225, 'network/data/latest/maricao_t60min.dat', 0.495000, 0.570000, 0.471000, 0.447000, 'Arcilla mucara', 'Diorita porfídica de hornblenda', 40.00, 400, 'Muy alta', '33 cm, 66 cm, 99 cm, 130 cm', 'stations_history/maricao_merged_60min.csv', '2022-12-01 00:00:00', 'Propietario privado', 30.39),
(11, 1, 52, 0.20, 'stations/maunabo.jpg', 'network/plots/Maunabo.jpeg', 'MAUN', 1, '2026-01-29 14:12:30', 18.035026, -65.910009, 'network/data/latest/maunabo_t60min.dat', 0.384000, 0.335000, 0.374000, 0.393000, 'Complejo de tierras muy pedregosas de Pandura', 'Granodiorita de San Lorenzo', 35.00, 350, 'Muy alta', '27,5 cm, 55 cm, 77,5 cm, 110 cm', 'stations_history/maunabo_merged_60min.csv', '2022-02-01 00:00:00', 'Para la Naturaleza', 27.14),
(12, 1, 74, 0.00, 'stations/naguabo.jpg', 'network/plots/Naguabo.jpeg', 'NAGU', 1, '2026-01-29 14:12:30', 18.255015, -65.794002, 'network/data/latest/naguabo_t60min.dat', 0.455000, 0.423000, 0.445000, 0.439000, 'Marga Pandura', 'Diorita de cuarzo de Río Blanco', 40.00, 370, 'Muy alta', '25 cm, 50 cm, 75 cm, 100 cm', 'stations_history/naguabo_merged_60min.csv', '2022-07-01 00:00:00', 'Propietario privado', 25.37),
(13, 1, 42, 0.00, 'stations/naranjito.jpg', 'network/plots/Naranjito.jpeg', 'NARA', 1, '2026-01-29 14:12:30', 18.296230, -66.246019, 'network/data/latest/naranjito_t60min.dat', 0.511000, 0.498000, 0.478000, 0.462000, 'Arcilla mucara', 'Grupo Río Orocovis, Formación Los Negros', 45.00, 300, 'Muy alta', '25 cm, 50 cm, 75 cm, 100 cm', 'stations_history/naranjito_merged_60min.csv', '2022-08-01 00:00:00', 'Municipio de Naranjito', 1.34),
(14, 1, 78, 0.25, 'stations/orocovis.jpg', 'network/plots/Orocovis.jpeg', 'OROC', 1, '2026-01-29 14:12:30', 18.177009, -66.415000, 'network/data/latest/orocovis_t60min.dat', 0.468000, 0.480000, 0.481000, 0.479000, 'A determinar', 'A determinar', NULL, NULL, 'A determinar', '25 cm, 50 cm, 75 cm, 100 cm', 'stations_history/orocovis_merged_60min.csv', '2024-04-01 00:00:00', 'colaborador privado', 91.79),
(15, 1, 58, 0.00, 'stations/ponce.jpg', 'network/plots/Ponce.jpeg', 'PONC', 1, '2026-01-29 14:12:30', 18.083024, -66.659999, 'network/data/latest/ponce_t60min.dat', 0.409000, 0.434000, 0.425000, 0.440000, 'Franco arcilloso con grava de Caguabo', 'Formación Lago Garzas', 35.00, 330, 'Muy alta', '25 cm, 50 cm, 75 cm, 100 cm', 'stations_history/ponce_merged_60min.csv', '2023-03-01 00:00:00', 'Para la Naturaleza', 25.37),
(16, 1, 97, 0.00, 'stations/sanlorenzo.jpg', 'network/plots/San_Lorenzo.jpeg', 'SANL', 1, '2026-01-29 14:12:30', 18.089038, -66.002001, 'network/data/latest/sanlorenzo_t60min.dat', 0.515000, 0.499000, 0.485000, 0.470000, 'Arcilla Los Guineos', 'Formación A', 35.00, 420, 'Alta', '17,5 cm, 35 cm, 52,5 cm, 70 cm', 'stations_history/sanlorenzo_merged_60min.csv', '2022-02-01 00:00:00', 'Para la Naturaleza', 22.80),
(17, 1, 95, 0.00, 'stations/toronegro.jpg', 'network/plots/Toro_Negro.jpeg', 'TORO', 1, '2026-01-29 14:12:30', 18.158017, -66.565002, 'network/data/latest/toronegro_t60min.dat', 0.544000, 0.535000, 0.524000, 0.536000, 'Asociación Los Guineos - Maricao', 'Conglomerado de achiote', 45.00, 1200, 'Muy alta', '30 cm, 50 cm, 90 cm, 100 cm', 'stations_history/toronegro_merged_60min.csv', '2018-01-01 00:00:00', 'Departamento de Recursos Naturales y Ambientales', 22.80),
(18, 1, 25, 0.00, 'stations/utuado.jpg', 'network/plots/Utuado.jpeg', 'UTUA', 1, '2026-01-29 14:12:30', 18.280025, -66.661001, 'network/data/latest/utuado_t60min.dat', 0.427000, 0.420000, 0.450000, 0.453000, 'Franco arcilloso Pellejas', 'Granodiorita de Utuado', 40.00, 500, 'Muy alta', '27 cm, 42 cm, 57 cm, 72 cm', 'stations_history/utuado_merged_60min.csv', '2018-01-01 00:00:00', 'Propietario privado', 1.34),
(19, 1, 57, 1.02, 'stations/yabucoa.jpg', 'network/plots/Yabucoa.jpeg', 'YABU', 1, '2026-01-29 14:12:30', 18.110012, -65.855998, 'network/data/latest/Yabucoa_t60min.dat', 0.423000, 0.379000, 0.397000, 0.499000, 'Marga Pandura', 'Granodiorita de San Lorenzo', NULL, 260, 'Muy alta', '25 cm, 50 cm, 75 cm, 100 cm', 'stations_history/yabucoa_merged_60min.csv', '2023-12-01 00:00:00', 'Propietario privado', 55.94),
(20, 1, 51, 0.00, 'stations/yauco.jpg', 'network/plots/Yauco.jpeg', 'YAUC', 1, '2026-01-29 14:12:30', 18.138996, -66.881000, 'network/data/latest/yauco_t60min.dat', 0.463000, 0.423000, 0.400000, 0.510000, 'Arcilla de Maricao', 'Lodolita de Yauco', 45.00, 770, 'Muy alta', '25 cm, 50 cm, 75 cm, 100 cm', 'stations_history/yauco_merged_60min.csv', '2023-01-01 00:00:00', 'Propietario privado', 14.37),
(21, 1, 0, 0.00, NULL, NULL, 'Carolina', 1, '2026-01-29 14:12:30', 18.290000, -65.922000, 'network/data/latest/carolina_t60min.dat', 0.000000, 0.000000, 0.000000, 0.000000, NULL, NULL, NULL, NULL, NULL, NULL, 'stations_history/carolina_merged_60min.csv', NULL, NULL, 1.52);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`admin_id`);

--
-- Indexes for table `landslide`
--
ALTER TABLE `landslide`
  ADD PRIMARY KEY (`landslide_id`),
  ADD KEY `admin_id` (`admin_id`);

--
-- Indexes for table `project`
--
ALTER TABLE `project`
  ADD PRIMARY KEY (`project_id`),
  ADD KEY `admin_id` (`admin_id`);

--
-- Indexes for table `publication`
--
ALTER TABLE `publication`
  ADD PRIMARY KEY (`publication_id`),
  ADD KEY `admin_id` (`admin_id`);

--
-- Indexes for table `report`
--
ALTER TABLE `report`
  ADD PRIMARY KEY (`report_id`),
  ADD KEY `idx_report_landslide` (`landslide_id`);

--
-- Indexes for table `station_info`
--
ALTER TABLE `station_info`
  ADD PRIMARY KEY (`station_id`),
  ADD KEY `admin_id` (`admin_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `admin_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- AUTO_INCREMENT for table `landslide`
--
ALTER TABLE `landslide`
  MODIFY `landslide_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=202;

--
-- AUTO_INCREMENT for table `project`
--
ALTER TABLE `project`
  MODIFY `project_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `publication`
--
ALTER TABLE `publication`
  MODIFY `publication_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- AUTO_INCREMENT for table `report`
--
ALTER TABLE `report`
  MODIFY `report_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=503;

--
-- AUTO_INCREMENT for table `station_info`
--
ALTER TABLE `station_info`
  MODIFY `station_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `landslide`
--
ALTER TABLE `landslide`
  ADD CONSTRAINT `landslide_ibfk_1` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`admin_id`);

--
-- Constraints for table `project`
--
ALTER TABLE `project`
  ADD CONSTRAINT `project_ibfk_1` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`admin_id`);

--
-- Constraints for table `publication`
--
ALTER TABLE `publication`
  ADD CONSTRAINT `publication_ibfk_1` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`admin_id`);

--
-- Constraints for table `report`
--
ALTER TABLE `report`
  ADD CONSTRAINT `report_ibfk_1` FOREIGN KEY (`landslide_id`) REFERENCES `landslide` (`landslide_id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `station_info`
--
ALTER TABLE `station_info`
  ADD CONSTRAINT `station_info_ibfk_1` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`admin_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
