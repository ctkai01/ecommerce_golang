-- Active: 1676744047596@@localhost@5432@ecommerce

-- Country Table
CREATE TABLE countries (
  id SERIAL PRIMARY KEY,
  country_name VARCHAR(50) NOT NULL
);

-- Address Table
CREATE TABLE addresses (
  id SERIAL PRIMARY KEY,
  street_number VARCHAR(100) NOT NULL,
  address_line VARCHAR(50) NOT NULL,
  city VARCHAR(50) NOT NULL,
  country_id INT NOT NULL,

  CONSTRAINT fk_country
    FOREIGN KEY(country_id)
        REFERENCES countries(id)
);


-- User Table
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(50) NOT NULL UNIQUE,
  phone VARCHAR(30) NOT NULL,
  password VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();



-- User Adress Table
CREATE TABLE user_addreses (
  user_id INT NOT NULL,
  address_id INT NOT NULL,
  is_default BOOLEAN DEFAULT TRUE,
  PRIMARY KEY (user_id, address_id),
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
  CONSTRAINT fk_address FOREIGN KEY(address_id) REFERENCES addresses(id),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON user_addreses
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Shopping cart table
CREATE TABLE shopping_carts (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
CONSTRAINT fk_user
    FOREIGN KEY(user_id)
        REFERENCES users(id),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON shopping_carts
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Payment type table
CREATE TYPE enum_value_payment AS ENUM('COD', 'CREDIT_CARD');
CREATE TABLE payment_types (
  id SERIAL PRIMARY KEY,
  value enum_value_payment NOT NULL, 

  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON payment_types
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- User payment method TABLE
CREATE TABLE user_payment_methods (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  payment_type_id INT NOT NULL,
  provider VARCHAR(50),
  account_number VARCHAR(50),
  expiry_date TIMESTAMP,
  is_default BOOLEAN,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
  CONSTRAINT fk_payment_type FOREIGN KEY(payment_type_id) REFERENCES payment_types(id),

  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON user_payment_methods
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();


-- Product Category TABLE
CREATE TABLE product_categories (
  id SERIAL PRIMARY KEY,
  category_id INT NOT NULL,
  category_name VARCHAR(100) NOT NULL,
  CONSTRAINT fk_product_category FOREIGN KEY(category_id) REFERENCES product_categories(id),

  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON product_categories
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();


-- Product TABLE
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  category_id INT NOT NULL,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  product_image TEXT NOT NULL,
  CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES product_categories(id),

  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();




-- Product Item TABLE
CREATE TABLE product_items (
  id SERIAL PRIMARY KEY,
  product_id INT NOT NULL,
  qty_in_stock INT NOT NULL,
  product_image TEXT NOT NULL,
  price INT NOT NULL,
  CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id),

  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON product_items
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();



-- Shopping cart item TABLE
CREATE TABLE shopping_cart_items (
  id SERIAL PRIMARY KEY,
  cart_id INT NOT NULL,
  product_item_id INT NOT NULL,
  qty INT   NOT NULL,
  CONSTRAINT fk_cart FOREIGN KEY(cart_id) REFERENCES users(id),
  CONSTRAINT fk_product_item FOREIGN KEY(product_item_id) REFERENCES product_items(id),

  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON shopping_cart_items
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();


-- Promotion  TABLE
CREATE TABLE promotions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  discount_rate INT NOT NULL,
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON promotions
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();


-- Promotion Category TABLE
CREATE TABLE promotion_categories (
  category_id INT NOT NULL,
  promotion_id INT NOT NULL,
  PRIMARY KEY (category_id, promotion_id),
  CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES product_categories(id),
  CONSTRAINT fk_promotion FOREIGN KEY(promotion_id) REFERENCES promotions(id),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON promotion_categories
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();


-- Variation TABLE
CREATE TABLE variations (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  category_id INT NOT NULL,
  CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES product_categories(id),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON variations
FOR EACH ROW 
EXECUTE FUNCTION update_timestamp();


-- Variation Option TABLE
CREATE TABLE variation_options (
  id SERIAL PRIMARY KEY,
  value VARCHAR(100) NOT NULL,
  variation_id INT NOT NULL,
  CONSTRAINT fk_variation FOREIGN KEY(variation_id) REFERENCES variations(id),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON variation_options
FOR EACH ROW 
EXECUTE FUNCTION update_timestamp();


-- Product Configuration TABLE
CREATE TABLE product_configurations (
  product_id INT NOT NULL,
  variation_option_id INT NOT NULL,
  PRIMARY KEY (product_id, variation_option_id),
  CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id),
  CONSTRAINT fk_variation_option FOREIGN KEY(variation_option_id) REFERENCES variation_options(id),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON product_configurations
FOR EACH ROW 
EXECUTE FUNCTION update_timestamp();


-- Order status TABLE
CREATE TYPE enum_order_status AS ENUM('PENDING', 'DONE');
CREATE TABLE order_status (
  id SERIAL PRIMARY KEY,
  value enum_order_status DEFAULT 'PENDING',
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON order_status
FOR EACH ROW 
EXECUTE FUNCTION update_timestamp();

-- Shiping method TABLE
CREATE TABLE shipping_methods (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  price INT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON shipping_methods
FOR EACH ROW 
EXECUTE FUNCTION update_timestamp();


-- Shop order TABLE
CREATE TABLE shop_orders (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  payment_method_id INT NOT NULL,
  shipping_address INT NOT NULL,
  shipping_method INT NOT NULL,
  order_status INT NOT NULL,

  order_date TIMESTAMP NOT NULL,
  order_total INT NOT NULL,

  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
  CONSTRAINT fk_payment_method FOREIGN KEY(payment_method_id) REFERENCES user_payment_methods(id),
  CONSTRAINT fk_shipping_address FOREIGN KEY(shipping_address) REFERENCES addresses(id),
  CONSTRAINT fk_order_status FOREIGN KEY(order_status) REFERENCES order_status(id),
 
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON shop_orders
FOR EACH ROW 
EXECUTE FUNCTION update_timestamp();


-- Order line TABLE
CREATE TABLE order_lines (
  id SERIAL PRIMARY KEY,
  product_item_id INT NOT NULL,
  order_id INT NOT NULL,

  qty INT NOT NULL,
  price INT NOT NULL,

  CONSTRAINT fk_product_item_id FOREIGN KEY(product_item_id) REFERENCES product_items(id),
  CONSTRAINT fk_order_id FOREIGN KEY(order_id) REFERENCES shop_orders(id),
 
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON order_lines
FOR EACH ROW 
EXECUTE FUNCTION update_timestamp();


-- User review TABLE
CREATE TYPE enum_rating AS ENUM('1', '2', '3', '4', '5');

CREATE TABLE user_reviews (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  ordered_product_id INT NOT NULL,

  rating_value enum_rating NOT NULL,
  comment TEXT ,

  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
  CONSTRAINT fk_ordered_product_id FOREIGN KEY(ordered_product_id) REFERENCES order_lines(id),
 
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER table_updated_at
BEFORE UPDATE ON user_reviews
FOR EACH ROW 
EXECUTE FUNCTION update_timestamp();