// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exampleexporter

import (
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration for the example exporter.
type Config struct {
	confighttp.ClientConfig   `mapstructure:",squash"`        // squash ensures fields are correctly decoded in embedded struct.
	QueueSettings             exporterhelper.QueueBatchConfig `mapstructure:"sending_queue"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`

	// insert business configs below
}
