-- name: GetPosts :many
select * from posts
    order by published_at DESC
limit $1;