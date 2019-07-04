package modinfo

import "fmt"

// NoteData data in raw binary format
type NoteDataRaw uint32

const (
	// _____byte 1_____   byte2_    _____byte 3_____   byte4_
	// /                 /        /                 /
	// 0000          0000-00000000  0000          0000-00000000
	//
	// Upper four    12 bits for    Lower four    Effect command.
	// bits of sam-  note period.   bits of sam-
	// ple number.                  ple number.

	// Bit masks for reading instrument number, note and effect
	NoteMaskSampleBitsUpper  = uint32(0xF0000000) // 04 11110000 00000000  00000000 00000000 Upper bits for sample number
	NoteMaskPeriod           = uint32(0x0FFF0000) // 12 00001111 11111111  00000000 00000000 Period (note)
	NoteMaskSampleBitsLower  = uint32(0x0000F000) // 04 00000000 00000000  11110000 00000000 Lower bits for sample number
	NoteMaskEffect           = uint32(0x00000F00) // 04 00000000 00000000  00001111 00000000 Effect number 0x0 - 0xF
	NoteMaskEffectParameters = uint32(0x000000FF) // 08 00000000 00000000  00000000 11111111 Parameter(s) for effect

	// split uint8 to uint4+uint4
	NoteSampleMaskUpper = uint8(0xF0) // 4 11110000
	NoteSampleMaskLower = uint8(0x0F) // 4 00001111
)

// Convert raw binary to proper note data
func (n NoteDataRaw) ToNote() NoteData {
	smpupperbits := NoteMaskSampleBitsUpper & uint32(n) >> 28
	period := NotePeriod(NoteMaskPeriod & uint32(n) >> 16)
	smplowerbits := NoteMaskSampleBitsLower & uint32(n) >> 12
	eff := NoteMaskEffect & uint32(n) >> 8
	effparam := NoteMaskEffectParameters & uint32(n)

	// Combine upper and lower bits
	smpnum := uint8(smpupperbits<<4 | smplowerbits)

	note, ok := NotePeriodMap[period]

	var notevar *NoteInfo

	if ok {
		notevar = &note
	} else {

		if period > 0 {
			panic(fmt.Errorf(`not found: %v`, period))
		}

		notevar = nil
	}

	return NoteData{
		SampleNumber:     smpnum,
		Note:             notevar,
		Effect:           EffectCommand(eff),
		EffectParameters: uint8(effparam),
	}
}

type NoteName uint8

const (
	NoteC      NoteName = iota // 0
	NoteCSharp                 // 1
	NoteD                      // 2
	NoteDSharp                 // 3
	NoteE                      // 4
	NoteF                      // 5
	NoteFSharp                 // 6
	NoteG                      // 7
	NoteGSharp                 // 8
	NoteA                      // 9
	NoteASharp                 // 10
	NoteB                      // 11
)

func (n NoteName) String() string {
	switch n {
	case NoteC:
		return "C-"
	case NoteCSharp:
		return "C#"
	case NoteD:
		return "D-"
	case NoteDSharp:
		return "D#"
	case NoteE:
		return "E-"
	case NoteF:
		return "F-"
	case NoteFSharp:
		return "F#"
	case NoteG:
		return "G-"
	case NoteGSharp:
		return "G#"
	case NoteA:
		return "A-"
	case NoteASharp:
		return "A#"
	case NoteB:
		return "B-"
	default:
		return "ERR"
	}
}

type NoteInfo struct {
	Note   NoteName
	Octave uint8
}

func (ni NoteInfo) String() string {
	return fmt.Sprintf(`%v%d`, ni.Note, ni.Octave)
}

// Find NotePeriod from a map
func (ni *NoteInfo) ToPeriod() NotePeriod {

	if ni == nil {
		return NotePeriod(0)
	}

	for period, ninfo := range NotePeriodMap {
		if ninfo.Note == ni.Note && ninfo.Octave == ni.Octave {
			return period
		}
	}

	return NotePeriod(0)
}

type NotePeriod uint16

// Map of notes and octaves for lookup
var NotePeriodMap = map[NotePeriod]NoteInfo{
	856: {NoteC, 1},
	808: {NoteCSharp, 1},
	762: {NoteD, 1},
	720: {NoteDSharp, 1},
	678: {NoteE, 1},
	640: {NoteF, 1},
	604: {NoteFSharp, 1},
	570: {NoteG, 1},
	538: {NoteGSharp, 1},
	508: {NoteA, 1},
	480: {NoteASharp, 1},
	453: {NoteB, 1},

	428: {NoteC, 2},
	404: {NoteCSharp, 2},
	381: {NoteD, 2},
	360: {NoteDSharp, 2},
	339: {NoteE, 2},
	320: {NoteF, 2},
	302: {NoteFSharp, 2},
	285: {NoteG, 2},
	269: {NoteGSharp, 2},
	254: {NoteA, 2},
	240: {NoteASharp, 2},
	226: {NoteB, 2},

	214: {NoteC, 3},
	202: {NoteCSharp, 3},
	190: {NoteD, 3},
	180: {NoteDSharp, 3},
	170: {NoteE, 3},
	160: {NoteF, 3},
	151: {NoteFSharp, 3},
	143: {NoteG, 3},
	135: {NoteGSharp, 3},
	127: {NoteA, 3},
	120: {NoteASharp, 3},
	113: {NoteB, 3},
}
