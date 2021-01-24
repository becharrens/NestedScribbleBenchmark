package messages

type Ring_Label int

const (
	Forward_RingNode_E Ring_Label = iota
	Forward_Start_End
	Msg
)
