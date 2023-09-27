package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"study/auth"
	"study/entity"
)

type ListTask struct {
	DB   *sqlx.DB
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	ts, err := l.Repo.ListTasks(ctx, l.DB, id)

	if err != nil {
		return nil, err
	}

	return ts, nil
}
