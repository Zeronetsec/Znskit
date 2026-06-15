// https://github.com/Zeronetsec/Znskit

package invinput

import (
    "fmt"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func MissingArgument() {
    fmt.Printf(
        "%s[!] %sMissing argument!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %sznskit --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec