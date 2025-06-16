// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exampleexporter

import (
	"context"

	"github.com/songy23/otel-community-day-example/exporter/exampleexporter/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// NewFactory creates a factory for the example exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
		exporter.WithLogs(createLogsExporter, metadata.LogsStability),
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createLogsExporter(ctx context.Context,
	params exporter.Settings,
	config component.Config,
) (exporter.Logs, error) {
	exp := newExampleExporter(ctx, params)
	return exporterhelper.NewLogs(
		ctx,
		params,
		config,
		exp.ConsumeLogs,
		exporterhelper.WithStart(exp.start))
}

func createMetricsExporter(ctx context.Context,
	params exporter.Settings,
	config component.Config,
) (exporter.Metrics, error) {
	exp := newExampleExporter(ctx, params)
	return exporterhelper.NewMetrics(
		ctx,
		params,
		config,
		exp.ConsumeMetrics,
		exporterhelper.WithStart(exp.start))
}

func createTracesExporter(ctx context.Context,
	params exporter.Settings,
	config component.Config,
) (exporter.Traces, error) {
	exp := newExampleExporter(ctx, params)
	return exporterhelper.NewTraces(
		ctx,
		params,
		config,
		exp.ConsumeTraces,
		exporterhelper.WithStart(exp.start))
}
