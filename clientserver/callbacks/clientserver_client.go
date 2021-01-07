package callbacks

import "fmt"

type ClientServer_Client_Env interface {
	Resp_From_Server_2(resp string)
	Req_To_Server_2() string
	Resp_From_Server(resp string)
	Req_To_Server() string
}

var requests = []string{"req", "short", "longrequest"}

type ClientServerClientState struct {
	Idx int
}

func (c *ClientServerClientState) Resp_From_Server_2(resp string) {
	fmt.Println("Client: Response from Server -", resp)
}

func (c *ClientServerClientState) Req_To_Server_2() string {
	req := requests[c.Idx]
	c.Idx = (c.Idx + 1) % len(requests)
	fmt.Println("Client: Request to Server -", req)
	return req
}

func (c *ClientServerClientState) Resp_From_Server(resp string) {
	fmt.Println("Client: Response from Server -", resp)
}

func (c *ClientServerClientState) Req_To_Server() string {
	req := requests[c.Idx]
	c.Idx = (c.Idx + 1) % len(requests)
	fmt.Println("Client: Request to Server -", req)
	return req
}
