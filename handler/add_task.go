package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"net/http"
	"study/entity"
	"study/store"
)

type AddTask struct {
	DB        *sqlx.DB
	Repo      store.Repository
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Title string `json:"title" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		ResposndJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	err := at.Validator.Struct(b)

	if err != nil {
		ResposndJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	t := &entity.Task{
		Title:  b.Title,
		Status: entity.TaskStatusTodo,
	}

	err = at.Repo.AddTask(ctx, at.DB, t)

	if err != nil {
		ResposndJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}

	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{
		ID: t.ID,
	}

	ResposndJSON(ctx, w, rsp, http.StatusOK)
}
