CREATE TABLE `DB_Name`.`example`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
   `test_field` varchar(32)  NOT NULL DEFAULT '' COMMENT '测试字段',
   `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
   `modify_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
   `deleted_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
   `is_del` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除（1 删除 0 未删除）',
   PRIMARY KEY (`id`) USING BTREE,
   UNIQUE INDEX `uk_test_field`(`test_field`, `deleted_time`) USING BTREE COMMENT '索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1338758788165255169 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '测试表' ROW_FORMAT = Dynamic;

create table `example_two`  (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
    `test_field` varchar(32)  NOT NULL DEFAULT '' COMMENT '测试字段',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `modify_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `deleted_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    `is_del` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除（1 删除 0 未删除）',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `uk_test_field`(`test_field`, `deleted_time`) USING BTREE COMMENT '索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1338758788165255169 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '测试表' ROW_FORMAT = Dynamic;