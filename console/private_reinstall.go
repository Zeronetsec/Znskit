// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "strings"
    "github.com/Zeronetsec/Znskit/module/privreinstall"
    "github.com/Zeronetsec/Znskit/utils/invinput"
)

type PrivReinstall struct{}
func (c PrivReinstall) Execute(args []string) {
    if len(args) < 4 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    toolName := args[2]
    privDataDir := args[3]

    var iFlags []string
    var uFlags []string

    for i := 4; i < len(args); i++ {
        if args[i] == "--iflag" && i+1 < len(args) {
            rawFlags := args[i+1]
            if strings.TrimSpace(rawFlags) != "" {
                iFlags = strings.Fields(rawFlags)
            }
            i++
        } else if args[i] == "--uflag" && i+1 < len(args) {
            rawFlags := args[i+1]
            if strings.TrimSpace(rawFlags) != "" {
                uFlags = strings.Fields(rawFlags)
            }
            i++
        }
    }

    privreinstall.PrivExec(
        toolName, privDataDir, iFlags, uFlags,
    )
}

// Copyright (c) 2026 Zeronetsec