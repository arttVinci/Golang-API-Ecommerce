CREATE TABLE `products` (
                           `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
                           `category_id` bigint unsigned,
                           `toko_id` bigint unsigned,
                           `nama_produk` longtext,
                           `slug` longtext,
                           `harga_reseller` longtext,
                           `harga_konsumen` longtext,
                           `stok` bigint,
                           `deskripsi` longtext,
                           `created_at` datetime(3) NULL,
                           `updated_at` datetime(3) NULL,
                           CONSTRAINT `fk_categories_produk` FOREIGN KEY (`category_id`) REFERENCES `category`(`id`),
                           CONSTRAINT `fk_tokos_produk` FOREIGN KEY (`toko_id`) REFERENCES `stores`(`id`) ON DELETE CASCADE
);