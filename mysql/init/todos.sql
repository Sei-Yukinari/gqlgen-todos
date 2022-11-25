CREATE TABLE todos (
                       `id`    INT(11) UNSIGNED PRIMARY KEY,
                       `text`    TEXT    NOT NULL,
                       done BOOLEAN DEFAULT FALSE
)