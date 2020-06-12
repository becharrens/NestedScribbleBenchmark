package ring

import "ScribbleBenchmark/ring/messages/ring"

type Start_Chan struct {
	End_Msg chan ring.Msg
	End_Msg_2 chan ring.Msg
	End_Msg_3 chan ring.Msg
}

type End_Chan struct {
	Start_Msg chan ring.Msg
	Start_Msg_2 chan ring.Msg
	Start_Msg_3 chan ring.Msg
}