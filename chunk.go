package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ChunkAudioFile chunks the audio file into smaller parts
func ChunkAudioFile(filePath string, chunkSize int64) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)
	partNum := 1

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read file: %w", err)
		}
		if n == 0 {
			break
		}

		partFileName := fmt.Sprintf("%s_part_%d%s", filePath, partNum, filepath.Ext(filePath))
		partFile, err := os.Create(partFileName)
		if err != nil {
			return fmt.Errorf("failed to create part file: %w", err)
		}

		if _, err := partFile.Write(buffer[:n]); err != nil {
			partFile.Close()
			return fmt.Errorf("failed to write to part file: %w", err)
		}
		partFile.Close()
		partNum++
	}

	return nil
}