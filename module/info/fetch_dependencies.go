// https://github.com/Zeronetsec/Znskit

package info

import (
    "fmt"
    "bufio"
    "strings"
    "net/http"
)

func fetchDependencies(toolName, branch string) string {
    url := fmt.Sprintf(
        "https://raw.githubusercontent.com/Zeronetsec/%s/%s/.install/packages.txt",
        toolName, branch,
    )

    resp, err := http.Get(url)
    if err != nil || resp.StatusCode != http.StatusOK {
        return "?"
    }
    defer resp.Body.Close()

    var deps []string
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" && !strings.HasPrefix(line, "#") {
            deps = append(deps, line)
        }
    }

    if len(deps) == 0 {
        return "?"
    }
    return strings.Join(deps, ", ")
}

// Copyright (c) 2026 Zeronetsec