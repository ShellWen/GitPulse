CREATE TABLE `star` (
                        `data_id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `data_create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `data_update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `developer_id` bigint unsigned NOT NULL DEFAULT '0',
                        `repo_id` bigint unsigned NOT NULL DEFAULT '0',
                        PRIMARY KEY (`data_id`),
                        UNIQUE KEY `star_pk` (`developer_id`,`repo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci



CREATE TABLE `fork` (
                        `data_id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `data_create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `data_update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `original_repo_id` bigint unsigned NOT NULL DEFAULT '0',
                        `fork_repo_id` bigint unsigned NOT NULL DEFAULT '0',
                        PRIMARY KEY (`data_id`),
                        UNIQUE KEY `fork_pk_2` (`fork_repo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `follow` (
                          `data_id` bigint unsigned NOT NULL AUTO_INCREMENT,
                          `data_create_at` timestamp NOT NULL DEFAULT (now()),
                          `data_update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `following_id` bigint unsigned NOT NULL DEFAULT '0',
                          `followed_id` bigint unsigned NOT NULL DEFAULT '0',
                          PRIMARY KEY (`data_id`),
                          UNIQUE KEY `follow_pk` (`following_id`,`followed_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `create_repo` (
                               `data_id` bigint unsigned NOT NULL AUTO_INCREMENT,
                               `data_create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `data_update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               `developer_id` bigint unsigned NOT NULL DEFAULT '0',
                               `repo_id` bigint unsigned NOT NULL DEFAULT '0',
                               PRIMARY KEY (`data_id`),
                               UNIQUE KEY `create_repo_pk_2` (`repo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

