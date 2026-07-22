// https://github.com/Zeronetsec/Znskit

package privinstall

import (
    "fmt"
    "os"
    "strings"
    "os/exec"
    "net/http"
    "path/filepath"
    "github.com/Zeronetsec/Znskit/utils"
    "github.com/Zeronetsec/Znskit/utils/color"
    "github.com/Zeronetsec/Znskit/utils/validator"
)

func PrivClone(
    toolName string,
    privDataDir string,
    iFlags []string,
) {
    if validator.Installed(toolName) {
        fmt.Printf(
            "%s[!] %sTool: %s%s %sis already installed!\n",
            color.R, color.N, color.GG, toolName, color.N,
        )
        os.Exit(1)
    }

    privInfo, err := os.Stat(privDataDir)
    if err != nil || !privInfo.IsDir() {
        fmt.Printf(
            "%s[!] %sPrivate data folder: %s%s %snot found!\n",
            color.R, color.N, color.GG, privDataDir, color.N,
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

    repoURL := fmt.Sprintf(
        "https://github.com/Zeronetsec/%s",
        toolName,
    )

    apiURL := fmt.Sprintf(
        "https://api.github.com/repos/Zeronetsec/%s",
        toolName,
    )

    req, _ := http.NewRequest("GET", apiURL, nil)
    req.Header.Set(
        "User-Agent",
        "https://github.com/Zeronetsec/Znskit",
    )

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil || resp.StatusCode != http.StatusOK {
        fmt.Printf(
            "%s[!] %sTool: %s%s %snot found!\n",
            color.R, color.N, color.GG, toolName, color.N,
        )
        os.Exit(1)
    }
    resp.Body.Close()

    prefix := utils.GetPrefix()
    tmpDir := filepath.Join(
        prefix, "tmp", strings.ToLower(toolName),
    )

    _ = os.RemoveAll(tmpDir)

    fmt.Printf(
        "%s[*] %sCloning: %s%s %s-> %s%s%s\n",
        color.B, color.N, color.GG, repoURL, color.DG,
        color.GG, tmpDir, color.N,
    )

    cloneCmd := exec.Command(
        "git", "clone", repoURL, tmpDir,
    )

    cloneCmd.Stdout = os.Stdout
    cloneCmd.Stderr = os.Stderr

    if err := cloneCmd.Run(); err != nil {
        fmt.Printf(
            "%s[!] %sGit clone error!\n",
            color.R, color.N, toolName,
        )
        os.Exit(1)
    }

    privFolderName := filepath.Base(
        filepath.Clean(privDataDir),
    )

    targetPrivPath := filepath.Join(
        tmpDir, privFolderName,
    )

    fmt.Printf(
        "%s[*] %sCopying private data: %s%s %s-> %s%s%s\n",
        color.B, color.N, color.GG, privDataDir, color.DG,
        color.GG, targetPrivPath, color.N,
    )

    if err := copyDir(privDataDir, targetPrivPath); err != nil {
        fmt.Printf(
            "%s[!] %sFailed to copy privat data: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        os.Exit(1)
    }

    installShPath := filepath.Join(tmpDir, "install.sh")
    if _, err := os.Stat(
        installShPath,
    ); os.IsNotExist(err) {
        fmt.Printf(
            "%s[!] %sFile: %sinstall.sh %snot found in %s%s%s\n",
            color.R, color.N, color.GG, color.N,
            color.GG, toolName, color.N,
        )
        os.Exit(1)
    }

    _ = os.Chmod(installShPath, 0755)
    fmt.Printf(
        "%s[*] %sExecute: %sinstall.sh %s%v%s\n",
        color.B, color.N, color.GG, color.CC, iFlags, color.N,
    )

    cmdArgs := append([]string{installShPath}, iFlags...)
    execCmd := exec.Command("bash", cmdArgs...)
    execCmd.Dir = tmpDir
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

    _ = os.RemoveAll(tmpDir)
}

// Copyright (c) 2026 Zeronetsec