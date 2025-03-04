package test

import (
	"context"
	"fmt"
	"testing"

	"albaranes/sqlc/generated"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	_ "github.com/go-sql-driver/mysql" // Importa el driver MySQL
)

// setupTestDB inicializa una base de datos de prueba con configuración hardcodeada
func setupTestDB(t *testing.T) (*generated.Queries, *sqlx.DB) {
	// Configuración hardcodeada para la conexión
	dbUser := "ufano"     // Cambia esto por tu usuario
	dbPassword := "ufano" // Cambia esto por tu contraseña
	dbHost := "127.0.0.1" // Dirección del host (localhost)
	dbPort := "3306"      // Puerto de MySQL
	dbName := "ufano_db"  // Nombre de la base de datos

	// Construir la cadena DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Conectar a la base de datos
	database, err := sqlx.Connect("mysql", dsn)
	require.NoError(t, err)

	// Limpia la tabla usuario antes de comenzar las pruebas
	_, err = database.Exec("TRUNCATE TABLE usuario")
	require.NoError(t, err)

	return generated.New(database), database
}

// teardownTestDB cierra la conexión a la base de datos después de las pruebas
func teardownTestDB(dbConn *sqlx.DB) {
	if dbConn != nil {
		err := dbConn.Close()
		if err != nil {
			fmt.Println("⚠️ Error al cerrar la conexión a la base de datos:", err)
		} else {
			fmt.Println("✅ Conexión a la base de datos cerrada correctamente")
		}
	}
}

func TestCreateUsuario(t *testing.T) {
	queries, dbConn := setupTestDB(t)
	defer teardownTestDB(dbConn)

	ctx := context.Background()
	params := generated.CreateUsuarioParams{
		Usuario:    "test_user",
		Contrasena: "test_password",
	}

	err := queries.CreateUsuario(ctx, params)
	require.NoError(t, err)

	user, err := queries.GetUsuario(ctx, params.Usuario)
	require.NoError(t, err)
	require.Equal(t, params.Usuario, user.Usuario)
	require.Equal(t, params.Contrasena, user.Contrasena)
}

func TestGetUsuario(t *testing.T) {
	queries, dbConn := setupTestDB(t)
	defer teardownTestDB(dbConn)

	ctx := context.Background()
	params := generated.CreateUsuarioParams{
		Usuario:    "test_user",
		Contrasena: "test_password",
	}
	err := queries.CreateUsuario(ctx, params)
	require.NoError(t, err)

	user, err := queries.GetUsuario(ctx, params.Usuario)
	require.NoError(t, err)
	require.Equal(t, params.Usuario, user.Usuario)
	require.Equal(t, params.Contrasena, user.Contrasena)

	// Prueba obtener un usuario inexistente
	_, err = queries.GetUsuario(ctx, "non_existent_user")
	require.Error(t, err)
}

func TestListUsuarios(t *testing.T) {
	queries, dbConn := setupTestDB(t)
	defer teardownTestDB(dbConn)

	ctx := context.Background()

	users := []generated.CreateUsuarioParams{
		{Usuario: "user1", Contrasena: "password1"},
		{Usuario: "user2", Contrasena: "password2"},
	}

	for _, u := range users {
		err := queries.CreateUsuario(ctx, u)
		require.NoError(t, err)
	}

	listedUsers, err := queries.ListUsuarios(ctx)
	require.NoError(t, err)
	require.Len(t, listedUsers, len(users))
}

func TestUpdateUsuario(t *testing.T) {
	queries, dbConn := setupTestDB(t)
	defer teardownTestDB(dbConn)

	ctx := context.Background()
	params := generated.CreateUsuarioParams{
		Usuario:    "test_user",
		Contrasena: "old_password",
	}
	err := queries.CreateUsuario(ctx, params)
	require.NoError(t, err)

	updateParams := generated.UpdateUsuarioParams{
		Contrasena: "new_password",
		Usuario:    "test_user",
	}
	err = queries.UpdateUsuario(ctx, updateParams)
	require.NoError(t, err)

	user, err := queries.GetUsuario(ctx, updateParams.Usuario)
	require.NoError(t, err)
	require.Equal(t, updateParams.Contrasena, user.Contrasena)
}

func TestDeleteUsuario(t *testing.T) {
	queries, dbConn := setupTestDB(t)
	defer teardownTestDB(dbConn)

	ctx := context.Background()
	params := generated.CreateUsuarioParams{
		Usuario:    "test_user",
		Contrasena: "test_password",
	}
	err := queries.CreateUsuario(ctx, params)
	require.NoError(t, err)

	err = queries.DeleteUsuario(ctx, params.Usuario)
	require.NoError(t, err)

	userList, err := queries.ListUsuarios(ctx)
	require.NoError(t, err)
	require.Empty(t, userList)
}
