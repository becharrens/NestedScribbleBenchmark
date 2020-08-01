package invitations

import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"

type FannkuchRecursive_RoleSetupChan struct {
	Source_Chan chan fannkuchrecursive.Source_Chan
	Worker_Chan chan fannkuchrecursive.Worker_Chan
}

type FannkuchRecursive_InviteSetupChan struct {
	Source_InviteChan chan FannkuchRecursive_Source_InviteChan
	Worker_InviteChan chan FannkuchRecursive_Worker_InviteChan
}

type FannkuchRecursive_Source_InviteChan struct {
	NewWorker_Invite_To_FannkuchRecursive_Source            chan fannkuchrecursive.Source_Chan
	NewWorker_Invite_To_FannkuchRecursive_Source_InviteChan chan FannkuchRecursive_Source_InviteChan
}

type FannkuchRecursive_Worker_InviteChan struct {
}

type FannkuchRecursive_NewWorker_InviteChan struct {
	Invite_NewWorker_To_FannkuchRecursive_Worker            chan fannkuchrecursive.Worker_Chan
	Invite_NewWorker_To_FannkuchRecursive_Worker_InviteChan chan FannkuchRecursive_Worker_InviteChan
	Invite_Source_To_FannkuchRecursive_Source               chan fannkuchrecursive.Source_Chan
	Invite_Source_To_FannkuchRecursive_Source_InviteChan    chan FannkuchRecursive_Source_InviteChan
}
