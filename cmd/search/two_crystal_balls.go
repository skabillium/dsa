package search

import "math"

func TwoCrystalBalls(breaks []bool) int {
	jumpAmountFl := math.Sqrt(float64(len(breaks)))
	jumpAmount := int(jumpAmountFl)

	i := 0
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount
	for j := 0; j < jumpAmount && i < len(breaks); i, j = i+1, j+1 {
		if breaks[i] {
			return i
		}
	}

	return -1
}
