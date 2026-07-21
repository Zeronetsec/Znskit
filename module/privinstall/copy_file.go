// https://github.com/Zeronetsec/Znskit

package privinstall

import (
    "os"
    "io"
)

func copyFile(src string, dst string) error {
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    if _, err = io.Copy(out, in); err != nil {
        return err
    }

    info, err := os.Stat(src)
    if err == nil {
        _ = os.Chmod(dst, info.Mode())
    }

    return nil
}

// Copyright (c) 2026 Zeronetsec