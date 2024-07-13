-- Create "testimonies" table
CREATE TABLE `testimonies` (`id` bigint NOT NULL AUTO_INCREMENT, `testimony` longtext NOT NULL, `name` varchar(255) NOT NULL, `picture` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `picture` (`picture`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
