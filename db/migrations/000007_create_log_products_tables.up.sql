CREATE TABLE `log_products` (
                               `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
                               `produk_id` bigint unsigned,
                               `category_id` bigint unsigned,
                               `toko_id` bigint unsigned,
                               `nama_produk` longtext,
                               `slug` longtext,
                               `harga_reseller` longtext,
                               `harga_konsumen` longtext,
                               `deskripsi` longtext,
                               `created_at` datetime(3) NULL,
                               `updated_at` datetime(3) NULL,
                               INDEX `idx_log_produk_id` (`produk_id`),
                               INDEX `idx_log_toko_id` (`toko_id`)
);