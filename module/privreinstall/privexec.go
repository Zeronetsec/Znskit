// https://github.com/Zeronetsec/Znskit

package privreinstall

import (
    "fmt"
    "github.com/Zeronetsec/Znskit/module/privinstall"
    "github.com/Zeronetsec/Znskit/module/uninstall"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func PrivExec(toolName string, privDataDir string, iFlags []string, uFlags []string) {
    fmt.Printf(
        "%s[*] %sReinstalling: %s%s%s\n",
        color.B, color.N, color.GG, toolName, color.N,
    )

    uninstall.Execute(toolName, uFlags)
    privinstall.PrivClone(toolName, privDataDir, iFlags)
}

// Copyright (c) 2026 Zeronetsec