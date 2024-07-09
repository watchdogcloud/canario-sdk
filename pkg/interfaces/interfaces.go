package interfaces

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type CPUVitals interface {
	FetchCPUPercent() ([]float64, error)
	FetchPhysicalOrLogicalCores() (int, error)
	FetchCPUInfo() ([]cpu.InfoStat, error) // detailed
}

type DiskVitals interface {
	FetchDiskUsage(path string) (*disk.UsageStat, error)
	FetchDiskPartitions(all bool) ([]disk.PartitionStat, error)
	FetchDiskIOCounters() (map[string]disk.IOCountersStat, error)
}

type NetworkVitals interface {
	FetchNetworkIO() ([]net.IOCountersStat, error)
	FetchNetworkInterfaces() ([]net.InterfaceStat, error)
}

type MemoryVitals interface {
	FetchMemoryStats() (*mem.VirtualMemoryStat, error)
}

type CanarioDisk struct{}
type CanarioCPU struct{}
type CanarioNetwork struct{}
type CanarioMemory struct{}

// returns array of cpu percentages
func (c *CanarioCPU) FetchCPUPercent() ([]float64, error) {
	return cpu.Percent(1*time.Second, false)
}

func (c *CanarioCPU) FetchPhysicalOrLogicalCores() (int, error) {
	needsLogical := true
	return cpu.Counts(needsLogical)
}

// FetchCPUInfo() ([]cpu.InfoStat, error) // detailed
func (c *CanarioCPU) FetchCPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

func (c *CanarioDisk) FetchDiskUsage(path string) (*disk.UsageStat, error) {
	return disk.Usage(path)
}

func (c *CanarioDisk) FetchDiskPartitions(all bool) ([]disk.PartitionStat, error) {
	return disk.Partitions(all)
}

func (c *CanarioDisk) FetchDiskIOCounters() (map[string]disk.IOCountersStat, error) {
	return disk.IOCounters()
}

func (c *CanarioNetwork) FetchNetworkIO() ([]net.IOCountersStat, error) {
	return net.IOCounters(true)
}

func (c *CanarioNetwork) FetchNetworkInterfaces() ([]net.InterfaceStat, error) {
	return net.Interfaces()
}

func (c *CanarioMemory) FetchMemoryStats() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}
