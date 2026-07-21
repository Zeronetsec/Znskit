// https://github.com/Zeronetsec/Znskit

package listins

import (
    "fmt"
)

func formatBytes(bytes int64) string {
    if bytes < 1024 {
        return fmt.Sprintf("%d B", bytes)
    }

    kb := float64(bytes) / 1024.0
    if kb < 1024 {
        return fmt.Sprintf("%.2f KB", kb)
    }

    mb := kb / 1024.0
    if mb < 1024 {
        return fmt.Sprintf("%.2f MB", mb)
    }

    gb := mb / 1024.0
    return fmt.Sprintf("%.2f GB", gb)
}

// Copyright (c) 2026 Zeronetsec