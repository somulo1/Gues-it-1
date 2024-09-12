package helperfunctions

import "math"

func StandardDeviation(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	mean := CalculateMean(data)
	var varianceSum float64
	for _, value := range data {
		varianceSum += math.Pow(value-mean, 2)//values obtained ^2
	}
	variance := varianceSum / float64(len(data))
	return math.Sqrt(variance)
}
