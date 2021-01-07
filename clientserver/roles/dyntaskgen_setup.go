package roles

import "NestedScribbleBenchmark/clientserver/messages"
import "NestedScribbleBenchmark/clientserver/channels/dyntaskgen"
import "NestedScribbleBenchmark/clientserver/invitations"
import "NestedScribbleBenchmark/clientserver/callbacks"
import "sync"

func DynTaskGen_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.DynTaskGen_RoleSetupChan, inviteChannels invitations.DynTaskGen_InviteSetupChan)  {
	w_s_string := make(chan string, 1)
	w_s_label := make(chan messages.ClientServer_Label, 1)
	s_invite_s := make(chan dyntaskgen.S_Chan, 1)
	s_invite_s_invitechan := make(chan invitations.DynTaskGen_S_InviteChan, 1)
	s_w_string := make(chan string, 1)
	s_w_label := make(chan messages.ClientServer_Label, 1)

	w_chan := dyntaskgen.W_Chan{
		String_To_S: w_s_string,
		String_From_S: s_w_string,
		Label_To_S: w_s_label,
		Label_From_S: s_w_label,
	}
	s_chan := dyntaskgen.S_Chan{
		String_To_W: s_w_string,
		String_From_W: w_s_string,
		Label_To_W: s_w_label,
		Label_From_W: w_s_label,
	}

	w_inviteChan := invitations.DynTaskGen_W_InviteChan{

	}
	s_inviteChan := invitations.DynTaskGen_S_InviteChan{
		Invite_S_To_DynTaskGen_S_InviteChan: s_invite_s_invitechan,
		Invite_S_To_DynTaskGen_S: s_invite_s,
	}

	roleChannels.S_Chan <- s_chan

	inviteChannels.S_InviteChan <- s_inviteChan

	wg.Add(1)

	w_env := callbacks.New_DynTaskGen_W_State()
	go DynTaskGen_W(wg, w_chan, w_inviteChan, w_env)
} 