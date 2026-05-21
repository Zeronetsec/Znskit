// https://github.com/Zeronetsec/Znskit

package invinput

import (
    "fmt"
    "znskit/utils/color"
)

func Unknown(input string) {
    tname := "znskit"

    fmt.Printf(
        "%s[!] %sInvalid command: %s%s%s\n",
        color.R, color.N, color.GG, input, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %s%s --help%s\n",
        color.R, color.N, color.GG, tname, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec