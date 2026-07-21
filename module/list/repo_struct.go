// https://github.com/Zeronetsec/Znskit

package list

type Repo struct {
    Name string `json:"name"`
    SvnURL string `json:"svn_url"`
    DefaultBranch string `json:"default_branch"`
    Language *string `json:"language"`
    Description *string `json:"description"`
    License *struct {
        Name string `json:"name"`
    } `json:"license"`
}

// Copyright (c) 2026 Zeronetsec