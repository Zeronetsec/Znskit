// https://github.com/Zeronetsec/Znskit

package module

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Uninstall(pkg string) {
    if pkg == "" {
        fmt.Printf(
            "%s[!] %sMissing package name!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[%s1%s/%s2%s] %sValidating package...\n",
        color.DG, color.GG, color.DG, color.GG, color.DG, color.N,
    )

    repo, err := getRepository(pkg)
    if err != nil {
        fmt.Printf(
            "%s[!] %sPackage: %s%s %snot found!\n",
            color.R, color.N, color.GG, pkg, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[%s2%s/%s2%s] %sRemoving package...\n",
        color.DG, color.GG, color.DG, color.GG, color.DG, color.N,
    )

    prefix := GetPrefix()
    paths := []string{
        filepath.Join(prefix, "bin", repo.Name),
        filepath.Join(prefix, "opt", repo.Name),
    }

    removed := false
    for _, path := range paths {
        if _, err := os.Lstat(path); err == nil {
            fmt.Printf(
                "%s[-] %sRemoving: %s%s%s\n",
                color.YY, color.N, color.GG, path, color.N,
            )

            if err := os.RemoveAll(path); err != nil {
                fmt.Printf(
                    "%s[!] %sFailed removing: %s%s %s(%s%v%s)%s\n",
                    color.R, color.N, color.GG, path, color.DG, color.CC, err, color.DG, color.N,
                )
                continue
            }
            removed = true
        }
    }

    if !removed {
        fmt.Printf(
            "%s[!] %sPackage: %s%s %sis not installed!\n",
            color.R, color.N, color.GG, repo.Name, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[+] %sSuccessfully uninstalled: %s%s%s\n",
        color.GG, color.N, color.GG, repo.Name, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec