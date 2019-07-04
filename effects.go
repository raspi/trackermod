package trackermod

type EffectCommand uint8

const (
	EffectNoneArpeggio   EffectCommand = iota // 0
	EffectPortamentoUp                        // 1
	EffectPortamentoDown                      // 2
	EffectTonePortamento                      // 3
	EffectVibrato                             // 4
	EffectTonePVolSlide                       // 5
	EffectVibraVolSlide                       // 6
	EffectTremolo                             // 7
	EffectNotUsed                             // 8
	EffectSampleOffset                        // 9
	EffectVolumeSlide                         // A 10
	EffectPositionJump                        // B 11
	EffectSetVolume                           // C 12
	EffectPatternBreak                        // D 13
	EffectMiscCmds                            // E 14
	EffectSetSpeed                            // F 15
)

func (e EffectCommand) String() string {
	switch e {
	case EffectNoneArpeggio:
		return "None/Arpeggio"
	case EffectPortamentoUp:
		return "PortamentoUp"
	case EffectPortamentoDown:
		return "PortamentoDown"
	case EffectTonePortamento:
		return "TonePortamento"
	case EffectVibrato:
		return "Vibrato"
	case EffectTonePVolSlide:
		return "TonePVolSlide"
	case EffectVibraVolSlide:
		return "VibraVolSlide"
	case EffectTremolo:
		return "Tremolo"
	case EffectNotUsed:
		return "NotUsed"
	case EffectSampleOffset:
		return "SampleOffset"
	case EffectVolumeSlide:
		return "VolumeSlide"
	case EffectPositionJump:
		return "PositionJump"
	case EffectSetVolume:
		return "SetVolume"
	case EffectPatternBreak:
		return "PatternBreak"
	case EffectMiscCmds:
		return "MiscCmds"
	case EffectSetSpeed:
		return "SetSpeed"

	default:
		return "ERR"
	}
}
