package crystalballs

import "math"

func TwoCrystalBalls(arr []bool) int {
	step := int(math.Sqrt(float64(len(arr))))
	i := step
	for i < len(arr) {
		if arr[i] == true {
			break
		}

		i += step
	}

	for j := i - step; j < i; j++ {
		if arr[j] == true {
			return j
		}
	}

	return -1
}
