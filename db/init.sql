DROP TABLE IF EXISTS tags CASCADE;
DROP TABLE IF EXISTS reviews CASCADE;
DROP TABLE IF EXISTS wines CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- USERS
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- WINES
CREATE TABLE wines (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(150) NOT NULL,
    winery VARCHAR(150),
    vintage INT CHECK (vintage >= 1900 AND vintage <= EXTRACT(YEAR FROM CURRENT_DATE)::INT + 5),
    varietal VARCHAR(100),
    region VARCHAR(150),
    category VARCHAR(100),
    purchase_date DATE,
    price NUMERIC(10,2),
    stock_quantity INT DEFAULT 0 CHECK (stock_quantity >= 0),
    drinking_window_start INT,
    drinking_window_end INT,
    storage_location VARCHAR(255),
    photo_url TEXT
);

-- REVIEWS
CREATE TABLE reviews (
    id BIGSERIAL PRIMARY KEY,
    wine_id BIGINT NOT NULL REFERENCES wines(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 100),
    tasting_notes TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    photo_urls TEXT[]  -- PostgreSQL array for multiple photo URLs
);

-- TAGS
CREATE TABLE tags (
    id BIGSERIAL PRIMARY KEY,
    wine_id BIGINT NOT NULL REFERENCES wines(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    tag_name VARCHAR(100) NOT NULL
);

-- optimize reads
CREATE INDEX idx_wines_user_id ON wines (user_id);
CREATE INDEX idx_reviews_wine_id ON reviews (wine_id);
CREATE INDEX idx_reviews_user_id ON reviews (user_id);
CREATE INDEX idx_tags_wine_id ON tags (wine_id);
CREATE INDEX idx_tags_user_id ON tags (user_id);
CREATE INDEX idx_tags_tag_name ON tags (tag_name);
