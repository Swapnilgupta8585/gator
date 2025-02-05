
-- name: GetFeedFollowsForUser :many
WITH GetFeedFollow AS (
    SELECT * 
    FROM feed_follows
    WHERE feed_follows.user_id=$1
)
SELECT GetFeedFollow.*, users.name AS user_name, feeds.name AS feed_name
FROM GetFeedFollow
INNER JOIN users
ON users.id = GetFeedFollow.user_id
INNER JOIN feeds
ON feeds.id = GetFeedFollow.feed_id;
