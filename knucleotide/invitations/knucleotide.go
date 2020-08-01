package invitations

import "NestedScribbleBenchmark/knucleotide/channels/schedulejobs"
import "NestedScribbleBenchmark/knucleotide/channels/knucleotide"

type KNucleotide_RoleSetupChan struct {
	Master_Chan chan knucleotide.Master_Chan
	Worker_Chan chan knucleotide.Worker_Chan
}

type KNucleotide_InviteSetupChan struct {
	Master_InviteChan chan KNucleotide_Master_InviteChan
	Worker_InviteChan chan KNucleotide_Worker_InviteChan
}

type KNucleotide_Master_InviteChan struct {
	Invite_Master_To_ScheduleJobs_M            chan schedulejobs.M_Chan
	Invite_Master_To_ScheduleJobs_M_InviteChan chan ScheduleJobs_M_InviteChan
}

type KNucleotide_Worker_InviteChan struct {
}
