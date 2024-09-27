CREATE DATABASE IF NOT EXISTS `monitoring`;
USE `monitoring`;

CREATE TABLE IF NOT EXISTS `random_data`(
    `id` bigint(20) primary key auto_increment,
    `x` decimal(18,2) not null,
    `y` decimal(18,2) not null,
    `created` datetime(6) not null default current_timestamp
)Engine=InnoDB;
