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

package commitlog

import (
	"time"

	"github.com/m3db/m3db/clock"
	"github.com/m3db/m3db/instrument"
	"github.com/m3db/m3db/persist/fs"
	"github.com/m3db/m3db/retention"
	"github.com/m3db/m3db/ts"
	xtime "github.com/m3db/m3x/time"
)

// Strategy describes the commit log writing strategy
type Strategy int

const (
	// StrategyWriteWait describes the strategy that waits
	// for the buffered commit log chunk that contains a write to flush
	// before acknowledging a write
	StrategyWriteWait Strategy = iota

	// StrategyWriteBehind describes the strategy that does not wait
	// for the buffered commit log chunk that contains a write to flush
	// before acknowledging a write
	StrategyWriteBehind
)

// CommitLog provides a synchronized commit log
type CommitLog interface {
	// Open the commit log
	Open() error

	// Write will write an entry in the commit log for a given series
	Write(
		series Series,
		datapoint ts.Datapoint,
		unit xtime.Unit,
		annotation ts.Annotation,
	) error

	// WriteBehind will write an entry in the commit log for a given series without waiting for completion
	WriteBehind(
		series Series,
		datapoint ts.Datapoint,
		unit xtime.Unit,
		annotation ts.Annotation,
	) error

	// Iter returns an iterator for accessing commit logs
	Iter() (Iterator, error)

	// Close the commit log
	Close() error
}

// Iterator provides an iterator for commit logs
type Iterator interface {
	// Next returns whether the iterator has the next value
	Next() bool

	// Current returns the current commit log entry
	Current() (Series, ts.Datapoint, xtime.Unit, ts.Annotation)

	// Err returns an error if an error occurred
	Err() error

	// Close the iterator
	Close()
}

// Series describes a series in the commit log
type Series struct {
	// UniqueIndex is the unique index assigned to this series
	UniqueIndex uint64

	// Namespace is the namespace the series belongs to
	Namespace ts.ID

	// ID is the series identifier
	ID ts.ID

	// Shard is the shard the series belongs to
	Shard uint32
}

// Options represents the options for the commit log
type Options interface {
	// SetClockOptions sets the clock options
	SetClockOptions(value clock.Options) Options

	// ClockOptions returns the clock options
	ClockOptions() clock.Options

	// SetInstrumentOptions sets the instrumentation options
	SetInstrumentOptions(value instrument.Options) Options

	// InstrumentOptions returns the instrumentation options
	InstrumentOptions() instrument.Options

	// SetRetentionOptions sets the retention options
	SetRetentionOptions(value retention.Options) Options

	// RetentionOptions returns the retention options
	RetentionOptions() retention.Options

	// SetFilesystemOptions sets the filesystem options
	SetFilesystemOptions(value fs.Options) Options

	// FilesystemOptions returns the filesystem options
	FilesystemOptions() fs.Options

	// SetFlushSize sets the flush size
	SetFlushSize(value int) Options

	// FlushSize returns the flush size
	FlushSize() int

	// SetStrategy sets the strategy
	SetStrategy(value Strategy) Options

	// Strategy returns the strategy
	Strategy() Strategy

	// SetFlushInterval sets the flush interval
	SetFlushInterval(value time.Duration) Options

	// FlushInterval returns the flush interval
	FlushInterval() time.Duration

	// SetBacklogQueueSize sets the backlog queue size
	SetBacklogQueueSize(value int) Options

	// BacklogQueueSize returns the backlog queue size
	BacklogQueueSize() int
}
