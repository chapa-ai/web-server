package authdb

import (
	"context"
	"web-server/pkg/db"
	"web-server/pkg/models/authmodels"
)

func GetUserAndPassword(ctx context.Context, login string) (*authmodels.User, error) {
	conn, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	sql := `SELECT "id", "login", "password" FROM "users" WHERE "login" = $1`
	rows, err := conn.QueryContext(ctx, sql, login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []*authmodels.User{}
	for rows.Next() {
		obj := authmodels.User{}
		if err := rows.Scan(&obj.Id, &obj.Login, &obj.Password); err != nil {
			return nil, err
		}

		list = append(list, &obj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list[0], nil
}

func SessionCreate(ctx context.Context, data *authmodels.Session) (int64, error) {
	conn, err := db.InitDB()
	if err != nil {
		return 0, err
	}

	var insertedID int64
	err = conn.QueryRowContext(ctx, `INSERT INTO sessions("userId", "token") VALUES($1, $2) RETURNING "id"`, data.UserId, data.Token).Scan(&insertedID)

	if err != nil {
		return 0, err
	}
	return insertedID, nil
}

func SessionDeleteToken(ctx context.Context, token string) (string, error) {
	conn, err := db.InitDB()
	if err != nil {
		return "", err
	}

	var deletedToken string

	err = conn.QueryRowContext(ctx, `DELETE FROM sessions WHERE "token" = $1 RETURNING "token"`, token).Scan(&deletedToken)
	if err != nil {
		return "", err
	}

	return deletedToken, nil
}
