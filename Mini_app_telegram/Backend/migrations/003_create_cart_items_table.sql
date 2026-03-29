CREATE TABLE IF NOT EXISTS cart_items(
  cart_item_id BIGSERIAL PRIMARY KEY,
  cart_id BIGINT NOT NULL REFERENCES carts(cart_id) ON DELETE CASCADE,
  product_id BIGINT NOT NULL REFERENCES products(product_id),
  quantity INT NOT NULL CHECK (quantity > 0)
)