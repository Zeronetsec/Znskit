// https://github.com/Zeronetsec/Znskit

package utils

import (
    "fmt"
    "time"
    "znskit/utils/color"
)

func Birthday() {
    bdate := "05-20"
    now := time.Now().Format("01-02")
    tname := "znskit"

    if now == bdate {
        fmt.Printf(
            "%s› %sHappy birthday for %s%s %s🎉\n",
            color.R, color.N, color.GG, tname, color.N,
        )
        fmt.Println()
    }
}

// Copyright (c) 2026 Zeronetsec