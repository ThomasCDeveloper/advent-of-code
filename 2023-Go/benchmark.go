package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func Benchmark(fn func(), n int) {
	originalStdout := os.Stdout

	defer func() {
		os.Stdout = originalStdout
	}()

	os.Stdout = nil

	times := make([]float64, n)
	var totalTime float64

	for i := 0; i < n; i++ {
		start := time.Now()
		fn()
		elapsed := time.Since(start).Seconds() * 1000
		times[i] = elapsed
		totalTime += elapsed
	}

	os.Stdout = originalStdout

	meanTime := totalTime / float64(n)

	var variance float64
	for _, t := range times {
		variance += (t - meanTime) * (t - meanTime)
	}
	variance /= float64(n)

	stddev := math.Sqrt(variance)

	minTime, maxTime := times[0], times[0]
	for _, t := range times {
		if t < minTime {
			minTime = t
		}
		if t > maxTime {
			maxTime = t
		}
	}

	fmt.Printf("Benchmark Results:\n")
	fmt.Printf("Mean Time: %.3f ms\n", meanTime)
	fmt.Printf("Variance: %.3f ms²\n", variance)
	fmt.Printf("Standard Deviation: %.3f ms\n", stddev)
	fmt.Printf("Min Time: %.3f ms\n", minTime)
	fmt.Printf("Max Time: %.3f ms\n", maxTime)
}
