// https://github.com/Zeronetsec/Znskit

package module

import (
    "fmt"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Info(pkg string) {
    if pkg == "" {
        fmt.Printf(
            "%s[!] %sMissing package name!\n",
            color.R, color.N,
        )
        return
    }

    repo, err := getRepository(pkg)
    if err != nil {
        fmt.Printf(
            "%s[!] %sPackage: %s%s %snot found!\n",
            color.R, color.N, color.GG, pkg, color.N,
        )
        return
    }

    fmt.Printf(
        "%sName: %s%s%s\n",
        color.N, color.GG, repo.Name, color.N,
    )

    fmt.Printf(
        "%sURL: %s%s%s\n",
        color.N, color.GG, repo.URL, color.N,
    )

    fmt.Printf(
        "%sDescription: %s%s%s\n",
        color.N, color.GG, repo.Description, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec