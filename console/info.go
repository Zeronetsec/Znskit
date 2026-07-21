// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "github.com/Zeronetsec/Znskit/module/info"
    "github.com/Zeronetsec/Znskit/utils/invinput"
)

type Info struct{}
func (c Info) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    toolName := args[2]
    info.Show(toolName)
}

// Copyright (c) 2026 Zeronetsec