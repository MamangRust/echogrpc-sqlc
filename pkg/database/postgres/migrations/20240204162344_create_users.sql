-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users"(
    "id" serial PRIMARY KEY,
    "firstname" VARCHAR(100) NOT NULL,
    "lastname" varchar(100) NOT NULL,
    "email" varchar(100) UNIQUE NOT NULL,
    "password" varchar(100) NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXITS "users";
-- +goose StatementEnd
