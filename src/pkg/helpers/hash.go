package helpers

import (
	"fmt"
	"math/rand"
)

func GenerateRandomStr() string {
	min := 1000.0
	max := 99999999999.9
	randomFloat := min + rand.Float64()*(max-min)
	return fmt.Sprintf("%.2f", randomFloat)
}
