// https://github.com/Zeronetsec/Znskit

package main

import (
    "os"
    "strings"
    "github.com/Zeronetsec/Znskit/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.ZnsConsole(input)
}

// Copyright (c) 2026 Zeronetsec