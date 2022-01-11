/**
 * trigger function 'update_purchase_date'
 */
 -- ensure that ROW.purchase_date is updated before it is created
CREATE OR REPLACE FUNCTION update_purchase_date()
RETURNS TRIGGER AS $$
BEGIN
  NEW.purchase_date = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

/**
 * trigger function 'update_date_time'
 */
 -- ensure that ROW.date_time is updated before it is created
CREATE OR REPLACE FUNCTION update_date_time()
RETURNS TRIGGER AS $$
BEGIN
  NEW.date_time = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER invoice_header_update_purchase_date
BEFORE INSERT ON invoice_header
FOR EACH ROW EXECUTE PROCEDURE update_purchase_date();

CREATE TRIGGER client_search_history_update_date_time
BEFORE INSERT ON client_search_history
FOR EACH ROW EXECUTE PROCEDURE update_date_time();