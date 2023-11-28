USE postgres;

DROP DATABASE IF EXISTS accounting;
CREATE DATABASE accounting with owner root;

USE accounting; 

CREATE TABLE transactions (
id bigserial primary key,
date date not null,
partner text not null,
account text not null,
payment_method text not null,
currency text not null,
sum NUMERIC(10,2) not null
);

GRANT all privileges on accounting.public.transactions to root;
GRANT USAGE, select on ALL SEQUENCES IN SCHEMA public to root;

