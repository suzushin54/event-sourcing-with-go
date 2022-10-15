package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"suzushin54/event-sourcing-with-go/app/usecase"
)

// ReceiveOrderV1 - 店舗注文確認
func (u *OrderController) ReceiveOrderV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID, err := strconv.Atoi(mux.Vars(r)["order_id"])
		if err != nil {
			http.Error(w, "Invalid parameter", http.StatusBadRequest)
			return
		}

		err = u.orderUseCaseV1.ReceiveV1(
			r.Context(), &usecase.ReceiveInputV1{OrderID: uint64(orderID)},
		)
		if err != nil {
			// FIXME: 400系と500系の判断をしてhttp statusを切り替える
			RespondJSON(w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		RespondJSON(w, nil, http.StatusOK)
		return
	}
}
