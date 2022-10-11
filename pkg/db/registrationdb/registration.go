package registrationdb

import (
	"context"
	"fmt"
	"web-server/pkg/db"
	"web-server/pkg/models/registrationModels"
)

func CreateCredentials(ctx context.Context, login string, password string) (int64, error) {
	conn, err := db.InitDB()
	if err != nil {
		return 0, err
	}

	var insertedID int64
	err = conn.QueryRowContext(ctx, `INSERT INTO users("login", "password") VALUES($1, $2) RETURNING "id"`, login, password).Scan(&insertedID)

	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func CheckCredentialsInDb(login string) (bool, error) {
	conn, err := db.InitDB()
	if err != nil {
		return false, err
	}

	stmt := `SELECT "id", "login", "password" FROM "users" WHERE "login" = $1`

	row := conn.QueryRow(stmt, login)
	s := registrationModels.Register{}

	err = row.Scan(&s.Id, &s.Login, &s.Password)
	//if err != nil {
	//	return false, err
	//}

	if s.Login == login {
		return false, fmt.Errorf("such login exists")
	}

	return true, nil

}
