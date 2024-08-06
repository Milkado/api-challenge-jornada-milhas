-- Create "destinies" table
CREATE TABLE `destinies` (`id` bigint NOT NULL AUTO_INCREMENT, `name` longtext NOT NULL, `picture` varchar(255) NOT NULL, `price` double NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `picture` (`picture`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
