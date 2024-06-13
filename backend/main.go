package main

import(
	"encoding/json"
	"net/http"
)

func main() {
    http.HandleFunc("/api/choiceShop", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            var requestData RequestData
            err := json.NewDecoder(r.Body).Decode(&requestData)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            // 現在地から1km圏内の店を取得
            FetchApi(requestData.Lat, requestData.Lng)

            // 店を選んで返す
            shops := ChoiceShop(requestData.Lat, requestData.Lng)

            responseData := ResponseData{Shops: shops}
            json.NewEncoder(w).Encode(responseData)
        }
    })

    http.ListenAndServe(":8080", nil)
}