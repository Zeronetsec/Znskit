// https://github.com/Zeronetsec/Znskit

package list

import (
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/Zeronetsec/Znskit/utils"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Show(isDetails bool) {
    if !utils.CheckConnection() {
        fmt.Printf(
            "%s[!] %sNo internet connection!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    url := "https://api.github.com/users/Zeronetsec/repos"
    req, err := http.NewRequest("GET", url, nil)
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

    if resp.StatusCode != http.StatusOK {
        fmt.Printf(
            "%s[!] %sGitHub API error: %s%s%s\n",
            color.R, color.N, color.GG, resp.Status, color.N,
        )
        os.Exit(1)
    }

    var repos []Repo
    if err := json.NewDecoder(
        resp.Body,
    ).Decode(&repos); err != nil {
        fmt.Printf(
            "%s[!] %sFailed to processing json data!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    for _, r := range repos {
        language := "?"
        if r.Language != nil {
            language = *r.Language
        }

        description := "?"
        if r.Description != nil {
            description = *r.Description
        }

        license := "?"
        if r.License != nil && r.License.Name != "" {
            license = r.License.Name
        }

        if isDetails {
            fmt.Printf(
                "%s%s%s/%s%s%s\n",
                color.GG, r.Name, color.DG,
                color.BB, r.SvnURL, color.N,
            )

            fmt.Printf(
                "%s├── %sbranch: %s%s%s, %slicense: %s%s%s, %slanguage: %s%s%s\n",
                color.DG, color.WW, color.GG, r.DefaultBranch, color.DG,
                color.WW, color.YY, license, color.DG,
                color.WW, color.CC, language, color.N,
            )

            fmt.Printf(
                "%s└── %s%s%s\n",
                color.DG, color.WW, description, color.N,
            )

            fmt.Println()
        } else {
            fmt.Printf(
                "%s%s%s/%s%s%s\n",
                color.GG, r.Name, color.DG,
                color.BB, r.SvnURL, color.N,
            )
        }
    }
}

// Copyright (c) 2026 Zeronetsec