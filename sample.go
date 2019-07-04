package modinfo

import "bytes"

// Sample in raw binary format
type SampleRaw struct {
	Name              [22]byte
	Length            uint16 // Multiply with 2 to get bytes
	Finetune          uint8  // Finetune 0x0 - 0xF
	Volume            uint8  // 0x0 - 0x40 (0 - 64)
	RepeatPointOffset uint16 // Multiply with 2 to get bytes
	RepeatLength      uint16 // Multiply with 2 to get bytes
}

func (s SampleRaw) getName() string {
	return string(bytes.TrimRight(s.Name[:], "\x00"))
}

// Sample
type Sample struct {
	Name              string `json:"name"`                          // Sample name
	WaveForm          []byte `json:"wave"`                          // Sample's binary waveform
	Finetune          uint8  `json:"finetune,omitempty"`            // Finetune
	Volume            uint8  `json:"volume,omitempty"`              // Sample's volume level
	RepeatPointOffset uint16 `json:"repeat_point_offset,omitempty"` // Point where sample repeats
	RepeatLength      uint16 `json:"repeat_length,omitempty"`       // How long the repeat is
}
