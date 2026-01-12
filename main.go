package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Metadata struct {
	Version string `json:"version"`
	Size    int64  `json:"size"`
	SHA256  string `json:"sha256"`
}

func sha256File(path string) (string, int64, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", 0, err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), int64(len(data)), nil
}

func readMetadata(path string) (*Metadata, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m Metadata
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func main() {
	metaPath := flag.String("meta", "testdata/metadata.json", "Path to update metadata JSON")
	flag.Parse()

	files := []string{
		"testdata/firmware.bin",
		"testdata/firmware_bad.bin",
	}

	results := make(chan string)

	for _, f := range files {
		go func(file string) {
			err := validateUpdate(*metaPath, file)
			if err != nil {
				results <- fmt.Sprintf("FAIL [%s]: %v", file, err)
				return
			}
			results <- fmt.Sprintf("PASS [%s]", file)
		}(f)
	}

	for i := 0; i < len(files); i++ {
		fmt.Println(<-results)
	}
}
