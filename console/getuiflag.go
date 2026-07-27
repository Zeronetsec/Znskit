// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "github.com/Zeronetsec/Znskit/module/getuiflag"
    "github.com/Zeronetsec/Znskit/utils/invinput"
)

type Getuiflag struct{}
func (c Getuiflag) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    toolName := args[2]
    getuiflag.Fetch(toolName)
}

// Copyright (c) 2026 Zeronetsec