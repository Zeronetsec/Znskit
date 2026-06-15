// https://github.com/Zeronetsec/Znskit

package module

import (
    "fmt"
    "strings"
    "path/filepath"
    "encoding/json"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func ListInstalled() {
    files, err := repositoryFS.ReadDir("repositories")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError reading repositories: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    prefix := GetPrefix()
    found := false

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        if !strings.HasSuffix(file.Name(), ".json") {
            continue
        }

        path := fmt.Sprintf("repositories/%s", file.Name())
        data, err := repositoryFS.ReadFile(path)
        if err != nil {
            continue
        }

        var repo Repository

        if err := json.Unmarshal(data, &repo); err != nil {
            continue
        }

        sourcePath := filepath.Join(prefix, "opt", repo.Name)
        executablePath := filepath.Join(prefix, "bin", repo.Name)

        sourceExists := exists(sourcePath)
        executableExists := exists(executablePath)

        if sourceExists || executableExists {

            found = true

            fmt.Printf(
                "%s[*] %sPackage: %s%s %sinstalled:\n",
                color.B, color.N, color.GG, repo.Name, color.N,
            )

            if sourceExists {
                fmt.Printf(
                    "    %s├── %sSource: %s%s%s\n",
                    color.DG, color.N, color.GG, sourcePath, color.N,
                )
            } else {
                fmt.Printf(
                    "    %s├── %sSource: %smissing%s\n",
                    color.DG, color.N, color.GG, color.N,
                )
            }

            if executableExists {
                fmt.Printf(
                    "    %s└── %sExecutable: %s%s%s\n\n",
                    color.DG, color.N, color.GG, executablePath, color.N,
                )
            } else {
                fmt.Printf(
                    "    %s└── %sExecutable: %smissing%s\n\n",
                    color.DG, color.N, color.GG, color.N,
                )
            }
        }
    }

    if !found {
        fmt.Printf(
            "%s[!] %sNo installed packages found!\n",
            color.R, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec