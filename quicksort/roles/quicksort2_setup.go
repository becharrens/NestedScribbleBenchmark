package roles

import "NestedScribbleBenchmark/quicksort/messages"
import "NestedScribbleBenchmark/quicksort/channels/quicksort2"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import "sync"

func QuickSort2_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.QuickSort2_RoleSetupChan, inviteChannels invitations.QuickSort2_InviteSetupChan) {
	r_p_intarr := make(chan []int, 1)
	r_p_label := make(chan messages.QuickSort_Label, 1)
	l_p_intarr := make(chan []int, 1)
	l_p_label := make(chan messages.QuickSort_Label, 1)
	r_invite_r := make(chan quicksort2.P_Chan, 1)
	r_invite_r_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	l_invite_l := make(chan quicksort2.P_Chan, 1)
	l_invite_l_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	p_r_intarr := make(chan []int, 1)
	p_r_label := make(chan messages.QuickSort_Label, 1)
	p_l_intarr := make(chan []int, 1)
	p_l_label := make(chan messages.QuickSort_Label, 1)

	r_chan := quicksort2.R_Chan{
		Label_To_P:    r_p_label,
		Label_From_P:  p_r_label,
		IntArr_To_P:   r_p_intarr,
		IntArr_From_P: p_r_intarr,
	}
	p_chan := quicksort2.P_Chan{
		Label_To_R:    p_r_label,
		Label_To_L:    p_l_label,
		Label_From_R:  r_p_label,
		Label_From_L:  l_p_label,
		IntArr_To_R:   p_r_intarr,
		IntArr_To_L:   p_l_intarr,
		IntArr_From_R: r_p_intarr,
		IntArr_From_L: l_p_intarr,
	}
	l_chan := quicksort2.L_Chan{
		Label_To_P:    l_p_label,
		Label_From_P:  p_l_label,
		IntArr_To_P:   l_p_intarr,
		IntArr_From_P: p_l_intarr,
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
