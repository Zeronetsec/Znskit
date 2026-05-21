// https://github.com/Zeronetsec/Znskit

package core

import (
    "os"
    "os/exec"
)

func runCommand(command string, args ...string) error {
    cmd := exec.Command(command, args...)

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin

    return cmd.Run()
}

// Copyright (c) 2026 Zeronetsec