package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlQueri := `
	INSERT INTO tasks(title, description, completed, created_at)
	VALUES ($1, $2, $3, $4 );
	`

	_, err := conn.Exec(ctx, sqlQueri, task)
	return err
}
