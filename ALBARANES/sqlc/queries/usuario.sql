-- name: CreateUsuario :exec
INSERT INTO usuario (usuario, contrasena) VALUES (?, ?);

-- name: GetUsuario :one
SELECT * FROM usuario WHERE usuario = ?;

-- name: ListUsuarios :many
SELECT * FROM usuario;

-- name: UpdateUsuario :exec
UPDATE usuario SET contrasena = ? WHERE usuario = ?;

-- name: DeleteUsuario :exec
DELETE FROM usuario WHERE usuario = ?;
