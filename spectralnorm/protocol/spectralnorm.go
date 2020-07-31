package protocol

import "ScribbleBenchmark/spectralnorm/messages/spectralnorm"
import "ScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"
import "ScribbleBenchmark/spectralnorm/channels/spectralnorm_times"
import spectralnorm_2 "ScribbleBenchmark/spectralnorm/channels/spectralnorm"
import "ScribbleBenchmark/spectralnorm/invitations"
import "ScribbleBenchmark/spectralnorm/callbacks"
import spectralnorm_3 "ScribbleBenchmark/spectralnorm/results/spectralnorm"
import "ScribbleBenchmark/spectralnorm/roles"
import "sync"

type SpectralNorm_Env interface {
	New_Master_Env() callbacks.SpectralNorm_Master_Env
	New_Worker_Env() callbacks.SpectralNorm_Worker_Env
	Master_Result(result spectralnorm_3.Master_Result) 
	Worker_Result(result spectralnorm_3.Worker_Result) 
}

func Start_SpectralNorm_Master(protocolEnv SpectralNorm_Env, wg *sync.WaitGroup, roleChannels spectralnorm_2.Master_Chan, inviteChannels invitations.SpectralNorm_Master_InviteChan, env callbacks.SpectralNorm_Master_Env)  {
	defer wg.Done()
	result := roles.SpectralNorm_Master(wg, roleChannels, inviteChannels, env)
	protocolEnv.Master_Result(result)
} 

func Start_SpectralNorm_Worker(protocolEnv SpectralNorm_Env, wg *sync.WaitGroup, roleChannels spectralnorm_2.Worker_Chan, inviteChannels invitations.SpectralNorm_Worker_InviteChan, env callbacks.SpectralNorm_Worker_Env)  {
	defer wg.Done()
	result := roles.SpectralNorm_Worker(wg, roleChannels, inviteChannels, env)
	protocolEnv.Worker_Result(result)
} 

func SpectralNorm(protocolEnv SpectralNorm_Env)  {
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
		Master_TimesTask: master_worker_timestask,
		Master_TimesResult: worker_master_timesresult,
		Master_Finish: master_worker_finish,
	}
	master_chan := spectralnorm_2.Master_Chan{
		Worker_TimesTask: master_worker_timestask,
		Worker_TimesResult: worker_master_timesresult,
		Worker_Finish: master_worker_finish,
	}

	worker_inviteChan := invitations.SpectralNorm_Worker_InviteChan{
		Master_Invite_To_SpectralNorm_Worker_InviteChan: master_invite_worker_invitechan,
		Master_Invite_To_SpectralNorm_Worker: master_invite_worker,
	}
	master_inviteChan := invitations.SpectralNorm_Master_InviteChan{
		Invite_Worker_To_SpectralNorm_Worker_InviteChan: master_invite_worker_invitechan,
		Invite_Worker_To_SpectralNorm_Worker: master_invite_worker,
		Invite_Master_To_SpectralNorm_Times_M_InviteChan_2: master_invite_master_invitechan_3,
		Invite_Master_To_SpectralNorm_Times_M_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_SpectralNorm_Times_M_2: master_invite_master_3,
		Invite_Master_To_SpectralNorm_Times_M: master_invite_master,
		Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan_2: master_invite_master_invitechan_4,
		Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan: master_invite_master_invitechan_2,
		Invite_Master_To_SpectralNorm_TimesTransp_M_2: master_invite_master_4,
		Invite_Master_To_SpectralNorm_TimesTransp_M: master_invite_master_2,
		Invite_Master_To_SpectralNorm_Master_InviteChan: master_invite_master_invitechan_5,
		Invite_Master_To_SpectralNorm_Master: master_invite_master_5,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	master_env := protocolEnv.New_Master_Env()
	worker_env := protocolEnv.New_Worker_Env()

	go Start_SpectralNorm_Master(protocolEnv, &wg, master_chan, master_inviteChan, master_env)
	go Start_SpectralNorm_Worker(protocolEnv, &wg, worker_chan, worker_inviteChan, worker_env)

	wg.Wait()
} 