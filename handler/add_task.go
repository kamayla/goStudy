package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"study/entity"
)

type AddTask struct {
	Service   AddTaskService
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

	t, err := at.Service.AddTask(ctx, b.Title)

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
