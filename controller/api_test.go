package controller

import (
	"bytes"
	"encoding/json"
	"github.com/reef-pi/reef-pi/controller/utils"
	"testing"
)

func TestAPI(t *testing.T) {
	r, err := New("0.1", "api-test.db")
	if err != nil {
		t.Fatal("Failed to create new reef-pi controller. Error:", err)
	}
	r.settings.Capabilities.DevMode = true
	if err := r.Start(); err != nil {
		t.Fatal("Failed to load subsystem. Error:", err)
	}
	tr := utils.NewTestRouter()

	r.loadAPI(tr.Router)
	if err := tr.Do("GET", "/api/health_stats/hour", new(bytes.Buffer), nil); err != nil {
		t.Fatal("Failed to get per minute health data.Error:", err)
	}
	if err := tr.Do("GET", "/api/health_stats/week", new(bytes.Buffer), nil); err != nil {
		t.Fatal("Failed to get per minute health data. Error:", err)
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&DefaultCredentials)
	if err := tr.Do("POST", "/api/credentials", body, nil); err != nil {
		t.Fatal("Failed to update creds via api")
	}
	if err := tr.Do("GET", "/api/settings", new(bytes.Buffer), nil); err != nil {
		t.Fatal("Failed to get settings via api")
	}
	body.Reset()
	json.NewEncoder(body).Encode(&DefaultSettings)
	if err := tr.Do("POST", "/api/settings", body, nil); err != nil {
		t.Fatal("Failed to update settings via api")
	}
	if err := tr.Do("GET", "/api/settings", new(bytes.Buffer), nil); err != nil {
		t.Fatal("Failed to get settings via api")
	}
	body.Reset()
	json.NewEncoder(body).Encode(&utils.TelemetryConfig{})
	if err := tr.Do("POST", "/api/telemetry", body, nil); err != nil {
		t.Fatal("Failed to update telemetry via api")
	}
	if err := tr.Do("GET", "/api/telemetry", new(bytes.Buffer), nil); err != nil {
		t.Fatal("Failed to get telemetry via api")
	}
	if err := r.Stop(); err != nil {
		t.Fatal(err)
	}
}
