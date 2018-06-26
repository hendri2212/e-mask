-- phpMyAdmin SQL Dump
-- version 4.5.1
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Jun 26, 2018 at 01:03 PM
-- Server version: 10.1.10-MariaDB
-- PHP Version: 7.0.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `e-mask`
--

-- --------------------------------------------------------

--
-- Table structure for table `tb_karya`
--

CREATE TABLE `tb_karya` (
  `Id_karya` int(11) NOT NULL,
  `Kategory` varchar(50) NOT NULL,
  `Judul` varchar(255) NOT NULL,
  `Deskripsi` varchar(500) NOT NULL,
  `Gambar` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tb_karya`
--

INSERT INTO `tb_karya` (`Id_karya`, `Kategory`, `Judul`, `Deskripsi`, `Gambar`) VALUES
(1, 'Puisi', 'Cinta Bertasbih', 'Mencintai seseorang karena tasbihnya', ''),
(2, 'Puisi', 'Mancing di malam hari', 'Pada suatu malam jum''at keliwon', ''),
(3, 'News', 'Hari Jadi Kotabaru', 'Pada bulan juni lalu telah dilaksanakan peringatan hari jadi kotabaru', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tb_karya`
--
ALTER TABLE `tb_karya`
  ADD PRIMARY KEY (`Id_karya`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tb_karya`
--
ALTER TABLE `tb_karya`
  MODIFY `Id_karya` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
