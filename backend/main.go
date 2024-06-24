package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// CORSミドルウェアを定義
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ここで許可するオリジンを指定
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// プリフライトリクエストに対応
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// CORSミドルウェアを適用するルートハンドラーを設定
	http.Handle("/api/choiceShop", enableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var requestData RequestData
			err := json.NewDecoder(r.Body).Decode(&requestData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
			defer cancel()

			// APIを叩いて、店を1店舗選ぶ
			shops := FetchApi(ctx, requestData.Lat, requestData.Lng)

			responseData := ResponseData{Shops: shops}
			json.NewEncoder(w).Encode(responseData)
		}
	})))

	// 環境変数からポートを取得し、デフォルトを設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// HTTPサーバーを起動
	http.ListenAndServe(":"+port, nil)
}
