-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE todo_status AS ENUM ('DRAFT', 'IN_PROGRESS', 'CANCEL', 'DONE');
CREATE TABLE todos (
    id          uuid        DEFAULT uuid_generate_v4(),
    title       text        not null,
    description text        not null,
    status      todo_status not null,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION IF EXISTS "uuid-ossp";
DROP TYPE IF EXISTS todo_status;
DROP TABLE IF EXISTS todos;
-- +goose StatementEnd
