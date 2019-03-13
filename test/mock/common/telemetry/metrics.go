// Code generated by MockGen. DO NOT EDIT.
// Source: ../../../../pkg/common/telemetry/metrics.go

// Package telemetry is a generated GoMock package.
package telemetry

import (
	gomock "github.com/golang/mock/gomock"
	telemetry "github.com/spiffe/spire/pkg/common/telemetry"
	reflect "reflect"
	time "time"
)

// MockMetrics is a mock of Metrics interface
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// SetGauge mocks base method
func (m *MockMetrics) SetGauge(key []string, val float32) {
	m.ctrl.Call(m, "SetGauge", key, val)
}

// SetGauge indicates an expected call of SetGauge
func (mr *MockMetricsMockRecorder) SetGauge(key, val interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGauge", reflect.TypeOf((*MockMetrics)(nil).SetGauge), key, val)
}

// SetGaugeWithLabels mocks base method
func (m *MockMetrics) SetGaugeWithLabels(key []string, val float32, labels []telemetry.Label) {
	m.ctrl.Call(m, "SetGaugeWithLabels", key, val, labels)
}

// SetGaugeWithLabels indicates an expected call of SetGaugeWithLabels
func (mr *MockMetricsMockRecorder) SetGaugeWithLabels(key, val, labels interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGaugeWithLabels", reflect.TypeOf((*MockMetrics)(nil).SetGaugeWithLabels), key, val, labels)
}

// EmitKey mocks base method
func (m *MockMetrics) EmitKey(key []string, val float32) {
	m.ctrl.Call(m, "EmitKey", key, val)
}

// EmitKey indicates an expected call of EmitKey
func (mr *MockMetricsMockRecorder) EmitKey(key, val interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitKey", reflect.TypeOf((*MockMetrics)(nil).EmitKey), key, val)
}

// IncrCounter mocks base method
func (m *MockMetrics) IncrCounter(key []string, val float32) {
	m.ctrl.Call(m, "IncrCounter", key, val)
}

// IncrCounter indicates an expected call of IncrCounter
func (mr *MockMetricsMockRecorder) IncrCounter(key, val interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrCounter", reflect.TypeOf((*MockMetrics)(nil).IncrCounter), key, val)
}

// IncrCounterWithLabels mocks base method
func (m *MockMetrics) IncrCounterWithLabels(key []string, val float32, labels []telemetry.Label) {
	m.ctrl.Call(m, "IncrCounterWithLabels", key, val, labels)
}

// IncrCounterWithLabels indicates an expected call of IncrCounterWithLabels
func (mr *MockMetricsMockRecorder) IncrCounterWithLabels(key, val, labels interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrCounterWithLabels", reflect.TypeOf((*MockMetrics)(nil).IncrCounterWithLabels), key, val, labels)
}

// AddSample mocks base method
func (m *MockMetrics) AddSample(key []string, val float32) {
	m.ctrl.Call(m, "AddSample", key, val)
}

// AddSample indicates an expected call of AddSample
func (mr *MockMetricsMockRecorder) AddSample(key, val interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSample", reflect.TypeOf((*MockMetrics)(nil).AddSample), key, val)
}

// AddSampleWithLabels mocks base method
func (m *MockMetrics) AddSampleWithLabels(key []string, val float32, labels []telemetry.Label) {
	m.ctrl.Call(m, "AddSampleWithLabels", key, val, labels)
}

// AddSampleWithLabels indicates an expected call of AddSampleWithLabels
func (mr *MockMetricsMockRecorder) AddSampleWithLabels(key, val, labels interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSampleWithLabels", reflect.TypeOf((*MockMetrics)(nil).AddSampleWithLabels), key, val, labels)
}

// MeasureSince mocks base method
func (m *MockMetrics) MeasureSince(key []string, start time.Time) {
	m.ctrl.Call(m, "MeasureSince", key, start)
}

// MeasureSince indicates an expected call of MeasureSince
func (mr *MockMetricsMockRecorder) MeasureSince(key, start interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeasureSince", reflect.TypeOf((*MockMetrics)(nil).MeasureSince), key, start)
}

// MeasureSinceWithLabels mocks base method
func (m *MockMetrics) MeasureSinceWithLabels(key []string, start time.Time, labels []telemetry.Label) {
	m.ctrl.Call(m, "MeasureSinceWithLabels", key, start, labels)
}

// MeasureSinceWithLabels indicates an expected call of MeasureSinceWithLabels
func (mr *MockMetricsMockRecorder) MeasureSinceWithLabels(key, start, labels interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeasureSinceWithLabels", reflect.TypeOf((*MockMetrics)(nil).MeasureSinceWithLabels), key, start, labels)
}