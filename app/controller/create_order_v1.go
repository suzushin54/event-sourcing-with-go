package controller

import (
	"encoding/json"
	"net/http"

	"suzushin54/event-sourcing-with-go/app/usecase"
)

// CreateOrderV1 - 新規注文作成
func (u *OrderController) CreateOrderV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req usecase.CreateInputV1
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			RespondJSON(w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		}

		_, err := u.orderUseCaseV1.CreateV1(r.Context(), &req)
		if err != nil {
			// FIXME: 400系と500系の判断をしてhttp statusを切り替える
			RespondJSON(w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		RespondJSON(w, nil, http.StatusOK)
		return
	}
}
