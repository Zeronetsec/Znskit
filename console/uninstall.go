// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "github.com/Zeronetsec/Znskit/module"
    "github.com/Zeronetsec/Znskit/utils/invinput"
)

type Uninstall struct{}
func (c Uninstall) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    module.Uninstall(args[2])
}

// Copyright (c) 2026 Zeronetsec