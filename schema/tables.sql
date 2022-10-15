CREATE TABLE IF NOT EXISTS `orders`
(
    `id`         BIGINT UNSIGNED   NOT NULL COMMENT '注文を一意に識別するID',
    `status`     INT UNSIGNED      NOT NULL COMMENT '注文ステータス',
    `contact`    VARCHAR(255)      NOT NULL COMMENT '連絡先情報',
    `version`    SMALLINT UNSIGNED NOT NULL COMMENT 'バージョン',
    `created_at` DATETIME          NOT NULL COMMENT '作成日時',
    `updated_at` DATETIME          NOT NULL COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `order_items`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `order_id`   BIGINT UNSIGNED NOT NULL COMMENT '注文商品が紐づく注文ID',
    `item_name`  VARCHAR(256)    NOT NULL COMMENT '注文時点の商品名',
    `price`      INT             NOT NULL COMMENT '注文時点の商品価格',
    `quantity`   INT             NOT NULL COMMENT '数量',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`order_id`) REFERENCES orders (`id`),
    INDEX `order_items_idx` (`order_id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `order_events`
(
    `id`         BIGINT UNSIGNED   NOT NULL AUTO_INCREMENT,
    `order_id`   BIGINT UNSIGNED   NOT NULL COMMENT '注文ID',
    `version`    SMALLINT UNSIGNED NOT NULL COMMENT 'イベントバージョン',
    `event_type` VARCHAR(30)       NOT NULL COMMENT 'イベント種別',
    `event_data` JSON              NOT NULL COMMENT '詳細',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    PRIMARY KEY (`id`),
    UNIQUE (`order_id`, `version`),
    INDEX `order_events_idx` (`order_id`)
) ENGINE = InnoDB;
