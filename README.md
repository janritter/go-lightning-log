# Lightning Log
 [![Build Status](https://travis-ci.org/janritter/go-lightning-log.svg?branch=master)](https://travis-ci.org/janritter/go-lightning-log)
 [![](https://godoc.org/github.com/janritter/go-lightning-log?status.svg)](http://godoc.org/github.com/janritter/go-lightning-log)
 [![Maintainability](https://api.codeclimate.com/v1/badges/b7ce8b1347d20e16ce35/maintainability)](https://codeclimate.com/github/janritter/go-lightning-log/maintainability)
 [![Coverage Status](https://coveralls.io/repos/github/janritter/go-lightning-log/badge.svg?branch=master)](https://coveralls.io/github/janritter/go-lightning-log?branch=master)

## Usage

### Get package
```bash
go get github.com/janritter/go-lightning-log
```

## Example using default logger

The default logger severity level is 4, which means all events are processed

```go
package main

import (
	lightning "github.com/janritter/go-lightning-log"
)

func main() {
	// This is an error created in your code
	err := errors.New("this is a debug error")

	// Tags can be used to add key value pairs to every logging operation
	tags := map[string]string{"module": "controller/bookings"}

	// Defines the severity level of the log event, it is used in the output and for the decision if the event should be logged
	// Allowed levels are 0 = FATAL" 1 = ERROR" 2 = "WARNING" 3 = "INFO" 4 = "DEBUG"
	severity := 4

	// Log with default logger
	lightning.Log(err, tags, severity)
}
```
For all available functions, check the godoc

## Example using a custom logger

```go
package main

import (
	lightning "github.com/janritter/go-lightning-log"
)

func main() {

	// This creates a new custom logger with a min severity of 2, which means all events with a severity level higher than 2 are ignored (all info and debug events)
	logger, _ := lightning.Init(2)

	// This is an error created in your code
	err := errors.New("this is a debug error")

	// Tags can be used to add key value pairs to every logging operation
	tags := map[string]string{"module": "controller/bookings"}

	// Defines the severity level of the log event, it is used in the output and for the decision if the event should be logged
	// Allowed levels are 0 = FATAL" 1 = ERROR" 2 = "WARNING" 3 = "INFO" 4 = "DEBUG"
	severity := 2

	// Log with custom logger
	logger.Log(err, tags, severity)
}
```
For all available functions, check the godoc

## Development

### Install dependencies
```bash
make prepare
```

### Execute tests
```bash
make tests
```

## License and Author

Author: Jan Ritter

License: MIT