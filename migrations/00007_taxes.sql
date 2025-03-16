-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS taxes(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) UNIQUE NOT NULL,
  label VARCHAR(100) NOT NULL,
  tax_percentage INT NOT NULL
)
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE taxes;
-- +goose StatementEnd
