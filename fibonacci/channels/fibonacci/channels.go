package fibonacci

import "NestedScribbleBenchmark/fibonacci/messages"

type Start_Chan struct {
	Int_To_F1 chan int
	Int_To_F2 chan int
	Label_To_F1 chan messages.Fibonacci_Label
	Label_To_F2 chan messages.Fibonacci_Label
}

type F1_Chan struct {
	Int_From_Start chan int
	Label_From_Start chan messages.Fibonacci_Label
}

type F2_Chan struct {
	Int_From_Start chan int
	Label_From_Start chan messages.Fibonacci_Label
}