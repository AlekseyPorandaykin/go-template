package util

import "math"

func RoundToPrecision(val float32, precision uint8) float32 {
	base := math.Pow10(int(precision))
	temp := val * float32(base)
	return float32(math.Round(float64(temp))) / float32(base)
}
