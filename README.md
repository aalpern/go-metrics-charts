# go-metrics-charts

Realtime charting of live go-metrics metrics data.

This project began as a fork of
[github.com/mkevac/debugcharts](http://github.com/mkevac/debugcharts),
until it seemed cleaner to keep it separate. It provides real-time
charting for applications that use the
[rcrowley/go-metrics](https://github.com/rcrowley/go-metrics) metrics
library.

This is intended as a simple live view of metrics in a single process
without the need for an entire metrics stack - it can be useful when
developing and testing locally. All third party dependencies are
loaded from a CDN to keep the amount of embedded data to a minimum.

## Installation

`go get github.com/aalpern/go-metrics-charts`

## Usage

```
package main

import _ "github.com/aalpern/go-metrics-charts"
```

You must have metrics exposed via expvars, as rather than having a
binary depedency on the metrics package, the UI works by hitting the
`/debug/metrics` JSON endpoint exposed by that package.

```
import "github.com/rcrowley/go-metrics/exp"

exp.Exp(metrics.DefaultRegistry)
```

The live charts page will be registered at `/debug/metrics/charts`.

## TO-DO

* [ ] Chart options - add stacked area, display data points
* [ ] pause/restart and change polling interval
* [ ] better handling of long list of metrics
* [ ] implement filtering of metrics list
