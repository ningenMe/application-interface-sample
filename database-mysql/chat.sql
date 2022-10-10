CREATE DATABASE `sample`;
USE `sample`;

CREATE TABLE `chat`
(
    `chat_id` bigint unsigned AUTO_INCREMENT NOT NULL,
    `body` varchar(511) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`chat_id`)
) ENGINE = InnoDB;
