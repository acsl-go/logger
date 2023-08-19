# acsl-go.logger
## Description
acsl-go.logger is a logger module that can be used to log messages to the console or an Elasticsearch instance.

## Installation
```bash
go get lab.afx.pub/ns/logger
```

## Usage

The basic logger function is:
```go
logger.Log(level int, format string, v ...interface{})
```
The `level` is used to determine what level the message is logged at. it could be one of the following values:

| CONSTANT NAME | INTEGER VALUE | DESCRIPTION |
|:--------------|:-------------:|:------------|
| logger.FATAL  | 0             | Fatal error |
| logger.ERROR  | 1             | Error       |
| logger.WARN   | 2             | Warning     |
| logger.INFO   | 3             | Information |
| logger.DEBUG  | 4             | Debugging message |

You can also use the `logger.Level` variable to specify the maximum level of messages to log. For example:
```go
logger.Level = logger.FATAL     // Only log messages with level FATAL
logger.Level = logger.INFO      // Only log messages with level INFO or higher
logger.Level = logger.DEBUG     // Log all messages
```

You can use the `logger.Log` directly, or use the following functions:
```go
logger.Fatal(format string, v ...interface{})
logger.Error(format string, v ...interface{})
logger.Warn(format string, v ...interface{})
logger.Info(format string, v ...interface{})
logger.Debug(format string, v ...interface{})
```

## The Extra Fields
The `logger` module will add some extra fields to the log message. The fields are:

| FIELD NAME | TYPE | METHOD | DESCRIPTION |
|:-----------|:----:|:------:|:------------|
| Timestamp  | int64/string | Stdout,Elastic | The timestamp of the message. |
| PID        | int | Stdout,Elastic | The process ID of the process that logged the message. |
| Process    | string | Elastic | The name of the process that logged the message. |
| Level      | int/string | Stdout,Elastic | The level of the message. |
| Identifier | string | Stdout,Elastic | The identifier of the message |

The data type will be different when using different output methods. For example, when using the `stdout` output method, the `Timestamp` field will be a string, and when using the `elastic` output method, the `Timestamp` field will be an int64 that represents the number of milliseconds since the Unix epoch.

By default, the identifier field will be the combination of the process id and the process name. you can specify the identifier by using the `logger.Identifier`:

```go
logger.Identifier = "my-identifier"
```

## The Output Methods
Presently, we support the following output methods:

| METHOD | FUNCTION NAME | TARGET |
|:-------|:--------------|:-------|
| stdout | logger.LogStdout | Standard output |
| elastic | logger.LogElastic | Elasticsearch |

By default, the `stdout` output method is used. You can change the output method by using the `logger.LogMethod` variable. The `logger.LogMethod` variable is an array of functions, and the index of the array is the level of the message. For example, if you want to use the `elastic` output method to log messages with level `ERROR`, you can use the following code:
```go
logger.LogMethod[logger.ERROR] = logger.LogElastic
```

### The usage of the `elastic` output method
The `elastic` output method will send the log message to an Elasticsearch instance. You `MUST` call `logger.InitElastic` function to initialize the `elastic` output method before specifying the `logger.LogElastic` to any element of the `logger.LogMethod` array. 

Here is an example:
```go
if err := InitElastic(&ElasticConfig{
    Addresses:     []string{"https://YOUR.ELASTIC:9200"},
    Username:      "elastic",
    Password:      "REPLACE WITH YOUR ELASTIC PASSWORD",
    CAFingerprint: "REPLACE WITH THE CA FINGERPRINT OF YOUR ELASTIC INSTANCE",
    Index:         "test",
}); err != nil {
    logger.LogMethod[logger.ERROR] = logger.LogStdout
} else {
    logger.LogMethod[logger.ERROR] = logger.LogElastic
}
```