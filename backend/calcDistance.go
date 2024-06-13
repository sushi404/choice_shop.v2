package main 

import (
	"math"
)

//現在地と店の距離を計算して返す
func CalcDistance (lat float64, lng float64, shopLat float64, shopLng float64) float64 {
	const r = 6371	//地球の半径(km)
	
	//緯度経度をラジアンに変換
	nowLatRad := lat * math.Pi / 180
	nowLngRad := lng * math.Pi / 180
	shopLatRad := shopLat * math.Pi / 180
	shopLngRad := shopLng * math.Pi / 180

	//2地点間の差
	deltaLat := shopLatRad - nowLatRad
	deltaLng := shopLngRad - nowLngRad

	//ハヴァサイン公式で最短距離を求める
	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(nowLatRad) * math.Cos(shopLatRad) * math.Pow(math.Sin(deltaLng/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))
	distance := r * c

	distance = distance / 100

	//距離を返す
	return distance
	
}