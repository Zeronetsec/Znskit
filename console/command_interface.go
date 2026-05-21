// https://github.com/Zeronetsec/Znskit

package console

import (
    "fmt"
    "time"
    "znskit/core"
    "znskit/utils/invinput"
)

type Command interface {
    Execute(args []string)
}

type Helper struct{}
func (c Helper) Execute(args []string) {
    core.Help()
}

type Version struct{}
func (c Version) Execute(args []string) {
    core.Version()
}

type UWU struct{}
func (c UWU) Execute(args []string) {
    fmt.Print("\033[?25l")
    core.Uwu(5 * time.Second)
    fmt.Print("\033[?25h")
}

type Install struct{}
func (c Install) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    core.Install(args[2])
}

type Uninstall struct{}
func (c Uninstall) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    core.Uninstall(args[2])
}

type Search struct{}
func (c Search) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    core.Search(args[2])
}

type Lister struct{}
func (c Lister) Execute(args []string) {
    core.List()
}

type ListInstalled struct{}
func (c ListInstalled) Execute(args []string) {
    core.ListInstalled()
}

type Reinstall struct{}
func (c Reinstall) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    core.Reinstall(args[2])
}

type Info struct{}
func (c Info) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    core.Info(args[2])
}

// Copyright (c) 2026 Zeronetsec