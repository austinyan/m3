// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db/storage/series/types.go

package series

import (
	gomock "github.com/golang/mock/gomock"
	clock "github.com/m3db/m3db/clock"
	context "github.com/m3db/m3db/context"
	encoding "github.com/m3db/m3db/encoding"
	instrument "github.com/m3db/m3db/instrument"
	persist "github.com/m3db/m3db/persist"
	retention "github.com/m3db/m3db/retention"
	block "github.com/m3db/m3db/storage/block"
	ts "github.com/m3db/m3db/ts"
	io "github.com/m3db/m3db/x/io"
	time0 "github.com/m3db/m3x/time"
	time "time"
)

// Mock of DatabaseSeries interface
type MockDatabaseSeries struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseSeriesRecorder
}

// Recorder for MockDatabaseSeries (not exported)
type _MockDatabaseSeriesRecorder struct {
	mock *MockDatabaseSeries
}

func NewMockDatabaseSeries(ctrl *gomock.Controller) *MockDatabaseSeries {
	mock := &MockDatabaseSeries{ctrl: ctrl}
	mock.recorder = &_MockDatabaseSeriesRecorder{mock}
	return mock
}

func (_m *MockDatabaseSeries) EXPECT() *_MockDatabaseSeriesRecorder {
	return _m.recorder
}

func (_m *MockDatabaseSeries) ID() ts.ID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(ts.ID)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockDatabaseSeries) Tick() error {
	ret := _m.ctrl.Call(_m, "Tick")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) Tick() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tick")
}

func (_m *MockDatabaseSeries) Write(ctx context.Context, timestamp time.Time, value float64, unit time0.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockDatabaseSeries) ReadEncoded(ctx context.Context, start time.Time, end time.Time) ([][]io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseSeriesRecorder) ReadEncoded(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2)
}

func (_m *MockDatabaseSeries) FetchBlocks(ctx context.Context, starts []time.Time) []block.FetchBlockResult {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, starts)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) FetchBlocks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1)
}

func (_m *MockDatabaseSeries) FetchBlocksMetadata(ctx context.Context, includeSizes bool, includeChecksums bool) block.FetchBlocksMetadataResult {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, includeSizes, includeChecksums)
	ret0, _ := ret[0].(block.FetchBlocksMetadataResult)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) FetchBlocksMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1, arg2)
}

func (_m *MockDatabaseSeries) IsEmpty() bool {
	ret := _m.ctrl.Call(_m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) IsEmpty() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsEmpty")
}

func (_m *MockDatabaseSeries) IsBootstrapped() bool {
	ret := _m.ctrl.Call(_m, "IsBootstrapped")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) IsBootstrapped() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsBootstrapped")
}

func (_m *MockDatabaseSeries) Bootstrap(rs block.DatabaseSeriesBlocks) error {
	ret := _m.ctrl.Call(_m, "Bootstrap", rs)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) Bootstrap(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0)
}

func (_m *MockDatabaseSeries) Flush(ctx context.Context, blockStart time.Time, persistFn persist.Fn) error {
	ret := _m.ctrl.Call(_m, "Flush", ctx, blockStart, persistFn)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseSeriesRecorder) Flush(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0, arg1, arg2)
}

func (_m *MockDatabaseSeries) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockDatabaseSeriesRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockDatabaseSeries) Reset(id ts.ID) {
	_m.ctrl.Call(_m, "Reset", id)
}

func (_mr *_MockDatabaseSeriesRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

// Mock of DatabaseSeriesPool interface
type MockDatabaseSeriesPool struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseSeriesPoolRecorder
}

// Recorder for MockDatabaseSeriesPool (not exported)
type _MockDatabaseSeriesPoolRecorder struct {
	mock *MockDatabaseSeriesPool
}

func NewMockDatabaseSeriesPool(ctrl *gomock.Controller) *MockDatabaseSeriesPool {
	mock := &MockDatabaseSeriesPool{ctrl: ctrl}
	mock.recorder = &_MockDatabaseSeriesPoolRecorder{mock}
	return mock
}

func (_m *MockDatabaseSeriesPool) EXPECT() *_MockDatabaseSeriesPoolRecorder {
	return _m.recorder
}

func (_m *MockDatabaseSeriesPool) Get() DatabaseSeries {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(DatabaseSeries)
	return ret0
}

func (_mr *_MockDatabaseSeriesPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockDatabaseSeriesPool) Put(block DatabaseSeries) {
	_m.ctrl.Call(_m, "Put", block)
}

func (_mr *_MockDatabaseSeriesPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

// Mock of databaseBuffer interface
type MockdatabaseBuffer struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseBufferRecorder
}

// Recorder for MockdatabaseBuffer (not exported)
type _MockdatabaseBufferRecorder struct {
	mock *MockdatabaseBuffer
}

func NewMockdatabaseBuffer(ctrl *gomock.Controller) *MockdatabaseBuffer {
	mock := &MockdatabaseBuffer{ctrl: ctrl}
	mock.recorder = &_MockdatabaseBufferRecorder{mock}
	return mock
}

func (_m *MockdatabaseBuffer) EXPECT() *_MockdatabaseBufferRecorder {
	return _m.recorder
}

func (_m *MockdatabaseBuffer) Write(ctx context.Context, timestamp time.Time, value float64, unit time0.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) Write(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockdatabaseBuffer) ReadEncoded(ctx context.Context, start time.Time, end time.Time) [][]io.SegmentReader {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, start, end)
	ret0, _ := ret[0].([][]io.SegmentReader)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) ReadEncoded(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2)
}

func (_m *MockdatabaseBuffer) FetchBlocks(ctx context.Context, starts []time.Time) []block.FetchBlockResult {
	ret := _m.ctrl.Call(_m, "FetchBlocks", ctx, starts)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) FetchBlocks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocks", arg0, arg1)
}

func (_m *MockdatabaseBuffer) FetchBlocksMetadata(ctx context.Context, includeSizes bool, includeChecksums bool) []block.FetchBlockMetadataResult {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadata", ctx, includeSizes, includeChecksums)
	ret0, _ := ret[0].([]block.FetchBlockMetadataResult)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) FetchBlocksMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadata", arg0, arg1, arg2)
}

func (_m *MockdatabaseBuffer) IsEmpty() bool {
	ret := _m.ctrl.Call(_m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) IsEmpty() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsEmpty")
}

func (_m *MockdatabaseBuffer) NeedsDrain() bool {
	ret := _m.ctrl.Call(_m, "NeedsDrain")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockdatabaseBufferRecorder) NeedsDrain() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NeedsDrain")
}

func (_m *MockdatabaseBuffer) DrainAndReset(forced bool) {
	_m.ctrl.Call(_m, "DrainAndReset", forced)
}

func (_mr *_MockdatabaseBufferRecorder) DrainAndReset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DrainAndReset", arg0)
}

func (_m *MockdatabaseBuffer) Reset() {
	_m.ctrl.Call(_m, "Reset")
}

func (_mr *_MockdatabaseBufferRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}

// Mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *_MockOptionsRecorder
}

// Recorder for MockOptions (not exported)
type _MockOptionsRecorder struct {
	mock *MockOptions
}

func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &_MockOptionsRecorder{mock}
	return mock
}

func (_m *MockOptions) EXPECT() *_MockOptionsRecorder {
	return _m.recorder
}

func (_m *MockOptions) SetClockOptions(value clock.Options) Options {
	ret := _m.ctrl.Call(_m, "SetClockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetClockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetClockOptions", arg0)
}

func (_m *MockOptions) ClockOptions() clock.Options {
	ret := _m.ctrl.Call(_m, "ClockOptions")
	ret0, _ := ret[0].(clock.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) ClockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClockOptions")
}

func (_m *MockOptions) SetInstrumentOptions(value instrument.Options) Options {
	ret := _m.ctrl.Call(_m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetInstrumentOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetInstrumentOptions", arg0)
}

func (_m *MockOptions) InstrumentOptions() instrument.Options {
	ret := _m.ctrl.Call(_m, "InstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) InstrumentOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstrumentOptions")
}

func (_m *MockOptions) SetRetentionOptions(value retention.Options) Options {
	ret := _m.ctrl.Call(_m, "SetRetentionOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetRetentionOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRetentionOptions", arg0)
}

func (_m *MockOptions) RetentionOptions() retention.Options {
	ret := _m.ctrl.Call(_m, "RetentionOptions")
	ret0, _ := ret[0].(retention.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) RetentionOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetentionOptions")
}

func (_m *MockOptions) SetDatabaseBlockOptions(value block.Options) Options {
	ret := _m.ctrl.Call(_m, "SetDatabaseBlockOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetDatabaseBlockOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDatabaseBlockOptions", arg0)
}

func (_m *MockOptions) DatabaseBlockOptions() block.Options {
	ret := _m.ctrl.Call(_m, "DatabaseBlockOptions")
	ret0, _ := ret[0].(block.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockOptions")
}

func (_m *MockOptions) SetContextPool(value context.Pool) Options {
	ret := _m.ctrl.Call(_m, "SetContextPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetContextPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetContextPool", arg0)
}

func (_m *MockOptions) ContextPool() context.Pool {
	ret := _m.ctrl.Call(_m, "ContextPool")
	ret0, _ := ret[0].(context.Pool)
	return ret0
}

func (_mr *_MockOptionsRecorder) ContextPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContextPool")
}

func (_m *MockOptions) SetEncoderPool(value encoding.EncoderPool) Options {
	ret := _m.ctrl.Call(_m, "SetEncoderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetEncoderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetEncoderPool", arg0)
}

func (_m *MockOptions) EncoderPool() encoding.EncoderPool {
	ret := _m.ctrl.Call(_m, "EncoderPool")
	ret0, _ := ret[0].(encoding.EncoderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) EncoderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncoderPool")
}

func (_m *MockOptions) SetMultiReaderIteratorPool(value encoding.MultiReaderIteratorPool) Options {
	ret := _m.ctrl.Call(_m, "SetMultiReaderIteratorPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetMultiReaderIteratorPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMultiReaderIteratorPool", arg0)
}

func (_m *MockOptions) MultiReaderIteratorPool() encoding.MultiReaderIteratorPool {
	ret := _m.ctrl.Call(_m, "MultiReaderIteratorPool")
	ret0, _ := ret[0].(encoding.MultiReaderIteratorPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) MultiReaderIteratorPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MultiReaderIteratorPool")
}
