// https://github.com/Zeronetsec/Znskit

package listins

import (
    "os"
)

func getFileSize(filePath string) (string, bool) {
    info, err := os.Stat(filePath)
    if err != nil {
        return "", false
    }
    return formatBytes(info.Size()), true
}

// Copyright (c) 2026 Zeronetsec