-- +goose Up
-- +goose StatementBegin
CREATE TABLE "comments" (
    "id" serial PRIMARY KEY,
    "id_post_comment" int NOT NULL,
    "user_name_comment" varchar(200) NOT NULL,
    "comment" varchar(200) NOT NULL,
    FOREIGN KEY (id_post_comment)
    REFERENCES posts(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXITS "comments";
-- +goose StatementEnd
