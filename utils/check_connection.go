// https://github.com/Zeronetsec/Znskit

package utils

import (
    "time"
    "net"
    "net/http"
)

func CheckConnection() bool {
    client := &http.Client{
        Timeout: 30 * time.Second,
        Transport: &http.Transport{
            DialContext: (&net.Dialer{
                Timeout: 5 * time.Second,
                KeepAlive: 30 * time.Second,
            }).DialContext,
            TLSHandshakeTimeout: 5 * time.Second,
            MaxIdleConns: 100,
            IdleConnTimeout: 90 * time.Second,
        },
    }

    _, err := client.Get("https://github.com/Zeronetsec")
    return err == nil
}

// Copyright (c) 2026 Zeronetsec