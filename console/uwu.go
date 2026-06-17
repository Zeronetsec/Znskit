// https://github.com/Zeronetsec/Znskit

package console

import (
    "time"
    "fmt"
    "github.com/Zeronetsec/Znskit/module"
)

type UWU struct{}
func (c UWU) Execute(args []string) {
    fmt.Print("\x1b[?25l")
    module.Uwu(5 * time.Second)
    fmt.Print("\x1b[?25h")
}

// Copyright (c) 2026 Zeronetsec