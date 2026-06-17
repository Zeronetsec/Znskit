// https://github.com/Zeronetsec/Znskit

package console

import (
    "github.com/Zeronetsec/Znskit/module"
)

type Version struct{}
func (c Version) Execute(args []string) {
    module.Version()
}

// Copyright (c) 2026 Zeronetsec