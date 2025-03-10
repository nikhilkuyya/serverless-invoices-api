-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bank_accounts (
  id BIGSERIAL PRIMARY KEY,
  -- user_id
  account_nick_name VARCHAR(100) NOT NULL,
  bank_name VARCHAR(255) NOT NULL,
  bank_account_number VARCHAR(100) UNIQUE NOT NULL,
  bank_ifsc_code VARCHAR(25) NOT NULL,
  bank_description VARCHAR(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT nick_name_account_number_notnull CHECK (
    NOT (
      (
        account_nick_name IS NULL
        OR account_nick_name = ''
      )
      AND (
        bank_account_number IS NULL
        OR bank_account_number = ''
      )
    )
  )
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bank_accounts;
-- +goose StatementEnd
