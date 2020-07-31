package roles

import "ScribbleBenchmark/knucleotide/messages/knucleotide"
import "ScribbleBenchmark/knucleotide/channels/schedulejobs"
import knucleotide_2 "ScribbleBenchmark/knucleotide/channels/knucleotide"
import "ScribbleBenchmark/knucleotide/invitations"
import "sync"

func KNucleotide_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.KNucleotide_RoleSetupChan, inviteChannels invitations.KNucleotide_InviteSetupChan)  {
	worker_master_sequenceresult := make(chan knucleotide.SequenceResult, 1)
	master_invite_master := make(chan schedulejobs.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.ScheduleJobs_M_InviteChan, 1)
	master_worker_sequencejob := make(chan knucleotide.SequenceJob, 1)

	worker_chan := knucleotide_2.Worker_Chan{
		Master_SequenceResult: worker_master_sequenceresult,
		Master_SequenceJob: master_worker_sequencejob,
	}
	master_chan := knucleotide_2.Master_Chan{
		Worker_SequenceResult: worker_master_sequenceresult,
		Worker_SequenceJob: master_worker_sequencejob,
	}

	worker_inviteChan := invitations.KNucleotide_Worker_InviteChan{

	}
	master_inviteChan := invitations.KNucleotide_Master_InviteChan{
		Invite_Master_To_ScheduleJobs_M_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_ScheduleJobs_M: master_invite_master,
	}

	roleChannels.Master_Chan <- master_chan
	roleChannels.Worker_Chan <- worker_chan

	inviteChannels.Master_InviteChan <- master_inviteChan
	inviteChannels.Worker_InviteChan <- worker_inviteChan
} 