package main

import (
	"NestedScribbleBenchmark/benchmark"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func PrintAvgResults(times1 benchmark.BenchmarkTimes, times2 benchmark.BenchmarkTimes) {
	for key := range times1 {
		avg1 := benchmark.Average(times1[key])
		sd1 := benchmark.StandardDeviation(times1[key])
		avg2 := benchmark.Average(times2[key])
		sd2 := benchmark.StandardDeviation(times2[key])
		fmt.Printf("%d - Scribble: %d, %f, Base: %d, %f\n", key, avg1, sd1, avg2, sd2)
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	nIterations := flag.Int("iterations", 1000, "Number of iterations for benchmark")
	minTime := flag.Int("time", 10, "Minimum number of seconds to execute the benchmark for")
	genInputs := flag.Bool("geninputs", false, "Run fasta to generate input files for benchmarks")
	runBoundedFib := flag.Bool("boundedfib", false, "Run benchmark on bounded old_fibonacci protocol")
	runFannkuch := flag.Bool("fannkuch", false, "Run benchmark on fannkuch protocol")
	runSieve := flag.Bool("sieve", false, "Run benchmark on primesieve protocol")
	runRegexRedux := flag.Bool("redux", false, "Run benchmark on regexredux protocol")
	runSpectralNorm := flag.Bool("snorm", false, "Run benchmark on spectralnorm protocol")
	runKNucleotide := flag.Bool("knucl", false, "Run benchmark on quicksort protocol")
	runQuickSort := flag.Bool("quicksort", false, "Run benchmark on quicksort protocol")
	runAll := flag.Bool("all", false, "Run all protocols")
	runUnboundedFib := flag.Bool("ubfib-scr", false, "Run Fibonacci Sequence protocol")
	runUnboundedFibBase := flag.Bool("ubfib-base", false, "Run Fibonacci Sequence protocol")
	flag.Parse()
	benchmark.MinExecTime = *minTime * benchmark.SECOND
	iterations := *nIterations
	numResults := boolToInt(*runFannkuch || *runAll) + boolToInt(*runSieve || *runAll) +
		boolToInt(*runRegexRedux || *runAll) + boolToInt(*runSpectralNorm || *runAll) + boolToInt(*runKNucleotide || *runAll) +
		boolToInt(*runBoundedFib || *runAll)
	strResults := make([]string, 2*numResults)
	if *genInputs {
		GenKNucleotideInputs()
		fmt.Println("Generated knucleotide input files")
		GenRegexInputs()
		fmt.Println("Generated regexredux input files")
		return
	}
	idx := 0

	if *runBoundedFib || *runAll {
		fmt.Println("BoundedFibonacci")
		scribbleResults, baseResults := BoundedFibonacciBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("boundedfibonacci-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("boundedfibonacci-base", baseResults) + "\n;;")
		idx++
	}

	if *runFannkuch || *runAll {
		fmt.Println("Fannkuch")
		scribbleResults, baseResults := FannkuchBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("fannkuch-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("fannkuch-base", baseResults) + "\n;;")
		idx++
	}

	if *runSieve || *runAll {
		fmt.Println("Primesieve")
		scribbleResults, baseResults := PrimeSieveBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("primesieve-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("primesieve-base", baseResults) + "\n;;")
		idx++
	}
	if *runRegexRedux || *runAll {
		fmt.Println("RegexRedux")
		scribbleResults, baseResults := RegexReduxBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("regexredux-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("regexredux-base", baseResults) + "\n;;")
		idx++
	}
	if *runSpectralNorm || *runAll {
		fmt.Println("SpectralNorm")
		scribbleResults, baseResults := SpectralNormBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("spectralnorm-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("spectralnorm-base", baseResults) + "\n;;")
		idx++
	}
	if *runKNucleotide || *runAll {
		fmt.Println("KNucleotide")
		scribbleResults, baseResults := KNucleotideBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("knucleotide-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("knucleotide-base", baseResults) + "\n;;")
		idx++
	}
	if *runQuickSort || *runAll {
		fmt.Println("QuickSort")
		fmt.Printf("\nThreshold search\n\n")
		QSThresholdSearch(iterations)
	}
	result := strings.Join(strResults, "\n")
	err := ioutil.WriteFile("benchmark-results.txt", []byte(result), 0644)
	if err != nil {
		panic("Error while writing to file")
	}

	if *runUnboundedFib {
		RunUboundedFibonacci()
	} else if *runUnboundedFibBase {
		RunUboundedFibonacciBase()
	}
}

// func main() {
// 	// bounded_fib_base.FibSequence()
// bounded_fib_base.Fibonacci(10)
// }
