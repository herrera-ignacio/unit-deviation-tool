package main

import (
	"conversions"
	"fmt"
	"math"
)

func testDeviation(transform conversions.Convert, revert conversions.Convert, start float64, limit float64, interval float64) {
	maxDeviation := 0.0
	for i := start; i < limit; i += interval {
		firstConversion := transform(i)
		secondConversion := revert(firstConversion)
		deviation := (secondConversion - i) / i * 100
		maxDeviation = math.Max(maxDeviation, math.Abs(deviation))
		// fmt.Printf("%e(A) to %f(B) back to %e(A) - Deviation %e%%\n", i, firstConversion, secondConversion, deviation)
	}
	fmt.Printf("Max deviation: +/- %e%%", maxDeviation)
}

func MetersToMilesDeviation(start float64, limit float64, interval float64) {
	testDeviation(
		conversions.MetersToMiles,
		conversions.MilesToMeters,
		start,
		limit,
		interval,
	)
}

func LitersToGallonsDeviation(start float64, limit float64, interval float64) {
	testDeviation(
		conversions.LitersToGallons,
		conversions.GallonsToLiters,
		start,
		limit,
		interval,
	)
}
