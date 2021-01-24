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
		speedup := float64(avg2) / float64(avg1)
		fmt.Printf("%d - Scribble: %d, %f, Base: %d, %f, Ratio: %f\n", key, avg1, sd1, avg2, sd2, speedup)
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	// Benchmark parameters
	nIterations := flag.Int("iterations", 1000, "Number of iterations for benchmark")
	minTime := flag.Int("time", 10, "Minimum number of seconds to execute the benchmark for")

	// Gen inputs for benchmarks
	genInputs := flag.Bool("geninputs", false, "Run fasta to generate input files for benchmarks")

	// Flags for benchmarks
	runBoundedFib := flag.Bool("boundedfib", false, "Run benchmark on bounded fibonacci protocol")
	cmpFibBaselines := flag.Bool("cmp-fib", false, "Compare bounded fibonacci performance against baselines with different degrees of optimisation")
	runFannkuch := flag.Bool("fannkuch", false, "Run benchmark on fannkuch protocol")
	runSieve := flag.Bool("sieve", false, "Run benchmark on primesieve protocol")
	cmpSieveBaselines := flag.Bool("cmp-sieve", false, "Compare primesieve performance against baselines with different degrees of optimisation")
	runRegexRedux := flag.Bool("redux", false, "Run benchmark on regexredux protocol")
	runSpectralNorm := flag.Bool("snorm", false, "Run benchmark on spectralnorm protocol")
	runKNucleotide := flag.Bool("knucl", false, "Run benchmark on quicksort protocol")
	runQuickSort := flag.Bool("quicksort", false, "Run benchmark for quicksort protocol with different thresholds")
	runAll := flag.Bool("all", false, "Run all benchmarks")

	// Flags for running protocols
	runUnboundedFib := flag.Bool("ubfib-scr", false, "Run Fibonacci Sequence protocol")
	runUnboundedFibBase := flag.Bool("ubfib-base", false, "Run Fibonacci Sequence protocol")
	ringSize := flag.Int("ring", 0, "Run Ring protocol of size n")
	runDynTaskGen := flag.Bool("dyntaskgen", false, "Run dynamic task generation protocol")
	runDNS := flag.Bool("dns", false, "Run DNS protocol")
	runSimpleDNS := flag.Bool("simpledns", false, "Run simple DNS protocol")
	runNoughtsAndCrosses := flag.Bool("nc", false, "Run Noughts and Crosses protocol")
	p1AI := flag.Bool("p1-ai", false, "Set player1 in Noughts and Crosses game as a computer AI")
	p2AI := flag.Bool("p2-ai", false, "Set player2 in Noughts and Crosses game as a computer AI")
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
	result := strings.Join(strResults, "\n")
	err := ioutil.WriteFile("benchmark-results.txt", []byte(result), 0644)
	if err != nil {
		panic("Error while writing to file")
	}

	if *cmpSieveBaselines {
		scribbleResults, allBaseResults := ComparePrimeSieveAgainstBaselines(iterations)
		baseResStr := benchmark.ResultsToString("primesieve-scribble", scribbleResults)
		err := ioutil.WriteFile("primesieve-scribble.txt", []byte(baseResStr), 0644)
		if err != nil {
			panic("Error while writing to file")
		}
		for name, baseResults := range allBaseResults {
			fmt.Println(name)
			PrintAvgResults(scribbleResults, baseResults)
			baseResStr := benchmark.ResultsToString(name, baseResults)
			err := ioutil.WriteFile(name+".txt", []byte(baseResStr+"\n;;"), 0644)
			if err != nil {
				panic("Error while writing to file")
			}
		}
		fmt.Println()
	}

	if *cmpFibBaselines {
		scribbleResults, allBaseResults := CompareFibonacciAgainstBaselines(iterations)
		baseResStr := benchmark.ResultsToString("bfib-scribble", scribbleResults)
		err := ioutil.WriteFile("bfib-scribble.txt", []byte(baseResStr), 0644)
		if err != nil {
			panic("Error while writing to file")
		}
		for name, baseResults := range allBaseResults {
			fmt.Println(name)
			PrintAvgResults(scribbleResults, baseResults)
			baseResStr := benchmark.ResultsToString(name, baseResults)
			err := ioutil.WriteFile(name+".txt", []byte(baseResStr+"\n;;"), 0644)
			if err != nil {
				panic("Error while writing to file")
			}
		}
		fmt.Println()
	}

	if *runQuickSort || *runAll {
		fmt.Println("QuickSort")
		fmt.Printf("\nThreshold search\n\n")
		QSThresholdSearch(iterations)
	}

	if *ringSize > 0 {
		if *ringSize < 2 {
			panic("ring size should be >= 2")
		}
		RunRing(*ringSize)
	}

	if *runDNS {
		RunDNS()
	}

	if *runSimpleDNS {
		RunSimpleDNS()
	}

	if *runDynTaskGen {
		RunClientServer()
	}

	if *runNoughtsAndCrosses {
		NoughtsAndCrosses(*p1AI, *p2AI)
	}

	if *runUnboundedFib {
		RunUboundedFibonacci()
	} else if *runUnboundedFibBase {
		RunUboundedFibonacciBase()
	}
}
