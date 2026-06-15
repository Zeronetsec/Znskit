// https://github.com/Zeronetsec/Znskit

package module

import (
    "fmt"
    "github.com/Zeronetsec/Znskit/utils/color"
)

const (
    name = "Znskit"
    version = "v0.1"
    creator = "Zeronetsec"
    homepage = "https://github.com/Zeronetsec/Znskit"
)

func Version() {
    fmt.Printf(
        "%sName: %s%s%s\n",
        color.N, color.GG, name, color.N,
    )

    fmt.Printf(
        "%sVersion: %s%s%s\n",
        color.N, color.GG, version, color.N,
    )

    fmt.Printf(
        "%sCreator: %s%s%s\n",
        color.N, color.GG, creator, color.N,
    )

    fmt.Printf(
        "%sHomepage: %s%s%s\n",
        color.N, color.GG, homepage, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec