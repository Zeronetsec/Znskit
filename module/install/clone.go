// https://github.com/Zeronetsec/Znskit

package install

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "encoding/json"
    "net/http"
    "github.com/Zeronetsec/Znskit/utils"
    "github.com/Zeronetsec/Znskit/utils/color"
    "github.com/Zeronetsec/Znskit/utils/validator"
)

func Clone(toolName string, iFlags []string) {
    if validator.Installed(toolName) {
        fmt.Printf(
            "%s[!] %sTool: %s%s %sis already installed!\n",
            color.R, color.N, color.GG, toolName, color.N,
        )
        os.Exit(1)
    }

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

    prefix := utils.GetPrefix()
    var targetDir string
    if prefix != "" {
        targetDir = filepath.Join(
            prefix, "tmp", toolName,
        )
    } else {
        targetDir = filepath.Join(
            "/tmp", toolName,
        )
    }

    if _, err := os.Stat(targetDir); err == nil {
        os.RemoveAll(targetDir)
    }

    fmt.Printf(
        "%s[*] %sCloning: %s%s %s-> %s%s%s\n",
        color.B, color.N, color.GG, repo.CloneURL, color.DG,
        color.GG, targetDir, color.N,
    )

    cloneCmd := exec.Command(
        "git", "clone", repo.CloneURL, targetDir,
    )

    cloneCmd.Stdout = os.Stdout
    cloneCmd.Stderr = os.Stderr

    if err := cloneCmd.Run(); err != nil {
        fmt.Printf(
            "%s[!] %sGit clone error!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    installShPath := filepath.Join(
        targetDir, "install.sh",
    )

    if _, err := os.Stat(installShPath); os.IsNotExist(err) {
        fmt.Printf(
            "%s[!] %sFile: %sinstall.sh %snot found in %s%s%s\n",
            color.R, color.N, color.GG, color.N,
            color.GG, targetDir, color.N,
        )
        os.Exit(1)
    }

    if err := os.Chmod(installShPath, 0755); err != nil {
        fmt.Printf(
            "%s[!] %sFailed to set permission for: %sinstall.sh %s(%s%v%s)%s\n",
            color.R, color.N, color.GG, color.DG,
            color.CC, err, color.DG, color.N,
        )
        os.Exit(1)
    }

    fmt.Printf(
        "%s[*] %sExecute: %sinstall.sh %s%v%s\n",
        color.B, color.N, color.GG, color.CC, iFlags, color.N,
    )

    cmdArgs := append([]string{installShPath}, iFlags...)

    execCmd := exec.Command("bash", cmdArgs...)
    execCmd.Dir = targetDir
    execCmd.Stdout = os.Stdout
    execCmd.Stderr = os.Stderr
    execCmd.Stdin = os.Stdin

    if err := execCmd.Run(); err != nil {
        fmt.Printf(
            "%s[!] %sInstallation failed: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        os.Exit(1)
    }

    _ = os.RemoveAll(targetDir)
}

// Copyright (c) 2026 Zeronetsec