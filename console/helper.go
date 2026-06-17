// https://github.com/Zeronetsec/Znskit

package console

import (
    "github.com/Zeronetsec/Znskit/module"
)

type Helper struct{}
func (c Helper) Execute(args []string) {
    module.Help()
}

// Copyright (c) 2026 Zeronetsec