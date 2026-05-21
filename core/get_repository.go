// https://github.com/Zeronetsec/Znskit

package core

import (
    "embed"
    "fmt"
    "path"
    "strings"
    "encoding/json"
)

//go:embed repositories/*.json
var repositoryFS embed.FS

func getRepository(pkg string) (*Repository, error) {
    if pkg == "" || strings.Contains(pkg, "..") || strings.Contains(pkg, "/") {
        return nil, fmt.Errorf(
            "Invalid package name!",
        )
    }

    filename := fmt.Sprintf("%s.json", pkg)
    finalPath := path.Join("repositories", filename)

    data, err := repositoryFS.ReadFile(finalPath)
    if err != nil {
        return nil, err
    }

    var repo Repository
    if err := json.Unmarshal(data, &repo); err != nil {
        return nil, err
    }

    return &repo, nil
}

// Copyright (c) 2026 Zeronetsec