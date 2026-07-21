// https://github.com/Zeronetsec/Znskit

package listins

import (
    "fmt"
    "strings"
    "path/filepath"
    "encoding/json"
    "net/http"
    "github.com/Zeronetsec/Znskit/utils"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Show() {
    if !utils.CheckConnection() {
        fmt.Printf(
            "%s[!] %sNo internet connection!\n",
            color.R, color.N,
        )
        return
    }

    url := "https://api.github.com/users/Zeronetsec/repos"
    req, err := http.NewRequest("GET", url, nil)
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

    var repos []Repo
    if err := json.NewDecoder(
        resp.Body,
    ).Decode(&repos); err != nil {
        fmt.Printf(
            "%s[!] %sFailed to processing json data!\n",
            color.R, color.N,
        )
        return
    }

    prefix := utils.GetPrefix()
    foundAny := false

    for _, r := range repos {
        toolName := strings.ToLower(r.Name)

        optPath := filepath.Join(prefix, "opt", toolName)
        binPath := filepath.Join(prefix, "bin", toolName)

        binSizeStr, binExists := getFileSize(binPath)
        filesCount, foldersCount, optSize, optExists := getDirStats(optPath)

        if binExists || optExists {
            foundAny = true

            var execDisplay string
            if binExists {
                execDisplay = fmt.Sprintf(
                    "%s%s %s(%ssize: %s%s%s)%s",
                    color.GG, binPath, color.DG,
                    color.WW, color.CC, binSizeStr, color.DG, color.N,
                )
            } else {
                execDisplay = fmt.Sprintf(
                    "%s%s %s(%snot found%s)%s",
                    color.GG, binPath, color.DG,
                    color.R, color.DG, color.N,
                )
            }

            var sourceDisplay string
            if optExists {
                sourceDisplay = fmt.Sprintf(
                    "%s%s %s(%sfiles: %s%d%s, %sfolders: %s%d%s, %ssize: %s%s%s)%s",
                    color.GG, optPath, color.DG,
                    color.WW, color.GG, filesCount, color.DG,
                    color.WW, color.YY, foldersCount, color.DG,
                    color.WW, color.CC, formatBytes(optSize), color.DG, color.N,
                )
            } else {
                sourceDisplay = fmt.Sprintf(
                    "%s%s %s(%snot found%s)%s",
                    color.GG, optPath, color.DG,
                    color.R, color.DG, color.N,
                )
            }

            fmt.Printf(
                "%s%s%s/%s%s%s\n",
                color.GG, toolName, color.DG,
                color.BB, r.SvnURL, color.N,
            )

            fmt.Printf(
                "%s├── %sexec: %s%s%s\n",
                color.DG, color.WW, color.N, execDisplay, color.N,
            )

            fmt.Printf(
                "%s└── %ssource: %s%s%s\n",
                color.DG, color.WW, color.N, sourceDisplay, color.N,
            )

            fmt.Println()
        }
    }

    if !foundAny {
        fmt.Printf(
            "%s[!] %sNo tools are installed!\n",
            color.R, color.N,
        )
        return
    }
}

// Copyright (c) 2026 Zeronetsec