// https://github.com/Zeronetsec/Znskit

package banner

import (
    "embed"
    "fmt"
    "github.com/Zeronetsec/Znskit/utils/color"
)

//go:embed ascii/*.txt
var asciiFS embed.FS

func Show() {
    data, err := asciiFS.ReadFile("ascii/znskit.txt")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError loading banner!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s%s%s\n",
        color.B, string(data), color.N,
    )
}

// Copyright (c) 2026 Zeronetsec