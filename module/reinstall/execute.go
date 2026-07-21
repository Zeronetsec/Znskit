// https://github.com/Zeronetsec/Znskit

package reinstall

import (
    "fmt"
    "github.com/Zeronetsec/Znskit/module/install"
    "github.com/Zeronetsec/Znskit/module/uninstall"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Execute(toolName string, iFlags []string, uFlags []string) {
    fmt.Printf(
        "%s[*] %sReinstalling: %s%s%s\n",
        color.B, color.N, color.GG, toolName, color.N,
    )

    uninstall.Execute(toolName, uFlags)
    install.Clone(toolName, iFlags)
}

// Copyright (c) 2026 Zeronetsec