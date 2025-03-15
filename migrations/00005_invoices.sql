-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS invoice_statuses (
  id INT PRIMARY KEY,
  name VARCHAR(20) UNIQUE NOT NULL
);

INSERT INTO invoice_statuses (id, name)
VALUES (1, 'draft'), (2, 'booked'), (3, 'sent'),  (4,'paid'), (5, 'cancelled'), (6, 'archieved');

CREATE TABLE IF NOT EXISTS invoices(
  id BIGSERIAL PRIMARY KEY,
  -- TODO: user_id
  number VARCHAR(50) UNIQUE NOT NULL,
  status_id INT REFERENCES invoice_statuses(id),
  issued_at DATE NOT NULL,
  due_at DATE,
  late_fee NUMERIC(5,2),
  currency VARCHAR(6),
  from_name VARCHAR(255) NOT NULL,
  from_gstin VARCHAR(255) NOT NULL,
  from_address VARCHAR(255) NOT NULL,
  from_postal_code VARCHAR(10) NOT NULL,
  from_city VARCHAR(100) NOT NULL,
  from_state VARCHAR(100) NOT NULL,
  from_country VARCHAR(100) NOT NULL,
  from_email VARCHAR(255) NOT NULL,
  from_phone VARCHAR(100) NOT NULL,
  team_id BIGINT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
  bank_name VARCHAR(255) NOT NULL,
  bank_account_number VARCHAR(100) NOT NULL,
  bank_ifsc_code VARCHAR(25) NOT NULL,
  client_name VARCHAR(255) NOT NULL,
  client_gstin VARCHAR(255) NOT NULL,
  client_address VARCHAR(255) NOT NULL,
  client_postal_code VARCHAR(10) NOT NULL,
  client_city VARCHAR(100) NOT NULL,
  client_state VARCHAR(100) NOT NULL,
  company_country VARCHAR(100) NOT NULL,
  company_email VARCHAR(255) NOT NULL,
  client_id BIGINT REFERENCES clients(id) ON DELETE CASCADE,
  consignee_name VARCHAR(255) NOT NULL,
  consignee_gstin VARCHAR(255) NOT NULL,
  consignee_address VARCHAR(255) NOT NULL,
  consignee_postal_code VARCHAR(10) NOT NULL,
  consignee_city VARCHAR(100) NOT NULL,
  consignee_state VARCHAR(100) NOT NULL,
  consignee_country VARCHAR(100) NOT NULL,
  consignee_email VARCHAR(255) NOT NULL,
  consignee_id BIGINT REFERENCES clients(id) ON DELETE CASCADE,
  notes VARCHAR(255) NOT NULL,
  total BIGINT NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE invoice_statuses;
DROP TABLE invoices;
-- +goose StatementEnd
