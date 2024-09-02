package app

import (
	"eyemonit/internal/agent/config"
	"eyemonit/internal/agent/service"
	"fmt"
)

func Run() {
	//fmt.Println(config.NewConfig())
	metricsconf, err := config.NewMetricsConfig()
	if err != nil {
		fmt.Println("error", err)
	}

	if metricsconf.Metrics.Cpu {
		//fmt.Println(service.GetCpu())
	}

	if metricsconf.Metrics.Disk {
		/*_, usagestat, _, err := service.GetDisk()
		if err != nil {
			fmt.Println("error", err)
		}
		fmt.Println(usagestat)*/
	}

	if metricsconf.Metrics.Memory {
		//fmt.Println(service.GetMemory())
	}

	if metricsconf.Metrics.Network {
		service.GetNetwork()
	}
}
