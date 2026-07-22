// https://github.com/Zeronetsec/Znskit

package info

import (
    "fmt"
    "strings"
    "encoding/json"
    "net/http"
    "github.com/Zeronetsec/Znskit/utils"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Show(toolName string) {
    if !utils.CheckConnection() {
        fmt.Printf(
            "%s[!] %sNo internet connection!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    apiUrl := fmt.Sprintf(
        "https://api.github.com/repos/Zeronetsec/%s",
        toolName,
    )

    req, err := http.NewRequest("GET", apiUrl, nil)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to create request: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        os.Exit(1)
    }

    req.Header.Set(
        "User-Agent",
        "https://github.com/Zeronetsec/Znskit",
    )

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to connect to API!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
        fmt.Printf(
            "%s[!] %sTool: %s%s %snot found!\n",
            color.R, color.N, color.GG, toolName, color.N,
        )
        os.Exit(1)
    } else if resp.StatusCode != http.StatusOK {
        fmt.Printf(
            "%s[!] %sGitHub API error: %s%s%s\n",
            color.R, color.N, color.GG, resp.Status, color.N,
        )
        os.Exit(1)
    }

    var repo RepoInfo
    if err := json.NewDecoder(
        resp.Body,
    ).Decode(&repo); err != nil {
        fmt.Printf(
            "%s[!] %sFailed to processing json data!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    language := "?"
    if repo.Language != nil {
        language = *repo.Language
    }

    description := "?"
    if repo.Description != nil {
        description = *repo.Description
    }

    license := "?"
    if repo.License != nil && repo.License.Name != "" {
        license = repo.License.Name
    }

    topicsStr := "?"
    if len(repo.Topics) > 0 {
        topicsStr = strings.Join(repo.Topics, ", ")
    }

    createdAt := formatDate(repo.CreatedAt)
    updatedAt := formatDate(repo.UpdatedAt)
    pushedAt := formatDate(repo.PushedAt)

    dependencies := fetchDependencies(
        repo.Name, repo.DefaultBranch,
    )

    fmt.Printf(
        "%s%s%s/%s%s%s\n",
        color.GG, repo.Name, color.DG,
        color.BB, repo.SvnURL, color.N,
    )

    fmt.Printf(
        "%s├── %sbranch: %s%s%s, %slicense: %s%s%s, %slanguage: %s%s%s\n",
        color.DG, color.WW, color.GG, repo.DefaultBranch, color.DG,
        color.WW, color.YY, license, color.DG,
        color.WW, color.CC, language, color.N,
    )

    fmt.Printf(
        "%s├── %ssize: %s%s%s\n",
        color.DG, color.WW, color.GG,
        formatSize(repo.Size), color.N,
    )

    fmt.Printf(
        "%s├── %screated_at: %s%s%s, %supdated_at: %s%s%s, %spushed_at: %s%s%s\n",
        color.DG, color.WW, color.GG, createdAt, color.DG,
        color.WW, color.YY, updatedAt, color.DG,
        color.WW, color.CC, pushedAt, color.N,
    )

    fmt.Printf(
        "%s├── %sstar: %s%d%s, %sfork: %s%d%s\n",
        color.DG, color.WW, color.GG, repo.StargazersCount, color.DG,
        color.WW, color.YY, repo.ForksCount, color.N,
    )

    fmt.Printf(
        "%s├── %stopic: %s%s%s\n",
        color.DG, color.WW, color.GG, topicsStr, color.N,
    )

    fmt.Printf(
        "%s├── %sdependency: %s%s%s\n",
        color.DG, color.WW, color.GG, dependencies, color.N,
    )

    fmt.Printf(
        "%s└── %s%s%s\n",
        color.DG, color.WW, description, color.N,
    )

    fmt.Println()
}

// Copyright (c) 2026 Zeronetsec