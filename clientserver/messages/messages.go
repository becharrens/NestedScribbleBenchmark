package messages

type ClientServer_Label int
const (
	DynTaskGen_S ClientServer_Label = iota
	DynTaskGen_Server
	LastReq
	Req
	Resp
)