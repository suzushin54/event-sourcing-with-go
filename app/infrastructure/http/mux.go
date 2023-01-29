package http

import (
	"net/http"
	"suzushin54/event-sourcing-with-go/app/controller"

	"github.com/gorilla/mux"
)

func NewMux(oc *controller.OrderController) http.Handler {
	r := mux.NewRouter()

	// NOTE: v1 State Sourcing
	// 新規注文
	r.HandleFunc("/v1/orders", oc.CreateOrderV1()).Methods("POST")
	// 店舗注文確認
	r.HandleFunc("/v1/orders/{order_id:[0-9]+}/receive", oc.ReceiveOrderV1()).Methods("PUT")
	// 受け渡し準備完了
	//r.HandleFunc("/v1/orders/{order_id:[0-9]+}/ready", oc.ReadyOrderV1()).Methods("PUT")
	// 受け渡し完了
	//r.HandleFunc("/v1/orders/{order_id:[0-9]+}/complete", oc.CompleteOrderV1()).Methods("PUT")
	// キャンセル
	r.HandleFunc("/v1/orders/{order_id:[0-9]+}/cancel", oc.CancelOrderV1()).Methods("PUT")
	// 提供時間変更
	r.HandleFunc("/v1/orders/{order_id:[0-9]+}/modify", oc.ModifyOrderV1()).Methods("PUT")

	// NOTE: v2 Event Sourcing
	// 新規注文
	r.HandleFunc("/v2/orders", oc.CreateOrderV2()).Methods("POST")
	// 店舗注文確認
	r.HandleFunc("/v2/orders/{order_id:[0-9]+}/receive", oc.ReceiveOrderV2()).Methods("PUT")
	// 受け渡し準備完了
	//r.HandleFunc("/v2/orders/{order_id:[0-9]+}/ready", oc.ReadyOrderV2()).Methods("PUT")
	// 受け渡し完了
	//r.HandleFunc("/v2/orders/{order_id:[0-9]+}/complete", oc.CompleteOrderV2()).Methods("PUT")
	// 注文キャンセル
	r.HandleFunc("/v2/orders/{order_id:[0-9]+}/cancel", oc.CancelOrderV2()).Methods("PUT")
	// 提供時間変更
	r.HandleFunc("/v2/orders/{order_id:[0-9]+}/modify", oc.ModifyOrderV2()).Methods("PUT")

	// 確認用
	r.HandleFunc("/v2/orders/{order_id:[0-9]+}", oc.FindOrderV2()).Methods("GET")

	// 注文詳細
	r.HandleFunc("/v1/orders/{order_id:[0-9]+}", oc.FindOrderV1()).Methods("GET")

	// 注文詳細
	//r.HandleFunc("/v2/orders/{order_id:[0-9]+}", oc.FindOrderV2()).Methods("GET")

	r.HandleFunc(
		"/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, err := w.Write([]byte(`{"status":"ok"}`))
			if err != nil {
				return
			}
		},
	)
	http.Handle("/", r)
	return r
}
