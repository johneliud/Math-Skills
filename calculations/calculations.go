package calculations

import "math"

func FindAverage(num []float64) float64 {
	var sum, average float64

	sliceLength := len(num)

	for i := 0; i < sliceLength; i++ {
		sum += num[i]
	}
	average = sum / float64(sliceLength)

	return average
}

func FindMedian(num []float64) float64 {
	var sum, median float64

	sliceLength := len(num)

	if sliceLength%2 == 1 {
		median = num[sliceLength/2]
	} else if sliceLength%2 == 0 {
		sum = num[(sliceLength/2)] + num[(sliceLength/2)-1]
		median = sum / 2
	}

	return median
}

func FindVariance(num []float64) float64 {
	var sum, variance float64

	meanOfSlice := FindAverage(num)
	sliceLength := len(num)

	for i := 0; i < sliceLength; i++ {
		sum += (num[i] - meanOfSlice) * (num[i] - meanOfSlice)
	}
	variance = sum / float64(sliceLength)

	return variance
}

func FindStandardDeviation(num float64) float64 {
	standardDeviation := math.Sqrt(num)

	return standardDeviation
}
