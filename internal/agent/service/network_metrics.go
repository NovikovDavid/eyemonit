package service

import (
	"time"

	pkg "eyemonit/pkg/agent/metrics"

	"github.com/shirou/gopsutil/v4/net"
)

func GetNetwork() ([]net.IOCountersStat, []net.ConnectionStat, net.InterfaceStatList, error) {
	iocountersstat, err := net.IOCounters(false) // true - означает, что будут возвращены данные по каждому интерфейсу в отдельности.
	if err != nil {                              // структура предоставляет данные о сетевых интерфейсах,
		return nil, nil, nil, err // такие как количество отправленных и полученных байт, пакетов и ошибки.
	}

	connectionstat, err := net.Connections("") // "" - означает все соединения
	if err != nil {                            // структура описывает текущее состояние сетевых соединений на системе.
		return nil, nil, nil, err
	}
	/*
		protocountersstat, err := net.ProtoCounters([]string{"tcp", "udp"}) // or nil
		if err != nil {											            // not implemented yet (еще не реализовано)
			return nil, nil, nil, nil, err
		}
	*/
	interfacestat, err := net.Interfaces() // информация о интерфейсах
	if err != nil {
		return nil, nil, nil, err
	}

	return iocountersstat, connectionstat, interfacestat, nil
}

func GetBandwidth() ([]pkg.Bandwidth, error) {
	var (
		interval      = 2 * time.Second
		bandwidthData []pkg.Bandwidth
	)

	initialCounters, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	time.Sleep(interval)

	finalCounters, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	for i := range initialCounters {
		bytesSentDiff := finalCounters[i].BytesSent - initialCounters[i].BytesSent
		bytesRecvDiff := finalCounters[i].BytesRecv - initialCounters[i].BytesRecv

		bandwidth := pkg.Bandwidth{
			Interface:     initialCounters[i].Name,
			UploadSpeed:   float64(bytesSentDiff) / interval.Seconds(),
			DownloadSpeed: float64(bytesRecvDiff) / interval.Seconds(),
		}
		bandwidthData = append(bandwidthData, bandwidth)
	}
	return bandwidthData, nil
}
