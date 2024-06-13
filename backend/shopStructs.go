package main

//フロントから緯度経度入手
type RequestData{
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

//フロントに返すデータ
type ResponseData struct{
	Shops []*ShopInfo `json:"shops"`
}

//output.jsonから切り分け
type Shop struct{
	Name string `json:"name"`
	Address string `json:"address"`
	Genre Genre `json:"genre"`
	OpenHour string `json:"open"`
	shopLat float64 `json:"lat"`
	shopLng float64 `json:"lng"`
	InverseDistance float64 //距離を基に店を選ぶために追加した。距離を逆数にしたもの
	Probability float64 //距離を基に店を選ぶために追加した。確率
}

//ジャンル名のみを抽出
type Genre struct{
	Name string `json:"name"`
	//Code string `json:"code"`
}

//APIから取得したもの
type Response struct{
	Results struct{
		Shop []Shop `json:"shop"`
	}`json:"results"`
}

//最終的な店の情報
type ShopInfo struct {
	Name string
	Genre string
	Address string
	OpenHour string
	GoogleMapURL string
}