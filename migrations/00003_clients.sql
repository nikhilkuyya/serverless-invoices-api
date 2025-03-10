-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS clients (
 id BIGSERIAL PRIMARY KEY,
 -- user_id
 company_name VARCHAR(150) NOT NULL,
 company_gstin VARCHAR(100) UNIQUE NOT NULL,
 company_address VARCHAR(200) NOT NULL,
 company_postalcode VARCHAR(10) NOT NULL,
 company_city VARCHAR(50) NOT NULL,
 company_state VARCHAR(50) NOT NULL,
 company_country VARCHAR(50) NOT NULL DEFAULT 'BHARAT',
 company_email VARCHAR(120) NOT NULL,
 company_bank_account_id BIGINT REFERENCES bank_accounts(id) ON DELETE CASCADE,
 created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clients;
-- +goose StatementEnd
