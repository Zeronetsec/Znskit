// https://github.com/Zeronetsec/Znskit

package console

import (
    "github.com/Zeronetsec/Znskit/module/list"
)

type Lister struct{}
func (c Lister) Execute(args []string) {
    isDetails := false
    if len(args) > 2 && args[2] == "details" {
        isDetails = true
    }

    list.Show(isDetails)
}

// Copyright (c) 2026 Zeronetsec