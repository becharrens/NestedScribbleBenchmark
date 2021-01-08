package fib

import "NestedScribbleBenchmark/fibonacci/messages"

type Res_Chan struct {
	Int_From_F3 chan int
	Label_From_F3 chan messages.Fibonacci_Label
}

type F1_Chan struct {
	Int_To_F3 chan int
	Label_To_F3 chan messages.Fibonacci_Label
}

type F2_Chan struct {
	Int_To_F3 chan int
	Label_From_F3 chan messages.Fibonacci_Label
	Label_To_F3 chan messages.Fibonacci_Label
}

type F3_Chan struct {
	Int_From_F1 chan int
	Int_From_F2 chan int
	Int_To_Res chan int
	Label_From_F1 chan messages.Fibonacci_Label
	Label_From_F2 chan messages.Fibonacci_Label
	Label_To_F2 chan messages.Fibonacci_Label
	Label_To_Res chan messages.Fibonacci_Label
}