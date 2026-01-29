-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        gen_random_uuid(),
        NOW(),
        NOW(),
        $1,
        $2
    )
    RETURNING *
)
SELECT
    iff.*,
    u.name AS user_name,
    f.name AS feed_name
FROM inserted_feed_follow iff
INNER JOIN users u ON iff.user_id = u.id
INNER JOIN feeds f ON iff.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT
  ff.*,
  u.name AS user_name,
  f.name AS feed_name
FROM feed_follows ff
INNER JOIN feeds f ON ff.feed_id = f.id
INNER JOIN users u ON ff.user_id = u.id
WHERE u.name = $1
ORDER BY f.name ASC;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;
