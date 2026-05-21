// https://github.com/Zeronetsec/Znskit

package core

import (
    "fmt"
    "znskit/utils/color"
)

func Reinstall(pkg string) {
    if pkg == "" {
        fmt.Printf(
            "%s[!] %sMissing packahe name!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sReinstalling: %s%s%s\n",
        color.B, color.N, color.GG, pkg, color.N,
    )

    fmt.Printf(
        "%s- %sExec: %sznskit --uninstall %s%s%s\n",
        color.DG, color.N, color.GG, color.CC, pkg, color.N,
    )

    Uninstall(pkg)

    fmt.Printf(
        "%s- %sExec: %sznskit --install %s%s%s\n",
        color.DG, color.N, color.GG, color.CC, pkg, color.N,
    )

    Install(pkg)
}

// Copyright (c) 2026 Zeronetsec