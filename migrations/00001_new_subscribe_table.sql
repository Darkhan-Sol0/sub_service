-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subs (
  id SERIAL PRIMARY KEY,
  service_name TEXT,
  price NUMERIC,
  user_id UUID,
  start_date TIMESTAMPTZ,
  end_date TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE subs;
-- +goose StatementEnd
