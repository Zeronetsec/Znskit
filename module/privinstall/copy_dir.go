// https://github.com/Zeronetsec/Znskit

package privinstall

import (
    "os"
    "path/filepath"
)

func copyDir(src string, dst string) error {
    srcInfo, err := os.Stat(src)
    if err != nil {
        return err
    }

    if err := os.MkdirAll(
        dst, srcInfo.Mode(),
    ); err != nil {
        return err
    }

    entries, err := os.ReadDir(src)
    if err != nil {
        return err
    }

    for _, entry := range entries {
        srcPath := filepath.Join(src, entry.Name())
        dstPath := filepath.Join(dst, entry.Name())

        if entry.IsDir() {
            if err := copyDir(
                srcPath, dstPath,
            ); err != nil {
                return err
            }
        } else {
            if err := copyFile(
                srcPath, dstPath,
            ); err != nil {
                return err
            }
        }
    }

    return nil
}

// Copyright (c) 2026 Zeronetsec