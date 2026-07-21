// https://github.com/Zeronetsec/Znskit

package console

import (
    "github.com/Zeronetsec/Znskit/module/listins"
)

type ListInstalled struct{}
func (c ListInstalled) Execute(args []string) {
    listins.Show()
}

// Copyright (c) 2026 Zeronetsec