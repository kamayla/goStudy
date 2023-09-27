CREATE TABLE `user`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name`     VARCHAR(20) NOT NULL COMMENT '名前',
    `password` VARCHAR(80) NOT NULL COMMENT 'パスワード',
    `role`     VARCHAR(80) NOT NULL COMMENT 'ロール',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード更新日',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `task`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'タスクの識別子',
    `user_id`  BIGINT UNSIGNED NOT NULL COMMENT 'タスクを作成したユーザーの識別子',
    `title`    VARCHAR(128) NOT NULL COMMENT '名前',
    `status`   VARCHAR(20)  NOT NULL COMMENT 'パスワード',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード更新日',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タスク';