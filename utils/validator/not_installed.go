// https://github.com/Zeronetsec/Znskit

package validator

func NotInstalled(toolName string) bool {
    return !Installed(toolName)
}

// Copyright (c) 2026 Zeronetsec