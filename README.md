# Lightning Log
 [![Build Status](https://travis-ci.org/janritter/go-lightning-log.svg?branch=master)](https://travis-ci.org/janritter/go-lightning-log)
 [![](https://godoc.org/github.com/janritter/go-lightning-log?status.svg)](http://godoc.org/github.com/janritter/go-lightning-log)
 [![Codacy Badge](https://api.codacy.com/project/badge/Grade/ea78dc8ca33c4269bbe8767e9b0a1fe4)](https://www.codacy.com/app/jan-ritter/go-lightning-log?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=janritter/go-lightning-log&amp;utm_campaign=Badge_Grade)

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