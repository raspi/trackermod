package trackermod

import "testing"

func TestEffectEnums(t *testing.T) {

	if EffectNoneArpeggio != 0 {
		t.Fatal(`0: enum mismatch`)
	}

	if EffectPortamentoUp != 1 {
		t.Fatal(`1: enum mismatch`)
	}

	if EffectPortamentoDown != 2 {
		t.Fatal(`2: enum mismatch`)
	}

	if EffectTonePortamento != 3 {
		t.Fatal(`3: enum mismatch`)
	}

	if EffectVibrato != 4 {
		t.Fatal(`4: enum mismatch`)
	}

	if EffectTonePVolSlide != 5 {
		t.Fatal(`5: enum mismatch`)
	}

	if EffectVibraVolSlide != 6 {
		t.Fatal(`6: enum mismatch`)
	}

	if EffectTremolo != 7 {
		t.Fatal(`7: enum mismatch`)
	}

	if EffectNotUsed != 8 {
		t.Fatal(`8: enum mismatch`)
	}

	if EffectSampleOffset != 9 {
		t.Fatal(`9: enum mismatch`)
	}

	if EffectVolumeSlide != 10 {
		t.Fatal(`A: enum mismatch`)
	}

	if EffectPositionJump != 11 {
		t.Fatal(`B: enum mismatch`)
	}

	if EffectSetVolume != 12 {
		t.Fatal(`C: enum mismatch`)
	}

	if EffectPatternBreak != 13 {
		t.Fatal(`D: enum mismatch`)
	}

	if EffectMiscCmds != 14 {
		t.Fatal(`E: enum mismatch`)
	}

	if EffectSetSpeed != 15 {
		t.Fatal(`F: enum mismatch`)
	}

}
