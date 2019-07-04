package trackermod

type TrackerType uint8

const (
	TrackerUnknown TrackerType = iota
	TrackerProtracker
	TrackerNoisetracker
	TrackerSoundtracker
	TrackerProtrackerClone
)

func (tt TrackerType) String() string {
	switch tt {
	case TrackerUnknown:
		return "Unknown"
	case TrackerProtracker:
		return "Protracker"
	case TrackerNoisetracker:
		return "Noisetracker"
	case TrackerSoundtracker:
		return "Soundtracker"
	case TrackerProtrackerClone:
		return "ProtrackerClone"
	default:
		return "ERR"
	}
}
