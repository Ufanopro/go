-- name: CreateUsuario :exec
CREATE TABLE usuario (
  usuario VARCHAR(20) PRIMARY KEY,
  contrasena VARCHAR(20) NOT NULL
);
