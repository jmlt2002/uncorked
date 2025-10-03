package db

import (
	"context"
	"database/sql"
)

func CreateStorageLocation(db *sql.DB, ctx context.Context, userID int64, name string) (*StorageLocation, error) {
	location := &StorageLocation{}
	err := db.QueryRowContext(ctx, `INSERT INTO storage_locations (user_id, name) VALUES ($1, $2) RETURNING id, user_id, name`,
		userID, name).Scan(&location.ID, &location.UserID, &location.Name)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func DeleteStorageLocation(db *sql.DB, ctx context.Context, id int64, userID int64) error {
	_, err := db.ExecContext(ctx, `DELETE FROM storage_locations WHERE id = $1 AND user_id = $2`, id, userID)
	return err
}
