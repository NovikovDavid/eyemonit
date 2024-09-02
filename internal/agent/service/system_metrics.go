package service

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetCpu() ([]cpu.TimesStat, []cpu.InfoStat, error) {
	// InfoStat - структура, которая содержит информацию о процессоре
	infostat, err := cpu.Info()
	if err != nil {
		return nil, nil, err
	}

	// TimesStat - структура, которая содержит информацию о времени использования процессора
	infotimes, err := cpu.Times(false) // true - подробный список ядер и их статистика
	if err != nil {
		return nil, nil, err
	}

	return infotimes, infostat, nil
}

func GetDisk() (map[string]disk.IOCountersStat, *disk.UsageStat, []disk.PartitionStat, error) {
	usagestat, err := disk.Usage("/") // используется для представления информации о использовании
	if err != nil {                   // дискового пространства в файловой системе
		return nil, nil, nil, err
	}

	// используется для хранения информации о разделе диска или точке монтирования
	partitionstate, err := disk.Partitions(false) // false - исключает виртуальные или псевдо
	if err != nil {                               // разделы(tmpfs, proc), true - вернет все
		return nil, nil, nil, err
	}

	iocounters, err := disk.IOCounters() // используется для хранения статистики ввода-вывода(I/O) дисков
	if err != nil {
		return nil, nil, nil, err
	}

	return iocounters, usagestat, partitionstate, nil
}

func GetMemory() (*mem.VirtualMemoryStat, *mem.SwapMemoryStat, error) {
	virtualmemorystat, err := mem.VirtualMemory() //оперативная память
	if err != nil {
		return nil, nil, err
	}

	swapmemorystat, err := mem.SwapMemory() //swap память
	if err != nil {
		return nil, nil, err
	}

	/*
		swapdevice, err := mem.SwapDevices()
		fmt.Println(swapdevice)					// not implemented yet (еще не реализовано)
		if err != nil {
			return nil, nil, err
		}
	*/

	return virtualmemorystat, swapmemorystat, nil
}
