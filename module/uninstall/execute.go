// https://github.com/Zeronetsec/Znskit

package uninstall

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "github.com/Zeronetsec/Znskit/utils"
    "github.com/Zeronetsec/Znskit/utils/color"
    "github.com/Zeronetsec/Znskit/utils/validator"
)

func Execute(toolName string, uFlags []string) {
    if validator.NotInstalled(toolName) {
        fmt.Printf(
            "%s[!] %sTool: %s%s %sis not installed!\n",
            color.R, color.N, color.GG, toolName, color.N,
        )
        os.Exit(1)
    }

    prefix := utils.GetPrefix()
    optDir := filepath.Join(prefix, "opt", toolName)
    uninstallShPath := filepath.Join(
        optDir, "uninstall.sh",
    )

    if _, err := os.Stat(uninstallShPath); os.IsNotExist(err) {
        fmt.Printf(
            "%s[!] %sUninstaller not found in: %s%s%s\n",
            color.R, color.N, color.GG, uninstallShPath, color.N,
        )

        fmt.Printf(
            "%s[!] %sEnsure the tool: %s%s %sis installed!\n",
            color.R, color.N, color.GG, toolName, color.N,
        )
        os.Exit(1)
    }

    _ = os.Chmod(uninstallShPath, 0755)
    fmt.Printf(
        "%s[*] %sExecute: %suninstall.sh %s%v%s\n",
        color.B, color.N, color.GG, color.CC, uFlags, color.N,
    )

    cmdArgs := append([]string{uninstallShPath}, uFlags...)
    execCmd := exec.Command("bash", cmdArgs...)
    execCmd.Dir = optDir
    execCmd.Stdout = os.Stdout
    execCmd.Stderr = os.Stderr
    execCmd.Stdin = os.Stdin

    if err := execCmd.Run(); err != nil {
        fmt.Printf(
            "%s[!] %sUninstallation failed: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        os.Exit(1)
    }
}

// Copyright (c) 2026 Zeronetsec