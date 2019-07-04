package trackermod

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang.org/x/xerrors"
	"io"
)

var Endianness = binary.BigEndian

type Module struct {
	Name           string      `json:"name"`
	Tracker        TrackerType `json:"tracker"`
	Channels       uint8       `json:"channels"`
	HighestPattern uint8       `json:"highest_pattern"`
	PatternOrders  []uint8     `json:"pattern_orders"`
	Patterns       []Pattern   `json:"patterns"`
	Samples        []Sample    `json:"samples"`
}

type Pattern struct {
	Index int        `json:"index"`
	Notes []NoteData `json:"notes"`
}

type Meta struct {
	Channels uint8
	Type     TrackerType
}

// Read magic string from offset 1080
func getMagic(r io.ReadSeeker) (m string, err error) {
	_, err = r.Seek(1080, io.SeekStart)
	if err != nil {
		return ``, err
	}

	var magic [4]byte
	err = binary.Read(r, Endianness, &magic)
	if err != nil {
		return ``, err
	}

	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		return ``, err
	}

	return string(magic[:]), nil
}

// Convert raw binary reader stream to a module
func NewFromReader(r io.ReadSeeker) (m Module, err error) {
	// Check magic
	magic, err := getMagic(r)

	meta, ok := Detect[magic]

	if !ok {
		return m, fmt.Errorf(`unknown magic: '%v'`, magic)
	}

	// Read raw rawBinaryModule data
	var rawBinaryModule moduleRaw

	err = binary.Read(r, Endianness, &rawBinaryModule)
	if err != nil {
		return m, err
	}

	err = rawBinaryModule.Validate()
	if err != nil {
		return m, err
	}

	m.Name = rawBinaryModule.getName()
	m.Tracker = meta.Type

	m.PatternOrders = rawBinaryModule.PatternOrders[:]

	// Calculate highest pattern
	highestPattern := uint8(0)

	for _, i := range rawBinaryModule.PatternOrders {
		if i > highestPattern {
			highestPattern = i
		}
	}

	m.HighestPattern = highestPattern
	m.Channels = meta.Channels

	// Read all patterns
	for i := uint8(0); i < m.HighestPattern; i++ {

		var notes []NoteData

		for patRow := uint8(0); patRow < 64; patRow++ {
			rawnotes := make([]NoteDataRaw, m.Channels)

			// Marshal
			err = binary.Read(r, Endianness, &rawnotes)
			if err != nil {
				offset, _ := r.Seek(0, io.SeekCurrent)
				return m, xerrors.Errorf(`couldn't read notes at pattern %v row %v at offset %v: %w`, i, patRow, offset, err)
			}

			for chidx, n := range rawnotes {
				note := n.ToNote()
				note.Channel = uint8(chidx)
				note.Row = patRow

				if note.Note == nil && note.Effect == EffectNoneArpeggio && note.EffectParameters == 0 {
					// Do not add "do nothing" note
					//continue
				}

				notes = append(notes, note)

			}
		}

		m.Patterns = append(m.Patterns, Pattern{Notes: notes, Index: int(i)})

	}

	highestSample := 0

	for idx, smp := range rawBinaryModule.Samples {
		if smp.Length > 0 {
			highestSample = idx
		}
	}

	// Read raw sample waveform data
	for i := 0; i < highestSample; i++ {
		smp := rawBinaryModule.Samples[i]
		smplen := smp.Length * 2

		var sd []byte

		sd = make([]byte, smplen)
		_, err = r.Read(sd)
		if err != nil {
			offset, _ := r.Seek(0, io.SeekCurrent)
			return m, xerrors.Errorf(`couldn't read sample #%[1]v (%[2]v B) at offset 0x%[3]x %04[3]d: %w`, i, smplen, offset, err)
		}

		m.Samples = append(m.Samples, Sample{
			Name:              smp.getName(),
			WaveForm:          sd,
			Finetune:          smp.Finetune,
			RepeatLength:      smp.RepeatLength * 2,
			RepeatPointOffset: smp.RepeatPointOffset * 2,
			Volume:            smp.Volume,
		})
	}

	return m, nil
}

// Custom JSON marshaler
func (m *Module) MarshalJSON() ([]byte, error) {

	var porders []int

	for _, p := range m.PatternOrders {
		porders = append(porders, int(p))
	}

	type Alias Module // Alias to prevent infinite loop
	return json.Marshal(&struct {
		Tracker       string `json:"tracker"`
		PatternOrders []int  `json:"pattern_orders,noescape"`
		*Alias
	}{
		Tracker:       m.Tracker.String(),
		PatternOrders: porders,
		Alias:         (*Alias)(m),
	})
}
