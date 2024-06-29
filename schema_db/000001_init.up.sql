CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    password_hash varchar(255) not null,
    email         varchar(255) not null unique
);

CREATE TABLE recipients
(
    id                serial              not null unique,
    name              varchar(255)        not null,
    surname           varchar(255)        not null,
    patronymic        varchar(255),
    email             varchar(255) unique not null,
    phone             varchar(255) unique not null,
    telegram_id       varchar(255) unique
);

CREATE TABLE channels
(
    id   serial      not null unique,
    name varchar(20) not null unique
);

INSERT INTO channels (name)
VALUES ('telegram'),
       ('email'),
       ('whatsapp');

CREATE TABLE user_channel
(
    id      serial                                                            not null unique,
    name_id int references channels (id) on delete cascade on update cascade  not null,
    recipient_id int references recipients (id) on delete cascade on update cascade not null
);

CREATE TABLE status
(
    id       serial                                                            not null unique,
    received bool                                                              not null default false,
    recipient_id  int references recipients (id) on delete cascade on update cascade not null
);