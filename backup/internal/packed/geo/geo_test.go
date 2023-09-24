package geo_test

import (
	"fmt"
	"login-demo/internal/logic"
	"login-demo/internal/packed/geo"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestGetDistance(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		location1 := "39.90469,116.40717"
		location2 := "31.23037,121.47370"

		distance := geo.GetDistance(location1, location2)
		fmt.Printf("两地之间的直线距离为：%.2f千米\n", distance)
		t.Assert(1067, int(distance))
	})
}

func TestCalculateAmountWithHourlyRates(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		startTime := time.Date(2023, 5, 15, 8, 20, 0, 0, time.Local)
		endTime := time.Date(2023, 5, 16, 18, 10, 0, 0, time.Local)
		hourlyRates := []logic.HourlyRate{
			{StartHour: 0, EndHour: 9, Rate: 0.75},
			{StartHour: 9, EndHour: 20, Rate: 1.5},
			{StartHour: 20, EndHour: 24, Rate: 0.75},
		}
		// 一共 33 小时 50 分钟，
		expectedAmount := 39.8

		amount := logic.CalculateAmount(startTime, endTime, hourlyRates)
		t.Assert(amount, expectedAmount)
		// (amount, expectedAmount)
		// if amount != expectedAmount {
		// 	t.Errorf("Expected amount %f, but got %f", expectedAmount, amount)
		// }
	})
}
