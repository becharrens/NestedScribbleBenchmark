package roles

import "NestedScribbleBenchmark/knucleotide/channels/schedulejobs"
import "NestedScribbleBenchmark/knucleotide/channels/knucleotide"
import "NestedScribbleBenchmark/knucleotide/invitations"
import "sync"

func KNucleotide_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.KNucleotide_RoleSetupChan, inviteChannels invitations.KNucleotide_InviteSetupChan) {
	master_invite_master := make(chan schedulejobs.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.ScheduleJobs_M_InviteChan, 1)

	master_chan := knucleotide.Master_Chan{}

	master_inviteChan := invitations.KNucleotide_Master_InviteChan{
		Invite_Master_To_ScheduleJobs_M_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_ScheduleJobs_M:            master_invite_master,
	}

	roleChannels.Master_Chan <- master_chan

	inviteChannels.Master_InviteChan <- master_inviteChan
}
