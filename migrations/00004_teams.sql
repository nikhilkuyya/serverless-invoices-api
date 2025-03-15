-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS teams(
  id BIGSERIAL PRIMARY KEY,
  company_name VARCHAR(255) NOT NULL,
  company_gstin VARCHAR(255) UNIQUE NOT NULL,
  company_address VARCHAR(255) NOT NULL,
  company_postal_code VARCHAR(10) NOT NULL,
  company_city VARCHAR(150) NOT NULL,
  company_state VARCHAR(150) NOT NULL,
  company_country VARCHAR(150) NOT NULL,
  website VARCHAR(255) NOT NULL,
  contact_email VARCHAR(255) NOT NULL,
  contact_phone VARCHAR(255) NOT NULL,
  currency VARCHAR(10) NOT NULL,
  invoice_due_days INT NOT NULL,
  invoice_late_fee INT NOT NULL,
  logo_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE teams;
-- +goose StatementEnd
