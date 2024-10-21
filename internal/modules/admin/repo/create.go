package repo_admin_auth

import "context"

type CreateReq struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age"  db:"age"`
}

func (w *Write) Create(ctx context.Context, p CreateReq) error {
	q := `
	insert into try (name,age) values ($1,$2)
	`
	_, err := w.TX.Exec(ctx, q, p.Name, p.Age)
	if err != nil {
		return err
	}
	return nil
}
