-- phpMyAdmin SQL Dump
-- version 4.6.5.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 21, 2018 at 03:48 PM
-- Server version: 10.1.21-MariaDB
-- PHP Version: 5.6.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_karya`
--

-- --------------------------------------------------------

--
-- Table structure for table `karya`
--

CREATE TABLE `karya` (
  `id` int(12) NOT NULL,
  `judul` varchar(100) NOT NULL,
  `isi` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `karya`
--

INSERT INTO `karya` (`id`, `judul`, `isi`) VALUES
(1, 'ini judul karya pertama', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod\n	tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,\n	quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo\n	consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse\n	cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non\n	proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),
(2, 'ini judul karya kedua', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod\n	tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,\n	quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo\n	consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse\n	cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non\n	proident, sunt in culpa qui officia deserunt mollit anim id est laborum.');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `karya`
--
ALTER TABLE `karya`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `karya`
--
ALTER TABLE `karya`
  MODIFY `id` int(12) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
