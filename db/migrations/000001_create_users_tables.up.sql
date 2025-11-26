CREATE TABLE `users` (
                         `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
                         `name` longtext,
                         `no_telp` varchar(191) UNIQUE,
                         `email` varchar(191) UNIQUE,
                         `password` longtext,
                         `tanggal_lahir` datetime(3) NULL,
                         `jenis_kelamin` longtext,
                         `tentang` longtext,
                         `pekerjaan` longtext,
                         `id_provinsi` longtext,
                         `id_kota` longtext,
                         `is_admin` boolean DEFAULT false,
                         `created_at` datetime(3) NULL,
                         `updated_at` datetime(3) NULL
);