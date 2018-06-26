-- phpMyAdmin SQL Dump
-- version 4.5.1
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Jun 26, 2018 at 11:51 AM
-- Server version: 10.1.10-MariaDB
-- PHP Version: 7.0.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `appe-surat`
--

-- --------------------------------------------------------

--
-- Table structure for table `data_disposisi`
--

CREATE TABLE `data_disposisi` (
  `id_disposisi` int(11) NOT NULL,
  `no_surat` varchar(100) NOT NULL,
  `surat_dari` varchar(255) NOT NULL,
  `tgl_surat` date NOT NULL,
  `tgl_diterima` datetime NOT NULL,
  `no_agenda` varchar(100) NOT NULL,
  `sifat` char(50) NOT NULL,
  `perihal` longtext NOT NULL,
  `diteruskan_kpd` varchar(100) NOT NULL,
  `dengan_hormat_hrp` varchar(100) NOT NULL,
  `catatan` longtext NOT NULL,
  `tembusan01` char(100) NOT NULL,
  `tembusan02` char(100) NOT NULL,
  `tembusan03` char(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `data_disposisi`
--

INSERT INTO `data_disposisi` (`id_disposisi`, `no_surat`, `surat_dari`, `tgl_surat`, `tgl_diterima`, `no_agenda`, `sifat`, `perihal`, `diteruskan_kpd`, `dengan_hormat_hrp`, `catatan`, `tembusan01`, `tembusan02`, `tembusan03`) VALUES
(2, '002H', '3', '0000-00-00', '2018-05-11 12:22:34', '3/3', 'Biasa', '3', 'Kasubbag. APK', '3', '3', '3', '3', '3'),
(3, '003H', '2', '2014-04-04', '2018-05-11 14:41:31', '2/2', 'Biasa', '2', '2', 'Tanggapan dan Saran', '2', '2', '2', '2'),
(4, '004HI', 'Dinas Komunikasi dan Informatika', '2010-10-31', '2018-05-11 14:39:03', '1/1', 'Biasa', 'Pembuatan aplikasi surat keluar masuk secara elektronik', 'Kasubag AP', 'Kepala dinas', 'Telah di ferivikasi', 'Presiden', 'Bupati', 'Sekda'),
(5, '005I/003/F9', 'BRI', '2018-06-01', '2018-06-10 20:58:29', '12/09A', 'Sangat Segera', 'Uji Coba', 'Kasubbag. APK', 'Tanggapan dan Saran', 'Tagihan', 'a', 'b', 'c');

-- --------------------------------------------------------

--
-- Table structure for table `data_users`
--

CREATE TABLE `data_users` (
  `id` int(11) NOT NULL,
  `username` varchar(150) NOT NULL,
  `password` varchar(255) NOT NULL,
  `level` varchar(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `data_users`
--

INSERT INTO `data_users` (`id`, `username`, `password`, `level`) VALUES
(1, 'admin', '21232f297a57a5a743894a0e4a801fc3', 'admin'),
(2, 'hendri', 'f68bac89085669468822b54a74b9b93e', 'kepala');

-- --------------------------------------------------------

--
-- Table structure for table `surat_keluar`
--

CREATE TABLE `surat_keluar` (
  `id_keluar` int(11) NOT NULL,
  `no_surat` varchar(100) NOT NULL,
  `lampiran` varchar(255) NOT NULL,
  `tujuan` varchar(255) NOT NULL,
  `alamat_tujuan` varchar(255) NOT NULL,
  `tgl_surat_keluar` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `perihal` varchar(255) NOT NULL,
  `isi` longtext NOT NULL,
  `file` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `surat_keluar`
--

INSERT INTO `surat_keluar` (`id_keluar`, `no_surat`, `lampiran`, `tujuan`, `alamat_tujuan`, `tgl_surat_keluar`, `perihal`, `isi`, `file`) VALUES
(1, '002H', 'Pembersihan Rutin', 'Pembersihan Sungai', '', '2018-05-11 12:22:34', 'Perintah Pembersihan Sungai Agus Salin', '2.jpg', ''),
(3, '004H', 'TPA', 'Pembuatan TPA Baru', '', '2018-05-11 14:38:43', 'Pembuatan TPA Baru di Tanjung Serdang', '4.jpg', ''),
(4, '005H', 'TPA', 'Pembuatan TPA Baru', '', '2018-05-11 14:39:03', 'Pembuatan TPA Baru di Tanjung Serdang', '5.jpg', 'dbd22b05d4015b7edbc23c014fe04907.png'),
(5, '006H', 'TPA', 'Pembuatan TPA Baru', '', '2018-05-11 14:39:59', 'Pembuatan TPA Baru di Tanjung Serdang', '3.jpg', 'cd0929977f4ecb01c9ca749aa95e9d5c.jpg'),
(8, '99/PRP/2017', '-', 'Yth. Direktur PT. Janaka Groups', 'Jl. Mega Kenangan Km. 9 Jakarta Pusat', '2018-05-11 12:20:04', 'Permintaan Penawaran', 'Dengan Hormat,', '8ca40f5f515a91e95b660865c7e93295.jpg'),
(9, 'BH01', '-', 'Kepala RT Kecamatan Pulau laut utara', 'alamat masing masing', '2018-06-13 22:49:24', 'Bersih RT dan warga', 'Salam Hormat,<div>Berhubungan dengan akan di adakan nya pemilihan umum pada tahun 2019, maka Dinas lingkungan hidup sangat menghimbau kepada seluruh masyarakat kotabaru khususnya kecamatan pulau laut utara</div>', '4ba2aa75a32e76c4f8434e15adba9b4e.jpg'),
(10, '005I/003/F90', '-', 'contoh', 'contoh', '2018-06-22 15:31:08', 'contoh', 'contoh<br>', '');

-- --------------------------------------------------------

--
-- Table structure for table `surat_masuk`
--

CREATE TABLE `surat_masuk` (
  `id_masuk` int(11) NOT NULL,
  `no_surat` varchar(100) NOT NULL,
  `judul_surat` varchar(255) NOT NULL,
  `nama_surat` varchar(255) NOT NULL,
  `tgl_surat_masuk` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `perihal` varchar(500) NOT NULL,
  `file` varchar(500) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `surat_masuk`
--

INSERT INTO `surat_masuk` (`id_masuk`, `no_surat`, `judul_surat`, `nama_surat`, `tgl_surat_masuk`, `perihal`, `file`) VALUES
(2, '002H-01', 'Pembersihan Rutin', 'Pembersihan Sungai', '2018-05-11 12:22:34', 'Perintah Pembersihan Sungai Agus Salin', '5111dd876764cdc14e1e615677683381.png'),
(3, '003H/02', 'Penambahan Bak Sampah', 'Pembelian bak sampah baru', '2018-05-11 14:41:31', 'hal hal berhubungan sampah jalanan', '309f2b1cdfd4b3a5b23218fe79f70335.jpg'),
(5, '005I/003/F9', 'Tes Mobil Sampah', 'Sampah berat tiada masalah', '2018-06-10 20:58:29', 'penyedotan lumpur di setiap selokan', '1b9aaea36f27a18df103f7dab969efba.jpg');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `data_disposisi`
--
ALTER TABLE `data_disposisi`
  ADD PRIMARY KEY (`no_surat`);

--
-- Indexes for table `data_users`
--
ALTER TABLE `data_users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `surat_keluar`
--
ALTER TABLE `surat_keluar`
  ADD PRIMARY KEY (`id_keluar`);

--
-- Indexes for table `surat_masuk`
--
ALTER TABLE `surat_masuk`
  ADD PRIMARY KEY (`id_masuk`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `data_users`
--
ALTER TABLE `data_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `surat_keluar`
--
ALTER TABLE `surat_keluar`
  MODIFY `id_keluar` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
--
-- AUTO_INCREMENT for table `surat_masuk`
--
ALTER TABLE `surat_masuk`
  MODIFY `id_masuk` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
