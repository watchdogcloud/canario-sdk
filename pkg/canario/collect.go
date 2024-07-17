package canario

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/watchdogcloud/canario/internal/conf"
	"github.com/watchdogcloud/canario/internal/conf/parse"
	"github.com/watchdogcloud/canario/pkg/client"
	"github.com/watchdogcloud/canario/pkg/interfaces"
)

const (
	BATCH_SIZE = 4
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
	CPU     interface{} `json:"cpu"`
	Disk    interface{} `json:"disk"`
	Mem     interface{} `json:"mem"`
	Network interface{} `json:"network"`
}

type MetricBatch struct {
	BufferMutex sync.Mutex
	Buffer      []MetricBufferUnitStructure
	Key         string
}

// constructer
func CreateNewMetricBatch(apiKey string) MetricBatch {
	return MetricBatch{
		Buffer: make([]MetricBufferUnitStructure, 0),
		Key:    apiKey,
	}
}

func (mb *MetricBatch) AddMetricToBatch(incomingMetric MetricBufferUnitStructure) {
	mb.BufferMutex.Lock()
	defer mb.BufferMutex.Unlock()

	mb.Buffer = append(mb.Buffer, incomingMetric)
	if len(mb.Buffer) >= BATCH_SIZE {
		go mb.apiCall()
	}
}

func (mb *MetricBatch) apiCall() {
	mb.BufferMutex.Lock()
	defer mb.BufferMutex.Unlock()

	client := client.CreateNewClient("API-Key", mb.Key)
	extraHeaders := make(map[string]string)

	payload := make(map[string]interface{})
	for _, m := range mb.Buffer {
		payload["cpu"] = m.CPU
		payload["network"] = m.Network
		payload["mem"] = m.Mem
		payload["disk"] = m.Disk
	}

	_, err := client.RecvData.PushMetricsToServer(payload, extraHeaders)
	if err != nil {
		log.Println("Error pushing metrics to server:", err)
		return
	}
	mb.Buffer = mb.Buffer[:0]
}

func (c *Canario) RunPeriodicMetrics() {

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

		c.mb.AddMetricToBatch(thisMetric)
	}
}

func NewCanario() *Canario {
	cfg := conf.CreateNewConf()
	return &Canario{
		cfg: &cfg,
		mb:  CreateNewMetricBatch(cfg.API.Key),
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
