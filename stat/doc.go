/*
  prometheus buried point library.

  simple usage:
    mgr := New(5 * time.Second)
    tick := mgr.NewTick("api_caller")
    defet tick.Close()

  Tick struct:
    this struct is used by api caller, it will automatically calculate
    qps, avg_cost, quantile_cost in p0, p50, p90, p99, p100 by samples
    by interval of New func.

    in prometheus it will export value like below:
    <name-by-NewTick>{tag="qps"} <qps-value>
    <name-by-NewTick>{tag="avg"} <avg-cost-value>
    <name-by-NewTick>{tag="p0"} <0%-cost-value>
    ...

  Counter struct:
    this struct is used by single value, it will export value in realtime,
    it is supported Set, Inc, Dec, Add, Sub with value.

    in prometheus it will export value like below:
    <name-by-NewCounter> <value>

  RawVec:
    you can use RawVec to create the GaugeVec of prometheus,
    GaugeVec usage: https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#GaugeVec
*/
package stat
