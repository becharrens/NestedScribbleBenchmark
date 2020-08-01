package invitations

import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"

type SpectralNorm_TimesTransp_RoleSetupChan struct {
	M_Chan chan spectralnorm_timestransp.M_Chan
}

type SpectralNorm_TimesTransp_InviteSetupChan struct {
	M_InviteChan chan SpectralNorm_TimesTransp_M_InviteChan
}

type SpectralNorm_TimesTransp_M_InviteChan struct {
	Invite_M_To_SpectralNorm_TimesTransp_M            chan spectralnorm_timestransp.M_Chan
	Invite_M_To_SpectralNorm_TimesTransp_M_InviteChan chan SpectralNorm_TimesTransp_M_InviteChan
}

type SpectralNorm_TimesTransp_W_InviteChan struct {
}
