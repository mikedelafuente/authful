CREATE TABLE `users` (
 `id` VARCHAR(64) NOT NULL,
 `username` varchar(320) NOT NULL,
 `password` varchar(255) NOT NULL,
 `create_datetime` DATETIME  DEFAULT NULL,
 `update_datetime` DATETIME  DEFAULT NULL,
 PRIMARY KEY (`id`)
) 
