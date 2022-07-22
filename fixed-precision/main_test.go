package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func StringFixedBank(value decimal.Decimal) string {
	return value.StringFixedBank(2)
}

func StringFixed(value decimal.Decimal) string {
	return value.StringFixed(2)
}

func W1(value decimal.Decimal) string {
	v, _ := value.Float64()
	return fmt.Sprintf("%.2f", v)
}

func W2(value decimal.Decimal) string {
	v, _ := value.Float64()
	return strconv.FormatFloat(v, 'f', 2, 64)
}

func BenchmarkStringFixedBank(b *testing.B) {
	values := []decimal.Decimal{
		decimal.NewFromFloat(2.00),
		decimal.NewFromFloat(1.50),
		decimal.NewFromFloat(3.553),
		decimal.NewFromFloat(4.5),
	}
	expects := []string{
		"2.00",
		"1.50",
		"3.55",
		"4.50",
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for idx, value := range values {
			if StringFixedBank(value) != expects[idx] {
				b.Error("FAIL")
			}
		}
	}
}

func BenchmarkStringFixed(b *testing.B) {
	values := []decimal.Decimal{
		decimal.NewFromFloat(2.00),
		decimal.NewFromFloat(1.50),
		decimal.NewFromFloat(3.553),
		decimal.NewFromFloat(4.5),
	}
	expects := []string{
		"2.00",
		"1.50",
		"3.55",
		"4.50",
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for idx, value := range values {
			if StringFixed(value) != expects[idx] {
				b.Error("FAIL")
			}
		}
	}
}

func BenchmarkW1(b *testing.B) {
	values := []decimal.Decimal{
		decimal.NewFromFloat(2.00),
		decimal.NewFromFloat(1.50),
		decimal.NewFromFloat(3.553),
		decimal.NewFromFloat(4.5),
	}
	expects := []string{
		"2.00",
		"1.50",
		"3.55",
		"4.50",
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for idx, value := range values {
			if W1(value) != expects[idx] {
				b.Error("FAIL")
			}
		}
	}
}

func BenchmarkW2(b *testing.B) {
	values := []decimal.Decimal{
		decimal.NewFromFloat(2.00),
		decimal.NewFromFloat(1.50),
		decimal.NewFromFloat(3.553),
		decimal.NewFromFloat(4.5),
	}
	expects := []string{
		"2.00",
		"1.50",
		"3.55",
		"4.50",
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for idx, value := range values {
			if W2(value) != expects[idx] {
				b.Error("FAIL")
			}
		}
	}
}
