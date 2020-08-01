package fannkuch

import "NestedScribbleBenchmark/fannkuch/messages/fannkuch"

type Main_Chan struct {
	Worker_Result   chan fannkuch.Result
	Worker_Result_2 chan fannkuch.Result
	Worker_Task     chan fannkuch.Task
}

type Worker_Chan struct {
	Main_Result   chan fannkuch.Result
	Main_Result_2 chan fannkuch.Result
	Main_Task     chan fannkuch.Task
}
