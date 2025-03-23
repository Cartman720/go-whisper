package cmd

import (
	"fmt"
	"os"

	"github.com/cartman720/go-whisper/lib"
	"github.com/spf13/cobra"
)

type options struct {
	videoPath      string
	outputLocation string
	outputLanguage string
}

var opts = options{}

var rootCmd = &cobra.Command{
	Use:   "go-whisper",
	Short: "go-whisper is a tool for transcribing video files",
	Long:  `go-whisper is a tool for transcribing video files. It uses the OpenAI API to transcribe the video files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Video path: %s\n", opts.videoPath)
		fmt.Printf("Output location: %s\n", opts.outputLocation)
		fmt.Printf("Output language: %s\n", opts.outputLanguage)

		lib.TranscribeAudio(opts.videoPath, opts.outputLocation, opts.outputLanguage)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&opts.videoPath, "video", "v", "", "path to the video file (required)")
	rootCmd.Flags().StringVarP(&opts.outputLocation, "output", "o", "./", "path to the output file")
	rootCmd.Flags().StringVarP(&opts.outputLanguage, "language", "l", "en", "language of the output file")

	// Video file path is required
	rootCmd.MarkFlagRequired("video")

	cobra.OnInitialize(func() {
		fmt.Println("Initializing...")
	})
}

func Execute() {
	_, err := rootCmd.ExecuteC()

	// if error, print error and exit
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
