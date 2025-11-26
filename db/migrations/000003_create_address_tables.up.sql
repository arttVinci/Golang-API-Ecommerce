CREATE TABLE `address` (
                           `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
                           `id_user` bigint unsigned,
                           `judul_alamat` longtext,
                           `nama_penerima` longtext,
                           `no_telp` longtext,
                           `detail_alamat` longtext,
                           `created_at` datetime(3) NULL,
                           `updated_at` datetime(3) NULL,
                           CONSTRAINT `fk_users_alamat` FOREIGN KEY (`id_user`) REFERENCES `users`(`id`) ON DELETE CASCADE
);