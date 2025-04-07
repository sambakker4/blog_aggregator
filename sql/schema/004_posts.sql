-- +goose Up
CREATE TABLE posts(
    id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMP,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_feeds
    FOREIGN KEY (feed_id)
    REFERENCES feeds(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
