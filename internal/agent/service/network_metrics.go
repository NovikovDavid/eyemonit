package service

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/net"
)

func GetNetwork() ([]net.IOCountersStat, []net.ConnectionStat, net.InterfaceStatList, error) {
	iocountersstat, err := net.IOCounters(true)
	if err != nil {
		return nil, nil, nil, err
	}

	connectionstat, err := net.Connections("")
	if err != nil {
		return nil, nil, nil, err
	}
	/*
		protocountersstat, err := net.ProtoCounters([]string{"tcp", "udp"}) // or nil
		if err != nil {											            // not implemented yet (еще не реализовано)
			return nil, nil, nil, nil, err
		}
	*/
	interfacestat, err := net.Interfaces()
	fmt.Println(interfacestat)
	if err != nil {
		return nil, nil, nil, err
	}

	return iocountersstat, connectionstat, interfacestat, nil
}
