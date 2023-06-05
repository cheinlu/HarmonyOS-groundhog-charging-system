package logic

import (
	"math"
	"strconv"
	"strings"
)

// 地理位置相关的功能工具文件

type GeoLogic struct{}

var Geo GeoLogic

// GetDistance 计算两个经纬度之间的直线距离，单位为千米
// 例子：
// distance := GetDistance(location1, location2)
// fmt.Printf("两地之间的直线距离为：%.2f千米\n", distance)
func (*GeoLogic) GetDistance(location1, location2 string) (distance float64) {
	// 将经纬度字符串转换为浮点数
	loc1 := strings.Split(location1, ",")
	lat1 := toFloat(loc1[0])
	lng1 := toFloat(loc1[1])

	loc2 := strings.Split(location2, ",")
	lat2 := toFloat(loc2[0])
	lng2 := toFloat(loc2[1])

	// 计算两地之间的直线距离
	R := 6371.0 // 地球半径，单位为千米
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lng2 - lng1) * math.Pi / 180.0
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180.0)*math.Cos(lat2*math.Pi/180.0)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance = R * c // 两地之间的直线距离，单位为千米

	return
}

// toFloat 将字符串转换为浮点数
func toFloat(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}
