// https://github.com/Zeronetsec/Znskit

package invinput

import (
    "fmt"
    "znskit/utils/color"
)

func MissingArgs() {
    tname := "znskit"

    fmt.Printf(
        "%s[!] %sMissing arguments!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %s%s --help%s\n",
        color.R, color.N, color.GG, tname, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec