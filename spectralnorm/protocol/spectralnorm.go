package protocol

import "NestedScribbleBenchmark/spectralnorm/messages"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_times"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm"
import "NestedScribbleBenchmark/spectralnorm/invitations"
import "NestedScribbleBenchmark/spectralnorm/callbacks"
import spectralnorm_2 "NestedScribbleBenchmark/spectralnorm/results/spectralnorm"
import "NestedScribbleBenchmark/spectralnorm/roles"
import "sync"

type SpectralNorm_Env interface {
	New_Master_Env() callbacks.SpectralNorm_Master_Env
	New_Worker_Env() callbacks.SpectralNorm_Worker_Env
	Master_Result(result spectralnorm_2.Master_Result)
	Worker_Result(result spectralnorm_2.Worker_Result)
}

func Start_SpectralNorm_Master(protocolEnv SpectralNorm_Env, wg *sync.WaitGroup, roleChannels spectralnorm.Master_Chan, inviteChannels invitations.SpectralNorm_Master_InviteChan, env callbacks.SpectralNorm_Master_Env) {
	defer wg.Done()
	result := roles.SpectralNorm_Master(wg, roleChannels, inviteChannels, env)
	protocolEnv.Master_Result(result)
}

func Start_SpectralNorm_Worker(protocolEnv SpectralNorm_Env, wg *sync.WaitGroup, roleChannels spectralnorm.Worker_Chan, inviteChannels invitations.SpectralNorm_Worker_InviteChan, env callbacks.SpectralNorm_Worker_Env) {
	defer wg.Done()
	result := roles.SpectralNorm_Worker(wg, roleChannels, inviteChannels, env)
	protocolEnv.Worker_Result(result)
}

func SpectralNorm(protocolEnv SpectralNorm_Env) {
	master_invite_worker := make(chan spectralnorm.Worker_Chan, 1)
	master_invite_worker_invitechan := make(chan invitations.SpectralNorm_Worker_InviteChan, 1)
	master_invite_master_3 := make(chan spectralnorm.Master_Chan, 1)
	master_invite_master_invitechan_3 := make(chan invitations.SpectralNorm_Master_InviteChan, 1)
	master_invite_master_2 := make(chan spectralnorm_timestransp.M_Chan, 1)
	master_invite_master_invitechan_2 := make(chan invitations.SpectralNorm_TimesTransp_M_InviteChan, 1)
	worker_master_vec := make(chan []float64, 1)
	worker_master_label := make(chan messages.SpectralNorm_Label, 1)
	master_invite_master := make(chan spectralnorm_times.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.SpectralNorm_Times_M_InviteChan, 1)
	master_worker_vec := make(chan []float64, 1)
	master_worker_int := make(chan int, 1)
	master_worker_label := make(chan messages.SpectralNorm_Label, 1)

	worker_chan := spectralnorm.Worker_Chan{
		Vec_To_Master:     worker_master_vec,
		Vec_From_Master:   master_worker_vec,
		Label_To_Master:   worker_master_label,
		Label_From_Master: master_worker_label,
		Int_From_Master:   master_worker_int,
	}
	master_chan := spectralnorm.Master_Chan{
		Vec_To_Worker:     master_worker_vec,
		Vec_From_Worker:   worker_master_vec,
		Label_To_Worker:   master_worker_label,
		Label_From_Worker: worker_master_label,
		Int_To_Worker:     master_worker_int,
	}

	worker_inviteChan := invitations.SpectralNorm_Worker_InviteChan{
		Master_Invite_To_SpectralNorm_Worker_InviteChan: master_invite_worker_invitechan,
		Master_Invite_To_SpectralNorm_Worker:            master_invite_worker,
	}
	master_inviteChan := invitations.SpectralNorm_Master_InviteChan{
		Invite_Worker_To_SpectralNorm_Worker_InviteChan:        master_invite_worker_invitechan,
		Invite_Worker_To_SpectralNorm_Worker:                   master_invite_worker,
		Invite_Master_To_SpectralNorm_Times_M_InviteChan:       master_invite_master_invitechan,
		Invite_Master_To_SpectralNorm_Times_M:                  master_invite_master,
		Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan: master_invite_master_invitechan_2,
		Invite_Master_To_SpectralNorm_TimesTransp_M:            master_invite_master_2,
		Invite_Master_To_SpectralNorm_Master_InviteChan:        master_invite_master_invitechan_3,
		Invite_Master_To_SpectralNorm_Master:                   master_invite_master_3,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	master_env := protocolEnv.New_Master_Env()
	worker_env := protocolEnv.New_Worker_Env()

	go Start_SpectralNorm_Master(protocolEnv, &wg, master_chan, master_inviteChan, master_env)
	go Start_SpectralNorm_Worker(protocolEnv, &wg, worker_chan, worker_inviteChan, worker_env)

	wg.Wait()
}
