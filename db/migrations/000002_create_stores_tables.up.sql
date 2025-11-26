CREATE TABLE `stores` (
                         `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
                         `user_id` bigint unsigned UNIQUE,
                         `nama_toko` longtext,
                         `url_foto` longtext,
                         `created_at` datetime(3) NULL,
                         `updated_at` datetime(3) NULL,
                         CONSTRAINT `fk_users_toko` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);