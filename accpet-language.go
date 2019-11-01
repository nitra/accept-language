package accept

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var byTLD map[string]string

func init() {
	// Open our jsonFile
	// jsonFile, err := os.Open("main.json")

	// if we os.Open returns an error then handle it
	resp, err := http.Get("https://raw.githubusercontent.com/nitra/accept-language/master/main.json")
	if err != nil {
		log.Panic(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(byteValue), &byTLD)
	fmt.Println(byTLD["ua"])
}

// Language preferred language value by domain TLD
func Language(domain string) string {

	if len(domain) < 3 {
		log.Print("very short domain")
		return ""
	}

	// tests a string to determine if it is a url or not.
	u, err := url.ParseRequestURI(domain)
	if err == nil {
		domain = u.Host
	}

	result := ""
	countryTLD := domain[len(domain)-3:]
	if countryTLD == ".ua" {
		result = byTLD["ua"]
	}

	return result
}
