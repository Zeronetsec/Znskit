// https://github.com/Zeronetsec/Znskit

package module

import (
    "os"
    "fmt"
    "path/filepath"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func EnsureDirs() error {
    prefix := GetPrefix()

    dirs := []string{
        filepath.Join(prefix, "opt"),
        filepath.Join(prefix, "bin"),
        filepath.Join(prefix, "tmp"),
    }

    for _, dir := range dirs {
        if _, err := os.Stat(dir); os.IsNotExist(err) {
            if err := os.MkdirAll(dir, 0755); err != nil {
                return fmt.Errorf(
                    "Failed creating directory: %s%s %s(%s%v%s)%s\n",
                    color.GG, dir, color.DG, color.CC, err, color.DG, color.N,
                )
            }

            fmt.Printf(
                "%s- %sCreating directory: %s%s%s\n",
                color.DG, color.N, color.GG, dir, color.N,
            )

        }
    }
    return nil
}

// Copyright (c) 2026 Zeronetsec