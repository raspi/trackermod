package modinfo

import (
	"encoding/json"
	"fmt"
)

// Note data
type NoteData struct {
	SampleNumber     uint8         `json:"sample"`
	Note             *NoteInfo     `json:"note"`
	Effect           EffectCommand `json:"effect"`
	EffectParameters uint8         `json:"effect_parameters"`
	Channel          uint8         `json:"channel"`
	Row              uint8         `json:"row"`
}

func (n NoteData) String() string {
	var note string

	if n.Note == nil {
		note = `   `
	} else {
		note = n.Note.String()
	}

	return fmt.Sprintf(`%d %02X %v %02d %01X%02X`, n.Channel, n.Row, note, n.SampleNumber, uint8(n.Effect), n.EffectParameters)
}

func (n NoteData) Validate() error {
	if n.SampleNumber > 0x1F {
		return fmt.Errorf(`invalid sample: %v`, n.SampleNumber)
	}

	return nil
}

// Convert NoteData back to raw binary
func (n NoteData) ToRaw() NoteDataRaw {
	period := uint32(0)

	if n.Note != nil {
		period = uint32(n.Note.ToPeriod())
	}

	smpupperbits := uint32(NoteSampleMaskUpper & n.SampleNumber >> 4)
	smplowerbits := uint32(NoteSampleMaskLower & n.SampleNumber)
	eff := uint32(n.Effect)
	effparam := uint32(n.EffectParameters)

	return NoteDataRaw(smpupperbits<<28 | period<<16 | smplowerbits<<12 | eff<<8 | effparam)
}

// Custom JSON marshaler
func (n *NoteData) MarshalJSON() ([]byte, error) {

	notestr := ""

	if n.Note != nil {
		notestr = n.Note.String()
	}

	type Alias NoteData // Alias to prevent infinite loop
	return json.Marshal(&struct {
		Note   string `json:"note,omitempty"`
		Effect string `json:"effect"`
		*Alias
	}{
		Note:   notestr,
		Effect: n.Effect.String(),
		Alias:  (*Alias)(n),
	})
}
