package config

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/gareisdev/systems-monitor/internal/monitor"

	"gopkg.in/yaml.v2"
)

type URLConfig struct {
	URL            string   `yaml:"url"`
	Timeout        int      `yaml:"timeout"`
	ExpectedStatus []string `yaml:"expected_status_codes"`
}

type Config struct {
	URLs []URLConfig `yaml:"urls"`
}

func ReadConfig(filename string) Config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error al leer el archivo YAML: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error al decodificar el archivo YAML: %v", err)
	}

	return config
}

func ConvertConfig(config Config) []monitor.URLConfig {
	var monitorConfig []monitor.URLConfig

	for _, url := range config.URLs {
		var expectedStatusCodes []int

		for _, statusCode := range url.ExpectedStatus {
			if strings.Contains(statusCode, "-") {
				parts := strings.Split(statusCode, "-")
				start, _ := strconv.Atoi(parts[0])
				end, _ := strconv.Atoi(parts[1])
				for i := start; i <= end; i++ {
					expectedStatusCodes = append(expectedStatusCodes, i)
				}
			} else {
				code, _ := strconv.Atoi(statusCode)
				expectedStatusCodes = append(expectedStatusCodes, code)
			}
		}

		monitorConfig = append(monitorConfig, monitor.URLConfig{
			URL:            url.URL,
			Timeout:        url.Timeout,
			ExpectedStatus: expectedStatusCodes,
		})
	}

	return monitorConfig
}
