package rand_ua

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/json"
	"math/rand"
)

//go:embed user-agents.jsonl
var browsersJsonl []byte

// Browser represents a user agent with its details.
type Browser struct {
	AppName    string `json:"appName"`
	Connection struct {
		Downlink      int    `json:"downlink"`
		EffectiveType string `json:"effectiveType"`
		Rtt           int    `json:"rtt"`
	} `json:"connection"`
	Language       string  `json:"language"`
	Platform       string  `json:"platform"`
	PluginsLength  int     `json:"pluginsLength"`
	ScreenHeight   int     `json:"screenHeight"`
	ScreenWidth    int     `json:"screenWidth"`
	UserAgent      string  `json:"userAgent"`
	Vendor         string  `json:"vendor"`
	ViewportHeight int     `json:"viewportHeight"`
	ViewportWidth  int     `json:"viewportWidth"`
	Weight         float64 `json:"weight"`
	DeviceCategory string  `json:"deviceCategory"`
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
