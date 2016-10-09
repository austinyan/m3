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
// Source: github.com/m3db/m3db/storage/block/types.go

package block

import (
	gomock "github.com/golang/mock/gomock"
	context "github.com/m3db/m3db/context"
	encoding "github.com/m3db/m3db/encoding"
	pool "github.com/m3db/m3db/pool"
	ts "github.com/m3db/m3db/ts"
	io "github.com/m3db/m3db/x/io"
	time "github.com/m3db/m3x/time"
	time0 "time"
)

// Mock of Metadata interface
type MockMetadata struct {
	ctrl     *gomock.Controller
	recorder *_MockMetadataRecorder
}

// Recorder for MockMetadata (not exported)
type _MockMetadataRecorder struct {
	mock *MockMetadata
}

func NewMockMetadata(ctrl *gomock.Controller) *MockMetadata {
	mock := &MockMetadata{ctrl: ctrl}
	mock.recorder = &_MockMetadataRecorder{mock}
	return mock
}

func (_m *MockMetadata) EXPECT() *_MockMetadataRecorder {
	return _m.recorder
}

func (_m *MockMetadata) Start() time0.Time {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockMetadataRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockMetadata) Size() int64 {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockMetadataRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Size")
}

func (_m *MockMetadata) Checksum() *uint32 {
	ret := _m.ctrl.Call(_m, "Checksum")
	ret0, _ := ret[0].(*uint32)
	return ret0
}

func (_mr *_MockMetadataRecorder) Checksum() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Checksum")
}

// Mock of BlocksMetadata interface
type MockBlocksMetadata struct {
	ctrl     *gomock.Controller
	recorder *_MockBlocksMetadataRecorder
}

// Recorder for MockBlocksMetadata (not exported)
type _MockBlocksMetadataRecorder struct {
	mock *MockBlocksMetadata
}

func NewMockBlocksMetadata(ctrl *gomock.Controller) *MockBlocksMetadata {
	mock := &MockBlocksMetadata{ctrl: ctrl}
	mock.recorder = &_MockBlocksMetadataRecorder{mock}
	return mock
}

func (_m *MockBlocksMetadata) EXPECT() *_MockBlocksMetadataRecorder {
	return _m.recorder
}

func (_m *MockBlocksMetadata) ID() ts.ID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(ts.ID)
	return ret0
}

func (_mr *_MockBlocksMetadataRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockBlocksMetadata) Blocks() []Metadata {
	ret := _m.ctrl.Call(_m, "Blocks")
	ret0, _ := ret[0].([]Metadata)
	return ret0
}

func (_mr *_MockBlocksMetadataRecorder) Blocks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Blocks")
}

// Mock of FetchBlockMetadataResult interface
type MockFetchBlockMetadataResult struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlockMetadataResultRecorder
}

// Recorder for MockFetchBlockMetadataResult (not exported)
type _MockFetchBlockMetadataResultRecorder struct {
	mock *MockFetchBlockMetadataResult
}

func NewMockFetchBlockMetadataResult(ctrl *gomock.Controller) *MockFetchBlockMetadataResult {
	mock := &MockFetchBlockMetadataResult{ctrl: ctrl}
	mock.recorder = &_MockFetchBlockMetadataResultRecorder{mock}
	return mock
}

func (_m *MockFetchBlockMetadataResult) EXPECT() *_MockFetchBlockMetadataResultRecorder {
	return _m.recorder
}

func (_m *MockFetchBlockMetadataResult) Start() time0.Time {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockFetchBlockMetadataResultRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockFetchBlockMetadataResult) Size() *int64 {
	ret := _m.ctrl.Call(_m, "Size")
	ret0, _ := ret[0].(*int64)
	return ret0
}

func (_mr *_MockFetchBlockMetadataResultRecorder) Size() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Size")
}

func (_m *MockFetchBlockMetadataResult) Checksum() *uint32 {
	ret := _m.ctrl.Call(_m, "Checksum")
	ret0, _ := ret[0].(*uint32)
	return ret0
}

func (_mr *_MockFetchBlockMetadataResultRecorder) Checksum() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Checksum")
}

func (_m *MockFetchBlockMetadataResult) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFetchBlockMetadataResultRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

// Mock of FetchBlocksMetadataResult interface
type MockFetchBlocksMetadataResult struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlocksMetadataResultRecorder
}

// Recorder for MockFetchBlocksMetadataResult (not exported)
type _MockFetchBlocksMetadataResultRecorder struct {
	mock *MockFetchBlocksMetadataResult
}

func NewMockFetchBlocksMetadataResult(ctrl *gomock.Controller) *MockFetchBlocksMetadataResult {
	mock := &MockFetchBlocksMetadataResult{ctrl: ctrl}
	mock.recorder = &_MockFetchBlocksMetadataResultRecorder{mock}
	return mock
}

func (_m *MockFetchBlocksMetadataResult) EXPECT() *_MockFetchBlocksMetadataResultRecorder {
	return _m.recorder
}

func (_m *MockFetchBlocksMetadataResult) ID() ts.ID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(ts.ID)
	return ret0
}

func (_mr *_MockFetchBlocksMetadataResultRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockFetchBlocksMetadataResult) Blocks() []FetchBlockMetadataResult {
	ret := _m.ctrl.Call(_m, "Blocks")
	ret0, _ := ret[0].([]FetchBlockMetadataResult)
	return ret0
}

func (_mr *_MockFetchBlocksMetadataResultRecorder) Blocks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Blocks")
}

// Mock of FilteredBlocksMetadataIter interface
type MockFilteredBlocksMetadataIter struct {
	ctrl     *gomock.Controller
	recorder *_MockFilteredBlocksMetadataIterRecorder
}

// Recorder for MockFilteredBlocksMetadataIter (not exported)
type _MockFilteredBlocksMetadataIterRecorder struct {
	mock *MockFilteredBlocksMetadataIter
}

func NewMockFilteredBlocksMetadataIter(ctrl *gomock.Controller) *MockFilteredBlocksMetadataIter {
	mock := &MockFilteredBlocksMetadataIter{ctrl: ctrl}
	mock.recorder = &_MockFilteredBlocksMetadataIterRecorder{mock}
	return mock
}

func (_m *MockFilteredBlocksMetadataIter) EXPECT() *_MockFilteredBlocksMetadataIterRecorder {
	return _m.recorder
}

func (_m *MockFilteredBlocksMetadataIter) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockFilteredBlocksMetadataIterRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Next")
}

func (_m *MockFilteredBlocksMetadataIter) Current() (ts.ID, Metadata) {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(ts.ID)
	ret1, _ := ret[1].(Metadata)
	return ret0, ret1
}

func (_mr *_MockFilteredBlocksMetadataIterRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Current")
}

// Mock of FetchBlockResult interface
type MockFetchBlockResult struct {
	ctrl     *gomock.Controller
	recorder *_MockFetchBlockResultRecorder
}

// Recorder for MockFetchBlockResult (not exported)
type _MockFetchBlockResultRecorder struct {
	mock *MockFetchBlockResult
}

func NewMockFetchBlockResult(ctrl *gomock.Controller) *MockFetchBlockResult {
	mock := &MockFetchBlockResult{ctrl: ctrl}
	mock.recorder = &_MockFetchBlockResultRecorder{mock}
	return mock
}

func (_m *MockFetchBlockResult) EXPECT() *_MockFetchBlockResultRecorder {
	return _m.recorder
}

func (_m *MockFetchBlockResult) Start() time0.Time {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockFetchBlockResultRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockFetchBlockResult) Readers() []io.SegmentReader {
	ret := _m.ctrl.Call(_m, "Readers")
	ret0, _ := ret[0].([]io.SegmentReader)
	return ret0
}

func (_mr *_MockFetchBlockResultRecorder) Readers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Readers")
}

func (_m *MockFetchBlockResult) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFetchBlockResultRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

// Mock of DatabaseBlock interface
type MockDatabaseBlock struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseBlockRecorder
}

// Recorder for MockDatabaseBlock (not exported)
type _MockDatabaseBlockRecorder struct {
	mock *MockDatabaseBlock
}

func NewMockDatabaseBlock(ctrl *gomock.Controller) *MockDatabaseBlock {
	mock := &MockDatabaseBlock{ctrl: ctrl}
	mock.recorder = &_MockDatabaseBlockRecorder{mock}
	return mock
}

func (_m *MockDatabaseBlock) EXPECT() *_MockDatabaseBlockRecorder {
	return _m.recorder
}

func (_m *MockDatabaseBlock) IsSealed() bool {
	ret := _m.ctrl.Call(_m, "IsSealed")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDatabaseBlockRecorder) IsSealed() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsSealed")
}

func (_m *MockDatabaseBlock) StartTime() time0.Time {
	ret := _m.ctrl.Call(_m, "StartTime")
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockDatabaseBlockRecorder) StartTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartTime")
}

func (_m *MockDatabaseBlock) Checksum() *uint32 {
	ret := _m.ctrl.Call(_m, "Checksum")
	ret0, _ := ret[0].(*uint32)
	return ret0
}

func (_mr *_MockDatabaseBlockRecorder) Checksum() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Checksum")
}

func (_m *MockDatabaseBlock) Write(timestamp time0.Time, value float64, unit time.Unit, annotation ts.Annotation) error {
	ret := _m.ctrl.Call(_m, "Write", timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseBlockRecorder) Write(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3)
}

func (_m *MockDatabaseBlock) Stream(blocker context.Context) (io.SegmentReader, error) {
	ret := _m.ctrl.Call(_m, "Stream", blocker)
	ret0, _ := ret[0].(io.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseBlockRecorder) Stream(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stream", arg0)
}

func (_m *MockDatabaseBlock) Reset(startTime time0.Time, encoder encoding.Encoder) {
	_m.ctrl.Call(_m, "Reset", startTime, encoder)
}

func (_mr *_MockDatabaseBlockRecorder) Reset(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0, arg1)
}

func (_m *MockDatabaseBlock) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockDatabaseBlockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockDatabaseBlock) Seal() {
	_m.ctrl.Call(_m, "Seal")
}

func (_mr *_MockDatabaseBlockRecorder) Seal() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Seal")
}

// Mock of DatabaseSeriesBlocks interface
type MockDatabaseSeriesBlocks struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseSeriesBlocksRecorder
}

// Recorder for MockDatabaseSeriesBlocks (not exported)
type _MockDatabaseSeriesBlocksRecorder struct {
	mock *MockDatabaseSeriesBlocks
}

func NewMockDatabaseSeriesBlocks(ctrl *gomock.Controller) *MockDatabaseSeriesBlocks {
	mock := &MockDatabaseSeriesBlocks{ctrl: ctrl}
	mock.recorder = &_MockDatabaseSeriesBlocksRecorder{mock}
	return mock
}

func (_m *MockDatabaseSeriesBlocks) EXPECT() *_MockDatabaseSeriesBlocksRecorder {
	return _m.recorder
}

func (_m *MockDatabaseSeriesBlocks) Options() Options {
	ret := _m.ctrl.Call(_m, "Options")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) Options() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Options")
}

func (_m *MockDatabaseSeriesBlocks) Len() int {
	ret := _m.ctrl.Call(_m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) Len() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Len")
}

func (_m *MockDatabaseSeriesBlocks) AddBlock(block DatabaseBlock) {
	_m.ctrl.Call(_m, "AddBlock", block)
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) AddBlock(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddBlock", arg0)
}

func (_m *MockDatabaseSeriesBlocks) AddSeries(other DatabaseSeriesBlocks) {
	_m.ctrl.Call(_m, "AddSeries", other)
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) AddSeries(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddSeries", arg0)
}

func (_m *MockDatabaseSeriesBlocks) MinTime() time0.Time {
	ret := _m.ctrl.Call(_m, "MinTime")
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) MinTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MinTime")
}

func (_m *MockDatabaseSeriesBlocks) MaxTime() time0.Time {
	ret := _m.ctrl.Call(_m, "MaxTime")
	ret0, _ := ret[0].(time0.Time)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) MaxTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MaxTime")
}

func (_m *MockDatabaseSeriesBlocks) BlockAt(t time0.Time) (DatabaseBlock, bool) {
	ret := _m.ctrl.Call(_m, "BlockAt", t)
	ret0, _ := ret[0].(DatabaseBlock)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) BlockAt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BlockAt", arg0)
}

func (_m *MockDatabaseSeriesBlocks) BlockOrAdd(t time0.Time) DatabaseBlock {
	ret := _m.ctrl.Call(_m, "BlockOrAdd", t)
	ret0, _ := ret[0].(DatabaseBlock)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) BlockOrAdd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BlockOrAdd", arg0)
}

func (_m *MockDatabaseSeriesBlocks) AllBlocks() map[time0.Time]DatabaseBlock {
	ret := _m.ctrl.Call(_m, "AllBlocks")
	ret0, _ := ret[0].(map[time0.Time]DatabaseBlock)
	return ret0
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) AllBlocks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AllBlocks")
}

func (_m *MockDatabaseSeriesBlocks) RemoveBlockAt(t time0.Time) {
	_m.ctrl.Call(_m, "RemoveBlockAt", t)
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) RemoveBlockAt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveBlockAt", arg0)
}

func (_m *MockDatabaseSeriesBlocks) RemoveAll() {
	_m.ctrl.Call(_m, "RemoveAll")
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) RemoveAll() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveAll")
}

func (_m *MockDatabaseSeriesBlocks) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockDatabaseSeriesBlocksRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of DatabaseBlockPool interface
type MockDatabaseBlockPool struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseBlockPoolRecorder
}

// Recorder for MockDatabaseBlockPool (not exported)
type _MockDatabaseBlockPoolRecorder struct {
	mock *MockDatabaseBlockPool
}

func NewMockDatabaseBlockPool(ctrl *gomock.Controller) *MockDatabaseBlockPool {
	mock := &MockDatabaseBlockPool{ctrl: ctrl}
	mock.recorder = &_MockDatabaseBlockPoolRecorder{mock}
	return mock
}

func (_m *MockDatabaseBlockPool) EXPECT() *_MockDatabaseBlockPoolRecorder {
	return _m.recorder
}

func (_m *MockDatabaseBlockPool) Init(alloc DatabaseBlockAllocate) {
	_m.ctrl.Call(_m, "Init", alloc)
}

func (_mr *_MockDatabaseBlockPoolRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

func (_m *MockDatabaseBlockPool) Get() DatabaseBlock {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(DatabaseBlock)
	return ret0
}

func (_mr *_MockDatabaseBlockPoolRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockDatabaseBlockPool) Put(block DatabaseBlock) {
	_m.ctrl.Call(_m, "Put", block)
}

func (_mr *_MockDatabaseBlockPoolRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
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

func (_m *MockOptions) SetDatabaseBlockAllocSize(value int) Options {
	ret := _m.ctrl.Call(_m, "SetDatabaseBlockAllocSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetDatabaseBlockAllocSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDatabaseBlockAllocSize", arg0)
}

func (_m *MockOptions) DatabaseBlockAllocSize() int {
	ret := _m.ctrl.Call(_m, "DatabaseBlockAllocSize")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockAllocSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockAllocSize")
}

func (_m *MockOptions) SetDatabaseBlockPool(value DatabaseBlockPool) Options {
	ret := _m.ctrl.Call(_m, "SetDatabaseBlockPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetDatabaseBlockPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDatabaseBlockPool", arg0)
}

func (_m *MockOptions) DatabaseBlockPool() DatabaseBlockPool {
	ret := _m.ctrl.Call(_m, "DatabaseBlockPool")
	ret0, _ := ret[0].(DatabaseBlockPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) DatabaseBlockPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseBlockPool")
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

func (_m *MockOptions) SetReaderIteratorPool(value encoding.ReaderIteratorPool) Options {
	ret := _m.ctrl.Call(_m, "SetReaderIteratorPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetReaderIteratorPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetReaderIteratorPool", arg0)
}

func (_m *MockOptions) ReaderIteratorPool() encoding.ReaderIteratorPool {
	ret := _m.ctrl.Call(_m, "ReaderIteratorPool")
	ret0, _ := ret[0].(encoding.ReaderIteratorPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) ReaderIteratorPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReaderIteratorPool")
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

func (_m *MockOptions) SetSegmentReaderPool(value io.SegmentReaderPool) Options {
	ret := _m.ctrl.Call(_m, "SetSegmentReaderPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetSegmentReaderPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSegmentReaderPool", arg0)
}

func (_m *MockOptions) SegmentReaderPool() io.SegmentReaderPool {
	ret := _m.ctrl.Call(_m, "SegmentReaderPool")
	ret0, _ := ret[0].(io.SegmentReaderPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) SegmentReaderPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SegmentReaderPool")
}

func (_m *MockOptions) SetBytesPool(value pool.BytesPool) Options {
	ret := _m.ctrl.Call(_m, "SetBytesPool", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetBytesPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetBytesPool", arg0)
}

func (_m *MockOptions) BytesPool() pool.BytesPool {
	ret := _m.ctrl.Call(_m, "BytesPool")
	ret0, _ := ret[0].(pool.BytesPool)
	return ret0
}

func (_mr *_MockOptionsRecorder) BytesPool() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BytesPool")
}
