package web

// Greeting returns a greeting message for the provided name.
func Greeting(name string) string {
    if name == "" {
        name = "World"
    }
    return "Hello, " + name + "!"
}

// Titles converts a list of words into title case strings.
func Titles(words []string) []string {
    titles := make([]string, len(words))
    for i, w := range words {
        if w == "" {
            titles[i] = ""
            continue
        }
        titles[i] = string(w[0]-32) + w[1:]
    }
    return titles
}

// GreetingWithEmoji returns greeting text appended with an emoji.
func GreetingWithEmoji(name, emoji string) string {
    if emoji == "" {
        emoji = "🙂"
    }
    return Greeting(name) + " " + emoji
}
