package input

import (
	"bufio"
	"os"
)

type FileInput struct {
	inputPath string
}

func NewFileInput(path string) FileInput {
	return FileInput{inputPath: path}
}

func (f FileInput) Read() ([]string, error) {
	file, err := os.Open(f.inputPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
