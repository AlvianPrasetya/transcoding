package main

import (
	"log"
	"time"

	"github.com/alvianprasetya/transcoding/transcoder"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	// Create transcode task
	transcodeTask := &transcoder.TranscodeTask{
		Input: &transcoder.TranscodeInput{
			URI: "in.mp4",
		},
		Outputs: []*transcoder.TranscodeOutput{
			&transcoder.TranscodeOutput{
				URI:              "out_720.mp4",
				Resolution:       720,
				VideoBitrateKbps: 3000,
				AudioBitrateKbps: 128,
				FrameRate:        30,
				GOPSize:          60,
			},
			&transcoder.TranscodeOutput{
				URI:              "out_480.mp4",
				Resolution:       480,
				VideoBitrateKbps: 1300,
				AudioBitrateKbps: 64,
				FrameRate:        30,
				GOPSize:          60,
			},
			&transcoder.TranscodeOutput{
				URI:              "out_360.mp4",
				Resolution:       360,
				VideoBitrateKbps: 900,
				AudioBitrateKbps: 64,
				FrameRate:        30,
				GOPSize:          60,
			},
			&transcoder.TranscodeOutput{
				URI:              "out_240.mp4",
				Resolution:       240,
				VideoBitrateKbps: 600,
				AudioBitrateKbps: 32,
				FrameRate:        30,
				GOPSize:          60,
			},
		},
	}

	stopChan := make(chan struct{})

	var totalCPU float64
	var dataPoints int
	go func() {
		for {
			select {
			case <-stopChan:
				return
			default:
				cpuUsages, err := cpu.Percent(0, false)
				if err != nil {
					log.Fatal(err)
				}

				totalCPU += cpuUsages[0]
				dataPoints++

				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	libx264 := transcoder.NVENCTranscoder{
		Benchmark: true,
		Preset:    transcoder.PresetMedium,
	}
	if err := libx264.Transcode(transcodeTask); err != nil {
		log.Print(err)
	}

	close(stopChan)

	log.Printf("Average CPU usage: %.2f", totalCPU/float64(dataPoints))
}
