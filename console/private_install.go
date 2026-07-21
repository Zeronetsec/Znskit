// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "strings"
    "github.com/Zeronetsec/Znskit/module/privinstall"
    "github.com/Zeronetsec/Znskit/utils/invinput"
)

type PrivInstall struct{}
func (c PrivInstall) Execute(args []string) {
    if len(args) < 4 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    toolName := args[2]
    privDataDir := args[3]
    var iFlags []string

    for i := 4; i < len(args); i++ {
        if args[i] == "--iflag" && i+1 < len(args) {
            rawFlags := args[i+1]
            if strings.TrimSpace(rawFlags) != "" {
                iFlags = strings.Fields(rawFlags)
            }
            break
        }
    }

    privinstall.PrivClone(toolName, privDataDir, iFlags)
}

// Copyright (c) 2026 Zeronetsec