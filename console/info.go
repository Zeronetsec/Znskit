// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "github.com/Zeronetsec/Znskit/module"
    "github.com/Zeronetsec/Znskit/utils/invinput"
)

type Info struct{}
func (c Info) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    module.Info(args[2])
}

// Copyright (c) 2026 Zeronetsec