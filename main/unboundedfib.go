package main

import (
	"NestedScribbleBenchmark/fibonacci_base"
	"NestedScribbleBenchmark/fibonacci_prev/callbacks"
	"NestedScribbleBenchmark/fibonacci_prev/protocol"
	fibonacci_2 "NestedScribbleBenchmark/fibonacci_prev/results/fibonacci"
)

type UnboundedFibEnv struct {
}

func (f *UnboundedFibEnv) New_Start_Env() callbacks.Fibonacci_Start_Env {
	return &callbacks.FibonacciStartState{}
}

func (f *UnboundedFibEnv) New_F1_Env() callbacks.Fibonacci_F1_Env {
	return &callbacks.FibonacciF1State{}
}

func (f *UnboundedFibEnv) New_F2_Env() callbacks.Fibonacci_F2_Env {
	return &callbacks.FibonacciF2State{}
}

func (f *UnboundedFibEnv) Start_Result(result fibonacci_2.Start_Result) {
}

func (f *UnboundedFibEnv) F1_Result(result fibonacci_2.F1_Result) {
}

func (f *UnboundedFibEnv) F2_Result(result fibonacci_2.F2_Result) {
}

func NewUnboundedFibEnv() *UnboundedFibEnv {
	return &UnboundedFibEnv{}
}

func RunUboundedFibonacci() {
	env := NewUnboundedFibEnv()
	protocol.Fibonacci(env)
}

func RunUboundedFibonacciBase() {
	fibonacci_base.FibSequence()
}
