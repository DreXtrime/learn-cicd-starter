package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey secret-token-123")

	got, err := GetAPIKey(headers)

	want := "secret-token-123"

	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}