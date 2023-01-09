CREATE TABLE IF NOT EXISTS todos(
    `ID` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
	`content` VARCHAR (50),
	`IsFinished` BOOLEAN,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    FOREIGN KEY (`user_id`) REFERENCES users(`ID`)
);
