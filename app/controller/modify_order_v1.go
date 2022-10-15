package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ModifyOrderV1 - 提供時間変更
func (u *OrderController) ModifyOrderV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := strconv.Atoi(mux.Vars(r)["order_id"])
		if err != nil {
			http.Error(w, "Invalid parameter", http.StatusBadRequest)
			return
		}

		// FIXME: 未実装（状態変更のみでReadyと同一処理のため）

		RespondJSON(w, nil, http.StatusOK)
		return
	}
}
