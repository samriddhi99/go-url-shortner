package shortener

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key := generateKey(6)
	if len(key) != 6 {
		t.Errorf("expected key length 6, got %d", len(key))
	}
}

func TestNormalizeURL(t *testing.T) {
	input := "example.com"
	expected := "http://example.com"
	result := NormalizeURL(input)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestShortenHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "/shorten", strings.NewReader(`{"url":"http://test.com"}`))
	w := httptest.NewRecorder()
	ShortenHandler(w, req)
	if w.Code != 200 {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
