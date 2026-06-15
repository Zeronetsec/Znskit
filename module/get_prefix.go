// https://github.com/Zeronetsec/Znskit

package module

import (
    "os"
    "path/filepath"
)

func GetPrefix() string {
    prefix := os.Getenv("PREFIX")

    if prefix == "" {
        prefix = "/usr"
    }

    return filepath.Clean(prefix)
}

// Copyright (c) 2026 Zeronetsec