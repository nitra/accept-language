package accept

import (
	"flag"
	"testing"

	log "github.com/sirupsen/logrus"
)

var domain string

func init() {
	log.SetLevel(log.DebugLevel)

	flag.StringVar(&domain, "domain", "nitra.ua", "Domain for dns check")
}

func TestLanguage(t *testing.T) {

	log.Debugf("Test start for: %s", domain)

	lang := Language(domain)
	if lang != byTLD["ua"] {
		t.Errorf("Lang incorrect, got: %s, want: %s.", lang, byTLD["ua"])
	}

	lang = Language("https://nitra.ua")
	if lang != byTLD["ua"] {
		t.Errorf("Lang incorrect, got: %s, want: %s.", lang, byTLD["ua"])
	}
}
