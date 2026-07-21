// https://github.com/Zeronetsec/Znskit

package console

import (
    "os"
    "strings"
    "github.com/Zeronetsec/Znskit/module/search"
    "github.com/Zeronetsec/Znskit/utils/invinput"
)

type Search struct{}
func (c Search) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    keyword := strings.Join(args[2:], " ")
    search.Searcher(keyword)
}

// Copyright (c) 2026 Zeronetsec