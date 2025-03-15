-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS taxes(
  id BIGSERIAL PRIMARY KEY,
  label VARCHAR(100) NOT NULL,
  value INT NOT NULL
)
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE taxes;
-- +goose StatementEnd
