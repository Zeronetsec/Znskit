// https://github.com/Zeronetsec/Znskit

package validator

import (
    "os"
    "strings"
    "path/filepath"
    "github.com/Zeronetsec/Znskit/utils"
)

func Installed(toolName string) bool {
    tool := strings.ToLower(toolName)
    prefix := utils.GetPrefix()

    binPath := filepath.Join(prefix, "bin", tool)
    optPath := filepath.Join(prefix, "opt", tool)

    binInfo, err := os.Stat(binPath)
    if (err != nil ||
        binInfo.IsDir() ||
        binInfo.Size() <= 0) {
            return false
        }

    optInfo, err := os.Stat(optPath)
    if err != nil || !optInfo.IsDir() {
        return false
    }

    var filesCount int
    var foldersCount int
    var totalSize int64

    _ = filepath.WalkDir(
        optPath, func(
            path string,
            d os.DirEntry,
            err error,
        ) error {
            if err != nil {
                return nil
            }

            if path == optPath {
                return nil
            }

            if d.IsDir() {
                foldersCount++
            } else {
                filesCount++
                info, err := d.Info()
                if err == nil {
                    totalSize += info.Size()
                }
            }
            return nil
        },
    )

    if (totalSize <= 0 ||
        filesCount < 1 ||
        foldersCount < 1) {
            return false
        }

    return true
}

// Copyright (c) 2026 Zeronetsec