package invitations

import "NestedScribbleBenchmark/generaldns/channels/cached"

type Cached_RoleSetupChan struct {
	Res_Chan chan cached.Res_Chan
}

type Cached_InviteSetupChan struct {
	Res_InviteChan chan Cached_res_InviteChan
}

type Cached_res_InviteChan struct {

}