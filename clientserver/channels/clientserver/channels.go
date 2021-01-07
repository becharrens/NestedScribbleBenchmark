package clientserver

import "NestedScribbleBenchmark/clientserver/messages"

type Client_Chan struct {
	Label_From_Server chan messages.ClientServer_Label
	Label_To_Server chan messages.ClientServer_Label
	String_From_Server chan string
	String_To_Server chan string
}

type Server_Chan struct {
	Label_From_Client chan messages.ClientServer_Label
	Label_To_Client chan messages.ClientServer_Label
	String_From_Client chan string
	String_To_Client chan string
}