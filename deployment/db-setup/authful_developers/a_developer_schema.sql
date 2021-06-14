DROP TABLE IF EXISTS `developers`;
DROP TABLE IF EXISTS `applications`;
DROP TABLE IF EXISTS `credentials`;

CREATE TABLE `developers` (
 `dev_id` VARCHAR(64) NOT NULL,
 `user_id` VARCHAR(64) NOT NULL,
 `organization_name` varchar(320) NOT NULL,
 `contact_email` varchar(255) NOT NULL,
 `agree_to_tos` BIT NOT NULL,
 `create_datetime` DATETIME  DEFAULT NULL,
 `update_datetime` DATETIME  DEFAULT NULL,
 PRIMARY KEY (`dev_id`)
);


CREATE TABLE `applications` (
 `app_id` VARCHAR(64) NOT NULL,
 `dev_id` VARCHAR(64) NOT NULL,
 `name` varchar(320) NOT NULL,
 `url` varchar(255) NOT NULL,
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

