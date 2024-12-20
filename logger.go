/*
 * IONOS Object Storage API for contract-owned buckets
 *
 * ## Overview The IONOS Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless compatibility with existing tools and libraries tailored for S3 systems.  ### API References - [S3 API Reference for contract-owned buckets](https://api.ionos.com/docs/s3-contract-owned-buckets/v2/) ### User documentation [IONOS Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage) * [Documentation on user-owned and contract-owned buckets](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/buckets) * [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility) * [S3 Tools](https://docs.ionos.com/cloud/managed-services/s3-object-storage/s3-tools)  ## Endpoints for contract-owned buckets | Location | Region Name | Bucket Type | Endpoint | | --- | --- | --- | --- | | **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |  ## Changelog - 30.05.2024 Initial version
 *
 * API version: 2.0.2
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"log"
	"os"
	"strings"
)

type LogLevel uint

func (l *LogLevel) Get() LogLevel {
	if l != nil {
		return *l
	}
	return Off
}

// Satisfies returns true if this LogLevel is at least high enough for v
func (l *LogLevel) Satisfies(v LogLevel) bool {
	return l.Get() >= v
}

const (
	Off LogLevel = 0x100 * iota
	Debug
	// Trace We recommend you only set this field for debugging purposes.
	// Disable it in your production environments because it can log sensitive data.
	// It logs the full request and response without encryption, even for an HTTPS call.
	// Verbose request and response logging can also significantly impact your application's performance.
	Trace
)

var LogLevelMap = map[string]LogLevel{
	"off":   Off,
	"debug": Debug,
	"trace": Trace,
}

// getLogLevelFromEnv - gets LogLevel type from env variable IONOS_LOG_LEVEL
// returns Off if an invalid log level is encountered
func getLogLevelFromEnv() LogLevel {
	strLogLevel := "off"
	if os.Getenv(IonosLogLevelEnvVar) != "" {
		strLogLevel = os.Getenv(IonosLogLevelEnvVar)
	}

	logLevel, ok := LogLevelMap[strings.ToLower(strLogLevel)]
	if !ok {
		log.Printf("Cannot set logLevel for value: %s, setting loglevel to Off", strLogLevel)
	}
	return logLevel
}

type Logger interface {
	Printf(format string, args ...interface{})
}

func NewDefaultLogger() Logger {
	return &defaultLogger{
		logger: log.New(os.Stderr, "IONOSLOG ", log.LstdFlags),
	}
}

type defaultLogger struct {
	logger *log.Logger
}

func (l defaultLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}
