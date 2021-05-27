CREATE TABLE `developers` (
 `dev_id` VARCHAR(64) NOT NULL,
 `username` varchar(320) NOT NULL,
 `password` varchar(255) NOT NULL,
 `create_datetime` DATETIME  DEFAULT NULL,
 `update_datetime` DATETIME  DEFAULT NULL,
 PRIMARY KEY (`dev_id`)
);


CREATE TABLE `applications` (
 `app_id` VARCHAR(64) NOT NULL,
 `dev_id` VARCHAR(64) NOT NULL,
 `username` varchar(320) NOT NULL,
 `password` varchar(255) NOT NULL,
 `create_datetime` DATETIME  DEFAULT NULL,
 `update_datetime` DATETIME  DEFAULT NULL,
 PRIMARY KEY (`app_id`)
);

CREATE TABLE `credentials` (
 `cred_id` VARCHAR(64) NOT NULL,   
 `app_id` VARCHAR(64) NOT NULL,
 `client_id` VARCHAR(64) NOT NULL,
 `client_secret` varchar(64) NOT NULL,
 `create_datetime` DATETIME  DEFAULT NULL,
 `update_datetime` DATETIME  DEFAULT NULL,
 PRIMARY KEY (`cred_id`)
);

