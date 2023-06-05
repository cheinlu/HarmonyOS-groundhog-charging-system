package logic_test

import (
	"fmt"
	"login-demo/internal/logic"
	"testing"
	"time"
)

func TestGetDistance(t *testing.T) {
	location1 := "39.90469,116.40717"
	location2 := "31.23037,121.47370"

	distance := logic.Geo.GetDistance(location1, location2)
	fmt.Printf("两地之间的直线距离为：%.2f千米\n", distance)
}

func TestCalculateAmountWithHourlyRates(t *testing.T) {
	startTime := time.Date(2023, 5, 15, 8, 20, 0, 0, time.Local)
	endTime := time.Date(2023, 5, 16, 18, 10, 0, 0, time.Local)
	hourlyRates := []logic.HourlyRate{
		{StartHour: 0, EndHour: 9, Rate: 0.75},
		{StartHour: 9, EndHour: 20, Rate: 1.5},
		{StartHour: 20, EndHour: 24, Rate: 0.75},
	}
	expectedAmount := 34 * (0.75*9 + 1.5*11 + 0.75*4)

	amount := logic.CalculateAmount(startTime, endTime, hourlyRates)

	if amount != expectedAmount {
		t.Errorf("Expected amount %f, but got %f", expectedAmount, amount)
	}
}
