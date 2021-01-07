package callbacks

type DynTaskGen_W_Env interface {
	Resp_To_S_2() string
	LastReq_From_S(req string)
	Done()
	Resp_To_S() string
	Req_From_S(req string)
}

type DynTaskGenWState struct {
	Req string
}

func (d *DynTaskGenWState) Resp_To_S_2() string {
	return d.Req + d.Req
}

func (d *DynTaskGenWState) LastReq_From_S(req string) {
	d.Req = req
}

func (d *DynTaskGenWState) Done() {
}

func (d *DynTaskGenWState) Resp_To_S() string {
	// Assumes req is a string with a single character
	return d.Req + d.Req
}

func (d *DynTaskGenWState) Req_From_S(req string) {
	d.Req = req
}

func New_DynTaskGen_W_State() DynTaskGen_W_Env {
	return &DynTaskGenWState{}
}
