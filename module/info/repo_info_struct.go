// https://github.com/Zeronetsec/Znskit

package info

type RepoInfo struct {
    Name string `json:"name"`
    Description *string `json:"description"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    PushedAt string `json:"pushed_at"`
    SvnURL string `json:"svn_url"`
    Size int `json:"size"`
    StargazersCount int `json:"stargazers_count"`
    Language *string `json:"language"`
    ForksCount int `json:"forks_count"`
    Topics []string `json:"topics"`
    DefaultBranch string `json:"default_branch"`
    License *struct {
        Name string `json:"name"`
    } `json:"license"`
}

// Copyright (c) 2026 Zeronetsec