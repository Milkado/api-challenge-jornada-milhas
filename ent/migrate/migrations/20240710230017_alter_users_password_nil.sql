-- Modify "users" table
ALTER TABLE `users` MODIFY COLUMN `password` varchar(255) NULL, MODIFY COLUMN `rand_security` varchar(255) NULL;
