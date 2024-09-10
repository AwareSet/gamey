package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	// "runtime/trace"

	"github.com/atotto/clipboard"
	"github.com/openai/openai-go"
)

// func main() {
// 	whisperTranscription("/Users/knighty/Documents/vid/wattsAudio.mp4")
// }



func WhisperTranscription(filePath string) string {
	client := openai.NewClient()
	ctx := context.Background()

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	transcription, err := client.Audio.Transcriptions.New(ctx, openai.AudioTranscriptionNewParams{
		Model: openai.F(openai.AudioModelWhisper1),
		File:  openai.F[io.Reader](file),
		
	})
	if err != nil {
		panic(err)
	}
	println(transcription.Text)
	clipboard.WriteAll(transcription.Text)
	saveToFile(transcription.Text)
	return transcription.Text

}

func saveToFile(text string) {
	file, err := os.Create("transcription_" + time.Now().Format("2006-01-02_15-04-05") + ".md") // to add date and time 
	// "transcription_<date and time>.md" to use function time.Now() 
	// use time.Format("2006-01-02_15-04-05") to get the date and time
	// the full will be os.Create("transcription_" + time.Now().Format("2006-01-02_15-04-05") + ".md")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println("Transcription saved to file")
	// ocount number of words
	words := strings.Split(text, " ")
	println("Number of words: ", len(words))
	
}

// WhisperTranscription transcribes the audio file at the given filepath.
func WhisperTranscription2(filepath string) string {
	// Example implementation
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return fmt.Sprintf("File %s does not exist", filepath)
	}
	return fmt.Sprintf("Transcription for %s", filepath)
}