package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ModifyOrderV2 - 提供時間変更
func (u *OrderController) ModifyOrderV2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := strconv.Atoi(mux.Vars(r)["order_id"])
		if err != nil {
			http.Error(w, "Invalid parameter", http.StatusBadRequest)
			return
		}

		// TODO
		//err = u.orderUseCaseV2.CancelV2(
		//	r.Context(), &usecase.CancelInputV2{OrderID: uint64(orderID)},
		//)
		if err != nil {
			// FIXME: 400系と500系の判断をしてhttp statusを切り替える
			RespondJSON(w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		RespondJSON(w, nil, http.StatusOK)
		return
	}
}
