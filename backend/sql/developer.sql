CREATE TABLE `developer` (
                             `data_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'Generated primary key, MUST NOT be changed.',
                             `data_create_at` timestamp NOT NULL DEFAULT (now()),
                             `data_update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'Unique id of GitHub user.',
                             `name` varchar(255) NOT NULL DEFAULT '',
                             `username` varchar(255) NOT NULL DEFAULT '',
                             `avatar_url` varchar(255) NOT NULL DEFAULT '',
                             `company` varchar(255) NOT NULL DEFAULT '',
                             `location` varchar(255) NOT NULL DEFAULT '',
                             `bio` varchar(255) NOT NULL DEFAULT '',
                             `blog` varchar(255) NOT NULL DEFAULT '',
                             `email` varchar(255) NOT NULL DEFAULT '',
                             `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             PRIMARY KEY (`data_id`),
                             UNIQUE KEY `developer_pk_2` (`id`),
                             UNIQUE KEY `developer_pk_3` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

