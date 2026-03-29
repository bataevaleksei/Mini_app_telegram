CREATE TABLE IF NOT EXISTS orders(
  order_id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(user_id),
  items_id BIGINT NOT NULL REFERENCES cart_items(cart_item_id),
  total_price NUMERIC(10, 2) NOT NULL CHECK (total_price >= 0),
  status TEXT NOT NULL CHECK (status IN ('pending', 'completed', 'cancelled')),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)