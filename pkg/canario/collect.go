package canario

import (
	"fmt"
	"sync"
	"time"

	"github.com/zakhaev26/canario/internal/conf"
	"github.com/zakhaev26/canario/internal/conf/parse"
	"github.com/zakhaev26/canario/pkg/interfaces"
)

const (
	BATCH_SIZE = 128
)

type Watchdog interface {
	NewCanario()
	RunPeriodicMetrics()
}
type Canario struct {
	cfg *parse.Config
	mb  MetricBatch // NOT FOR USE IN CLIENT!!
}

type MetricBufferUnitStructure struct {
	CPU     interface{}
	Disk    interface{}
	Mem     interface{}
	Network interface{}
}

type MetricBatch struct {
	BufferMutex sync.Mutex
	Buffer      []MetricBufferUnitStructure
}

// constructer
func CreateNewMetricBatch() MetricBatch {
	return MetricBatch{
		Buffer: make([]MetricBufferUnitStructure, 0),
	}
}

func (mb *MetricBatch) addMetricToBatch(incomingMetric MetricBufferUnitStructure) {
	mb.BufferMutex.Lock()
	defer mb.BufferMutex.Unlock()

	mb.Buffer = append(mb.Buffer, incomingMetric)

	if len(mb.Buffer) >= BATCH_SIZE {
		// make an api call with expo backoff..
		mb.Buffer = nil
	}
}

func (c *Canario) RunPeriodicMetrics() {
	fmt.Println(c.cfg.Monitoring.IntervalSeconds)
	ticker := time.NewTicker(time.Duration(c.cfg.Monitoring.IntervalSeconds) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		thisMetric := MetricBufferUnitStructure{}
		if c.cfg.Metrics.CPU.Enabled {
			cpuMetrics, err := fetchAndLogCPU()
			if err == nil {
				thisMetric.CPU = cpuMetrics
			}
		}

		if c.cfg.Metrics.Memory.Enabled {
			memMetrics, err := fetchAndLogMemory()
			if err == nil {
				thisMetric.Mem = memMetrics
			}
		}

		if c.cfg.Metrics.Disk.Enabled {
			diskMetrics, err := fetchAndLogDisk()
			if err == nil {
				thisMetric.Disk = diskMetrics
			}
		}

		if c.cfg.Metrics.Network.Enabled {
			netMetrics, err := fetchAndLogNetwork()
			if err == nil {
				thisMetric.Network = netMetrics
			}
		}

		c.mb.addMetricToBatch(thisMetric)
		fmt.Println(c.mb.Buffer)
	}
}

func NewCanario() *Canario {
	cfg := conf.CreateNewConf()
	return &Canario{
		cfg: &cfg,
		mb:  CreateNewMetricBatch(),
	}
}

func fetchAndLogCPU() (interface{}, error) {
	cpu := interfaces.CanarioCPU{}
	cpuPercent, err := cpu.FetchCPUPercent()
	if err != nil {
		fmt.Printf("Error fetching CPU metrics: %v\n", err)
		return nil, err
	}
	return cpuPercent, nil
}

func fetchAndLogMemory() (interface{}, error) {
	mem := interfaces.CanarioMemory{}
	memStats, err := mem.FetchMemoryStats()
	if err != nil {
		fmt.Printf("Error fetching memory metrics: %v\n", err)
		return nil, err
	}
	return memStats, nil
}

func fetchAndLogDisk() (interface{}, error) {
	disk := interfaces.CanarioDisk{}
	diskUsage, err := disk.FetchDiskUsage("/")
	if err != nil {
		fmt.Printf("Error fetching disk metrics: %v\n", err)
		return nil, err
	}
	return diskUsage, nil
}

func fetchAndLogNetwork() (interface{}, error) {
	net := interfaces.CanarioNetwork{}
	netStats, err := net.FetchNetworkIO()
	if err != nil {
		fmt.Printf("Error fetching network metrics: %v\n", err)
		return nil, err
	}
	return netStats, nil
}
