package accept

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

var byTLD map[string]string

func init() {
	// Open our jsonFile
	jsonFile, err := os.Open("main.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &byTLD)
	fmt.Println(byTLD["ua"])
}

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
		result = byTLD["ua"]
	}

	return result
}
