-- +goose Up
-- +goose StatementBegin
CREATE TABLE "categories" (
    "id" serial PRIMARY KEY,
    "name" varchar(200) NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXITS "categories";
-- +goose StatementEnd
