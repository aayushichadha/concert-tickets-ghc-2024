// Copyright 2023 Google LLC
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

package compute

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newZoneOperationsClientHook clientHook

// ZoneOperationsCallOptions contains the retry settings for each method of ZoneOperationsClient.
type ZoneOperationsCallOptions struct {
	Delete []gax.CallOption
	Get    []gax.CallOption
	List   []gax.CallOption
	Wait   []gax.CallOption
}

func defaultZoneOperationsRESTCallOptions() *ZoneOperationsCallOptions {
	return &ZoneOperationsCallOptions{
		Delete: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
		Get: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		List: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		Wait: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
		},
	}
}

// internalZoneOperationsClient is an interface that defines the methods available from Google Compute Engine API.
type internalZoneOperationsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteZoneOperationRequest, ...gax.CallOption) (*computepb.DeleteZoneOperationResponse, error)
	Get(context.Context, *computepb.GetZoneOperationRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListZoneOperationsRequest, ...gax.CallOption) *OperationIterator
	Wait(context.Context, *computepb.WaitZoneOperationRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// ZoneOperationsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The ZoneOperations API.
type ZoneOperationsClient struct {
	// The internal transport-dependent client.
	internalClient internalZoneOperationsClient

	// The call options for this service.
	CallOptions *ZoneOperationsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ZoneOperationsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ZoneOperationsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ZoneOperationsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified zone-specific Operations resource.
func (c *ZoneOperationsClient) Delete(ctx context.Context, req *computepb.DeleteZoneOperationRequest, opts ...gax.CallOption) (*computepb.DeleteZoneOperationResponse, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get retrieves the specified zone-specific Operations resource.
func (c *ZoneOperationsClient) Get(ctx context.Context, req *computepb.GetZoneOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// List retrieves a list of Operation resources contained within the specified zone.
func (c *ZoneOperationsClient) List(ctx context.Context, req *computepb.ListZoneOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Wait waits for the specified Operation resource to return as DONE or for the request to approach the 2 minute deadline, and retrieves the specified Operation resource. This method waits for no more than the 2 minutes and then returns the current state of the operation, which might be DONE or still in progress. This method is called on a best-effort basis. Specifically: - In uncommon cases, when the server is overloaded, the request might return before the default deadline is reached, or might return after zero seconds. - If the default deadline is reached, there is no guarantee that the operation is actually done when the method returns. Be prepared to retry if the operation is not DONE.
func (c *ZoneOperationsClient) Wait(ctx context.Context, req *computepb.WaitZoneOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Wait(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type zoneOperationsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing ZoneOperationsClient
	CallOptions **ZoneOperationsCallOptions
}

// NewZoneOperationsRESTClient creates a new zone operations rest client.
//
// The ZoneOperations API.
func NewZoneOperationsRESTClient(ctx context.Context, opts ...option.ClientOption) (*ZoneOperationsClient, error) {
	clientOpts := append(defaultZoneOperationsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultZoneOperationsRESTCallOptions()
	c := &zoneOperationsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &ZoneOperationsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultZoneOperationsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *zoneOperationsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *zoneOperationsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *zoneOperationsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Delete deletes the specified zone-specific Operations resource.
func (c *zoneOperationsRESTClient) Delete(ctx context.Context, req *computepb.DeleteZoneOperationRequest, opts ...gax.CallOption) (*computepb.DeleteZoneOperationResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/operations/%v", req.GetProject(), req.GetZone(), req.GetOperation())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "operation", url.QueryEscape(req.GetOperation()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.DeleteZoneOperationResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// Get retrieves the specified zone-specific Operations resource.
func (c *zoneOperationsRESTClient) Get(ctx context.Context, req *computepb.GetZoneOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/operations/%v", req.GetProject(), req.GetZone(), req.GetOperation())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "operation", url.QueryEscape(req.GetOperation()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// List retrieves a list of Operation resources contained within the specified zone.
func (c *zoneOperationsRESTClient) List(ctx context.Context, req *computepb.ListZoneOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	it := &OperationIterator{}
	req = proto.Clone(req).(*computepb.ListZoneOperationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.Operation, string, error) {
		resp := &computepb.OperationList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/operations", req.GetProject(), req.GetZone())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Wait waits for the specified Operation resource to return as DONE or for the request to approach the 2 minute deadline, and retrieves the specified Operation resource. This method waits for no more than the 2 minutes and then returns the current state of the operation, which might be DONE or still in progress. This method is called on a best-effort basis. Specifically: - In uncommon cases, when the server is overloaded, the request might return before the default deadline is reached, or might return after zero seconds. - If the default deadline is reached, there is no guarantee that the operation is actually done when the method returns. Be prepared to retry if the operation is not DONE.
func (c *zoneOperationsRESTClient) Wait(ctx context.Context, req *computepb.WaitZoneOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/operations/%v/wait", req.GetProject(), req.GetZone(), req.GetOperation())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "operation", url.QueryEscape(req.GetOperation()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Wait[0:len((*c.CallOptions).Wait):len((*c.CallOptions).Wait)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
