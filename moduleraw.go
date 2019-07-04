package trackermod

import (
	"bytes"
	"fmt"
)

// Module in raw binary
type moduleRaw struct {
	Name            [20]byte      // offset 0
	Samples         [31]SampleRaw // offset 20
	Length          uint8         // offset 950 Value can be 1-128
	RestartObsolete uint8         // offset 951 In old .mods 127 which means repeat
	PatternOrders   [128]uint8    // offset 952 Order of patterns
	Magic           [4]byte       // offset 1080 'M.K.'
}

func (s moduleRaw) Validate() error {
	_, ok := Detect[string(s.Magic[:])]

	if !ok {
		return fmt.Errorf(`unknown: %v`, s.Magic)
	}

	if s.Length > 128 {
		return fmt.Errorf(`song length can't be > 128'`)
	}

	for _, p := range s.PatternOrders {
		if p > 64 {
			return fmt.Errorf(`pattern number can't be > 64`)
		}
	}

	return nil
}

func (s moduleRaw) getName() string {
	return string(bytes.TrimRight(s.Name[:], "\x00"))
}

// Detect module type
var Detect = map[string]Meta{
	"M.K.": {Channels: 4, Type: TrackerProtracker},
	"M!K!": {Channels: 4, Type: TrackerProtracker},
	"M&K!": {Channels: 4, Type: TrackerNoisetracker},
	"N.T.": {Channels: 4, Type: TrackerNoisetracker},
}
