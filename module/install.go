// https://github.com/Zeronetsec/Znskit

package module

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/Zeronetsec/Znskit/utils"
    "github.com/Zeronetsec/Znskit/utils/color"
)

func Install(pkg string) {
    if pkg == "" {
        fmt.Printf(
            "%s[!] %sMissing package name!\n",
            color.R, color.N,
        )
        return
    }

    if !utils.CheckConnection() {
        fmt.Printf(
            "%s[!] %sNo internet connection!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[%s1%s/%s6%s] %sEnsuring directories...\n",
        color.DG, color.GG, color.DG, color.GG, color.DG, color.N,
    )

    if err := EnsureDirs(); err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[%s2%s/%s6%s] %sChecking installed package...\n",
        color.DG, color.GG, color.DG, color.GG, color.DG, color.N,
    )

    if CheckPackage(pkg) {
        fmt.Printf(
            "%s[!] %sPackage: %s%s %sis already installed!\n",
            color.R, color.N, color.GG, pkg, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[%s3%s/%s6%s] %sChecking internet connection...\n",
        color.DG, color.GG, color.DG, color.GG, color.DG, color.N,
    )

    fmt.Printf(
        "%s[%s4%s/%s6%s] %sValidating package...\n",
        color.DG, color.GG, color.DG, color.GG, color.DG, color.N,
    )

    repo, err := getRepository(pkg)
    if err != nil {
        fmt.Printf(
            "%s[!] %sPackage: %s%s %snot found!\n",
            color.R, color.N, color.GG, pkg, color.N,
        )
        return
    }

    fmt.Printf(
        "%s- %sFound package: %s%s%s\n",
        color.DG, color.N, color.GG, repo.Name, color.N,
    )

    fmt.Printf(
        "    %s└── %s%s%s\n",
        color.DG, color.GG, repo.URL, color.N,
    )

    targetDir := repo.Name
    if _, err := os.Stat(targetDir); err == nil {
        fmt.Printf(
            "%s[*] %sExisting folder found, removing...\n",
            color.R, color.N,
        )

        if err := os.RemoveAll(targetDir); err != nil {
            fmt.Printf(
                "%s[!] %sFailed to remove old installation: %s%v%s\n",
                color.R, color.N, color.GG, err, color.N,
            )
            return
        }
    }

    fmt.Printf(
        "%s[%s5%s/%s6%s] %sCloning repository...\n",
        color.DG, color.GG, color.DG, color.GG, color.DG, color.N,
    )

    if err := runCommand(
        "git", "clone", repo.URL, targetDir,
    ); err != nil {
        fmt.Printf(
            "%s[!] %sClone failed: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[%s6%s/%s6%s] %sRunning install script...\n",
        color.DG, color.GG, color.DG, color.GG, color.DG, color.N,
    )

    installPath := filepath.Join(
        targetDir, "install.sh",
    )

    if _, err := os.Stat(installPath); err != nil {
        fmt.Printf(
            "%s[!] %sScript: %sinstall.sh %snot found!\n",
            color.R, color.N, color.GG, color.N,
        )
        return
    }

    _ = runCommand("chmod", "+x", installPath)

    if err := runCommand("bash", installPath); err != nil {
        fmt.Printf(
            "%s[!] %sInstallation failed: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[+] %sInstallation completed!\n",
        color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec