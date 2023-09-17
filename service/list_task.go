package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"study/entity"
)

type ListTask struct {
	DB   *sqlx.DB
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)

	if err != nil {
		return nil, err
	}

	return ts, nil
}
