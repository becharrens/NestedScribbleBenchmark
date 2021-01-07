package main

import (
	"NestedScribbleBenchmark/clientserver/callbacks"
	"NestedScribbleBenchmark/clientserver/protocol"
	clientserver_2 "NestedScribbleBenchmark/clientserver/results/clientserver"
	"fmt"
)

type ClientServerEnv struct {
}

func (c *ClientServerEnv) New_Client_Env() callbacks.ClientServer_Client_Env {
	return &callbacks.ClientServerClientState{
		Idx: 0,
	}
}

func (c *ClientServerEnv) New_Server_Env() callbacks.ClientServer_Server_Env {
	return &callbacks.ClientServerServerState{
		Req:  "",
		Resp: "",
	}
}

func (c *ClientServerEnv) Client_Result(result clientserver_2.Client_Result) {
}

func (c *ClientServerEnv) Server_Result(result clientserver_2.Server_Result) {
}

func NewClientServerEnv() *ClientServerEnv {
	return &ClientServerEnv{}
}

func RunClientServer() {
	fmt.Println("Running Client-Server protocol")
	env := NewClientServerEnv()
	protocol.ClientServer(env)
}
