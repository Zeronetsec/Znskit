// https://github.com/Zeronetsec/Znskit

package getuiflag

import (
    "fmt"
    "os"
    "os/exec"
    "net/http"
    "encoding/json"
    "github.com/Zeronetsec/Znskit/utils"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Fetch(toolName string) {
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

    var repo Repo
    if err := json.NewDecoder(
        resp.Body,
    ).Decode(&repo); err != nil {
        fmt.Printf(
            "%s[!] %sFailed to processing json data!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    rawUrl := fmt.Sprintf(
        "https://raw.githubusercontent.com/Zeronetsec/%s/%s/.gitaction/uiflag_params.flg",
        repo.Name, repo.DefaultBranch,
    )

    flgReq, _ := http.NewRequest("GET", rawUrl, nil)
    flgReq.Header.Set(
        "User-Agent",
        "https://github.com/Zeronetsec/Znskit",
    )

    flgResp, err := client.Do(flgReq)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to fetch uiflag params data!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }
    defer flgResp.Body.Close()

    if flgResp.StatusCode == http.StatusNotFound {
        fmt.Printf(
            "%s[!] %sFile: %s.gitaction/uiflag_params.flg %snot found in tool %s%s%s\n",
            color.R, color.N, color.GG, color.N, color.GG, toolName, color.N,
        )
        os.Exit(1)
    } else if flgResp.StatusCode != http.StatusOK {
        fmt.Printf(
            "%s[!] %sFailed to download uiflag params data: %s%s%s\n",
            color.R, color.N, color.GG, flgResp.Status, color.N,
        )
        os.Exit(1)
    }

    execCmd := exec.Command("bash")
    execCmd.Stdin = flgResp.Body
    execCmd.Stdout = os.Stdout
    execCmd.Stderr = os.Stderr

    if err := execCmd.Run(); err != nil {
        fmt.Printf(
            "%s[!] %sFailed to execute uiflag params: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        os.Exit(1)
    }
}

// Copyright (c) 2026 Zeronetsec