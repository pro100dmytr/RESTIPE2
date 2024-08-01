-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE IF NOT EXISTS users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);

CREATE TABLE IF NOT EXISTS lists_items
(
    id      serial                                           not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS lists_items;
DROP TABLE IF EXISTS todo_items;
DROP TABLE IF EXISTS users_lists;
DROP TABLE IF EXISTS todo_lists;
DROP TABLE IF EXISTS users;