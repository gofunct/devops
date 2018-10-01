# Prometheus Study Guide

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
