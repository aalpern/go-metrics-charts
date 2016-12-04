# go-metrics-charts

Realtime charting of live go-metrics metrics data.

This project began as a fork of
[github.com/mkevac/debugcharts](http://github.com/mkevac/debugcharts),
until it seemed cleaner to keep it separate. It provides a real-time
charting for applications that use the
[rcrowley/go-metrics](https://github.com/rcrowley/go-metrics) metrics
library.

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
