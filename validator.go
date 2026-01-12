package main

import (
	"fmt"
)

func validateUpdate(metaPath, filePath string) error {
	m, err := readMetadata(metaPath)
	if err != nil {
		return fmt.Errorf("read metadata:%w", err)
	}
	actualHash, actualSize, err := sha256File(filePath)
	if err != nil {
		return fmt.Errorf("read firmware file:%w", err)
	}
	if m.Size != actualSize {
		return fmt.Errorf("size mismatch (meta=%d, actual=%d)", m.Size, actualSize)
	}
	if m.SHA256 != actualHash {
		return fmt.Errorf("hash mismatch (meta=%s, actual=%s)", m.SHA256, actualHash)
	}
	return nil
}
