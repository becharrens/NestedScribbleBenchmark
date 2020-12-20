package boundedfib

import "NestedScribbleBenchmark/boundedfibonacci/messages"

type Res_Chan struct {
	Int_From_F3   chan int
	Label_From_F3 chan messages.BoundedFibonacci_Label
}

type F1_Chan struct {
	Int_To_F3   chan int
	Label_To_F3 chan messages.BoundedFibonacci_Label
}

type F2_Chan struct {
	Int_To_F3     chan int
	Label_From_F3 chan messages.BoundedFibonacci_Label
	Label_To_F3   chan messages.BoundedFibonacci_Label
}

type F3_Chan struct {
	Int_From_F1   chan int
	Int_From_F2   chan int
	Int_To_Res    chan int
	Label_From_F1 chan messages.BoundedFibonacci_Label
	Label_From_F2 chan messages.BoundedFibonacci_Label
	Label_To_F2   chan messages.BoundedFibonacci_Label
	Label_To_Res  chan messages.BoundedFibonacci_Label
}
