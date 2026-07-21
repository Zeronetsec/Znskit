// https://github.com/Zeronetsec/Znskit

package listins

import (
    "os"
    "path/filepath"
)

func getDirStats(dirPath string) (
    files int,
    folders int,
    totalSize int64,
    exists bool,
) {
    _, err := os.Stat(dirPath)
    if err != nil {
        return 0, 0, 0, false
    }

    _ = filepath.WalkDir(
        dirPath, func(
            path string, d os.DirEntry, err error,
        ) error {
            if err != nil {
                return nil
            }

            if path == dirPath {
                return nil
            }

            if d.IsDir() {
                folders++
            } else {
                files++
                info, err := d.Info()
                if err == nil {
                    totalSize += info.Size()
                }
            }

            return nil
        },
    )

    return files, folders, totalSize, true
}

// Copyright (c) 2026 Zeronetsec