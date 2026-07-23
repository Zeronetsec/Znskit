// https://github.com/Zeronetsec/Znskit

package info

import (
    "bufio"
    "fmt"
    "net/http"
    "strings"
)

func fetchDependencies(toolName, branch string) string {
    paths := getPackagePaths()
    uniqueDeps := make(map[string]bool)
    var depsList []string

    for _, path := range paths {
        cleanPath := strings.TrimPrefix(path, "/")
        url := fmt.Sprintf(
            "https://raw.githubusercontent.com/Zeronetsec/%s/%s/%s",
            toolName, branch, cleanPath,
        )

        resp, err := http.Get(url)
        if err != nil {
            continue
        }

        if resp.StatusCode != http.StatusOK {
            resp.Body.Close()
            continue
        }

        scanner := bufio.NewScanner(resp.Body)
        for scanner.Scan() {
            line := strings.TrimSpace(scanner.Text())
            if line != "" && !strings.HasPrefix(
                line, "#",
            ) {
                if !uniqueDeps[line] {
                    uniqueDeps[line] = true
                    depsList = append(depsList, line)
                }
            }
        }
        resp.Body.Close()
    }

    if len(depsList) == 0 {
        return "?"
    }

    return strings.Join(depsList, ", ")
}

// Copyright (c) 2026 Zeronetsec