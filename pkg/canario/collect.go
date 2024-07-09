package canario

import (
	"log"
	"time"

	"github.com/zakhaev26/canario/internal/conf"
	"github.com/zakhaev26/canario/internal/conf/parse"
	types "github.com/zakhaev26/canario/pkg/interfaces"
)

type Watchdog interface {
	NewCanario()
	RunPeriodicMetrics()
}

type Canario struct {
	cfg *parse.Config
}

type MetricBufferUnitStructure map[string]interface{}{}

func (c *Canario) RunPeriodicMetrics() {
	ticker := time.NewTicker(time.Duration(c.cfg.Monitoring.IntervalSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			if c.cfg.Metrics.CPU.Enabled {
				MetricBufferUnitStructure[""] = x
			}
			if c.cfg.Metrics.Memory.Enabled {
			}
			// if cfg.Metrics.Disk.Enabled {
			// 	fetchAndLogDisk()
			// }
			// if cfg.Metrics.Network.Enabled {
			// 	fetchAndLogNetwork()
			// }
		}
	}
}

func NewCanario() *Canario {
	cfg := conf.CreateNewConf()
	return &Canario{
		cfg: &cfg,
	}
}

// func fetchAndLogCPU() ([]float64, error){
// 	cpu := types.CanarioCPU{}
// 	cpuPercent, err := cpu.FetchCPUPercent()
// 	if err != nil {
// 		// Handle error, e.g., log it or send alert
// 		log.Printf("Error fetching CPU metrics: %v", err)
// 		return
// 	}
// 	return cpuPercent, err
// }
