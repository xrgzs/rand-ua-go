package rand_ua

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/json"
	"math/rand"
)

//go:embed fake_data/src/fake_useragent/data/browsers.jsonl
var browsersJsonl []byte

// Browser represents a user agent with its details.
type Browser struct {
	UserAgent                string  `json:"useragent"`
	Percent                  float64 `json:"percent"`
	Type                     string  `json:"type"`
	DeviceBrand              string  `json:"device_brand"`
	Browser                  string  `json:"browser"`
	BrowserVersion           string  `json:"browser_version"`
	BrowserVersionMajorMinor float64 `json:"browser_version_major_minor"`
	Os                       string  `json:"os"`
	OsVersion                string  `json:"os_version"`
	Platform                 string  `json:"platform"`
}

// MustGetRandomUA returns a random User-Agent string from the embedded browsers data.
func MustGetRandomUA() string {
	return getRandomUAFromData(browsersJsonl)
}

// getRandomUAFromData returns a random User-Agent string from the provided JSONL data.
func getRandomUAFromData(data []byte) string {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		return "curl/8.9.1"
	}
	target := rand.Intn(len(lines))
	line := lines[target]

	var obj Browser
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return "curl/8.9.1"
	}
	return obj.UserAgent
}
