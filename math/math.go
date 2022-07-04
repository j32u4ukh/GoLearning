package math

func Mean(data []float64) (mean float64) {
	var sum float64
	for _, v := range data {
		sum += v
	}
	mean = sum / float64(len(data))
	return
}
