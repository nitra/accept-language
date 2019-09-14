package accept

import (
	"testing"
)

func TestLanguage(t *testing.T) {
	expected := "uk-UA, uk;q=0.9, ru;q=0.8, *;q=0.5"
	lang := Language("nitra.ua")
	if lang != expected {
		t.Errorf("Lang incorrect, got: %s, want: %s.", lang, expected)
	}
}
