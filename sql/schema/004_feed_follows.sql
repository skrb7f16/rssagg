-- +goose Up
CREATE TABLE feeds_follow(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE NOT NULL,
    UNIQUE(user_id, feed_id)
);
-- +goose Down
DROP TABLE feeds_follow;