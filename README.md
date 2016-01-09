# go-logger
Simple Golang logger with only Debug, Info, and Fatal log levels

# Description
This logger is based on the recommendations of [Dave Cheney](Dave Cheney)

The Fatal log call should only be used from main and only after all necessary cleanup has been performed.

# Examples

Example code:
```
package main

import (
	log "github.com/alexrudd/go-logger"
)

func main() {
  log.Info("The default log level is:", "INFO")
  log.Debug("That means that this message is ignored")

  log.SetLogLevel(log.DEBUG)
  log.Debug("Changing the log level is easy")

  log.Fatal("And a 'Fatal' log call will exit the application")
}
```

Outputs:
```
[2016-01-09T21:06:50Z] INFO (main) The default log level is:INFO
[2016-01-09T21:06:50Z] DEBUG (main/main.go:12) Changing the log level is easy
[2016-01-09T21:06:50Z] FATAL (main/main.go:14) And a 'Fatal' log call will exit the application
```
