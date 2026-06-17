// https://github.com/Zeronetsec/Znskit

package console

import (
    "github.com/Zeronetsec/Znskit/module"
)

type Lister struct{}
func (c Lister) Execute(args []string) {
    module.List()
}

// Copyright (c) 2026 Zeronetsec