package db

import (
	"context"
	"database/sql"
)

func CreateUser(db *sql.DB, ctx context.Context, name, email, password string) (*User, error) {
	user := &User{}
	err := db.QueryRowContext(ctx, `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id, name, email, password_hash`,
		name, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AuthenticateUser(db *sql.DB, ctx context.Context, email, password string) (*User, error) {
	user := &User{}
	err := db.QueryRowContext(ctx, `SELECT id, name, email, password_hash FROM users WHERE email = $1 AND password_hash = $2`,
		email, password).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
