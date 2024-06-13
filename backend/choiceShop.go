package main

import(
	"encoding/json"
	"fmt"
	"io"
	"os"
	"math/rand"
	"time"
	"net/url"
)

//店をランダムで1つ選んで返す
func ChoiceShop(lat float64, lng float64)[]*ShopInfo{

	//jsonファイルを開く
	file, err := os.Open("output.json")
	if err != nil{
		fmt.Println("Failed to open file:",err)
		return nil
	}
	defer file.Close()

	//ファイル読み込み
	bytes, err := io.ReadAll(file)
	if err != nil{
		fmt.Println("Failed to read file",err)
		return nil
	}

	//レスポンス作成
	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil{
		fmt.Println("Failed to unmarshal JSON:",err)
		return nil
	}

	//店をランダム(距離)で選ぶためのスコア付け

	//各店の距離の逆数を計算
	inverseSum := 0.0
	for i, shop := range response.Results.Shop{

		fmt.Println("------------------------")
		//現在地と店の距離を計算
		distance := CalcDistance(lat, lng, shop.shopLat, shop.shopLng)
		fmt.Println("distance:",distance)
		
		//距離の逆数を計算
		inverse := 1 / distance
		response.Results.Shop[i].InverseDistance = inverse
		fmt.Println("inverse:",inverse)

		//逆数の合計
		inverseSum += inverse
		fmt.Println("inverseSum:",inverseSum)

		//店に逆数を設定
		shop.InverseDistance = inverse
	}

	//店が選ばれる確率を計算
	for i, shop := range response.Results.Shop{
		probability := shop.InverseDistance / inverseSum
		response.Results.Shop[i].Probability = probability

		fmt.Println("------------------------")
		fmt.Println("Probability:",shop.Probability)
		fmt.Println("InverseDistance:",shop.InverseDistance)

	}

	//確率に基づいてランダムに店を選ぶ
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Float64()
	cumulativeProbability := 0.0

	for _, shop := range response.Results.Shop{
		cumulativeProbability += shop.Probability
		if randomValue < cumulativeProbability {
			fmt.Println("------------------------")
			fmt.Println("Name:",shop.Name)
			fmt.Println("Genre:",shop.Genre)
			fmt.Println("Address:",shop.Address)
			fmt.Println("Open Hour:",shop.OpenHour)
			
			
			// 店舗の名前からGoogleマップの検索URLを生成
			encodedName := url.QueryEscape(shop.Name)
			googleMapURL := "https://www.google.com/maps/search/?api=1&query=" + encodedName
			fmt.Println("Google Map:", googleMapURL)
			fmt.Println("------------------------")

			return []*ShopInfo{
				{
					Name:        shop.Name,
					Genre:       shop.Genre.Name,
					Address:     shop.Address,
					OpenHour:    shop.OpenHour,
					GoogleMapURL: googleMapURL,
				},				
			}

			break
		}
	}
	return nil
}