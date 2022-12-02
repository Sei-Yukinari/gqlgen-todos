-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `todos`
(
    `id`      int unsigned NOT NULL AUTO_INCREMENT,
    `text`    text         NOT NULL,
    `done`    tinyint(1)            DEFAULT '0',
    `user_id` int          NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 9
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `todos`;
