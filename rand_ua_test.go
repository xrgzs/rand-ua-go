package rand_ua

import (
	"testing"
)

func TestMustGetRandomUA(t *testing.T) {
	ua := MustGetRandomUA()
	t.Logf("Random UA: %s", ua)
	if ua == "" {
		t.Error("UserAgent should not be empty")
	}
}

func TestGetRandomUAFromData_ValidSingleLine(t *testing.T) {
	data := []byte(`{"useragent":"TestAgent/1.0","percent":1.0,"type":"desktop","device_brand":"TestBrand","browser":"TestBrowser","browser_version":"1.0","browser_version_major_minor":1.0,"os":"TestOS","os_version":"1.0","platform":"TestPlatform"}`)
	ua := getRandomUAFromData(data)
	if ua != "TestAgent/1.0" {
		t.Errorf("expected TestAgent/1.0, got: %s", ua)
	}
}

func TestGetRandomUAFromData_ValidMultipleLines(t *testing.T) {
	data := []byte(
		`{"useragent":"AgentA/1.0","percent":1.0,"type":"desktop","device_brand":"A","browser":"A","browser_version":"1.0","browser_version_major_minor":1.0,"os":"A","os_version":"1.0","platform":"A"}
{"useragent":"AgentB/2.0","percent":2.0,"type":"mobile","device_brand":"B","browser":"B","browser_version":"2.0","browser_version_major_minor":2.0,"os":"B","os_version":"2.0","platform":"B"}`,
	)
	// Run multiple times to check both lines can be picked
	foundA, foundB := false, false
	for i := 0; i < 20; i++ {
		ua := getRandomUAFromData(data)
		if ua == "AgentA/1.0" {
			foundA = true
		}
		if ua == "AgentB/2.0" {
			foundB = true
		}
	}
	if !foundA || !foundB {
		t.Errorf("expected both AgentA/1.0 and AgentB/2.0 to be possible, got foundA=%v, foundB=%v", foundA, foundB)
	}
}

func TestGetRandomUAFromData_InvalidJSONLine(t *testing.T) {
	data := []byte(`{"useragent":"AgentA/1.0"}
not a json`)
	ua := getRandomUAFromData(data)
	// Could be valid or fallback depending on which line is picked
	if ua != "AgentA/1.0" && ua != "curl/8.9.1" {
		t.Errorf("expected AgentA/1.0 or fallback, got: %s", ua)
	}
}

func TestGetRandomUAFromData_EmptyLines(t *testing.T) {
	data := []byte("\n\n")
	ua := getRandomUAFromData(data)
	if ua != "curl/8.9.1" {
		t.Errorf("expected fallback ua for only empty lines, got: %s", ua)
	}
}

func TestGetRandomUAFromData_ZeroLines(t *testing.T) {
	data := []byte("")
	ua := getRandomUAFromData(data)
	if ua != "curl/8.9.1" {
		t.Errorf("expected fallback ua for only empty lines, got: %s", ua)
	}
}
