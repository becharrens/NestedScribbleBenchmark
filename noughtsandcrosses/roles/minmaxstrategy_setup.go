package roles

import (
	"NestedScribbleBenchmark/noughtsandcrosses/callbacks"
	"NestedScribbleBenchmark/noughtsandcrosses/channels/minmaxstrategy"
	"NestedScribbleBenchmark/noughtsandcrosses/channels/minmaxstrategy_evalboard"
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/invitations"
	"NestedScribbleBenchmark/noughtsandcrosses/messages"
	"sync"
)

func MinMaxStrategy_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.MinMaxStrategy_RoleSetupChan, inviteChannels invitations.MinMaxStrategy_InviteSetupChan) {
	worker_invite_worker_2 := make(chan minmaxstrategy_evalboard.W_Chan, 1)
	worker_invite_worker_invitechan_2 := make(chan invitations.MinMaxStrategy_EvalBoard_W_InviteChan, 1)
	worker_master_int := make(chan int, 1)
	worker_master_label := make(chan messages.NoughtsAndCrosses_Label, 1)
	worker_invite_worker := make(chan minmaxstrategy.Master_Chan, 1)
	worker_invite_worker_invitechan := make(chan invitations.MinMaxStrategy_Master_InviteChan, 1)
	master_invite_master := make(chan minmaxstrategy.Master_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.MinMaxStrategy_Master_InviteChan, 1)
	master_worker_player := make(chan impl.Player, 1)
	master_worker_board := make(chan impl.Board, 1)
	master_worker_label := make(chan messages.NoughtsAndCrosses_Label, 1)

	worker_chan := minmaxstrategy.Worker_Chan{
		Player_From_Master: master_worker_player,
		Label_To_Master:    worker_master_label,
		Label_From_Master:  master_worker_label,
		Int_To_Master:      worker_master_int,
		Board_From_Master:  master_worker_board,
	}
	master_chan := minmaxstrategy.Master_Chan{
		Player_To_Worker:  master_worker_player,
		Label_To_Worker:   master_worker_label,
		Label_From_Worker: worker_master_label,
		Int_From_Worker:   worker_master_int,
		Board_To_Worker:   master_worker_board,
	}

	worker_inviteChan := invitations.MinMaxStrategy_Worker_InviteChan{
		Invite_Worker_To_MinMaxStrategy_Master_InviteChan:      worker_invite_worker_invitechan,
		Invite_Worker_To_MinMaxStrategy_Master:                 worker_invite_worker,
		Invite_Worker_To_MinMaxStrategy_EvalBoard_W_InviteChan: worker_invite_worker_invitechan_2,
		Invite_Worker_To_MinMaxStrategy_EvalBoard_W:            worker_invite_worker_2,
	}
	master_inviteChan := invitations.MinMaxStrategy_Master_InviteChan{
		Invite_Master_To_MinMaxStrategy_Master_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_MinMaxStrategy_Master:            master_invite_master,
	}

	roleChannels.Master_Chan <- master_chan

	inviteChannels.Master_InviteChan <- master_inviteChan

	wg.Add(1)

	worker_env := callbacks.New_MinMaxStrategy_Worker_State()
	go MinMaxStrategy_Worker(wg, worker_chan, worker_inviteChan, worker_env)
}
