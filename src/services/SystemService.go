package service

import (
    "fmt"
    "github.com/shirou/gopsutil/v4/mem"
    "github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/net"
)
func GetCPUInfo() (map[string]interface{}, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
    if err != nil {
        return nil, err
    }
	result := map[string]interface{}{
        "usage_percent": cpuPercent[0], // Uso total como porcentaje
        "times": map[string]float64{
            "user":   times[0].User,
            "system": times[0].System,
            "idle":   times[0].Idle,
            "nice":   times[0].Nice,
            "iowait": times[0].Iowait,
        },
    }
	return result, nil
}
func GetMemoryInfo() (map[string]interface{}, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	memoryInfo := map[string]interface{}{
		"total":     memInfo.Total,
		"available": memInfo.Available,
		"used":      memInfo.Used,
		"free":      memInfo.Free,
	}
	return memoryInfo, nil
}
func GetNetworkInfo() (map[string]interface{}, error) {
    ioStats, err := net.IOCounters(true)
    if err != nil {
        return nil, err
    }

    conns, err := net.Connections("all")
    if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "interfaces": ioStats,
        "connections_count": len(conns),
    }, nil
}