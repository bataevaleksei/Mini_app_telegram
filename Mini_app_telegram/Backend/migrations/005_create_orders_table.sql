CREATE TABLE IF NOT EXISTS orders(
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  items_id BIGINT NOT NULL,
  price INT NOT NULL,
  status TEXT NOT NULL CHECK (status IN ('pending', 'completed', 'cancelled')),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)