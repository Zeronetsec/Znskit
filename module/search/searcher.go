// https://github.com/Zeronetsec/Znskit

package search

import (
    "fmt"
    "strings"
    "net/url"
    "net/http"
    "encoding/json"
    "github.com/Zeronetsec/Znskit/utils"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Searcher(keyword string) {
    if !utils.CheckConnection() {
        fmt.Printf(
            "%s[!] %sNo internet connection!\n",
            color.R, color.N,
        )
        return
    }

    safeKeyword := url.QueryEscape(keyword)
    apiUrl := fmt.Sprintf(
        "https://api.github.com/search/repositories?q=user:Zeronetsec+%s",
        safeKeyword,
    )

    req, err := http.NewRequest("GET", apiUrl, nil)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to create request: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
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
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf(
            "%s[!] %sGitHub API error: %s%s%s\n",
            color.R, color.N, color.GG, resp.Status, color.N,
        )
        return
    }

    var result SearchResponse
    if err := json.NewDecoder(
        resp.Body,
    ).Decode(&result); err != nil {
        fmt.Printf(
            "%s[!] %sFailed to processing json data!\n",
            color.R, color.N,
        )
        return
    }

    if len(result.Items) == 0 {
        fmt.Printf(
            "%s[!] %sTool with keyword: %s%s %snot found!\n",
            color.R, color.N, color.GG, keyword, color.N,
        )
        return
    }

    for _, r := range result.Items {
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

        topicsStr := "?"
        if len(r.Topics) > 0 {
            topicsStr = strings.Join(r.Topics, ", ")
        }

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
            "%s├── %stopic: %s%s%s\n",
            color.DG, color.WW, color.GG, topicsStr, color.N,
        )

        fmt.Printf(
            "%s└── %s%s%s\n",
            color.DG, color.WW, description, color.N,
        )

        fmt.Println()
    }
}

// Copyright (c) 2026 Zeronetsec