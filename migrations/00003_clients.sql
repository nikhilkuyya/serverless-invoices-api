-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS clients (
 id BIGSERIAL PRIMARY KEY,
 -- TODO: user_id
 company_name VARCHAR(255) NOT NULL,
 company_gstin VARCHAR(255) UNIQUE NOT NULL,
 company_address VARCHAR(255) NOT NULL,
 company_postalcode VARCHAR(10) NOT NULL,
 company_city VARCHAR(100) NOT NULL,
 company_state VARCHAR(100) NOT NULL,
 company_country VARCHAR(100) NOT NULL DEFAULT 'BHARAT',
 company_email VARCHAR(255) NOT NULL,
 company_bank_account_id BIGINT REFERENCES bank_accounts(id) ON DELETE CASCADE,
 created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clients;
-- +goose StatementEnd
