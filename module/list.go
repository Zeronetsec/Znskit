// https://github.com/Zeronetsec/Znskit

package module

import (
    "fmt"
    "strings"
    "encoding/json"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func List() {
    files, err := repositoryFS.ReadDir("repositories")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError reading repositories: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

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

        found = true
        fmt.Printf(
            "%s%s%s/%s%s%s\n",
            color.GG, repo.Name, color.N, color.BB, repo.URL, color.N,
        )
    }

    if !found {
        fmt.Printf(
            "%s[!] %sNo packages available!\n",
            color.R, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec