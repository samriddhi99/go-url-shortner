package shortener

import (
    "encoding/json"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

var (
    store = make(map[string]string)
    mu    sync.Mutex
)

func generateKey(n int) string {
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    rand.Seed(time.Now().UnixNano())
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var data map[string]string
    json.NewDecoder(r.Body).Decode(&data)
    original := data["url"]
    key := generateKey(6)

    mu.Lock()
    store[key] = original
    mu.Unlock()

    resp := map[string]string{"short_url": "http://localhost:8080/" + key}
    json.NewEncoder(w).Encode(resp)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Path[1:]

    mu.Lock()
    original, ok := store[key]
    mu.Unlock()

    if !ok {
        http.NotFound(w, r)
        return
    }

    http.Redirect(w, r, original, http.StatusSeeOther)
}
