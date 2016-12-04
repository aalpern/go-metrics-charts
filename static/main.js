/**
 * Base class for objects which are event sources.
 */
class Observable {
  on(event, handler) {
    bean.on(this, event, handler)
  }

  fire(event, ...data) {
    bean.fire(this, event, ...data)
  }
}

/**
 * MetricSet encapsulates the set of selected metrics to be graphed
 * and the collected data for that graph.
 */
class MetricSet extends Observable {
  /** parent is a Metrics instance which this MetricSet is a subset of */
  constructor(parent) {
    super()
    this.parent = parent
    this.metrics = new Set()
    this.data = {}

    this.parent.on('update', () => {
      // Update the metrics listing table
      this.render()

      // Capture datapoints for any selected metrics
      for (let name of this.metrics) {
        this.store_point(name)
      }

      this.fire('update', this)
    })

    this.on('update', () => { this.render() } )
  }

  render() {
    $('#metric-table').html(templates.metric_table_template({
      data: this.parent.names.map(name => {
        return {
          name: name,
          value: this.parent.value(name),
          selected: Data.selected.has(name)
        }
      })
    }))
  }

  add(metric) {
    if (!this.has(metric)) {
      this.metrics.add(metric)
      this.store_point(metric)
      this.fire('add', metric)
      this.fire('update')
    }
  }

  remove(metric) {
    if (this.has(metric)) {
      this.metrics.delete(metric)
      delete this.data[metric]
      this.fire('remove', metric)
      this.fire('update')
    }
  }

  clear() {
    this.metrics.clear()
    this.data = {}
    this.fire('clear')
    this.fire('update')
  }

  has(metric) {
    return this.metrics.has(metric)
  }

  store_point(metric) {
    let points = this.data[metric] || []
    points.push({
      timestamp: this.parent.timestamp,
      value: this.parent.metrics[metric]
    })
    this.data[metric] = points
  }

  plotly_data() {
    return Object.keys(this.data).map(name => {
      let raw = this.data[name]
      return {
        x: raw.map(point => point.timestamp),
        y: raw.map(point => point.value),
        mode: 'lines',
        name: name
      }
    })
  }
}

/**
 * Metrics encapsulates the set of all metrics and their most recent
 * values, fetched from the /debug/metrics endpoint.
 */
class Metrics extends Observable {
  constructor() {
    super()
    this.metrics = {}
    this.names = []
    this.timestamp = null
    this.period = 5000
  }

  value(metric) {
    return this.metrics[metric]
  }

  update() {
    axios.get('/debug/metrics')
      .then((rsp) => {
        this.metrics = rsp.data.metrics
        this.names = Object.keys(this.metrics)
        this.timestamp = new Date()
      })
      .then(() => {
        this.fire('update', this)
      })
  }

  start() {
    stop()
    this.interval = window.setInterval(() => this.update(), this.period)
  }

  stop() {
    if (this.interval) {
      window.clearInterval(this.interval)
    }
  }
}

function redraw() {
    Data.plot.data = Data.selected.plotly_data()
    Plotly.redraw(Data.plot)
}

function compile_template(selector) {
  return Handlebars.compile($(selector).html())
}

function toggle_metric(name) {
  let set = Data.selected
  if (set.has(name)) {
    set.remove(name)
  } else {
    set.add(name)
  }
}

function init() {
  templates = {
    metric_table_template: compile_template('#metric-table-template')
  }
  Data = {}
  Data.metrics = new Metrics()
  Data.selected = new MetricSet(Data.metrics)
  Data.metrics.update()

  Plotly.newPlot('chart', Data.selected.plotly_data())
  Data.plot = $('#chart')[0]
  Data.selected.on('update', redraw)
  Data.metrics.start()
}

$(document).ready(() => {
  init()

  $(document).on('click', 'tr.metric-row td.name', (el) => {
    let metric = el.target.textContent
    toggle_metric(metric)
  })
  $(document).on('click', '#clear-button', (el) => {
    Data.selected.clear()
  })
})
