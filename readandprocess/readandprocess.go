package readandprocess

import (
	"Math-Skills/calculations"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ReadAndProcess(str string) string {
	dataTxt, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening input file!", err)
		os.Exit(1)
	}
	defer dataTxt.Close()

	dataTxtDetails, err := dataTxt.Stat()
	if err != nil {
		fmt.Println("Error getting file details!", err)
		os.Exit(1)
	}

	if dataTxtDetails.Size() == 0 {
		fmt.Println("Error. Input file is empty!")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(dataTxt)

	var valueSlice []float64

	for scanner.Scan() {
		readLine := scanner.Text()

		if readLine == "" {
			continue
		}

		floatValue, err := strconv.ParseFloat(readLine, 64)
		if err != nil {
			fmt.Println("Error converting to float64!", err)
			os.Exit(1)
		}

		if floatValue >= math.MaxInt64 {
			fmt.Println("Error. Float value is out of range!")
			os.Exit(1)
		}

		if floatValue < 0 {
			fmt.Println("Error. Float value should not be a negative value!")
			os.Exit(1)
		}
		valueSlice = append(valueSlice, floatValue)
	}

	if len(valueSlice) == 1 {
		fmt.Println("Error. Statistics cannot be calculated from a single value!")
		os.Exit(1)
	}

	average := calculations.FindAverage(valueSlice)
	median := calculations.FindMedian(valueSlice)
	variance := calculations.FindVariance(valueSlice)
	standardDeviation := calculations.FindStandardDeviation(valueSlice)

	// Round results of the calculations
	roundedAverage := math.Round(average)
	roundedMedian := math.Round(median)
	roundedVariance := math.Round(variance)
	roundedStandardDeviation := math.Round(standardDeviation)

	var statisticsResult string

	// Check if rounded calculations are greater than allowed range. Display statisticsResult in a formatted manner with 0 decimal places.
	if roundedAverage >= math.MaxInt64 {
		statisticsResult += "Error. Average value out of range!\n"
	} else {
		statisticsResult += fmt.Sprintf("Average: %.0f\n", roundedAverage)
	}

	if roundedMedian >= math.MaxInt64 {
		statisticsResult += "Error. Median value out of range!\n"
	} else {
		statisticsResult += fmt.Sprintf("Median: %.0f\n", roundedMedian)
	}

	if roundedVariance >= math.MaxInt64 {
		statisticsResult += "Error. Variance value out of range!\n"
	} else {
		statisticsResult += fmt.Sprintf("Variance: %.0f\n", roundedVariance)
	}

	if roundedStandardDeviation >= math.MaxInt64 {
		statisticsResult += "Error. Standard Deviation value out of range!\n"
	} else {
		statisticsResult += fmt.Sprintf("Standard Deviation: %.0f\n", roundedStandardDeviation)
	}

	return statisticsResult
}
