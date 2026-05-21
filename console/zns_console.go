// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "znskit/utils/invinput"
)

func ZnsConsole(input string) {
    args := os.Args
    if len(args) < 2 {
        invinput.Invalid()
        return
    }

    commands := map[string]Command{
        "--help": Helper{},
        "--version": Version{},
        "--uwu": UWU{},
        "--install": Install{},
        "--uninstall": Uninstall{},
        "--search": Search{},
        "--list": Lister{},
        "--list-installed": ListInstalled{},
        "--reinstall": Reinstall{},
        "--info": Info{},
    }

    if cmd, ok := commands[args[1]]; ok {
        cmd.Execute(args)
    } else {
        invinput.Unknown(args[1])
    }
}

// Copyright (c) 2026 Zeronetsec