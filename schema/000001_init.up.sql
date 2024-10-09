CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  public_id UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
  password_hash VARCHAR(256) NOT NULL,
  email VARCHAR(256) UNIQUE NOT NULL,
  name VARCHAR(64) NOT NULL,
  surname VARCHAR(64) NOT NULL,
  patronymic VARCHAR(64),
  passport VARCHAR(64),
  is_inactive BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE accounts (
  id SERIAL PRIMARY KEY,
  account_number UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
  owner_id INTEGER NOT NULL,
  balance DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
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
  amount DECIMAL(20, 2) NOT NULL,
  is_conversion BOOLEAN NOT NULL DEFAULT false,
  conversion_rate DECIMAL(14, 2),
  FOREIGN KEY (sender_account_id) REFERENCES accounts(id)
  ON DELETE RESTRICT ON UPDATE CASCADE,
  FOREIGN KEY (receiver_account_id) REFERENCES accounts(id)
  ON DELETE RESTRICT ON UPDATE CASCADE
);
