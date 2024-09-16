package app

import (
	"eyemonit/internal/config"
	"eyemonit/internal/service"
	"fmt"
)

func Run() {
	//fmt.Println(config.NewConfig())
	metricsconf, err := config.NewMetricsConfig()
	if err != nil {
		fmt.Println("error", err)
	}

	if metricsconf.Metrics.Cpu {
		//service.GetCpu()
	}

	if metricsconf.Metrics.Disk {
		//service.GetDisk()
	}

	if metricsconf.Metrics.Memory {
		//service.GetMemory()
	}

	if metricsconf.Metrics.Network {
		fmt.Println(service.GetNetwork())
		fmt.Println(service.GetBandwidth())
	}
}
