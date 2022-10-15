package controller

import (
	"encoding/json"
	"net/http"
	"suzushin54/event-sourcing-with-go/app/usecase"
)

// CreateOrderV2 - 新規注文作成
func (u *OrderController) CreateOrderV2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req usecase.CreateInputV2
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			RespondJSON(w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		}

		_, err := u.orderUseCaseV2.CreateV2(r.Context(), &req)
		if err != nil {
			// FIXME: 400系と500系の判断をしてhttp statusを切り替える
			RespondJSON(w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		RespondJSON(w, nil, http.StatusOK)
		return
	}
}
