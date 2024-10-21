package repo_admin_auth

import (
	"context"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type Res struct {
	ID        uuid.UUID `db:"id"         json:"id"`
	Name      string    `db:"name"       json:"name"`
	Age       int       `db:"age"        json:"age"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (r *Read) Find(ctx context.Context) ([]Res, error) {
	q := `
	select * from try
	`
	rows, err := r.DB.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Res])
	if err != nil {
		return nil, err
	}
	return res, nil
}
