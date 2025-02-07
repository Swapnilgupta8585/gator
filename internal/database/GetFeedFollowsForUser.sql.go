// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: GetFeedFollowsForUser.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
WITH GetFeedFollow AS (
    SELECT id, created_at, updated_at, user_id, feed_id 
    FROM feed_follows
    WHERE feed_follows.user_id=$1
)
SELECT getfeedfollow.id, getfeedfollow.created_at, getfeedfollow.updated_at, getfeedfollow.user_id, getfeedfollow.feed_id, users.name AS user_name, feeds.name AS feed_name
FROM GetFeedFollow
INNER JOIN users
ON users.id = GetFeedFollow.user_id
INNER JOIN feeds
ON feeds.id = GetFeedFollow.feed_id
`

type GetFeedFollowsForUserRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	UserName  string
	FeedName  string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
			&i.UserName,
			&i.FeedName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
