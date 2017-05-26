// Copyright (c) 2017 Uber Technologies, Inc.
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

// @generated Code generated by thrift-gen. Do not modify.

// Package rpc is generated code used to make or handle TChannel calls using Thrift.
package rpc

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"
)

// Interfaces for the service and client for the services defined in the IDL.

// TChanCluster is the interface that defines the server handler and client interface.
type TChanCluster interface {
	Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error)
	Health(ctx thrift.Context) (*HealthResult_, error)
	Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error)
	Write(ctx thrift.Context, req *WriteRequest) error
}

// TChanNode is the interface that defines the server handler and client interface.
type TChanNode interface {
	Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error)
	FetchBatchRaw(ctx thrift.Context, req *FetchBatchRawRequest) (*FetchBatchRawResult_, error)
	FetchBlocksMetadataRaw(ctx thrift.Context, req *FetchBlocksMetadataRawRequest) (*FetchBlocksMetadataRawResult_, error)
	FetchBlocksRaw(ctx thrift.Context, req *FetchBlocksRawRequest) (*FetchBlocksRawResult_, error)
	FetchMetadataBatchRaw(ctx thrift.Context, req *FetchMetadataBatchRawRequest) (*FetchMetadataBatchRawResult_, error)
	GetPersistRateLimit(ctx thrift.Context) (*NodePersistRateLimitResult_, error)
	GetWriteNewSeriesAsync(ctx thrift.Context) (*NodeWriteNewSeriesAsyncResult_, error)
	Health(ctx thrift.Context) (*NodeHealthResult_, error)
	Repair(ctx thrift.Context) error
	SetPersistRateLimit(ctx thrift.Context, req *NodeSetPersistRateLimitRequest) (*NodePersistRateLimitResult_, error)
	SetWriteNewSeriesAsync(ctx thrift.Context, req *NodeSetWriteNewSeriesAsyncRequest) (*NodeWriteNewSeriesAsyncResult_, error)
	Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error)
	Write(ctx thrift.Context, req *WriteRequest) error
	WriteBatchRaw(ctx thrift.Context, req *WriteBatchRawRequest) error
}

// Implementation of a client and service handler.

type tchanClusterClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanClusterInheritedClient(thriftService string, client thrift.TChanClient) *tchanClusterClient {
	return &tchanClusterClient{
		thriftService,
		client,
	}
}

// NewTChanClusterClient creates a client that can be used to make remote calls.
func NewTChanClusterClient(client thrift.TChanClient) TChanCluster {
	return NewTChanClusterInheritedClient("Cluster", client)
}

func (c *tchanClusterClient) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	var resp ClusterFetchResult
	args := ClusterFetchArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetch", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanClusterClient) Health(ctx thrift.Context) (*HealthResult_, error) {
	var resp ClusterHealthResult
	args := ClusterHealthArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "health", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanClusterClient) Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error) {
	var resp ClusterTruncateResult
	args := ClusterTruncateArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "truncate", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanClusterClient) Write(ctx thrift.Context, req *WriteRequest) error {
	var resp ClusterWriteResult
	args := ClusterWriteArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "write", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return err
}

type tchanClusterServer struct {
	handler TChanCluster
}

// NewTChanClusterServer wraps a handler for TChanCluster so it can be
// registered with a thrift.Server.
func NewTChanClusterServer(handler TChanCluster) thrift.TChanServer {
	return &tchanClusterServer{
		handler,
	}
}

func (s *tchanClusterServer) Service() string {
	return "Cluster"
}

func (s *tchanClusterServer) Methods() []string {
	return []string{
		"fetch",
		"health",
		"truncate",
		"write",
	}
}

func (s *tchanClusterServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "fetch":
		return s.handleFetch(ctx, protocol)
	case "health":
		return s.handleHealth(ctx, protocol)
	case "truncate":
		return s.handleTruncate(ctx, protocol)
	case "write":
		return s.handleWrite(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanClusterServer) handleFetch(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterFetchArgs
	var res ClusterFetchResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Fetch(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanClusterServer) handleHealth(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterHealthArgs
	var res ClusterHealthResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Health(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanClusterServer) handleTruncate(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterTruncateArgs
	var res ClusterTruncateResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Truncate(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanClusterServer) handleWrite(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterWriteArgs
	var res ClusterWriteResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.Write(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

type tchanNodeClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanNodeInheritedClient(thriftService string, client thrift.TChanClient) *tchanNodeClient {
	return &tchanNodeClient{
		thriftService,
		client,
	}
}

// NewTChanNodeClient creates a client that can be used to make remote calls.
func NewTChanNodeClient(client thrift.TChanClient) TChanNode {
	return NewTChanNodeInheritedClient("Node", client)
}

func (c *tchanNodeClient) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	var resp NodeFetchResult
	args := NodeFetchArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetch", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchBatchRaw(ctx thrift.Context, req *FetchBatchRawRequest) (*FetchBatchRawResult_, error) {
	var resp NodeFetchBatchRawResult
	args := NodeFetchBatchRawArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchBatchRaw", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchBlocksMetadataRaw(ctx thrift.Context, req *FetchBlocksMetadataRawRequest) (*FetchBlocksMetadataRawResult_, error) {
	var resp NodeFetchBlocksMetadataRawResult
	args := NodeFetchBlocksMetadataRawArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchBlocksMetadataRaw", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchBlocksRaw(ctx thrift.Context, req *FetchBlocksRawRequest) (*FetchBlocksRawResult_, error) {
	var resp NodeFetchBlocksRawResult
	args := NodeFetchBlocksRawArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchBlocksRaw", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchMetadataBatchRaw(ctx thrift.Context, req *FetchMetadataBatchRawRequest) (*FetchMetadataBatchRawResult_, error) {
	var resp NodeFetchMetadataBatchRawResult
	args := NodeFetchMetadataBatchRawArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchMetadataBatchRaw", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) GetPersistRateLimit(ctx thrift.Context) (*NodePersistRateLimitResult_, error) {
	var resp NodeGetPersistRateLimitResult
	args := NodeGetPersistRateLimitArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "getPersistRateLimit", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) GetWriteNewSeriesAsync(ctx thrift.Context) (*NodeWriteNewSeriesAsyncResult_, error) {
	var resp NodeGetWriteNewSeriesAsyncResult
	args := NodeGetWriteNewSeriesAsyncArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "getWriteNewSeriesAsync", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Health(ctx thrift.Context) (*NodeHealthResult_, error) {
	var resp NodeHealthResult
	args := NodeHealthArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "health", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Repair(ctx thrift.Context) error {
	var resp NodeRepairResult
	args := NodeRepairArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "repair", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return err
}

func (c *tchanNodeClient) SetPersistRateLimit(ctx thrift.Context, req *NodeSetPersistRateLimitRequest) (*NodePersistRateLimitResult_, error) {
	var resp NodeSetPersistRateLimitResult
	args := NodeSetPersistRateLimitArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "setPersistRateLimit", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) SetWriteNewSeriesAsync(ctx thrift.Context, req *NodeSetWriteNewSeriesAsyncRequest) (*NodeWriteNewSeriesAsyncResult_, error) {
	var resp NodeSetWriteNewSeriesAsyncResult
	args := NodeSetWriteNewSeriesAsyncArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "setWriteNewSeriesAsync", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error) {
	var resp NodeTruncateResult
	args := NodeTruncateArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "truncate", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Write(ctx thrift.Context, req *WriteRequest) error {
	var resp NodeWriteResult
	args := NodeWriteArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "write", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return err
}

func (c *tchanNodeClient) WriteBatchRaw(ctx thrift.Context, req *WriteBatchRawRequest) error {
	var resp NodeWriteBatchRawResult
	args := NodeWriteBatchRawArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "writeBatchRaw", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return err
}

type tchanNodeServer struct {
	handler TChanNode
}

// NewTChanNodeServer wraps a handler for TChanNode so it can be
// registered with a thrift.Server.
func NewTChanNodeServer(handler TChanNode) thrift.TChanServer {
	return &tchanNodeServer{
		handler,
	}
}

func (s *tchanNodeServer) Service() string {
	return "Node"
}

func (s *tchanNodeServer) Methods() []string {
	return []string{
		"fetch",
		"fetchBatchRaw",
		"fetchBlocksMetadataRaw",
		"fetchBlocksRaw",
		"fetchMetadataBatchRaw",
		"getPersistRateLimit",
		"getWriteNewSeriesAsync",
		"health",
		"repair",
		"setPersistRateLimit",
		"setWriteNewSeriesAsync",
		"truncate",
		"write",
		"writeBatchRaw",
	}
}

func (s *tchanNodeServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "fetch":
		return s.handleFetch(ctx, protocol)
	case "fetchBatchRaw":
		return s.handleFetchBatchRaw(ctx, protocol)
	case "fetchBlocksMetadataRaw":
		return s.handleFetchBlocksMetadataRaw(ctx, protocol)
	case "fetchBlocksRaw":
		return s.handleFetchBlocksRaw(ctx, protocol)
	case "fetchMetadataBatchRaw":
		return s.handleFetchMetadataBatchRaw(ctx, protocol)
	case "getPersistRateLimit":
		return s.handleGetPersistRateLimit(ctx, protocol)
	case "getWriteNewSeriesAsync":
		return s.handleGetWriteNewSeriesAsync(ctx, protocol)
	case "health":
		return s.handleHealth(ctx, protocol)
	case "repair":
		return s.handleRepair(ctx, protocol)
	case "setPersistRateLimit":
		return s.handleSetPersistRateLimit(ctx, protocol)
	case "setWriteNewSeriesAsync":
		return s.handleSetWriteNewSeriesAsync(ctx, protocol)
	case "truncate":
		return s.handleTruncate(ctx, protocol)
	case "write":
		return s.handleWrite(ctx, protocol)
	case "writeBatchRaw":
		return s.handleWriteBatchRaw(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanNodeServer) handleFetch(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchArgs
	var res NodeFetchResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Fetch(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleFetchBatchRaw(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchBatchRawArgs
	var res NodeFetchBatchRawResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchBatchRaw(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleFetchBlocksMetadataRaw(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchBlocksMetadataRawArgs
	var res NodeFetchBlocksMetadataRawResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchBlocksMetadataRaw(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleFetchBlocksRaw(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchBlocksRawArgs
	var res NodeFetchBlocksRawResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchBlocksRaw(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleFetchMetadataBatchRaw(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchMetadataBatchRawArgs
	var res NodeFetchMetadataBatchRawResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchMetadataBatchRaw(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleGetPersistRateLimit(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeGetPersistRateLimitArgs
	var res NodeGetPersistRateLimitResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetPersistRateLimit(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleGetWriteNewSeriesAsync(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeGetWriteNewSeriesAsyncArgs
	var res NodeGetWriteNewSeriesAsyncResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetWriteNewSeriesAsync(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleHealth(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeHealthArgs
	var res NodeHealthResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Health(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleRepair(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeRepairArgs
	var res NodeRepairResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.Repair(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleSetPersistRateLimit(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeSetPersistRateLimitArgs
	var res NodeSetPersistRateLimitResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.SetPersistRateLimit(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleSetWriteNewSeriesAsync(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeSetWriteNewSeriesAsyncArgs
	var res NodeSetWriteNewSeriesAsyncResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.SetWriteNewSeriesAsync(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleTruncate(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeTruncateArgs
	var res NodeTruncateResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Truncate(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleWrite(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeWriteArgs
	var res NodeWriteResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.Write(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleWriteBatchRaw(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeWriteBatchRawArgs
	var res NodeWriteBatchRawResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.WriteBatchRaw(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *WriteBatchRawErrors:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *WriteBatchRawErrors but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}
