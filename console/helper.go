// https://github.com/Zeronetsec/Znskit

package console

import (
    "github.com/Zeronetsec/Znskit/module/help"
)

type Helper struct{}
func (c Helper) Execute(args []string) {
    help.Show()
}

// Copyright (c) 2026 Zeronetsec