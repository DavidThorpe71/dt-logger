## DT Logger idea

### Usage

- Create a new log (should do once per request) using `NewLog()`
- For each function call pass the `log` as an argument to the function
- Inside each function first open a new log context using `l.OpenContext()`
- You can immediately defer a call to close the log context, which will run before the function completes its execution: `defer l.CloseContext`
- There are four additional log method`s you can use:
  - `AddArg(key string, arg interface{})` - use to add the function arguments (one at a time)
  - `AddResponse(res interface{})`- use to add a function response (not any error responses) to the log
  - `AddError(err error)` - use to add an Error response to the log
  - `Write() string` - write the log as JSON to stdout
- Additional methods could be added, such as :
  - `AddMetric(metricName string)` - for use adding a specific Metric value to record (to aid in creating dashboards)

See `example/exampleApp.go` for an example of the usage and output, run the example using `go run exampleApp.go`