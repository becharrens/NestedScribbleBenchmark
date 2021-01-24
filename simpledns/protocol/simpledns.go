package protocol

import "NestedScribbleBenchmark/simpledns/messages"
import "NestedScribbleBenchmark/simpledns/channels/simpledns_cached"
import "NestedScribbleBenchmark/simpledns/channels/simpledns"
import "NestedScribbleBenchmark/simpledns/channels/dnslookup"
import "NestedScribbleBenchmark/simpledns/invitations"
import "NestedScribbleBenchmark/simpledns/callbacks"
import simpledns_2 "NestedScribbleBenchmark/simpledns/results/simpledns"
import "NestedScribbleBenchmark/simpledns/roles"
import "sync"

type SimpleDNS_Env interface {
	New_App_Env() callbacks.SimpleDNS_app_Env
	New_IspDNS_Env() callbacks.SimpleDNS_ispDNS_Env
	App_Result(result simpledns_2.App_Result) 
	IspDNS_Result(result simpledns_2.IspDNS_Result) 
}

func Start_SimpleDNS_app(protocolEnv SimpleDNS_Env, wg *sync.WaitGroup, roleChannels simpledns.App_Chan, inviteChannels invitations.SimpleDNS_app_InviteChan, env callbacks.SimpleDNS_app_Env)  {
	defer wg.Done()
	result := roles.SimpleDNS_app(wg, roleChannels, inviteChannels, env)
	protocolEnv.App_Result(result)
} 

func Start_SimpleDNS_ispDNS(protocolEnv SimpleDNS_Env, wg *sync.WaitGroup, roleChannels simpledns.IspDNS_Chan, inviteChannels invitations.SimpleDNS_ispDNS_InviteChan, env callbacks.SimpleDNS_ispDNS_Env)  {
	defer wg.Done()
	result := roles.SimpleDNS_ispDNS(wg, roleChannels, inviteChannels, env)
	protocolEnv.IspDNS_Result(result)
} 

func SimpleDNS(protocolEnv SimpleDNS_Env)  {
	ispdns_invite_ispdns_2 := make(chan dnslookup.Res_Chan, 1)
	ispdns_invite_ispdns_invitechan_2 := make(chan invitations.DNSLookup_res_InviteChan, 1)
	ispdns_app_string := make(chan string, 1)
	ispdns_app_label := make(chan messages.SimpleDNS_Label, 1)
	ispdns_invite_ispdns := make(chan simpledns_cached.Res_Chan, 1)
	ispdns_invite_ispdns_invitechan := make(chan invitations.SimpleDNS_Cached_res_InviteChan, 1)
	app_ispdns_string := make(chan string, 1)
	app_ispdns_label := make(chan messages.SimpleDNS_Label, 1)

	ispdns_chan := simpledns.IspDNS_Chan{
		String_To_app: ispdns_app_string,
		String_From_app: app_ispdns_string,
		Label_To_app: ispdns_app_label,
		Label_From_app: app_ispdns_label,
	}
	app_chan := simpledns.App_Chan{
		String_To_ispDNS: app_ispdns_string,
		String_From_ispDNS: ispdns_app_string,
		Label_To_ispDNS: app_ispdns_label,
		Label_From_ispDNS: ispdns_app_label,
	}

	ispdns_inviteChan := invitations.SimpleDNS_ispDNS_InviteChan{
		Invite_IspDNS_To_SimpleDNS_Cached_res_InviteChan: ispdns_invite_ispdns_invitechan,
		Invite_IspDNS_To_SimpleDNS_Cached_res: ispdns_invite_ispdns,
		Invite_IspDNS_To_DNSLookup_res_InviteChan: ispdns_invite_ispdns_invitechan_2,
		Invite_IspDNS_To_DNSLookup_res: ispdns_invite_ispdns_2,
	}
	app_inviteChan := invitations.SimpleDNS_app_InviteChan{

	}

	var wg sync.WaitGroup

	wg.Add(2)

	app_env := protocolEnv.New_App_Env()
	ispdns_env := protocolEnv.New_IspDNS_Env()

	go Start_SimpleDNS_app(protocolEnv, &wg, app_chan, app_inviteChan, app_env)
	go Start_SimpleDNS_ispDNS(protocolEnv, &wg, ispdns_chan, ispdns_inviteChan, ispdns_env)

	wg.Wait()
} 