package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/3d0c/gmf"
	"github.com/go-resty/resty/v2"
)

func startFFmpeg() (*exec.Cmd, error) {
	ffmpegCmd := exec.Command("ffmpeg",
		"-f", "avfoundation",         // input format (macOS)
		"-i", "0",                    // input device index (webcam)
		"-f", "avfoundation",         // input format (macOS)
		"-i", ":1",                   // input device index (microphone)
		"-vf", "format=yuv420p",      // video format
		"-c:v", "rawvideo",
		"-pix_fmt", "yuv420p",
		"-an",                        // disable audio
		"-f", "avfoundation",         // output format (macOS)
		"-pix_fmt", "uyvy422",
		"-c:v", "rawvideo",
		"-an",                        // disable audio
		"-vf", "format=uyvy422",      // video format
		"-vf", "fps=30",
		"-f", "avfoundation",         // output format (macOS)
		"-framerate", "30",
		"-pix_fmt", "yuv420p",
		"-c:v", "rawvideo",
		"-an",                        // disable audio
		"-f", "v4l2",                 // output format (Linux)
		"/dev/video2",                // virtual camera device (Linux)
	)

	err := ffmpegCmd.Start()
	if err != nil {
		return nil, err
	}

	time.Sleep(2 * time.Second) // wait for ffmpeg to start

	return ffmpegCmd, nil
}

func connectToJitsiMeet() error {
	client := resty.New()

	response, err := client.R().Get("https://meet.jit.si/http-bind")
	if err != nil {
		return err
	}

	// Extract the session ID from the response
	sessionID := response.Header().Get("JSESSIONID")

	// Generate a random room name or use your desired room name
	roomName := "C0MFYC0W0RK1NGCLUB"

	// Join the Jitsi Meet room URL
	url := fmt.Sprintf("https://meet.jit.si/%s/%s", sessionID, roomName)

	// Open the room URL in a browser or handle it programmatically
	fmt.Println("Join the Jitsi Meet room: ", url)

	return nil
}

func main() {
	// Start ffmpeg process
	ffmpegCmd, err := startFFmpeg()
	if err != nil {
		fmt.Println("Failed to start ffmpeg:", err)
		os.Exit(1)
	}
	defer ffmpegCmd.Process.Kill()

	// Connect to the Jitsi Meet room
	err = connectToJitsiMeet()
	if err != nil {
		fmt.Println("Failed to connect to Jitsi Meet:", err)
		os.Exit(1)
	}

	// Keep the program running until interrupted
	select {}
}
