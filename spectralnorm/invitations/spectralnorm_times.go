package invitations

import "ScribbleBenchmark/spectralnorm/channels/spectralnorm_times"

type SpectralNorm_Times_RoleSetupChan struct {
	M_Chan chan spectralnorm_times.M_Chan
}

type SpectralNorm_Times_InviteSetupChan struct {
	M_InviteChan chan SpectralNorm_Times_M_InviteChan
}

type SpectralNorm_Times_M_InviteChan struct {
	Invite_M_To_SpectralNorm_Times_M chan spectralnorm_times.M_Chan
	Invite_M_To_SpectralNorm_Times_M_InviteChan chan SpectralNorm_Times_M_InviteChan
}

type SpectralNorm_Times_W_InviteChan struct {

}