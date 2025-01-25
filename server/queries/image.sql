-- name: GetImageByID :one
SELECT sqlc.embed(images) FROM images
WHERE images.id = sqlc.arg(id);

-- name: CreateImage :one
INSERT INTO images (id, bucket_name, object_id, width, height, file_size) 
VALUES (sqlc.arg(id), sqlc.arg(bucket_name), sqlc.arg(object_id), sqlc.arg(width), sqlc.arg(height), sqlc.arg(file_size)) RETURNING *;

-- name: GetCollectionByID :one
SELECT sqlc.embed(collections) FROM collections
WHERE collections.id = sqlc.arg(id);

-- name: GetCollectionByOwnerID :many
SELECT sqlc.embed(collections) FROM collections
WHERE collections.owner_id = sqlc.arg(owner_id);

-- name: CreateCollection :one
INSERT INTO collections (id, collection_name, owner_id, public) 
VALUES (sqlc.arg(id), sqlc.arg(collection_name), sqlc.arg(owner_id), sqlc.arg(public)) RETURNING *;

-- name: AddImageToCollection :one
INSERT INTO collection_images (collection_id, image_id, sort_key)
VALUES (sqlc.arg(collection_id), sqlc.arg(image_id), sqlc.arg(sort_key)) RETURNING *;

-- name: RemoveImageFromCollection :exec
DELETE FROM collection_images
WHERE collection_id = sqlc.arg(collection_id) AND image_id = sqlc.arg(image_id);

-- name: GetCollectionImagesByCollectionID :many
SELECT sqlc.embed(images) FROM images
JOIN collection_images ON images.id = collection_images.image_id
WHERE collection_images.collection_id = sqlc.arg(collection_id);