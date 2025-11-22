package web

import "testing"

func TestGreeting(t *testing.T) {
    if got := Greeting("Go"); got != "Hello, Go!" {
        t.Fatalf("unexpected greeting: %s", got)
    }
    if got := Greeting(""); got != "Hello, World!" {
        t.Fatalf("unexpected greeting with fallback: %s", got)
    }
}

func TestTitles(t *testing.T) {
    input := []string{"go", "lang"}
    want := []string{"Go", "Lang"}
    got := Titles(input)
    for i := range want {
        if got[i] != want[i] {
            t.Fatalf("unexpected title at %d: %s", i, got[i])
        }
    }
}
