// https://github.com/Zeronetsec/Znskit

package module

import (
    "os"
    "path/filepath"
)

func CheckPackage(pkg string) bool {
    prefix := GetPrefix()
    sourcePath := filepath.Join(prefix, "opt", pkg)
    executablePath := filepath.Join(prefix, "bin", pkg)

    _, sourceErr := os.Lstat(sourcePath)
    _, executableErr := os.Lstat(executablePath)

    return sourceErr == nil || executableErr == nil
}

// Copyright (c) 2026 Zeronetsec