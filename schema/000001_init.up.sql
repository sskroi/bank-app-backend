CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  public_id UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
  email VARCHAR(256) UNIQUE NOT NULL,
  password_hash VARCHAR(256) NOT NULL,
  passport VARCHAR(64),
  name VARCHAR(64) NOT NULL,
  surname VARCHAR(64) NOT NULL,
  patronymic VARCHAR(64),
  is_inactive BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE accounts (
  id SERIAL PRIMARY KEY,
  number UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
  owner_id INTEGER NOT NULL,
  balance DECIMAL(20, 2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0),
  currency VARCHAR(16) NOT NULL,
  is_close BOOLEAN NOT NULL DEFAULT false,
  FOREIGN KEY (owner_id) REFERENCES users(id)
  ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE transactions (
  id BIGSERIAL PRIMARY KEY,
  public_id UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
  sender_account_id INTEGER NOT NULL,
  receiver_account_id INTEGER NOT NULL,
  sent DECIMAL (20, 2) NOT NULL,
  received DECIMAL(20, 2) NOT NULL,
  is_conversion BOOLEAN NOT NULL DEFAULT false,
  conversion_rate DECIMAL(16, 4),
  dt TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (sender_account_id) REFERENCES accounts(id)
  ON DELETE RESTRICT ON UPDATE CASCADE,
  FOREIGN KEY (receiver_account_id) REFERENCES accounts(id)
  ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE conversion_rates (
  currency_from CHAR(3),
  currency_to CHAR(3),
  rate DECIMAL(20, 8),
  dt_updated TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY (currency_from, currency_to)
);

-- test conversion rates data
INSERT INTO conversion_rates (currency_from, currency_to, rate) VALUES
  ('usd', 'rub', 103.5), ('rub', 'usd', 0.00966),
  ('usd', 'eur', 0.9649), ('eur', 'usd', 1.0365),
  ('eur', 'rub', 107.273), ('rub', 'eur', 0.00932);

