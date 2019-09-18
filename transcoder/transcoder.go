package transcoder

const (
	PresetMedium = Preset("medium")
	PresetFast   = Preset("fast")
	PresetFaster = Preset("faster")
)

type Preset string

type Transcoder interface {
	Transcode(TranscodeTask) error
}

// TranscodeTask represents a transcoding task.
type TranscodeTask struct {
	Input   *TranscodeInput
	Outputs []*TranscodeOutput
}

// TranscodeInput ...
type TranscodeInput struct {
	URI string
}

// TranscodeOutput ...
type TranscodeOutput struct {
	URI              string
	Resolution       uint32
	VideoBitrateKbps uint32
	AudioBitrateKbps uint32
	FrameRate        uint32
	GOPSize          uint32
}
