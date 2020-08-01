package forward

import "NestedScribbleBenchmark/ring/messages/forward"

type S_Chan struct {
	RingNode_Msg chan forward.Msg
}

type E_Chan struct {
	RingNode_Msg chan forward.Msg
}

type RingNode_Chan struct {
	E_Msg chan forward.Msg
	S_Msg chan forward.Msg
}
