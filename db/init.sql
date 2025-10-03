DROP TABLE IF EXISTS wine_tags CASCADE;
DROP TABLE IF EXISTS tags CASCADE;
DROP TABLE IF EXISTS wines CASCADE;
DROP TABLE IF EXISTS storage_locations CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- USERS
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- STORAGE LOCATIONS
CREATE TABLE storage_locations (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    location_name VARCHAR(255) NOT NULL
);



-- WINES
CREATE TABLE wines (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(150) NOT NULL,
    wine_producer VARCHAR(150),
    region VARCHAR(150),
    stock_quantity INT DEFAULT 0 CHECK (stock_quantity >= 0),
    storage_location_id BIGINT REFERENCES storage_locations(id) ON DELETE SET NULL,
    photo_url TEXT
);

-- TAGS (reusable definitions)
CREATE TABLE tags (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    tag_name VARCHAR(100) NOT NULL,
    color VARCHAR(7) CHECK (color ~ '^#[0-9A-Fa-f]{6}$'),
    UNIQUE(user_id, tag_name)
);

-- WINE_TAGS (associates tags with wines)
CREATE TABLE wine_tags (
    id BIGSERIAL PRIMARY KEY,
    wine_id BIGINT NOT NULL REFERENCES wines(id) ON DELETE CASCADE,
    tag_id BIGINT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    UNIQUE(wine_id, tag_id)
);

CREATE INDEX idx_wines_user_id ON wines (user_id);
CREATE INDEX idx_wines_storage_location_id ON wines (storage_location_id);
CREATE INDEX idx_tags_user_id ON tags (user_id);
CREATE INDEX idx_tags_tag_name ON tags (tag_name);
CREATE INDEX idx_wine_tags_wine_id ON wine_tags (wine_id);
CREATE INDEX idx_wine_tags_tag_id ON wine_tags (tag_id);
