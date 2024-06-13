package main

import(
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func FetchApi(lat,lng float64) {

	//envファイル読み取り
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env file:",err)
		return
	}

	//apikeyがなかったとき
	apiKey := os.Getenv("HOTPEPPER_API_KEY")
	if apiKey == ""{
		fmt.Println("API key not found in environment variables.")
		return
	}

	//api叩く
	url := fmt.Sprintf("https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?key=%s&lat=%f&lng=%f&range=3&format=json",apiKey,lat,lng)

	//レスポンスをfetch
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch data:",err)
		return
	}
	defer resp.Body.Close()

	//本体の情報を取得
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("Failed to read response body:",err)
		return
	}

	//jsonファイル作成
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Failed to create file:",err)
	}
	defer file.Close()

	//作成したファイルに保存
	_, err = file.WriteString(string(body))
	if err != nil{
		fmt.Println("Failed to write to file:",err)
		return
	}

	fmt.Println("Successfully wrote to output.json")
}