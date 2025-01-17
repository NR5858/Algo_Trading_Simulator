package order

import (
	"fmt"

	"github.com/NR5858/Algo_Trading_Simulator/pkg/models"
)

func ConvertToFix(o *models.Order) string {
	var action string
	switch o.Action {
	case "BUY":
		action = "1"
	case "SELL":
		action = "2"
	}

	var priceType string
	switch o.PriceType {
	case "MARKET":
		priceType = "1"
	case "LIMIT":
		priceType = "2"
	}

	var duration string
	switch o.Duration {
	case "DAY":
		duration = "0"
	case "IOC":
		duration = "3"
	}

	return fmt.Sprintf(
		"8=FIX.4.2;35=D;55=%s;54=%s;38=%d;40=%s;59=%s;10=000;",
		o.Symbol, action, o.Quantity, priceType, duration,
	)
}
