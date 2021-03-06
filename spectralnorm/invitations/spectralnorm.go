package invitations

import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_times"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm"

type SpectralNorm_RoleSetupChan struct {
	Master_Chan chan spectralnorm.Master_Chan
	Worker_Chan chan spectralnorm.Worker_Chan
}

type SpectralNorm_InviteSetupChan struct {
	Master_InviteChan chan SpectralNorm_Master_InviteChan
	Worker_InviteChan chan SpectralNorm_Worker_InviteChan
}

type SpectralNorm_Master_InviteChan struct {
	Invite_Master_To_SpectralNorm_Master                   chan spectralnorm.Master_Chan
	Invite_Master_To_SpectralNorm_Master_InviteChan        chan SpectralNorm_Master_InviteChan
	Invite_Master_To_SpectralNorm_TimesTransp_M            chan spectralnorm_timestransp.M_Chan
	Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan chan SpectralNorm_TimesTransp_M_InviteChan
	Invite_Master_To_SpectralNorm_Times_M                  chan spectralnorm_times.M_Chan
	Invite_Master_To_SpectralNorm_Times_M_InviteChan       chan SpectralNorm_Times_M_InviteChan
	Invite_Worker_To_SpectralNorm_Worker                   chan spectralnorm.Worker_Chan
	Invite_Worker_To_SpectralNorm_Worker_InviteChan        chan SpectralNorm_Worker_InviteChan
}

type SpectralNorm_Worker_InviteChan struct {
	Master_Invite_To_SpectralNorm_Worker            chan spectralnorm.Worker_Chan
	Master_Invite_To_SpectralNorm_Worker_InviteChan chan SpectralNorm_Worker_InviteChan
}
