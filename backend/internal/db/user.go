package db

import (
	"context"
	"database/sql"
)

func CreateUser(db *sql.DB, ctx context.Context, name, username, email, password string) (*User, error) {
	user := &User{}
	err := db.QueryRowContext(ctx, `INSERT INTO users (name, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id, name, username, email, password_hash`,
		name, username, email, password).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AuthenticateUser(db *sql.DB, ctx context.Context, username, password string) (*User, error) {
	user := &User{}
	err := db.QueryRowContext(ctx, `SELECT id, name, username, email, password_hash FROM users WHERE username = $1 AND password_hash = $2`,
		username, password).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
