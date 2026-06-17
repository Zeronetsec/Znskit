// https://github.com/Zeronetsec/Znskit

package console

import (
    "github.com/Zeronetsec/Znskit/module"
)

type ListInstalled struct{}
func (c ListInstalled) Execute(args []string) {
    module.ListInstalled()
}

// Copyright (c) 2026 Zeronetsec