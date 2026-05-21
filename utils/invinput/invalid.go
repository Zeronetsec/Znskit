// https://github.com/Zeronetsec/Znskit

package invinput

import (
    "fmt"
    "znskit/utils/color"
)

func Invalid() {
    tname := "znskit"

    fmt.Printf(
        "%s[!] %sInvalid input!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %s%s --help%s\n",
        color.R, color.N, color.GG, tname, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec