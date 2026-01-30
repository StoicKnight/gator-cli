-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3
)
RETURNING *;

-- name: ListFeeds :many
SELECT f.name, f.url, u.name AS user_name
FROM feeds f
JOIN users u ON u.id = f.user_id
ORDER BY f.created_at DESC;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET
last_fetched_at = NOW(),
updated_at = NOW()
WHERE user_id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1
FOR UPDATE SKIP LOCKED;
