package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/lghartmann/go-rss-web-scraper/internal/database"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	APIKey    string     `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Url       string     `json:"url"`
	UserID    uuid.UUID  `json:"user_id"`
}

type FeedFollows struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	FeedID    uuid.UUID  `json:"feed_id"`
	UserID    uuid.UUID  `json:"user_id"`
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	PublishedAt time.Time  `json:"published_at"`
	Url         string     `json:"url"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: &dbUser.UpdatedAt.Time,
		APIKey:    dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: &dbFeed.UpdatedAt.Time,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}

	return feeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: &dbFeedFollow.UpdatedAt.Time,
		FeedID:    dbFeedFollow.ID,
		UserID:    dbFeedFollow.UserID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollows {
	feedFollows := []FeedFollows{}

	for _, feedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}

	return feedFollows
}

func databasePostToPost(dbPost database.Post) Post {
	var updatedAt *time.Time

	if dbPost.UpdatedAt.Valid {
		updatedAt = &dbPost.UpdatedAt.Time
	}

	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   updatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}

	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}

	return posts
}
