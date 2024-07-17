# Canario SDK

<div align='center'> 
<img src=".github/logo.png" alt="canario.png" width='150'></img>
</div>

The Canario SDK is a Go SDK designed for easy metric collection and export. It is a part of `watchdogcloud` and is heavily dependent on a configuration and management server - Terrier (watch <a href="https://github.com/watchdogcloud/terrier">this</a>) that helps watchdog to send critical alerts to stakeholders when the hardware goes wrong.

Internally, it uses `github.com/shirou/gopsutil` to expose the OS and hardware level details with metrics like disk, CPU, and memory.

In batches, these metrics are then pushed to the Terrier management server that processes these metrics, to be precise - stores them in persistent storage (MongoDB), streams realtime metrics to client apps, and constantly computes over these datapoints to calculate if there is a Spike trigger or a Cumulative trigger.

### Spike Triggers
These are triggers that detect sudden and significant changes in the metric values, indicating potential issues that need immediate attention.

### Cumulative Triggers
These triggers detect gradual changes in metric values over time, indicating potential issues that may arise if the trend continues.

Current work involves making a nice console for watchdog so that the dirty configuration changes need not be done via touching the code. It's in progress and will hopefully be out soon.

```bash
go get github.com/zakhaev26/canario
```

## Configuration File (`canario.yml`)
```yaml
# canario.yml
version: 0.8.1

api:
  baseuri: 'http://localhost:3030' 
  key: '1bf42e28aa0ac838e47aabac10a4439567b289460bb784e57fc8304baf9ff095'

metrics:
  cpu:
    enabled: true
  memory:
    enabled: true
  disk:
    enabled: true
  network:
    enabled: true

monitoring:
  interval_seconds: 1
  retention_hours: 24
```

## Configuration Attributes

|Attribute                     | Description  |
|------------------------------|--------------|
| `version` |The version of the Canario SDK. Ensure this matches with your installed version. |
| `api.baseuri` | The base URI of the terrier management server. |
| `api.key` | The API key used for authenticating requests to the terrier management server. |
| `metrics.cpu.enabled` | Enable or disable CPU metrics collection. |
| `metrics.memory.enabled` | Enable or disable Memory metrics collection. |
| `metrics.disk.enabled` | Enable or disable Disk metrics collection. |
|`metrics.network.enabled` | Enable or disable Network metrics collection. |

> [!NOTE]
> You can opt-in for the metrics you want to measure.


> [!CAUTION]
> Start the metric collection by invoking this in a different goroutine!Please don't invoke in the main thread, as the function is blocking and will not allow your application logic to run if you place it above. It is strongly advised to run it in a different goroutine.

For updates, documentation, and community support, visit the [Canario GitHub repository](https://github.com/watchdogcloud/canario). 

Enjoy monitoring with Canario! 🐕