package accept

import (
	"testing"
)

func TestLanguage(t *testing.T) {
	lang := Language("nitra.ua")
	if lang != byTLD["ua"] {
		t.Errorf("Lang incorrect, got: %s, want: %s.", lang, byTLD["ua"])
	}

	lang = Language("https://nitra.ua")
	if lang != byTLD["ua"] {
		t.Errorf("Lang incorrect, got: %s, want: %s.", lang, byTLD["ua"])
	}
}
