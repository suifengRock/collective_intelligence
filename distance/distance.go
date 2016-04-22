package distance

import (
	"math"
)

func SameitemInMap(dim1, dim2 map[string]interface{}) map[string]int {
	same := map[string]int{}
	for key, _ := range dim1 {
		if _, ok := dim2[key]; ok {
			same[key] = 1
		}
	}

}

// 欧几里得距离
func EuclideanDistance(dim1, dim2 map[string]float64) float64 {
	same := SameitemInMap(dim1, dim2)
	if len(same) == 0 {
		return 0
	}
	var powSum float64
	for key, _ := range same {
		powSum += math.Pow(dim1[key], 2) + math.Pow(dim2[key], 2)
	}

	return 1 / (1 + math.Sqrt(powSum))
}

// 皮尔逊价值评价
func PearsonScore(dim1, dim2 map[string]float64) float64 {
	same := SameitemInMap(dim1, dim2)
	n := len(same)
	if n == 0 {
		return 0
	}
	var sum1 float64
	var sum2 float64
	var powSum1 float64
	var powSum2 float64
	var pSum float64

	for key, _ := range same {
		sum1 += dim1[key]
		sum2 += dim2[key]
		powSum1 += math.Pow(dim1[key], 2)
		powSum2 += math.Pow(dim2[key], 2)
		pSum += dim1[key] * dim2[key]
	}

	num := pSum - (sum1 * sum2 / n)
	den := math.Sqrt(powSum1-math.Pow(sum1, 2)/n) * (powSum2 - math.Pow(sum2, 2)/n)
	if den == 0 {
		return 0
	}
	return num / den
}
