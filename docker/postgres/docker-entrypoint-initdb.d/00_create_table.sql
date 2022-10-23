--
-- テーブルの作成
--
CREATE TABLE user_profile
(
    user_id    VARCHAR(50) NOT NULL,
    user_name       VARCHAR(50) NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id)
);

CREATE TABLE todos
(
    todo_id    VARCHAR(50) NOT NULL,
    title       VARCHAR(50) NOT NULL,
    user_id    VARCHAR(50) NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES user_profile (user_id)
);
