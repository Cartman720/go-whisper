# Go-Whisper

Go-Whisper is a command-line tool for transcribing video files using OpenAI's Whisper API. It extracts audio from videos, transcribes it using OpenAI's speech-to-text model, and can optionally embed the transcription as subtitles back into the video.

## Author

**Aren Hovsepyan**

## Features

- Extract audio from video files
- Transcribe audio using OpenAI's Whisper model
- Generate SRT subtitle files
- Embed subtitles into the original video (optional)
- Support for multiple languages

## Requirements

- Go 1.24+
- FFmpeg (for audio extraction and subtitle embedding)
- OpenAI API key

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/cartman720/go-whisper.git
cd go-whisper
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Build the application

#### Quick Build

```bash
go build -o go-whisper
```

#### Cross-Compilation

You can compile the application for different platforms:

For Windows:
```bash
GOOS=windows GOARCH=amd64 go build -o go-whisper.exe
```

For macOS:
```bash
GOOS=darwin GOARCH=amd64 go build -o go-whisper
```

For Linux:
```bash
GOOS=linux GOARCH=amd64 go build -o go-whisper
```

### 4. Set up your environment variables

Create a `.env` file in the project root with your OpenAI API key:

```
OPENAI_API_KEY=your_openai_api_key
```

## Usage

### Using the Compiled Binary

```bash
./go-whisper --video /path/to/your/video.mp4
```

### Using go run (Without Compiling)

You can also run the application directly without compiling:

```bash
go run main.go --video /path/to/your/video.mp4
```

Or with additional parameters:

```bash
go run main.go --video /path/to/your/video.mp4 --output /path/to/output/dir --language fr
```

### Basic Usage

```bash
./go-whisper --video /path/to/your/video.mp4
```

This will:
1. Extract audio from the video
2. Transcribe the audio using OpenAI's Whisper model
3. Save the transcription as an SRT file in the current directory
4. Create a new video with embedded subtitles

### Options

```bash
./go-whisper --video /path/to/your/video.mp4 --output /path/to/output/dir --language en
```

- `--video`, `-v`: Path to the video file (required)
- `--output`, `-o`: Directory where output files will be saved (default: current directory)
- `--language`, `-l`: Language code for transcription (default: "en" for English)

### Translation Capability

Go-Whisper can also be used for translation. By setting the `--language` parameter to a different language than the original video's spoken language, the tool will transcribe and translate simultaneously. This is useful for creating subtitles in a language different from the audio.

For example, if you have a video in English but want Spanish subtitles:

```bash
./go-whisper --video ~/Videos/english_lecture.mp4 --language es
```

This will transcribe the English audio and produce Spanish subtitles embedded in the video.

## Output Files

The tool generates several files in the specified output directory:

1. `audio-[timestamp].wav`: Extracted audio from the video
2. `transcription-[timestamp].txt`: Transcription in SRT format
3. `transcribed-[timestamp].mp4`: Video with embedded subtitles

## Supported Languages

Go-Whisper supports all languages that OpenAI's Whisper model can transcribe. Some common language codes:

- `en`: English
- `es`: Spanish
- `fr`: French
- `de`: German
- `it`: Italian
- `ja`: Japanese
- `ko`: Korean
- `zh`: Chinese

## Example

```bash
# Transcribe a video in English
./go-whisper --video ~/Videos/lecture.mp4 --output ~/Documents/Transcriptions

# Transcribe a video in Spanish
./go-whisper --video ~/Videos/spanish_interview.mp4 --language es
```

## Development

To contribute to this project:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgements

- [OpenAI Whisper](https://openai.com/research/whisper) for the speech recognition model
- [Cobra](https://github.com/spf13/cobra) for the CLI framework
- [FFmpeg](https://ffmpeg.org/) for media processing 