# OTA Metadata Validator

A Linux-based CLI tool written in Go to validate OTA update metadata and firmware
integrity before deployment.

## Overview

This tool validates OTA update artifacts by checking:
- Metadata correctness (JSON format)
- Firmware file size
- SHA-256 checksum integrity

It is designed to model reliability-focused validation workflows commonly used
in Linux-based OTA update systems.

## Features

- JSON-based OTA metadata parsing
- Firmware integrity verification using SHA-256
- Explicit failure handling with clear error messages
- Concurrent validation of multiple firmware artifacts

## Project Structure 
cmd/validator        CLI entrypoint
internal/validator   Core validation logic
testdata             Sample OTA metadata and firmware files

## Usage

```bash
go run ./cmd/validator --meta testdata/metadata.json --file testdata/firmware.bin
```

### Example Output
```bash
PASS [testdata/firmware.bin]
FAIL [testdata/firmware_bad.bin]: sha256 mismatch
```

## Environment
- Linux
- Go 1.22