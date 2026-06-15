// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "fmt"
    "time"
    "github.com/Zeronetsec/Znskit/module"
    "github.com/Zeronetsec/Znskit/utils/invinput"
)

type Command interface {
    Execute(args []string)
}

type Helper struct{}
func (c Helper) Execute(args []string) {
    module.Help()
}

type Version struct{}
func (c Version) Execute(args []string) {
    module.Version()
}

type UWU struct{}
func (c UWU) Execute(args []string) {
    fmt.Print("\x1b[?25l")
    module.Uwu(5 * time.Second)
    fmt.Print("\x1b[?25h")
}

type Install struct{}
func (c Install) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    module.Install(args[2])
}

type Uninstall struct{}
func (c Uninstall) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    module.Uninstall(args[2])
}

type Search struct{}
func (c Search) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    module.Search(args[2])
}

type Lister struct{}
func (c Lister) Execute(args []string) {
    module.List()
}

type ListInstalled struct{}
func (c ListInstalled) Execute(args []string) {
    module.ListInstalled()
}

type Reinstall struct{}
func (c Reinstall) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    module.Reinstall(args[2])
}

type Info struct{}
func (c Info) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    module.Info(args[2])
}

// Copyright (c) 2026 Zeronetsec