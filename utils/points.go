package utils

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"receiptprocessor/models"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0
	breakdown := []string{}

	// Alphanumeric characters in retailer name; One point per alphanumeric character in retailer name
	alphanum := 0
	for _, c := range receipt.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			alphanum++
		}
	}
	points += alphanum
	breakdown = append(breakdown, fmt.Sprintf("%3d points - retailer name has %d alphanumeric characters", alphanum, alphanum))

	total := parseFloat(receipt.Total)

	// Round dollar amount
	if total == float64(int(total)) {
		points += 50
		breakdown = append(breakdown, " 50 points - total is a round dollar amount")
	}

	// Total is multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
		breakdown = append(breakdown, " 25 points - total is a multiple of 0.25")
	}

	// 5 points for every two items
	itemCount := len(receipt.Items)
	itemPoints := (itemCount / 2) * 5
	points += itemPoints
	breakdown = append(breakdown, fmt.Sprintf("%3d points - %d items (%d pairs @ 5 points each)", itemPoints, itemCount, itemCount/2))

	// Item descriptions length divisible by 3
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price := parseFloat(item.Price)
			itemPoints := int(math.Ceil(price * 0.2))
			points += itemPoints
			breakdown = append(breakdown, fmt.Sprintf("%3d points - \"%s\" (%d chars, %.2f x 0.2)", itemPoints, desc, len(desc), price))
		}
	}

	// Odd purchase day
	if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil && date.Day()%2 == 1 {
		points += 6
		breakdown = append(breakdown, "  6 points - purchase day is odd")
	}

	// Time between 2:00pm and 4:00pm
	if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil && t.Hour() == 14 {
		points += 10
		breakdown = append(breakdown, "10 points - purchase time is between 2pm and 4pm")
	}

	// Print results breakdown to terminal
	fmt.Println("📊 Points Breakdown:")
	for _, line := range breakdown {
		fmt.Println(" ", line)
	}
	fmt.Printf("🎯 Total Points: %d\n", points)
	return points
}

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Invalid float value: %q — defaulting to 0.0", s)
		return 0.0
	}
	return f
}
