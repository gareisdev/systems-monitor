package main

import (
	"github.com/gareisdev/systems-monitor/internal/config"
	"github.com/gareisdev/systems-monitor/internal/logger"
	"github.com/gareisdev/systems-monitor/internal/monitor"
)

func main() {
	// Init
	logger.Info("Init system")

	configFile := "urls.yml"
	yamlConfig := config.ReadConfig(configFile)

	monitorConfig := config.ConvertConfig(yamlConfig)

	go monitor.StartMonitoring(monitorConfig)

	select {}
}
