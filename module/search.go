// https://github.com/Zeronetsec/Znskit

package module

import (
    "strings"
    "fmt"
    "encoding/json"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Search(query string) {
    if query == "" {
        fmt.Printf(
            "%s[!] %sMissing search query!\n",
            color.R, color.N,
        )
        return
    }

    files, err := repositoryFS.ReadDir("repositories")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError reading repositories: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    found := false
    query = strings.ToLower(query)

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

        name := strings.ToLower(repo.Name)
        description := strings.ToLower(repo.Description)

        if strings.Contains(name, query) || strings.Contains(description, query) {
            found = true
            fmt.Printf(
                "%s%s%s/%s%s%s\n",
                color.GG, repo.Name, color.N, color.BB, repo.URL, color.N,
            )

            fmt.Printf(
                "    %s└── %s%s%s\n\n",
                color.DG, color.WW, repo.Description, color.N,
            )
        }
    }

    if !found {
        fmt.Printf(
            "%s[!] %sNo packages found for: %s%s%s\n",
            color.R, color.N, color.GG, query, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec