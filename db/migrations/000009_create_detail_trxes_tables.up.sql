CREATE TABLE `detail_trxes` (
                                `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
                                `trx_id` bigint unsigned,
                                `toko_id` bigint unsigned,
                                `log_product_id` bigint unsigned,
                                `quantity` bigint,
                                `harga_total` bigint,
                                `created_at` datetime(3) NULL,
                                `updated_at` datetime(3) NULL,
                                CONSTRAINT `fk_trxes_detail_trx` FOREIGN KEY (`trx_id`) REFERENCES `trxes`(`id`) ON DELETE CASCADE,
                                CONSTRAINT `fk_tokos_detail_trx` FOREIGN KEY (`toko_id`) REFERENCES `stores`(`id`),
                                CONSTRAINT `fk_log_produks_detail_trx` FOREIGN KEY (`log_product_id`) REFERENCES `log_products`(`id`)
);