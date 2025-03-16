-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS invoice_taxes(
  id BIGSERIAL PRIMARY KEY,
  invoice_row_id BIGINT REFERENCES invoice_rows(id) ON DELETE CASCADE,
  name VARCHAR(100) NOT NULL,
  label VARCHAR(100) NOT NULL,
  tax_percentage INT NOT NULL
)
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DELETE TABLE invoice_taxes;
-- +goose StatementEnd
