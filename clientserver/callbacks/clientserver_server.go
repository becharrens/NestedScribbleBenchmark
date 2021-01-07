package callbacks

import "NestedScribbleBenchmark/clientserver/results/dyntaskgen"

type ClientServer_Server_Env interface {
	Resp_To_Client_2() string
	ResultFrom_DynTaskGen_S_2(result dyntaskgen.S_Result)
	To_DynTaskGen_S_Env_2() DynTaskGen_S_Env
	DynTaskGen_Setup_2()
	Req_From_Client_2(req string)
	Resp_To_Client() string
	ResultFrom_DynTaskGen_S(result dyntaskgen.S_Result)
	To_DynTaskGen_S_Env() DynTaskGen_S_Env
	DynTaskGen_Setup()
	Req_From_Client(req string)
}

type ClientServerServerState struct {
	Req  string
	Resp string
}

func (c *ClientServerServerState) Resp_To_Client_2() string {
	return c.Resp
}

func (c *ClientServerServerState) ResultFrom_DynTaskGen_S_2(result dyntaskgen.S_Result) {
	c.Resp = result.Resp
}

func (c *ClientServerServerState) To_DynTaskGen_S_Env_2() DynTaskGen_S_Env {
	return &DynTaskGenSState{
		Req:  c.Req,
		Idx:  0,
		Resp: "",
	}
}

func (c *ClientServerServerState) DynTaskGen_Setup_2() {
}

func (c *ClientServerServerState) Req_From_Client_2(req string) {
	c.Req = req
}

func (c *ClientServerServerState) Resp_To_Client() string {
	return c.Resp
}

func (c *ClientServerServerState) ResultFrom_DynTaskGen_S(result dyntaskgen.S_Result) {
	c.Resp = result.Resp
}

func (c *ClientServerServerState) To_DynTaskGen_S_Env() DynTaskGen_S_Env {
	return &DynTaskGenSState{
		Req:  c.Req,
		Idx:  0,
		Resp: "",
	}
}

func (c *ClientServerServerState) DynTaskGen_Setup() {
}

func (c *ClientServerServerState) Req_From_Client(req string) {
	c.Req = req
}
