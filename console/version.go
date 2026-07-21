// https://github.com/Zeronetsec/Znskit

package console

import (
    "github.com/Zeronetsec/Znskit/module/version"
)

type Version struct{}
func (c Version) Execute(args []string) {
    version.Show()
}

// Copyright (c) 2026 Zeronetsec