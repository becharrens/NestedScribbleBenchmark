package invitations

import "NestedScribbleBenchmark/knucleotide/channels/schedulejobs"

type ScheduleJobs_RoleSetupChan struct {
	M_Chan chan schedulejobs.M_Chan
}

type ScheduleJobs_InviteSetupChan struct {
	M_InviteChan chan ScheduleJobs_M_InviteChan
}

type ScheduleJobs_M_InviteChan struct {
	Invite_M_To_ScheduleJobs_M            chan schedulejobs.M_Chan
	Invite_M_To_ScheduleJobs_M_InviteChan chan ScheduleJobs_M_InviteChan
}

type ScheduleJobs_W_InviteChan struct {
}
