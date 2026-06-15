// https://github.com/Zeronetsec/Znskit

package module

import (
    "os"
)

func exists(path string) bool {
    _, err := os.Lstat(path)
    return err == nil
}

// Copyright (c) 2026 Zeronetsec