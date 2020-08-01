package regexredux2

import "NestedScribbleBenchmark/regexredux/messages/regexredux2"

type M_Chan struct {
	W_CalcLength chan regexredux2.CalcLength
	W_Length     chan regexredux2.Length
	W_NumMatches chan regexredux2.NumMatches
	W_Task       chan regexredux2.Task
}

type W_Chan struct {
	M_CalcLength chan regexredux2.CalcLength
	M_Length     chan regexredux2.Length
	M_NumMatches chan regexredux2.NumMatches
	M_Task       chan regexredux2.Task
}
