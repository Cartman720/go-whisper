package lib

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var timePrefix string

// Extract audio from video file
func extractAudio(inputFile, outputPathDir string) (string, error) {
	audioFilePath := filepath.Join(outputPathDir, fmt.Sprintf("audio-%s.wav", timePrefix))

	// Create ffmpeg command with loglevel quiet to suppress output
	cmd := exec.Command(
		"ffmpeg",
		"-loglevel",
		"quiet",
		"-i",
		inputFile,
		"-vn",
		"-acodec",
		"pcm_s16le",
		"-ar",
		"44100",
		"-ac",
		"2",
		audioFilePath,
	)

	// Redirect stderr to null device
	cmd.Stderr = nil
	// Redirect stdout to null device
	cmd.Stdout = nil

	fmt.Println("Running ffmpeg to extract audio...")

	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return audioFilePath, nil
}

func getTimestampPrefix() string {
	now := time.Now()

	return now.Format("20060102_150405")
}

func initOpenAI() *openai.Client {
	openaiKey, ok := os.LookupEnv("OPENAI_API_KEY")

	if !ok {
		fmt.Println("OPENAI_API_KEY is not set")
		os.Exit(1)
	}

	client := openai.NewClient(option.WithAPIKey(openaiKey))
	return &client
}

func readAudioFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening audio file: ", err)
		return nil, err
	}

	return file, nil
}

func addSubtitlesToVideo(videoFilePath string, outputPath string, subtitlesFilePath string) {
	transcribedVideoFilePath := filepath.Join(outputPath, fmt.Sprintf("transcribed-%s.mp4", timePrefix))

	fmt.Println("Adding subtitles to video...")

	cmd := exec.Command(
		"ffmpeg",
		"-i", videoFilePath,
		"-vf", fmt.Sprintf("subtitles=%s", subtitlesFilePath),
		transcribedVideoFilePath,
	)

	cmd.Stderr = nil
	cmd.Stdout = nil

	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error adding subtitles to video: ", err)
		os.Exit(1)
	}

	fmt.Println("Subtitles added to video: ", videoFilePath)
}

func TranscribeAudio(inputPath string, outputPath string, outputLanguage string) {
	// Set time prefix for this run
	timePrefix = getTimestampPrefix()

	client := initOpenAI()

	fmt.Println("Extracting audio from video...")

	audioFilePath, err := extractAudio(inputPath, outputPath)

	if err != nil {
		fmt.Println("Error extracting audio: ", err)
		os.Exit(1)
	}

	fmt.Println("Running AI models to transcribe audio...")

	audioBytes, err := readAudioFile(audioFilePath)

	// Handle error if audio file cannot be read
	if err != nil {
		fmt.Println("Error reading audio file: ", err)
		os.Exit(1)
	}

	// For SRT response format, we need to use text response rather than JSON
	// When using non-JSON formats, we should let the SDK handle the response directly
	var transcription string

	_, err = client.Audio.Transcriptions.New(
		context.TODO(),
		openai.AudioTranscriptionNewParams{
			File:                   audioBytes,
			Model:                  openai.AudioModelWhisper1,
			Language:               openai.String(outputLanguage),
			Include:                []openai.TranscriptionInclude{openai.TranscriptionIncludeLogprobs},
			ResponseFormat:         openai.AudioResponseFormatSRT,
			TimestampGranularities: []string{"word"},
		},
		option.WithResponseBodyInto(&transcription),
	)

	// Handle error if transcription fails
	if err != nil {
		var apierr *openai.Error

		if errors.As(err, &apierr) {
			fmt.Println(string(apierr.DumpRequest(true)))
		}

		fmt.Println("Error transcribing audio: ", err)

		os.Exit(1)
	}

	// Save transcription to file
	transcriptionFilePath := filepath.Join(outputPath, fmt.Sprintf("transcription-%s.svr", timePrefix))
	err = os.WriteFile(transcriptionFilePath, []byte(transcription), 0644)

	if err != nil {
		fmt.Println("Error writing transcription to file: ", err)
		os.Exit(1)
	}

	fmt.Println("Transcription saved to: ", transcriptionFilePath)

	addSubtitlesToVideo(inputPath, outputPath, transcriptionFilePath)
}
