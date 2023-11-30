CREATE OR REPLACE FUNCTION find_or_create_partner(
  p_short_name  text
)
RETURNS int 
AS $$
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

RETURN R_PARTNER_ID;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE create_transaction_item(
  p_transaction_id INT,
  p_item_name TEXT,
  p_amount NUMERIC(10,4),
  p_unit_price NUMERIC(10,2)
)
AS $$ 
BEGIN
  INSERT INTO accounting.public.transaction_items(
    transaction_id, 
    name,
    amount,
    unit_price
  ) VALUES (
    p_transaction_id, 
    p_item_name,
    p_amount,
    p_unit_price
  );
END;
$$
LANGUAGE plpgsql;
