-- Modify "destinies" table
ALTER TABLE `destinies` DROP COLUMN `picture`, ADD COLUMN `meta` varchar(160) NOT NULL, ADD COLUMN `description` longtext NULL;
