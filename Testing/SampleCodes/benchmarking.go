package main

import (
    "net/http"
    "testing"
)

func BenchmarkAPICall(b *testing.B) {
    client := &http.Client{}
    url := "https://api.example.com/endpoint"

    for i := 0; i < b.N; i++ {
        resp, err := client.Get(url)
        if err != nil {
            b.Fatal(err)
        }
        resp.Body.Close()
    }
}