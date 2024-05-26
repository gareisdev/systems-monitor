package monitor

import (
	"net/http"
	"time"

	"github.com/gareisdev/systems-monitor/internal/logger"
)

func StartMonitoring(configs []URLConfig) {
	for _, config := range configs {
		go monitorURL(config)
	}
}

func monitorURL(config URLConfig) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			checkURL(config)
		}
	}
}

func checkURL(config URLConfig) {
	client := http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
	}

	resp, err := client.Get(config.URL)
	if err != nil {
		logger.Error("GET ", config.URL, ": ", err)
		return
	}
	defer resp.Body.Close()

	found := false
	for _, expected := range config.ExpectedStatus {
		if resp.StatusCode == expected {
			found = true
			break
		}
	}

	if found {
		logger.Info("URL ", config.URL, " OK. Status Code: ", resp.StatusCode)
	} else {
		logger.Warn("URL ", config.URL, " BAD. Status Code: ", resp.StatusCode)
	}
}
