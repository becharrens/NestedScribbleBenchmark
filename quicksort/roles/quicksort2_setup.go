package roles

import "NestedScribbleBenchmark/quicksort/messages/quicksort2"
import quicksort2_2 "NestedScribbleBenchmark/quicksort/channels/quicksort2"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import "sync"

func QuickSort2_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.QuickSort2_RoleSetupChan, inviteChannels invitations.QuickSort2_InviteSetupChan) {
	p_r_done := make(chan quicksort2.Done, 1)
	p_l_done := make(chan quicksort2.Done, 1)
	r_p_sortedright := make(chan quicksort2.SortedRight, 1)
	l_p_sortedleft := make(chan quicksort2.SortedLeft, 1)
	r_invite_r := make(chan quicksort2_2.P_Chan, 1)
	r_invite_r_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	l_invite_l := make(chan quicksort2_2.P_Chan, 1)
	l_invite_l_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	p_r_rightpartition := make(chan quicksort2.RightPartition, 1)
	p_l_leftparitition := make(chan quicksort2.LeftParitition, 1)

	r_chan := quicksort2_2.R_Chan{
		P_SortedRight:    r_p_sortedright,
		P_RightPartition: p_r_rightpartition,
		P_Done:           p_r_done,
	}
	p_chan := quicksort2_2.P_Chan{
		R_SortedRight:    r_p_sortedright,
		R_RightPartition: p_r_rightpartition,
		R_Done:           p_r_done,
		L_SortedLeft:     l_p_sortedleft,
		L_LeftParitition: p_l_leftparitition,
		L_Done:           p_l_done,
	}
	l_chan := quicksort2_2.L_Chan{
		P_SortedLeft:     l_p_sortedleft,
		P_LeftParitition: p_l_leftparitition,
		P_Done:           p_l_done,
	}

	r_inviteChan := invitations.QuickSort2_R_InviteChan{
		Invite_R_To_QuickSort2_P_InviteChan: r_invite_r_invitechan,
		Invite_R_To_QuickSort2_P:            r_invite_r,
	}
	p_inviteChan := invitations.QuickSort2_P_InviteChan{}
	l_inviteChan := invitations.QuickSort2_L_InviteChan{
		Invite_L_To_QuickSort2_P_InviteChan: l_invite_l_invitechan,
		Invite_L_To_QuickSort2_P:            l_invite_l,
	}

	roleChannels.P_Chan <- p_chan

	inviteChannels.P_InviteChan <- p_inviteChan

	wg.Add(2)

	l_env := callbacks.New_QuickSort2_L_State()
	go QuickSort2_L(wg, l_chan, l_inviteChan, l_env)

	r_env := callbacks.New_QuickSort2_R_State()
	go QuickSort2_R(wg, r_chan, r_inviteChan, r_env)
}
