package roles

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/generaldns"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import generaldns_2 "NestedScribbleBenchmark/generaldns/results/generaldns"
import "sync"

func GeneralDNS_dnsRes(wg *sync.WaitGroup, roleChannels generaldns.DnsRes_Chan, inviteChannels invitations.GeneralDNS_dnsRes_InviteChan, env callbacks.GeneralDNS_dnsRes_Env) generaldns_2.DnsRes_Result {
	app_choice_4 := <-roleChannels.Label_From_app
	switch app_choice_4 {
	case messages.Done:
		env.Done_From_App()

		return env.Done()
	case messages.Query:
		host := <-roleChannels.String_From_app
		env.Query_From_App(host)

		dnsres_choice_4 := env.DnsRes_Choice()
		switch dnsres_choice_4 {
		case callbacks.GeneralDNS_dnsRes_RecDNSLookup:
			env.RecDNSLookup_Setup()
			
			recdnslookup_rolechan := invitations.RecDNSLookup_RoleSetupChan{
				Res_Chan: inviteChannels.Invite_DnsRes_To_RecDNSLookup_res,
			}
			recdnslookup_invitechan := invitations.RecDNSLookup_InviteSetupChan{
				Res_InviteChan: inviteChannels.Invite_DnsRes_To_RecDNSLookup_res_InviteChan,
			}
			RecDNSLookup_SendCommChannels(wg, recdnslookup_rolechan, recdnslookup_invitechan)

			recdnslookup_res_chan := <-inviteChannels.Invite_DnsRes_To_RecDNSLookup_res
			recdnslookup_res_inviteChan := <-inviteChannels.Invite_DnsRes_To_RecDNSLookup_res_InviteChan
			recdnslookup_res_env := env.To_RecDNSLookup_res_Env()
			recdnslookup_res_result := RecDNSLookup_res(wg, recdnslookup_res_chan, recdnslookup_res_inviteChan, recdnslookup_res_env)
			env.ResultFrom_RecDNSLookup_res(recdnslookup_res_result)

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
				case messages.Query:
					host_2 := <-roleChannels.String_From_app
					env.Query_From_App_2(host_2)

					dnsres_choice := env.DnsRes_Choice_2()
					switch dnsres_choice {
					case callbacks.GeneralDNS_dnsRes_RecDNSLookup_2:
						env.RecDNSLookup_Setup_2()
						
						recdnslookup_rolechan_2 := invitations.RecDNSLookup_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_DnsRes_To_RecDNSLookup_res,
						}
						recdnslookup_invitechan_2 := invitations.RecDNSLookup_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_DnsRes_To_RecDNSLookup_res_InviteChan,
						}
						RecDNSLookup_SendCommChannels(wg, recdnslookup_rolechan_2, recdnslookup_invitechan_2)

						recdnslookup_res_chan_2 := <-inviteChannels.Invite_DnsRes_To_RecDNSLookup_res
						recdnslookup_res_inviteChan_2 := <-inviteChannels.Invite_DnsRes_To_RecDNSLookup_res_InviteChan
						recdnslookup_res_env_2 := env.To_RecDNSLookup_res_Env_2()
						recdnslookup_res_result_2 := RecDNSLookup_res(wg, recdnslookup_res_chan_2, recdnslookup_res_inviteChan_2, recdnslookup_res_env_2)
						env.ResultFrom_RecDNSLookup_res_2(recdnslookup_res_result_2)

						ip_2 := env.IP_To_App_2()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_2

						continue REC
					case callbacks.GeneralDNS_dnsRes_IterDNSLookup_2:
						env.IterDNSLookup_Setup()
						
						iterdnslookup_rolechan := invitations.IterDNSLookup_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_DnsRes_To_IterDNSLookup_res,
						}
						iterdnslookup_invitechan := invitations.IterDNSLookup_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_DnsRes_To_IterDNSLookup_res_InviteChan,
						}
						IterDNSLookup_SendCommChannels(wg, iterdnslookup_rolechan, iterdnslookup_invitechan)

						iterdnslookup_res_chan := <-inviteChannels.Invite_DnsRes_To_IterDNSLookup_res
						iterdnslookup_res_inviteChan := <-inviteChannels.Invite_DnsRes_To_IterDNSLookup_res_InviteChan
						iterdnslookup_res_env := env.To_IterDNSLookup_res_Env()
						iterdnslookup_res_result := IterDNSLookup_res(wg, iterdnslookup_res_chan, iterdnslookup_res_inviteChan, iterdnslookup_res_env)
						env.ResultFrom_IterDNSLookup_res(iterdnslookup_res_result)

						ip_3 := env.IP_To_App_3()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_3

						continue REC
					case callbacks.GeneralDNS_dnsRes_Cached_2:
						env.Cached_Setup()
						
						cached_rolechan := invitations.Cached_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_DnsRes_To_Cached_res,
						}
						cached_invitechan := invitations.Cached_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_DnsRes_To_Cached_res_InviteChan,
						}
						Cached_SendCommChannels(wg, cached_rolechan, cached_invitechan)

						cached_res_chan := <-inviteChannels.Invite_DnsRes_To_Cached_res
						cached_res_inviteChan := <-inviteChannels.Invite_DnsRes_To_Cached_res_InviteChan
						cached_res_env := env.To_Cached_res_Env()
						cached_res_result := Cached_res(wg, cached_res_chan, cached_res_inviteChan, cached_res_env)
						env.ResultFrom_Cached_res(cached_res_result)

						ip_4 := env.IP_To_App_4()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_4

						continue REC
					default:
						panic("Invalid choice was made")
					}
				default:
					panic("Invalid choice was made")
				}
			}
		case callbacks.GeneralDNS_dnsRes_IterDNSLookup:
			env.IterDNSLookup_Setup_2()
			
			iterdnslookup_rolechan_2 := invitations.IterDNSLookup_RoleSetupChan{
				Res_Chan: inviteChannels.Invite_DnsRes_To_IterDNSLookup_res,
			}
			iterdnslookup_invitechan_2 := invitations.IterDNSLookup_InviteSetupChan{
				Res_InviteChan: inviteChannels.Invite_DnsRes_To_IterDNSLookup_res_InviteChan,
			}
			IterDNSLookup_SendCommChannels(wg, iterdnslookup_rolechan_2, iterdnslookup_invitechan_2)

			iterdnslookup_res_chan_2 := <-inviteChannels.Invite_DnsRes_To_IterDNSLookup_res
			iterdnslookup_res_inviteChan_2 := <-inviteChannels.Invite_DnsRes_To_IterDNSLookup_res_InviteChan
			iterdnslookup_res_env_2 := env.To_IterDNSLookup_res_Env_2()
			iterdnslookup_res_result_2 := IterDNSLookup_res(wg, iterdnslookup_res_chan_2, iterdnslookup_res_inviteChan_2, iterdnslookup_res_env_2)
			env.ResultFrom_IterDNSLookup_res_2(iterdnslookup_res_result_2)

			ip_5 := env.IP_To_App_5()
			roleChannels.Label_To_app <- messages.IP
			roleChannels.String_To_app <- ip_5

REC_2:
			for {
				app_choice_2 := <-roleChannels.Label_From_app
				switch app_choice_2 {
				case messages.Done:
					env.Done_From_App_3()

					return env.Done()
				case messages.Query:
					host_3 := <-roleChannels.String_From_app
					env.Query_From_App_3(host_3)

					dnsres_choice_2 := env.DnsRes_Choice_3()
					switch dnsres_choice_2 {
					case callbacks.GeneralDNS_dnsRes_RecDNSLookup_3:
						env.RecDNSLookup_Setup_3()
						
						recdnslookup_rolechan_3 := invitations.RecDNSLookup_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_DnsRes_To_RecDNSLookup_res,
						}
						recdnslookup_invitechan_3 := invitations.RecDNSLookup_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_DnsRes_To_RecDNSLookup_res_InviteChan,
						}
						RecDNSLookup_SendCommChannels(wg, recdnslookup_rolechan_3, recdnslookup_invitechan_3)

						recdnslookup_res_chan_3 := <-inviteChannels.Invite_DnsRes_To_RecDNSLookup_res
						recdnslookup_res_inviteChan_3 := <-inviteChannels.Invite_DnsRes_To_RecDNSLookup_res_InviteChan
						recdnslookup_res_env_3 := env.To_RecDNSLookup_res_Env_3()
						recdnslookup_res_result_3 := RecDNSLookup_res(wg, recdnslookup_res_chan_3, recdnslookup_res_inviteChan_3, recdnslookup_res_env_3)
						env.ResultFrom_RecDNSLookup_res_3(recdnslookup_res_result_3)

						ip_6 := env.IP_To_App_6()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_6

						continue REC_2
					case callbacks.GeneralDNS_dnsRes_IterDNSLookup_3:
						env.IterDNSLookup_Setup_3()
						
						iterdnslookup_rolechan_3 := invitations.IterDNSLookup_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_DnsRes_To_IterDNSLookup_res,
						}
						iterdnslookup_invitechan_3 := invitations.IterDNSLookup_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_DnsRes_To_IterDNSLookup_res_InviteChan,
						}
						IterDNSLookup_SendCommChannels(wg, iterdnslookup_rolechan_3, iterdnslookup_invitechan_3)

						iterdnslookup_res_chan_3 := <-inviteChannels.Invite_DnsRes_To_IterDNSLookup_res
						iterdnslookup_res_inviteChan_3 := <-inviteChannels.Invite_DnsRes_To_IterDNSLookup_res_InviteChan
						iterdnslookup_res_env_3 := env.To_IterDNSLookup_res_Env_3()
						iterdnslookup_res_result_3 := IterDNSLookup_res(wg, iterdnslookup_res_chan_3, iterdnslookup_res_inviteChan_3, iterdnslookup_res_env_3)
						env.ResultFrom_IterDNSLookup_res_3(iterdnslookup_res_result_3)

						ip_7 := env.IP_To_App_7()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_7

						continue REC_2
					case callbacks.GeneralDNS_dnsRes_Cached_3:
						env.Cached_Setup_2()
						
						cached_rolechan_2 := invitations.Cached_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_DnsRes_To_Cached_res,
						}
						cached_invitechan_2 := invitations.Cached_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_DnsRes_To_Cached_res_InviteChan,
						}
						Cached_SendCommChannels(wg, cached_rolechan_2, cached_invitechan_2)

						cached_res_chan_2 := <-inviteChannels.Invite_DnsRes_To_Cached_res
						cached_res_inviteChan_2 := <-inviteChannels.Invite_DnsRes_To_Cached_res_InviteChan
						cached_res_env_2 := env.To_Cached_res_Env_2()
						cached_res_result_2 := Cached_res(wg, cached_res_chan_2, cached_res_inviteChan_2, cached_res_env_2)
						env.ResultFrom_Cached_res_2(cached_res_result_2)

						ip_8 := env.IP_To_App_8()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_8

						continue REC_2
					default:
						panic("Invalid choice was made")
					}
				default:
					panic("Invalid choice was made")
				}
			}
		case callbacks.GeneralDNS_dnsRes_Cached:
			env.Cached_Setup_3()
			
			cached_rolechan_3 := invitations.Cached_RoleSetupChan{
				Res_Chan: inviteChannels.Invite_DnsRes_To_Cached_res,
			}
			cached_invitechan_3 := invitations.Cached_InviteSetupChan{
				Res_InviteChan: inviteChannels.Invite_DnsRes_To_Cached_res_InviteChan,
			}
			Cached_SendCommChannels(wg, cached_rolechan_3, cached_invitechan_3)

			cached_res_chan_3 := <-inviteChannels.Invite_DnsRes_To_Cached_res
			cached_res_inviteChan_3 := <-inviteChannels.Invite_DnsRes_To_Cached_res_InviteChan
			cached_res_env_3 := env.To_Cached_res_Env_3()
			cached_res_result_3 := Cached_res(wg, cached_res_chan_3, cached_res_inviteChan_3, cached_res_env_3)
			env.ResultFrom_Cached_res_3(cached_res_result_3)

			ip_9 := env.IP_To_App_9()
			roleChannels.Label_To_app <- messages.IP
			roleChannels.String_To_app <- ip_9

REC_3:
			for {
				app_choice_3 := <-roleChannels.Label_From_app
				switch app_choice_3 {
				case messages.Done:
					env.Done_From_App_4()

					return env.Done()
				case messages.Query:
					host_4 := <-roleChannels.String_From_app
					env.Query_From_App_4(host_4)

					dnsres_choice_3 := env.DnsRes_Choice_4()
					switch dnsres_choice_3 {
					case callbacks.GeneralDNS_dnsRes_RecDNSLookup_4:
						env.RecDNSLookup_Setup_4()
						
						recdnslookup_rolechan_4 := invitations.RecDNSLookup_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_DnsRes_To_RecDNSLookup_res,
						}
						recdnslookup_invitechan_4 := invitations.RecDNSLookup_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_DnsRes_To_RecDNSLookup_res_InviteChan,
						}
						RecDNSLookup_SendCommChannels(wg, recdnslookup_rolechan_4, recdnslookup_invitechan_4)

						recdnslookup_res_chan_4 := <-inviteChannels.Invite_DnsRes_To_RecDNSLookup_res
						recdnslookup_res_inviteChan_4 := <-inviteChannels.Invite_DnsRes_To_RecDNSLookup_res_InviteChan
						recdnslookup_res_env_4 := env.To_RecDNSLookup_res_Env_4()
						recdnslookup_res_result_4 := RecDNSLookup_res(wg, recdnslookup_res_chan_4, recdnslookup_res_inviteChan_4, recdnslookup_res_env_4)
						env.ResultFrom_RecDNSLookup_res_4(recdnslookup_res_result_4)

						ip_10 := env.IP_To_App_10()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_10

						continue REC_3
					case callbacks.GeneralDNS_dnsRes_IterDNSLookup_4:
						env.IterDNSLookup_Setup_4()
						
						iterdnslookup_rolechan_4 := invitations.IterDNSLookup_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_DnsRes_To_IterDNSLookup_res,
						}
						iterdnslookup_invitechan_4 := invitations.IterDNSLookup_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_DnsRes_To_IterDNSLookup_res_InviteChan,
						}
						IterDNSLookup_SendCommChannels(wg, iterdnslookup_rolechan_4, iterdnslookup_invitechan_4)

						iterdnslookup_res_chan_4 := <-inviteChannels.Invite_DnsRes_To_IterDNSLookup_res
						iterdnslookup_res_inviteChan_4 := <-inviteChannels.Invite_DnsRes_To_IterDNSLookup_res_InviteChan
						iterdnslookup_res_env_4 := env.To_IterDNSLookup_res_Env_4()
						iterdnslookup_res_result_4 := IterDNSLookup_res(wg, iterdnslookup_res_chan_4, iterdnslookup_res_inviteChan_4, iterdnslookup_res_env_4)
						env.ResultFrom_IterDNSLookup_res_4(iterdnslookup_res_result_4)

						ip_11 := env.IP_To_App_11()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_11

						continue REC_3
					case callbacks.GeneralDNS_dnsRes_Cached_4:
						env.Cached_Setup_4()
						
						cached_rolechan_4 := invitations.Cached_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_DnsRes_To_Cached_res,
						}
						cached_invitechan_4 := invitations.Cached_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_DnsRes_To_Cached_res_InviteChan,
						}
						Cached_SendCommChannels(wg, cached_rolechan_4, cached_invitechan_4)

						cached_res_chan_4 := <-inviteChannels.Invite_DnsRes_To_Cached_res
						cached_res_inviteChan_4 := <-inviteChannels.Invite_DnsRes_To_Cached_res_InviteChan
						cached_res_env_4 := env.To_Cached_res_Env_4()
						cached_res_result_4 := Cached_res(wg, cached_res_chan_4, cached_res_inviteChan_4, cached_res_env_4)
						env.ResultFrom_Cached_res_4(cached_res_result_4)

						ip_12 := env.IP_To_App_12()
						roleChannels.Label_To_app <- messages.IP
						roleChannels.String_To_app <- ip_12

						continue REC_3
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