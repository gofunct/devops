# Instrumenting Request Duration Metrics

## Overview

<details><summary>show</summary>
<p>

Prometheus bunch of very useful functions like rate(), increase() & histogram_quantile().

Adding metrics to your app is easy, just import ﻿prometheus client﻿ and register metrics HTTP handler `http.Handle("/metrics", promhttp.Handler())`

This one-liner adds HTTP /metrics endpoint to HTTP router. 

By default client exports memory usage, number of goroutines, gc information and other runtime information. 

</p>
</details>

## Using a Histogram

<details><summary>show</summary>
<p>

A histogram is made up 3 counters:

* a counter, which counts number of events that happened
* a counter for a sum of event values 
* a counter for each of a bucket

Histogram buckets:

* count how many times event value was less than or equal to the bucket’s value.

</p>
</details>

## Calculating Request Duration

<details><summary>show</summary>
<p>

It turns out that client library allows you to create a timer using: prometheus.NewTimer(o Observer) and record duration using ObserveDuration() method. Provided Observer can be either Summary, Histogram or a Gauge.

</p>
</details>

## Histogram Example 

<details><summary>show</summary>
<p>

Imagine that you create a histogram with 5 buckets with values: 0.5, 1, 2, 3, 5. Let’s call this histogram *http_request_duration_seconds* and 3 requests come in with durations 1s, 2s, 3s. Then you would see that */metrics* endpoint contains:

```

# HELP http_request_duration_seconds request duration histogram
# TYPE http_request_duration_seconds histogram
http_request_duration_seconds_bucket{le="0.5"} 0
http_request_duration_seconds_bucket{le="1"} 1
http_request_duration_seconds_bucket{le="2"} 2
http_request_duration_seconds_bucket{le="3"} 3
http_request_duration_seconds_bucket{le="5"} 3
http_request_duration_seconds_bucket{le="+Inf"} 3
http_request_duration_seconds_sum 6
http_request_duration_seconds_count 3

```

Here we can see that:
sum is 1s + 2s + 3s = *6*,
count is *3*, because of 3 requests
bucket {le=”0.5″} is *0*, because none of the requests where <= 0.5 seconds
bucket {le=”1″} is *1*, because one of the requests where <= 1 seconds
bucket {le=”2″} is *2*, because two of the requests where <= 2 seconds
bucket {le=”3″} is *3*, because all of the requests where <= 3 seconds

</p>
</details>

## Tips for Histograms

<details><summary>show</summary>
<p>

* when using Histogram we don’t need to have a separate counter to count total HTTP requests, as it creates one for us.

* We can calculate average request time by dividing sum over count. 

In PromQL it would be:

```
http_request_duration_seconds_sum / http_request_duration_seconds_count
```

* Also we could calculate percentiles (https://en.wikipedia.org/wiki/Percentile) from it. 
* Prometheus comes with a handy histogram_quantile function for it. 

For example calculating 50% percentile (second quartile) for last 10 minutes in PromQL would be:
```
histogram_quantile(0.5, rate(http_request_duration_seconds_bucket[10m])
```

Which results in 1.5.
Wait, 1.5? Shouldn’t it be 2? (50th percentile is supposed to be the median, the number in the middle)

* this value is only an *approximation* of computed quantile. 

* creating a new histogram requires you to specify bucket boundaries up front.

* The default values, which are 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10 are tailored to broadly measure the response time in seconds and probably won’t fit your app’s behavior.


* if you are instrumenting HTTP server or client, prometheus library has some helpers around it in promhttp package

</p>
</details>
