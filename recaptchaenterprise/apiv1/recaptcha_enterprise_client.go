// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package recaptchaenterprise

import (
	"context"
	"fmt"
	"math"
	"net/url"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateAssessment   []gax.CallOption
	AnnotateAssessment []gax.CallOption
	CreateKey          []gax.CallOption
	ListKeys           []gax.CallOption
	GetKey             []gax.CallOption
	UpdateKey          []gax.CallOption
	DeleteKey          []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("recaptchaenterprise.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateAssessment:   []gax.CallOption{},
		AnnotateAssessment: []gax.CallOption{},
		CreateKey:          []gax.CallOption{},
		ListKeys:           []gax.CallOption{},
		GetKey:             []gax.CallOption{},
		UpdateKey:          []gax.CallOption{},
		DeleteKey:          []gax.CallOption{},
	}
}

// Client is a client for interacting with reCAPTCHA Enterprise API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type Client struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	client recaptchaenterprisepb.RecaptchaEnterpriseServiceClient

	// The call options for this service.
	CallOptions *CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new recaptcha enterprise service client.
//
// Service to determine the likelihood an event is legitimate.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	connPool, err := gtransport.DialPool(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		connPool:    connPool,
		CallOptions: defaultCallOptions(),

		client: recaptchaenterprisepb.NewRecaptchaEnterpriseServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateAssessment creates an Assessment of the likelihood an event is legitimate.
func (c *Client) CreateAssessment(ctx context.Context, req *recaptchaenterprisepb.CreateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Assessment, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateAssessment[0:len(c.CallOptions.CreateAssessment):len(c.CallOptions.CreateAssessment)], opts...)
	var resp *recaptchaenterprisepb.Assessment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateAssessment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AnnotateAssessment annotates a previously created Assessment to provide additional information
// on whether the event turned out to be authentic or fradulent.
func (c *Client) AnnotateAssessment(ctx context.Context, req *recaptchaenterprisepb.AnnotateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.AnnotateAssessmentResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.AnnotateAssessment[0:len(c.CallOptions.AnnotateAssessment):len(c.CallOptions.AnnotateAssessment)], opts...)
	var resp *recaptchaenterprisepb.AnnotateAssessmentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.AnnotateAssessment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateKey creates a new reCAPTCHA Enterprise key.
func (c *Client) CreateKey(ctx context.Context, req *recaptchaenterprisepb.CreateKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateKey[0:len(c.CallOptions.CreateKey):len(c.CallOptions.CreateKey)], opts...)
	var resp *recaptchaenterprisepb.Key
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListKeys returns the list of all keys that belong to a project.
func (c *Client) ListKeys(ctx context.Context, req *recaptchaenterprisepb.ListKeysRequest, opts ...gax.CallOption) *KeyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListKeys[0:len(c.CallOptions.ListKeys):len(c.CallOptions.ListKeys)], opts...)
	it := &KeyIterator{}
	req = proto.Clone(req).(*recaptchaenterprisepb.ListKeysRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recaptchaenterprisepb.Key, string, error) {
		var resp *recaptchaenterprisepb.ListKeysResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListKeys(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.Keys, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// GetKey returns the specified key.
func (c *Client) GetKey(ctx context.Context, req *recaptchaenterprisepb.GetKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetKey[0:len(c.CallOptions.GetKey):len(c.CallOptions.GetKey)], opts...)
	var resp *recaptchaenterprisepb.Key
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateKey updates the specified key.
func (c *Client) UpdateKey(ctx context.Context, req *recaptchaenterprisepb.UpdateKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "key.name", url.QueryEscape(req.GetKey().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateKey[0:len(c.CallOptions.UpdateKey):len(c.CallOptions.UpdateKey)], opts...)
	var resp *recaptchaenterprisepb.Key
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteKey deletes the specified key.
func (c *Client) DeleteKey(ctx context.Context, req *recaptchaenterprisepb.DeleteKeyRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteKey[0:len(c.CallOptions.DeleteKey):len(c.CallOptions.DeleteKey)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// KeyIterator manages a stream of *recaptchaenterprisepb.Key.
type KeyIterator struct {
	items    []*recaptchaenterprisepb.Key
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*recaptchaenterprisepb.Key, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *KeyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *KeyIterator) Next() (*recaptchaenterprisepb.Key, error) {
	var item *recaptchaenterprisepb.Key
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *KeyIterator) bufLen() int {
	return len(it.items)
}

func (it *KeyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}