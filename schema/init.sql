CREATE DATABASE gourly CHARACTER SET utf8 COLLATE utf8_general_ci;

USE gourly;

CREATE TABLE urls(
    id INT(16) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    redirect_to VARCHAR(128) NOT NULL,
    unique_id VARCHAR(32) NOT NULL,
    clicks INT(8) DEFAULT 0
);