package roles

import "NestedScribbleBenchmark/spectralnorm/messages/spectralnorm"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_times"
import spectralnorm_2 "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm"
import "NestedScribbleBenchmark/spectralnorm/invitations"
import "sync"

func SpectralNorm_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.SpectralNorm_RoleSetupChan, inviteChannels invitations.SpectralNorm_InviteSetupChan) {
	master_worker_finish := make(chan spectralnorm.Finish, 1)
	master_invite_worker := make(chan spectralnorm_2.Worker_Chan, 1)
	master_invite_worker_invitechan := make(chan invitations.SpectralNorm_Worker_InviteChan, 1)
	master_invite_master_5 := make(chan spectralnorm_2.Master_Chan, 1)
	master_invite_master_invitechan_5 := make(chan invitations.SpectralNorm_Master_InviteChan, 1)
	master_invite_master_4 := make(chan spectralnorm_timestransp.M_Chan, 1)
	master_invite_master_invitechan_4 := make(chan invitations.SpectralNorm_TimesTransp_M_InviteChan, 1)
	master_invite_master_3 := make(chan spectralnorm_times.M_Chan, 1)
	master_invite_master_invitechan_3 := make(chan invitations.SpectralNorm_Times_M_InviteChan, 1)
	master_invite_master_2 := make(chan spectralnorm_timestransp.M_Chan, 1)
	master_invite_master_invitechan_2 := make(chan invitations.SpectralNorm_TimesTransp_M_InviteChan, 1)
	worker_master_timesresult := make(chan spectralnorm.TimesResult, 1)
	master_invite_master := make(chan spectralnorm_times.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.SpectralNorm_Times_M_InviteChan, 1)
	master_worker_timestask := make(chan spectralnorm.TimesTask, 1)

	worker_chan := spectralnorm_2.Worker_Chan{
		Master_TimesTask:   master_worker_timestask,
		Master_TimesResult: worker_master_timesresult,
		Master_Finish:      master_worker_finish,
	}
	master_chan := spectralnorm_2.Master_Chan{
		Worker_TimesTask:   master_worker_timestask,
		Worker_TimesResult: worker_master_timesresult,
		Worker_Finish:      master_worker_finish,
	}

	worker_inviteChan := invitations.SpectralNorm_Worker_InviteChan{
		Master_Invite_To_SpectralNorm_Worker_InviteChan: master_invite_worker_invitechan,
		Master_Invite_To_SpectralNorm_Worker:            master_invite_worker,
	}
	master_inviteChan := invitations.SpectralNorm_Master_InviteChan{
		Invite_Worker_To_SpectralNorm_Worker_InviteChan:          master_invite_worker_invitechan,
		Invite_Worker_To_SpectralNorm_Worker:                     master_invite_worker,
		Invite_Master_To_SpectralNorm_Times_M_InviteChan_2:       master_invite_master_invitechan_3,
		Invite_Master_To_SpectralNorm_Times_M_InviteChan:         master_invite_master_invitechan,
		Invite_Master_To_SpectralNorm_Times_M_2:                  master_invite_master_3,
		Invite_Master_To_SpectralNorm_Times_M:                    master_invite_master,
		Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan_2: master_invite_master_invitechan_4,
		Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan:   master_invite_master_invitechan_2,
		Invite_Master_To_SpectralNorm_TimesTransp_M_2:            master_invite_master_4,
		Invite_Master_To_SpectralNorm_TimesTransp_M:              master_invite_master_2,
		Invite_Master_To_SpectralNorm_Master_InviteChan:          master_invite_master_invitechan_5,
		Invite_Master_To_SpectralNorm_Master:                     master_invite_master_5,
	}

	roleChannels.Master_Chan <- master_chan
	roleChannels.Worker_Chan <- worker_chan

	inviteChannels.Master_InviteChan <- master_inviteChan
	inviteChannels.Worker_InviteChan <- worker_inviteChan
}
