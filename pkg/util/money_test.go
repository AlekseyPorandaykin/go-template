package util

import (
	"fmt"
	"testing"
)

func TestRoundToPrecision(t *testing.T) {
	f := RoundToPrecision(0.123456789, 5)
	fmt.Println(f)
}
