-- Modify "testimonies" table
ALTER TABLE `testimonies` ADD COLUMN `destiny_id` bigint NOT NULL, ADD INDEX `testimonies_destinies_testimonies` (`destiny_id`), ADD CONSTRAINT `testimonies_destinies_testimonies` FOREIGN KEY (`destiny_id`) REFERENCES `destinies` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
