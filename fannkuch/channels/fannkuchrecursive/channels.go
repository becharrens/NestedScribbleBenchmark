package fannkuchrecursive

import "NestedScribbleBenchmark/fannkuch/messages/fannkuchrecursive"

type Source_Chan struct {
	NewWorker_Result   chan fannkuchrecursive.Result
	NewWorker_Result_2 chan fannkuchrecursive.Result
}

type Worker_Chan struct {
	NewWorker_Task chan fannkuchrecursive.Task
}

type NewWorker_Chan struct {
	Source_Result   chan fannkuchrecursive.Result
	Source_Result_2 chan fannkuchrecursive.Result
	Worker_Task     chan fannkuchrecursive.Task
}
