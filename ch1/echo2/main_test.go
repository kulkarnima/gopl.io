package main

import(
    "testing"
)

func BenchmarkEcho2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Echo2([]string{"gopher", "go", "google"});
    }
}
