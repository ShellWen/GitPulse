CREATE TABLE `repo` (
                        `data_id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `data_create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `data_update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `id` bigint unsigned NOT NULL DEFAULT '0',
                        `name` varchar(255) NOT NULL DEFAULT '',
                        `gist` boolean NOT NULL DEFAULT '0',
                        `star_count` bigint NOT NULL DEFAULT '0',
                        `fork_count` bigint NOT NULL DEFAULT '0',
                        `issue_count` bigint NOT NULL DEFAULT '0',
                        `commit_count` bigint NOT NULL DEFAULT '0',
                        `pr_count` bigint NOT NULL DEFAULT '0',
                        `language` json NOT NULL,
                        `description` varchar(255) NOT NULL DEFAULT '',
                        `readme` longtext NOT NULL,
                        PRIMARY KEY (`data_id`),
                        UNIQUE KEY `repo_pk` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

