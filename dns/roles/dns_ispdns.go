package roles

import "NestedScribbleBenchmark/dns/messages"
import "NestedScribbleBenchmark/dns/channels/dns"
import "NestedScribbleBenchmark/dns/invitations"
import "NestedScribbleBenchmark/dns/callbacks"
import dns_2 "NestedScribbleBenchmark/dns/results/dns"
import "sync"

func DNS_ispDNS(wg *sync.WaitGroup, roleChannels dns.IspDNS_Chan, inviteChannels invitations.DNS_ispDNS_InviteChan, env callbacks.DNS_ispDNS_Env) dns_2.IspDNS_Result {
	dnsres_choice_3 := <-roleChannels.Label_From_dnsRes
	switch dnsres_choice_3 {
	case messages.RecQuery:
		host := <-roleChannels.String_From_dnsRes
		env.RecQuery_From_DnsRes(host)

		ispdns_choice_3 := env.IspDNS_Choice()
		switch ispdns_choice_3 {
		case callbacks.DNS_ispDNS_DNS_Cached:
			env.DNS_Cached_Setup()
			
			dns_cached_rolechan := invitations.DNS_Cached_RoleSetupChan{
				Res_Chan: inviteChannels.Invite_IspDNS_To_DNS_Cached_res,
			}
			dns_cached_invitechan := invitations.DNS_Cached_InviteSetupChan{
				Res_InviteChan: inviteChannels.Invite_IspDNS_To_DNS_Cached_res_InviteChan,
			}
			DNS_Cached_SendCommChannels(wg, dns_cached_rolechan, dns_cached_invitechan)

			dns_cached_res_chan := <-inviteChannels.Invite_IspDNS_To_DNS_Cached_res
			dns_cached_res_inviteChan := <-inviteChannels.Invite_IspDNS_To_DNS_Cached_res_InviteChan
			dns_cached_res_env := env.To_DNS_Cached_res_Env()
			dns_cached_res_result := DNS_Cached_res(wg, dns_cached_res_chan, dns_cached_res_inviteChan, dns_cached_res_env)
			env.ResultFrom_DNS_Cached_res(dns_cached_res_result)

			ip := env.IP_To_DnsRes()
			roleChannels.Label_To_dnsRes <- messages.IP
			roleChannels.String_To_dnsRes <- ip

REC:
			for {
				dnsres_choice := <-roleChannels.Label_From_dnsRes
				switch dnsres_choice {
				case messages.RecQuery:
					host_2 := <-roleChannels.String_From_dnsRes
					env.RecQuery_From_DnsRes_2(host_2)

					ispdns_choice := env.IspDNS_Choice_2()
					switch ispdns_choice {
					case callbacks.DNS_ispDNS_DNS_Cached_2:
						env.DNS_Cached_Setup_2()
						
						dns_cached_rolechan_2 := invitations.DNS_Cached_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_IspDNS_To_DNS_Cached_res,
						}
						dns_cached_invitechan_2 := invitations.DNS_Cached_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_IspDNS_To_DNS_Cached_res_InviteChan,
						}
						DNS_Cached_SendCommChannels(wg, dns_cached_rolechan_2, dns_cached_invitechan_2)

						dns_cached_res_chan_2 := <-inviteChannels.Invite_IspDNS_To_DNS_Cached_res
						dns_cached_res_inviteChan_2 := <-inviteChannels.Invite_IspDNS_To_DNS_Cached_res_InviteChan
						dns_cached_res_env_2 := env.To_DNS_Cached_res_Env_2()
						dns_cached_res_result_2 := DNS_Cached_res(wg, dns_cached_res_chan_2, dns_cached_res_inviteChan_2, dns_cached_res_env_2)
						env.ResultFrom_DNS_Cached_res_2(dns_cached_res_result_2)

						ip_2 := env.IP_To_DnsRes_2()
						roleChannels.Label_To_dnsRes <- messages.IP
						roleChannels.String_To_dnsRes <- ip_2

						continue REC
					case callbacks.DNS_ispDNS_DNSLookup_2:
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

						ip_3 := env.IP_To_DnsRes_3()
						roleChannels.Label_To_dnsRes <- messages.IP
						roleChannels.String_To_dnsRes <- ip_3

						continue REC
					default:
						panic("Invalid choice was made")
					}
				case messages.Done:
					env.Done_From_DnsRes()

					return env.Done()
				default:
					panic("Invalid choice was made")
				}
			}
		case callbacks.DNS_ispDNS_DNSLookup:
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

			ip_4 := env.IP_To_DnsRes_4()
			roleChannels.Label_To_dnsRes <- messages.IP
			roleChannels.String_To_dnsRes <- ip_4

REC_2:
			for {
				dnsres_choice_2 := <-roleChannels.Label_From_dnsRes
				switch dnsres_choice_2 {
				case messages.RecQuery:
					host_3 := <-roleChannels.String_From_dnsRes
					env.RecQuery_From_DnsRes_3(host_3)

					ispdns_choice_2 := env.IspDNS_Choice_3()
					switch ispdns_choice_2 {
					case callbacks.DNS_ispDNS_DNS_Cached_3:
						env.DNS_Cached_Setup_3()
						
						dns_cached_rolechan_3 := invitations.DNS_Cached_RoleSetupChan{
							Res_Chan: inviteChannels.Invite_IspDNS_To_DNS_Cached_res,
						}
						dns_cached_invitechan_3 := invitations.DNS_Cached_InviteSetupChan{
							Res_InviteChan: inviteChannels.Invite_IspDNS_To_DNS_Cached_res_InviteChan,
						}
						DNS_Cached_SendCommChannels(wg, dns_cached_rolechan_3, dns_cached_invitechan_3)

						dns_cached_res_chan_3 := <-inviteChannels.Invite_IspDNS_To_DNS_Cached_res
						dns_cached_res_inviteChan_3 := <-inviteChannels.Invite_IspDNS_To_DNS_Cached_res_InviteChan
						dns_cached_res_env_3 := env.To_DNS_Cached_res_Env_3()
						dns_cached_res_result_3 := DNS_Cached_res(wg, dns_cached_res_chan_3, dns_cached_res_inviteChan_3, dns_cached_res_env_3)
						env.ResultFrom_DNS_Cached_res_3(dns_cached_res_result_3)

						ip_5 := env.IP_To_DnsRes_5()
						roleChannels.Label_To_dnsRes <- messages.IP
						roleChannels.String_To_dnsRes <- ip_5

						continue REC_2
					case callbacks.DNS_ispDNS_DNSLookup_3:
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

						ip_6 := env.IP_To_DnsRes_6()
						roleChannels.Label_To_dnsRes <- messages.IP
						roleChannels.String_To_dnsRes <- ip_6

						continue REC_2
					default:
						panic("Invalid choice was made")
					}
				case messages.Done:
					env.Done_From_DnsRes_2()

					return env.Done()
				default:
					panic("Invalid choice was made")
				}
			}
		default:
			panic("Invalid choice was made")
		}
	case messages.Done:
		env.Done_From_DnsRes_3()

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 