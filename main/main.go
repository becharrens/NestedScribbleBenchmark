package main

import (
	"ScribbleBenchmark/benchmark"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
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
	iterations := 1
	genInputs := flag.Bool("geninputs", false, "Run fasta to generate input files for benchmarks")
	runFib := flag.Bool("fib", false, "Run benchmark on fibonacci protocol")
	runFannkuch := flag.Bool("fannkuch", false, "Run benchmark on fannkuch protocol")
	runSieve := flag.Bool("sieve", false, "Run benchmark on primesieve protocol")
	runRegexRedux := flag.Bool("redux", false, "Run benchmark on regexredux protocol")
	runSpectralNorm := flag.Bool("snorm", false, "Run benchmark on spectralnorm protocol")
	runKNucleotide := flag.Bool("knucl", false, "Run benchmark on quicksort protocol")
	runQuickSort := flag.Bool("quicksort", false, "Run benchmark on quicksort protocol")

	flag.Parse()
	numResults := boolToInt(*runFib) + boolToInt(*runFannkuch) + boolToInt(*runSieve) +
		boolToInt(*runRegexRedux) + boolToInt(*runSpectralNorm) + boolToInt(*runKNucleotide) +
		boolToInt(*runQuickSort)
	strResults := make([]string, 2*numResults)
	if *genInputs {
		GenKNucleotideInputs()
		fmt.Println("Generated knucleotide input files")
		GenRegexInputs()
		fmt.Println("Generated regexredux input files")
		return
	}
	idx := 0
	if *runFib {
		fmt.Println("Fibonacci")
		scribbleResults, baseResults := FibonacciBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("fibonacci-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("fibonacci-base", baseResults) + "\n;;")
		idx++
	}

	if *runFannkuch {
		fmt.Println("Fannkuch")
		scribbleResults, baseResults := FannkuchBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("fannkuch-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("fannkuch-base", baseResults) + "\n;;")
		idx++
	}

	if *runSieve {
		fmt.Println("Primesieve")
		scribbleResults, baseResults := PrimeSieveBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("primesieve-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("primesieve-base", baseResults) + "\n;;")
		idx++
	}
	if *runRegexRedux {
		fmt.Println("RegexRedux")
		scribbleResults, baseResults := RegexReduxBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("regexredux-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("regexredux-base", baseResults) + "\n;;")
		idx++
	}
	if *runSpectralNorm {
		fmt.Println("SpectralNorm")
		scribbleResults, baseResults := SpectralNormBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("spectralnorm-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("spectralnorm-base", baseResults) + "\n;;")
		idx++
	}
	if *runKNucleotide {
		fmt.Println("KNucleotide")
		scribbleResults, baseResults := KNucleotideBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("quicksort-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("quicksort-base", baseResults) + "\n;;")
		idx++
	}
	if *runQuickSort {
		fmt.Println("QuickSort")
		scribbleResults, baseResults := QuickSortBenchmark(iterations)
		PrintAvgResults(scribbleResults, baseResults)
		strResults[idx] = (benchmark.ResultsToString("quicksort-scribble", scribbleResults) + "\n;;")
		idx++
		strResults[idx] = (benchmark.ResultsToString("quicksort-base", baseResults) + "\n;;")
		idx++
	}
	result := strings.Join(strResults, "\n")
	err := ioutil.WriteFile("benchmark-results.txt", []byte(result), 0644)
	if err != nil {
		panic("Error while writing to file")
	}
}

func PrintResult(n int, res int, chk int) {
	fmt.Printf("%d\nPfannkuchen(%d) = %d\n", chk, n, res)
}

func PrintPrimes(n int, primes []int) {
	strPrimes := make([]string, len(primes))
	for i, prime := range primes {
		strPrimes[i] = strconv.Itoa(prime)
	}
	fmt.Printf("Primes up to %d: %s\n", n, strings.Join(strPrimes, ", "))
}
