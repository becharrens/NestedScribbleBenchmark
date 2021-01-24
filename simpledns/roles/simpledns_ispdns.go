package roles

import "NestedScribbleBenchmark/simpledns/messages"
import "NestedScribbleBenchmark/simpledns/channels/simpledns"
import "NestedScribbleBenchmark/simpledns/invitations"
import "NestedScribbleBenchmark/simpledns/callbacks"
import simpledns_2 "NestedScribbleBenchmark/simpledns/results/simpledns"
import "sync"

func SimpleDNS_ispDNS(wg *sync.WaitGroup, roleChannels simpledns.IspDNS_Chan, inviteChannels invitations.SimpleDNS_ispDNS_InviteChan, env callbacks.SimpleDNS_ispDNS_Env) simpledns_2.IspDNS_Result {
	app_choice_3 := <-roleChannels.Label_From_app
	switch app_choice_3 {
	case messages.Done:
		env.Done_From_App()

		return env.Done()
	case messages.RecQuery:
		host := <-roleChannels.String_From_app
		env.RecQuery_From_App(host)

		ispdns_choice_3 := env.IspDNS_Choice()
		switch ispdns_choice_3 {
		case callbacks.SimpleDNS_ispDNS_SimpleDNS_Cached:
			env.SimpleDNS_Cached_Setup()
			
			simpledns_cached_rolechan := invitations.SimpleDNS_Cached_RoleSetupChan{
				Res_Chan: inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res,
			}
			simpledns_cached_invitechan := invitations.SimpleDNS_Cached_InviteSetupChan{
				Res_InviteChan: inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res_InviteChan,
			}
			SimpleDNS_Cached_SendCommChannels(wg, simpledns_cached_rolechan, simpledns_cached_invitechan)

			simpledns_cached_res_chan := <-inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res
			simpledns_cached_res_inviteChan := <-inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res_InviteChan
			simpledns_cached_res_env := env.To_SimpleDNS_Cached_res_Env()
			simpledns_cached_res_result := SimpleDNS_Cached_res(wg, simpledns_cached_res_chan, simpledns_cached_res_inviteChan, simpledns_cached_res_env)
			env.ResultFrom_SimpleDNS_Cached_res(simpledns_cached_res_result)

			ip := env.IP_To_App()
			roleChannels.Label_To_app <- messages.IP
			roleChannels.String_To_app <- ip

REC:
			for {
				app_choice := <-roleChannels.Label_From_app
				switch app_choice {
				case messages.Done:
					env.Done_From_App_2()

					return env.Done()
				case messages.RecQuery:
					host_2 := <-roleChannels.String_From_app
					env.RecQuery_From_App_2(host_2)

					ispdns_choice := env.IspDNS_Choice_2()
					switch ispdns_choice {
					case callbacks.SimpleDNS_ispDNS_SimpleDNS_Cached_2:
						env.SimpleDNS_Cached_Setup_2()
						
						simpledns_cached_rolechan_2 := invitations.SimpleDNS_Cached_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res,
						}
						simpledns_cached_invitechan_2 := invitations.SimpleDNS_Cached_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res_InviteChan,
						}
						SimpleDNS_Cached_SendCommChannels(wg, simpledns_cached_rolechan_2, simpledns_cached_invitechan_2)

						simpledns_cached_res_chan_2 := <-inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res
						simpledns_cached_res_inviteChan_2 := <-inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res_InviteChan
						simpledns_cached_res_env_2 := env.To_SimpleDNS_Cached_res_Env_2()
						simpledns_cached_res_result_2 := SimpleDNS_Cached_res(wg, simpledns_cached_res_chan_2, simpledns_cached_res_inviteChan_2, simpledns_cached_res_env_2)
						env.ResultFrom_SimpleDNS_Cached_res_2(simpledns_cached_res_result_2)

						ip_2 := env.IP_To_App_2()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_2

						continue REC
					case callbacks.SimpleDNS_ispDNS_DNSLookup_2:
						env.DNSLookup_Setup()
						
						dnslookup_rolechan := invitations.DNSLookup_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_IspDNS_To_DNSLookup_res,
						}
						dnslookup_invitechan := invitations.DNSLookup_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_IspDNS_To_DNSLookup_res_InviteChan,
						}
						DNSLookup_SendCommChannels(wg, dnslookup_rolechan, dnslookup_invitechan)

						dnslookup_res_chan := <-inviteChannels.Invite_IspDNS_To_DNSLookup_res
						dnslookup_res_inviteChan := <-inviteChannels.Invite_IspDNS_To_DNSLookup_res_InviteChan
						dnslookup_res_env := env.To_DNSLookup_res_Env()
						dnslookup_res_result := DNSLookup_res(wg, dnslookup_res_chan, dnslookup_res_inviteChan, dnslookup_res_env)
						env.ResultFrom_DNSLookup_res(dnslookup_res_result)

						ip_3 := env.IP_To_App_3()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_3

						continue REC
					default:
						panic("Invalid choice was made")
					}
				default:
					panic("Invalid choice was made")
				}
			}
		case callbacks.SimpleDNS_ispDNS_DNSLookup:
			env.DNSLookup_Setup_2()
			
			dnslookup_rolechan_2 := invitations.DNSLookup_RoleSetupChan{
				Res_Chan: inviteChannels.Invite_IspDNS_To_DNSLookup_res,
			}
			dnslookup_invitechan_2 := invitations.DNSLookup_InviteSetupChan{
				Res_InviteChan: inviteChannels.Invite_IspDNS_To_DNSLookup_res_InviteChan,
			}
			DNSLookup_SendCommChannels(wg, dnslookup_rolechan_2, dnslookup_invitechan_2)

			dnslookup_res_chan_2 := <-inviteChannels.Invite_IspDNS_To_DNSLookup_res
			dnslookup_res_inviteChan_2 := <-inviteChannels.Invite_IspDNS_To_DNSLookup_res_InviteChan
			dnslookup_res_env_2 := env.To_DNSLookup_res_Env_2()
			dnslookup_res_result_2 := DNSLookup_res(wg, dnslookup_res_chan_2, dnslookup_res_inviteChan_2, dnslookup_res_env_2)
			env.ResultFrom_DNSLookup_res_2(dnslookup_res_result_2)

			ip_4 := env.IP_To_App_4()
			roleChannels.Label_To_app <- messages.IP
			roleChannels.String_To_app <- ip_4

REC_2:
			for {
				app_choice_2 := <-roleChannels.Label_From_app
				switch app_choice_2 {
				case messages.Done:
					env.Done_From_App_3()

					return env.Done()
				case messages.RecQuery:
					host_3 := <-roleChannels.String_From_app
					env.RecQuery_From_App_3(host_3)

					ispdns_choice_2 := env.IspDNS_Choice_3()
					switch ispdns_choice_2 {
					case callbacks.SimpleDNS_ispDNS_SimpleDNS_Cached_3:
						env.SimpleDNS_Cached_Setup_3()
						
						simpledns_cached_rolechan_3 := invitations.SimpleDNS_Cached_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res,
						}
						simpledns_cached_invitechan_3 := invitations.SimpleDNS_Cached_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res_InviteChan,
						}
						SimpleDNS_Cached_SendCommChannels(wg, simpledns_cached_rolechan_3, simpledns_cached_invitechan_3)

						simpledns_cached_res_chan_3 := <-inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res
						simpledns_cached_res_inviteChan_3 := <-inviteChannels.Invite_IspDNS_To_SimpleDNS_Cached_res_InviteChan
						simpledns_cached_res_env_3 := env.To_SimpleDNS_Cached_res_Env_3()
						simpledns_cached_res_result_3 := SimpleDNS_Cached_res(wg, simpledns_cached_res_chan_3, simpledns_cached_res_inviteChan_3, simpledns_cached_res_env_3)
						env.ResultFrom_SimpleDNS_Cached_res_3(simpledns_cached_res_result_3)

						ip_5 := env.IP_To_App_5()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_5

						continue REC_2
					case callbacks.SimpleDNS_ispDNS_DNSLookup_3:
						env.DNSLookup_Setup_3()
						
						dnslookup_rolechan_3 := invitations.DNSLookup_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_IspDNS_To_DNSLookup_res,
						}
						dnslookup_invitechan_3 := invitations.DNSLookup_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_IspDNS_To_DNSLookup_res_InviteChan,
						}
						DNSLookup_SendCommChannels(wg, dnslookup_rolechan_3, dnslookup_invitechan_3)

						dnslookup_res_chan_3 := <-inviteChannels.Invite_IspDNS_To_DNSLookup_res
						dnslookup_res_inviteChan_3 := <-inviteChannels.Invite_IspDNS_To_DNSLookup_res_InviteChan
						dnslookup_res_env_3 := env.To_DNSLookup_res_Env_3()
						dnslookup_res_result_3 := DNSLookup_res(wg, dnslookup_res_chan_3, dnslookup_res_inviteChan_3, dnslookup_res_env_3)
						env.ResultFrom_DNSLookup_res_3(dnslookup_res_result_3)

						ip_6 := env.IP_To_App_6()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_6

						continue REC_2
					default:
						panic("Invalid choice was made")
					}
				default:
					panic("Invalid choice was made")
				}
			}
		default:
			panic("Invalid choice was made")
		}
	default:
		panic("Invalid choice was made")
	}
} 