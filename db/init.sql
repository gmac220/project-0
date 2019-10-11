create table customers(
    username varchar primary key,
    password varchar NOT NULL,
    firstname varchar NOT NULL,
    lastname varchar NOT NULL
);

create table employees(
    username varchar primary key,
    password varchar NOT NULL,
    firstname varchar NOT NULL,
    lastname varchar NOT NULL
);

create table applications(
    acntnumber serial primary key,
    username varchar NOT NULL,
    firstname varchar NOT NULL,
    lastname varchar NOT NULL,
    acntname varchar NOT NULL,
    joint boolean NOT NULL,
    username2 varchar NOT NULL,
    firstname2 varchar NOT NULL,
    lastname2 varchar NOT NULL
);

create table accounts(
    acntnumber serial primary key,
    acntname varchar NOT NULL,
    balance decimal NOT NULL,
    username varchar NOT NULL,
    username2 varchar
);