CREATE TABLE `foto_products` (
                                `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
                                `produk_id` bigint unsigned,
                                `url` longtext,
                                `created_at` datetime(3) NULL,
                                `updated_at` datetime(3) NULL,
                                CONSTRAINT `fk_produks_foto_produk` FOREIGN KEY (`produk_id`) REFERENCES `products`(`id`) ON DELETE CASCADE
);