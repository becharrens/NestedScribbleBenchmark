package regexredux

import "NestedScribbleBenchmark/regexredux/messages/regexredux"

type Master_Chan struct {
	Worker_NumMatches chan regexredux.NumMatches
	Worker_Task       chan regexredux.Task
}

type Worker_Chan struct {
	Master_NumMatches chan regexredux.NumMatches
	Master_Task       chan regexredux.Task
}
