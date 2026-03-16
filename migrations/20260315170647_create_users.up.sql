CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_passwords varchar not null
)