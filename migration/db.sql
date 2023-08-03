CREATE TABLE `tbl_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama_lengkap` varchar(90) DEFAULT NULL,
  `email` varchar(250) DEFAULT NULL,
  `username` varchar(60) DEFAULT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `status` varchar(32) DEFAULT 'OFF',
  `password` text,
  `cpassword` text,
  `created_at` datetime,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB ;

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext DEFAULT NULL,
  `occupation` varchar(255) DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `username` longtext DEFAULT NULL,
  `password_hash` varchar(255) DEFAULT NULL,
  `avatar_file_name` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `no_hp` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `mem_type` smallint(3) DEFAULT NULL,
  `pin` longtext DEFAULT NULL,
  `Photodiriktp` varchar(255) DEFAULT NULL,
  `Photoktp` varchar(255) DEFAULT NULL,
  `nik` varchar(60) DEFAULT NULL,
  `mother_name` varchar(60) DEFAULT NULL,
  `status` varchar(22) DEFAULT 'OFF',
  `balance` int(11) DEFAULT 0,
  `fcm_token` text DEFAULT NULL,
  `anggota_perkemi` varchar(2) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
