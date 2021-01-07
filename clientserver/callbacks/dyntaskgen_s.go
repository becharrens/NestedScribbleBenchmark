package callbacks

import "NestedScribbleBenchmark/clientserver/results/dyntaskgen"

type DynTaskGen_S_Choice int

const (
	DynTaskGen_S_Req DynTaskGen_S_Choice = iota
	DynTaskGen_S_LastReq
)

type DynTaskGen_S_Env interface {
	Resp_From_W_2(resp string)
	LastReq_To_W() string
	Done() dyntaskgen.S_Result
	Resp_From_W(resp string)
	ResultFrom_DynTaskGen_S(result dyntaskgen.S_Result)
	To_DynTaskGen_S_Env() DynTaskGen_S_Env
	DynTaskGen_Setup()
	Req_To_W() string
	S_Choice() DynTaskGen_S_Choice
}

type DynTaskGenSState struct {
	Req  string
	Idx  int
	Resp string
}

func (d *DynTaskGenSState) Resp_From_W_2(resp string) {
	d.Resp = resp + d.Resp
}

func (d *DynTaskGenSState) LastReq_To_W() string {
	req := ""
	if len(d.Req) > 0 {
		req = d.Req[d.Idx : d.Idx+1]
	}
	return req
}

func (d *DynTaskGenSState) Done() dyntaskgen.S_Result {
	return dyntaskgen.S_Result{
		Resp: d.Resp,
	}
}

func (d *DynTaskGenSState) Resp_From_W(resp string) {
	d.Resp = resp + d.Resp
}

func (d *DynTaskGenSState) ResultFrom_DynTaskGen_S(result dyntaskgen.S_Result) {
	// d.Resp = result.Resp
}

func (d *DynTaskGenSState) To_DynTaskGen_S_Env() DynTaskGen_S_Env {
	return d
}

func (d *DynTaskGenSState) DynTaskGen_Setup() {
}

func (d *DynTaskGenSState) Req_To_W() string {
	req := d.Req[d.Idx : d.Idx+1]
	d.Idx++
	return req
}

func (d *DynTaskGenSState) S_Choice() DynTaskGen_S_Choice {
	if d.Idx < len(d.Req)-1 {
		return DynTaskGen_S_Req
	}
	return DynTaskGen_S_LastReq
}
