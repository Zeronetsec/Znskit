// https://github.com/Zeronetsec/Znskit

package info

import (
    "fmt"
)

func formatSize(kb int) string {
    if kb < 1024 {
        return fmt.Sprintf("%d KB", kb)
    }

    mb := float64(kb) / 1024.0
    if mb < 1024 {
        return fmt.Sprintf("%.2f MB", mb)
    }

    gb := mb / 1024.0
    return fmt.Sprintf("%.2f GB", gb)
}

// Copyright (c) 2026 Zeronetsec