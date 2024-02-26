package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/skrb7f16/rssagg/internal/database"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	APIKeyString string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}

type FeedsFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

func convertDbUserToNormalUser(dbUser database.User) User {
	return User{
		ID:           dbUser.ID,
		Name:         dbUser.Name,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
		APIKeyString: dbUser.ApiKey,
	}
}

func convertDbFeedToNormalFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Name:      feed.Name,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		URL:       feed.Url,
		UserId:    feed.UserID,
	}
}

func convertDbFeedsToNormalFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, convertDbFeedToNormalFeed(feed))
	}
	return feeds
}

func convertDbFeedsFollowToNormalFeedsFollow(dbFeedsFollow database.FeedsFollow) FeedsFollow {
	return FeedsFollow{
		ID:        dbFeedsFollow.ID,
		CreatedAt: dbFeedsFollow.CreatedAt,
		UpdatedAt: dbFeedsFollow.UpdatedAt,
		UserId:    dbFeedsFollow.UserID,
		FeedId:    dbFeedsFollow.FeedID,
	}
}
