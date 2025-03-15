-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS invoice_rows(
  id BIGSERIAL PRIMARY KEY,
  invoice_id BIGINT REFERENCES invoices(id) ON DELETE CASCADE,
  serial_no VARCHAR(100) NOT NULL,
  item VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  hsn_code VARCHAR(10) NOT NULL,
  quantity BIGINT NOT NULL,
  price BIGINT NOT NULL,
  unit VARCHAR(10) NOT NULL,
  invoice_row_order INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE invoice_rows;
-- +goose StatementEnd
