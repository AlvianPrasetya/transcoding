package transcoder

import (
	"fmt"
	"os"
	"os/exec"
)

// Libx264Transcoder uses libx264 software codec.
type Libx264Transcoder struct {
	Benchmark bool
	Preset    Preset
}

// Transcode executes the transcoder on the specified task.
func (transcoder *Libx264Transcoder) Transcode(task *TranscodeTask) error {
	if transcoder.Preset == "" {
		transcoder.Preset = PresetMedium
	}

	var args []string
	args = append(args, "-y")
	args = append(args, "-loglevel", "quiet", "-stats")
	if !transcoder.Benchmark {
		args = append(args, "-re") // Ingest at playback speed
	}
	args = append(args, "-i", task.Input.URI)

	for _, output := range task.Outputs {
		args = append(args, "-vf", fmt.Sprintf("scale=-2:%d", output.Resolution))
		args = append(args, "-c:v", "libx264")
		args = append(args, "-preset", string(transcoder.Preset))
		args = append(args, "-b:v", fmt.Sprintf("%dk", output.VideoBitrateKbps))
		args = append(args, "-maxrate", fmt.Sprintf("%dk", output.VideoBitrateKbps))
		args = append(args, "-bufsize", fmt.Sprintf("%dk", output.VideoBitrateKbps))
		args = append(args, "-r", fmt.Sprintf("%d", output.FrameRate))
		args = append(args, "-g", fmt.Sprintf("%d", output.GOPSize))
		args = append(args, "-c:a", "aac")
		args = append(args, "-b:a", fmt.Sprintf("%dk", output.AudioBitrateKbps))
		args = append(args, output.URI)
	}

	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
