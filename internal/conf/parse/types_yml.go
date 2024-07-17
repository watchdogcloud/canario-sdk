package parse

import (
	"sync"
)

type Config struct {
	Version string `yaml:"version"`

	API struct {
		Key     string `yaml:"key"`
		BaseURI string `yaml:"baseuri"`
	} `yaml:"api"`

	Metrics struct {
		CPU     MetricConfig `yaml:"cpu"`
		Memory  MetricConfig `yaml:"memory"`
		Disk    MetricConfig `yaml:"disk"`
		Network MetricConfig `yaml:"network"`
	} `yaml:"metrics"`

	Monitoring struct {
		IntervalSeconds int `yaml:"interval_seconds"`
		RetentionHours  int `yaml:"retention_hours"`
	} `yaml:"monitoring"`

	Alerting struct {
		Enabled              bool             `yaml:"enabled"`
		Thresholds           ThresholdsConfig `yaml:"thresholds"`
		NotificationChannels struct {
			Emails []string `yaml:"emails"`
		} `yaml:"notification_channels"`
	} `yaml:"alerting"`

	Mu sync.Mutex
}

type MetricConfig struct {
	Enabled bool `yaml:"enabled"`
}

type ThresholdsConfig struct {
	CPUUsagePercentage      ThresholdValues `yaml:"cpu_usage_percentage"`
	MemoryUsagePercentage   ThresholdValues `yaml:"memory_usage_percentage"`
	DiskUsagePercentage     ThresholdValues `yaml:"disk_usage_percentage"`
	NetworkTrafficThreshold ThresholdValues `yaml:"network_traffic_threshold"`
}

type ThresholdValues struct {
	Warning  int `yaml:"warning"`
	Critical int `yaml:"critical"`
}

func (c *Config) SetDefaultsIfFieldsMissing() {
	if !c.Metrics.CPU.Enabled {
		c.Metrics.CPU.Enabled = true
	}
	if !c.Metrics.Memory.Enabled {
		c.Metrics.Memory.Enabled = true
	}
	if !c.Metrics.Disk.Enabled {
		c.Metrics.Disk.Enabled = true
	}
	if !c.Metrics.Network.Enabled {
		c.Metrics.Network.Enabled = true
	}

	if c.Monitoring.IntervalSeconds == 0 {
		c.Monitoring.IntervalSeconds = 60
	}
	if c.Monitoring.RetentionHours == 0 {
		c.Monitoring.RetentionHours = 24
	}

	if !c.Alerting.Enabled {
		c.Alerting.Enabled = true
	}

	if c.Alerting.Thresholds.CPUUsagePercentage.Warning == 0 {
		c.Alerting.Thresholds.CPUUsagePercentage.Warning = 80
	}
	if c.Alerting.Thresholds.CPUUsagePercentage.Critical == 0 {
		c.Alerting.Thresholds.CPUUsagePercentage.Critical = 90
	}

	if c.Alerting.Thresholds.MemoryUsagePercentage.Warning == 0 {
		c.Alerting.Thresholds.MemoryUsagePercentage.Warning = 70
	}
	if c.Alerting.Thresholds.MemoryUsagePercentage.Critical == 0 {
		c.Alerting.Thresholds.MemoryUsagePercentage.Critical = 80
	}

	if c.Alerting.Thresholds.DiskUsagePercentage.Warning == 0 {
		c.Alerting.Thresholds.DiskUsagePercentage.Warning = 75
	}
	if c.Alerting.Thresholds.DiskUsagePercentage.Critical == 0 {
		c.Alerting.Thresholds.DiskUsagePercentage.Critical = 85
	}

	if c.Alerting.Thresholds.NetworkTrafficThreshold.Warning == 0 {
		c.Alerting.Thresholds.NetworkTrafficThreshold.Warning = 1000
	}
	if c.Alerting.Thresholds.NetworkTrafficThreshold.Critical == 0 {
		c.Alerting.Thresholds.NetworkTrafficThreshold.Critical = 1500
	}
}
