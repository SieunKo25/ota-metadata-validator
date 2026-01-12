package main

import (
	"flag"
	"fmt"
	"ota-metadata-validator/internal/validator"
)

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
			err := validator.ValidateUpdate(*metaPath, file)
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
