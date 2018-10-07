# Prometheus Study Guide

- [Prometheus Study Guide](#prometheus-study-guide)
- [[Metric types](https://prometheus.io/docs/concepts/metric_types)](#-metric-types--https---prometheusio-docs-concepts-metric-types-)
    + [Counters](#counters)
    + [Gauge](#gauge)
    + [Histogram](#histogram)
    + [Summary](#summary)
- [Aggregation Basics](#aggregation-basics)
    + [Aggregation basics on gauge type metrics](#aggregation-basics-on-gauge-type-metrics)
    + [Aggregation basics on counter type metrics](#aggregation-basics-on-counter-type-metrics)
    + [Aggregation basics on histogram type metrics](#aggregation-basics-on-histogram-type-metrics)
- [Aggregation Operators](#aggregation-operators)
    + [Grouping](#grouping)
    + [Operators](#operators)
- [References](#references)
- [Prometheus Exporters](#prometheus-exporters)
  * [What is Prometheus?](#what-is-prometheus-)
  * [Prometheus Text Form](#prometheus-text-form)
  * [Whats a Prometheus Exporter?](#whats-a-prometheus-exporter-)
  * [Exporter Examples](#exporter-examples)
  * [Finding Prometheus Exporters](#finding-prometheus-exporters)
  * [Building an Exporter](#building-an-exporter)
    + [Registering the collector](#registering-the-collector)
    + [Set up a handler for metrics](#set-up-a-handler-for-metrics)
    + [Listen for connections](#listen-for-connections)
  * [The Collector Interface](#the-collector-interface)
    + [Metric description](#metric-description)
    + [Gathering metrics](#gathering-metrics)
  * [Constructing a collector](#constructing-a-collector)
  * [Implementing the Collector](#implementing-the-collector)
  * [Sources of Metrics](#sources-of-metrics)
  * [Gathering Metrics from /proc/stat](#gathering-metrics-from--proc-stat)
    + [Parsing /proc/stat in Go](#parsing--proc-stat-in-go)
  * [Parsing in Go](#parsing-in-go)
    + [CPUStat](#cpustat)
    + [Scan](#scan)
    + [Handle slice bounds](#handle-slice-bounds)
  * [Try out your API for /proc/stat](#try-out-your-api-for--proc-stat)
  * [Wire up Dependencies](#wire-up-dependencies)


***


# [Metric types](https://prometheus.io/docs/concepts/metric_types)

### Counters

<details><summary>show</summary>
<p>

- Always increasing metric, e.g. http_requests_total
- Gets reset(to zero) every now and then, because of process restarts
- Prometheus has safe methods e.g. rate(), increase() which understands and handles counter resets

</p>
</details>

### Gauge

<details><summary>show</summary>
<p>

- Absolute value (or counts) which can go up and down, e.g. cpu_memory_usage, number_of_go_routines etc
- Between scrapes, spikes are missed

</p>
</details>

### Histogram

<details><summary>show</summary>
<p>

- Client side sampling of obervations and counted in buckets
- Internally actually exposes various counter type metrics:
    - x_sum
      → Sum of total observed valuse
    - x_count
      → Count of total number of observations
    - x_bucket{le=v1}
    - x_bucket{le=v2}
    - x_bucket{le=+Inf}
      → Each bucket counting number of times observed value was less than or equal to v1, v2.., +Inf respectively
- Unlike gauge, doesn't misses spikes between scrapes because client is sampling all the observations anyway

</p>
</details>

### Summary

<details><summary>show</summary>
<p>

- Use with caution! [Reference](https://prometheus.io/docs/concepts/metric_types/#summary)
- It is not possible for you to aggregate the quantiles of a summary (the time series with the quantile label) from a statistical standpoint

</p>
</details>

# Aggregation Basics

### Aggregation basics on gauge type metrics

<details><summary>show</summary>
<p>

Consider metric *node_filesystem_size_bytes* from Node exporter, which reports the size of each of your mounted filesystems, and has device, fstype, and mountpoint labels.

→ Sums everything up with same labels, gets total sum of filesystem size of all machines being monitored:
```
sum(node_filesystem_size_bytes)
```

→ Sums everything up with same labels only taking in those in *by*, gets filesystem size of all devices on each machines:
```
sum by(instance, device)(node_filesystem_size_bytes)
```

→ Sums everything up with same lables ignoring those in *without*, gets filesystem size of each machines(because this label is not in without):
```
sum without(device, fstype, mountpoint)(node_filesystem_size_bytes)
```

→ Gets size of the biggest mounted filesystem on each machine:
```
# Using by
max by(instance)(node_filesystem_size_bytes)

# Using without
max without(device, fstype, mountpoint)(node_filesystem_size_bytes)
```
- When math has been performed on a metric, metric name is no longer returned. The response is always some value against combination(s) of instrumentation and/or target labels or {}.

→ Gets change in memory usage in the Node exporter over past hour
```
process_resident_memory_bytes{job="node"}
-
process_resident_memory_bytes{job="node"} offset 1h
```

</p>
</details>

### Aggregation basics on counter type metrics

<details><summary>show</summary>
<p>

→ Gets amount of network traffic received per second:
```
rate(node_network_receive_bytes_total[5m])
```
- The [5m] says to provide rate with 5 minutes of data, so the returned value will be an average over the last 5 minutes
- The output of rate is gauge, so further aggregations could be applied similarly.

→ Gets total bytes received per machine per second:
```
sum by(instance)(rate(node_network_receive_bytes_total[5m]))
```

### Aggregation basics on histogram type metrics

Consider Prometheus 2.2.1 exposing a histogram metric called *prometheus_tsdb_compaction_duration_seconds* that tracks how many seconds compaction takes for the time series database. Obviously it will under the hood expose thre counter metrics.

→ Gets total number of times compaction happens per second per instance
```
sum by(instance)(rate(prometheus_tsdb_compaction_duration_seconds_count[5m]))
```

→ Gets average compaction seconds per instance over a time period of 5m
```
sum by(instance)(rate(prometheus_tsdb_compaction_duration_seconds_sum[5m]))
/
sum by(instance)(rate(prometheus_tsdb_compaction_duration_seconds_count[5m]))
```

→ Gets 90%ile value of compaction seconds over a  of 1d:
```
histogram_quantile(
    0.90,
    rate(prometheus_tsdb_compaction_duration_seconds_bucket[1d]))
```

</p>
</details>

# Aggregation Operators

### Grouping

<details><summary>show</summary>
<p>

→ Both **without** & **by** work on gauge type metric as in many examples above

</p>
</details>

### Operators

<details><summary>show</summary>
<p>

→ **sum**
Adds of all the values in a group and returns that as a value for the group.

→ **count**
Counts the number of time series in a group, and returns it as the value for the group.

→ **avg**
Returns the average of the values4 of the time series in the group as the value for the group.

→ **stddev**, **stdvar**

→ **min**, **max**
The min and max aggregators return the minimum or maximum value within a group as the value of the group.

→ **topk**
Returns the K time series with the biggest values
```
topk without(device, fstype, mountpoint)(K, node_filesystem_size_bytes)
```
- The labels of time series they return for a group are not the labels of the group
- They can return more than one time series per group

→ **bottomk**
Returns the K time series with the lowest values

→ **quantile**
Returns the specified quantile of the values of the group as the group’s return value.
Gets 90th percentile of the system mode CPU usage across the different CPUs in each of machines:
```
quantile without(cpu)(0.9, rate(node_cpu_seconds_total{mode="system"}[5m]))
```

→ **count_values**
Builds a frequency histogram of the values of the time series in the group, with the count of each value as the value of the output time series and the original value as a new label.
E.g. given following time series:
```
software_version{instance="a",job="j"} 7
software_version{instance="b",job="j"} 4
software_version{instance="c",job="j"} 8
software_version{instance="d",job="j"} 4
software_version{instance="e",job="j"} 7
software_version{instance="f",job="j"} 4
```
Following query:
```
count_values without(instance)("version", software_version)
```
Returns:
```
{job="j",version="7"} 2
{job="j",version="8"} 1
{job="j",version="4"} 3
```

</p>
</details>

# References

<details><summary>show</summary>
<p>

- [Offical documentation](https://prometheus.io/docs/introduction/overview/)
- [Book - Prometheus: Up & Running](https://www.safaribooksonline.com/library/view/prometheus-up/9781492034131/)

</p>
</details>

# Prometheus Exporters

## What is Prometheus?

<details><summary>show</summary>
<p>

* Open sourced monitoring and alerts toolkit

* One of the graduated projects from the Cloud Native Computing Foundation

* Can be used with a client called PromQL- a query language

</p>
</details>


## Prometheus Text Form 

<details><summary>show</summary>
<p>

```bash

curl -s http://localhost:9100/metrics | grep node

```

```
# HELP node_arp_entries ARP entries by device
# TYPE node_arp_entries gauge
node_arp_entries{device="br0"} 7
# HELP node_boot_time Node boot time, in unixtime.
# TYPE node_boot_time gauge
node_boot_time 1.521387979e+
# HELP node_context_switches Total number of context switches.
# TYPE node_context_switches counter
node_context_switches 1.55007032e+

```

</p>
</details>


## Whats a Prometheus Exporter?

<details><summary>show</summary>
<p>

* Exporters bridge the gap between Prometheus and systems that don't export metrics in the Prometheus format

* Typically run on the same machine as a service, but not always!

</p>
</details>

## Exporter Examples

<details><summary>show</summary>
<p>

node exporter:
* exposes Unix-like system metrics

mysqld exporter:
* Exposes metrics from a MySQL server

blackbox exporter:
* exposes metrics from blackbox systems via HTTP, ICMP, etc

</p>
</details>

## Finding Prometheus Exporters

<details><summary>show</summary>
<p>

* [Prometheus.io](https://prometheus.io/docs/instrumenting/exporters/)

* Search Github!

* Otherwise, code your own!

</p>
</details>

## Building an Exporter

<details><summary>show</summary>
<p>

### Registering the collector

    c := newCollector(x)
    prometheus.MustRegister(c)

### Set up a handler for metrics
    
    mux := http.NewServeMux()
    mux.Handle("/metrics", promhttp.Handler())

### Listen for connections

    const addr = ":8080"
    log.Printf("starting exporter on %q", addr)
    if err := http.ListenAndServe(addr, mux); err != nil {
        log.Fatalf("cannot start exporter: %s", err)
    }
    
</p>
</details>

## The Collector Interface

<details><summary>show</summary>
<p>

```go

type collector struct {
    
}
```

### Metric description

```go
RequestsTotal *prometheus.Desc

```

### Gathering metrics

```go
requests func() (int, error)
```



### Constructing a collector

   
```go

return &collector{
    RequestsTotal: prometheus.NewDesc(
        // metric name
        "total_requests",
        // help text
        "The total number of requests that occur.",
        // label dimensions
        nil, nil,
    ),
    requests: /* dependencies */,
    }
}

```



### Implementing the Collector


```go

func (c *collector) Collect(ch chan<- prometheus.Metric) {
    // metrics snapshot
    // must be concurrency safe
    requests, err := c.requests()
    if err != nil {
        // notify Promeheus of error
        ch <- prometheus.NewInvalidMetric(c.RequestsTotal, err)
        return
    }
    // use "const metric" for constructors
    ch <- prometheus.MustNewConstMetric(
    c.RequestsTotal, prometheus.CounterValue, requests,
    )
    
}
```

</p>
</details>

## Tips

<details><summary>show</summary>
<p>

* Build Reusable Metrics
* Write unit tests to ping your metrics endpoint to check that the output is what you wanted
* use promtool to check metrics for linting

```bash
curl http://localhost:8080/metrics | promtool check metrics 
x_gigabytes counter metrics should have "_total" suffix
x_gigabytes use base unit bytes instead of "gigabytes"
```

* Use io.Reader interface whenever possible
* bufio.Scanner is your friend
* always check slice/array bounds
* checkout [procfs]github.com/prometheus/procfs



</p>
</details>

## Sources of Metrics

<details><summary>show</summary>
<p>

* Files
* System Calls
* Hardware Devices

</p>
</details>

## Gathering Metrics from /proc/stat

<details><summary>show</summary>
<p>
* Kernel/System statistics
* Number indicates CPU spent in various states like "user", "system", "idle"

</p>
</details>

### Parsing /proc/stat in Go

<details><summary>show</summary>
<p>

## Parsing in Go
* create a clear and concise exporter API

### CPUStat

```go
// CPUStat contains stats for an individual CPU
type CPUStat struct {
    //ID of CPU
    ID string
    
    // Time in User_HZ(1/100th of second)
    User, System, Idle int
}
```

### Scan

* reads and parses CPUStat from r

```go
func Scan(r io.Reader) ([]CPUStat, error) {
    s := bufio.NewScanner(r)
    s.Scan() //skip first summarized line
    
    var stats []CPUStat
    for s.Scan() {/* ... */ }
    
    // check the error
    if err := s.Err(); err != nil {
        return nil, err
    }
    return stats, nil
}

```

### Handle slice bounds

```go

for s.Scan() {
    // Each CPU stats line should have cpu prefix
    // 11 fields
    
    const nFields = 11
    fields := strings.Fields(string(s.Bytes()))
    if len(fields) != nFields(
        continue
    )
    if !strings.HasPrefix(fields[0], "cpu") {
        continue
    }
    
    // Values we care about: (user, system, idle) 
    // lie at indecies 1, 3 , 4
    // Parse them into an array
    var times [3]int
    for i, idx := range []int{1,3,4} {
        v, err := strconv.Atoi(fields[idx])
        if err != nil {
            return nil, err
        }
        time[i] = v
    }
    stats = append(stats, CPUStat{
        //first files is CPU ID
        Id: fields[0],
        User: times[0],
        System: times[1],
        Idle: times[2],
    })
} // End of loop

```
</p>
</details>

## Try out your API for /proc/stat

<details><summary>show</summary>
<p>

```go

f, err := os.Open("/proc/stat")
if err != nil {
log.Fatalf("failed to open /proc/stat: %v", err)
}
defer f.Close()

stats, err := cpustat.Scan(f)
if err != nil {
log.Fatalf("failed to scan: %v", err)
}

for _, s := range stats {
fmt.Printf("%4s: user: %06d, system: %06d, idle: %06d\n",
s.ID, s.User, s.System, s.Idle)
}

```

```go
go build

```

```bash
./cpustat
```

</p>
</details>

## Wire up Dependencies

<details><summary>show</summary>
<p>

```go

// Called on each collector.Collect.
stats := func() ([]cpustat.CPUStat, error) {
f, err := os.Open("/proc/stat")
if err != nil {
return nil, fmt.Errorf("failed to open: %v", rr)
}
defer f.Close()

return cpustat.Scan(f)
}

// Make Prometheus client aware of our collector.
c := newCollector(stats)
prometheus.MustRegister(c)

 // A collector is a prometheus.Collector for Linux CPU stats.

type collector struct {
// Possible metric descriptions.
TimeUserHertzTotal *prometheus.Desc

// A parameterized function used to gather metrics.
stats func() ([]cpustat.CPUStat, error)
}

stats, err := c.stats()
if err != nil {
ch <- prometheus.NewInvalidMetric(c.TimeUserHertzTotal, err)
return
}
for _, s := range stats {
tuples := []struct {
mode string
v int
}{
{mode: "user", v: s.User},
{mode: "system", v: s.System},
{mode: "idle", v: s.Idle},
}
}

for _, t := range tuples { (^)
// prometheus.Collector implementations should always use
// "const metric" constructors.
ch <- prometheus.MustNewConstMetric(
c.TimeUserHertzTotal,
prometheus.CounterValue,
float64(t.v),
s.ID, t.mode,
)
}

```

```bash 
curl [http://localhost:8888/metrics](http://localhost:8888/metrics) | head -n 5
# HELP cpustat_time_user_hertz_total Time in USER_HZ a given CPU
spent in a given mode.
# TYPE cpustat_time_user_hertz_total counter
cpustat_time_user_hertz_total{cpu="cpu0",mode="idle"}
1.597421e+06
cpustat_time_user_hertz_total{cpu="cpu0",mode="system"} 39621
cpustat_time_user_hertz_total{cpu="cpu0",mode="user"} 160345**

```

</p>
</details>


