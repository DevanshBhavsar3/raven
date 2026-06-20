CREATE TABLE `conversations` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(30) NOT NULL,
    `user_id` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT NOW(),
    `updated_at` DATETIME NOT NULL DEFAULT NOW()
);

CREATE TABLE `messages` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `content` TEXT NOT NULL,
    `role` ENUM('USER', 'AGENT') NOT NULL,
    `conversation_id` BIGINT UNSIGNED,
    `created_at` DATETIME NOT NULL DEFAULT NOW(),
    `updated_at` DATETIME NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_conversation_id FOREIGN KEY (`conversation_id`) REFERENCES `conversations`(`id`)
);
