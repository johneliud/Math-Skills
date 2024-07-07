package calculations

import "math"

func FindAverage(nums []float64) float64 {
	var sum, average float64

	sliceLength := len(nums)

	for i := 0; i < sliceLength; i++ {
		sum += nums[i]
	}
	average = sum / float64(sliceLength)

	return average
}

func FindMedian(nums []float64) float64 {
	var sum, median float64

	sliceLength := len(nums)

	if sliceLength%2 == 1 {
		median = nums[sliceLength/2]
	} else if sliceLength%2 == 0 {
		sum = nums[(sliceLength/2)] + nums[(sliceLength/2)-1]
		median = sum / 2
	}

	return median
}

func FindVariance(nums []float64) float64 {
	var sum, variance float64

	meanOfSlice := FindAverage(nums)
	sliceLength := len(nums)

	for i := 0; i < sliceLength; i++ {
		sum += (nums[i] - meanOfSlice) * (nums[i] - meanOfSlice)
	}
	variance = sum / float64(sliceLength)

	return variance
}

func FindStandardDeviation(nums []float64) float64 {
	standardDeviation := math.Sqrt(FindVariance(nums))

	return standardDeviation
}
