// Package provides a logging utility designed for AWS Lambda functions
// and API Gateway. It offers a globally accessible logger configured to output
// JSON formatted logs. The package also includes utility functions for adding
// additional metadata fields to logs, either from custom fields or directly
// from AWS API Gateway and Custom Authorizer request objects.
//
// Example Usage:
//
//	// Initializing and getting the logger
//	logger := GetLogger()
//
//	// Logging with a simple message
//	PrepareLog("A simple message").Info()
//
//	// Logging with additional fields
//	PrepareLog("With extra fields").Add("field1", "value1").Warn()
//
//	// Logging with API Gateway request details
//	PrepareLog("API Gateway request info").AddGatewayRequest(request).Info()
//
// Global Logger:
//
// The logger is initialized only once and is globally accessible via GetLogger().
// It's configured to output JSON-formatted logs and includes metadata like 'env'
// and 'app_name' from environment variables.
//
// Custom Log Messages:
//
// The Log struct provides a way to prepare a log message with additional
// metadata fields. You can add custom fields using the Add() method or append
// AWS request details using AddAuthorizerRequest() and AddGatewayRequest().
package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var globalLogger *log.Entry // Singleton logger instance

// GetLogger initializes and returns the global logger.
func GetLogger() *log.Entry {
	if globalLogger == nil {
		// Logger setup
		log.SetFormatter(&log.TextFormatter{
			ForceColors:      true,
			DisableTimestamp: true,
		})
		log.SetReportCaller(false)
		log.SetOutput(os.Stdout)
		globalLogger = log.WithFields(log.Fields{})
	}

	return globalLogger
}

// Logf creates a new Log with formatted message.
func Logf(format string, args ...interface{}) *Log {
	return &Log{Message: fmt.Sprintf(format, args...), Fields: make(map[string]interface{})}
}

// Log holds the log message and additional fields.
type Log struct {
	Message string     `json:"message"`
	Fields  log.Fields `json:"fields"`
}

// Add adds a key-value pair to the Log's Fields
// Chainable: Can be chained with other methods.
func (l Log) Add(key string, value string) Log {
	l.Fields[key] = value

	return l
}

func (l Log) AddError(err error) Log {
	l.Fields["error"] = err.Error()

	return l
}

// Info logs the message at Info level
// Chainable: Can be chained with other methods.
func (l Log) Info() {
	GetLogger().WithFields(l.Fields).Info(l.Message)
}

// Debug logs the message at Debug level
// Chainable: Can be chained with other methods.
func (l Log) Debug() {
	GetLogger().WithFields(l.Fields).Debug("🐛 " + l.Message)
}

// Warn logs the message at Warn level
// Chainable: Can be chained with other methods.
func (l Log) Warn() {
	GetLogger().WithFields(l.Fields).Warn("⚠️ " + l.Message)
}

// Error logs the message at Error level
// Chainable: Can be chained with other methods.
func (l Log) Error() {
	GetLogger().WithFields(l.Fields).Error(l.Message)
}

// Fatal logs the message at Fatal level
// Chainable: Can be chained with other methods.
func (l Log) Fatal() {
	GetLogger().WithFields(l.Fields).Fatal(l.Message)
}
