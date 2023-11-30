USE postgres;

DROP DATABASE IF EXISTS accounting;
CREATE DATABASE accounting with owner root;

USE accounting; 

CREATE TABLE partners (
  id bigserial primary key,
  short_name text not null,
  long_name text,
  address text,
  misc text, 
  alias_of_id bigserial
);

CREATE TABLE transactions (
id bigserial primary key,
date date not null,
partner_id bigserial,
account text not null,
payment_method text not null,
currency text not null,
sum NUMERIC(10,2) not null,
CONSTRAINT fk_partner_id
  FOREIGN KEY(partner_id)
  REFERENCES partners(id)
);

CREATE TABLE transaction_items (
  id bigserial primary key,
  transaction_id bigserial NOT NULL,
  name text NOT NULL,
  item_id bigserial,
  ammount NUMERIC(10,4) NOT NULL,
  unit_price NUMERIC(10,2) NOT NULL,
  CONSTRAINT fk_transaction_id
    FOREIGN KEY(transaction_id)
      REFERENCES transactions(id)
);

CREATE OR REPLACE VIEW transactions_view (
  Date, 
  Partner,
  Account,
  payment_method, 
  currency,
  sum
) AS
  select date, partners.short_name, account, payment_method, currency, sum
  from transactions 
  join
  partners on transactions.partner_id = partners.id
  ORDER BY date
  ;

CREATE OR REPLACE FUNCTION find_or_create_partner(
  p_short_name  text
)
RETURNS int AS 
$$
DECLARE
  r_partner_id INT;
BEGIN 
  SELECT id INTO r_partner_id
  FROM accounting.public.partners
  WHERE short_name = p_short_name;

  IF r_partner_id IS NULL THEN
    INSERT INTO accounting.public.partners (short_name)
    VALUES (p_short_name)
    RETURNING id INTO r_partner_id;
  END IF;

return r_partner_id;
END;
$$
LANGUAGE plpgsql;

GRANT INSERT, SELECT on accounting.public.transactions to root;
GRANT INSERT, SELECT on accounting.public.transaction_items to root;
GRANT INSERT, SELECT on accounting.public.partners to root;
GRANT USAGE, select on ALL SEQUENCES IN SCHEMA public to root;

