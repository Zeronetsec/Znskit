// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "github.com/Zeronetsec/Znskit/utils/invinput"
)

func ZnsConsole(input string) {
    args := os.Args
    if len(args) < 2 {
        invinput.MissingArgument()
        os.Exit(1)
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
        "--private-install": PrivInstall{},
        "--private-reinstall": PrivReinstall{},
    }

    if cmd, ok := commands[args[1]]; ok {
        cmd.Execute(args)
    } else {
        invinput.InvalidOption(args[1])
        os.Exit(1)
    }
}

// Copyright (c) 2026 Zeronetsec