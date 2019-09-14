package accept

import (
	"net/url"
)

// Language preferred language value by domain TLD
func Language(domain string) string {

	// tests a string to determine if it is a url or not.
	u, err := url.ParseRequestURI(domain)
	if err == nil {
		domain = u.Host
	}

	result := ""
	countryTLD := domain[len(domain)-3:]
	if countryTLD == ".ua" {
		result = "uk-UA, uk;q=0.9, ru;q=0.8, *;q=0.5"
	}

	return result
}
