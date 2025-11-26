CREATE TABLE `trxes` (
                         `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
                         `user_id` bigint unsigned,
                         `alamat_id` bigint unsigned,
                         `toko_id` bigint unsigned,
                         `harga_total` bigint,
                         `code_invoice` longtext,
                         `method_bayar` longtext,
                         `created_at` datetime(3) NULL,
                         `updated_at` datetime(3) NULL,
                         CONSTRAINT `fk_users_trx` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
                         CONSTRAINT `fk_alamats_trx` FOREIGN KEY (`alamat_id`) REFERENCES `address`(`id`),
                         CONSTRAINT `fk_tokos_trx` FOREIGN KEY (`toko_id`) REFERENCES `stores`(`id`)
);