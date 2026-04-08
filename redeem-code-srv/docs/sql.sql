CREATE TABLE `user` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL DEFAULT '',
    `password` varchar(255) NOT NULL DEFAULT '',
    `nickname` varchar(255) NOT NULL DEFAULT '',
    `mobile` char(16) NOT NULL DEFAULT '',
    `gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '性别：0，未知；1，男；2，女；',
    `birthday` date DEFAULT NULL,
    `is_del` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '是否删除：0，默认；1，已删除；2，未删除；',
    `deleted_at` datetime DEFAULT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `redeem_code_batch` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT '批次名称',
    `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '批次描述',
    `usage_limit` int unsigned NOT NULL DEFAULT '1' COMMENT '可使用次数',
    `total_count` int unsigned NOT NULL DEFAULT '0' COMMENT '生成数量',
    `used_count` int unsigned NOT NULL DEFAULT '0' COMMENT '已使用数量',
    `started_at` datetime NOT NULL COMMENT '开始时间',
    `ended_at` datetime NOT NULL COMMENT '结束时间',
    `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态：0，默认；1，启用；2，停用；',
    `creator_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人 ID',
    `creator_name` varchar(255) NOT NULL DEFAULT '' COMMENT '创建人名称',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `redeem_code` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `redeem_code_batch_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '兑换码批次id',
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
    `value` varchar(255) NOT NULL DEFAULT '' COMMENT '兑换码',
    `usage_limit` int unsigned NOT NULL DEFAULT '1' COMMENT '可使用次数',
    `used_count` int unsigned NOT NULL DEFAULT '0' COMMENT '已使用数量',
    `expiration_at` datetime NOT NULL COMMENT '过期时间',
    `is_del` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '是否删除：0，默认；1，已删除；2，未删除；',
    `deleted_at` datetime DEFAULT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `redeem_code_record` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '用户id',
    `redeem_code_id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '兑换码id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
