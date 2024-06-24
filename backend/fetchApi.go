package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func FetchApi(ctx context.Context, lat, lng float64) []*ShopInfo {

	//envファイル読み取り
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env file:", err)
	}

	//apikeyを環境変数から取得
	apiKey := os.Getenv("HOTPEPPER_API_KEY")
	if apiKey == "" {
		fmt.Println("API key not found in environment variables.")
	}

	//api叩く
	url := fmt.Sprintf("https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?key=%s&lat=%f&lng=%f&range=3&format=json", apiKey, lat, lng)

	//コンテキストからリクエスト作成
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
	}

	//http.Clientでリクエスト実行
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to fetch API:", err)
	}
	defer resp.Body.Close()

	//本体の情報を取得
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
	}

	//jsonファイル作成
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Failed to create file:", err)
	}
	defer file.Close()

	//作成したファイルに保存
	_, err = file.WriteString(string(body))
	if err != nil {
		fmt.Println("Failed to write to file:", err)
	}

	fmt.Println("Successfully wrote to output.json")

	//店を1店舗選ぶ
	shops := ChoiceShop(lat, lng)

	return shops

}
