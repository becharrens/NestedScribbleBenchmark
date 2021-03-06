package invitations

import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuch"

type Fannkuch_RoleSetupChan struct {
	Main_Chan   chan fannkuch.Main_Chan
	Worker_Chan chan fannkuch.Worker_Chan
}

type Fannkuch_InviteSetupChan struct {
	Main_InviteChan   chan Fannkuch_Main_InviteChan
	Worker_InviteChan chan Fannkuch_Worker_InviteChan
}

type Fannkuch_Main_InviteChan struct {
	Worker_Invite_To_FannkuchRecursive_Source            chan fannkuchrecursive.Source_Chan
	Worker_Invite_To_FannkuchRecursive_Source_InviteChan chan FannkuchRecursive_Source_InviteChan
}

type Fannkuch_Worker_InviteChan struct {
	Invite_Main_To_FannkuchRecursive_Source              chan fannkuchrecursive.Source_Chan
	Invite_Main_To_FannkuchRecursive_Source_InviteChan   chan FannkuchRecursive_Source_InviteChan
	Invite_Worker_To_FannkuchRecursive_Worker            chan fannkuchrecursive.Worker_Chan
	Invite_Worker_To_FannkuchRecursive_Worker_InviteChan chan FannkuchRecursive_Worker_InviteChan
}
