package simpledns

import "NestedScribbleBenchmark/simpledns/messages"

type App_Chan struct {
	Label_From_ispDNS chan messages.SimpleDNS_Label
	Label_To_ispDNS chan messages.SimpleDNS_Label
	String_From_ispDNS chan string
	String_To_ispDNS chan string
}

type IspDNS_Chan struct {
	Label_From_app chan messages.SimpleDNS_Label
	Label_To_app chan messages.SimpleDNS_Label
	String_From_app chan string
	String_To_app chan string
}