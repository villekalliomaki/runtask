-- name: GetTaskProperties :one
SELECT *
FROM task_properties
WHERE id = $1 LIMIT 1;