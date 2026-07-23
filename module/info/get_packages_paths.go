// https://github.com/Zeronetsec/Znskit

package info

import (
    "os"
    "bufio"
    "strings"
    "path/filepath"
)

func getPackagePaths() []string {
    defaultPath := []string{".install/packages.txt"}
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return defaultPath
    }

    lstPath := filepath.Join(
        homeDir, ".znskit", "packages.lst",
    )

    file, err := os.Open(lstPath)
    if err != nil {
        return defaultPath
    }
    defer file.Close()

    var paths []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" && !strings.HasPrefix(
            line, "#",
        ) {
            paths = append(paths, line)
        }
    }

    if len(paths) == 0 {
        return defaultPath
    }

    return paths
}

// Copyright (c) 2026 Zeronetsec