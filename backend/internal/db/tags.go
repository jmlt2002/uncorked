package db

import (
	"context"
	"database/sql"
)

func CreateTag(db *sql.DB, ctx context.Context, userID int64, name string, color string) (*Tag, error) {
	tag := &Tag{}
	err := db.QueryRowContext(ctx,
		`INSERT INTO tags (user_id, tag_name, color) VALUES ($1, $2, $3) RETURNING id, user_id, tag_name, color`,
		userID, name, color,
	).Scan(&tag.ID, &tag.UserID, &tag.Name, &tag.Color)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func DeleteTag(db *sql.DB, ctx context.Context, tagID int64, userID int64) error {
	_, err := db.ExecContext(ctx, `DELETE FROM tags WHERE id = $1 AND user_id = $2`, tagID, userID)
	return err
}
