// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exampleexporter

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type exampleExporter struct {
	params exporter.Settings
	ctx    context.Context
}

func newExampleExporter(ctx context.Context, params exporter.Settings) *exampleExporter {
	exp := &exampleExporter{
		params: params,
		ctx:    ctx,
	}
	return exp
}

func (e *exampleExporter) start(_ context.Context, _ component.Host) error {
	return nil
}

func (e *exampleExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (e *exampleExporter) ConsumeMetrics(_ context.Context, m pmetric.Metrics) error {
	// insert business logic here
	fmt.Printf("received %d metrics\n", m.MetricCount())
	return nil
}

func (e *exampleExporter) ConsumeLogs(_ context.Context, l plog.Logs) error {
	// insert business logic here
	fmt.Printf("received %d logs\n", l.LogRecordCount())
	return nil
}

func (e *exampleExporter) ConsumeTraces(_ context.Context, t ptrace.Traces) error {
	// insert business logic here
	fmt.Printf("received %d spans\n", t.SpanCount())
	return nil
}
