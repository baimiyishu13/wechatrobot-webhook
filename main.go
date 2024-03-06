package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/baimiyishu13/wechatrobot-webhook/model"
	"github.com/baimiyishu13/wechatrobot-webhook/notifier"
	"github.com/go-chi/chi/v5"
)

var (
	h        bool
	RobotKey string
	addr     string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&RobotKey, "RobotKey", "", "global wechatrobot webhook key")
	flag.StringVar(&addr, "addr", ":3000", "listen addr")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	// 路由
	r := chi.NewRouter()

	r.Post("/webhook", func(w http.ResponseWriter, r *http.Request) {
		var notification model.Notification

		err := json.NewDecoder(r.Body).Decode(&notification)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 从 HTTP 请求的 URL 查询参数中获取名为 "key" 的值
		RobotKey := r.URL.Query().Get("key")

		err = notifier.Send(notification, RobotKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "send to wechatbot successful!"}`))
	})

	fmt.Printf("Starting the server on %v ...\n", addr)
	http.ListenAndServe(addr, r)
}
